package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/shopspring/decimal"
)

var db *sqlx.DB

func init() {
	dsn := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True"
	// 也可以使用MustConnect连接不成功就panic
	var err error
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return
}

type Book struct {
	Id     int             `db:"id"`
	Title  string          `db:"title"`
	Author string          `db:"author"`
	Price  decimal.Decimal `db:"price"`
}

func queryByPrice(price decimal.Decimal) ([]Book, error) {
	sqlStr := "SELECT id,title,author,price FROM books t WHERE t.price>?"
	var books []Book
	err := db.Select(&books, sqlStr, price)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return nil, err
	}
	fmt.Printf("books:%#v\n", books)

	return books, nil
}

func main() {
	bookList, err := queryByPrice(decimal.NewFromFloat(50))
	if err != nil {
		fmt.Printf("queryByPrice failed, err:%v\n", err)
	}
	fmt.Println("length = ", len(bookList))
	fmt.Printf("bookList:%v\n", bookList)
}
