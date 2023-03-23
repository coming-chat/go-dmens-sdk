package dmens

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/coming-chat/go-sui/client"
	"github.com/coming-chat/go-sui/types"
	"github.com/coming-chat/wallet-SDK/core/base"
)

type NFTAvatar struct {
	Id    string `json:"id"`
	Image string `json:"image"`
	Type  string `json:"type"`
}

func NewNFTAvatar(nftId, image, typ string) *NFTAvatar {
	return &NFTAvatar{
		Id:    nftId,
		Image: image,
		Type:  typ,
	}
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

func (p *Poster) BatchQueryNFTAvatarByIds(nftIds []string) (map[string]*NFTAvatar, error) {
	if len(nftIds) <= 0 {
		return nil, nil
	}
	var elems []client.BatchElem
	for _, nftid := range nftIds {
		elems = append(elems, client.BatchElem{
			Method: "sui_getObject",
			Args:   []interface{}{nftid},
			Result: &types.ObjectRead{},
		})
	}
	client, err := p.chain.Client()
	if err != nil {
		return nil, err
	}
	err = client.BatchCallContext(context.Background(), elems)
	if err != nil {
		return nil, err
	}
	results := make(map[string]*NFTAvatar)
	for _, ele := range elems {
		if ele.Error != nil {
			return nil, ele.Error
		}
		obj := ele.Result.(*types.ObjectRead)
		avatar := mapToNFTAvatar(obj)
		if avatar != nil {
			results[avatar.Id] = avatar
		}
	}
	return results, nil
}

func mapToNFTAvatar(obj *types.ObjectRead) *NFTAvatar {
	if obj == nil || obj.Status != types.ObjectStatusExists {
		return nil
	}

	meta := struct {
		Type   string `json:"type"`
		Fields struct {
			Id struct {
				Id string `json:"id"`
			} `json:"id"`
			Url string `json:"url"`
		} `json:"fields"`
	}{}
	metaBytes, err := json.Marshal(obj.Details.Data)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(metaBytes, &meta)
	if err != nil {
		return nil
	}

	return &NFTAvatar{
		Id:    meta.Fields.Id.Id,
		Image: strings.Replace(meta.Fields.Url, "ipfs://", "https://ipfs.io/ipfs/", 1),
		Type:  meta.Type,
	}
}

func (p *Poster) QuerySuiNameByAddress(address string) (*base.OptionalString, error) {
	client, err := p.chain.Client()
	if err != nil {
		return nil, err
	}
	addr, err := types.NewAddressFromHex(address)
	if err != nil {
		return nil, err
	}
	objs, err := client.BatchGetFilteredObjectsOwnedByAddress(context.Background(), *addr, func(oi types.ObjectInfo) bool {
		if strings.HasSuffix(oi.Type, "::base_registrar::RegistrationNFT") {
			return true
		}
		return false
	})
	if err != nil {
		return nil, err
	}
	for _, obj := range objs {
		var out struct {
			Fields struct {
				Name string `json:"name"`
			} `json:"fields"`
		}
		data, err := json.Marshal(obj.Details.Data)
		if err != nil {
			continue
		}
		err = json.Unmarshal(data, &out)
		if err != nil {
			continue
		}
		if out.Fields.Name != "" {
			return &base.OptionalString{Value: out.Fields.Name}, nil
		}
	}
	return nil, fmt.Errorf("sui name by address `%v` not found", address)
}
