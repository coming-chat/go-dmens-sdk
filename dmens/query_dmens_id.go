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
		Fields rawFieldsId `json:"fields"`
	}
	var out []struct {
		ObjectId string `json:"objectId"`
		Fields   struct {
			Follows    Field `json:"follows"`
			DmensTable Field `json:"dmens_table"`
		} `json:"fields"`
	}
	query := p.QueryDmensObjectId()
	err := p.makeQueryOut(query, "allSuiObjects.nodes", &out)
	if err != nil {
		return err
	}
	if len(out) == 0 {
		p.DmensNftId = ""
		p.dmensTableId = ""
		p.followsId = ""
		return ErrUserNotRegistered
	}
	node := out[0]
	p.DmensNftId = node.ObjectId
	p.dmensTableId = node.Fields.DmensTable.Fields.Id.Id
	p.followsId = node.Fields.Follows.Fields.Id.Id
	return nil
}
