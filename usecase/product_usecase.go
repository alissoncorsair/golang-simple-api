package usecase

import (
	"github.com/alissoncorsair/golang-api/model"
	"github.com/alissoncorsair/golang-api/repository"
)

type ProductUseCase struct {
	repository repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) ProductUseCase {
	return ProductUseCase{
		repository: repo,
	}
}

func (p *ProductUseCase) GetProducts() ([]model.Product, error) {
	return p.repository.GetProducts()
}

func (p *ProductUseCase) CreateProduct(product model.Product) (model.Product, error) {
	productID, err := p.repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}

	product.ID = productID

	return product, nil
}

func (p *ProductUseCase) GetProductByID(id int) (*model.Product, error) {
	product, err := p.repository.GetProductByID(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}
