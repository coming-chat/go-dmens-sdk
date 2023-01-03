package dmens

import (
	"encoding/json"
	"github.com/coming-chat/wallet-SDK/core/sui"
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

func (p *Poster) Register(profile *Profile) (*sui.Transaction, error) {
	profileBytes, err := json.Marshal(profile)
	if err != nil {
		return nil, err
	}
	return p.chain.BaseMoveCall(
		p.Address,
		p.ContractAddress,
		profileModule,
		FunctionRegister,
		[]any{
			p.GlobalProfileId,
			string(profileBytes),
			"", //signature is disabled
		})
}
