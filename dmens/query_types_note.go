package dmens

import (
	"time"

	"github.com/coming-chat/wallet-SDK/core/base"
)

type Note struct {
	CreateTime int64  `json:"createTime"`
	NoteId     string `json:"noteId"`

	Action NoteAction `json:"action"`
	Text   string     `json:"text"`
	Poster string     `json:"poster"`
	RefId  string     `json:"refId"`

	Status *NoteStatus `json:"status"`
}

func (n *Note) JsonString() (string, error) {
	return JsonString(n)
}

func (n *Note) AsAny() *base.Any {
	return &base.Any{Value: n}
}

func AsNote(any *base.Any) *Note {
	if res, ok := any.Value.(*Note); ok {
		return res
	}
	if res, ok := any.Value.(Note); ok {
		return &res
	}
	return nil
}

type NotePage struct {
	*Pageable
	Notes []*Note `json:"notes"`
}

func (p *NotePage) JsonString() (string, error) {
	return JsonString(p)
}

func (p *NotePage) NoteArray() *base.AnyArray {
	if p.anyArray == nil {
		a := make([]any, len(p.Notes))
		for idx, n := range p.Notes {
			a[idx] = n
		}
		p.anyArray = &base.AnyArray{Values: a}
	}
	return p.anyArray
}

func (a *rawNote) MapToNote() *Note {
	var timestamp int64 = 0
	t, err := time.Parse("2006-01-02T15:04:05.999999", a.CreateTime)
	if err == nil {
		timestamp = t.Unix()
	}
	fields := a.Fields.Value.Fields
	return &Note{
		CreateTime: timestamp,
		NoteId:     a.ObjectId,
		Action:     fields.Action,
		Text:       fields.Text,
		Poster:     fields.Poster,
		RefId:      fields.RefId,
	}
}

// MapToNotePage
// @param poster If you need to query the status of notes in batches, you need to provide the poster.
func (a *rawNotePage) MapToNotePage(poster *Poster, pageSize int) *NotePage {
	notes := make([]*Note, len(a.Edges))
	for idx, n := range a.Edges {
		notes[idx] = n.Node.MapToNote()
	}
	basePage := a.mapToBasePage(pageSize)
	page := &NotePage{
		Pageable: basePage,
		Notes:    notes,
	}
	if poster != nil {
		_ = poster.BatchQueryNoteStatus(page, "")
	}
	return page
}

func (a *rawNotePage) FirstObject() *rawNote {
	if len(a.Edges) <= 0 {
		return nil
	}
	return &a.Edges[0].Node
}
