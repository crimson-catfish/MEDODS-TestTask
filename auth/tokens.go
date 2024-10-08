package auth

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/golang-jwt/jwt"
	"os"
	"time"
)

const (
	jwtExpirationTimeMinuets       = 1
	RefreshTokenExpirationTimeDays = 30
	refreshTokenLength             = 32
)

type AccessTokenClaims struct {
	IP string
	jwt.StandardClaims
}

func GenerateTokens(guid, clientIP string) (string, string, error) {
	expirationTime := time.Now().Add(jwtExpirationTimeMinuets * time.Minute)
	clm := &AccessTokenClaims{
		IP: clientIP,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Id:        guid,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, clm)

	accessToken, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", "", err
	}

	bytes := make([]byte, refreshTokenLength)
	_, err = rand.Read(bytes)
	if err != nil {
		return "", "", err
	}

	refreshToken := base64.URLEncoding.EncodeToString(bytes)

	return accessToken, refreshToken, nil
}
