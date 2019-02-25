package controllers

import (
	"aws-api/persistence"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type userServiceHandler struct {
	dbhandler persistence.DatabaseHandler
}

func NewUserHandler(databasehandler persistence.DatabaseHandler) *userServiceHandler {
	return &userServiceHandler{
		dbhandler: databasehandler,
	}
}

// GetUser zema eden korisnik
func (eh *userServiceHandler) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	u, err := eh.dbhandler.GetUser(StoI(id))
	if err != nil {
		panic(err.Error())
	}
	// Marshal vcituva Json struktura
	uj, _ := json.Marshal(u)

	// Setiranje na header-i i vnesuvanje na json strukturata
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}

func (eh *userServiceHandler) GetAllUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	us, err := eh.dbhandler.GetAllUsers()
	if err != nil {
		panic(err.Error())
	}
	// Marshal vcituva Json struktura
	uj, _ := json.Marshal(us)

	// Setiranje na header-i i vnesuvanje na json strukturata
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}

// CreateUser kreira nov korisnik
func (eh *userServiceHandler) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	name := r.FormValue("name")
	gender := r.FormValue("gender")
	age := r.FormValue("age")
	uid, err := eh.dbhandler.AddUser(name, gender, StoI(age))
	if err != nil {
		panic(err.Error())
	}
	u := persistence.User{
		Id:     int(uid),
		Name:   name,
		Gender: gender,
		Age:    StoI(age),
	}
	// Marshal vcituva Json struktura
	uj, _ := json.Marshal(u)

	// Setiranje na header-i i vnesuvanje na json strukturata
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", uj)
}

func (eh *userServiceHandler) UpdateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	id := p.ByName("id")
	name := r.FormValue("name")
	age := r.FormValue("age")
	err := eh.dbhandler.UpdateUser(StoI(id), StoI(age), name)
	if err != nil {
		panic(err.Error())
	}
	u, _ := eh.dbhandler.GetUser(StoI(id))
	// Marshal vcituva Json struktura
	uj, _ := json.Marshal(u)

	// Setiranje na header-i i vnesuvanje na json strukturata
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", uj)
}

// RemoveUser brise eden korisnik
func (eh *userServiceHandler) RemoveUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// Vcituva ID od GET requestot
	id := p.ByName("id")

	eh.dbhandler.DeleteUser(StoI(id))
	fmt.Println("DELETED")
	w.WriteHeader(200)
}

func StoI(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err.Error())
		return 0
	}
	return i
}
