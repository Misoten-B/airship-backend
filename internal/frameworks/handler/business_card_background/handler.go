//nolint:cyclop // 時間がないため一旦
package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Misoten-B/airship-backend/internal/domain/shared"
	"github.com/Misoten-B/airship-backend/internal/drivers"
	"github.com/Misoten-B/airship-backend/internal/drivers/config"
	"github.com/Misoten-B/airship-backend/internal/drivers/database/model"
	"github.com/Misoten-B/airship-backend/internal/file"
	"github.com/Misoten-B/airship-backend/internal/frameworks"
	"github.com/Misoten-B/airship-backend/internal/frameworks/handler/business_card_background/dto"
	"github.com/gin-gonic/gin"
)

const (
	containerName = "background-images"
)

// @Tags BusinessCardBackground
// @Router /v1/users/business_card_backgrounds [POST]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Accept multipart/form-data
// @Param BusinessCardBackgroundImage formData file false "Image file to be uploaded"
// @Param dto.CreateBackgroundRequest formData dto.CreateBackgroundRequest true "BusinessCardBackground"
// @Success 201 {object} dto.BackgroundResponse
func CreateBusinessCardBackground(c *gin.Context) {
	uid, err := frameworks.GetUID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	request := dto.CreateBackgroundRequest{}
	if err = c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: pngでバリデーション
	// TODO: 余裕があれば解像度を低くする
	formFile, fileHeader, err := c.Request.FormFile("BusinessCardBackgroundImage")
	if err != nil && !errors.Is(err, http.ErrMissingFile) {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bcbID, err := shared.NewID()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var fileName string

	if fileHeader != nil {
		var appConfig *config.Config
		appConfig, err = config.GetConfig()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ab := drivers.NewAzureBlobDriver(appConfig)
		castedFile := file.NewMyFile(formFile, fileHeader)

		fileName = fmt.Sprintf("%s.png", bcbID.String())
		castedFile.FileHeader().Filename = fileName

		if err = ab.SaveBlob(containerName, castedFile); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	bcb := model.BusinessCardBackground{
		ID:        bcbID.String(),
		ColorCode: request.BusinessCardBackgroundColor,
		ImagePath: fileName,
	}

	pbcb := model.PersonalBusinessCardBackground{
		UserID: uid,
		ID:     bcb.ID,
	}

	db, err := frameworks.GetDB(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err = db.Create(&bcb).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err = db.Create(&pbcb).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Header("Location", fmt.Sprintf("/%s", uid))
	c.JSON(http.StatusCreated, dto.BackgroundResponse{
		ID:                          bcb.ID,
		BusinessCardBackgroundColor: bcb.ColorCode,
		BusinessCardBackgroundImage: bcb.ImagePath,
	})
}

// @Tags BusinessCardBackground
// @Router /v1/users/business_card_backgrounds [GET]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Success 200 {object} []dto.BackgroundResponse
func ReadAllBusinessCardBackground(c *gin.Context) {
	uid, err := frameworks.GetUID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	db, err := frameworks.GetDB(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	bcbst := []model.BusinessCardBackground{}
	if err = db.Joins("LEFT JOIN business_card_background_templates on "+
		"business_card_backgrounds.id = business_card_background_templates.id").
		Joins("LEFT JOIN personal_business_card_backgrounds on "+
			"personal_business_card_backgrounds.id = business_card_backgrounds.id").
		Where("personal_business_card_backgrounds.user_id = ? OR "+
			"business_card_backgrounds.id = business_card_background_templates.id", uid).
		Order("id desc").
		Find(&bcbst).
		Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	appConfig, err := config.GetConfig()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ab := drivers.NewAzureBlobDriver(appConfig)
	containerURL, err := ab.GetContainerURL(containerName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	responses := []dto.BackgroundResponse{}

	for _, bcb := range bcbst {
		backgroundImageURL := ""
		if bcb.ImagePath != "" {
			backgroundImageURL = containerURL.Path(bcb.ImagePath)
		}
		responses = append(responses, dto.BackgroundResponse{
			ID:                          bcb.ID,
			BusinessCardBackgroundColor: bcb.ColorCode,
			BusinessCardBackgroundImage: backgroundImageURL,
		})
	}

	c.JSON(http.StatusOK, responses)
}

// @Tags BusinessCardBackground
// @Router /v1/users/business_card_backgrounds/{id} [Delete]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Success 204 {object} nil
func DeleteBusinessCardBackground(c *gin.Context) {
	uid, err := frameworks.GetUID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	bcbID := c.Param("id")

	db, err := frameworks.GetDB(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	bcb := model.BusinessCardBackground{}
	if err = db.First(&bcb, "id = ?", bcbID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// TODO: 権限確認

	if bcb.ImagePath != "" {
		var appConfig *config.Config
		appConfig, err = config.GetConfig()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ab := drivers.NewAzureBlobDriver(appConfig)
		if err = ab.DeleteBlob(containerName, bcb.ImagePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err = tx.Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 外部キー制約は400に

	if err = tx.Delete(&model.BusinessCard{}, "business_card_background_id = ?", bcbID).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err = tx.Delete(&model.PersonalBusinessCardBackground{}, "id = ? AND user_id = ?", bcbID, uid).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err = tx.Delete(&model.BusinessCardBackground{}, "id = ?", bcbID).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err = tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
