package handler

import (
	"assignment/app/helpers"
	"assignment/app/models"
	"assignment/app/repository"
	"assignment/app/resource"
	"assignment/app/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	// "bytes"
	// uuid "github.com/satori/go.uuid"
)

type UserHandler struct {
	repo repository.UserRepository
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		repository.NewUserRepository(),
	}
}

type UserOut struct {
	ID           uint      `json:"order_id"`
	CustomerName string    `json:"customer_name"`
	UseredAt     time.Time `json:"ordered_at"`
	Items        []ItemOut `gorm:"foreignKey:UserID"`
}

type ItemOut struct {
	ItemID      uint   `json:"item_id"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
	UserID      uint   `json:"order_id"`
}

// AddUser
// @Summary Add new User
// @Decription Add new User
// @Tags User
// @Accept json
// @Produce json
// @Success 200
// @Router /User [post]
// @Param data body resource.InputUser true "body data"
func (h *UserHandler) RegisterUser(c *gin.Context) {
	repo := h.repo
	var req resource.InputUser
	err := c.ShouldBind(&req)
	if err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"error_messages": errors}
		response := helpers.APIResponse("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	var User models.User
	err = repo.RegisterUser(&User, req)
	if err != nil {
		response := helpers.APIResponse2("Failed when trying to register, perhaps already registered.", http.StatusBadRequest, 0, 0, 0, "")
		c.JSON(http.StatusOK, response)
		return
	}
	response := helpers.APIResponse2("Success Register User", http.StatusOK, 0, 0, 0, map[string]interface{}{
		"id":       User.ID,
		"email":    User.Email,
		"username": User.Username,
		"age":      User.Age,
	})
	c.JSON(http.StatusOK, response)
}

func (h *UserHandler) LoginUser(c *gin.Context) {
	repo := h.repo
	var req resource.Login
	err := c.ShouldBind(&req)
	if err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"error_messages": errors}
		response := helpers.APIResponse("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	token, err := repo.Login(req.Email, req.Password)
	if err != nil {
		response := helpers.APIResponse2(err.Error(), http.StatusBadRequest, 0, 0, 0, err)
		c.JSON(http.StatusOK, response)
		return
	}
	response := helpers.APIResponse2("Success Login", http.StatusOK, 0, 0, 0, map[string]interface{}{
		"token": token,
	})
	c.JSON(http.StatusOK, response)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	repo := h.repo
	userId := c.GetInt("UserID")
	err := repo.DeleteUser(userId)
	if err != nil {
		response := helpers.APIResponse2(err.Error(), http.StatusBadRequest, 0, 0, 0, err)
		c.JSON(http.StatusOK, response)
		return
	}
	response := helpers.APIResponse2("Delete Successful.", http.StatusOK, 0, 0, 0, "")
	c.JSON(http.StatusOK, response)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	repo := h.repo
	var req resource.UpdateUser
	err := c.ShouldBind(&req)
	if err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"error_messages": errors}
		response := helpers.APIResponse("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	userId := c.GetInt("UserID")
	var User models.User
	User.ID = uint(userId)
	err = repo.UpdateUser(&User, req)
	if err != nil {
		response := helpers.APIResponse2("Failed when trying to update", http.StatusBadRequest, 0, 0, 0, "")
		c.JSON(http.StatusOK, response)
		return
	}
	response := helpers.APIResponse2("Success Update User", http.StatusOK, 0, 0, 0, map[string]interface{}{
		"id":         User.ID,
		"email":      User.Email,
		"username":   User.Username,
		"age":        User.Age,
		"updated_at": User.UpdatedAt,
	})
	c.JSON(http.StatusOK, response)
}
