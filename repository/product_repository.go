package repository

import (
	"database/sql"
	"fmt"

	"github.com/alissoncorsair/golang-api/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return ProductRepository{
		connection: db,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT id, product_name, price FROM product"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println("Error to get products: ", err)
		return []model.Product{}, err
	}

	defer rows.Close()

	var productList []model.Product
	var product model.Product

	for rows.Next() {
		err := rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			fmt.Println("Error to get product: ", err)
			return []model.Product{}, err
		}

		productList = append(productList, product)
	}

	return productList, nil
}

func (pr *ProductRepository) CreateProduct(product model.Product) (int, error) {
	var id int
	query := "INSERT INTO product(product_name, price) VALUES($1, $2) RETURNING id"
	stmt, err := pr.connection.Prepare(query)

	if err != nil {
		fmt.Println("Error to prepare statement")
		return 0, err
	}

	defer stmt.Close()

	err = stmt.QueryRow(product.Name, product.Price).Scan(&id)

	if err != nil {
		fmt.Println("Error to insert product: ", err)
		return 0, err
	}

	return id, nil
}

func (pr *ProductRepository) GetProductByID(id int) (*model.Product, error) {
	query, err := pr.connection.Prepare("SELECT id, product_name, price FROM product WHERE id = $1")

	if err != nil {
		fmt.Println("Error to prepare statement")
		return nil, err
	}

	defer query.Close()

	var product model.Product

	err = query.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		fmt.Println("Error to get product: ", err)
		return nil, err
	}

	return &product, nil

}
