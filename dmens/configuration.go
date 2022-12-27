package dmens

const (
	APP_ID_FOR_COMINGCHAT_APP = 0
	SuiDevNet                 = "https://fullnode.devnet.sui.io"
	Module                    = "dmens"
	GasBudGet                 = 1000
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
		ContractAddress: "0xc5ee772bc6cabe728810e3282aef5333aaaa8cfd",
	}
)
