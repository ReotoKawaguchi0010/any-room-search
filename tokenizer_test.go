package search

import (
	"testing"
)

func TestJapaneseTokenizer(t *testing.T) {
	var tokens []string
	s := `これは、日本語のトークンのテスト用テキストです。`
	tokens = BiGram(s)

	r := []string{
		"これ",
		"れは",
		"は日",
		"日本",
		"本語",
		"語の",
		"のト",
		"トー",
		"ーク",
		"クン",
		"ンの",
		"のテ",
		"テス",
		"スト",
		"ト用",
		"用テ",
		"テキ",
		"キス",
		"スト",
		"トで",
		"です",
	}

	if len(tokens) != len(r) {
		t.Errorf("token create error \n")
	}

}
