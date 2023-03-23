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
			actionPost,
			text,
		},
	)
}

func (p *Poster) DmensPostWithRef(action int, text, refIdentifier string) (*sui.Transaction, error) {
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
	)
}
