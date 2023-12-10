package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Misoten-B/airship-backend/config"
	"github.com/Misoten-B/airship-backend/internal/drivers"
	"github.com/Misoten-B/airship-backend/internal/drivers/database"
	"github.com/Misoten-B/airship-backend/internal/drivers/database/model"
	"github.com/Misoten-B/airship-backend/internal/file"
	"github.com/Misoten-B/airship-backend/internal/frameworks"
	"github.com/Misoten-B/airship-backend/internal/frameworks/handler/business_card_background/dto"
	"github.com/Misoten-B/airship-backend/internal/id"
	"github.com/gin-gonic/gin"
)

const (
	containerName = "background-images"
)

// @Tags BusinessCardBackground
// @Router /v1/users/business_card_backgrounds [POST]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Param BusinessCardBackgroundImage formData file true "Image file to be uploaded"
// @Param dto.CreateBackgroundRequest formData dto.CreateBackgroundRequest true "BusinessCardBackground"
// @Success 201 {object} dto.BackgroundResponse
func CreateBusinessCardBackground(c *gin.Context) {
	log.Printf("Authorization: %s", c.GetHeader("Authorization"))

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
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bcbID, err := id.NewID()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ab := drivers.NewAzureBlobDriver(config.GetConfig())
	castedFile := file.NewMyFile(formFile, fileHeader)
	castedFile.FileHeader().Filename = fmt.Sprintf("%s.png", bcbID.String())
	if err = ab.SaveBlob(containerName, *castedFile); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	bcb := model.BusinessCardBackground{
		ID:        bcbID.String(),
		ColorCode: request.BusinessCardBackgroundColor,
		ImagePath: castedFile.FileHeader().Filename,
	}

	pbcb := model.PersonalBusinessCardBackground{
		UserID: uid,
		ID:     bcb.ID,
	}

	db, err := database.ConnectDB()
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
	log.Printf("Authorization: %s", c.GetHeader("Authorization"))

	uid, err := frameworks.GetUID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	db, err := database.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// TODO: どうにかする
	bcbst := []model.BusinessCardBackground{}
	if err = db.Joins("JOIN business_card_background_templates on " +
		"business_card_backgrounds.id = " +
		"business_card_background_templates.id").
		Find(&bcbst).
		Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	pbcbs := []model.BusinessCardBackground{}
	if err = db.Joins("JOIN personal_business_card_backgrounds on "+
		"personal_business_card_backgrounds.id = "+
		"business_card_backgrounds.id").
		Where("personal_business_card_backgrounds.user_id = ?", uid).
		Find(&pbcbs).
		Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ab := drivers.NewAzureBlobDriver(config.GetConfig())
	containerURL, err := ab.GetContainerURL(containerName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	responses := []dto.BackgroundResponse{}

	for _, bcb := range bcbst {
		responses = append(responses, dto.BackgroundResponse{
			ID:                          bcb.ID,
			BusinessCardBackgroundColor: bcb.ColorCode,
			BusinessCardBackgroundImage: containerURL.Path(bcb.ImagePath),
		})
	}

	for _, bcb := range pbcbs {
		responses = append(responses, dto.BackgroundResponse{
			ID:                          bcb.ID,
			BusinessCardBackgroundColor: bcb.ColorCode,
			BusinessCardBackgroundImage: containerURL.Path(bcb.ImagePath),
		})
	}

	c.JSON(http.StatusOK, responses)
}
