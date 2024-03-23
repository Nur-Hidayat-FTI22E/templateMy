package categorymodel

import (
	"ccnt/config"
	"ccnt/entities"
)

func GetAll() []entities.Categories {
	rows, err := config.DB.Query("Select * from categoriess")

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

	return categories
}

func Create(category entities.Categories) bool {
	result, err := config.DB.Exec(`
		INSERT INTO categoriess (name, created_at, updated_at) 
		VALUE (?, ?, ?)`,
		category.Name,
		category.CreatedAt,
		category.UpdatedAt,
	)

	if err != nil {
		panic(err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return lastInsertId > 0
}

func Detail(id int) entities.Categories {
	row := config.DB.QueryRow("Select id, name from categoriess where id = ?", id)

	var category entities.Categories

	if err := row.Scan(&category.Id, &category.Name); err != nil {
		panic(err)
	}

	return category
}

func Update(id int, category entities.Categories) bool {
	query, err := config.DB.Exec("UPDATE categoriess SET name = ?, updated_at = ? where id = ?", category.Name, category.UpdatedAt, id)

	if err != nil {
		panic(err)
	}

	result, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}

	return result > 0

}

func Delete(id int) error {
	_, err := config.DB.Exec("Delete from categoriess where id = ?", id)

	return err
}
