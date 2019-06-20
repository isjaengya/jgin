package model

type User struct {
	Id int `json:"id" from:"id"`
	CreateAt int32 `json:"create_at" from:"create_at"`
	UpdateAt int32 `json:"update_at" from:"update_at"`
	Uid int32 `json:"uid" from:"uid"`
	FamilyId int32 `json:"family_id" form:"family_id"`
}
