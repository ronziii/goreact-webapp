package services

import (
	"context"

	"github.com/ictlife/infra-interview-may-2020/backend/app/db"
	"github.com/ictlife/infra-interview-may-2020/backend/app/entities"
	"github.com/ictlife/infra-interview-may-2020/backend/app/forms"
	"github.com/ictlife/infra-interview-may-2020/backend/app/repos"
)

type (
	ProductService interface {
		CreateProduct(context.Context, db.DB, *forms.ProductForm) (*entities.Product, error)
		DeleteProduct(context.Context, db.DB, int64) (*entities.Product, error)
		GetProduct(context.Context, db.DB, int64) (*entities.Product, error)
		ListProducts(context.Context, db.DB) (*entities.ProductList, error)
		UpdateProduct(context.Context, db.DB, int64, *forms.ProductForm) (*entities.Product, error)
	}

	AppProductService struct {
		productRepository repos.ProductRepository
	}
)

func NewProductService(
	productRepository repos.ProductRepository,
) *AppProductService {

	return &AppProductService{
		productRepository: productRepository,
	}
}

func (s *AppProductService) CreateProduct(ctx context.Context, dB db.DB, form *forms.ProductForm) (*entities.Product, error) {

	product := &entities.Product{
		Name:  form.Name,
		Price: form.Price,
	}

	err := s.productRepository.Save(ctx, dB, product)

	return product, err

}
func (s *AppProductService) DeleteProduct(ctx context.Context, dB db.DB, productID int64) (*entities.Product, error) {

	product, err := s.productRepository.FindByID(ctx, dB, productID)
	if err != nil {
		return product, err
	}

	err = s.productRepository.Delete(ctx, dB, product)

	return product, err

}
func (s *AppProductService) GetProduct(ctx context.Context, dB db.DB, productID int64) (*entities.Product, error) {
	return s.productRepository.FindByID(ctx, dB, productID)
}

func (s *AppProductService) ListProducts(ctx context.Context, dB db.DB) (*entities.ProductList, error) {

	productList := &entities.ProductList{}

	products, err := s.productRepository.All(ctx, dB)
	if err != nil {
		return productList, err
	}

	productList.Products = products

	return productList, nil
}

func (s *AppProductService) UpdateProduct(ctx context.Context, dB db.DB, productID int64, form *forms.ProductForm) (*entities.Product, error) {

	product, err := s.productRepository.FindByID(ctx, dB, productID)
	if err != nil {
		return product, err
	}

	product.Name = form.Name
	product.Price = form.Price

	err = s.productRepository.Save(ctx, dB, product)

	return product, err
}
