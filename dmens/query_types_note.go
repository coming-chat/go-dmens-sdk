package dmens

import (
	"github.com/coming-chat/wallet-SDK/core/base"
	"github.com/coming-chat/wallet-SDK/core/base/inter"
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

func NewNote() *Note {
	return &Note{}
}

func (n *Note) JsonString() (*base.OptionalString, error) {
	return base.JsonString(n)
}
func NewNoteWithJsonString(str string) (*Note, error) {
	var o Note
	err := base.FromJsonString(str, &o)
	return &o, err
}

type RepostNote struct {
	*Note        // origin note info
	Repost *Note `json:"repost"` // repost note info
}

func NewRepostNote() *RepostNote {
	return &RepostNote{}
}

func (n *RepostNote) JsonString() (*base.OptionalString, error) {
	return base.JsonString(n)
}
func NewRepostNoteWithJsonString(str string) (*RepostNote, error) {
	var o RepostNote
	err := base.FromJsonString(str, &o)
	return &o, err
}

type NotePage struct {
	*inter.SdkPageable[*Note]
}

func NewNotePage() *NotePage {
	return &NotePage{SdkPageable: &inter.SdkPageable[*Note]{}}
}

func NewNotePageWithJsonString(str string) (*NotePage, error) {
	var o inter.SdkPageable[*Note]
	err := base.FromJsonString(str, &o)
	if err != nil {
		return nil, err
	}
	return &NotePage{&o}, nil
}

type RepostNotePage struct {
	*inter.SdkPageable[*RepostNote]
}

func NewRepostNotePage() *RepostNotePage {
	return &RepostNotePage{}
}

func NewRepostNotePageWithJsonString(str string) (*RepostNotePage, error) {
	var o inter.SdkPageable[*RepostNote]
	err := base.FromJsonString(str, &o)
	if err != nil {
		return nil, err
	}
	return &RepostNotePage{&o}, nil
}
