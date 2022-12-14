package define

import (
	"fmt"
	"gorm.io/gorm"
	"yifan/app/db"
)

const (
	YfUserTypePersonalSeller    = 1 //个人卖家
	YfUserTypePersonalBuyer     = 2 //个人买家
	YfUserTypeMerchantSeller    = 3 //商户卖家
	YfUserTypeMerchantBuyer     = 4 //商户买家
	YfUserTypeStockOutAssistant = 5 //仓库出库员
	YfUserTypeStockInAssistant  = 6 //仓库验收员

	YfUserRoleAdmin    = 1 //管理员
	YfUserRoleDevelop  = 2 //开发人员
	YfUserRoleOperator = 3 //运营人员

	YfUserPermissionLogin  = 1 //登录权限
	YfUserPermissionQuery  = 2 //查询权限
	YfUserPermissionAdd    = 3 //增加权限
	YfUserPermissionModify = 4 //修改权限
	YfUserPermissionDelete = 5 //删除权限

	YfUserClassFirst  = 1 //用户等级1
	YfUserClassSecond = 2 //用户等级2
	YfUserClassThird  = 3 //用户等级3

	YfConsignmentTypeSelf         = 1 //自行邮寄
	YfConsignmentType9youPlatform = 2 //上面收货

	YfConsignmentOrderNeedToPay       = 1 //待支付：支付货物入库检验费用
	YfConsignmentOrderNeedToDeliver   = 2 //待发货：等待用户填写快递单号
	YfConsignmentOrderNeedToArrive    = 3 //待收货：等待快递到达仓库
	YfConsignmentOrderNeedToStorage   = 4 //待揽收：等待仓库揽收
	YfConsignmentOrderNeedToCheck     = 5 //待检验：等到仓库工作人员开箱检验货物
	YfConsignmentOrderNeedToVerify    = 6 //待审核：等单云仓运营审核检验货物
	YfConsignmentOrderNeedToGrounding = 7 //待上架：等待用户确认检验货物后，标记商品款式和价格后上架
	YfConsignmentOrderFinished        = 8 //完成：寄售单完结

	YfGoodsPkgStatus       = 0
	YfGoodsPkgStatusNewOld = 1 //拆盒未拆袋
	YfGoodsPkgStatusNewNew = 2 //全新（未拆盒未拆袋）
	YfGoodsPkgStatusOldOld = 3 //拆盒拆袋

	//YfSoldStatus        = 0
	YfSoldStatusPreSell = 1 //预售
	YfSoldStatusExist   = 2 //现货

	//奖品状态
	YfPrizeStatusSoldOut    = 1 //奖品售罄
	YfPrizeStatusNotSoldOut = 2 //奖品未售罄
	YfPrizeStatusDelete     = 6

	//箱子状态
	YfBoxStatusPrizeLeft    = 1 //箱子上架有奖品 (展示)
	YfBoxStatusNotOnBoard   = 2 //箱子未上架 (不展示)
	YfBoxStatusPrizeNotLeft = 3 //箱子上架无奖品 (展示)
	YfBoxStatusDelete       = 6 //删除

	//蕃状态
	YfFanStatusOnBoardByMan     = 1 //手动蕃上架 (展示)
	YfFanStatusOnBoardAuto      = 2 //自动蕃上架(到时间) (展示)
	YfFanStatusNotOnBoard       = 3 //蕃未上架 (展示)
	YfFanStatusNotOffBoardByMan = 4 //手动蕃下架 (展示)
	YfFanStatusNotOffBoardAuto  = 5 //自动蕃下架(到时间) (展示)
	YfFanStatusDelete           = 6 //删除

	//广告位状态
	YfAdvartanceOnBoardByMan  = 1 //广告位手动上架 (展示)
	YfAdvartanceOnBoardAuto   = 2 //广告位自动上架 (展示)
	YfAdvartanceOffBoardByMan = 3 //广告位手动下架 (不展示)
	YfAdvartanceOffBoardAuto  = 4 //广告位自动下架 (不展示)

	PrizeIndexNameFirst  = "First"
	PrizeIndexNameLast   = "Last"
	PrizeIndexNameGlobal = "全局"

	OrderStatusToBePickUp  = 1 // 1.待提货
	OrderStatusToDelivered = 2 // 2.待发货
	OrderStatusToReceive   = 3 // 3.等收货
	OrderStatusFinish      = 4 // 4.已完成
	OrderStatusAll         = 5 //全部

	PrizeIndexNameA = "A"
	PrizeIndexNameB = "B"
	PrizeIndexNameC = "C"
	PrizeIndexNameD = "D"
	PrizeIndexNameE = "E"
	PrizeIndexNameF = "F"
	PrizeIndexNameG = "G"
	PrizeIndexNameH = "H"
	PrizeIndexNameI = "I"
	PrizeIndexNameJ = "J"
	PrizeIndexNameK = "K"
	PrizeIndexNameL = "L"
	PrizeIndexNameM = "M"
	PrizeIndexNameN = "N"
	PrizeIndexNameO = "O"
	PrizeIndexNameP = "P"
	PrizeIndexNameQ = "Q"
	PrizeIndexNameR = "R"
	PrizeIndexNameS = "S"
	PrizeIndexNameT = "T"
	PrizeIndexNameU = "U"
	PrizeIndexNameV = "V"
	PrizeIndexNameW = "W"
	PrizeIndexNameX = "X"
	PrizeIndexNameY = "Y"
	PrizeIndexNameZ = "Z"
)

func RedisTarget(fanId, boxId uint) string {
	return fmt.Sprintf("fanId%d-boxId%d-target", fanId, boxId)
}
func RedisLeft(fanId, boxId uint) string {
	return fmt.Sprintf("fanId%d-boxId%d-left", fanId, boxId)
}
func RedisSure(fanId, boxId uint) string {
	return fmt.Sprintf("fanId%d-boxId%d-sure", fanId, boxId)
}
func RedisUser(fanId, boxId uint) string {
	return fmt.Sprintf("fanId%d-boxId%d-user", fanId, boxId)
}

func RedisFirstPrize(fanId, boxId uint) string {
	return fmt.Sprintf("fanId%d-boxId%d-first-prize", fanId, boxId)
}
func RedisLastPrize(fanId, boxId uint) string {
	return fmt.Sprintf("fanId%d-boxId%d-last-prize", fanId, boxId)
}
func RedisGlobalPrize(fanId, boxId uint) string {
	return fmt.Sprintf("fanId%d-boxId%d-global-prize", fanId, boxId)
}

func RedisSpecealRecordPosition(fanId, boxId uint) string {
	return fmt.Sprintf("fanId%d-boxId%d-specialrecord-position", fanId, boxId)
}

func RedisFirstPrizePosition(fanId, boxId uint) string {
	return fmt.Sprintf("fanId%d-boxId%d-first-position", fanId, boxId)
}
func RedisLastPrizePosition(fanId, boxId uint) string {
	return fmt.Sprintf("fanId%d-boxId%d-last-position", fanId, boxId)
}
func RedisGlobalPrizePosition(fanId, boxId uint) string {
	return fmt.Sprintf("fanId%d-boxId%d-global-position", fanId, boxId)
}
func IsHasSameEle(nums []int) bool {
	if len(nums) >= 1 {
		if nums[0] == -1 {
			return false
		}
	}
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
		if m[nums[i]] > 1 {
			return true
		}
	}
	return false
}
func DeleteSlice(a []db.Prize, elem uint) []db.Prize {
	for i := 0; i < len(a); i++ {
		if a[i].ID == elem {
			a = append(a[:i], a[i+1:]...)
			i--
		}
	}
	return a
}

type User struct {
	UserId         uint   `json:"userId,omitempty"`
	Time           int64  `json:"time,omitempty"`
	FanId          uint   `json:"fanId,omitempty"`
	FanName        string `json:"fanName,omitempty"`
	BoxId          uint   `json:"boxId,omitempty"`
	BoxName        string `json:"boxName,omitempty"`
	PrizeIndexName string `json:"prizeIndexName,omitempty"`
	PrizeName      string `json:"prizeName,omitempty"`
	Position       int    `json:"position,omitempty"`
}
type PrizeIdIndexName struct {
	PrizeId           uint
	PrizeName         string
	PrizeIndex        int32
	FanId             uint
	FanName           string
	GoodID            uint
	GoodName          string
	PrizeIndexName    string
	Remark            string
	Price             float64
	SoldStatus        int
	TimeForSoldStatus string
	PkgStatus         int
	Pic               string
	IpID              uint
	IpName            string
	SeriesID          uint
	SeriesName        string
}
type AccessTokenErrorResponse struct {
	ErrMsg  string `json:"err_msg"`
	ErrCode string `json:"err_code"`
}
type SnsOauth2 struct {
	Sessionkey string `json:"session_key"`
	Openid     string `json:"openid"`
}

type Ip struct {
	IpName string
}
type Series struct {
	IpName     string
	SeriesName string
}
type Goods struct {
	Ip        string
	Series    string
	GoodsName string
	PkgStatus string
	PreStore  string
	Integral  int
	Pic       string
	Price     float64
	CreatTime string
}
type GoodsNormal struct {
	GoodId         int
	PrizeIndexName string
	PrizeNum       int
	Range          string
	Ip             string
	Series         string
	GoodName       string
	Pic            string
	Status         string
}
type GoodsPosition struct {
	PrizeName string
	PrizeNum  int
	Range     string
	Remark    string
}

func DealWithOneIp(DB *gorm.DB, ipName string) uint {
	var tmpIp db.Ip
	tmpIp.Name = ipName
	DB.Create(&tmpIp)
	return tmpIp.ID
}
func DealWithOneSeries(DB *gorm.DB, ipName, serName string) uint {
	ser := db.Series{}
	ser.Name = serName
	ser.IpName = ipName
	DB.Create(&ser)
	return ser.ID
}
func DealWithOneGood(og Goods) {

}
func DealWithOneGoodNormal(nor GoodsNormal) {

}
func DealWithOneGoodPosition(pos GoodsPosition) {

}
