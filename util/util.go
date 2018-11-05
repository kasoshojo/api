package util

import (
	"crypto/rsa"
	"crypto/sha1"
	"fmt"
	"log"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/goadesign/goa/uuid"
)

func DateEqual(date1, date2 time.Time) bool {
	y1, m1, d1 := date1.Date()
	y2, m2, d2 := date2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

func HashPassword(password string) string {
	data := []byte(password)
	chksum := sha1.Sum(data)
	hash := fmt.Sprintf("%x", chksum)
	return hash
}

func GenerateJWTToken(id int, nickname *string, privatekey *rsa.PrivateKey, scope []string) string {
	// Generate JWT
	nick := "user"

	if nickname != nil && len(*nickname) > 0 {
		nick = *nickname
	}

	token := jwtgo.New(jwtgo.SigningMethodRS512)
	in10d := time.Now().Add(time.Duration(10) * 24 * time.Hour).Unix()
	token.Claims = jwtgo.MapClaims{
		"iss":    "Kasoshojo",
		"aud":    "Users",
		"exp":    in10d,
		"jti":    uuid.NewV4().String(),
		"iat":    time.Now().Unix(),
		"nbf":    2,
		"sub":    "user",
		"user":   id,
		"scopes": scope,
		"name":   nick,
	}
	signedToken, err := token.SignedString(privatekey)
	if err != nil {
		log.Println("failed to sign token: %s", err) // internal error
	}
	return signedToken
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
