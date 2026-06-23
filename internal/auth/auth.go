package auth

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/alexedwards/argon2id"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		return "", err
	}

	return hashedPassword, err
}

func CheckPasswordHash(password, hash string) (bool, error) {
	matchPassword, err := argon2id.ComparePasswordAndHash(password, hash)

	return matchPassword, err
}

func MakeJWT(userID uuid.UUID, tokenSecret string, expiresIn time.Duration) (string, error) {
	claims := jwt.RegisteredClaims{
		Issuer:    "chirpy-access",
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresIn)),
		Subject:   userID.String(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(tokenSecret))

	return signedToken, err
}

func ValidateJWT(tokenString, tokenSecret string) (uuid.UUID, error) {
	claims := &jwt.RegisteredClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (any, error) {
		return []byte(tokenSecret), nil
	})
	if err != nil {
		return uuid.Nil, err
	}

	claimSbjct, err := uuid.Parse(claims.Subject)
	if err != nil {
		return uuid.Nil, err
	}
	return claimSbjct, nil
}

func GetBearerToken(headers http.Header) (string, error) {
	prefix := "Bearer "
	authString := headers.Get("Authorization")
	if authString == "" {
		return "", errors.New("authorization header missing")
	}

	if !strings.HasPrefix(authString, prefix) {
		return "", errors.New("no prefix")
	}

	token := strings.TrimPrefix(authString, prefix)
	token = strings.TrimSpace(token)

	if token == "" {
		return "", errors.New("token is empty")
	}

	return token, nil
}

func MakeRefreshToken() string {
	key := make([]byte, 32)

	rand.Read(key)

	return hex.EncodeToString(key)
}

func GetAPIKey(headers http.Header) (string, error) {
	prefix := "ApiKey "
	authString := headers.Get("Authorization")
	if authString == "" {
		return "", errors.New("authorization header missing")
	}

	if !strings.HasPrefix(authString, prefix) {
		return "", errors.New("no prefix")
	}

	apiKey := strings.TrimPrefix(authString, prefix)
	apiKey = strings.TrimSpace(apiKey)

	if apiKey == "" {
		return "", errors.New("token is empty")
	}

	return apiKey, nil
}
