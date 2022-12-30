package dmens

type NoteAction = int

const (
	ACTION_POST       NoteAction = 0
	ACTION_REPOST     NoteAction = 1
	ACTION_QUOTE_POST NoteAction = 2
	ACTION_REPLY      NoteAction = 3
	ACTION_LIKE       NoteAction = 4
)
