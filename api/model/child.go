package model

type Child struct {
	Id int `json:"id" from:"id"`
	CreateAt int32 `json:"create_at" from:"create_at"`
	UpdateAt int32 `json:"update_at" from:"update_at"`
	ChildName string `json:"child_name" from:"child_name"`
	ChildSex string `json:"child_sex" from:"child_sex"`
	ChildAge int `json:"child_age" json:"child_age"`
}
