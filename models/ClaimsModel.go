package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtClaims struct {
	CompanyID string `json:"companyId,omitempty"`
	Username  string `json:"username,omitempty"`
	Roles     []int  `json:"roles,omitempty"`
	jwt.StandardClaims
}

const ip = "190.151.192.116" // 192.168.0.107

func (claims *JwtClaims) Valid() error {
	var now = time.Now().Unix()
	if claims.VerifyExpiresAt(now, true) && claims.VerifyIssuer(ip, true) {
		return nil
	}
	return fmt.Errorf("token is invalid")
}

func (claims *JwtClaims) VerifyAudience(origin string) bool {
	return strings.Compare(claims.Audience, origin) == 0
}
