package handler

import (
	"ai_admin_project/internal/request"
	"ai_admin_project/internal/response"
	"ai_admin_project/internal/service"
	"ai_admin_project/pkg/utils"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	Service *service.ProductService
}

func NewProductHandler(s *service.ProductService) *ProductHandler {
	return &ProductHandler{Service: s}
}

func (h *ProductHandler) Create(c *gin.Context) {
	var req request.CreateProductRequest
	if err := utils.BindAndValidate(c, &req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	product, err := h.Service.Create(req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(product, c)
}

func (h *ProductHandler) Get(c *gin.Context) {
	var req request.GetProductRequest
	if err := utils.BindAndValidate(c, &req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	product, err := h.Service.GetByID(req)
	if err != nil {
		response.FailWithMessage("Product not found", c)
		return
	}

	response.OkWithData(product, c)
}

func (h *ProductHandler) Update(c *gin.Context) {
	var req request.UpdateProductRequest
	if err := utils.BindAndValidate(c, &req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	product, err := h.Service.Update(req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(product, c)
}

func (h *ProductHandler) Delete(c *gin.Context) {
	var req request.DeleteProductRequest
	if err := utils.BindAndValidate(c, &req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := h.Service.Delete(req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithMessage("Product deleted successfully", c)
}

func (h *ProductHandler) List(c *gin.Context) {
	var req request.PaginationRequest
	if err := utils.BindAndValidate(c, &req, "page", "page_size"); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	products, total, err := h.Service.List(req.Page, req.PageSize)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(gin.H{
		"list":  products,
		"total": total,
	}, c)
}
