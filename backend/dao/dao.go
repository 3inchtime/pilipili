package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Dao struct {
	DB *sql.DB
}

func NewDao() *Dao {
	return &Dao{
		DB: NewDB(),
	}
}

func NewDB() *sql.DB {
	Mysql, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/pilipili?charset=utf8mb4")
	if err != nil {
		panic(err)
	}
	return Mysql
}
