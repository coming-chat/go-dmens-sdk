package dmens

type NoteStatus struct {
	NoteId string `json:"noteId"`
	Viewer string `json:"viewer"`

	ReplyCount  int `json:"replyCount"`
	RepostCount int `json:"repostCount"`
	LikeCount   int `json:"likeCount"`

	// Whether the viewer reposted it
	IsReposted bool `json:"isReposted"`
	// Whether the viewer liked it
	IsLiked bool `json:"isLiked"`
}

// QueryNoteStatusById
// @param noteId the note's id
// @param viewer the note's viewer, if the viewer is empty, the poster's address will be queried.
func (p *Poster) QueryNoteStatusById(noteId string, viewer string) (*NoteStatus, error) {
	statuses, err := p.BatchQueryNoteStatusByIds([]string{noteId}, viewer)
	if err != nil {
		return nil, err
	}
	return statuses[noteId], nil
}

// BatchQueryNoteStatus
// 批量查询 page 中所有 note 的状态，数据会直接同步到 page 中每一个 note 对象中
// @param viewer the note's viewer, if the viewer is empty, the poster's address will be queried.
func (p *Poster) BatchQueryNoteStatus(page *NotePage, viewer string) error {
	if len(page.Items) == 0 {
		return nil
	}
	if viewer == "" {
		viewer = p.Address
	}

	noteids := make([]string, len(page.Items))
	for idx, note := range page.Items {
		noteids[idx] = note.NoteId
	}
	statuses, err := p.BatchQueryNoteStatusByIds(noteids, viewer)
	if err != nil {
		return err
	}

	for _, note := range page.Items {
		if status, ok := statuses[note.NoteId]; ok {
			note.Status = status
		}
	}

	return nil
}

// BatchQueryNoteStatusByIds
func (p *Poster) BatchQueryNoteStatusByIds(noteIds []string, viewer string) (map[string]*NoteStatus, error) {
	if len(noteIds) == 0 {
		return make(map[string]*NoteStatus, 0), nil
	}
	if viewer == "" {
		viewer = p.Address
	}

	query := &Query{
		Query: `
		query MyQuery($dmensObjectType: String, $watcher: String, $dmensId: [String]) {
			fetchDmensStatus(
			  dmensId: $dmensId
			  dmensType: $dmensObjectType
			  watcher: $watcher
			) {
			  edges {
				node {
				  id
				  actionCount
				  action
				  poster
				}
			  }
			}
		  }
		`,
		Variables: map[string]interface{}{
			"dmensId":         noteIds,
			"dmensObjectType": p.dmensObjectType(),
			"watcher":         viewer,
		},
	}

	type rawCounter struct {
		Id          string     `json:"id"`
		ActionCount int        `json:"actionCount,string"`
		Action      NoteAction `json:"action"`
		Poster      string     `json:"poster,omitempty"`
	}
	var out []struct {
		Node rawCounter `json:"node"`
	}
	err := p.makeQueryOut(query, "fetchDmensStatus.edges", &out)
	if err != nil {
		return nil, err
	}

	result := make(map[string]*NoteStatus)
	for _, node := range out {
		counter := node.Node
		status, exists := result[counter.Id]
		if !exists {
			status = &NoteStatus{
				NoteId: counter.Id,
				Viewer: viewer,
			}
			result[counter.Id] = status
		}

		isViewerActioned := (counter.Poster == viewer)
		count := counter.ActionCount
		switch counter.Action {
		case ACTION_REPLY:
			status.ReplyCount = count
		case ACTION_REPOST:
			status.RepostCount = count
			status.IsReposted = isViewerActioned
		case ACTION_LIKE:
			status.LikeCount = count
			status.IsLiked = isViewerActioned
		}
	}

	return result, nil
}
