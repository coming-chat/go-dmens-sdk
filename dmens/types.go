package dmens

type NoteAction int

const (
	ACTION_POST        NoteAction = 0
	ACTION_RETWEET     NoteAction = 1
	ACTION_QUOTE_TWEET NoteAction = 2
	ACTION_REPLY       NoteAction = 3
	ACTION_ATTACH      NoteAction = 4
	ACTION_LIKE        NoteAction = 5
)
