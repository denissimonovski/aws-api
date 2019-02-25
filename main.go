package main

import (
	"aws-api/controllers"
	"aws-api/persistence/dblayer"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	// Iniciranje na nov ruter
	r := httprouter.New()

	// Iniciranje sesija do baza
	dbhandle, err := dblayer.NewPersistenceLayer(dblayer.MYSQLDB, "root:supersecret@tcp(172.17.0.2:3306)/banka")
	if err != nil {
		fmt.Printf("Can't connect to db: %s", err.Error())
	}
	// Iniciranje na kontroler za korisnici
	uc := controllers.NewUserHandler(dbhandle)
	// Get a user resource
	r.GET("/user/:id", uc.GetUser)

	r.GET("/user", uc.GetAllUser)

	r.POST("/user", uc.CreateUser)

	r.DELETE("/user/:id", uc.RemoveUser)

	r.POST("/user/:id", uc.UpdateUser)

	// Start na server
	http.ListenAndServe(":3000", r)
}
