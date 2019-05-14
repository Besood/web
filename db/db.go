package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)
type DB struct {
	id int
	name string 
	password string
}

func SelectDB() {
	db:=OpenDB()
	rows, err := db.Query("SELECT * FROM user.user_id where user.user_id.user_id = 1;")
	if err != nil {
        panic(err)
    }
    defer rows.Close()

    for rows.Next() {
		table := DB{}
        err := rows.Scan(&table.id,&table.password, &table.name)
        if err != nil {
            panic(err)
        }
        fmt.Println(table.id, table.name,table.password);
	}
	defer db.Close()
}

func OpenDB() *sql.DB{
	db, err := sql.Open("mysql", "root:methodist@tcp(127.0.0.1:3306)/user")

    if err != nil {
        panic(err)
    }
	return db
}

func InsertDB(name,password string) bool{
	db := OpenDB()
	defer db.Close()
	rows, err := db.Query("INSERT INTO user_id (username,password) VALUES (?,?)",name,password)
	if err != nil {
       return false
    }
	defer rows.Close()
	return true
}