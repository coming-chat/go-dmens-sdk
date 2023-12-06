package dmens

import (
	"github.com/coming-chat/wallet-SDK/core/base/inter"
)

// MapToNotePage
// @param poster If you need to query the status of notes in batches, you need to provide the poster.
func (a *rawNotePage) MapToNotePage(poster *Poster, pageSize int) *NotePage {
	sdkPage := a.mapToSdkPage(pageSize, func(rn rawNote) *Note {
		return rn.MapToNote()
	})
	page := &NotePage{SdkPageable: sdkPage}
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
	return &UserPage{SdkPageable: sdkPage}
}

// MapToUserPage
func (a *rawUserFollowPage) MapToUserPage(pageSize int) *UserPage {
	sdkPage := a.mapToSdkPage(pageSize, func(ruf rawTrendUser) *UserInfo {
		return ruf.UserInfo
	})
	return &UserPage{SdkPageable: sdkPage}
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
				break
			}
		}
	}
	return &RepostNotePage{
		SdkPageable: &inter.SdkPageable[*RepostNote]{
			TotalCount_:    repostPage.TotalCount_,
			CurrentCount_:  repostPage.CurrentCount_,
			CurrentCursor_: repostPage.CurrentCursor_,
			HasNextPage_:   repostPage.HasNextPage_,

			Items: notes,
		},
	}
}
