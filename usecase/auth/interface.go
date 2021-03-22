package auth

import (
	"StoreManager-DDD/entity"
	"github.com/dgrijalva/jwt-go"
)

type auth interface {
	VerifyToken(token string) (bool, error)
	GenerateToken(email string, userId entity.ID) (*jwt.Token, error)
}

type Repository interface {
	auth
}
