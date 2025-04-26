package jsonwebtoken

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type tokenClaims struct {
	jwt.StandardClaims
	ClickCount int64 `json:"click_count"`
	UpdatedAt  int64 `json:"updated_at"`
}

func GenerateAccessToken(clickCount int64, updatedAt int64, jwtSigningKey string, accessTokenTTL time.Duration) (string, error) {
	claims := &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(accessTokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		ClickCount: clickCount,
		UpdatedAt:  updatedAt,
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessTokenStr, err := accessToken.SignedString([]byte(jwtSigningKey))
	if err != nil {
		return "", err
	}

	return accessTokenStr, nil
}

func ParseToken(accessToken string, signingKey string) (int64, int64, error) {
	return parseJWTToken(accessToken, signingKey)
}

func parseJWTToken(accessToken string, signingKey string) (int64, int64, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, 0, err
	}

	if !token.Valid {
		return 0, 0, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.ClickCount, claims.UpdatedAt, nil
}
