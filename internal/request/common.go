package request

// PaginationRequest defines the common structure for pagination query parameters.
type PaginationRequest struct {
	Page     int `form:"page" binding:"omitempty,gte=1" default:"1"`
	PageSize int `form:"page_size" binding:"omitempty,gte=1,lte=100" default:"10"`
}