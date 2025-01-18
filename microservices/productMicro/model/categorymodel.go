package model

import (
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
