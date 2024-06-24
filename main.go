package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
)

func main() {
	cfg := mysql.Config{
		User: Envs.DBUser,
		Passwd: Envs.DBPasswrod,
		Addr: Envs.DBAddress,
		DBName: Envs.DBName,
		Net: "tcp",
		AllowNativePasswords: true,
		ParseTime: true,
	}

	sqlStorage := NewMySQLStorage(cfg)
	db, err := sqlStorage.Init()
	if err != nil {
		log.Fatal(err)
	}

	// DOUBT : what is the diff bt DB and Store? 
	// DOUBT : where and how is the DB instance passed in?

	store := NewStore(db)

	api := NewAPIServer(":8080", store)
	api.Serve()
}