package model

import (
	"database/sql"
	"errors"
	"fmt"
	data "productmicro/Database"
	"time"
)

type Product struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	Image_url   string    `json:"image"`
	CreatedAt   time.Time `json:"created_at"`
	Category_id int       `category_id`
}

func Addproduct(product Product) (Product, error) {

	product.CreatedAt = time.Now()

	query := `INSERT into product(name,description,price,stock,image,created_at,category_id)
	        Values (?,?,?,?,?,?,?)`

	result, err := data.DB.Exec(query, product.Name, product.Description, product.Price, product.Stock, product.Image_url, product.CreatedAt, product.Category_id)

	if err != nil {
		return product, err
	}

	CurrentID, err := result.LastInsertId()

	if err != nil {
		return product, err
	}

	row := data.DB.QueryRow(`SELECT id,name,description,price,stock,image,category_id From product WHERE id= ? `, CurrentID)

	err = row.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Stock, &product.Image_url, &product.Category_id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return product, errors.New("product not found after insertion")
		}
		return product, err
	}

	return product, nil
}

func Updateproduct(product Product, ID int) (Product, error) {

	update := make(map[string]interface{})

	if product.Name != "" {
		update["name"] = product.Name
	}
	if product.Description != "" {
		update["description"] = product.Description
	}
	if product.Price != 0 {
		update["price"] = product.Price
	}
	if product.Stock != 0 {
		update["stock"] = product.Stock
	}
	if product.Image_url != "" {
		update["image"] = product.Image_url
	}
	if product.Category_id != 0 {
		update["category_id"] = product.Category_id
	}

	if len(update) == 0 {
		return product, errors.New("no fields provided for update")
		return product, nil
	}

	query := "UPDATE product SET "
	values := []interface{}{}
	i := 0

	for feild, value := range update {

		if i > 0 {
			query += ", "
		}

		query = query + fmt.Sprintf("%s = ?", feild)
		values = append(values, value)
		i++
	}

	query += " WHERE id=?"
	values = append(values, ID)

	_, err := data.DB.Exec(query, values...)
	if err != nil {
		return product, err
	}

	return product, nil
}
