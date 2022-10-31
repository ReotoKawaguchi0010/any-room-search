package btree

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewLeafNode(t *testing.T) {
	ln := NewLeafNode()
	ln = ln.Add("偽善", map[interface{}]interface{}{"PostingId": 1})
	ln = ln.Add("偽善2", map[interface{}]interface{}{"PostingId": 1})
	fmt.Println(ln.Front().Value)
	fmt.Println(ln.Front().Next().Value)
}

func TestNewBPlusTreeInsert(t *testing.T) {
	btree := NewBPlusTree()
	btree.Insert("偽善", map[interface{}]interface{}{"PostingId": 1})
	btree.Insert("偽物", map[interface{}]interface{}{"PostingId": 2})
	btree.Insert("曲者", map[interface{}]interface{}{"PostingId": 3})

	if !reflect.DeepEqual(btree.LeafNode.Front().Value.(*Leaf).Key, "偽善") {
		t.Errorf("s \n")
	}

	if !reflect.DeepEqual(btree.LeafNode.Front().Next().Value.(*Leaf).Key, "偽物") {
		t.Errorf("s \n")
	}
}
