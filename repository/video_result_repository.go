package repository

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"strings"
	"surpreedz-backend/model"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type VideoResultRepository interface {
	Create(videoResult *model.VideoResult) error
	FindById(id int) (model.VideoResult, error)
	FindAll(page int, itemPerPage int) ([]model.VideoResult, error)
	UpdateByID(videoResult *model.VideoResult, by map[string]interface{}) error
	Delete(videoResult *model.VideoResult) error
}

type videoResultRepository struct {
	db  *gorm.DB
	azr *azblob.ServiceClient
}

func (v *videoResultRepository) Create(videoResult *model.VideoResult) error {

	tx := v.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}
	containerClient, err := v.azr.NewContainerClient("videoresult")
	if err != nil {
		log.Fatalln("Error getting container client")
	}

	uid, err := uuid.NewV4()
	if err != nil {
		fmt.Println(err)
	}

	splittedUrl := strings.Split(videoResult.DataUrl, ",")
	//contentType := splittedUrl[0]
	dataUrl := splittedUrl[1]
	video, err := base64.StdEncoding.DecodeString(dataUrl)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println("Image : ", image)

	blockBlobClient, err := containerClient.NewBlockBlobClient(time.Now().Format("20060102") + uid.String() + ".mp4")
	if err != nil {
		fmt.Println(err)
	}

	blobUploadResponse, err := blockBlobClient.UploadBuffer(context.TODO(), video, azblob.UploadOption{})
	if err != nil {
		fmt.Println(err)
	}

	videoResult.VideoLink = time.Now().Format("20060102") + uid.String() + ".mp4"

	fmt.Println("Upload Response : ", blobUploadResponse)

	err = v.db.Create(videoResult).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (v *videoResultRepository) FindById(id int) (model.VideoResult, error) {
	var videoResult model.VideoResult
	result := v.db.First(&videoResult, "order_id = ?", id)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return videoResult, nil
		} else {
			return videoResult, err
		}
	}
	return videoResult, nil
}

func (v *videoResultRepository) FindAll(page int, itemPerPage int) ([]model.VideoResult, error) {
	var videoResult []model.VideoResult
	offset := itemPerPage * (page - 1)
	result := v.db.Unscoped().Order("created_at").Limit(itemPerPage).Offset(offset).Find(&videoResult)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return videoResult, nil
		} else {
			return videoResult, err
		}
	}
	return videoResult, nil
}

func (v *videoResultRepository) UpdateByID(videoResult *model.VideoResult, by map[string]interface{}) error {
	result := v.db.Model(videoResult).Updates(by).Error
	return result
}

func (v *videoResultRepository) Delete(videoResult *model.VideoResult) error {
	result := v.db.Delete(videoResult).Error
	return result
}

func NewVideoResultRepository(db *gorm.DB, azr *azblob.ServiceClient) VideoResultRepository {
	repo := new(videoResultRepository)
	repo.db = db
	repo.azr = azr
	return repo
}
