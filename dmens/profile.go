package dmens

import (
	"bytes"
	"encoding/json"
	"github.com/coming-chat/wallet-SDK/core/sui"
	"io"
	"net/http"
)

const (
	profileModule    = "profile"
	FunctionRegister = "register"
)

type Profile struct {
	Name   string `json:"name"`
	Bio    string `json:"bio"`
	Avatar string `json:"avatar"`
}

type ValidProfile struct {
	Profile   string
	Signature string
}

func (p *Poster) Register(validProfile ValidProfile) (*sui.Transaction, error) {
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
		IsValid   bool   `json:"isValid"`
		Signature string `json:"signature"`
	}

	err = json.Unmarshal(respBody, &respData)
	if err != nil {
		return nil, err
	}
	if respData.IsValid {
		return &ValidProfile{
			Profile:   string(profileBytes),
			Signature: respData.Signature,
		}, nil
	}
	return nil, ErrNotValidPorfile
}
