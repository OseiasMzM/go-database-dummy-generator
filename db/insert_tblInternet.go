package db

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const file string = "./../ClientData2.mkl"

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

	apps := []string{
		"chrome.exe",
		"firefox.exe",
		"iexplore.exe",
		"msedge.exe",
		"opera.exe",
	}

	window_text := []string{
		"WhatsApp - Google Chrome",
		"WhatsApp — Mozilla Firefox",
		"Internet Explorer 11",
		"Sem nome - Pessoal — Microsoft? Edge",
		"WhatsApp - Opera",
	}
	computer := []string{
		"EX1CO04",
		"EX2CO01",
		"EX2RE01",
		"EX2CB01",
		"CA3CO03",
	}

	user := []string{
		"user",
		"dell",
		"Omniman",
		"CapitapPatria",
		"capitalcar",
	}

	host := []string{
		"web.whatsapp.com",
		"web.whatsapp.com",
		"portal.omie.com.br",
		"europa.hinova.com.br",
		"web.whatsapp.com",
	}

	url := []string{
		"web.whatsapp.com",
		"web.whatsapp.com",
		"portal.omie.com.br/view/2023-09-04/72b264edf15b1ebc9d81622aeb18d5b74db3db0b?a=7hxz5c8&d=4ny78rdq1kfgw&t=r",
		"https://europa.hinova.com.br/sga/sgav4_exclusive/veiculo/consultarVeiculo.php?key=MjAyMDk1NTU4MTUy",
		"web.whatsapp.com",
	}
	for i := 0; i < 2; i++ {
		idRandom := getRandomID()
		computerIdRandom := getRandomComputerID
		userIdRandom := getRandomUserID
		unixTime := time.Now().Format("20060102150405")

		appsRandom := rand.Int() % len(apps)
		fmt.Print("\n", apps[appsRandom])

		windowTextRandom := rand.Int() % len(window_text)
		fmt.Print("\n", window_text[windowTextRandom])

		computerRandom := rand.Int() % len(computer)
		fmt.Print("\n", computer[computerRandom])

		userRandom := rand.Int() % len(user)
		fmt.Print("\n", user[userRandom])

		hostRandom := rand.Int() % len(host)
		fmt.Print("\n", host[hostRandom])

		urlRandom := rand.Int() % len(url)
		fmt.Print("\n", url[urlRandom])

		result, err := db.Exec(`
		INSERT INTO table_internet
			(id,
			timestamp,
			computer,
			computer_id,
			users,
			user_id,
			alloweds,
			is_idle,
			duration,
			app,
			window_text,
			url,
			host)
		VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?) `,
			idRandom, unixTime, computerRandom, computerIdRandom, userRandom, userIdRandom, appsRandom, windowTextRandom, urlRandom, hostRandom)

		if err != nil {
			log.Fatal(err)
		}

		lastInsertID, err := result.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("ID da última linha inserida: %d", lastInsertID)
	}
}

func getRandomID() int {
	random := rand.Intn(9999999-0) + 0
	return random
}
func getRandomComputerID() int {
	random := rand.Intn(9999999-0) + 0
	return random
}

func getRandomUserID() int {
	random := rand.Intn(9999999-0) + 0
	return random
}
