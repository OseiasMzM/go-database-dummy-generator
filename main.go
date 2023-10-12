package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Populating SQLite with dummy data...")
	// db.DummyTblInternet()
	fmt.Println(rand.Intn(9999999-0) + 0)
	rand.Seed(time.Now().Unix())
	reasons := []string{
		"Locked out",
		"Pipes broke",
		"Food poisoning",
		"Not feeling well",
		"Locked 1",
		"Pipes 1",
		"Food 1",
		"Not feeling 1",
	}
	for i := 0; i < 100; i++ {
		n := rand.Int() % len(reasons)
		fmt.Print("\nGonna work from home...", reasons[n])
	}
}
