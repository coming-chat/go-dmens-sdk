package dmens

const (
	appIdForComingchatApp = 0
)

type Configuration struct {
	Name            string
	FullNodeUrl     string
	GraphqlUrl      string
	ContractAddress string
	GlobalProfileId string
	// ProfileTableId string
	IsMainNet bool
}

type PosterConfig struct {
	Address    string
	DmensNftId string

	dmensTableId string
	followsId    string
}

var (
	MainnetConfig = &Configuration{
		Name:        "mainnet",
		FullNodeUrl: "https://fullnode.mainnet.sui.io:443",
		IsMainNet:   true,
	}

	TestnetConfig = &Configuration{
		Name:            "testnet",
		FullNodeUrl:     "https://fullnode.testnet.sui.io:443",
		GraphqlUrl:      "https://graphql.coming.chat/sui-testnet/graphql",
		ContractAddress: "0x7a3ff93380660c4fa3ea8df8de13acb2cadf7052",
	}

	DevnetConfig = &Configuration{
		Name:            "devnet",
		FullNodeUrl:     "https://fullnode.devnet.sui.io:443",
		GraphqlUrl:      "https://graphql.coming.chat/sui-devnet/graphql",
		ContractAddress: "0x307772d54928d34b9a45bac2f436db7e3e33fe5e",
		GlobalProfileId: "0x6a0402a6c37fb14446683ff13bc97d1ee0474ac2",
	}
)
