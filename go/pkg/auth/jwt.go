package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/beckbikang/WriteCodeEveryDay/global"
)

type Claims struct {
	UserName string `json:"user_name"`
	jwt.StandardClaims
}

func GenerateToken(username string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(global.JWTSetting.JwtExpire) * time.Second)

	claims := Claims{
		UserName: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    global.JWTSetting.JwtSecret,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(global.JWTSetting.JwtSecret))
}

func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.JWTSetting.JwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}
