package main

import (
	"myApp/main/dao"
)

func demo() {
	//dao.CreateDeskTable("restaurant_a")
	var item = dao.Desk{}
	item.SetName("masa2")
	dao.InsertDesk("restaurant_a", item)
}
