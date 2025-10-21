
package model

import "gorm.io/gorm"

type {{.ModelName}} struct {
	gorm.Model
	Name  string `gorm:"not null"`
	Price uint   `gorm:"not null"`
}
