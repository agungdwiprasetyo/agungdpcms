package token

// Token module abstraction
type Token interface {
	Generate(cl *Claim) (string, error)
	Refresh(tokenString string) (string, error)
	Extract(tokenString string) (*Claim, bool)
}
