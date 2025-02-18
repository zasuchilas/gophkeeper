package secret

import (
	"reflect"
	"testing"
)

func TestLogoPass_EncryptToBase64(t *testing.T) {
	type fields struct {
		Login    string
		Password string
		Info     string
		Meta     string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "successful",
			fields: fields{
				Login:    "Анна",
				Password: "123",
				Info:     "qwe",
				Meta:     "007",
			},
			want:    []byte("eyJsb2dpbiI6ItCQ0L3QvdCwIiwicGFzc3dvcmQiOiIxMjMiLCJpbmZvIjoicXdlIiwibWV0YSI6IjAwNyJ9"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sec := &LogoPass{
				Login:    tt.fields.Login,
				Password: tt.fields.Password,
				Info:     tt.fields.Info,
				Meta:     tt.fields.Meta,
			}
			got, err := sec.EncryptToBase64()
			if (err != nil) != tt.wantErr {
				t.Errorf("EncryptToBase64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EncryptToBase64() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogoPass_DecryptFromBase64(t *testing.T) {
	type fields struct {
		Login    string
		Password string
		Info     string
		Meta     string
	}
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "successful",
			fields: fields{
				Login:    "Анна",
				Password: "123",
				Info:     "qwe",
				Meta:     "007",
			},
			args: args{
				data: []byte("eyJsb2dpbiI6ItCQ0L3QvdCwIiwicGFzc3dvcmQiOiIxMjMiLCJpbmZvIjoicXdlIiwibWV0YSI6IjAwNyJ9"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sec := NewEmptyLogoPass()
			expected := &LogoPass{
				Login:    tt.fields.Login,
				Password: tt.fields.Password,
				Info:     tt.fields.Info,
				Meta:     tt.fields.Meta,
			}
			if err := sec.DecryptFromBase64(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("DecryptFromBase64() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(sec, expected) {
				t.Errorf("EncryptToBase64() got = %v, want %v", sec, expected)
			}
		})
	}
}
