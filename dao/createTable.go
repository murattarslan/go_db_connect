package dao

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func CreateDeskTable(tableName string) sql.Result {
	conf :=
		"host=" + host +
			" port=" + port +
			" user=" + user +
			" password=" + password +
			" dbname=" + dbname +
			" sslmode=disable"

	db, err := sql.Open("postgres", conf)
	if err != nil {
		panic(err)
	}
	fmt.Println("connect")
	defer db.Close()

	createQuery :=
		"CREATE TABLE IF NOT EXISTS " + tableName + " (" +
			" ID SERIAL PRIMARY KEY," +
			" NAME TEXT NOT NULL," +
			" ACTIVE BOOL NOT NULL" +
			" ); "

	r, err := db.Exec(createQuery)
	if err != nil {
		panic(err)
	}
	fmt.Println("complete")
	return r
}

func CreateRestaurantTable(tableName string) sql.Result {
	conf :=
		"host=" + host +
			" port=" + port +
			" user=" + user +
			" password=" + password +
			" dbname=" + dbname +
			" sslmode=disable"

	db, err := sql.Open("postgres", conf)
	if err != nil {
		panic(err)
	}
	fmt.Println("connect")
	defer db.Close()

	createQuery :=
		"CREATE TABLE IF NOT EXISTS " + tableName +
			" (" +
			"ID SERIAL PRIMARY KEY, " +
			"NAME TEXT NOT NULL, " +
			"ADDRESS TEXT NOT NULL," +
			")"

	r, err := db.Exec(createQuery)
	if err != nil {
		panic(err)
	}
	fmt.Println("complete")
	return r
}
