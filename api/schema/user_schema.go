package schema

type UserLoginSchema struct {
	Id int `json:"id" binding:"required"`
	Uid int `json:"uid" binding:"required"`
}

type UserInfoSchema struct {
	Uid     string    `form:"uid"`
}

