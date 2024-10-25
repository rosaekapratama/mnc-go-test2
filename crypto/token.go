package crypto

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rosaekapratama/go-starter/log"
	"time"
)

// Generate JWT token (access and refresh)
func GenerateTokens(ctx context.Context, sub string, phoneNumber, secret string) (string, string, error) {
	byteSecret := []byte(secret)

	// Set token expiration times
	accessTokenExpiry := time.Now().Add(15 * time.Minute)    // 15 minutes for access token
	refreshTokenExpiry := time.Now().Add(7 * 24 * time.Hour) // 7 days for refresh token

	// Create claims for access token
	accessClaims := Claims{
		Sub:         sub,
		PhoneNumber: phoneNumber,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(accessTokenExpiry),
		},
	}

	// Create claims for refresh token (you can omit sensitive data like pin here)
	refreshClaims := Claims{
		Sub:         sub,
		PhoneNumber: phoneNumber,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(refreshTokenExpiry),
		},
	}

	// Generate access token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := accessToken.SignedString(byteSecret)
	if err != nil {
		log.Error(ctx, err)
		return "", "", err
	}

	// Generate refresh token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString(byteSecret)
	if err != nil {
		log.Error(ctx, err)
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}
