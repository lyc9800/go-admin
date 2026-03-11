package helper

import (
	"server/define"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// 生成token
func GenerateToken(id uint, name string, expireAt int64) (string, error) {
	uc := define.UserClaim{
		Id:   id,
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(expireAt, 0)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := token.SignedString([]byte(define.Jwtkey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
