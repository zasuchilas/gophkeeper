package secret

var _ Secret = (*LogoPass)(nil)

type LogoPass struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Info     string `json:"info"`
	Meta     string `json:"meta"`
}

func NewLogoPass(login, password, info, meta string) *LogoPass {
	return &LogoPass{Login: login, Password: password, Info: info, Meta: meta}
}

func NewEmptyLogoPass() *LogoPass {
	return &LogoPass{}
}

func (sec LogoPass) EncryptToBase64() ([]byte, error) {
	b, err := encodeToBase64[LogoPass](sec)
	return b, err
}

func (sec *LogoPass) DecryptFromBase64(data []byte) error {
	var next LogoPass
	err := decodeFromBase64[LogoPass](&next, data)
	if err != nil {
		return err
	}

	sec.Login = next.Login
	sec.Password = next.Password
	sec.Info = next.Info
	sec.Meta = next.Meta

	return nil
}
