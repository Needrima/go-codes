package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"log"
)

var db *sql.DB

type user struct {
	Firstname, Lastname, Username, Email, Password string
}

func main() {
	//"root:demopassword@tcp(127.0.0.1:3306)/usersdb" follows the pattern
	//"name_of_connection:password@tcp("connection_host")/database_name"
	var err error
	db, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/dbname")
	if err != nil {
		fmt.Println("Could not connect to mysql workbench:", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Println(err)
		return
	}
	query := `select * from signup where Username= ?;`

	row := db.QueryRow(query, "user")
	var u user

	err = row.Scan(&u.Firstname, &u.Lastname, &u.Username, &u.Email, &u.Password)

	if err != nil {
		fmt.Println("Error getting row")
	}

	fmt.Println(u)

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte("passw"))
	if err != nil {
		fmt.Println("Do not match")
	} else {
		fmt.Println("match")
	}
}
