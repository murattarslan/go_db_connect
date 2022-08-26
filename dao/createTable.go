package dao

import (
	"fmt"

	_ "github.com/lib/pq"
)

func CreateDeskTable(tableName string) {

	db, err := connect()
	if err != nil {
		panic(err)
	}
	fmt.Println("connect")
	defer db.Close()

	createQuery := fmt.Sprintf("create table if not exists %s (id serial primary key, name text not null, active bool not null)", tableName)

	r, err := db.Exec(createQuery)
	if err != nil && r == nil {
		panic(err)
	}
	fmt.Println("complete")
}

func CreateRestaurantTable(tableName string) {

	db, err := connect()
	if err != nil {
		panic(err)
	}
	fmt.Println("connect")
	defer db.Close()

	createQuery := fmt.Sprintf("create table if not exists %s (id serial primary key, name text not null, address text not null)", tableName)

	r, err := db.Exec(createQuery)
	if err != nil && r == nil {
		panic(err)
	}
	fmt.Println("complete")
}
