package model

type Product struct {
	BaseModel
	Name  string `gorm:"not null" json:"name"`
	Price uint   `gorm:"not null" json:"price"`
}
