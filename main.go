package main

import (
	// "fmt"
	"ccnt/config"
	"ccnt/controllers/categorycontrollers"
	"ccnt/controllers/homecontrollers"
	"log"
	"net/http"
)

// type duct struct{

// }

func main() {
	config.ConnectDB()

	http.HandleFunc("/", homecontrollers.Welcome)

	// Func Categories
	http.HandleFunc("/categories", categorycontrollers.Index)
	http.HandleFunc("/categories/Add", categorycontrollers.Add)
	// http.HandleFunc("/categories/Edit", categorycontrollers.Edit)
	// http.HandleFunc("/categories/Delete", categorycontrollers.Delete)

	log.Printf("listen and serve my brody")
	http.ListenAndServe(":8080", nil)
}
