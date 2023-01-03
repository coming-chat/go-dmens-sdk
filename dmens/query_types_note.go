package dmens

type Note struct {
	CreateTime string `json:"createTime"`
	NoteId     string `json:"noteId"`

	Action NoteAction `json:"action"`
	Text   string     `json:"text"`
	Poster string     `json:"poster"`
	RefId  string     `json:"refId"`
}

type NotePage struct {
	Notes         []Note `json:"notes"`
	CurrentCursor string `json:"currentCursor"`
	CurrentCount  int    `json:"currentCount"`
	TotalCount    int    `json:"totalCount"`
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
			// Type   string `json:"type"`
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
	fields := a.Fields.Value.Fields
	return &Note{
		CreateTime: a.CreateTime,
		NoteId:     a.ObjectId,
		Action:     fields.Action,
		Text:       fields.Text,
		Poster:     fields.Poster,
		RefId:      fields.RefId,
	}
}

func (a *rawNotePage) MapToNotePage() *NotePage {
	length := len(a.Edges)
	if length == 0 {
		return &NotePage{
			TotalCount: a.TotalCount,
		}
	}
	notes := make([]Note, 0)
	for _, n := range a.Edges {
		notes = append(notes, *n.Node.MapToNote())
	}
	return &NotePage{
		TotalCount:    a.TotalCount,
		Notes:         notes,
		CurrentCount:  len(notes),
		CurrentCursor: a.Edges[length-1].Cursor,
	}
}

func (a *rawNotePage) FirstObject() *rawNote {
	if len(a.Edges) <= 0 {
		return nil
	}
	return &a.Edges[0].Node
}

func (n *Note) JsonString() (string, error) {
	return JsonString(n)
}

func (n *NotePage) JsonString() (string, error) {
	return JsonString(n)
}
