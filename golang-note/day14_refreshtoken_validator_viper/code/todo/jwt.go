package main

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// jwt 相关代码

const TokenExpireDuration = time.Hour * 1

var MySecret = []byte("夏天夏天悄悄过去") // 加盐

// MyClaims 自定义负载（payload）数据
type MyClaims struct {
	// 根据业务需求定义私有字段
	Uid  int64  `json:"uid"`
	Name string `json:"name"`

	jwt.StandardClaims // JWT官方字段
}

// GenToken 生成JWT
func GenToken(uid int64, name string) (string, error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		uid,
		name, // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "todo-app",                                 // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(MySecret)
}

// ParseToken 解析JWT
// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjE2NTAxNjkxNDcsIm5hbWUiOiLnjovnn7PlpLQiLCJleHAiOjE2NTAxNzI3NTgsImlzcyI6InRvZG8tYXBwIn0.tzTt3mp7vj15aE_HmKMgbwx6h0hc0Z_j2j31XtubUBs
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token, 得到的token对象就表示一个 jwt 对象
	token, err := jwt.ParseWithClaims(
		tokenString,
		&MyClaims{},
		func(token *jwt.Token) (i interface{}, err error) {
			return MySecret, nil // 加盐的字符串通过这个函数返回
		},
	)
	if err != nil {
		return nil, err
	}
	// token.Header
	// token.Claims
	// token.Signature
	// token.Valid
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
