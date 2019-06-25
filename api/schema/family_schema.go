package schema

type FamilySchema struct {
	FamilyName string `json:"family_name" binding:"required"`
	MemberName string `json:"member_name" binding:"required"`
	MemberSex string `json:"member_sex" binding:"required"`
	MemberAge string `json:"member_age" binding:"required"`
	MemberCity string `json:"member_city" binding:"required"`
	ChildName string `json:"child_name" binding:"required"`
	ChildSex string `json:"child_sex" binding:"required"`
	ChildAge string `json:"child_age" binding:"required"`
}

