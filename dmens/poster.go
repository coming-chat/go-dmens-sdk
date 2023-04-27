package dmens

import (
	"errors"

	"github.com/coming-chat/wallet-SDK/core/sui"
)

type Poster struct {
	*Configuration
	*PosterConfig
	chain *sui.Chain
}

func NewPoster(posterConfig *PosterConfig, configuration *Configuration) (*Poster, error) {
	if posterConfig == nil || configuration == nil {
		return nil, errors.New("invalid poster params")
	}
	poster := &Poster{
		Configuration: configuration,
		PosterConfig:  posterConfig,
	}
	err := poster.FetchDmensGlobalConfig()
	if err != nil {
		return nil, err
	}
	poster.chain = sui.NewChainWithRpcUrl(poster.FullNodeUrl)
	_ = poster.FetchDmensObjecId()
	return poster, nil
}

func (p *Poster) SwitchRpcUrl(rpc string) {
	p.FullNodeUrl = rpc
	p.chain = sui.NewChainWithRpcUrl(rpc)
}

func (p *Poster) IsRegister() bool {
	if p.DmensNftId != "" {
		return true
	}
	if _ = p.FetchDmensObjecId(); p.DmensNftId == "" {
		return false
	}
	return true
}
