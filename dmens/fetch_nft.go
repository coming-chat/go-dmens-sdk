package dmens

import (
	"context"
	"fmt"
	"strings"

	"github.com/coming-chat/go-sui/v2/sui_types"
	"github.com/coming-chat/go-sui/v2/types"
	"github.com/coming-chat/wallet-SDK/core/base"
	"github.com/coming-chat/wallet-SDK/core/sui"
)

type NFTAvatar struct {
	Id    string `json:"id"`
	Image string `json:"image"`
	Type  string `json:"type"`
	Name  string `json:"name"`
}

func NewNFTAvatar() *NFTAvatar {
	return &NFTAvatar{}
}

func NewNFTAvatarWithId(nftId, image, typ string) *NFTAvatar {
	return &NFTAvatar{
		Id:    nftId,
		Image: image,
		Type:  typ,
	}
}

func (n *NFTAvatar) JsonString() (*base.OptionalString, error) {
	return base.JsonString(n)
}
func NewNFTAvatarWithJsonString(str string) (*NFTAvatar, error) {
	var o NFTAvatar
	err := base.FromJsonString(str, &o)
	return &o, err
}

func (p *Poster) BatchQueryNFTAvatarForUserPage(page *UserPage) error {
	ids := make([]string, 0)
	for _, user := range page.Items {
		l := len(user.Item)
		if l > 0 && len(user.Item[l-1]) > 0 {
			ids = append(ids, user.Item[l-1])
		}
	}
	if len(ids) <= 0 {
		return nil
	}
	results, err := p.BatchQueryNFTAvatarByIds(ids)
	if err != nil {
		return err
	}
	if len(results) == 0 {
		return nil
	}
	for _, user := range page.Items {
		l := len(user.Item)
		if l > 0 && len(user.Item[l-1]) > 0 {
			key := user.Item[l-1]
			if avatar, ok := results[key]; ok {
				user.NFTAvatar = avatar
			}
		}
	}
	return nil
}

func (p *Poster) QueryNFTAvatar(nftId string) (*NFTAvatar, error) {
	nfts, err := p.BatchQueryNFTAvatarByIds([]string{nftId})
	if err != nil {
		return nil, err
	}
	if nft, ok := nfts[nftId]; ok {
		return nft, nil
	}
	return nil, fmt.Errorf("nft %v not found", nftId)
}

func (p *Poster) BatchQueryNFTAvatarByIds(nftIds []string) (res map[string]*NFTAvatar, err error) {
	defer base.CatchPanicAndMapToBasicError(&err)

	if len(nftIds) <= 0 {
		return
	}

	ids := make([]sui_types.ObjectID, 0)
	for _, nftId := range nftIds {
		id, err := sui_types.NewObjectIdFromHex(nftId)
		if err != nil {
			continue
		}
		ids = append(ids, *id)
	}
	client, err := p.chain.Client()
	if err != nil {
		return
	}
	elems, err := client.MultiGetObjects(context.Background(), ids, &types.SuiObjectDataOptions{
		ShowType:    true,
		ShowDisplay: true,
	})
	if err != nil {
		return
	}

	results := make(map[string]*NFTAvatar)
	for _, ele := range elems {
		avatar := mapToNFTAvatar(ele)
		if avatar != nil {
			results[avatar.Id] = avatar
		}
	}
	return results, nil
}

func mapToNFTAvatar(obj types.SuiObjectResponse) *NFTAvatar {
	nft := sui.TransformNFT(&obj)
	if nft == nil {
		return nil
	}
	typ := ""
	if obj.Data.Type != nil {
		typ = *obj.Data.Type
	}
	return &NFTAvatar{
		Id:    nft.Id,
		Image: nft.Image,
		Type:  typ,
		Name:  nft.Name,
	}
}

func (p *Poster) QuerySuiNameByAddress(address string) (name *base.OptionalString, err error) {
	defer base.CatchPanicAndMapToBasicError(&err)

	client, err := p.chain.Client()
	if err != nil {
		return
	}
	addr, err := sui_types.NewAddressFromHex(address)
	if err != nil {
		return
	}
	options := types.SuiObjectDataOptions{
		ShowType:    true,
		ShowDisplay: true}
	objs, err := client.BatchGetFilteredObjectsOwnedByAddress(context.Background(), *addr, options, func(oi *types.SuiObjectData) bool {
		if strings.HasSuffix(*oi.Type, "::registrar::RegistrationNFT") {
			return true
		}
		return false
	})
	if err != nil {
		return
	}
	for _, obj := range objs {
		nft := mapToNFTAvatar(obj)
		if nft != nil {
			return &base.OptionalString{Value: nft.Name}, nil
		}
	}
	return nil, fmt.Errorf("sui name by address `%v` not found", address)
}
