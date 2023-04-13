package dmens

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_rawNoteAction_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   interface{}
		want    NoteAction
		wantErr bool
	}{
		{
			name:    "int value",
			input:   100,
			want:    100,
			wantErr: false},
		{
			name:    "string value",
			input:   "123",
			want:    123,
			wantErr: false},
		{
			name:    "error string value",
			input:   "11abc",
			want:    0,
			wantErr: true},
		{
			name:    "error input",
			input:   map[string]string{"name": "comingchat"},
			want:    0,
			wantErr: true},
		{
			name:    "any byte",
			input:   []byte{0x0, 0x2, 0x3, 0xff},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(tt.input)
			require.NoError(t, err)

			var res rawNoteAction
			err = json.Unmarshal(data, &res)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, res.Value, tt.want)
			}
		})
	}
}

func Test_rawNoteAction_MarshalJSON(t *testing.T) {
	raw := rawNoteAction{Value: 123}

	data, err := json.Marshal(raw)
	require.NoError(t, err)

	var new rawNoteAction
	err = json.Unmarshal(data, &new)
	require.NoError(t, err)
	require.Equal(t, raw, new)
}
