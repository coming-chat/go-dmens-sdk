package dmens

func (p *Poster) QueryDmensObjectId() *Query {
	return &Query{
		Query: `
		query dmensObjectId($owner: JSON, $type: String) {
			allSuiObjects(
			  filter: { owner: { equalTo: $owner }, type: { equalTo: $type } }
			) {
			  nodes {
				objectId
			  }
			}
		  }
		`,
		Variables: map[string]interface{}{
			"owner": map[string]string{
				"AddressOwner": p.Address,
			},
			"type": p.ContractAddress + "::dmens::DmensMeta",
		},
	}
}

// FetchDmensObjecId After ios or android call profileRegister and send that transaction,
// this func should be recalled again to fetch the registered dmens object id
func (p *Poster) FetchDmensObjecId() error {
	var res = struct {
		AllSuiObjects struct {
			Nodes []struct {
				ObjectId string `json:"objectId"`
			} `json:"nodes"`
		} `json:"allSuiObjects"`
	}{}
	query := p.QueryDmensObjectId()
	err := p.makeQueryOut(query, &res)
	if err != nil {
		return err
	}
	if len(res.AllSuiObjects.Nodes) == 0 {
		p.DmensNftId = ""
		return ErrUserNotRegistered
	}
	p.DmensNftId = res.AllSuiObjects.Nodes[0].ObjectId
	return nil
}
