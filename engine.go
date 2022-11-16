package search

import (
	"database/sql"
	"os"
	"path/filepath"
)

type Engine struct {
	tokenizer *Tokenizer
	indexDir  string
}

func NewSearchEngine(db *sql.DB) *Engine {
	tokenizer := NewTokenizer()

	path, ok := os.LookupEnv("INDEX_DIR_PATH")
	if !ok {
		current, _ := os.Getwd()
		path = filepath.Join(current, "_index_data")
	}
	return &Engine{
		tokenizer: tokenizer,
		indexDir:  path,
	}
}
