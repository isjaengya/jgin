package e

var msgFlags = map[int]string{
	INVALID_SIGNATURE: "签名出错",
	WRONG_SIGN_VERSION: "sign version 错误",
	MISSING_SIGN_HEADER: "缺少 ts 或 sign 或 sv",
	SUCCESS:                        "ok",
	ERROR:                          "fail",
	INVALID_PARAMS:                 "请求参数错误",
	ERROR_EXIST_TAG:                "已存在该标签名称",
	ERROR_NOT_EXIST_TAG:            "该标签不存在",
	ERROR_NOT_EXIST_ARTICLE:        "该文章不存在",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",
	SHOULD_ERROR:                   "请求参数不全",
	SELECT_ERROR:                   "查询出错,请重试",
	UNKONWN_ERROE:                  "未知错误",
	ADD_FAMILY_ERROR: "添加家庭出错，请重试",
	FAMILY_EXIST: "家庭已存在",
	CODE_EXIST: "邀请码生成重复，请重试",
	}

// GetMessage 通过code获取message
func GetMessage(code int) string {
	msg, ok := msgFlags[code]
	if ok {
		return msg
	}

	return msgFlags[ERROR]
}
