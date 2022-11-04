package bpt

import "fmt"

func IsHigher(x, y string) bool {
	rx := []rune(x)
	ry := []rune(y)

	for i, v := range rx {
		if ry[i] > v {
			return true
		}

	}
	return false
}

type BPlusTree struct {
	AverageSize uint
	PageOffset  uint
	PointerSize uint
}

func NewBPlusTree() *BPlusTree {
	return &BPlusTree{
		AverageSize: 10,
		PageOffset:  2,
	}
}

func (bpt *BPlusTree) Insert(key string, value LeafInterface) {
	root := NewNode("root.log", RootNode)
	err := root.open()
	if err != nil {
	}
	if root.Key == "" {
		root.insert([]byte(key))
	} else {
		var insertRoot []byte
		var insertInternal []byte

		if IsHigher(root.Key, key) {
			insertRoot = []byte(key)
			insertInternal = []byte(root.Key)
		} else {
			insertRoot = []byte(root.Key)
			insertInternal = []byte(key)
		}
		root.insert(insertRoot)
		internal := NewNode("internal1.log", InternalNode)
		internal.insert(insertInternal)
	}

	leaf := NewNode("leaf_1.log", LeafNode)
	v := fmt.Sprintf("%s:%s", key, value.ToString())
	leaf.insert([]byte(v))

}

func (bpt *BPlusTree) search(value string) {
	root := NewNode("root.log", RootNode)
	key := []byte(root.Key)
	v := []byte(value)
	fmt.Println(key)
	fmt.Println(v)

}
