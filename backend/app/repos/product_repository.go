package repos

import (
	"context"

	"github.com/ictlife/infra-interview-may-2020/backend/app/db"
	"github.com/ictlife/infra-interview-may-2020/backend/app/entities"
)

const (
	selectProductsSQL             = "SELECT id, name, price, created_at, updated_at FROM products"
	selectProductByIDSQL          = selectProductsSQL + " WHERE id = $1"
	selectProductsSortedByNameSQL = selectProductsSQL + " ORDER BY name"
	deleteProductSQL              = "DELETE FROM products WHERE id = $1"
	insertProductSQL              = "INSERT INTO products (name, price, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id"
	updateProductSQL              = "UPDATE products SET (name, price, updated_at) = ($1, $2, $3) WHERE id = $4"
)

type (
	ProductRepository interface {
		All(ctx context.Context, operations db.SQLOperations) ([]*entities.Product, error)
		Delete(ctx context.Context, operations db.SQLOperations, product *entities.Product) error
		FindByID(ctx context.Context, operations db.SQLOperations, productID int64) (*entities.Product, error)
		Save(ctx context.Context, operations db.SQLOperations, product *entities.Product) error
	}

	AppProductRepository struct {
	}
)

func NewProductRepository() *AppProductRepository {
	return &AppProductRepository{}
}

func (r *AppProductRepository) All(ctx context.Context, operations db.SQLOperations) ([]*entities.Product, error) {

	products := make([]*entities.Product, 0)

	rows, err := operations.QueryContext(ctx, selectProductsSortedByNameSQL)
	if err != nil {
		return products, err
	}

	defer rows.Close()

	for rows.Next() {
		product, err := r.scanRow(rows)
		if err != nil {
			return products, err
		}

		products = append(products, product)
	}

	return products, rows.Err()
}

func (r *AppProductRepository) Delete(ctx context.Context, operations db.SQLOperations, product *entities.Product) error {

	_, err := operations.ExecContext(ctx, deleteProductSQL, product.ID)
	return err
}

func (r *AppProductRepository) FindByID(ctx context.Context, operations db.SQLOperations, productID int64) (*entities.Product, error) {

	row := operations.QueryRowContext(ctx, selectProductByIDSQL, productID)
	return r.scanRow(row)
}

func (r *AppProductRepository) Save(ctx context.Context, operations db.SQLOperations, product *entities.Product) error {

	product.Touch()

	if product.IsNew() {
		return operations.QueryRowContext(
			ctx,
			insertProductSQL,
			product.Name,
			product.Price,
			product.CreatedAt,
			product.UpdatedAt,
		).Scan(&product.ID)
	}

	_, err := operations.ExecContext(
		ctx,
		updateProductSQL,
		product.Name,
		product.Price,
		product.UpdatedAt,
		product.ID,
	)
	return err
}

func (r *AppProductRepository) scanRow(rowScanner db.RowScanner) (*entities.Product, error) {
	var product entities.Product

	err := rowScanner.Scan(
		&product.ID,
		&product.Name,
		&product.Price,
		&product.Timestamps.CreatedAt,
		&product.Timestamps.UpdatedAt,
	)

	return &product, err
}
