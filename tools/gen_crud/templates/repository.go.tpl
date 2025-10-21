
package repository

import (
	"ai_admin_project/internal/model"
	"gorm.io/gorm"
)

type {{.ModelName}}Repository struct {
	DB *gorm.DB
}

func New{{.ModelName}}Repository(db *gorm.DB) *{{.ModelName}}Repository {
	return &{{.ModelName}}Repository{DB: db}
}

func (r *{{.ModelName}}Repository) Create({{.LowerModelName}} *model.{{.ModelName}}) error {
	return r.DB.Create({{.LowerModelName}}).Error
}

func (r *{{.ModelName}}Repository) FindByID(id uint) (*model.{{.ModelName}}, error) {
	var {{.LowerModelName}} model.{{.ModelName}}
	if err := r.DB.First(&{{.LowerModelName}}, id).Error; err != nil {
		return nil, err
	}
	return &{{.LowerModelName}}, nil
}

func (r *{{.ModelName}}Repository) Update({{.LowerModelName}} *model.{{.ModelName}}) error {
	return r.DB.Save({{.LowerModelName}}).Error
}

func (r *{{.ModelName}}Repository) Delete(id uint) error {
	return r.DB.Delete(&model.{{.ModelName}}{}, id).Error
}

func (r *{{.ModelName}}Repository) FindAll(page int, pageSize int) ([]model.{{.ModelName}}, int64, error) {
	var {{.LowerModelNamePlural}} []model.{{.ModelName}}
	var total int64

	offset := (page - 1) * pageSize

	if err := r.DB.Model(&model.{{.ModelName}}{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.DB.Offset(offset).Limit(pageSize).Find(&{{.LowerModelNamePlural}}).Error; err != nil {
		return nil, 0, err
	}

	return {{.LowerModelNamePlural}}, total, nil
}
