package main

import (
	"aws-api/controllers"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	// Instantiate a new router
	r := httprouter.New()

	// Get a UserController instance
	uc := controllers.NewUserController(getSession())

	// Get a user resource
	r.GET("/user/:id", uc.GetUser)

	r.GET("/user", uc.GetAllUser)

	r.POST("/user", uc.CreateUser)

	r.DELETE("/user/:id", uc.RemoveUser)

	r.POST("/user/:id", uc.UpdateUser)

	// Fire up the server
	http.ListenAndServe(":3000", r)
}

func getSession() *sql.DB {
	// Connect to docker mysql
	db, err := sql.Open("mysql", "root:supersecret@tcp(172.17.0.2:3306)/banka")

	if err != nil {
		panic(err)
	}
	return db
}
