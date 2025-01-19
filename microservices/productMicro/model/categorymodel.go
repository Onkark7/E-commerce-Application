package model

import (
	"errors"
	"fmt"
	data "productmicro/Database"
	"time"
)

type Category struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Parent_id   int       `json:"parent_id"`
	CreatedAt   time.Time ` json:"created_at"`
}

func AddCategory(category Category) (Category, error) {

	category.CreatedAt = time.Now()

	query := `INSERT into category (name,description,parent_id,created_at) Values(?,?,?,?)`

	result, err := data.DB.Exec(query, category.Name, category.Description, category.Parent_id, category.CreatedAt)

	if err != nil {
		return category, err
	}

	LastID, err := result.LastInsertId()

	if err != nil {
		return category, err
	}

	row := data.DB.QueryRow(`SELECT id,name,description,parent_id,created_at FROM category WHERE id =?`, LastID)

	err = row.Scan(&category.Id, &category.Name, &category.Description, &category.Parent_id, &category.CreatedAt)

	if err != nil {
		return category, err
	}

	return category, nil
}

func UpdateCategory(category Category, id int) (Category, error) {

	catUpdate := make(map[string]interface{})

	if category.Name != "" {
		catUpdate["name"] = category.Name
	}

	if category.Description != "" {
		catUpdate["description"] = category.Description
	}

	if category.Parent_id != 0 {
		catUpdate["parent_id"] = category.Parent_id
	}

	if len(catUpdate) == 0 {
		return category, errors.New("no fields provided for update")
	}

	query := "UPDATE category SET "
	values := []interface{}{}
	i := 0

	for field, value := range catUpdate {
		if i > 0 {
			query += ", "
		}
		query += fmt.Sprintf("%s = ?", field)
		values = append(values, value)
		i++
	}

	query += " WHERE id = ?"
	values = append(values, id)

	_, err := data.DB.Exec(query, values...)
	if err != nil {
		return category, err
	}

	return category, nil
}

func GetallCategory() ([]Category, error) {
	var categorys []Category

	query := `Select id,name,description,parent_id FROM category`

	rows, err := data.DB.Query(query)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var category Category
		err := rows.Scan(
			&category.Id,
			&category.Name,
			&category.Description,
			&category.Parent_id,
		)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}
		categorys = append(categorys, category)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Rows iteration error:", err)
		return nil, err
	}

	return categorys, nil
}

func GetCategoryWithID(id int) (Category, error) {
	var category Category

	query := `select id,name,description,parent_id from category where id=?`

	result := data.DB.QueryRow(query, id)

	err := result.Scan(&category.Id, &category.Name, &category.Description, &category.Parent_id)

	if err != nil {
		return category, err
	}

	return category, nil
}
