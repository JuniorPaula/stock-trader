package models

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	Email    string `json:"email,omitempty" bson:"email,omitempty"`
	Password string `json:"password,omitempty" bson:"password,omitempty"`
}

type Credentials struct {
	Email     string    `json:"email,omitempty" bson:"email,omitempty"`
	Token     string    `json:"token,omitempty" bson:"token,omitempty"`
	UserID    string    `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Username  string    `json:"username,omitempty" bson:"username,omitempty"`
	Hash      []byte    `json:"hash,omitempty" bson:"hash,omitempty"`
	Scope     string    `json:"scope,omitempty" bson:"scope,omitempty"`
	ExpiresAt time.Time `json:"expires_at,omitempty" bson:"expires_at,omitempty"`
}

func (a *Auth) PasswordMatch(hash, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}
	return true, nil
}

func GenerateToken(userID, username, email string, ttl time.Duration, scope string) (*Credentials, error) {
	token := &Credentials{
		UserID:    userID,
		Username:  username,
		Email:     email,
		Scope:     scope,
		ExpiresAt: time.Now().Add(ttl),
	}

	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return &Credentials{}, err
	}

	token.Token = base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randomBytes)
	hash := sha256.Sum256([]byte(token.Token))
	token.Hash = hash[:]
	return token, nil
}
