package auth

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type CustomClaims struct {
	jwt.RegisteredClaims
	Username string `json:"username"`
}

func HashPassword(passord string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(passord), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CheckPasswordHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func MakeJWT(username string, userId int, tokenSecret string, expiresIn time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "scrapy",
			Subject:   strconv.Itoa(userId),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresIn)),
		},
	})
	return token.SignedString([]byte(tokenSecret))
}

func ValidateJWT(tokenString, tokenSecret string) (int32, error) {
	claims := jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(tokenSecret), nil
	})
	if err != nil {
		return 0, err
	}
	subject, err := token.Claims.GetSubject()
	if err != nil {
		return 0, err
	}
	issuer, err := token.Claims.GetIssuer()
	if err != nil {
		return 0, err
	}
	if issuer != string("scrapy") {
		return 0, errors.New("invalid issuer")
	}
	id, err := strconv.Atoi(subject)
	if err != nil {
		return 0, fmt.Errorf("invalid user ID '%s'", subject)
	}
	return int32(id), nil
}

func MakeRefreshToken() (string, error) {
	token := make([]byte, 32)
	_, err := rand.Read(token)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(token), nil
}

func GetBearerToken(headers http.Header) (string, error) {
	header := headers.Get("Authorization")
	parts := strings.Split(header, " ")
	if len(parts) != 2 {
		return "", errors.New("no auth header found")
	}
	if parts[0] != "Bearer" {
		return "", errors.New("wrong auth method")
	}
	return parts[1], nil
}
