package categorymodel

import (
	"ccnt/config"
	"ccnt/entities"
)

func GetAll() []entities.Categories {
	rows, err := config.DB.Query("Select * from categories")

	if err != nil {
		panic(err)
	}

	defer rows.Close()


	var categories []entities.Categories

	for rows.Next() {
		var category entities.Categories
		if err := rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt); err != nil {
			panic(err)
		}

		categories = append(categories, category)
	}

	return	categories
}