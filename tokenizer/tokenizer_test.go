package tokenizer

import (
	"reflect"
	"testing"
)

func TestTokenizer(t *testing.T) {
	page := `
これは、test用のテキストです。
これはtokenizerで利用します。
`
	tokens := Tokenizer(page)
	test := []string{
		"これ",
		"れは",
		"は、",
		"test",
		"用の",
		"のテ",
		"テキ",
		"キス",
		"スト",
		"トで",
		"です",
		"す。",
		"これ",
		"れは",
		"tokenizer",
		"で利",
		"利用",
		"用し",
		"しま",
		"ます",
		"す。",
	}

	if !reflect.DeepEqual(tokens, test) {
		t.Errorf("token可失敗\n")
	}

}
