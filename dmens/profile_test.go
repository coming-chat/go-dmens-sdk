package dmens

import (
	"github.com/coming-chat/wallet-SDK/core/sui"
	"reflect"
	"testing"
)

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
