package request

type GetProductRequest struct {
	ID uint `uri:"id" binding:"required,gte=1"`
}

type CreateProductRequest struct {
	Name  string `json:"name" binding:"required"`
	Price uint   `json:"price" binding:"required"`
}

type UpdateProductRequest struct {
	ID    uint   `json:"id" binding:"required,gte=1"`
	Name  string `json:"name" binding:"omitempty"`
	Price uint   `json:"price" binding:"omitempty,gte=0"`
}

type DeleteProductRequest struct {
	ID uint `json:"id" binding:"required,gte=1"`
}
