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

func (n *Note) JsonString() (*base.OptionalString, error) {
	return base.JsonString(n)
}
func NewNoteWithJsonString(str string) (*Note, error) {
	var o Note
	err := base.FromJsonString(str, &o)
	return &o, err
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

func (n *RepostNote) JsonString() (*base.OptionalString, error) {
	return base.JsonString(n)
}
func NewRepostNoteWithJsonString(str string) (*RepostNote, error) {
	var o RepostNote
	err := base.FromJsonString(str, &o)
	return &o, err
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

func NewNotePageWithJsonString(str string) (*NotePage, error) {
	var o sdkPageable[Note]
	err := base.FromJsonString(str, &o)
	if err != nil {
		return nil, err
	}
	return &NotePage{&o}, nil
}

type RepostNotePage struct {
	*sdkPageable[RepostNote]
}

func NewRepostNotePageWithJsonString(str string) (*RepostNotePage, error) {
	var o sdkPageable[RepostNote]
	err := base.FromJsonString(str, &o)
	if err != nil {
		return nil, err
	}
	return &RepostNotePage{&o}, nil
}
