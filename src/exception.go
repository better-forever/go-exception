package exception

func (e *Exception) Error() string {
	return e.Msg
}

func New(code int, msg string) *Exception {
	return &Exception{
		Code: code,
		Msg:  msg,
	}
}

/*
* 0 表示没有错误
* 1 开头表示系统级错误
* 2 开头表示数据库错误
* 3 开头表示网络错误
* 4 开头表示业务错误
* 5 开头表示其他错误
 */
var (
	//系统错误
	EX_SYSTEM = New(1001, "系统异常")
	EX_BUSY   = New(1002, "服务器繁忙")
	//数据库错误
	EX_CREATE = New(2001, "新增失败")
	EX_UPDATE = New(2002, "更新失败")
	EX_QUERY  = New(2003, "查询失败")
	EX_DELETE = New(2004, "删除失败")
	//网络错误
	EX_NETWORK = New(3001, "网络错误")
	EX_TIMEOUT = New(3002, "请求超时")
	EX_REMOTE  = New(3003, "远程服务错误")
	//业务错误
	EX_NOT_FOUND      = New(4001, "资源不存在")
	EX_FORBIDDEN      = New(4002, "无权限访问")
	EX_USER_UNKNOWN   = New(4003, "用户不存在")
	EX_LOGIN          = New(4004, "账号或密码错误")
	EX_PARAM          = New(4005, "参数错误")
	EX_VALID          = New(4006, "验证失败")
	EX_CODE           = New(4007, "验证码错误")
	EX_TOKEN          = New(4008, "无效的令牌")
	EX_TASK           = New(4009, "无效的任务")
	EX_SIGNED         = New(4010, "签到失败")
	EX_CAPACITY       = New(4011, "数量超过上限")
	EX_SOURCE_EXISTED = New(4012, "资源已存在")
	EX_PHONE_EXISTED  = New(4013, "手机号已存在")
	EX_EMAIL_EXISTED  = New(4014, "邮箱已存在")
	EX_USER_EXISTED   = New(4015, "用户已存在")
	EX_REQUEST_LIMIT  = New(4016, "请求太频繁了")
	//其他错误

)
