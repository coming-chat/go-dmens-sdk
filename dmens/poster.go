package dmens

import (
	"github.com/coming-chat/wallet-SDK/core/sui"
)

type Poster struct {
	*Configuration
	*PosterConfig
	chain *sui.Chain
}

func NewPoster(posterConfig *PosterConfig, configuration *Configuration) (*Poster, error) {
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

func NewPosterWithAddress(posterAddress string, configuration *Configuration) (*Poster, error) {
	posterConfig := &PosterConfig{
		Address: posterAddress,
	}
	return NewPoster(posterConfig, configuration)
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
