package dmens

const (
	appIdForComingchatApp = 0
)

type Configuration struct {
	Name                 string
	FullNodeUrl          string
	GraphqlUrl           string
	ContractAddress      string
	GlobalProfileId      string
	GlobalProfileTableId string
	IsMainNet            bool
}

type PosterConfig struct {
	Address    string
	DmensNftId string

	dmensTableId string
	followsId    string
}

func NewPosterConfig(address string) *PosterConfig {
	return &PosterConfig{Address: address}
}

var (
	MainnetConfig = &Configuration{
		Name:       "mainnet",
		IsMainNet:  true,
		GraphqlUrl: "https://graphql.coming.chat/sui-mainnet/graphql",
	}

	TestnetConfig = &Configuration{
		Name:       "testnet",
		GraphqlUrl: "https://graphql.coming.chat/sui-testnet/graphql",
	}

	DevnetConfig = &Configuration{
		Name:       "devnet",
		GraphqlUrl: "https://graphql.coming.chat/sui-devnet/graphql",
	}
)
