package token

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	models "github.com/vegadelalyra/go_jwt_authenticate/models"
)

const (
	JWTPrivateToken = "SecretTokenSecret"
	ip              = "190.151.192.116"
)

func GenerateToken(claims *models.JwtClaims, expirationTime time.Time) (string, error) {
	claims.ExpiresAt = expirationTime.Unix()
	claims.IssuedAt = time.Now().Unix()
	claims.Issuer = ip

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(JWTPrivateToken))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(tokenString, origin string) (bool, *models.JwtClaims) {
	claims := &models.JwtClaims{}
	token, err := getTokenFromString(tokenString, claims)
	if err != nil {
		return false, claims
	}
	if token.Valid {
		if e := claims.Valid(); e == nil {
			return true, claims
		}
	}
	return false, claims
}

func GetClaims(tokenString string) models.JwtClaims {
	claims := &models.JwtClaims{}

	_, err := getTokenFromString(tokenString, claims)
	if err != nil {
		return *claims
	}
	return *claims
}

func getTokenFromString(tokenString string, claims *models.JwtClaims) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method :%v", token.Header["alg"])
		}
		return []byte(JWTPrivateToken), nil
	})
}
