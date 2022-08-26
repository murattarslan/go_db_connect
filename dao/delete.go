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

	deleteQuery := fmt.Sprintf("delete from %s where id=%v", tableName, id)

	_, err = db.Exec(deleteQuery)
	if err != nil {
		panic(err)
	}
	fmt.Printf("complete delete")
}
