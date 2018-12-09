package mysql_layer

import (
	"aws-api/persistence"
	"database/sql"
)

type MySqlLayer struct {
	session *sql.DB
}

func NewMySqlLayer(connection string) (persistence.DatabaseHandler, error) {
	db, err := sql.Open("mysql", connection)
	// Check if connection error, is mysql running?
	return &MySqlLayer{
		session: db,
	}, err
}

func (mySqlLayer *MySqlLayer) AddUser(name, gender string, age int) (int64, error) {
	db := mySqlLayer.getSession()
	defer db.Close()

	in, err := db.Prepare(`INSERT INTO Users(Name, Gender, Age) VALUES(?,?,?)`)
	if err != nil {
		return 0, err
	}

	res, errRes := in.Exec(name, gender, age)
	if errRes != nil {
		uid, erruid := res.LastInsertId()
		if erruid != nil {
			return uid, erruid
		}
	}
	return 0, err
}

func (mySqlLayer *MySqlLayer) GetUser(uid int) (persistence.User, error) {
	db := mySqlLayer.getSession()
	defer db.Close()

	u := persistence.User{}

	selDB, err := db.Query("SELECT * FROM Users WHERE id=?", uid)
	if err != nil {
		return u, err
	}

	for selDB.Next() {
		var id, age int
		var name, gender string
		err = selDB.Scan(&id, &name, &gender, &age)
		if err != nil {
			return u, err
		}
		u.Id = id
		u.Name = name
		u.Gender = gender
		u.Age = age
	}

	return u, nil
}

func (mySqlLayer *MySqlLayer) GetAllUsers() ([]persistence.User, error) {
	db := mySqlLayer.getSession()
	defer db.Close()

	us := []persistence.User{}

	// Fetch users
	selDB, err := db.Query("SELECT * FROM Users")
	if err != nil {
		return nil, err
	}

	for selDB.Next() {
		u := persistence.User{}
		var id, age int
		var name, gender string
		err = selDB.Scan(&id, &name, &gender, &age)
		if err != nil {
			return nil, err
		}
		u.Id = id
		u.Name = name
		u.Gender = gender
		u.Age = age
		us = append(us, u)
	}

	return us, nil
}

func (mySqlLayer *MySqlLayer) DeleteUser(uid int) (err error) {
	db := mySqlLayer.getSession()
	defer db.Close()

	selDB, err := db.Prepare(`DELETE FROM Users WHERE id=?`)
	if err != nil {
		panic(err.Error())
	}
	_, err = selDB.Exec(uid)
	return
}

func (mySqlLayer *MySqlLayer) UpdateUser(uid, age int, name string) (err error) {
	db := mySqlLayer.getSession()
	defer db.Close()

	in, err := db.Prepare(`UPDATE Users SET Name=?, Age=? WHERE Id=?`)
	if err != nil {
		return
	}
	_, err = in.Exec(name, age, uid)
	return
}

func (mySqlLayer *MySqlLayer) getSession() *sql.DB {
	db, err := sql.Open("mysql", "root:supersecret@tcp(172.17.0.2:3306)/banka")
	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	return db
}
