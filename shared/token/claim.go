package token

import (
	"github.com/agungdwiprasetyo/agungdpcms/src/user/domain"
	jwt "github.com/dgrijalva/jwt-go"
)

// Claim for token claim data
type Claim struct {
	User *domain.User `json:"user,omitempty"`
	jwt.StandardClaims
}

// NewClaim construct token claim data
func NewClaim(dataUser *domain.User) *Claim {
	aud := "guest"
	if dataUser.Role != nil {
		aud = dataUser.Role.Slug
	}
	userClaims := domain.User{
		ID:       dataUser.ID,
		Username: dataUser.Username,
	}

	cl := new(Claim)
	cl.Audience = aud
	cl.User = &userClaims
	return cl
}
