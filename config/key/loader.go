package key

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/agungdwiprasetyo/go-utils/debug"
	jwt "github.com/dgrijalva/jwt-go"
)

// LoadPrivateKey load rsa private key
func LoadPrivateKey() *rsa.PrivateKey {
	signBytes, err := ioutil.ReadFile(fmt.Sprintf("%s/config/key/private.key", os.Getenv("APP_PATH")))
	if err != nil {
		debug.Println("Error when load private key.", err)
		os.Exit(1)
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		debug.Println("Error when load private key.", err)
		os.Exit(1)
	}
	return privateKey
}

// LoadPublicKey load rsa public key
func LoadPublicKey() *rsa.PublicKey {
	verifyBytes, err := ioutil.ReadFile(fmt.Sprintf("%s/config/key/public.pem", os.Getenv("APP_PATH")))
	if err != nil {
		debug.Println("Error when load public key.", err)
		os.Exit(1)
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		debug.Println("Error when load public key.", err)
		os.Exit(1)
	}
	return publicKey
}
