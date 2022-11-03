package bpt

import (
	"os"
	"testing"
)

type Postings struct {
	Key   string
	Value string
}

func NewPostings(key, value string) *Postings {
	return &Postings{
		Key:   key,
		Value: value,
	}
}

func (p *Postings) ToString() string {
	return p.Value
}

func TestNewBPlusTree(t *testing.T) {

	btree := NewBPlusTree()
	btree.Insert("偽善", NewPostings("PostingId", "1"))
	btree.Insert("曲者", NewPostings("PostingId", "3"))
	btree.Insert("偽物", NewPostings("PostingId", "2"))
	_, err := os.Stat("__index")
	if os.IsNotExist(err) {
		t.Errorf("error \n")
	}

	root := NewNode("root.log", RootNode)
	err = root.open()
	if err != nil {
		t.Errorf("open error")
	}
	if root.Key != "曲者" {
		t.Errorf("not root")
	}
	leaf := NewNode("leaf_1.log", LeafNode)
	err = leaf.open()
	if err != nil {
		t.Errorf("not open")
	}
	if root.Values != "1\n2\n3" {
		t.Errorf("not insert")
	}

}
