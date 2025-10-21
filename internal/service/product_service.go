package service

import (
	"ai_admin_project/internal/model"
	"ai_admin_project/internal/repository"
	"ai_admin_project/internal/request"
	"errors"
)

type ProductService struct {
	Repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{Repo: repo}
}

func (s *ProductService) Create(req request.CreateProductRequest) (*model.Product, error) {
	product := &model.Product{
		Name:  req.Name,
		Price: req.Price,
	}

	if err := s.Repo.Create(product); err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductService) GetByID(req request.GetProductRequest) (*model.Product, error) {
	return s.Repo.FindByID(req.ID)
}

func (s *ProductService) Update(req request.UpdateProductRequest) (*model.Product, error) {
	product, err := s.Repo.FindByID(req.ID)
	if err != nil {
		return nil, errors.New("product not found")
	}

	if req.Name != "" {
		product.Name = req.Name
	}
	if req.Price != 0 {
		product.Price = req.Price
	}

	if err := s.Repo.Update(product); err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductService) Delete(req request.DeleteProductRequest) error {
	return s.Repo.Delete(req.ID)
}

func (s *ProductService) List(page int, pageSize int) ([]model.Product, int64, error) {
	return s.Repo.FindAll(page, pageSize)
}