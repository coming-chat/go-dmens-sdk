package dmens

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueryTwittersMyFollowed(t *testing.T) {
	poster := DefaultPoster(t)
	res, err := poster.QueryTwittersMyFollowed(10, 0)
	require.Nil(t, err)
	t.Log(res)
}

func TestQueryTrendTwitters(t *testing.T) {
	poster := DefaultPoster(t)
	res, err := poster.QueryTrendTwittersList(10, 0)
	require.Nil(t, err)
	t.Log(res)
}

func TestQueryTwitterByNoteId(t *testing.T) {
	poster := DefaultPoster(t)
	res, err := poster.QueryTwitterByNoteId("0xc4dbf29b5513f7695d6370e094c5ad03fb44acc2", 10, 0)
	require.Nil(t, err)
	t.Log(res)
}
