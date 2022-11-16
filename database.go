package search

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type TokensScheme struct {
	Token     string
	DocsCount int
	Postings  []byte
}

type Tokens struct {
	conn *sql.DB
	Name string
}

func NewTokens() *Tokens {
	return &Tokens{}
}

func (d *Tokens) Conn(name string) (*sql.DB, error) {
	conn, err := sql.Open("sqlite3", name)
	if err != nil {
		log.Printf("sql open error %v\n", err)
		return nil, err
	}
	d.conn = conn
	return conn, nil
}

func (d *Tokens) Create(name string) error {
	d.Name = name
	query := `
CREATE TABLE %s (
    token      TEXT PRIMARY KEY,
    docs_count INT NOT NULL,
    postings   BLOB NOT NULL
);
`
	query = fmt.Sprintf(query, name)
	_, err := d.conn.Exec(query)
	if err != nil {
		log.Printf("exec error %v\n", err)
		return err
	}

	return nil
}

func (d *Tokens) Close() error {
	return d.conn.Close()
}

func (d *Tokens) Insert(token string, docsCount int, postings []byte) {
	query := "INSERT INTO %s(token, docs_count, postings) VALUES(?, ?, ?);"
	query = fmt.Sprintf(query, d.Name)
	stmt, err := d.conn.Prepare(query)
	if err != nil {
		log.Printf("prepare error %v\n", err)
	}
	stmt.Exec(token, docsCount, postings)
}
