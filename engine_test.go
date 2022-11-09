package search

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var testDB *sql.DB

func setup() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/anyroom_search")
	if err != nil {
		log.Fatalln(err)
	}
	_, err = db.Exec(`TRUNCATE TABLE documents`)
	if err != nil {
		log.Fatalln(err)
	}
	if err := os.RemoveAll("_index_data"); err != nil {
		log.Fatalln(err)
	}
	if err := os.Mkdir("_index_data", 0777); err != nil {
		log.Fatalln(err)
	}
	return db
}

func TestMain(m *testing.M) {
	testDB = setup()
	defer testDB.Close()
	exitCode := m.Run()
	os.Exit(exitCode)
}
