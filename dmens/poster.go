package dmens

import (
	"context"
	"errors"
	"github.com/coming-chat/go-sui/client"
	"github.com/coming-chat/go-sui/types"
	"net/http"
	"time"
)

type Poster struct {
	Configuration
	chain     *client.Client
	address   *types.Address
	packageId *types.HexData
	PosterConfig
}

func NewPoster(posterConfig PosterConfig, configuration Configuration) (*Poster, error) {
	hc := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:    20,
			IdleConnTimeout: 20 * time.Second,
		},
		Timeout: 15 * time.Second,
	}
	ct, _ := client.DialWithClient(configuration.FullNodeUrl, hc)
	addr, err := types.NewAddressFromHex(posterConfig.Address)
	if err != nil {
		return nil, err
	}
	packId, err := types.NewHexData(configuration.ContractAddress)
	if err != nil {
		return nil, err
	}
	poster := &Poster{
		Configuration: configuration,
		chain:         ct,
		address:       addr,
		packageId:     packId,
		PosterConfig:  posterConfig,
	}
	err = poster.initialDmensObjecId()
	if err != nil {
		return nil, err
	}
	return poster, nil
}

func (p *Poster) IsRegister() bool {
	if p.DmensNftId == "" {
		return false
	}
	return true
}

func (p *Poster) ExecTransaction(signedTxn *types.SignedTransaction) (string, error) {
	response, err := p.chain.ExecuteTransaction(context.Background(), *signedTxn, types.TxnRequestTypeWaitForLocalExecution)
	if err != nil {
		return "", err
	}
	if response.EffectsCert.ConfirmedLocalExecution {
		return response.TransactionDigest(), nil
	}
	return "", errors.New("exec transaction failed")
}
