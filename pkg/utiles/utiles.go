package utiles

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var JWTSecret = []byte("ABABAAA")

type Claims struct {
	Id       uint   `json:"id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// GenerateToken 签发token
func GenerateToken(id uint, username, password string) (string, error) {
	notTime := time.Now()
	//时间戳 12 小时过期 需要重新登录
	expirTime := notTime.Add(12 * time.Hour)
	claims := Claims{
		Id:       id,
		UserName: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			//Unix()将int 转化为 time类型
			ExpiresAt: expirTime.Unix(),
			Issuer:    "todo_list",
		},
	}
	//报错 因为 SigningMethodES256 没有SignedString方法  应改为 SigningMethodHS256
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(JWTSecret)
	if err != nil {
		fmt.Println(err)
	}
	return token, err
}

// ParseToken 验证token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JWTSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
