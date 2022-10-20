package handler

import (
	"assignment/app/helpers"
	"assignment/app/models"
	"assignment/app/repository"
	"assignment/app/resource"
	"assignment/app/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	// "bytes"
	// uuid "github.com/satori/go.uuid"
)

type PhotoHandler struct {
	repo repository.PhotoRepository
}

func NewPhotoHandler() *PhotoHandler {
	return &PhotoHandler{
		repository.NewPhotoRepository(),
	}
}

type PhotoOut struct {
	ID           uint      `json:"order_id"`
	CustomerName string    `json:"customer_name"`
	PhotoedAt    time.Time `json:"ordered_at"`
	Items        []ItemOut `gorm:"foreignKey:PhotoID"`
}

// AddPhoto
// @Summary Add new Photo
// @Decription Add new Photo
// @Tags Photo
// @Accept json
// @Produce json
// @Success 200
// @Router /Photo [post]
// @Param data body resource.InputPhoto true "body data"
func (h *PhotoHandler) AddPhoto(c *gin.Context) {
	repo := h.repo
	var req resource.InputPhoto
	err := c.ShouldBind(&req)
	if err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"error_messages": errors}
		response := helpers.APIResponse("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	userId := c.GetInt("UserID")
	var Photo models.Photo
	Photo.UserID = uint(userId)
	err = repo.AddPhoto(&Photo, req)
	if err != nil {
		response := helpers.APIResponse2("Failed when trying to add photo.", http.StatusBadRequest, 0, 0, 0, "")
		c.JSON(http.StatusOK, response)
		return
	}
	response := helpers.APIResponse2("Success Add Photo", http.StatusOK, 0, 0, 0, map[string]interface{}{
		"id":         Photo.ID,
		"title":      Photo.Title,
		"caption":    Photo.Caption,
		"photo_url":  Photo.PhotoUrl,
		"user_id":    Photo.UserID,
		"created_at": Photo.CreatedAt,
	})
	c.JSON(http.StatusOK, response)
}

func (h *PhotoHandler) GetPhoto(c *gin.Context) {
	repo := h.repo
	userId := c.GetInt("UserID")
	var Photos []models.Photo
	err := repo.GetPhoto(&Photos, uint(userId))
	if err != nil {
		response := helpers.APIResponse2("Failed when trying to get photo.", http.StatusBadRequest, 0, 0, 0, "")
		c.JSON(http.StatusOK, response)
		return
	}
	var photoList []map[string]interface{}
	for _, eachPhoto := range Photos {
		data := map[string]interface{}{
			"id":         eachPhoto.ID,
			"title":      eachPhoto.Title,
			"caption":    eachPhoto.Caption,
			"photo_url":  eachPhoto.PhotoUrl,
			"user_id":    eachPhoto.UserID,
			"created_at": eachPhoto.CreatedAt,
		}
		photoList = append(photoList, data)
	}
	response := helpers.APIResponse2("Success Get Photo", http.StatusOK, 0, 0, 0, photoList)
	c.JSON(http.StatusOK, response)
}

func (h *PhotoHandler) DeletePhoto(c *gin.Context) {
	photoId := c.Param("photoId")
	photoIdInt, err := strconv.Atoi(photoId)
	repo := h.repo
	userId := c.GetInt("UserID")
	err = repo.DeletePhoto(uint(userId), uint(photoIdInt))
	if err != nil {
		response := helpers.APIResponse2(err.Error(), http.StatusBadRequest, 0, 0, 0, err)
		c.JSON(http.StatusOK, response)
		return
	}
	response := helpers.APIResponse2("Delete Successful.", http.StatusOK, 0, 0, 0, "")
	c.JSON(http.StatusOK, response)
}

func (h *PhotoHandler) UpdatePhoto(c *gin.Context) {
	photoId := c.Param("photoId")
	photoIdInt, err := strconv.Atoi(photoId)
	repo := h.repo
	var req resource.InputPhoto
	err = c.ShouldBind(&req)
	if err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"error_messages": errors}
		response := helpers.APIResponse("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	userId := c.GetInt("UserID")
	var Photo models.Photo
	Photo.ID = uint(photoIdInt)
	err = repo.UpdatePhoto(&Photo, req, uint(userId))
	if err != nil {
		response := helpers.APIResponse2(err.Error(), http.StatusBadRequest, 0, 0, 0, "")
		c.JSON(http.StatusOK, response)
		return
	}
	response := helpers.APIResponse2("Success Update Photo", http.StatusOK, 0, 0, 0, map[string]interface{}{
		"id":         Photo.ID,
		"title":      Photo.Title,
		"caption":    Photo.Caption,
		"photo_url":  Photo.PhotoUrl,
		"user_id":    Photo.UserID,
		"created_at": Photo.UpdatedAt,
	})
	c.JSON(http.StatusOK, response)
}
