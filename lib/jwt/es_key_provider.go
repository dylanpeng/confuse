package jwt

import (
	"crypto/ecdsa"
	gjwt "github.com/golang-jwt/jwt/v4"
)

const (
	SignMethodES256 = "ES256"
	SignMethodES384 = "ES384"
	SignMethodES512 = "ES512"
)

var esSignMethod = map[string]gjwt.SigningMethod{
	SignMethodES256: gjwt.SigningMethodES256,
	SignMethodES384: gjwt.SigningMethodES384,
	SignMethodES512: gjwt.SigningMethodES512,
}

type EsTokenProducer struct {
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
	signMethod gjwt.SigningMethod
}

func (p *EsTokenProducer) GetSigningMethod() gjwt.SigningMethod {
	return p.signMethod
}

func (p *EsTokenProducer) GetPrivateKey() any {
	return p.privateKey
}

func (p *EsTokenProducer) GetPublicKey(token *gjwt.Token) (interface{}, error) {
	return p.publicKey, nil
}

func newEsTokenProducer(privateKeyString, publicKeyString, signMethod string) (result *EsTokenProducer, err error) {
	privateKey, err := gjwt.ParseECPrivateKeyFromPEM([]byte(privateKeyString))
	if err != nil {
		return
	}

	publicKey, err := gjwt.ParseECPublicKeyFromPEM([]byte(publicKeyString))
	if err != nil {
		return
	}

	result = &EsTokenProducer{
		privateKey: privateKey,
		publicKey:  publicKey,
		signMethod: gjwt.SigningMethodES256,
	}

	if signMethod, ok := esSignMethod[signMethod]; ok {
		result.signMethod = signMethod
	}

	return
}
