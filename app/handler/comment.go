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

type CommentHandler struct {
	repo repository.CommentRepository
}

func NewCommentHandler() *CommentHandler {
	return &CommentHandler{
		repository.NewCommentRepository(),
	}
}

type CommentOut struct {
	ID           uint      `json:"order_id"`
	CustomerName string    `json:"customer_name"`
	CommentedAt  time.Time `json:"ordered_at"`
	Items        []ItemOut `gorm:"foreignKey:CommentID"`
}

// AddComment
// @Summary Add new Comment
// @Decription Add new Comment
// @Tags Comment
// @Accept json
// @Produce json
// @Success 200
// @Router /Comment [post]
// @Param data body resource.InputComment true "body data"
func (h *CommentHandler) AddComment(c *gin.Context) {
	repo := h.repo
	var req resource.InputComment
	err := c.ShouldBind(&req)
	if err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"error_messages": errors}
		response := helpers.APIResponse("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	userId := c.GetInt("UserID")
	var Comment models.Comment
	Comment.UserID = uint(userId)
	err = repo.AddComment(&Comment, req)
	if err != nil {
		response := helpers.APIResponse2("Failed when trying to add comment.", http.StatusBadRequest, 0, 0, 0, "")
		c.JSON(http.StatusOK, response)
		return
	}
	response := helpers.APIResponse2("Success Add Comment", http.StatusOK, 0, 0, 0, map[string]interface{}{
		"id":         Comment.ID,
		"message":    Comment.Message,
		"photo_id":   Comment.PhotoID,
		"user_id":    Comment.UserID,
		"created_at": Comment.CreatedAt,
	})
	c.JSON(http.StatusOK, response)
}

func (h *CommentHandler) GetComment(c *gin.Context) {
	repo := h.repo
	userId := c.GetInt("UserID")
	var Comments []models.Comment
	err := repo.GetComment(&Comments, uint(userId))
	if err != nil {
		response := helpers.APIResponse2("Failed when trying to get photo.", http.StatusBadRequest, 0, 0, 0, "")
		c.JSON(http.StatusOK, response)
		return
	}
	var photoList []map[string]interface{}
	for _, eachComment := range Comments {
		data := map[string]interface{}{
			"id":         eachComment.ID,
			"message":    eachComment.Message,
			"photo_id":   eachComment.PhotoID,
			"user_id":    eachComment.UserID,
			"created_at": eachComment.CreatedAt,
			"updated_at": eachComment.UpdatedAt,
			"user": map[string]interface{}{
				"id":       eachComment.User.ID,
				"email":    eachComment.User.Email,
				"username": eachComment.User.Username,
			},
			"photo": map[string]interface{}{
				"id":        eachComment.Photo.ID,
				"title":     eachComment.Photo.Title,
				"caption":   eachComment.Photo.Caption,
				"photo_url": eachComment.Photo.PhotoUrl,
				"user_id":   eachComment.Photo.UserID,
			},
		}
		photoList = append(photoList, data)
	}
	response := helpers.APIResponse2("Success Get Comment", http.StatusOK, 0, 0, 0, photoList)
	c.JSON(http.StatusOK, response)
}

func (h *CommentHandler) DeleteComment(c *gin.Context) {
	commentId := c.Param("commentId")
	commentIdInt, err := strconv.Atoi(commentId)
	repo := h.repo
	userId := c.GetInt("UserID")
	err = repo.DeleteComment(uint(userId), uint(commentIdInt))
	if err != nil {
		response := helpers.APIResponse2(err.Error(), http.StatusBadRequest, 0, 0, 0, err)
		c.JSON(http.StatusOK, response)
		return
	}
	response := helpers.APIResponse2("Delete Successful.", http.StatusOK, 0, 0, 0, "")
	c.JSON(http.StatusOK, response)
}

func (h *CommentHandler) UpdateComment(c *gin.Context) {
	commentId := c.Param("commentId")
	commentIdInt, err := strconv.Atoi(commentId)
	repo := h.repo
	var req resource.UpdateComment
	err = c.ShouldBind(&req)
	if err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"error_messages": errors}
		response := helpers.APIResponse("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	userId := c.GetInt("UserID")
	var Comment models.Comment
	Comment.ID = uint(commentIdInt)
	err = repo.UpdateComment(&Comment, req, uint(userId))
	if err != nil {
		response := helpers.APIResponse2(err.Error(), http.StatusBadRequest, 0, 0, 0, "")
		c.JSON(http.StatusOK, response)
		return
	}
	response := helpers.APIResponse2("Success Update Comment", http.StatusOK, 0, 0, 0, map[string]interface{}{
		"id":         Comment.ID,
		"message":    Comment.Message,
		"photo_id":   Comment.PhotoID,
		"user_id":    Comment.UserID,
		"updated_at": Comment.UpdatedAt,
	})
	c.JSON(http.StatusOK, response)
}
