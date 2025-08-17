package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type JWTManager struct {
	accessSecret  string
	refreshSecret string
	accessExpire  time.Duration
	refreshExpire time.Duration
}

func NewJWTManager(accessSecret, refreshSecret string, accessExpire, refreshExpire time.Duration) *JWTManager {
	return &JWTManager{
		accessSecret:  accessSecret,
		refreshSecret: refreshSecret,
		accessExpire:  accessExpire,
		refreshExpire: refreshExpire,
	}
}

// GenerateTokenPair membuat access token dan refresh token
func (j *JWTManager) GenerateTokenPair(userID int, email, role string) (*TokenPair, error) {
	// Generate access token
	accessToken, err := j.generateToken(userID, email, role, j.accessSecret, j.accessExpire)
	if err != nil {
		return nil, err
	}

	// Generate refresh token
	refreshToken, err := j.generateToken(userID, email, role, j.refreshSecret, j.refreshExpire)
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// generateToken membuat JWT token
func (j *JWTManager) generateToken(userID int, email, role, secret string, expire time.Duration) (string, error) {
	claims := Claims{
		UserID: userID,
		Email:  email,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expire)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// ValidateAccessToken memvalidasi access token
func (j *JWTManager) ValidateAccessToken(tokenString string) (*Claims, error) {
	return j.validateToken(tokenString, j.accessSecret)
}

// ValidateRefreshToken memvalidasi refresh token
func (j *JWTManager) ValidateRefreshToken(tokenString string) (*Claims, error) {
	return j.validateToken(tokenString, j.refreshSecret)
}

// validateToken memvalidasi JWT token
func (j *JWTManager) validateToken(tokenString, secret string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}