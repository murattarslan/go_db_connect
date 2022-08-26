package dao

import (
	"fmt"

	_ "github.com/lib/pq"
)

func UpdateDeskName(tableName string, id int, name string) {
	db, err := connect()
	if err != nil {
		panic(err)
	}
	fmt.Println("connect")
	defer db.Close()

	updateQuery := fmt.Sprintf("update %s set name='%s' where id=%v", tableName, name, id)

	_, err = db.Exec(updateQuery)
	if err != nil {
		panic(err)
	}
	fmt.Printf("complete edit")
}
