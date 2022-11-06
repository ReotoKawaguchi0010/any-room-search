package bpt

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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
	if root.Key != "曲者\n" {
		t.Errorf("not root")
	}
	leaf := NewNode("leaf_1.log", LeafNode)
	err = leaf.open()
	if err != nil {
		t.Errorf("not open")
	}
	if leaf.Values != "偽善:1\n偽物:2\n曲者:3\n" {
		t.Errorf("not insert")
	}

	r, _ := ioutil.ReadFile("__index/root.log")
	i, _ := ioutil.ReadFile("__index/internal1.log")
	l, _ := ioutil.ReadFile("__index/leaf_1.log")
	separate := strings.Repeat("=", 25)
	fmt.Println(separate)
	fmt.Println("root.log")
	fmt.Println(strings.Repeat("*", 25))
	fmt.Println(string(r))
	fmt.Println(separate)
	fmt.Println("internal1.log")
	fmt.Println(strings.Repeat("*", 25))
	fmt.Println(string(i))
	fmt.Println(separate)
	fmt.Println("leaf_1.log")
	fmt.Println(strings.Repeat("*", 25))
	fmt.Println(string(l))
	fmt.Println(separate)

	os.Remove("__index/root.log")
	os.Remove("__index/internal1.log")
	os.Remove("__index/leaf_1.log")

}
