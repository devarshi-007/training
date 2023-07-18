package database

import (
	"context"
	"database/sql"
	"time"
)

const (
	queryForGetAllProducts = "SELECT ID, NAME FROM inventorydb.products;"
	queryForGetProductById = "SELECT ID, NAME FROM inventorydb.products WHERE ID = ?;"
	queryForUpdateProduct  = "UPDATE products SET NAME=? WHERE ID=?"
	queryForInsertProduct  = "INSERT INTO products (NAME) VALUES (?)"
	queryForDeleteProduct  = "DELETE FROM products WHERE ID = ?"
)

func GetProduct(productId int) (*Product, error) {
	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()
	row := DbConn.QueryRowContext(ctx, queryForGetProductById, productId)
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
	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()
	results, err := DbConn.QueryContext(ctx, queryForGetAllProducts)

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
	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()
	result, err := DbConn.ExecContext(ctx, queryForInsertProduct, product.Name)

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
	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()
	_, err := DbConn.ExecContext(ctx, queryForUpdateProduct, product.Name, product.Id)

	if err != nil {
		return err
	}

	return nil
}

func RemoveProduct(productID int) error {
	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()
	_, err := DbConn.QueryContext(ctx, queryForDeleteProduct, productID)

	if err != nil {
		return err
	}
	return nil
}
