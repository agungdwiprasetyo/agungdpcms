package helper

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

// ComputeHmac256 to encrypt string
func ComputeHmac256(str string) string {
	key := []byte(salt)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
