package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// DB is a global variable to store the connection to the database
// It is of type *sql.DB
var DB *sql.DB

// ConnectDB is a function to connect to the database
// It uses the Open function from the database/sql package
func ConnectDB()  {

	// Open the connection to the database
	// The first parameter is the name of the RDBMS
	// The second parameter is the connection string
	db, err := sql.Open("NameofRDBMS", "root@/YourDatabaseName")
	if err != nil {
		panic(err)
	}
	
	// Print a log if the connection is successful
	log.Println("koneksi Database berhasil")
	
	// Assign the connection to the global variable DB
	DB = db
}