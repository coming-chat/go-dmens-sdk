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
				fields
			  }
			}
		  }
		`,
		Variables: map[string]interface{}{
			"owner": p.addressOwner(),
			"type":  p.dmensMetaObjectType(),
		},
	}
}

// FetchDmensObjecId After ios or android call profileRegister and send that transaction,
// this func should be recalled again to fetch the registered dmens object id
func (p *Poster) FetchDmensObjecId() error {
	type Field struct {
		Fields struct {
			Id struct {
				Id string `json:"id"`
			} `json:"id"`
		} `json:"fields"`
	}
	var res = struct {
		AllSuiObjects struct {
			Nodes []struct {
				ObjectId string `json:"objectId"`
				Fields   struct {
					Follows    Field `json:"follows"`
					DmensTable Field `json:"dmens_table"`
				} `json:"fields"`
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
		p.dmensTableId = ""
		p.followsId = ""
		return ErrUserNotRegistered
	}
	node := res.AllSuiObjects.Nodes[0]
	p.DmensNftId = node.ObjectId
	p.dmensTableId = node.Fields.DmensTable.Fields.Id.Id
	p.followsId = node.Fields.Follows.Fields.Id.Id
	return nil
}
