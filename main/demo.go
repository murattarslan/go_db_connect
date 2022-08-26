package main

import (
	"fmt"
	"myApp/main/dao"
)

func demo() {
	// create table
	//dao.CreateDeskTable("restaurant_b")
	// insert item
	//var item = dao.Desk{}
	//item.SetName("masa2")
	//dao.InsertDesk("restaurant_a", item)
	// get items
	all := dao.GetAllDesk("restaurant_a")
	fmt.Print(all)
	// delete items
	//dao.Delete("restaurant_a", 1)
}
