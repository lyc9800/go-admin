package helper

import (
	"errors"
	"server/define"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// 生成token
func GenerateToken(id uint, roleId uint, name string, expireAt int64) (string, error) {
	uc := define.UserClaim{
		Id:     id,
		Name:   name,
		RoleId: roleId,
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

func AnalyzeToken(token string) (*define.UserClaim, error) {
	uc := new(define.UserClaim)
	claims, err := jwt.ParseWithClaims(token, uc, func(token *jwt.Token) (interface{}, error) {
		return []byte(define.Jwtkey), nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return uc, errors.New("token不正确")
	}
	return uc, nil
}
