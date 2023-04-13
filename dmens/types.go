package dmens

import (
	"encoding/json"
	"errors"
	"strconv"
)

type NoteAction = int

const (
	ACTION_POST       NoteAction = 0
	ACTION_REPOST     NoteAction = 1
	ACTION_QUOTE_POST NoteAction = 2
	ACTION_REPLY      NoteAction = 3
	ACTION_LIKE       NoteAction = 4
)

type rawNoteAction struct {
	Value NoteAction
}

func (a *rawNoteAction) UnmarshalJSON(d []byte) error {
	var i NoteAction
	err := json.Unmarshal(d, &i)
	if err == nil {
		a.Value = i
		return nil
	}
	var str string
	err = json.Unmarshal(d, &str)
	if err == nil {
		i64, err := strconv.ParseInt(str, 10, 64)
		if err == nil {
			a.Value = int(i64)
			return nil
		}
	}
	return errors.New("invalid data")
}

func (a rawNoteAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.Value)
}
