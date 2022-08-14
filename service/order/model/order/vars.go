package order

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var ErrNotFound = sqlx.ErrNotFound

// GetOrderStatusStr 根据Status代码获取字符串描述
func GetOrderStatusStr(status int64) string {
	switch status {
	case 0:
		return "待支付"
	case 1:
		return "已支付"
	case 2:
		return "配货完成"
	case 3:
		return "出库成功"
	case 4:
		return "交易成功"
	case -1:
		return "手动关闭"
	case -2:
		return "超时关闭"
	case -3:
		return "商家关闭"
	default:
		return "error"
	}
}

// GetPayTypeStr 根据payType代码获取字符串描述
func GetPayTypeStr(payType int64) string {
	switch payType {
	case 0:
		return "无"
	case 1:
		return "支付宝支付"
	case 2:
		return "微信支付"
	default:
		return "error"
	}
}
