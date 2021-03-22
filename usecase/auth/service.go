package auth

import (
	"StoreManager-DDD/entity"
	"StoreManager-DDD/usecase/user"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Service struct {
	repo Repository
}

type AuthClaims struct {
	jwt.StandardClaims
	Email string `json:"email"`
	ID entity.ID `json:"id"`
}

type AuthUseCase struct {
	userRepo       user.Repository
	signingKey     []byte
	expireDuration time.Duration
}

func NewAuthUseCase(
	userRepo user.Repository,
	signingKey []byte,
	tokenTTLSeconds time.Duration) *AuthUseCase {
	return &AuthUseCase{
		userRepo:       userRepo,
		signingKey:     signingKey,
		expireDuration: time.Second * tokenTTLSeconds,
	}
}

// SignIn handles user sign in with email and password
func (a *AuthUseCase) SignIn(email, password string) (string, error) {
	user_, err := a.userRepo.GetByEmail(email)
	if err != nil {
		return "", entity.ErrNotFound
	}
	err = user_.ValidatePassword(password)
	if err != nil {
		return "", entity.ErrInvalidUserEmailOrPassword
	}
	claims := AuthClaims{
		Email: user_.Email,
		ID: user_.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(a.signingKey)
}

// ValidateToken handles token validation
func (a *AuthUseCase) ValidateToken(encodedToken string)(*AuthClaims, error) {
	token, err := jwt.ParseWithClaims(
		encodedToken,
		&AuthClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return a.signingKey, nil
		},
	)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*AuthClaims)
	if !ok {
		err = errors.New("couldn't parse claims")
		return nil, err
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token has expired")
		return nil, err
	}

	return claims, nil
}