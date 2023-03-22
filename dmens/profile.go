package dmens

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/coming-chat/wallet-SDK/core/sui"
)

const (
	profileModule      = "profile"
	FunctionRegister   = "register"
	FunctionAddItem    = "add_item"
	FunctionRemoveItem = "remove_item"
)

type Profile struct {
	Name           string `json:"name"`
	Bio            string `json:"bio"`
	Avatar         string `json:"avatar"`
	Background     string `json:"background"`
	Website        string `json:"website"`
	Identification string `json:"identification"`
}

type ValidProfile struct {
	Profile   string
	Signature string
}

func (p *Poster) Register(validProfile *ValidProfile) (*sui.Transaction, error) {
	return p.chain.BaseMoveCall(
		p.Address,
		p.ContractAddress,
		profileModule,
		FunctionRegister,
		[]any{
			p.GlobalProfileId,
			validProfile.Profile,
			validProfile.Signature,
		})
}

func (p *Poster) CheckProfile(profile *Profile) (*ValidProfile, error) {
	profileBytes, err := json.Marshal(profile)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPut, p.ProfileCheckUrl+p.Address, bytes.NewBuffer(profileBytes))
	if err != nil {
		return nil, err
	}
	req.Header["Content-Type"] = []string{"application/json"}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var respData struct {
		IsValid   bool    `json:"isValid"`
		Signature string  `json:"signature"`
		Profile   Profile `json:"profile"`
	}

	err = json.Unmarshal(respBody, &respData)
	if err != nil {
		return nil, err
	}
	newProfileData, err := json.Marshal(respData.Profile)
	if err != nil {
		return nil, err
	}
	if respData.IsValid {
		return &ValidProfile{
			Profile:   string(newProfileData),
			Signature: respData.Signature,
		}, nil
	}
	return nil, ErrNotValidPorfile
}

func (p *Poster) SetNftAvatar(nftId string) (*sui.Transaction, error) {
	nft, err := p.QueryNFTAvatar(nftId)
	if err != nil {
		return nil, err
	}
	return p.chain.BaseMoveCall(
		p.Address,
		p.ContractAddress,
		profileModule,
		FunctionAddItem,
		[]any{
			nft.Type,
			p.GlobalProfileId,
			nft.Id,
		})
}

func (p *Poster) RemoveNftAvatar(avatar *NFTAvatar) (*sui.Transaction, error) {
	return p.chain.BaseMoveCall(
		p.Address,
		p.ContractAddress,
		profileModule,
		FunctionRemoveItem,
		[]any{
			avatar.Type,
			p.GlobalProfileId,
			avatar.Id,
		})
}
