
package service

import (
	"ai_admin_project/internal/model"
	"ai_admin_project/internal/repository"
	"ai_admin_project/internal/request"
)

type {{.ModelName}}Service struct {
	Repo *repository.{{.ModelName}}Repository
}

func New{{.ModelName}}Service(repo *repository.{{.ModelName}}Repository) *{{.ModelName}}Service {
	return &{{.ModelName}}Service{Repo: repo}
}

func (s *{{.ModelName}}Service) Create(req request.Create{{.ModelName}}Request) (*model.{{.ModelName}}, error) {
	{{.LowerModelName}} := &model.{{.ModelName}}{
		Name:  req.Name,
		Price: req.Price,
	}

	if err := s.Repo.Create({{.LowerModelName}}); err != nil {
		return nil, err
	}

	return {{.LowerModelName}}, nil
}

func (s *{{.ModelName}}Service) GetByID(id uint) (*model.{{.ModelName}}, error) {
	return s.Repo.FindByID(id)
}

func (s *{{.ModelName}}Service) Update(id uint, req request.Update{{.ModelName}}Request) (*model.{{.ModelName}}, error) {
	{{.LowerModelName}}, err := s.Repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		{{.LowerModelName}}.Name = req.Name
	}
	if req.Price != 0 {
		{{.LowerModelName}}.Price = req.Price
	}

	if err := s.Repo.Update({{.LowerModelName}}); err != nil {
		return nil, err
	}

	return {{.LowerModelName}}, nil
}

func (s *{{.ModelName}}Service) Delete(id uint) error {
	return s.Repo.Delete(id)
}

func (s *{{.ModelName}}Service) List(page int, pageSize int) ([]model.{{.ModelName}}, int64, error) {
	return s.Repo.FindAll(page, pageSize)
}
