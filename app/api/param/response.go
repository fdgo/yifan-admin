package param

import (
	"time"
	"yifan/app/db"
)

type RespGetUser struct {
	UserName string `json:"user_name,omitempty"`
	ID       int64  `json:"id,omitempty"`
}

type RespUpLoadGoods struct {
	IpIdSerId []IpIdSerId `json:"ipIdSerId,omitempty"`
}
type IpIdSerId struct {
	IpName   string `json:"ipName,omitempty"`
	SerName  string `json:"serName,omitempty"`
	GoodName string `json:"goodName,omitempty"`
	Tip      string `json:"tip,omitempty"`
}
type RespSearchGoods struct {
	Goods Goods `json:"goods"`
}
type RespQueryGoods struct {
	AllPages  float64   `json:"allPages,omitempty"`
	GoodsInfo GoodsInfo `json:"goodsInfo,omitempty"`
}
type GoodsInfo struct {
	Num   int     `json:"num,omitempty"`
	Goods []Goods `json:"goods,omitempty"`
}
type Goods struct {
	GoodsId         uint      `json:"goodsId,omitempty"`
	Pic             string    `json:"pic,omitempty"`
	PicIndex        string    `json:"picIndex,omitempty"`
	Price           float64   `json:"price,omitempty"`
	Name            string    `json:"name,omitempty"`
	PkgStatus       int8      `json:"pkgStatus,omitempty"`
	Introduce       string    `json:"introduce,omitempty"`
	CreateTime      time.Time `json:"createTime,omitempty"`
	IpID            *uint     `json:"ipId,omitempty"`
	IpName          string    `json:"ipName,omitempty"`
	SeriesID        *uint     `json:"seriesId,omitempty"`
	SeriesName      string    `json:"seriesName,omitempty"`
	SingleOrMuti    int       `json:"singleOrMuti,omitempty"`
	MultiIds        []int     `json:"multiIds,omitempty"`
	WhoUpdate       string    `json:"whoUpdate,omitempty"`
	Integral        int32     `json:"integral,omitempty"`
	PreStore        string    `json:"preStore,omitempty"`
	ActiveBeginTime int64     `json:"activeBeginTime,omitempty"`
	ActiveEndTime   int64     `json:"activeEndTime,omitempty"`
}

//////////////////////////////////////////////////////////////////////
type RespUpLoadIPs struct {
	IpIdNames []IpIdName `json:"ipIdNames,omitempty"`
}
type IpIdName struct {
	Id   uint   `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Tip  string `json:"tip,omitempty"`
}
type RespSearchIp struct {
	ID         uint      `json:"id,omitempty"`
	Name       string    `json:"name,omitempty"`
	CreateName string    `json:"createName,omitempty"`
	CreateTime time.Time `json:"createTime,omitempty"`
}
type RespQueryIP struct {
	AllPages float64 `json:"allPages,omitempty"`
	IpInfo   IpInfo  `json:"ipInfo,omitempty"`
}
type IpInfo struct {
	Num    int      `json:"num,omitempty"`
	RespIp []RespIp `json:"respIp,omitempty"`
}
type RespIp struct {
	ID         uint      `json:"id,omitempty"`
	Name       string    `json:"name,omitempty"`
	CreateName string    `json:"createName,omitempty"` //创建人
	CreateTime time.Time `json:"createTime,omitempty"`
}

////////////////////////////////////////////////////////////////////
type RespUpLoadSeries struct {
	SeriIdNames []SeriIdName `json:"seriIdNames,omitempty"`
}
type SeriIdName struct {
	IpId uint   `json:"ipId,omitempty"`
	Name string `json:"name,omitempty"`
	Tip  string `json:"tip"`
}

type RespSearchSeries struct {
	SerInfo []SerInfo `json:"serInfo"`
}
type SerInfo struct {
	Id         uint      `json:"id,omitempty"`
	Name       string    `json:"name,omitempty"`
	CreateName string    `json:"createName,omitempty"`
	IpId       uint      `json:"ipId,omitempty"`
	IpName     string    `json:"ipName,omitempty"`
	CreateTime time.Time `json:"createTime,omitempty"`
}

type RespQuerySeries struct {
	AllPages   float64    `json:"allPages,omitempty"`
	ServieInfo ServieInfo `json:"servieInfo,omitempty"`
}
type ServieInfo struct {
	Num     int       `json:"num,omitempty"`
	Servies []Servies `json:"servies,omitempty"`
}
type Servies struct {
	Id         *uint     `json:"id,omitempty"`
	Name       string    `json:"name,omitempty"`
	CreateName string    `json:"createName,omitempty"`
	IpId       *uint     `json:"ipId,omitempty"`
	IpName     string    `json:"ipName,omitempty"`
	CreateTime time.Time `json:"createTime,omitempty"`
}

///////////////////////////////////////////////////////////
type RespQueryBoxAllPrize struct {
	TotalBoxPrizeNum int            `json:"totalBoxPrizeNum,omitempty"`
	LeftBoxPrizeNum  int            `json:"leftBoxPrizeNum,omitempty"`
	BoxStatus        int            `json:"boxStatus"`
	BoxIndex         int32          `json:"boxIndex,omitempty"` //箱子的顺序
	BoxId            uint           `json:"boxId,omitempty"`
	EachBoxPrize     []EachBoxPrize `json:"eachBoxPrize,omitempty"`
}
type RespQueryBoxAllRecord struct {
	TotalBoxPrizeNum int           `json:"totalBoxPrizeNum,omitempty"`
	LeftBoxPrizeNum  int           `json:"leftBoxPrizeNum,omitempty"`
	BoxId            uint          `json:"boxId"`
	BoxIndex         int32         `json:"boxIndex"`
	BoxStatus        int           `json:"boxStatus"`
	AllPages         float64       `json:"allPages,omitempty"`
	BoxRecordInfo    BoxRecordInfo `json:"boxRecordInfo,omitempty"`
}
type BoxRecordInfo struct {
	Num     int      `json:"num,omitempty"`
	Records []Record `json:"records,omitempty"`
}
type Record struct {
	Id             uint    `json:"id"`
	FanId          uint    `json:"fanId"`
	FanName        string  `json:"fanName"`
	BoxId          uint    `json:"boxId"`
	BoxIndex       int     `json:"boxIndex"`
	BoxName        string  `json:"boxName"`
	PrizeIndex     uint    `json:"prizeIndex"`
	PrizeName      string  `json:"prizeName"`
	PrizeIndexName string  `json:"prizeIndexName"`
	Position       string  `json:"position"`
	UserId         uint    `json:"userId"`
	NickName       string  `json:"nickName"`
	Price          float64 `json:"price"`
	Time           string  `json:"time"`
}
type RespQueryBoxLeftPrize struct {
	TotalBoxPrizeNum int            `json:"totalBoxPrizeNum,omitempty"`
	LeftBoxPrizeNum  int            `json:"leftBoxPrizeNum,omitempty"`
	BoxStatus        int            `json:"boxStatus"`
	BoxIndex         int32          `json:"boxIndex,omitempty"` //箱子的顺序
	BoxId            uint           `json:"boxId,omitempty"`
	EachBoxPrize     []EachBoxPrize `json:"eachBoxPrize,omitempty"`
}
type RespAddBox struct {
	Tips []Tips `json:"tips,omitempty"`
}
type Tips struct {
	PrizeName      string `json:"prizeName,omitempty"`
	PrizeIndex     int32  `json:"prizeIndex,omitempty"`
	PrizeIndexName string `json:"prizeIndexName,omitempty"`
	Message        string `json:"message,omitempty"`
}
type RespQueryGoodsForBox struct {
	AllPages float64 `json:"allPages,omitempty"`
	GInfo    GInfo   `json:"goodsInfo,omitempty"`
}
type GInfo struct {
	Num  int    `json:"num,omitempty"`
	Good []Good `json:"goods,omitempty"`
}

type Good struct {
	ID           uint    `json:"id"`
	IpID         *uint   `json:"ipId"`
	IpName       string  `json:"ipName"`
	SeriesID     *uint   `json:"seriesId"`
	SeriesName   string  `json:"seriesName"`
	Pic          string  `json:"pic"`
	Price        float64 `json:"price"`
	Name         string  `json:"name"`
	SingleOrMuti int     `json:"singleOrMuti"`
	MultiIds     []int   `json:"multiIds"`
	PkgStatus    int8    `json:"pkgStatus"`
	Introduce    string  `json:"introduce"`
	Integral     int32   `json:"integral"`
	Prestore     string  `json:"prestore"`
}

//type X struct {
//	Ret []Ret
//}
type Ele struct {
	FanId          uint   `json:"fanId"`
	FanTitle       string `json:"fanTitle"`
	BoxId          uint   `json:"boxId"`
	Num            int32  `json:"num"`
	PrizeIndexName string `json:"prizeIndexName"`
	PrizeName      string `json:"prizeName"`
	Status         int    `json:"status"`
	Postion        string `json:"postion"`
}
type RespPageOfPosition struct {
	Ele []Ele `json:"eles"`
}
type RespPageOfPositionCondition struct {
	Ele []Ele `json:"eles"`
}

type Boxes struct {
	PrizeNum int      `json:"prizeNum"`
	PrizeA   []PrizeA `json:"prizes"`
}
type PrizeA struct {
	BoxId          uint   `json:"boxId,omitempty"`
	PrizeIndexName string `json:"prizeIndexName,omitempty"`
	Num            int32  `json:"num"`
	PrizeName      string `json:"prizeName"`
	PrizeIndex     int32  `json:"prizeIndex,omitempty"`
	Position       []int  `json:"position,omitempty"`
}

//////////////////////////////////////////////////////////////
type RespAddFan struct {
}
type RespQueryFanStatus struct {
	FanId     uint        `json:"fanId"`
	FanTitle  string      `json:"fanTitle"`
	BoxStatus []BoxStatus `json:"boxStatus"`
}
type RespQueryFanStatusCondition struct {
	FanId           uint      `json:"fanId,omitempty"`
	FanTitle        string    `json:"fanTitle"`
	TotalBoxNum     int       `json:"totalBoxNum"`
	TotalPrizeNum   int32     `json:"totalPrizeNum"`
	LeftPrizeNum    int32     `json:"leftPrizeNum"`
	Status          int       `json:"status,omitempty"`
	Price           float64   `json:"price,omitempty"`
	SharePic        string    `json:"sharePic"`
	DetailPic       string    `json:"detailPic"`
	ActiveBeginTime int64     `json:"activeBeginTime,omitempty"`
	ActiveEndTime   int64     `json:"activeEndTime,omitempty"`
	CreateTime      time.Time `json:"createTime,omitempty"`
	WhoUpdate       string    `json:"whoUpdate,omitempty"`
}
type BoxStatus struct {
	BoxId  uint `json:"boxId"`
	Status int  `json:"status"`
}
type RespQueryFan struct {
	AllPages float64 `json:"allPages,omitempty"`
	FanInfos FanInfo `json:"fanInfos,omitempty"`
}
type FanInfo struct {
	Num  int   `json:"num,omitempty"`
	Fans []Fan `json:"fans,omitempty"`
}
type Fan struct {
	ID              uint      `json:"Id,omitempty"`
	Title           string    `json:"title,omitempty"`
	TotalBoxNum     int       `json:"totalBoxNum"`
	TotalPrizeNum   int32     `json:"totalPrizeNum"`
	LeftPrizeNum    int32     `json:"leftPrizeNum"`
	Status          int       `json:"status,omitempty"`
	Price           float64   `json:"price,omitempty"`
	SharePic        string    `json:"sharePic"`
	DetailPic       string    `json:"detailPic"`
	ActiveBeginTime int64     `json:"activeBeginTime,omitempty"`
	ActiveEndTime   int64     `json:"activeEndTime,omitempty"`
	CreateTime      time.Time `json:"createTime,omitempty"`
	WhoUpdate       string    `json:"whoUpdate,omitempty"`
}
type RespModifyFan struct {
	FanId           uint           `json:"fanId,omitempty"`           //蕃的Id
	FanName         string         `json:"fanName,omitempty"`         //蕃的名字
	Type            string         `json:"ty,omitempty"`              //玩法类型
	Rule            string         `json:"rule,omitempty"`            //活动规则
	Title           string         `json:"title,omitempty"`           //活动标题
	FanPrice        float64        `json:"fanPrice,omitempty"`        //蕃的价格
	ActiveBeginTime int64          `json:"activeBeginTime,omitempty"` //活动开始时间
	ActiveEndTime   int64          `json:"activeEndTime,omitempty"`   //活动结束时间
	DetailPic       string         `json:"detailPic,omitempty"`       //详细图片
	SharePic        string         `json:"sharePic,omitempty"`        //分享图片
	BoxNum          int            `json:"boxNum,omitempty"`          //一个蕃的箱数
	WhoUpdate       string         `json:"whoUpdate,omitempty"`
	EachBoxPrize    []EachBoxPrize `json:"eachBoxPrize,omitempty"`
}
type EachBoxPrize struct {
	PrizeId        uint        `json:"prizeId,omitempty"`   //奖品id
	GoodId         uint        `json:"goodId,omitempty"`    //形成奖品的商品id
	PrizeName      string      `json:"prizeName,omitempty"` //奖品名
	PrizeNum       int32       `json:"prizeNum"`            //某一个种类奖品数量
	PrizeLeftNum   int32       `json:"prizeLeftNum"`
	PrizeIndex     int32       `json:"prizeIndex,omitempty"` //某一个种类奖品在箱子中的序号
	PrizeIndexName string      `json:"prizeIndexName,omitempty"`
	PrizeStyle     string      `json:"prizeStyle,omitempty"` //抽取方式
	PrizeRate      string      `json:"prizeRate,omitempty"`
	PrizeStatus    int         `json:"prizeStatus"`
	Rate           string      `json:"rate"`
	Remark         string      `json:"remark"`
	Position       []int       `json:"position"`
	IpId           uint        `json:"ipId,omitempty"`         //该奖品所属IP
	IpName         string      `json:"ipName,omitempty"`       //该奖品所属IP的名字
	SeriId         uint        `json:"seriId,omitempty"`       //该奖品所属系列
	SeriName       string      `json:"seriName,omitempty"`     //该奖品所属系列名字
	Pic            string      `json:"pic,omitempty"`          //奖品图片
	PkgStatus      int         `json:"pkgStatus,omitempty"`    //品相状态
	SingleOrMuti   int         `json:"singleOrMuti,omitempty"` //单一商品填1, 有n个组合就写n
	MultiIds       db.GormList `json:"multiIds,omitempty"`     //商品id组合,单一商品[435],
}
type RespModifySaveFan struct {
	Tips []Tips `json:"tips,omitempty"`
}
type RespEnterFan struct {
	FanId            uint           `json:"fanId,omitempty"`           //蕃的Id
	FanName          string         `json:"fanName,omitempty"`         //蕃的名字
	Type             string         `json:"ty,omitempty"`              //玩法类型
	Rule             string         `json:"rule,omitempty"`            //活动规则
	Title            string         `json:"title,omitempty"`           //活动标题
	FanPrice         float64        `json:"fanPrice,omitempty"`        //蕃的价格
	Status           int            `json:"status,omitempty"`          //蕃的状态  1.上架  2.下架
	ActiveBeginTime  int64          `json:"activeBeginTime,omitempty"` //活动开始时间
	ActiveEndTime    int64          `json:"activeEndTime,omitempty"`   //活动结束时间
	DetailPic        string         `json:"detailPic,omitempty"`       //详细图片
	SharePic         string         `json:"sharePic,omitempty"`        //分享图片
	TotalBoxPrizeNum int            `json:"totalBoxPrizeNum,omitempty"`
	LeftBoxPrizeNum  int            `json:"leftBoxPrizeNum,omitempty"`
	BoxIndex         int32          `json:"boxIndex,omitempty"` //箱子的顺序
	TotalBoxNum      int            `json:"totalBoxNum,omitempty"`
	BoxId            uint           `json:"boxId,omitempty"`
	EachBoxPrize     []EachBoxPrize `json:"eachBoxPrize,omitempty"`
}

type RespQueryPrizePostion struct {
	QueryPrizePostions []QueryPrizePostion `json:"queryPrizePostions,omitempty"`
}
type QueryPrizePostion struct {
	FanId          uint   `json:"fanId,omitempty"`
	FanTitle       string `json:"fanTitle,omitempty"`
	BoxId          uint   `json:"boxId,omitempty"`
	BoxTitle       string `json:"boxTitle,omitempty"`
	PrizeNum       int32  `json:"prizeNum,omitempty"`
	PrizeIndexName string `json:"prizeIndexName,omitempty"`
	PrizeName      string `json:"prizeName,omitempty"`
	Position       []int  `json:"position,omitempty"`
}

type RespModifyGoodsPosition struct {
	BoxId uint `json:"boxId,omitempty"`
}

type RespBuy struct {
	Money float64 `json:"money,omitempty"`
}

type RespBuySures struct {
	BuySures []BuySure `json:"buySures,omitempty"`
}
type RespBuySure struct {
	BuySure BuySure `json:"buySure,omitempty"`
}
type BuySure struct {
	Index          int       `json:"index"`
	PrizeIndexName string    `json:"prizeIndexName,omitempty"`
	PrizeIndex     int       `json:"prizeIndex,omitempty"`
	PrizeName      string    `json:"prizeName,omitempty"`
	Price          float64   `json:"price,omitempty"`
	Pic            string    `json:"pic,omitempty"`
	Time           time.Time `json:"time"`
}

type RespBuyQuerys struct {
	RespBuyQuerys []RespBuyQuery `json:"respBuyQuerys"`
}
type RespBuyQuery struct {
	Index          int     `json:"index"`
	PrizeIndexName string  `json:"prizeIndexName,omitempty"`
	PrizeIndex     int     `json:"prizeIndex,omitempty"`
	PrizeName      string  `json:"prizeName,omitempty"`
	Price          float64 `json:"price,omitempty"`
	Pic            string  `json:"pic,omitempty"`
}

//////////////////////
type RespGetOpenId struct {
	JwtToken string `json:"jwt_token"`
	UserId   uint   `json:"user_id"`
	NickName string `json:"nick_name"`
	Avatar   string `json:"avatar"`
}

type RespPageOfOrder struct {
	AllPages float64 `json:"allPages,omitempty"`
	Num      int     `json:"num"`
	Orders   []Order `json:"orders"`
}
type Order struct {
	ID         uint     `json:"-"`
	OutTradeNo string   `json:"orderId"`
	FanPic     string   `json:"fanPic"`
	Price      float64  `json:"price"`
	PrizeNum   int      `json:"prizeNum"`
	OrderType  string   `json:"orderType"`
	UserName   string   `json:"userName"`
	PayStyle   string   `json:"payStyle"`
	CreateTime int64    `json:"createTime"`
	Status     string   `json:"status"`
	FanId      uint     `json:"-"`
	FanName    string   `json:"-"`
	BoxId      uint     `json:"-"`
	BoxIndex   int      `json:"-"`
	OpenId     string   `json:"-"`
	UserId     uint     `json:"-"`
	Avatar     string   `json:"-"`
	UserMobile string   `json:"-"`
	PrepayId   string   `json:"-"`
	Appid      string   `json:"-"`
	Remark     string   `json:"remark"`
	Detail     string   `json:"detail"`
	Operator   string   `json:"operator"`
	Goods      []Goodxs `json:"goods"`
}
type RespPageOfOrderCondition struct {
	AllPages float64 `json:"allPages,omitempty"`
	Num      int     `json:"num"`
	Orders   []Order `json:"orders"`
}
type RespPageOfOrderDetail struct {
	Orders Order `json:"orders"`
}
type Goodxs struct {
	IpID           uint   `json:"ipId"`
	IpName         string `json:"ipName"`
	SeriesID       uint   `json:"seriesId"`
	SeriesName     string `json:"seriesName"`
	PrizeName      string `json:"prizeName"`
	PrizeIndexName string `json:"prizeIndexName"`
	PrizeId        uint   `json:"prizeId"`
	Pic            string `json:"pic"`
}

type RespActiveByMan struct {
}
type RespSingleClick struct {
	FanPicTitle []FanPicTitle `json:"fanPicTitle"`
}
type FanPicTitle struct {
	Pic   string `json:"pic"`
	Title string `json:"title"`
}
type RespGetBannerPic struct {
	AdverTip        string   `json:"adverTip"`
	BannerPic       []string `json:"bannerPic"`
	ActiveBeginTime int64    `json:"activeBeginTime"`
	ActiveEndTime   int64    `json:"activeEndTime"`
}
type RespAddSecondTab struct {
}
type RespAddSecondTabSon struct {
}
type RespQuerySecondTab struct {
	SecondTab []string `json:"secondTab"`
}

type RespQuerySecondSonTab struct {
	Tab []Tab `json:"tab"`
}
type Tab struct {
	TabTag          string `json:"tabTag"`
	TabSon          string `json:"tabSon"`
	RedirectType    string `json:"redirect_type"`
	RedirectAddress string `json:"redirect_address"`
	ActiveBeginTime int64  `json:"active_begin_time"`
	ActiveEndTime   int64  `json:"active_end_time"`
	Remark          string `json:"remark"`
	Title           string `json:"title"`
	Pic             string `json:"pic"`
}

type RespShowOrHideSecondTab struct {
}
type RespModifyAndSaveSecondTab struct {
}
