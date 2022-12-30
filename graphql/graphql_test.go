package graphql

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const graphqlURL = "https://graphql.coming.chat/sui-devnet/graphql"

func TestXxx(t *testing.T) {
	query := `
	query trendNoteList($type: String) {
		trendingNotes(
		  filter: { status: { equalTo: "Exists" } }
		  objectType: $type
		  first: 5
		) {
		  edges {
			node {
			  objectId
			  owner
			  fields
			  dataType
			  createTime
			  digest
			  hasPublicTransfer
			  isUpdate
			  previousTransaction
			  status
			  storageRebate
			  type
			  updateTime
			  version
			  actionNumber
			}
		  }
		}
	  }
	`

	variables := map[string]interface{}{
		"type": "0x2::dynamic_field::Field<u64, 0xc5ee772bc6cabe728810e3282aef5333aaaa8cfd::dmens::Dmens>",
	}

	var out interface{}
	err := FetchGraphQL(query, "", variables, graphqlURL, &out)
	require.Nil(t, err)
	t.Log(out)
}
