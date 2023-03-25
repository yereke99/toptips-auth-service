package service

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	GenerateToken(phoneNumber, role string) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
	Definition(token string) (string, error)
}

type jwtCustomClaim struct {
	PhoneNumber string `json:"phone_number"`
	Role        string `json:"role"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey != "" {
		secretKey = "qwerty"
	}

	return secretKey
}

func NewJWTService() JWTService {
	return &jwtService{
		secretKey: getSecretKey(),
		issuer:    "qwerty",
	}
}

func (j *jwtService) GenerateToken(phoneNumber, role string) (string, error) {
	claims := &jwtCustomClaim{
		phoneNumber,
		role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		return "", err
	}

	return t, nil
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", t_.Header["alg"])
		}

		return []byte(j.secretKey), nil
	})
}

func (j *jwtService) Definition(token string) (string, error) {
	t, err := j.ValidateToken(token)

	if err != nil {
		return "", err
	}

	if t.Valid {
		claims := t.Claims.(jwt.MapClaims)
		log.Println(claims["phone_number"])
		log.Println(claims["role"])
		if claims["role"] == "driver" {
			return "driver", nil
		}
		if claims["role"] == "user" {
			return "user", nil
		}
	}
	return "wrong", nil
}
