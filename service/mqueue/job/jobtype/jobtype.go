package jobtype

// DeferCloseOrder 延迟关闭订单
const DeferCloseOrder = "defer:order:close"

// DeferCloseOrderPayload 延迟关闭订单
type DeferCloseOrderPayload struct {
	OrderNo string
}
