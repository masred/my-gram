package helper

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/masred/my-gram/app/exception"
)

type UserClaims struct {
	UserID   uuid.UUID `json:"id"`
	Username string    `json:"name"`
	Email    string    `json:"email"`
	jwt.RegisteredClaims
}

type UserJWTService struct {
	SecretKey []byte
}

func NewUserJWTService(secretKey []byte) *UserJWTService {
	return &UserJWTService{SecretKey: secretKey}
}

func (s UserJWTService) GenerateUserToken(id uuid.UUID, username, email string) (encodedToken string, err error) {
	claims := &UserClaims{
		UserID:   id,
		Username: username,
		Email:    email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "mygram",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	encodedToken, err = token.SignedString(s.SecretKey)
	if err != nil {
		return
	}

	return
}

func (s UserJWTService) ParseUserToken(tokenString string) (parsedToken *jwt.Token, err error) {
	parsedToken, err = jwt.ParseWithClaims(tokenString, &UserClaims{}, s.keyFunc)
	if err != nil {
		return
	}

	if !parsedToken.Valid {
		return nil, exception.ErrInvalidAuthToken
	}

	return
}

func (s UserJWTService) keyFunc(token *jwt.Token) (any, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, exception.ErrInvalidAuthToken
	}
	return s.SecretKey, nil
}
