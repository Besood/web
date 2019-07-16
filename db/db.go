package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func SelectDB(username, password string) bool {
	db := OpenDB()
	rows, err := db.Query("SELECT user.user_id.password FROM user.user_id WHERE user.user_id.username = ?", username)
	if err != nil {
		log.Println("err:",err)
		return false
	}
	var pass string
	rows.Next()
	rows.Scan(&pass)
	if result := UserPassMach(pass,password); result == false{
		log.Println("ID or Password is different")
		return false
	}
	defer db.Close()
	defer rows.Close()
	return true
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
	hash,err := UserPassHash(password)
	if err != nil{
		log.Println(err)
		return false
	}
	rows, err := db.Query("INSERT INTO user.user_id (username, password) VALUES (?, ?)", name, hash)
	if err != nil {
		log.Println(err)
		return false
	}
	defer db.Close()
	defer rows.Close()
	return true
}
