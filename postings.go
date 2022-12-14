package search

import (
	"container/list"
	"encoding/json"
	"fmt"
	"strings"
)

type DocumentID int64

type Posting struct {
	DocID         DocumentID
	Positions     []int
	TermFrequency int
}

func NewPosting(docID DocumentID, positions ...int) *Posting {
	return &Posting{
		DocID:         docID,
		Positions:     positions,
		TermFrequency: len(positions),
	}
}

func (p *Posting) String() string {
	return fmt.Sprintf("(%v, %v, %v)",
		p.DocID, p.TermFrequency, p.Positions)
}

type PostingsList struct {
	*list.List
}

func NewPostingsList(postings ...*Posting) PostingsList {
	l := list.New()
	for _, posting := range postings {
		l.PushBack(posting)
	}
	return PostingsList{l}
}

func (pl *PostingsList) add(p *Posting) {
	pl.PushBack(p)
}

func (pl *PostingsList) last() *Posting {
	e := pl.List.Back()
	if e == nil {
		return nil
	}
	return e.Value.(*Posting)
}

func (pl *PostingsList) Add(new *Posting) {
	last := pl.last()
	if last == nil || last.DocID != new.DocID {
		pl.add(new)
		return
	}
	last.Positions = append(last.Positions, new.Positions...)
	last.TermFrequency++
}

func (pl *PostingsList) String() string {
	str := make([]string, 0, pl.Len())
	for e := pl.Front(); e != nil; e = e.Next() {
		str = append(str, e.Value.(*Posting).String())
	}
	return strings.Join(str, "=>")
}

func (pl *PostingsList) MarshallJSON() ([]byte, error) {

	postings := make([]*Posting, 0, pl.Len())

	for e := pl.Front(); e != nil; e = e.Next() {
		postings = append(postings, e.Value.(*Posting))
	}
	return json.Marshal(postings)
}

func (pl *PostingsList) UnmarshallJSON(b []byte) error {

	var postings []*Posting
	if err := json.Unmarshal(b, &postings); err != nil {
		return err
	}
	pl.List = list.New()
	for _, posting := range postings {
		pl.add(posting)
	}
	return nil

}
