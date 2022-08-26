package dao

import (
	"fmt"

	_ "github.com/lib/pq"
)

func GetAllDesk(tableName string) []Desk {
	db, err := connect()
	if err != nil {
		panic(err)
	}
	fmt.Println("connect")
	defer db.Close()

	selectQuery := fmt.Sprintf("select id, name, active from %s", tableName)

	rows, err := db.Query(selectQuery)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var result []Desk

	for rows.Next() {

		var id int64
		var name string
		var active bool

		err = rows.Scan(&id, &name, &active)
		if err != nil {
			panic(err)
		}

		item := Desk{name, id, active}
		result = append(result, item)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return result
}

func GetAllRestaurant(tableName string) []Restaurant {
	db, err := connect()
	if err != nil {
		panic(err)
	}
	fmt.Println("connect")
	defer db.Close()

	selectQuery := fmt.Sprintf("select id, name, address from %s", tableName)

	rows, err := db.Query(selectQuery)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var result []Restaurant

	for rows.Next() {

		var id int64
		var name string
		var address string

		err = rows.Scan(&id, &name, &address)
		if err != nil {
			panic(err)
		}

		item := Restaurant{name, id, address}
		result = append(result, item)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return result
}
