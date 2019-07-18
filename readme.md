### 封装gin框架, 更符合一个pythoner的开发习惯

#### 启动测试
go run main.go 启动gin框架
go run machinery_master.go 启动machinery任务队列
go run machinery_sender.go 模拟一个异步任务的发送，发送完毕之后在machinery_master界面能看到结果输出

#### curl测试
```bash
登录
curl -i -X POST \
   -H "Content-Type:application/json" \
   -H "sv:1" \
   -H "sign:kWzyW23DOnMpGXz9Iqj2fWkaenYz0Qw7JiJrLqA5gZ2DnVGlhSWfoOvZqsa6opoc2m3DwJmfWhuwQRDQLTVY0QHCKR9JoycLljBH" \
   -H "ts:23452" \
   -d \
'{
  "id": 1,
  "uid": 99475266
}' \
 'http://127.0.0.1:8000/user/login'

查看用户信息
curl -i -X GET \
   -H "sv:1" \
   -H "sign:kWzyW23DOnMpGXz9Iqj2fWkaenYz0Qw7JiJrLqA5gZ2DnVGlhSWfoOvZqsa6opoc2m3DwJmfWhuwQRDQLTVY0QHCKR9JoycLljBH" \
   -H "ts:23452" \
   -H "Authorization:这里是登录返回的jwt，前面没有Bearer！！！！！！！！！！！" \
   # 长这个样子，记住没有前面的
   -H "Authorization:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjM0Mzc3OTYxMDY5MjYwMDAsInVpZCI6Ijk5NDc1MjY2In0.OtGO53UGVkSWGCQtbScsOKXFLhRMwhBoOlEZPG7qOG0" \
 'http://127.0.0.1:8000/v1/user?uid=99475266&uid=99475268&uid=99475267'
 
退出
curl -i -X GET \
   -H "Authorization:这里是登录返回的jwt" \
   -H "sign:kWzyW23DOnMpGXz9Iqj2fWkaenYz0Qw7JiJrLqA5gZ2DnVGlhSWfoOvZqsa6opoc2m3DwJmfWhuwQRDQLTVY0QHCKR9JoycLljBH" \
   -H "sv:1" \
   -H "ts:23452" \
   -H "sign1:12" \
 'http://127.0.0.1:8000/v1/user/logout'

```

#### 目录结构
```
 - common     一些通用配置，返回正常信息、错误信息
 - config     配置文件
 - handler    逻辑处理
 - lib        常量、状态码、middleware基础函数
 - middlewre  权限验证中间件
 - model      用来存放表结构信息,暂时无用
 - route      路由
 - schema     用来解析请求参数，代替model层
 - service    db初始化
 - util       其他工具函数
 - tasks      异步任务
```

##### 异步任务(machinery)
```go
register_func 里面注册异步任务的map，每增加一个新的异步任务在里面写对应关系
send_func 里面构造一个signature，在业务层面只调用这里面的方法即可，更加简洁
```

##### 参数解析
```go
func (v *defaultValidator) lazyInit() {
	v.once.Do(func() {
		v.validate = validator.New()
		v.validate.SetTagName("binding")
		_ = v.validate.RegisterValidation("family_name-uniq", schema.ValidateUniqFamilyName)
		// 在里添加自定义的验证函数
	})
}
--------------------------------------------------------------------------------
返回自定义错误信息
上面添加了一个名为`family_name-uniq`的验证关键字, 在文件/lib/e/message.go validateMsg下面增加一行
var validateMsg = map[string]string{
	"required": "字段是必须的",
	"max": "最大值或长度超出",
	"min": "最小值或长度超出",
	"family_name-uniq": "家庭名字重复",
}
```
##### 参数验证
```go
结构体:
type UserLoginSchema struct {
	Id int `json:"id" binding:"required,min=1,max=999999"` // 如果不是整形返回错误信息 您输入的不是数字
	Uid int `json:"uid" binding:"required"`
}
---------------------------------------------------
定义bind方法用来验证参数是否合法:
func (u *UserLoginSchema) Bind (c *gin.Context) error {
	b := binding.Default(c.Request.Method, c.ContentType())
	err := c.ShouldBindWith(u, b)

	if err != nil {
		return err
	}
	return nil
}
```

#### 其他
作为用户中心服务需要大量序列化与反序列化, 所以需要一个高效的json解析库, 对比之下选择了gojay。具体用法可以查看gojay的官网文档。
```go
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
```
##### 以下为用到的mysql信息
```go
mysql:
  user: "root"
  password: "root"
  host: "127.0.0.1"
  port: 3306
  database: yinyu_gin

CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `created_at` int(11) DEFAULT NULL,
  `updated_at` int(11) DEFAULT NULL,
  `uid` varchar(256) DEFAULT NULL COMMENT '用户uid',
  `family_id` int(11) DEFAULT NULL COMMENT '家庭id',
  PRIMARY KEY (`id`),
  KEY `ix_user_created_at` (`created_at`),
  KEY `ix_user_uid` (`uid`),
  KEY `ix_user_updated_at` (`updated_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `yinyu_gin`.`user`(`id`, `created_at`, `updated_at`, `uid`, `family_id`) VALUES (1, 1560258391, 1560258639, '99475266', 2);
INSERT INTO `yinyu_gin`.`user`(`id`, `created_at`, `updated_at`, `uid`, `family_id`) VALUES (2, 1560258391, 1560258639, '99475267', 2);
INSERT INTO `yinyu_gin`.`user`(`id`, `created_at`, `updated_at`, `uid`, `family_id`) VALUES (3, 1560258391, 1560258639, '99475268', 2);
```