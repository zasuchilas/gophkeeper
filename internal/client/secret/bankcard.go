package secret

type BankCard struct {
	Number string `json:"number"`
	Exp    string `json:"exp"`
	Cvv    int    `json:"cvv"`
	Owner  string `json:"owner"`
	Bank   string `json:"bank"`
	Meta   string `json:"meta"`
}

func NewBankCard(number string, exp string, cvv int, owner string, bank string) *BankCard {
	return &BankCard{Number: number, Exp: exp, Cvv: cvv, Owner: owner, Bank: bank}
}
