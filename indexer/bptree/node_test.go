package bpt

import (
	"fmt"
	"testing"
)

func TestRootNodeOpen(t *testing.T) {
	node := NewNode("root.log", RootNode)
	err := node.open()
	if err != nil {
		t.Errorf("index file not create")
	}
	fmt.Println(node.Key)
}

func TestRootNodeInsert(t *testing.T) {
	node := NewNode("root.log", RootNode)
	node.insert([]byte("test"))
	node.open()
	if node.Key != "test" {
		t.Errorf("key not insert")
	}
}

func TestLeafNodeInsert(t *testing.T) {
	node := NewNode("leaf.1.log", LeafNode)
	node.insert([]byte("1,TF,1,2,3"))
	node.open()
	if node.Values != "1,TF,1,2,3" {
		t.Errorf("node not values")
	}
}
