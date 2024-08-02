package main

import (
	"fmt"
	"go-rest-api/db"
	"go-rest-api/model"
)

func main() {
	// DBのインスタンスのアドレスを受け取る
	dbConn := db.NewDB()
	// defer：後入れ先出し法(stack)で処理される
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}
