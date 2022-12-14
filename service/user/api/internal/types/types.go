// Code generated by goctl. DO NOT EDIT.
package types

type LoginRequest struct {
	LoginName   string `json:"loginName"`
	PasswordMd5 string `json:"passwordMd5"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterRequest struct {
	LoginName string `json:"loginName"`
	Password  string `json:"password"`
}

type UpdateUserInfoRequest struct {
	Token         string `header:"token"`
	NickName      string `json:"nickName"`
	PasswordMd5   string `json:"passwordMd5"`
	IntroduceSign string `json:"introduceSign"`
}

type GetUserInfoRequest struct {
	Token string `header:"token"`
}

type GetUserInfoResponse struct {
	NickName      string `json:"nickName"`
	LoginName     string `json:"loginName"`
	IntroduceSign string `json:"introduceSign"`
}

type LogoutRequest struct {
	Token string `header:"token"`
}

type GetGoodsDetailRequest struct {
	Id int64 `path:"id"`
}

type GetGoodsDetailResponse struct {
	GoodsId            int64    `json:"goodsId"`
	GoodsName          string   `json:"goodsName"`
	GoodsIntro         string   `json:"goodsIntro"`
	GoodsCoverImg      string   `json:"goodsCoverImg"`
	SellingPrice       int64    `json:"sellingPrice"`
	GoodsDetailContent string   `json:"goodsDetailContent"`
	OriginalPrice      int64    `json:"originalPrice"`
	Tag                string   `json:"tag" form:"tag"`
	GoodsCarouselList  []string `json:"goodsCarouselList"`
}

type Response struct {
	ResultCode int         `json:"resultCode"`
	Msg        string      `json:"message"`
	Data       interface{} `json:"data"`
}

type SaveAddressRequest struct {
	UserName      string `json:"userName"`
	UserPhone     string `json:"userPhone"`
	DefaultFlag   int64  `json:"defaultFlag"` // 0-不是 1-是
	ProvinceName  string `json:"provinceName"`
	CityName      string `json:"cityName"`
	RegionName    string `json:"regionName"`
	DetailAddress string `json:"detailAddress"`
	Token         string `header:"token"`
}

type UpdateAddressRequest struct {
	AddressId     int64  `json:"addressId"`
	UserId        int64  `json:"userId, optional"`
	UserName      string `json:"userName"`
	UserPhone     string `json:"userPhone"`
	DefaultFlag   int64  `json:"defaultFlag"` // 0-不是 1-是
	ProvinceName  string `json:"provinceName"`
	CityName      string `json:"cityName"`
	RegionName    string `json:"regionName"`
	DetailAddress string `json:"detailAddress"`
	Token         string `header:"token"`
}

type DeleteAddressRequest struct {
	AddressId int64  `path:"addressId"`
	Token     string `header:"token"`
}

type GetMyAddressRequest struct {
	Token string `header:"token"`
}

type Address struct {
	AddressId     int64  `json:"addressId"`
	UserId        int64  `json:"userId"`
	UserName      string `json:"userName"`
	UserPhone     string `json:"userPhone"`
	DefaultFlag   int64  `json:"defaultFlag"`
	ProvinceName  string `json:"provinceN"`
	CityName      string `json:"cityName"`
	RegionName    string `json:"regionName"`
	DetailAddress string `json:"detailAddress"`
	IsDeleted     int64  `json:"isDeleted"`
	CreateTime    string `json:"createTime"`
	UpdateTime    string `json:"updateTime"`
}

type GetMyAddressResponse struct {
	AddressList []Address `json:"addressList"`
}

type GetAddressByIdRequest struct {
	AddressId int64  `path:"addressId"`
	Token     string `header:"token"`
}

type GetAddressByIdResponse struct {
	Address Address `json:"address"`
}

type GetDefaultAddressRequest struct {
	Token string `header:"token"`
}

type GetDefaultAddressResponse struct {
	Address Address `json:"address"`
}

type AddCartItemRequest struct {
	GoodsId    int64  `json:"goodsId"`
	GoodsCount int64  `json:"goodsCount"`
	Token      string `header:"token"`
}

type UpdateCartItemRequest struct {
	CartItemId int64  `json:"cartItemId"`
	GoodsCount int64  `json:"goodsCount"`
	Token      string `header:"token"`
}

type DeleteCartItemRequest struct {
	Id    int64  `path:"id"`
	Token string `header:"token"`
}

type GetCartListRequest struct {
	Token string `header:"token"`
}

type CartItem struct {
	CartItemId    int64  `json:"cartItemId"`
	GoodsId       int64  `json:"goodsId"`
	GoodsCount    int64  `json:"goodsCount"`
	GoodsName     string `json:"goodsName"`
	GoodsCoverImg string `json:"goodsCoverImg"`
	SellingPrice  int64  `json:"sellingPrice"`
}

type GetCartListResponse struct {
	CartItemList []CartItem `json:"cartItemList"`
}

type GetCartItemsRequest struct {
	CartItemIds string `form:"cartItemIds"`
	Token       string `header:"token"`
}

type AddOrderRequest struct {
	CartItemIds []int64 `json:"cartItemIds"`
	AddressId   int64   `json:"addressId"`
	Token       string  `header:"token"`
}

type PaySuccessRequest struct {
	OrderNo string `form:"orderNo"`
	PayType int64  `form:"payType"`
}

type OrderRequest struct {
	OrderNo string `path:"orderNo"`
	Token   string `header:"token"`
}

type GetOrderListRequest struct {
	Status     string `form:"status, optional"`
	PageNumber int64  `form:"pageNumber"`
	Token      string `header:"token"`
}

type OrderItem struct {
	GoodsId       int64  `json:"goodsId"`
	GoodsName     string `json:"goodsName"`
	GoodsCount    int64  `json:"goodsCount"`
	GoodsCoverImg string `json:"goodsCoverImg"`
	SellingPrice  int64  `json:"sellingPrice"`
}

type GetOrderDetailResponse struct {
	OrderNo           string      `json:"orderNo"`
	TotalPrice        int64       `json:"totalPrice"`
	PayStatus         int64       `json:"payStatus"`
	PayType           int64       `json:"payType"`
	PayTypeString     string      `json:"payTypeString"`
	PayTime           string      `json:"payTime"`
	OrderStatus       int64       `json:"orderStatus"`
	OrderStatusString string      `json:"orderStatusString"`
	CreateTime        string      `json:"createTime"`
	OrderItems        []OrderItem `json:"orderItems"`
}

type ListItem struct {
	OrderId           int64       `json:"orderId"`
	OrderNo           string      `json:"orderNo"`
	TotalPrice        int64       `json:"totalPrice"`
	PayType           int64       `json:"payType"`
	PayTypeString     string      `json:"payTypeString"`
	OrderStatus       int64       `json:"orderStatus"`
	OrderStatusString string      `json:"orderStatusString"`
	CreateTime        string      `json:"createTime"`
	OrderItems        []OrderItem `json:"orderItems"`
}

type GetListResponse struct {
	List       []ListItem `json:"list"`
	TotalCount int64      `json:"totalCount"`
	TotalPage  int64      `json:"totalPage"`
	CurrPage   int64      `json:"currPage"`
	PageSize   int64      `json:"pageSize"`
}
