package bpt

import (
	"fmt"
	"log"
	"os"
)

const (
	PageSize     = 4096
	RootNode     = 0
	InternalNode = 1
	LeafNode     = 2
)

type LeafInterface interface {
	ToString() string
}

type Node struct {
	FileName string
	NodeType uint
	Key      string
	Values   string
}

func NewNode(name string, nodeType uint) *Node {
	return &Node{
		FileName: name,
		NodeType: nodeType,
	}
}

func (n *Node) isRoot() bool {
	return n.NodeType == RootNode
}

func (n *Node) isInternal() bool {
	return n.NodeType == InternalNode
}

func (n *Node) isLeaf() bool {
	return n.NodeType == LeafNode
}

func (n *Node) getDir() string {
	dir := "__index"
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		err = os.Mkdir(dir, 0744)
		if err != nil {
			log.Printf("Node getDir() error %s\n", err.Error())
		}
	}
	return dir
}

func (n *Node) insert(in []byte) {
	dir := n.getDir()
	name := fmt.Sprintf("./%s/%s", dir, n.FileName)
	var flag int

	if n.isRoot() {
		flag = os.O_CREATE | os.O_RDWR | os.O_TRUNC
	} else {
		flag = os.O_CREATE | os.O_RDWR | os.O_APPEND
	}

	file, err := os.OpenFile(name, flag, 0744)
	if err != nil {
	}
	_, err = file.Write(in)
	if err != nil {
	}
}

func (n *Node) open() error {
	dir := n.getDir()
	name := fmt.Sprintf("./%s/%s", dir, n.FileName)
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		_, err = os.Create(name)
		if err != nil {
			return err
		}
	}
	data := make([]byte, PageSize)
	file, err := os.OpenFile(name, os.O_RDONLY, 0744)
	if err != nil {
		return err
	}
	i, err := file.Read(data)
	if err != nil {
		return err
	}

	if n.isRoot() {
		n.Key = string(data[:i])
	} else if n.isInternal() {
		n.Key = string(data[:i])
	} else if n.isLeaf() {
		n.Values = string(data[:i])
	}

	return nil
}
