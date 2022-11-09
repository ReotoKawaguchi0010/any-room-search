package search

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"unicode"
)

type Tokenizer struct {
}

func NewTokenizer() *Tokenizer {
	return &Tokenizer{}
}

func BiGram(text string) []string {
	var tokens []string
	var token []rune
	text = strings.ReplaceAll(text, "\n", "")
	text = strings.ReplaceAll(text, "\r", "")
	text = strings.ReplaceAll(text, "、", "")
	text = strings.ReplaceAll(text, "。", "")
	chars := []rune(text)

	for i := 0; i < len(chars); i++ {
		token = []rune{}
		var bc rune
		var ac rune
		if i == 0 {
			continue
		} else {
			bc = chars[i-1]
			ac = chars[i]
		}
		token = append(token, bc)
		token = append(token, ac)
		tokens = append(tokens, string(token))
	}

	return tokens
}

func replace(r rune) rune {
	if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') && !unicode.IsNumber(r) {
		return -1
	}
	return unicode.ToLower(r)
}

func (t *Tokenizer) SplitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {
	advance, token, err = bufio.ScanWords(data, atEOF)

	if err == nil && token != nil {
		token = bytes.Map(replace, token)
		if len(token) == 0 {
			token = nil
		}
	}
	return advance, token, err
}

func (t *Tokenizer) TextToWordSequence(text string) []string {
	scanner := bufio.NewScanner(strings.NewReader(text))
	scanner.Split(t.SplitFunc)
	var result []string
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		result = append(result, scanner.Text())
	}
	return result
}
