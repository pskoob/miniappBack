package jsonwebtoken

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type tapTokenClaims struct {
	jwt.StandardClaims
	ClickCount int64 `json:"click_count"`
	UpdatedAt  int64 `json:"updated_at"`
}

type energyTokenClaims struct {
	jwt.StandardClaims
	Energy    int64 `json:"energy"`
	UpdatedAt int64 `json:"updated_at"`
}

func ParseTapToken(tapToken string, signingKey string) (int64, int64, error) {
	return parseJWTTapToken(tapToken, signingKey)
}

func parseJWTTapToken(tapToken string, signingKey string) (int64, int64, error) {
	token, err := jwt.ParseWithClaims(tapToken, &tapTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
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

	claims, ok := token.Claims.(*tapTokenClaims)
	if !ok {
		return 0, 0, errors.New("token claims are not of type *tapTokenClaims")
	}

	return claims.ClickCount, claims.UpdatedAt, nil
}

func ParseEnergyToken(energyToken string, signingKey string) (int64, int64, error) {
	return parseJWTEnergyToken(energyToken, signingKey)
}

func parseJWTEnergyToken(energyToken string, signingKey string) (int64, int64, error) {
	token, err := jwt.ParseWithClaims(energyToken, &energyTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
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

	claims, ok := token.Claims.(*energyTokenClaims)
	if !ok {
		return 0, 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.Energy, claims.UpdatedAt, nil
}

func GenerateTapToken(clickCount int64, updatedAt int64, jwtSigningKey string, accessTokenTTL time.Duration) (string, error) {
	claims := &tapTokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(accessTokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		ClickCount: clickCount,
		UpdatedAt:  updatedAt,
	}

	tapToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tapTokenStr, err := tapToken.SignedString([]byte(jwtSigningKey))
	if err != nil {
		return "", err
	}

	return tapTokenStr, nil
}

func GenerateEnergyToken(energy int64, updatedAt int64, jwtSigningKey string, accessTokenTTL time.Duration) (string, error) {
	claims := &energyTokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(accessTokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		Energy:    energy,
		UpdatedAt: updatedAt,
	}

	energyToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	energyTokenStr, err := energyToken.SignedString([]byte(jwtSigningKey))
	if err != nil {
		return "", err
	}

	return energyTokenStr, nil
}
