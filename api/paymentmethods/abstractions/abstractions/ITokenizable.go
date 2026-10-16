package abstractions

import "github.com/globalpayments/go-sdk/api/abstractions"

type ITokenizable interface {
	GetToken() string
	SetToken(token string)
	Tokenize() (abstractions.ExecutableGateway, error)
}
