package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"
)

type ProductReprository struct {
	connection *sql.DB
}

func NewProductReprository(connection *sql.DB) ProductReprository {
	return ProductReprository{
		connection: connection,
	}
}

// (pr *ProductReprository) antes do nome da funcao atribui essa funcao à struct
func (pr *ProductReprository) GetProducts() ([]model.Product, error) {

	query := "SELECT id, product_name, price FROM product"

	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)

		return []model.Product{}, err
	}

	var productList []model.Product
	var productObj model.Product

	for rows.Next() {
		err := rows.Scan(
			&productObj.ID,
			&productObj.Name,
			&productObj.Price)

		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}

		productList = append(productList, productObj)
	}

	rows.Close()

	return productList, nil
}

func (pr *ProductReprository) CreateProduct(product model.Product) (int, error) {
	
	var id int

	query, err := pr.connection.Prepare("INSERT INTO product" +
		"(product_name, price)" +
		" VALUES ($1, $2) RETURNING id") //$1 e $2 sao os parametros que a query vai receber
	// inserir os valores nas colunas do banco de dados e retorna o id gerado automaticamente

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(product.Name, product.Price).Scan(&id) 
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()

	return id, nil

}
