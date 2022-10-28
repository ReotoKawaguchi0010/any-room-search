package indexer

import "container/list"

type DocumentID uint64

type Posting struct {
	DocID         DocumentID
	Positions     []int
	TermFrequency int
}

type PostingList struct {
	*list.List
}
