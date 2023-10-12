package main

import (
	"fmt"
	"time"

	"github.com/OseiasMzM/go-database-dummy-generator/db"
)

func main() {
	fmt.Println("Populating SQLite with dummy data...")
	db.DummyTblInternet()
	sec := time.Now()
	fmt.Println(sec.Format("20060102150405"))

}
