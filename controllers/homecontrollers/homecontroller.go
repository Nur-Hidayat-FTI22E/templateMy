package homecontrollers

import (
	"net/http"
	"text/template"
)

// Welcome is a function to render home page
// it will render views/home/index.html
// and return the html to the client
func Welcome(w http.ResponseWriter, r *http.Request){
	temp, err := template.ParseFiles("views/home/index.html")

	if err != nil {
		panic(err)
	}

	temp.Execute(w, nil)
}