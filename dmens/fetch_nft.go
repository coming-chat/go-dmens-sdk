package dmens

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/coming-chat/go-sui/types"
	"github.com/coming-chat/wallet-SDK/core/base"
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

func (n *NFTAvatar) AsAny() *base.Any {
	return &base.Any{Value: n}
}

func AsNFTAvatar(any *base.Any) *NFTAvatar {
	if res, ok := any.Value.(*NFTAvatar); ok {
		return res
	}
	if res, ok := any.Value.(NFTAvatar); ok {
		return &res
	}
	return nil
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

	ids := make([]types.ObjectId, 0)
	for _, nftId := range nftIds {
		id, err := types.NewHexData(nftId)
		if err != nil {
			continue
		}
		ids = append(ids, *id)
	}
	client, err := p.chain.Client()
	if err != nil {
		return nil, err
	}
	elems, err := client.MultiGetObjects(context.Background(), ids, &types.SuiObjectDataOptions{
		ShowType:    true,
		ShowContent: true,
	})
	if err != nil {
		return nil, err
	}

	results := make(map[string]*NFTAvatar)
	for _, ele := range elems {
		if ele.Error != nil {
			return nil, fmt.Errorf("some nft not found")
		}
		avatar := mapToNFTAvatar(ele.Data)
		if avatar != nil {
			results[avatar.Id] = avatar
		}
	}
	return results, nil
}

func mapToNFTAvatar(obj *types.SuiObjectData) *NFTAvatar {
	if obj == nil || obj.Content == nil {
		return nil
	}

	content := struct {
		Type   string `json:"type"`
		Fields struct {
			Url  string `json:"url"`
			Name string `json:"name"`
		} `json:"fields"`
	}{}
	contentData, err := json.Marshal(obj.Content)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(contentData, &content)
	if err != nil {
		return nil
	}

	return &NFTAvatar{
		Id:    obj.ObjectId.String(),
		Image: strings.Replace(content.Fields.Url, "ipfs://", "https://ipfs.io/ipfs/", 1),
		Type:  content.Type,
		Name:  content.Fields.Name,
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
	options := types.SuiObjectDataOptions{ShowType: true, ShowContent: true}
	objs, err := client.BatchGetFilteredObjectsOwnedByAddress(context.Background(), *addr, options, func(oi *types.SuiObjectData) bool {
		if strings.HasSuffix(*oi.Type, "::registrar::RegistrationNFT") {
			return true
		}
		return false
	})
	if err != nil {
		return nil, err
	}
	for _, obj := range objs {
		nft := mapToNFTAvatar(obj.Data)
		if nft != nil && nft.Name != "" {
			return &base.OptionalString{Value: nft.Name}, nil
		}
	}
	return nil, fmt.Errorf("sui name by address `%v` not found", address)
}
