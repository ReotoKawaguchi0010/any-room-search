package btree

import (
	"container/list"
)

type RootNode struct {
	Key          string
	InternalNode *InternalNode
}

type InternalNode struct {
	Key       string
	ChildNode *InternalNode
}

type Leaf struct {
	Key   string
	Value interface{}
}

type LeafNode struct {
	*list.List
}

func NewLeafNode() *LeafNode {
	return &LeafNode{}
}

func (ln *LeafNode) Add(key string, value interface{}) *LeafNode {
	leaf := &Leaf{
		Key:   key,
		Value: value,
	}
	if ln.List == nil {
		l := list.New()
		l.PushBack(leaf)
		return &LeafNode{l}
	}
	ln.List.PushBack(leaf)
	return &LeafNode{ln.List}
}

func (ln *LeafNode) Next() *Leaf {
	return ln.Front().Value.(*Leaf)
}

type BPlusTree struct {
	BlockSize   uint
	PageSize    uint
	AverageSize uint
	PageOffset  uint
	PointerSize uint
	RootNode    *RootNode
	LeafNode    *LeafNode
}

func NewBPlusTree() *BPlusTree {
	return &BPlusTree{
		BlockSize:   4096,
		PageSize:    4096,
		AverageSize: 10,
		PageOffset:  2,
	}
}

func (bt *BPlusTree) Insert(key string, value interface{}) {
	if bt.LeafNode == nil {
		bt.LeafNode = NewLeafNode()
	}
	bt.LeafNode = bt.LeafNode.Add(key, value)
}
