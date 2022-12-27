package dmens

import (
	"context"
	"encoding/json"
	"github.com/coming-chat/go-sui/types"
)

const (
	FunctionRegister = "register"
)

type Profile struct {
	Name   string `json:"name"`
	Bio    string `json:"bio"`
	Avatar string `json:"avatar"`
}

func (p *Poster) Register(profile Profile) (*types.TransactionBytes, error) {
	profileBytes, err := json.Marshal(profile)
	if err != nil {
		return nil, err
	}
	coins, err := p.chain.GetSuiCoinsOwnedByAddress(context.Background(), *p.address)
	if err != nil {
		return nil, err
	}
	coin, err := coins.PickCoinNoLess(1000)
	if err != nil {
		return nil, err
	}
	tx, err := p.chain.MoveCall(
		context.Background(),
		*p.address,
		*p.packageId,
		Module,
		FunctionRegister,
		[]string{},
		[]any{
			p.GlobalProfileId,
			string(profileBytes),
			"", //this field is disabled
		},
		&coin.Reference.ObjectId,
		GasBudGet,
	)
	if err != nil {
		return nil, err
	}
	return tx, nil
}
