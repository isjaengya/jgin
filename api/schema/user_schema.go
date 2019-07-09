package schema

import (
	"errors"
	"fmt"
	"github.com/francoispqt/gojay"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"tebu_go/api/lib"
	"tebu_go/api/service"
	"time"
)

type UserLoginSchema struct {
	Id int `json:"id" binding:"required,min=1,max=999999"` // 如果不是整形返回错误信息 您输入的不是数字
	Uid int `json:"uid" binding:"required"`
}

type UserInfoSchema struct {
	Uid     string    `form:"uid"`
}

func (u *UserLoginSchema) Bind (c *gin.Context) error {
	b := binding.Default(c.Request.Method, c.ContentType())
	err := c.ShouldBindWith(u, b)

	if err != nil {
		return err
	}
	return nil
}

type UserQuerySchema struct {
	Uids     []string    `form:"uid" binding:"required"`
}

func (u *UserQuerySchema) Bind (c *gin.Context) error {
	b := binding.Default(c.Request.Method, c.ContentType())
	err := c.ShouldBindWith(u, b)

	if err != nil {
		return err
	}
	return nil
}


type UnmarshalerJSONObject interface {
	UnmarshalJSONObject(*gojay.Decoder, string) error
	NKeys() int
}

type User struct {
	Id int `json:"id" from:"id"`
	CreateAt int32 `json:"create_at" from:"create_at"`
	UpdateAt int32 `json:"update_at" from:"update_at"`
	Uid string `json:"uid" from:"uid"`
	FamilyId int `json:"family_id" form:"family_id"`
}

func (u *User) MarshalJSONObject(enc *gojay.Encoder) {
	enc.IntKey("id", u.Id)
	enc.StringKey("uid", u.Uid)
	enc.IntKey("family_id", u.FamilyId)
	enc.Int32Key("create_at", u.CreateAt)
}

func (u *User) IsNil() bool {
	return u == nil
}

func (u *User) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
    switch key {
    case "id":
        return dec.Int(&u.Id)
    case "uid":
        return dec.String(&u.Uid)
    case "family_id":
        return dec.Int(&u.FamilyId)
	case "create_at":
		return dec.Int32(&u.CreateAt)

    }
    return nil
}
func (u *User) NKeys() int {
    return 4
}

func GetCacheInfoToUser(uid string) (u *User, err error) {
	user := &User{}
	redisClient := service.GetRedisClient()
	key := lib.UserRedisKey + uid
	s, err := redisClient.Get(key).Result()
	if err != nil {
		// 没有在redis中查询到数据, 扫db, 重新缓存
		fmt.Println(err.Error(), "not find data")
		user, err = FindUserByUid(uid)
		if err == nil {
			// 查询到
			go user.UpdateRedisCache()
			return user, err
		} else {
			// 没查到
			s := "没有查询到该用户, " + uid
			return user, errors.New(s)
		}

	}
	fmt.Println("s", s)
	err = gojay.UnmarshalJSONObject([]byte(s), user)
	//err = json.Unmarshal([]byte(s), user)
	if err != nil {
		fmt.Println(err.Error())
	}
	return user, err
}

func (u *User) SetUserJwtLast10(jwt string) {
	redisClient := service.GetRedisClient()
	key := lib.UserJwtRedisKey + u.Uid
	err := redisClient.Set(key, jwt, 0).Err()
	if err != nil {
		fmt.Println("set user jwt last 10 error, ", err.Error())
	}
}

func (u *User) UpdateRedisCache() {
	json := u.ToJson()
	redisClient := service.GetRedisClient()

	key := lib.UserRedisKey + u.Uid
	err := redisClient.Set(key, json, time.Second*3600).Err()
	if err != nil {
		fmt.Println("set user info error, ", err.Error())
	}
}

func (u *User) ToJson() (s string) {
	b, err := gojay.MarshalJSONObject(u)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}

func VerifyUserLogin(u UserLoginSchema) (user User, b bool) {
	id := u.Id
	uid := u.Uid

	mysqlClient := service.GetMysqlClient()
	err := mysqlClient.QueryRow("select * from user where id = ? and uid = ?", id, uid).Scan(&user.Id, &user.CreateAt, &user.UpdateAt, &user.Uid, &user.FamilyId)
	if err != nil{
		fmt.Println(err.Error(), "没有查询到user")
		return user, false
	}
	fmt.Println("查询到user", user)
	return user, true
}

func FindUserByUid(uid string) (u *User, err error) {
	user := &User{}
	mysqlClient := service.GetMysqlClient()

	err = mysqlClient.QueryRow("select * from user where uid = ?", uid).Scan(&user.Id, &user.CreateAt, &user.UpdateAt, &user.Uid, &user.FamilyId)
	if err != nil{
		fmt.Println(err.Error(), "没有查询到user")
		return user, err
	}
	fmt.Println("查询到user", user)
	return user, err
}

