package util

import (
	"api/pkg/setting"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	ID uint
	jwt.StandardClaims
}

func GenerateToken(id uint) (string, error) {
	jwtSecret := []byte(setting.AppSetting.JwtSecret)
	nowTime := time.Now()
	expireTime := nowTime.Add(setting.AppSetting.JwtExpireTime * time.Hour)

	// 创建 Claims
	claims := Claims{
		id,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	jwtSecret := []byte(setting.AppSetting.JwtSecret)
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
