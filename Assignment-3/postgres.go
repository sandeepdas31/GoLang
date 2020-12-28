package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type products struct {
	proid   int
	proname string
	price   float64
}

func main() {
	connstr := "user=postgres dbname=product password=Sandip@123 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection established successfully")

	insertstmt := "Insert into product (proname, price) values ('keyboard', 140);"
	_, err = db.Exec(insertstmt)
	if err != nil {
		fmt.Println("Near Insert statement")
		panic(err)
	}

	var user products
	row, err := db.Query("select * from product;")
	if err != nil {
		panic(err)
	}

	for row.Next() {

		err = row.Scan(&user.proid, &user.proname, &user.price)

		fmt.Println(user.proid, user.proname, user.price, "\n")
	}
}
