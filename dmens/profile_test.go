package dmens

import (
	"reflect"
	"testing"

	"github.com/coming-chat/wallet-SDK/core/sui"
	"github.com/stretchr/testify/require"
)

func TestRegister(t *testing.T) {
	profile := Profile{
		Name: "zhiuaanngggg",
	}

	acc, err := sui.NewAccountWithMnemonic(M1)
	require.Nil(t, err)

	poster := DefaultPoster(t)
	poster.Address = acc.Address()

	p, err := poster.CheckProfile(&profile)
	require.Nil(t, err)

	txn, err := poster.Register(p)
	require.Nil(t, err)

	signedTxn, err := txn.SignWithAccount(acc)
	require.Nil(t, err)

	hash, err := poster.chain.SendRawTransaction(signedTxn.Value)
	require.Nil(t, err)
	t.Log(hash)
}

func TestPoster_CheckProfile(t *testing.T) {
	type fields struct {
		Configuration *Configuration
		PosterConfig  *PosterConfig
		chain         *sui.Chain
	}
	type args struct {
		profile *Profile
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ValidProfile
		wantErr bool
	}{
		{
			name: "test1",
			fields: fields{
				Configuration: &Configuration{
					ProfileCheckUrl: "https://coming-zero-dmens-pre.coming.chat/profile/check/",
				},
				PosterConfig: NewPosterConfig("0x6fc6148816617c3c3eccb1d09e930f73f6712c9c"),
				chain:        nil,
			},
			args: args{
				profile: &Profile{
					Name:   "Gkirito",
					Bio:    "Hello",
					Avatar: "ipfs://bafkreiahy2mdbxcvf4ftsqpfykbt7o37elvyn7uknmj7bxqrgdii5aabri",
				},
			},
			want: &ValidProfile{
				Profile:   "{\"name\":\"Gkirito\",\"bio\":\"Hello\",\"avatar\":\"ipfs://bafkreiahy2mdbxcvf4ftsqpfykbt7o37elvyn7uknmj7bxqrgdii5aabri\"}",
				Signature: "0xd485020c6ac369e6f2b28be2dcca24ebfd827c53893b6462e9e65cf16dba3cedf004e8740b8c8c3579a4391269b9e103bcfc39627c6af729abb7675bc8004301",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Poster{
				Configuration: tt.fields.Configuration,
				PosterConfig:  tt.fields.PosterConfig,
				chain:         tt.fields.chain,
			}
			got, err := p.CheckProfile(tt.args.profile)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckProfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckProfile() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetNFTAvatar(t *testing.T) {
	poster := DefaultPoster(t)
	nftId := "0xb6bb02c3e96df53b2bca7933fb14f2ca610ec737" // dmens nft
	_, err := poster.SetNftAvatar(nftId)
	require.Error(t, err)
	// t.Log(txn)
}

func TestRemoveNFTAvatar(t *testing.T) {
	poster := DefaultPoster(t)
	nftId := "0x94391ecd915b8df4e8273df1d9aa5c5e1f10f909"
	nftType := "0x8521368cac606257ce902ccf1735ac41d9acc709::capy::Capy"
	txn, err := poster.RemoveNftAvatar(&NFTAvatar{Id: nftId, Type: nftType})
	require.Nil(t, err)
	t.Log(txn)

	// hash := signAndSendTxn(t, poster, M1, txn)
	// t.Log(hash)
}

func signAndSendTxn(t *testing.T, poster *Poster, mnemonic string, txn *sui.Transaction) string {
	acc, err := sui.NewAccountWithMnemonic(mnemonic)
	require.Nil(t, err)

	signedTxn, err := txn.SignWithAccount(acc)
	require.Nil(t, err)

	hash, err := poster.chain.SendRawTransaction(signedTxn.Value)
	require.Nil(t, err)

	return hash
}
