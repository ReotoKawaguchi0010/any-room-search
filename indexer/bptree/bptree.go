package bpt

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
	}

	leaf := NewNode("leaf_1.log", LeafNode)
	leaf.insert([]byte(value.ToString()))

}

func (bpt *BPlusTree) search(value string) {
	root := NewNode("root.log", RootNode)
	key := []byte(root.Key)
	v := []byte(value)

}

func (bpt *BPlusTree) save() {}
