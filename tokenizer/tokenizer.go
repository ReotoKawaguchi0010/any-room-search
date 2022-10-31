package tokenizer

import (
	"fmt"
	"strings"
)

func Tokenizer(text string) []string {

	var tokens []string
	var token string
	text = strings.ReplaceAll(text, "\n", "")
	text = strings.ReplaceAll(text, "\r", "")
	for _, c := range text {
		fmt.Println(token)
		if len(token) == 2 {
			token = ""
			continue
		}
		token = fmt.Sprintf("%s%s", token, string(c))
		tokens = append(tokens, token)
	}

	return []string{}
}
