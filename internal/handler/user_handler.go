package handler

import (
	"ai_admin_project/internal/request"
	"ai_admin_project/internal/response"
	"ai_admin_project/internal/service"
	"ai_admin_project/pkg/utils"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{Service: s}
}

func (h *UserHandler) Register(c *gin.Context) {
	var req request.RegisterRequest
	if err := utils.BindAndValidate(c, &req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	user, err := h.Service.Register(req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(response.ToUserResponse(user), c)
}

func (h *UserHandler) Login(c *gin.Context) {
	var req request.LoginRequest
	if err := utils.BindAndValidate(c, &req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	token, err := h.Service.Login(req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(gin.H{"token": token}, c)
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	// No request struct needed for GetProfile, but we still call BindAndValidate
	// to ensure no unexpected query parameters are passed.
	if err := utils.BindAndValidate(c, nil); err != nil {
		response.FailWithMessage(err.Error(), c)
		return	
	}

	userID, _ := c.Get("userID")
	// In a real application, you would fetch user details from the service
	// For simplicity, we just return the user ID from the token
	response.OkWithData(gin.H{"user_id": userID}, c)
}