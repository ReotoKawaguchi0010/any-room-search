package indexer

import "crypto/sha256"

func hash(text string) string {
	h := sha256.New()
	h.Write([]byte(text))
	return string(h.Sum(nil))
}
