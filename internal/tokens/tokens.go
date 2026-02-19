package tokens

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"time"
)

const (
	ScopeAuth = "authentication"
)

type Token struct {
	PlainText string    `json:"token"`
	Hash      []byte    `json:"_"`
	UserID    int       `json:"-"`
	Expiry    time.Time `json:"expiry"`
	Scope     string    `json:"-"`
}

func GenerateToken(userID int, ttl time.Duration, scope string) (*Token, error) {
	token := &Token{
		UserID: userID,
		Expiry: time.Now().Add(ttl),
		Scope:  scope,
	}

	emptyBytes := make([]byte, 32)
	_, err := rand.Read(emptyBytes)
	if err != nil {
		return nil, err
	}

	token.PlainText = base64.StdEncoding.WithPadding(base64.NoPadding).EncodeToString(emptyBytes)
	hash := sha256.Sum256([]byte(token.PlainText))
	token.Hash = hash[:]
	return token, nil
}
