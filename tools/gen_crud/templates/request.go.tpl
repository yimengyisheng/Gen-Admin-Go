
package request

type Create{{.ModelName}}Request struct {
	Name  string `json:"name" binding:"required"`
	Price uint   `json:"price" binding:"required"`
}

type Update{{.ModelName}}Request struct {
	Name  string `json:"name"`
	Price uint   `json:"price"`
}
