package controllers

import (
	"aws-api/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type (
	// UserController represents the controller for operating on the User resource
	UserController struct {
		session *sql.DB
	}
)

func getSession() *sql.DB {
	// Connect to our local mongo
	db, err := sql.Open("mysql", "root:supersecret@tcp(172.17.0.2:3306)/banka")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	return db
}

func NewUserController(s *sql.DB) *UserController {
	return &UserController{s}
}

// GetUser retrieves an individual user resource
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	db := getSession()
	defer db.Close()
	// Grab id
	id := p.ByName("id")

	// Stub user
	u := models.User{}

	// Fetch user
	selDB, err := db.Query("SELECT * FROM Users WHERE id=?", id)
	if err != nil {
		panic(err.Error())
	}

	for selDB.Next() {
		var id, age int
		var name, gender string
		err = selDB.Scan(&id, &name, &gender, &age)
		if err != nil {
			panic(err.Error())
		}
		u.Id = strconv.Itoa(id)
		u.Name = name
		u.Gender = gender
		u.Age = age
	}
	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(u)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}

func (uc UserController) GetAllUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	db := getSession()
	defer db.Close()
	// Stub user
	us := []models.User{}

	// Fetch user
	selDB, err := db.Query("SELECT * FROM Users")
	if err != nil {
		panic(err.Error())
	}

	for selDB.Next() {
		u := models.User{}
		var id, age int
		var name, gender string
		err = selDB.Scan(&id, &name, &gender, &age)
		if err != nil {
			panic(err.Error())
		}
		u.Id = strconv.Itoa(id)
		u.Name = name
		u.Gender = gender
		u.Age = age
		us = append(us, u)
	}
	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(us)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}

// CreateUser creates a new user resource
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	db := getSession()
	defer db.Close()
	// Stub an user to be populated from the body
	u := models.User{}

	// Populate the user data
	json.NewDecoder(r.Body).Decode(&u)

	in, err := db.Prepare(`INSERT INTO Users(Name, Gender, Age) VALUES(?,?,?)`)
	if err != nil {
		panic(err.Error())
	}
	name := r.FormValue("name")
	gender := r.FormValue("gender")
	age := r.FormValue("age")
	in.Exec(name, gender, age)
	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(u)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", uj)
}

// RemoveUser removes an existing user resource
func (uc UserController) RemoveUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	db := getSession()
	defer db.Close()
	// Grab id
	id := p.ByName("id")

	selDB, err := db.Prepare(`DELETE * FROM Users WHERE id=?`)
	if err != nil {
		panic(err.Error())
	}
	selDB.Exec(id)
	fmt.Println("DELETED")
	// TODO: only write status for now
	w.WriteHeader(200)
}
