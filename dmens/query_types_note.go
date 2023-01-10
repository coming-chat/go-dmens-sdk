package dmens

import (
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

type RepostNote struct {
	*Note        // origin note info
	Repost *Note `json:"repost"` // repost note info
}

func (p *RepostNote) JsonString() (string, error) {
	return JsonString(p)
}

func (n *RepostNote) AsAny() *base.Any {
	return &base.Any{Value: n}
}

func AsRepostNote(any *base.Any) *RepostNote {
	if res, ok := any.Value.(*RepostNote); ok {
		return res
	}
	if res, ok := any.Value.(RepostNote); ok {
		return &res
	}
	return nil
}

type NotePage struct {
	*sdkPageable[Note]
}

type RepostNotePage struct {
	*sdkPageable[RepostNote]
}
