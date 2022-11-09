package search

import (
	"os"
	"testing"
)

var name = "./testdata/test.db"

func TestConnection(t *testing.T) {
	db := NewTokens()
	conn, err := db.Conn(name)
	if err != nil {
		t.Errorf("open error\n")
	}
	defer func() {
		err = conn.Close()
		if err != nil {
		}
		err = os.Remove(name)
		if err != nil {
		}
	}()
}

func TestNewTokens(t *testing.T) {
	table := "user_tokens"
	db := NewTokens()
	conn, _ := db.Conn(name)
	defer func() {
		err := db.Close()
		if err != nil {
		}
		err = os.Remove(name)
		if err != nil {
		}
	}()
	_, err := conn.Exec("DROP TABLE user_tokens;")
	if err != nil {
	}
	err = db.Create(table)
	if err != nil {
		t.Errorf("create error %v\n", err)
	}
	conn, _ = db.Conn(name)
	_, err = conn.Exec("SELECT * FROM user_tokens;")
	if err != nil {
		t.Errorf("does not exits err, %v\n", err)
	}
}

func TestTokensInsert(t *testing.T) {
	table := "user_tokens"
	db := NewTokens()
	conn, _ := db.Conn(name)
	defer func() {
		err := db.Close()
		if err != nil {
		}
		err = os.Remove(name)
		if err != nil {
		}
	}()
	_, err := conn.Exec("DROP TABLE user_tokens;")
	if err != nil {
	}
	db.Create(table)
	db.Insert("これ", 2, []byte("1,2,4,6"))
	row := conn.QueryRow("SELECT token FROM user_tokens WHERE id = ?", 1)
	var token string
	row.Scan(&token)
	if token != "これ" {
		t.Errorf("get error \n")
	}
}
