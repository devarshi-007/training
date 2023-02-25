package database

import (
	"database/sql"
)

const (
	queryForGetAllProducts = "SELECT ID, NAME FROM inventorydb.products;"
	queryForGetProductById = "SELECT ID, NAME FROM inventorydb.products WHERE ID = ?;"
	queryForUpdateProduct  = "UPDATE products SET NAME=? WHERE ID=?"
	queryForInsertProduct  = "INSERT INTO products (NAME) VALUES (?)"
	queryForDeleteProduct  = "DELETE FROM products WHERE ID = ?"
)

func GetProduct(productId int) (*Product, error) {
	row := DbConn.QueryRow(queryForGetProductById, productId)
	product := &Product{}
	err := row.Scan(&product.Id, &product.Name)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return product, nil
}

func GetProductList() ([]Product, error) {
	results, err := DbConn.Query(queryForGetAllProducts)

	if err != nil {
		return nil, err
	}

	defer results.Close()

	products := make([]Product, 0)
	for results.Next() {
		var product Product
		results.Scan(&product.Id, &product.Name)
		products = append(products, product)
	}

	return products, nil
}

func InsertProduct(product Product) (int, error) {
	result, err := DbConn.Exec(queryForInsertProduct, product.Name)

	if err != nil {
		return 0, nil
	}

	insertID, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return int(insertID), nil
}

func UpdatedProduct(product Product) error {
	_, err := DbConn.Exec(queryForUpdateProduct, product.Name, product.Id)

	if err != nil {
		return err
	}

	return nil
}

func RemoveProduct(productID int) error {
	_, err := DbConn.Query(queryForDeleteProduct, productID)

	if err != nil {
		return err
	}
	return nil
}
