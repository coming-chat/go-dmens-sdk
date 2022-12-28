package dmens

import (
	"github.com/coming-chat/wallet-SDK/core/sui"
)

type Poster struct {
	Configuration
	PosterConfig
	chain *sui.Chain
}

func NewPoster(posterConfig PosterConfig, configuration Configuration) (*Poster, error) {
	poster := &Poster{
		Configuration: configuration,
		chain:         sui.NewChainWithRpcUrl(configuration.FullNodeUrl),
		PosterConfig:  posterConfig,
	}
	err := poster.initialDmensObjecId()
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
