package lib

import "time"

const (
	JwtKey = "pXMiBAr331"

	UserLockKye     = "user_lock_"
	UserLockExpire  = 20 * time.Second
	UserLockRetry   = 5
	UserRedisKey    = "user_info_"
	UserJwtRedisKey = "user_jwt_"

	UserLoginDays  = "user_login_days_"
	UserTaskStatus = "user_task_status_"
)
