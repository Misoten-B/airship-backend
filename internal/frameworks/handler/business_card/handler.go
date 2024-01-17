//nolint:funlen,gocognit,cyclop // 時間がないため一旦
package handler

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/Misoten-B/airship-backend/internal/domain/shared"
	"github.com/Misoten-B/airship-backend/internal/drivers"
	"github.com/Misoten-B/airship-backend/internal/drivers/config"
	"github.com/Misoten-B/airship-backend/internal/drivers/database/model"
	"github.com/Misoten-B/airship-backend/internal/frameworks"
	"github.com/Misoten-B/airship-backend/internal/frameworks/handler/business_card/dto"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	backgroundContainer            = "background-images"
	qrcodeContainer                = "qrcode-images"
	threeDimentionalModelContainer = "three-dimentional-models"
	audioContainer                 = "voice-sounds"
)

// @Tags BusinessCard
// @Router /v1/users/business_cards [POST]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Param CreateBusinessCardRequest formData dto.CreateBusinessCardRequest true "BusinessCard"
// @Success 201 {object} dto.BusinessCardResponse
func CreateBusinessCard(c *gin.Context) {
	request := dto.CreateBusinessCardRequest{}
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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

	id, err := shared.NewID()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	businesscard := model.BusinessCard{
		BusinessCardBackgroundID:      request.BusinessCardBackgroundID,
		ARAssetID:                     request.ArAssetsID,
		BusinessCardPartsCoordinateID: request.BusinessCardPartsCoordinateID,
		BusinessCardName:              request.BusinessCardName,
		ID:                            id.String(),
		UserID:                        uid,
		DisplayName:                   request.DisplayName,
		CompanyName:                   request.CompanyName,
		Department:                    request.Department,
		OfficialPosition:              request.OfficialPosition,
		PhoneNumber:                   request.PhoneNumber,
		Email:                         request.Email,
		PostalCode:                    request.PostalCode,
		Address:                       request.Address,
		CreatedAt:                     time.Now(),
	}
	if err = db.Create(&businesscard).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	bcc := model.BusinessCardPartsCoordinate{ID: request.BusinessCardPartsCoordinateID}
	if err = db.First(&bcc).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	bcb := model.BusinessCardBackground{ID: request.BusinessCardBackgroundID}
	if err = db.First(&bcb).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	arassets := model.ARAsset{ID: request.ArAssetsID}
	if err = db.First(&arassets).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := dto.ConvertBC(businesscard, bcc, bcb, arassets)

	c.Header("Location", fmt.Sprintf("/%s", "1"))
	c.JSON(http.StatusCreated, response)
}

// @Tags BusinessCard
// @Router /v1/users/business_cards [GET]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Success 200 {object} []dto.BusinessCardResponse
func ReadAllBusinessCard(c *gin.Context) {
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

	businesscards := []model.BusinessCard{}
	if err = db.Where("user_id = ?", uid).Find(&businesscards).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "failed to fetch businesscards"})
		return
	}

	appConfig, err := config.GetConfig()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ab := drivers.NewAzureBlobDriver(appConfig)

	bcURL, err := ab.GetContainerURL(backgroundContainer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	qcURL, err := ab.GetContainerURL(qrcodeContainer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	tdmcURL, err := ab.GetContainerURL(threeDimentionalModelContainer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	acURL, err := ab.GetContainerURL(audioContainer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	responses := []dto.BusinessCardResponse{}
	for _, businesscard := range businesscards {
		bcc := model.BusinessCardPartsCoordinate{ID: businesscard.BusinessCardPartsCoordinateID}
		if err = db.First(&bcc).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "failed to fetch bcc"})
			return
		}

		bcb := model.BusinessCardBackground{ID: businesscard.BusinessCardBackgroundID}
		if err = db.First(&bcb).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "failed to fetch bcb"})
			return
		}
		bcb.ImagePath = bcURL.Path(bcb.ImagePath)

		arassets := model.ARAsset{ID: businesscard.ARAssetID}
		if err = db.First(&arassets).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "failed to fetch arassets"})
			return
		}
		arassets.ThreeDimentionalModel.ModelPath = tdmcURL.Path(arassets.ThreeDimentionalModel.ModelPath)
		arassets.SpeakingAsset.AudioPath = acURL.Path(arassets.SpeakingAsset.AudioPath)

		var qrCodeImagePath string
		if arassets.QRCodeImagePath != "" {
			qrCodeImagePath = qcURL.Path(arassets.QRCodeImagePath)
		}
		arassets.QRCodeImagePath = qrCodeImagePath

		response := dto.ConvertBC(businesscard, bcc, bcb, arassets)

		responses = append(responses, response)
	}

	c.JSON(http.StatusOK, responses)
}

// @Tags BusinessCard
// @Router /v1/users/business_cards/{business_card_id} [GET]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Param business_card_id path string true "BusinessCard ID"
// @Success 200 {object} dto.BusinessCardResponse
func ReadBusinessCardByID(c *gin.Context) {
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

	appConfig, err := config.GetConfig()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ab := drivers.NewAzureBlobDriver(appConfig)

	bcURL, err := ab.GetContainerURL(backgroundContainer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	qcURL, err := ab.GetContainerURL(qrcodeContainer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	tdmcURL, err := ab.GetContainerURL(threeDimentionalModelContainer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	acURL, err := ab.GetContainerURL(audioContainer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	businesscard := model.BusinessCard{}
	if err = db.Where("id = ? AND user_id = ?", c.Param("business_card_id"), uid).
		Order("id desc").First(&businesscard).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	bcc := model.BusinessCardPartsCoordinate{ID: businesscard.BusinessCardPartsCoordinateID}
	if err = db.First(&bcc).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	bcb := model.BusinessCardBackground{ID: businesscard.BusinessCardBackgroundID}
	if err = db.First(&bcb).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	bcb.ImagePath = bcURL.Path(bcb.ImagePath)

	arassets := model.ARAsset{ID: businesscard.ARAssetID}
	if err = db.First(&arassets).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	arassets.ThreeDimentionalModel.ModelPath = tdmcURL.Path(arassets.ThreeDimentionalModel.ModelPath)
	arassets.SpeakingAsset.AudioPath = acURL.Path(arassets.SpeakingAsset.AudioPath)

	var qrCodeImagePath string
	if arassets.QRCodeImagePath != "" {
		qrCodeImagePath = qcURL.Path(arassets.QRCodeImagePath)
	}
	arassets.QRCodeImagePath = qrCodeImagePath

	response := dto.ConvertBC(businesscard, bcc, bcb, arassets)

	c.JSON(http.StatusOK, response)
}

// @Tags BusinessCard
// @Router /v1/users/business_cards/{business_card_id} [PUT]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Param business_card_id path string true "BusinessCard ID"
// @Param CreateBusinessCardRequest formData dto.CreateBusinessCardRequest true "BusinessCard"
// @Success 204 {object} nil
func UpdateBusinessCard(c *gin.Context) {
	// FIXME: ベタ書き

	// コンテキストから取得
	userID, err := frameworks.GetUID(c)
	if err != nil {
		frameworks.ErrorHandling(c, err, http.StatusInternalServerError)
		return
	}

	db, err := frameworks.GetDB(c)
	if err != nil {
		frameworks.ErrorHandling(c, err, http.StatusInternalServerError)
		return
	}

	// リクエストから取得
	id := c.Param("business_card_id")
	if id == "" {
		reqErr := errors.New("business_card_id is empty")
		frameworks.ErrorHandling(c, reqErr, http.StatusBadRequest)
	}

	request := dto.CreateBusinessCardRequest{}
	if err = c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ユースケース

	// 名刺の取得（存在確認）
	var prevModel model.BusinessCard

	if err = db.Where("id = ?", id).First(&prevModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			frameworks.ErrorHandling(c, errors.New("business card not found"), http.StatusNotFound)
			return
		}
		frameworks.ErrorHandling(c, fmt.Errorf("failed to fetch business card: %w", err), http.StatusInternalServerError)
		return
	}

	if prevModel.UserID != userID {
		frameworks.ErrorHandling(c, errors.New("user does not have permission"), http.StatusForbidden)
		return
	}

	// バリデーション

	// （ARアセットを変える場合は）権限確認

	// （名刺背景を変える場合は）権限確認

	// データベース更新
	// 許して

	if request.BusinessCardBackgroundID != "" {
		prevModel.BusinessCardBackgroundID = request.BusinessCardBackgroundID
	}
	if request.ArAssetsID != "" {
		prevModel.ARAssetID = request.ArAssetsID
	}
	if request.BusinessCardPartsCoordinateID != "" {
		prevModel.BusinessCardPartsCoordinateID = request.BusinessCardPartsCoordinateID
	}
	if request.BusinessCardName != "" {
		prevModel.BusinessCardName = request.BusinessCardName
	}
	if request.DisplayName != "" {
		prevModel.DisplayName = request.DisplayName
	}
	if request.CompanyName != "" {
		prevModel.CompanyName = request.CompanyName
	}
	if request.Department != "" {
		prevModel.Department = request.Department
	}
	if request.OfficialPosition != "" {
		prevModel.OfficialPosition = request.OfficialPosition
	}
	if request.PhoneNumber != "" {
		prevModel.PhoneNumber = request.PhoneNumber
	}
	if request.Email != "" {
		prevModel.Email = request.Email
	}
	if request.PostalCode != "" {
		prevModel.PostalCode = request.PostalCode
	}
	if request.Address != "" {
		prevModel.Address = request.Address
	}

	if err = db.Debug().Save(&prevModel).Error; err != nil {
		frameworks.ErrorHandling(c, fmt.Errorf("failed to update business card: %w", err), http.StatusInternalServerError)
		return
	}

	// レスポンス
	c.JSON(http.StatusNoContent, nil)
}

// @Tags BusinessCard
// @Router /v1/users/business_cards/{business_card_id} [DELETE]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Param business_card_id path string true "BusinessCard ID"
// @Success 204 {object} nil
func DeleteBusinessCard(c *gin.Context) {
	// FIXME: ベタ書きだけど許して

	// コンテキストから取得
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

	// リクエストから取得
	id := c.Param("business_card_id")
	if id == "" {
		reqErr := errors.New("business_card_id is empty")
		frameworks.ErrorHandling(c, reqErr, http.StatusBadRequest)
	}

	// データベースからBusinessCardを取得
	var dbModel model.BusinessCard

	if err = db.Where("id = ?", id).First(&dbModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			frameworks.ErrorHandling(c, errors.New("business card not found"), http.StatusNotFound)
			return
		}
		frameworks.ErrorHandling(c, fmt.Errorf("failed to fetch business card: %w", err), http.StatusInternalServerError)
		return
	}

	// BusinessCardの権限確認
	if dbModel.UserID != uid {
		frameworks.ErrorHandling(c, errors.New("user does not have permission"), http.StatusForbidden)
		return
	}

	// データベースから削除
	if err = db.Delete(&dbModel).Error; err != nil {
		frameworks.ErrorHandling(c, fmt.Errorf("failed to delete business card: %w", err), http.StatusInternalServerError)
		return
	}

	// レスポンス
	c.JSON(http.StatusNoContent, nil)
}

// @Tags BusinessCard
// @Router /v1/business_cards/{business_card_id} [GET]
// @Param business_card_id path string true "BusinessCard ID"
// @Success 200 {object} dto.BusinessCardResponse
func ReadBusinessCardByIDPublic(c *gin.Context) {
	// FIXME: まるごとコピペ
	db, err := frameworks.GetDB(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	appConfig, err := config.GetConfig()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ab := drivers.NewAzureBlobDriver(appConfig)

	bcURL, err := ab.GetContainerURL(backgroundContainer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	qcURL, err := ab.GetContainerURL(qrcodeContainer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	tdmcURL, err := ab.GetContainerURL(threeDimentionalModelContainer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	acURL, err := ab.GetContainerURL(audioContainer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	businesscard := model.BusinessCard{}
	if err = db.Where("id = ?", c.Param("business_card_id")).First(&businesscard).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	bcc := model.BusinessCardPartsCoordinate{ID: businesscard.BusinessCardPartsCoordinateID}
	if err = db.First(&bcc).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	bcb := model.BusinessCardBackground{ID: businesscard.BusinessCardBackgroundID}
	if err = db.First(&bcb).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	bcb.ImagePath = bcURL.Path(bcb.ImagePath)

	arassets := model.ARAsset{ID: businesscard.ARAssetID}
	if err = db.Preload("SpeakingAsset").
		Preload("ThreeDimentionalModel").
		First(&arassets).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	arassets.ThreeDimentionalModel.ModelPath = tdmcURL.Path(arassets.ThreeDimentionalModel.ModelPath)
	arassets.SpeakingAsset.AudioPath = acURL.Path(arassets.SpeakingAsset.AudioPath)

	var qrCodeImagePath string
	if arassets.QRCodeImagePath != "" {
		qrCodeImagePath = qcURL.Path(arassets.QRCodeImagePath)
	}
	arassets.QRCodeImagePath = qrCodeImagePath

	response := dto.ConvertBC(businesscard, bcc, bcb, arassets)

	c.JSON(http.StatusOK, response)
}
