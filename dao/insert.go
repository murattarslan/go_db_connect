package dao

import (
	"fmt"

	_ "github.com/lib/pq"
)

func InsertDesk(tableName string, item Desk) int {
	db, err := connect()
	if err != nil {
		panic(err)
	}
	fmt.Println("connect")
	defer db.Close()

	insertQuery := fmt.Sprintf("insert into %s (name, active) values ('%s', %v) returning id;", tableName, item.name, item.active)

	id := 0
	err = db.QueryRow(insertQuery).Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Printf("add item. id:%d", id)
	return id
}

func InsertRestaurant(tableName string, item Restaurant) int {

	db, err := connect()
	if err != nil {
		panic(err)
	}
	fmt.Println("connect")
	defer db.Close()

	insertQuery := fmt.Sprintf("insert into %s (name, address) values ('%s', %s) returning id;", tableName, item.name, item.address)

	id := 0
	err = db.QueryRow(insertQuery).Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Printf("add item. id:%d", id)
	return id
}
