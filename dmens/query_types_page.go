package dmens

import "github.com/coming-chat/wallet-SDK/core/base"

type sdkPageable[T Note | UserInfo | RepostNote] struct {
	TotalCount    int    `json:"totalCount"`
	CurrentCount  int    `json:"currentCount"`
	CurrentCursor string `json:"currentCursor"`
	HasNextPage   bool   `json:"hasNextPage"`

	Items    []*T `json:"items"`
	anyArray *base.AnyArray
}

func (p *sdkPageable[T]) JsonString() (string, error) {
	return JsonString(p)
}

func (p *sdkPageable[T]) ItemArray() *base.AnyArray {
	if p.anyArray == nil {
		a := make([]any, len(p.Items))
		for idx, n := range p.Items {
			a[idx] = n
		}
		p.anyArray = &base.AnyArray{Values: a}
	}
	return p.anyArray
}

func (p *sdkPageable[T]) FirstObject() *T {
	if len(p.Items) <= 0 {
		return nil
	}
	return p.Items[0]
}

// MapToNotePage
// @param poster If you need to query the status of notes in batches, you need to provide the poster.
func (a *rawNotePage) MapToNotePage(poster *Poster, pageSize int) *NotePage {
	sdkPage := a.mapToSdkPage(pageSize, func(rn rawNote) *Note {
		return rn.MapToNote()
	})
	page := &NotePage{sdkPageable: sdkPage}
	if poster != nil {
		_ = poster.BatchQueryNoteStatus(page, "")
	}
	return page
}

// MapToUserPage
func (a *rawUserPage) MapToUserPage(pageSize int) *UserPage {
	sdkPage := a.mapToSdkPage(pageSize, func(ui UserInfo) *UserInfo {
		return &ui
	})
	return &UserPage{sdkPageable: sdkPage}
}

// MapToUserPage
func (a *rawUserFollowPage) MapToUserPage(pageSize int) *UserPage {
	sdkPage := a.mapToSdkPage(pageSize, func(ruf rawUserFollow) *UserInfo {
		return ruf.MapToUserInfo()
	})
	return &UserPage{sdkPageable: sdkPage}
}

// 合并完成后，originPage 的 notes 会被清空
func combineRepostPage(repostPage, originPage *NotePage) *RepostNotePage {
	notes := make([]*RepostNote, len(repostPage.Items))
	for idx, note := range repostPage.Items {
		notes[idx] = &RepostNote{
			Repost: note,
		}
		// 在 originPage 中找到对应的 note
		for oidx, originNote := range originPage.Items {
			if originNote.NoteId == note.RefId {
				notes[idx].Note = originNote
				originPage.Items = append(originPage.Items[:oidx], originPage.Items[oidx+1:]...)
				continue
			}
		}
	}
	return &RepostNotePage{
		sdkPageable: &sdkPageable[RepostNote]{
			TotalCount:    repostPage.TotalCount,
			CurrentCount:  repostPage.CurrentCount,
			CurrentCursor: repostPage.CurrentCursor,
			HasNextPage:   repostPage.HasNextPage,

			Items: notes,
		},
	}
}
