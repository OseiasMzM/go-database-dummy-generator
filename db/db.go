package db

import (
	"database/sql"
	"log"
	"math/rand"

	_ "github.com/mattn/go-sqlite3"
)

const file string = "ClientData2.mkl"

/*type table_internet struct {
	id          uint32
	timestamp   uint32
	computer    string
	computer_id uint8
	users       string
	user_id     uint8
	alloweds    int
	is_idle     uint8
	duration    uint8
	app         string
	window_text string
	url         string
	host        string
}*/

func DummyTblInternet() {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 50000; i++ {
		idRandom := getRandomNumber()
		// Executar uma consulta de inserção
		result, err := db.Exec("INSERT INTO table_internet (campo1, campo2) VALUES (?, ?)", idRandom, 42)
		if err != nil {
			log.Fatal(err)
		}

		// Obter o ID da linha inserida, se necessário
		lastInsertID, err := result.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("ID da última linha inserida: %d", lastInsertID)
	}
}

func getRandomNumber() int {
	random := rand.Intn(9999999-0) + 0
	return random
}
