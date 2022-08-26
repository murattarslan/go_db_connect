package dao

import (
	"fmt"

	_ "github.com/lib/pq"
)

func Delete(tableName string, id int) {
	db, err := connect()
	if err != nil {
		panic(err)
	}
	fmt.Println("connect")
	defer db.Close()

	insertQuery := fmt.Sprintf("delete from %s where id=%v", tableName, id)

	_, err = db.Exec(insertQuery)
	if err != nil {
		panic(err)
	}
	fmt.Printf("complete delete")
}
