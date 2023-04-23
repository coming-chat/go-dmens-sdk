package dmens

const (
	appIdForComingChatApp = 0
)

type Configuration struct {
	Name                 string
	FullNodeUrl          string
	GraphqlUrl           string
	ContractAddress      string
	GlobalProfileId      string
	GlobalProfileTableId string
	ProfileCheckUrl      string
	IsMainNet            bool
}

type PosterConfig struct {
	Address    string
	DmensNftId string
	// Default false
	Reviewing bool

	dmensTableId string
	followsId    string
}

func NewPosterConfig(address string, reviewing bool) *PosterConfig {
	return &PosterConfig{
		Address:   address,
		Reviewing: reviewing,
	}
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
