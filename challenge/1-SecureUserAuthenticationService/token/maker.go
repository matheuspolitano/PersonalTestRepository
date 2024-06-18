package token

type Maker interface {
	CreateToken(*Payload) (string, error)
	CheckToken(string) (*Payload, error)
}
