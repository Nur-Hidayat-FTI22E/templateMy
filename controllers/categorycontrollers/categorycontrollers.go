package categorycontrollers

import (
	"ccnt/entities"
	"ccnt/models/categorymodel"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

// This function is used to render the index page
func Index(w http.ResponseWriter, r *http.Request) {
	
	// Get all data from the table
	categories := categorymodel.GetAll()

	// Create a map to store the data
	// The key is a string and the value is an interface
	data := map[string]any{
		"categories": categories,
	}

	// Parse the template file
	// Execute the template file
	temp, err := template.ParseFiles("views/category/index.html")

	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)	
}

// This function is used to render the add page
// It will render views/category/created.html
// and return the html to the client
func Add(w http.ResponseWriter, r *http.Request) {

	// If the method is GET
	// Parse the template file
	// Execute the template file
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/category/created.html")

		if err != nil {
			panic(err)
		}

		temp.Execute(w, nil)
	}

	// If the method is POST
	// Create a new data in the table
	// Redirect to /categories
	// The status is SeeOther
	if r.Method == "POST" {
		var category entities.NameYourStruct

		category.Name = r.FormValue("name")
		category.CreatedAt = time.Now()
		category.UpdatedAt = time.Now()

		if ok := categorymodel.Create(category); !ok {
			temp, _ := template.ParseFiles("views/category/created.html")
			temp.Execute(w, nil)
		}

		http.Redirect(w, r, "/categories", http.StatusSeeOther)
	}
}


// This function is used to render the edit page
// It will render views/category/edit.html
// and return the html to the client
func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/category/edit.html")

		if err != nil {
			panic(err)
		}

		idString := r.URL.Query().Get("id")
		Id, err := strconv.Atoi(idString)

		if err != nil {
			panic(err)
		}

		category := categorymodel.Detail(Id)
		data := map[string]any{
			"category": category,
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var category entities.NameYourStruct

		idString := r.FormValue("id")
		Id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		category.Name = r.FormValue("name")
		category.UpdatedAt = time.Now()

		if ok := categorymodel.Update(Id, category); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
			return
		}
		http.Redirect(w, r, "/categories", http.StatusSeeOther)
	}
}

// This function is used to delete a data in the table
func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		panic(err)
	}

	if err := categorymodel.Delete(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/categories", http.StatusSeeOther)

}
