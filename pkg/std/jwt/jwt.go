package jwt

import (
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secret string

func init() {
	secret = os.Getenv("JWT_SECRET")

	if secret == "" {
		log.Fatal("error: JWT_SECRET not defined")
	}
}

func New(subject interface{}, audience []string) *jwt.Token {
	issuedAt := time.Now().Format(time.UnixDate)

	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "account",
		"sub": subject,
		"aud": audience,
		"iat": issuedAt,
	})
}

func NewSigned(subject interface{}, audience []string) (string, error) {
	signedToken, err := New(subject, audience).SignedString([]byte(secret))
	if err != nil {
		log.Printf("error: cannot create signed JWT token (%s)", err)
	}
	return signedToken, err
}
