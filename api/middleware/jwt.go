/*
	jwt里面只存储uid, 和exp(一个纳秒级时间戳), exp用来生成jwt时, 保证同一个用户也生成不同的jwt, redis只存储jwt后10位, 减少空间占用
*/

package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"jgin/api/lib"
	"jgin/api/service"
	time2 "time"
)

func CreateJwt(m map[string]interface{}) string {
	key := lib.JwtKey
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

func parseToken(tokenString string) (jwt.MapClaims, bool) {
	key := lib.JwtKey
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})
	if err != nil {
		return map[string]interface{}{"nil": "nil"}, false
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		return map[string]interface{}{"nil": "nil"}, false
	}
}

func ParseTokenUid(tokenString string) (uid string, b bool) {
	// TODO 解析jwt里面的信息在这里解析
	claims, ok := parseToken(tokenString)
	if ok {
		uid := claims["uid"].(string)
		return uid, true
	} else {
		return "", false
	}
}

func GetUserJwtLast10(uid string) (s string) {
	redisClient := service.GetRedisClient()
	key := lib.UserJwtRedisKey + uid
	s, err := redisClient.Get(key).Result()
	if err != nil {
		fmt.Println("set user jwt last 10 error, ", err.Error())
		return ""
	}
	return s
}

func GinGetJwt(c *gin.Context, uid string) (s string) {
	// TODO 在jwt里面放置什么信息在这里添加
	m := map[string]interface{}{"uid": uid}
	_jwt := CreateJwt(m)
	c.Header("Authorization", _jwt)
	fmt.Println("jwt: --> ", _jwt)
	return _jwt[len(_jwt)-10:]
}
