package token

import (
	"crypto/rsa"
	"fmt"
	"strings"
	"time"

	"github.com/agungdwiprasetyo/agungdpcms/src/user/domain"
	jwt "github.com/dgrijalva/jwt-go"
)

type tokenImpl struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
	age        time.Duration
}

// New construct token
func New(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey, age time.Duration) Token {
	tok := new(tokenImpl)
	tok.privateKey = privateKey
	tok.publicKey = publicKey
	tok.age = age
	return tok
}

func (t *tokenImpl) Generate(cl *Claim) (string, error) {
	token := jwt.New(jwt.SigningMethodRS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(t.age).Unix()
	claims["iat"] = time.Now().Unix()
	claims["aud"] = cl.Audience
	claims["account"] = cl.User
	token.Claims = claims
	tokenString, err := token.SignedString(t.privateKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (t *tokenImpl) Refresh(tokenString string) (string, error) {
	splitToken := strings.Split(tokenString, ".")
	if len(splitToken) != 3 {
		return "", fmt.Errorf("Token Invalid")
	}
	result, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return t.publicKey, nil
	})
	if !result.Valid {
		return "", fmt.Errorf("Token Invalid")
	}
	claims := result.Claims.(jwt.MapClaims)

	tok := jwt.New(jwt.SigningMethodRS256)
	claims["exp"] = time.Now().Add(t.age).Unix()
	tok.Claims = claims
	tokenString, err := tok.SignedString(t.privateKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (t *tokenImpl) Extract(tokenString string) (*Claim, bool) {
	splitToken := strings.Split(tokenString, ".")
	if len(splitToken) != 3 {
		return nil, false
	}
	result, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return t.publicKey, nil
	})
	if err != nil {
		return nil, false
	}

	claims := new(Claim)
	if result.Valid {
		mapClaims := result.Claims.(jwt.MapClaims)
		claims.Audience = fmt.Sprint(mapClaims["aud"])
		account, ok := mapClaims["account"].(map[string]interface{})
		if !ok {
			return claims, false
		}
		claims.User = &domain.User{
			ID:       int(account["id"].(float64)),
			Username: fmt.Sprint(account["username"]),
		}
	}
	return claims, result.Valid
}
