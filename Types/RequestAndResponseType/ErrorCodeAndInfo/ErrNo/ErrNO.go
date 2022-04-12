package ErrNo

type ErrNo int

const (
	OK             ErrNo = 0
	ParamInvalid   ErrNo = 1 // 参数不合法
	UserHasExisted ErrNo = 2 // 该用户已存在
	UserNotExisted ErrNo = 4 // 该用户不存在
	WrongPassword  ErrNo = 5 // 密码错误
	LoginRequired  ErrNo = 6 // 用户未登录
	PermDenied     ErrNo = 7 // 没有操作权限
	MoneyNotEnough ErrNo = 8 // 金额不足

	UnknownError ErrNo = 255 // 未知错误
)
