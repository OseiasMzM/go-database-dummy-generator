package main

import (
	"fmt"

	"github.com/OseiasMzM/go-database-dummy-generator/db"
)

func main() {
	fmt.Println("Populating SQLite with dummy data...")
	db.DummyTblInternet()
}
