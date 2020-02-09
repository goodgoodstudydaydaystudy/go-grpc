package status

const (
	Success = 0 // 成功

	// -1 ~ -99 Internal System error
	ErrSys = -2 // 服务器错误, 默认错误
	ErrDB  = -3 // 数据库错误
	ErrDBTypeNotSupport = -4 // 数据库类型不支持

	// -1000 ~ -1999 Account
	ErrAccountExists = -1000 // 账号已存在
	ErrAccountNotExists = -1001 //账号不存在
	ErrPasswordError = -1002 //密码错误

	// -2000 ~ -2999 Pay
	ErrBalanceNotEnough = -2000 // 余额不足
	ErrRechargeFailed = -2001	// 充值失败
	ErrGetUserBalanceFailed = -2002 //查询余额失败

	// -3000 ~ -3999 token
	ErrNotLogin = -3001 // 未登陆

)

var CodeMessageMap = map[int]string{
	Success:           		  	"成功",
	ErrAccountExists:    		"账号已存在",
	ErrAccountNotExists:		"账号不存在",
	ErrPasswordError:  			"密码错误",
	ErrBalanceNotEnough:     	"余额不足",
	ErrRechargeFailed:			"充值失败",
	ErrGetUserBalanceFailed:   	"查询余额失败",
	ErrDB:               		"数据库错误",
	ErrDBTypeNotSupport:		"数据库类型不支持",
	ErrNotLogin:				"未登陆",
}

// MessageFromCode get message associated with the code
func MessageFromCode(code int) string {
	if m, ok := CodeMessageMap[code]; ok {
		return m
	}

	return ""
}
