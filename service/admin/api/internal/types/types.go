// Code generated by goctl. DO NOT EDIT.
package types

type Admin struct {
	AdminUserId   int64  `json:"adminUserId"`
	LoginUserName string `json:"loginUserName"`
	LoginPassword string `json:"loginPassword"`
	NickName      string `json:"nickName"`
	Locked        int64  `json:"locked"`
}

type AdminLoginRequest struct {
	UserName    string `json:"userName"`
	PasswordMd5 string `json:"passwordMd5"`
}

type AdminLoginResponse struct {
	Token string `json:"token"`
}

type UpdateAdminNameRequest struct {
	Token         string `header:"token"`
	LoginUserName string `json:"loginUserName"`
	NickName      string `json:"nickName"`
}

type UpdateAdminPwdRequest struct {
	Token            string `header:"token"`
	OriginalPassword string `json:"originalPassword"`
	NewPassword      string `json:"newPassword"`
}

type GetAdminProfileRequest struct {
	Token string `header:"token"`
}

type GetAdminProfileResponse struct {
	Admin Admin `json:"admin"`
}

type GetUserListRequest struct {
	PageNumber int64 `json:"pageNumber"` // 页码
	PageSize   int64 `json:"pageSize"`   // 每页大小
}

type GetListResponse struct {
	List       interface{} `json:"list"`
	TotalCount int64       `json:"totalCount"`
	TotalPage  int64       `json:"totalPage"`
	CurrPage   int64       `json:"currPage"`
	PageSize   int64       `json:"pageSize"`
}

type LockUserRequest struct {
	Ids        []int64 `json:"ids"`
	LockStatus int64   `path:"lockStatus"`
}

type LogoutRequest struct {
	Token string `header:"token"`
}

type AddGoodsInfoRequest struct {
	GoodsName          string `json:"goodsName"`
	GoodsIntro         string `json:"goodsIntro"`
	GoodsCategoryId    int64  `json:"goodsCategoryId"`
	GoodsCoverImg      string `json:"goodsCoverImg"`
	GoodsCarousel      string `json:"goodsCarousel"`
	GoodsDetailContent string `json:"goodsDetailContent"`
	OriginalPrice      int64  `json:"originalPrice"`
	SellingPrice       int64  `json:"sellingPrice"`
	StockNum           int64  `json:"stockNum"`
	Tag                string `json:"tag"`
	GoodsSellStatus    int64  `json:"goodsSellStatus"`
}

type UpdateGoodsInfoRequest struct {
	GoodsId            int64  `json:"goodsId"`
	GoodsName          string `json:"goodsName, optional"`
	GoodsIntro         string `json:"goodsIntro, optional"`
	GoodsCategoryId    int64  `json:"goodsCategoryId, optional"`
	GoodsCoverImg      string `json:"goodsCoverImg, optional"`
	GoodsCarousel      string `json:"goodsCarousel, optional"`
	GoodsDetailContent string `json:"goodsDetailContent, optional"`
	OriginalPrice      int64  `json:"originalPrice, optional"`
	SellingPrice       int64  `json:"sellingPrice, optional"`
	StockNum           int64  `json:"stockNum, optional"`
	Tag                string `json:"tag, optional"`
	GoodsSellStatus    int64  `json:"goodsSellStatus, optional"`
}

type AlterGoodsStatusRequest struct {
	Ids    []int64 `json:"ids"`
	Status int64   `path:"status"`
}

type DeleteGoodsInfoRequest struct {
	GoodsId int64 `json:"goodsId"`
}

type GetGoodsInfoRequest struct {
	Id int64 `path:"id"`
}

type GoodsInfo struct {
	GoodsId            int64  `json:"goodsId"`
	GoodsName          string `json:"goodsName"`
	GoodsIntro         string `json:"goodsIntro"`
	GoodsCategoryId    int64  `json:"goodsCategoryId"`
	GoodsCoverImg      string `json:"goodsCoverImg"`
	GoodsCarousel      string `json:"goodsCarousel"`
	GoodsDetailContent string `json:"goodsDetailContent"`
	OriginalPrice      int64  `json:"originalPrice"`
	SellingPrice       int64  `json:"sellingPrice"`
	StockNum           int64  `json:"stockNum"`
	Tag                string `json:"tag"`
	GoodsSellStatus    int64  `json:"goodsSellStatus"`
	CreateUser         int64  `json:"createUser"`
	CreateTime         string `json:"createTime"`
	UpdateUser         int64  `json:"updateUser"`
	UpdateTime         string `json:"updateTime"`
}

type GoodsCategory struct {
	CategoryId    int64  `json:"categoryId"`
	CategoryLevel int64  `json:"categoryLevel"`
	ParentId      int64  `json:"parentId"`
	CategoryName  string `json:"categoryName"`
	CategoryRank  int64  `json:"categoryRank"`
	IsDeleted     int64  `json:"isDeleted"`
	CreateTime    string `json:"createTime"`
	UpdateTime    string `json:"updateTime"`
}

type GetGoodsInfoResponse struct {
	Goods          GoodsInfo     `json:"goods"`
	ThirdCategory  GoodsCategory `json:"thirdCategory"`
	SecondCategory GoodsCategory `json:"secondCategory"`
	FirstCategory  GoodsCategory `json:"firstCategory"`
}

type GetGoodsListRequest struct {
	PageNumber      int64  `json:"pageNumber"`
	PageSize        int64  `json:"pageSize"`
	GoodsName       string `form:"goodsName, optional"`
	GoodsSellStatus string `form:"goodsSellStatus, optional"`
}

type Response struct {
	ResultCode int         `json:"resultCode"`
	Msg        string      `json:"message"`
	Data       interface{} `json:"data"`
}