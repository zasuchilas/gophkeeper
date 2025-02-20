package secret

import (
	"reflect"
	"testing"
)

func Test_encodeToBase64(t *testing.T) {
	type args struct {
		v LogoPass
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "successful",
			args: args{
				v: LogoPass{
					Login:    "Anna",
					Password: "123",
					Info:     "qwe",
					Meta:     "007",
				},
			},
			want:    []byte("eyJsb2dpbiI6IkFubmEiLCJwYXNzd29yZCI6IjEyMyIsImluZm8iOiJxd2UiLCJtZXRhIjoiMDA3In0K"),
			wantErr: false,
		},
		{
			name: "successful",
			args: args{
				v: LogoPass{
					Login:    "Анна",
					Password: "123",
					Info:     "qwe",
					Meta:     "007",
				},
			},
			want:    []byte("eyJsb2dpbiI6ItCQ0L3QvdCwIiwicGFzc3dvcmQiOiIxMjMiLCJpbmZvIjoicXdlIiwibWV0YSI6IjAwNyJ9"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := encodeToBase64[LogoPass](tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("encodeToBase64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("encodeToBase64() got = %s, want %v", got, tt.want)
			}
		})
	}
}

func Test_decodeFromBase64_LogoPass(t *testing.T) {
	type args struct {
		v   *LogoPass
		enc []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "successful",
			args: args{
				v:   NewEmptyLogoPass(),
				enc: []byte("eyJsb2dpbiI6IkFubmEiLCJwYXNzd29yZCI6IjEyMyIsImluZm8iOiJxd2UiLCJtZXRhIjoiMDA3In0K"),
			},
			wantErr: false,
		},
		{
			name: "successful rus",
			args: args{
				v:   NewEmptyLogoPass(),
				enc: []byte("eyJsb2dpbiI6ItCQ0L3QvdCwIiwicGFzc3dvcmQiOiIxMjMiLCJpbmZvIjoicXdlIiwibWV0YSI6IjAwNyJ9"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := decodeFromBase64[LogoPass](tt.args.v, tt.args.enc); (err != nil) != tt.wantErr {
				t.Errorf("decodeFromBase64() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
