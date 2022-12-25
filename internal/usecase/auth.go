package usecase

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/glhfuck/turbo-waffle/internal/domain"
	"github.com/glhfuck/turbo-waffle/internal/infrastructure/repository"
)

const (
	salt       = "jHFGYkV7TvruGlni7ynIr5"
	signingKey = "JKv89Hfdf98ffSD"
	tokenTTL   = 12 * time.Hour
)

type Authorization interface {
	CreateUser(u domain.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type authUsecase struct {
	repo repository.Authorization
}

func (au *authUsecase) CreateUser(u domain.User) (int, error) {
	u.Password = au.passwordHash(u.Password)
	return au.repo.CreateUser(u)
}

func (au *authUsecase) GenerateToken(username, password string) (string, error) {
	u, err := au.repo.GetUser(username, au.passwordHash(password))

	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		u.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func (au *authUsecase) ParseToken(token string) (int, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := parsedToken.Claims.(*tokenClaims)

	if !ok {
		return 0, errors.New("token claims are invalid")
	}

	return claims.UserId, nil
}

func (au *authUsecase) passwordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func newAuthUsecase(repo repository.Authorization) *authUsecase {
	return &authUsecase{repo: repo}
}
