/*
	jwt里面只存储uid, 和exp(一个纳秒级时间戳), exp用来生成jwt时, 保证同一个用户也生成不同的jwt, redis只存储jwt后10位, 减少空间占用
 */

package util

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
    time2 "time"
)

func CreateJwt(key string, m map[string] interface{}) string {
	NanoTime := time2.Now().UnixNano()
	m["exp"] = NanoTime
	token := jwt.New(jwt.SigningMethodHS256)
    claims := make(jwt.MapClaims)

    for index, val := range m {
        claims[index] = val
    }
    token.Claims = claims
    tokenString, _ := token.SignedString([]byte(key))
    return tokenString

}

func parseToken(tokenString string, key string) (jwt.MapClaims, bool){
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
        }
        return []byte(key), nil
    }); if err != nil{
    	return map[string]interface{}{"nil": "nil"}, false
	}
    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        return claims, true
    } else {
        fmt.Println(err)
        return map[string]interface{}{"nil": "nil"}, false
    }
}

func ParseTokenUid(tokenString string, key string) (uid string, b bool){
	claims, ok := parseToken(tokenString, key)
	if ok {
		uid := claims["uid"].(string)
		return uid, true
	} else {
		return "", false
	}
}