package main

import "log"

func main() {
	cfg := NewConfig()

	sqlStorage := NewSQLStorage(cfg)

	db, err := sqlStorage.Init()
	if err != nil {
		log.Fatal(err)
	}

	storage := NewStorage(db)

	api := NewAPIServer(":4005", storage)
	api.Serve()	
}