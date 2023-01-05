package dmens

func (p *Poster) QueryDmensGlobalConfig() *Query {
	return &Query{
		Query: `
		query DmensGlobalConfig {
			allDmensGlobalConfigs{
			  	nodes {
					chainName
      				contractAddress
      				fullNodeUrl
      				globalProfileId
      				globalProfileTableId
			  	}
			}
		}
		`,
	}
}

func (p *Poster) FetchDmensGlobalConfig() error {
	var out = []struct {
		ChainName            string `json:"chainName"`
		ContractAddress      string `json:"contractAddress"`
		FullNodeUrl          string `json:"fullNodeUrl"`
		GlobalProfileId      string `json:"globalProfileId"`
		GlobalProfileTableId string `json:"globalProfileTableId"`
	}{}
	query := p.QueryDmensGlobalConfig()
	err := p.makeQueryOut(query, "allDmensGlobalConfigs.nodes", &out)
	if err != nil {
		return err
	}
	if len(out) == 0 {
		return ErrGetNullConfiguration
	}
	node := out[0]
	p.Name = node.ChainName
	p.FullNodeUrl = node.FullNodeUrl
	p.ContractAddress = node.ContractAddress
	p.GlobalProfileId = node.GlobalProfileId
	p.GlobalProfileTableId = node.GlobalProfileTableId
	return nil
}