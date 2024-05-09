package categorymodel

import (
	"ccnt/config"
	"ccnt/entities"
)

// This function is used to get all data from the table
// The data is then stored in a slice of struct
// The struct is defined in folder entities in file category.go 
// The struct is then used to store the data from the table
func GetAll() []entities.NameYourStruct {
	
	// Query the database to get all data from the table 
	rows, err := config.DB.Query("Select * from YourDatabase")

	// If there is an error, panic the error
	if err != nil {
		panic(err)
	}

	// Defer the closing of the rows
	defer rows.Close()

	var categories []entities.NameYourStruct

	// Loop through the rows
	for rows.Next() {
		var category entities.NameYourStruct

		// Scan the data from the rows and store it in the struct
		if err := rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt); err != nil {
			panic(err)
		}

		categories = append(categories, category)
	}

	return categories
}

// This function is used to create a new data in the table
// The data is then stored in a struct
// The struct is defined in folder entities in file category.go
// The struct is then used to store the data in the table
// The function returns a boolean value
func Create(category entities.NameYourStruct) bool {

	// Insert the data into the table
	// The data is stored in the struct
	result, err := config.DB.Exec(`
		INSERT INTO YourDatabase (name, created_at, updated_at) 
		VALUE (?, ?, ?)`,
		category.Name,
		category.CreatedAt,
		category.UpdatedAt,
	)

	if err != nil {
		panic(err)
	}

	// Get the last inserted id
	// The id is then used to check if the data is successfully inserted
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	// Return a boolean value
	// If the last inserted id is greater than 0, return true
	return lastInsertId > 0
}

// This function is used to get a single data from the table
// The function returns a struct
// The struct is then used to display the data in the view
func Detail(id int) entities.NameYourStruct {

	// Query the database to get a single data from the table
	row := config.DB.QueryRow("Select id, name from YourDatabase where id = ?", id)

	var category entities.NameYourStruct

	if err := row.Scan(&category.Id, &category.Name); err != nil {
		panic(err)
	}

	return category
}

// This function is used to update a data in the table
// The function returns a boolean value
// The boolean value is used to check if the data is successfully updated
func Update(id int, category entities.NameYourStruct) bool {

	// Update the data in the table
	query, err := config.DB.Exec("UPDATE YourDatabase SET name = ?, updated_at = ? where id = ?", category.Name, category.UpdatedAt, id)

	if err != nil {
		panic(err)
	}

	// Get the number of rows affected
	result, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}
	
	// If the number of rows affected is greater than 0, return true
	return result > 0

}

// This function is used to delete a data in the table
func Delete(id int) error {
	_, err := config.DB.Exec("Delete from YourDatabase where id = ?", id)

	return err
}
