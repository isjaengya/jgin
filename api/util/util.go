package util

import (
	"jgin/api/lib"
	"math/rand"
	"strings"
	"time"
)

func RandString(len int) string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

func CountVerify(i int32) (b bool) {
	if i >= 1 {
		return true
	} else {
		return false
	}
}

func SplitUid(uidS string) (uid []string) {
	return strings.Split(uidS, ",")[:10]
}

func GetTodayStr() (s string) {
	timeStr := time.Now().Format("2006-01-02")
	return timeStr
}

func GetUserTaskKey(uid string) (s string) {
	// key: user_task_status_2019_7_19_uid
	prefix := GetTodayStr()
	return lib.UserTaskStatus + prefix + "_" + uid
}

func GetUserLoginDaysKey(uid string) (s string) {
	return lib.UserLoginDays + uid
}
