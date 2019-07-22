package e

type Errmsg []interface{}

var (
	INVALID_SIGNATURE          = Errmsg{10005, "签名出错"}
	WRONG_SIGN_VERSION         = Errmsg{10006, "sign version 错误"}
	MISSING_SIGN_HEADER        = Errmsg{10007, "缺少 ts 或 sign 或 sv"}
	SUCCESS                    = Errmsg{200, "ok"}
	ERROR                      = Errmsg{500, "fail"}
	SHOULD_ERROR               = Errmsg{300001, "请求参数不正确"}
	JWT_PARSE_ERROE            = Errmsg{10008, "JWT字符串解析失败"}
	JWT_INVALID                = Errmsg{10012, "JWT已失效"}
	PASSWORD_OR_USERNAME_ERROR = Errmsg{10011, "账号或密码不正确"}
	MIDDLEWARE_GET_USER_ERROR  = Errmsg{10012, "middleware获取用户信息失败"}
	USER_NOT_LOGIN_TODAY       = Errmsg{10013, "该用户今天还未登录"}
)

var validateMsg = map[string]string{
	"required": "字段是必须的",
	"max":      "最大值或长度超出",
	"min":      "最小值或长度超出",
}

func GetValidateMessage(s string) string {
	msg, ok := validateMsg[s]
	if ok {
		return msg
	}
	return ""
}
