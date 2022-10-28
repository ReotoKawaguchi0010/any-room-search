package indexer

import "crypto/sha256"

func hash(text string) string {
	h := sha256.New()
	h.Write([]byte(text))
	return string(h.Sum(nil))
}

type RootNode struct {
	Key          string
	InternalNode *InternalNode
}

type InternalNode struct {
	Key       string
	ChildNode *InternalNode
}

type LeafNode struct {
}

type BPlusTree struct {
	BlockSize   uint
	PageSize    uint
	AverageSize uint
	PageOffset  uint
	PointerSize uint
	RootNode    *RootNode
}

func NewBPlusTree() *BPlusTree {
	return &BPlusTree{
		BlockSize:   4096,
		PageSize:    4096,
		AverageSize: 10,
		PageOffset:  2,
	}
}
