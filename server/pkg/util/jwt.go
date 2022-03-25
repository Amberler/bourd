package util

import (
	"api/pkg/setting"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte(setting.AppSetting.JwtSecret) // jwt秘钥

type Claims struct {
	ID uint
	jwt.StandardClaims
}

func GeterateToken(id uint, mobile string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(setting.AppSetting.JwtExpireTime * time.Hour)

	claims := Claims{
		id,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func ParseToken(token string) (*Claims, error) {
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
