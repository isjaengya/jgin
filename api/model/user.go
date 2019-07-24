package model

import (
	"errors"
	"fmt"
	"github.com/francoispqt/gojay"
	"jgin/api/lib"
	"jgin/api/schema"
	"jgin/api/service"
	"jgin/api/util"
	"log"
	"time"
)

type User struct {
	Id         int    `json:"id" from:"id"`
	CreateAt   int32  `json:"create_at" from:"create_at"`
	UpdateAt   int32  `json:"update_at" from:"update_at"`
	Uid        string `json:"uid" from:"uid"`
	FamilyId   int    `json:"family_id" form:"family_id"`
	SelectTime int64
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

type UnmarshalerJSONObject interface {
	UnmarshalJSONObject(*gojay.Decoder, string) error
	NKeys() int
}

func GetCacheInfoToUser(uid string) (u *User, err error) {
	user := &User{Uid: uid}
	redisClient := service.GetRedisClient()
	key := lib.UserRedisKey + uid
	s, err := redisClient.Get(key).Result()
	if err != nil {
		// 没有在redis中查询到数据, 扫db, 重新缓存
		ok := user.FindUserByUid()
		if ok {
			go user.UpdateRedisCache()
			return user, nil
		} else {
			s = "没有查询到该用户, " + uid
			return user, errors.New(s)
		}
	}
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

func (u *User) DeleteUserJwtLast10() {
	redisClient := service.GetRedisClient()
	key := lib.UserJwtRedisKey + u.Uid
	if err := redisClient.Del(key).Err(); err != nil {
		fmt.Println("delte user jwt last 10 err, ", err.Error())
	}

}

func (u *User) RedisCache() {
	json := u.ToJson()
	redisClient := service.GetRedisClient()

	key := lib.UserRedisKey + u.Uid
	err := redisClient.Set(key, json, time.Second*3600).Err()
	if err != nil {
		fmt.Println("set user info error, ", err.Error())
	}
}

func (u *User) UpdateRedisCache() {
	// 真正的锁逻辑
	i := 0
	for {
		if i >= lib.UserLockRetry {
			fmt.Println("超出重试上限，", i)
			break
		}
		lock := u.GetLock()
		fmt.Println(lock, "lock status")
		if lock {
			// 拿到锁
			u.RedisCache()
			// 删除锁
			u.DelLock()
			break
		} else {
			// 判断当前时间是否比锁时间大，如果小于等于锁时间不执行缓存, 更新用户信息时SelectTime为0，此时不管之前lock里面的值为多少
			if u.SelectTime == 0 {
				continue
			}
			userLockValue := u.GetLockValue()
			t := u.SelectTime
			if t <= userLockValue {
				fmt.Println("mysql time < user lock time, pass, lock time and current time is : ", userLockValue, t)
				break
			}
		}
		time.Sleep(1)
		i++
	}
}

func (u *User) GetLock() (b bool) {
	redisClient := service.GetRedisClient()
	key := lib.UserLockKye + u.Uid

	result, err := redisClient.SetNX(key, u.SelectTime, lib.UserLockExpire).Result()
	if err != nil {
		fmt.Println(err.Error(), "setnx error")
		return false
	}
	fmt.Println(result, "setnx status")
	return result
}

func (u *User) GetLockValue() (i int64) {
	redisClient := service.GetRedisClient()
	key := lib.UserLockKye + u.Uid

	result, err := redisClient.Get(key).Int64()
	if err != nil {
		fmt.Println("get user lock error, ", err.Error())
	}
	return result
}

func (u *User) DelLock() {
	redisClient := service.GetRedisClient()
	key := lib.UserLockKye + u.Uid

	if u.SelectTime == 0 {
		i, err := redisClient.Del(key).Result()
		fmt.Println(i, err, "删除 user lock, SelectTime = 0")
		return
	}
	userLockValue := u.GetLockValue()
	t := u.SelectTime
	if t >= userLockValue {
		i, err := redisClient.Del(key).Result()
		fmt.Println(i, err, "删除 user lock")
	}

}

func (u *User) ToJson() (s string) {
	b, err := gojay.MarshalJSONObject(u)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}

func VerifyUserLogin(u schema.UserLoginSchema) (user User, b bool) {
	id := u.Id
	uid := u.Uid

	mysqlClient := service.GetMysqlClient()
	err := mysqlClient.QueryRow("select * from user where id = ? and uid = ?", id, uid).Scan(&user.Id, &user.CreateAt, &user.UpdateAt, &user.Uid, &user.FamilyId)
	if err != nil {
		fmt.Println(err.Error(), "没有查询到user")
		return user, false
	}
	fmt.Println("查询到user", user)
	return user, true
}

func (u *User) FindUserByUid() (b bool) {
	mysqlClient := service.GetMysqlClient()

	err := mysqlClient.QueryRow("select id, COALESCE(created_at, 0), COALESCE(updated_at, 0), uid, family_id, unix_timestamp(now())  from user where uid = ?", u.Uid).Scan(&u.Id, &u.CreateAt, &u.UpdateAt, &u.Uid, &u.FamilyId, &u.SelectTime)
	if err != nil {
		fmt.Println(err.Error(), "没有查询到user")
		return false
	}
	fmt.Println("查询到user", u)
	return true
}

func (u User) GetUserTodayLoginF() (b bool) {
	// 今天登录奖励是否领取
	redisClient := service.GetRedisClient()
	// key: user_task_status_2019_7_19_uid
	key := util.GetUserTaskKey(u.Uid)
	result, err := redisClient.Get(key).Int()
	if err != nil {
		// key不存在，说明用户今天第一次登录，incr这个key同时增加用户连续登录天数
		_, _ = redisClient.Incr(key).Result()
		userLoginDaysKey := util.GetUserLoginDaysKey(u.Uid)
		_, _ = redisClient.Incr(userLoginDaysKey).Result()
		return true
	}
	if result <= 1 {
		// 未领取
		return true
	} else {
		return false
	}
}

func (u User) GetUserLoginDays() (i int) {
	redisClient := service.GetRedisClient()
	key := util.GetUserLoginDaysKey(u.Uid)
	loginDays, err := redisClient.Get(key).Int()
	if err != nil {
		fmt.Println(err.Error())
		return 1
	}
	return loginDays
}

func (u User) VerifyUserReceiveAward() (b bool) {
	redisClient := service.GetRedisClient()
	key := util.GetUserTaskKey(u.Uid)
	result, err := redisClient.Get(key).Int()
	if err != nil {
		// 说明用户今天还未登录
		return false
	} else {
		if result == 1 {
			// 可以领取
			return true
		}
		return false
	}

}
