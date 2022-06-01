package main

import (
	"errors"
	"log"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

type MyCustomClaims struct {
	UID string `json:"uid"`
	jwt.RegisteredClaims
}

var mySigningKey = []byte("夏天夏天悄悄过去")

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyCustomClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return mySigningKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// valid 校验认证信息.
func valid(authorization []string) bool {
	if len(authorization) < 1 {
		return false
	}

	// HTTP header
	// authorization: Bearer saasdasda.adsadasdasda.daqwfggod

	// TrimPrefix。去掉字符串首 `Bearer ` 字符
	token := strings.TrimPrefix(authorization[0], "Bearer ")
	// Perform the token validation here. For the sake of this example, the code
	// here forgoes any of the usual OAuth2 token validation and instead checks
	// for a token matching an arbitrary string.
	mc, err := ParseToken(token)
	log.Printf("uid: %s", mc.UID)
	if err != nil || mc.UID != "uid" {
		return false
	}
	return true
}
