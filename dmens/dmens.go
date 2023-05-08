package dmens

import (
	"github.com/coming-chat/wallet-SDK/core/base"
	"github.com/coming-chat/wallet-SDK/core/sui"
)

const (
	actionPost = iota
	ActionRePost
	ActionQuotePost
	ActionReply
	ActionLike

	dmensModule = "dmens"

	functionPost     = "post"
	functionPostRef  = "post_with_ref"
	functionLike     = "like"
	functionFollow   = "follow"
	functionUnfollow = "unfollow"
)

func (p *Poster) DmensPost(text string) (*sui.Transaction, error) {
	return p.chain.BaseMoveCall(
		p.Address,
		p.ContractAddress,
		dmensModule,
		functionPost,
		[]string{},
		[]any{
			p.DmensNftId,
			appIdForComingChatApp,
			text,
		},
		0,
	)
}

func (p *Poster) DmensPostWithRef(action int, text, refIdentifier, refIdPoster string) (*sui.Transaction, error) {
	if action == ActionLike {
		return p.chain.BaseMoveCall(
			p.Address,
			p.ContractAddress,
			dmensModule,
			functionLike,
			[]string{},
			[]any{
				p.DmensNftId,
				appIdForComingChatApp,
				text,
				refIdentifier,
				refIdPoster,
			},
			0,
		)
	}
	return p.chain.BaseMoveCall(
		p.Address,
		p.ContractAddress,
		dmensModule,
		functionPostRef,
		[]string{},
		[]any{
			p.DmensNftId,
			appIdForComingChatApp,
			action,
			text,
			refIdentifier,
		},
		0,
	)
}

func (p *Poster) DmensFollow(addresses *base.StringArray) (*sui.Transaction, error) {
	return p.chain.BaseMoveCall(
		p.Address,
		p.ContractAddress,
		dmensModule,
		functionFollow,
		[]string{},
		[]any{
			p.DmensNftId,
			addresses.Values,
		},
		0,
	)
}

func (p *Poster) DmensUnfollow(addresses *base.StringArray) (*sui.Transaction, error) {
	return p.chain.BaseMoveCall(
		p.Address,
		p.ContractAddress,
		dmensModule,
		functionUnfollow,
		[]string{},
		[]any{
			p.DmensNftId,
			addresses.Values,
		},
		0,
	)
}
