package infrastructure

import (
	"crypto/rand"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4/middleware"
)

type jwtSecret struct {
	gorm.Model
	secret []byte
}

func newJWTSecret() jwtSecret {
	b := make([]byte, 128)
	_, err := rand.Read(b)
	if err != nil {
		panic("failed to generate JWT secret")
	}

	return jwtSecret{secret: b}
}

// JWTProvider generates tokens for users
type JWTProvider struct {
	secret []byte
}

// NewJWTProvider returns JWT provider
func NewJWTProvider(secret []byte) *JWTProvider {
	router.Use(middleware.JWT(secret))
	return &JWTProvider{secret}
}

// Generate tokens for users
func (provider *JWTProvider) Generate(userID string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS512)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 48)

	return token.SignedString(provider.secret)
}
