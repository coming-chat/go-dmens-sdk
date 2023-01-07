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
	return nil
}

type NotePage struct {
	Notes         []*Note `json:"notes"`
	CurrentCursor string  `json:"currentCursor"`
	CurrentCount  int     `json:"currentCount"`
	TotalCount    int     `json:"totalCount"`

	notesArray *base.AnyArray
}

func (p *NotePage) JsonString() (string, error) {
	return JsonString(p)
}

func (p *NotePage) NoteArray() *base.AnyArray {
	if p.notesArray == nil {
		a := make([]any, len(p.Notes))
		for _, n := range p.Notes {
			a = append(a, n)
		}
		p.notesArray = &base.AnyArray{Values: a}
	}
	return p.notesArray
}

type rawFieldsId struct {
	Id struct {
		Id string `json:"id"`
	} `json:"id"`
}

type rawNote struct {
	CreateTime string `json:"createTime,omitempty"`
	ObjectId   string `json:"objectId,omitempty"`

	Fields *struct {
		// rawFieldsId
		// Name  string `json:"name"`
		Value struct {
			// Type string `json:"type"`
			Fields struct {
				Action NoteAction `json:"action"`
				Text   string     `json:"text"`
				Poster string     `json:"poster"`
				RefId  string     `json:"ref_id"`

				// rawFieldsId
				// Url    string     `json:"url"`
				// AppId  int        `json:"app_id"`
			} `json:"fields"`
		} `json:"value"`
	} `json:"fields,omitempty"`

	// Owner      interface{} `json:"owner,omitempty"`
	// UpdateTime string `json:"updateTime,omitempty"`
	// Status     string `json:"status,omitempty"`
	// DataType   string `json:"dataType,omitempty"`
	// Type       string `json:"type,omitempty"`
	// NodeId     string `json:"nodeId,omitempty"`
	// Digest     string `json:"digest,omitempty"`
	// Version    string `json:"version,omitempty"`
	// IsUpdate   bool   `json:"isUpdate,omitempty"`
	// StorageRebate       string `json:"storageRebate,omitempty"`
	// PreviousTransaction string `json:"previousTransaction,omitempty"`
	// HasPublicTransfer   bool   `json:"hasPublicTransfer,omitempty"`
}

type rawNotePage struct {
	TotalCount int `json:"totalCount,omitempty"`
	Edges      []struct {
		Node   rawNote `json:"node"`
		Cursor string  `json:"cursor"`
	} `json:"edges,omitempty"`
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
func (a *rawNotePage) MapToNotePage(poster *Poster) *NotePage {
	length := len(a.Edges)
	if length == 0 {
		return &NotePage{
			TotalCount: a.TotalCount,
		}
	}
	notes := make([]*Note, 0)
	for _, n := range a.Edges {
		notes = append(notes, n.Node.MapToNote())
	}
	page := &NotePage{
		TotalCount:    a.TotalCount,
		Notes:         notes,
		CurrentCount:  len(notes),
		CurrentCursor: a.Edges[length-1].Cursor,
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
