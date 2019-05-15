package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	//"fmt"
)

type DB struct {
	name     string
	password string
}

func SelectDB(username, password string) (bool,string) {
	db := OpenDB()
	rows, err := db.Query("SELECT user.user_id.username, user.user_id.password FROM user.user_id WHERE user.user_id.username = ? AND user.user_id.password = ?", username, password)
	if err != nil {
		msg := "DB error"
		log.Println(err)
		return false,msg
	}
	if 	msg := rows.Next(); msg == false{
		return false ,"ID or password is different"
	}
	defer db.Close()
	defer rows.Close()
	return true,""
}

func OpenDB() *sql.DB {
	db, err := sql.Open("mysql", "root:methodist@tcp(127.0.0.1:3306)/user")

	if err != nil {
		panic(err)
	}
	return db
}

func InsertDB(name, password string) bool {
	db := OpenDB()
	defer db.Close()
	rows, err := db.Query("INSERT INTO user_id (username,password) VALUES (?,?)", name, password)
	if err != nil {
		return false
	}
	defer rows.Close()
	return true
}
