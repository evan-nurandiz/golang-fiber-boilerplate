package services

import (
	"time"

	"github.com/evan_nurandiz/go_fiber_boilerplate/config"
	"github.com/golang-jwt/jwt/v4"
)

type JWTClaim struct {
	User_id int
	Name    string
	Email   string
	jwt.RegisteredClaims
}

type JWTpayload struct {
	User_id int
	Name    string
	Email   string
}

func GenerateToken(payload JWTpayload) (jwtError string, accessToken string, refreshToken string) {
	expAccesToken := time.Now().Add(time.Hour * 3)
	expRefreshToken := time.Now().Add(time.Hour * 100)
	accessTokenClaims := JWTClaim{
		User_id: payload.User_id,
		Name:    payload.Email,
		Email:   payload.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    config.GetConfig("APP_NAME"),
			ExpiresAt: jwt.NewNumericDate(expAccesToken),
		},
	}

	refreshTokenClaims := JWTClaim{
		User_id: payload.User_id,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    config.GetConfig("APP_NAME"),
			ExpiresAt: jwt.NewNumericDate(expRefreshToken),
		},
	}

	accessToken, errAccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims).SignedString([]byte(config.GetConfig("JWT_KEY")))

	if errAccessToken != nil {
		return errAccessToken.Error(), "", ""
	}

	refreshToken, errRefreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims).SignedString([]byte(config.GetConfig("JWT_KEY")))

	if errRefreshToken != nil {
		return errRefreshToken.Error(), "", ""
	}

	return "", accessToken, refreshToken

}
