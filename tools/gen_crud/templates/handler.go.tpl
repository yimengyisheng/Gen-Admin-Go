
package handler

import (
	"ai_admin_project/internal/request"
	"ai_admin_project/internal/response"
	"ai_admin_project/internal/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

type {{.ModelName}}Handler struct {
	Service *service.{{.ModelName}}Service
}

func New{{.ModelName}}Handler(s *service.{{.ModelName}}Service) *{{.ModelName}}Handler {
	return &{{.ModelName}}Handler{Service: s}
}

func (h *{{.ModelName}}Handler) Create(c *gin.Context) {
	var req request.Create{{.ModelName}}Request
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	{{.LowerModelName}}, err := h.Service.Create(req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData({{.LowerModelName}}, c)
}

func (h *{{.ModelName}}Handler) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.FailWithMessage("Invalid ID", c)
		return
	}

	{{.LowerModelName}}, err := h.Service.GetByID(uint(id))
	if err != nil {
		response.FailWithMessage("{{.ModelName}} not found", c)
		return
	}

	response.OkWithData({{.LowerModelName}}, c)
}

func (h *{{.ModelName}}Handler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.FailWithMessage("Invalid ID", c)
		return
	}

	var req request.Update{{.ModelName}}Request
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	{{.LowerModelName}}, err := h.Service.Update(uint(id), req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData({{.LowerModelName}}, c)
}

func (h *{{.ModelName}}Handler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.FailWithMessage("Invalid ID", c)
		return
	}

	if err := h.Service.Delete(uint(id)); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithMessage("{{.ModelName}} deleted successfully", c)
}

func (h *{{.ModelName}}Handler) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	{{.LowerModelNamePlural}}, total, err := h.Service.List(page, pageSize)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(gin.H{
		"list":  {{.LowerModelNamePlural}},
		"total": total,
	}, c)
}
