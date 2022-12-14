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
type RespGlobalConfig struct {
}
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
	CreateName string    `json:"createName,omitempty"` //?????????
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
	BoxIndex         int32          `json:"boxIndex,omitempty"` //???????????????
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
	BoxIndex         int32          `json:"boxIndex,omitempty"` //???????????????
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
	FanId           uint           `json:"fanId,omitempty"`           //??????Id
	FanName         string         `json:"fanName,omitempty"`         //????????????
	Type            string         `json:"ty,omitempty"`              //????????????
	Rule            string         `json:"rule,omitempty"`            //????????????
	Title           string         `json:"title,omitempty"`           //????????????
	FanPrice        float64        `json:"fanPrice,omitempty"`        //????????????
	ActiveBeginTime int64          `json:"activeBeginTime,omitempty"` //??????????????????
	ActiveEndTime   int64          `json:"activeEndTime,omitempty"`   //??????????????????
	DetailPic       string         `json:"detailPic,omitempty"`       //????????????
	SharePic        string         `json:"sharePic,omitempty"`        //????????????
	BoxNum          int            `json:"boxNum,omitempty"`          //??????????????????
	WhoUpdate       string         `json:"whoUpdate,omitempty"`
	EachBoxPrize    []EachBoxPrize `json:"eachBoxPrize,omitempty"`
}
type EachBoxPrize struct {
	PrizeId        uint        `json:"prizeId,omitempty"`   //??????id
	GoodId         uint        `json:"goodId,omitempty"`    //?????????????????????id
	PrizeName      string      `json:"prizeName,omitempty"` //?????????
	PrizeNum       int32       `json:"prizeNum"`            //???????????????????????????
	PrizeLeftNum   int32       `json:"prizeLeftNum"`
	PrizeIndex     int32       `json:"prizeIndex,omitempty"` //??????????????????????????????????????????
	PrizeIndexName string      `json:"prizeIndexName,omitempty"`
	PrizeStyle     string      `json:"prizeStyle,omitempty"` //????????????
	PrizeRate      string      `json:"prizeRate,omitempty"`
	PrizeStatus    int         `json:"prizeStatus"`
	Rate           string      `json:"rate"`
	Remark         string      `json:"remark"`
	Position       []int       `json:"position"`
	IpId           uint        `json:"ipId,omitempty"`         //???????????????IP
	IpName         string      `json:"ipName,omitempty"`       //???????????????IP?????????
	SeriId         uint        `json:"seriId,omitempty"`       //?????????????????????
	SeriName       string      `json:"seriName,omitempty"`     //???????????????????????????
	Pic            string      `json:"pic,omitempty"`          //????????????
	PkgStatus      int         `json:"pkgStatus,omitempty"`    //????????????
	SingleOrMuti   int         `json:"singleOrMuti,omitempty"` //???????????????1, ???n???????????????n
	MultiIds       db.GormList `json:"multiIds,omitempty"`     //??????id??????,????????????[435],
}
type RespModifySaveFan struct {
	Tips []Tips `json:"tips,omitempty"`
}
type RespEnterFan struct {
	FanId            uint           `json:"fanId,omitempty"`           //??????Id
	FanName          string         `json:"fanName,omitempty"`         //????????????
	Type             string         `json:"ty,omitempty"`              //????????????
	Rule             string         `json:"rule,omitempty"`            //????????????
	Title            string         `json:"title,omitempty"`           //????????????
	FanPrice         float64        `json:"fanPrice,omitempty"`        //????????????
	Status           int            `json:"status,omitempty"`          //????????????  1.??????  2.??????
	ActiveBeginTime  int64          `json:"activeBeginTime,omitempty"` //??????????????????
	ActiveEndTime    int64          `json:"activeEndTime,omitempty"`   //??????????????????
	DetailPic        string         `json:"detailPic,omitempty"`       //????????????
	SharePic         string         `json:"sharePic,omitempty"`        //????????????
	TotalBoxPrizeNum int            `json:"totalBoxPrizeNum,omitempty"`
	LeftBoxPrizeNum  int            `json:"leftBoxPrizeNum,omitempty"`
	BoxIndex         int32          `json:"boxIndex,omitempty"` //???????????????
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
type RespUserList struct {
	AllPages float64 `json:"allPages"`
	Num      int     `json:"num"`
	Users    []User  `json:"users"`
}
type RespUserListCondition struct {
	AllPages float64 `json:"allPages"`
	Num      int     `json:"num"`
	Users    []User  `json:"users"`
}
type User struct {
	UserId         uint      `json:"userId"`
	NickName       string    `json:"nickName"`
	Mobile         string    `json:"mobile"`
	AppId          string    `json:"appId"`
	OpenId         string    `json:"openId"`
	Avatar         string    `json:"avatar"`
	CreatTime      string    `json:"creatTime"`
	SessionKey     string    `json:"sessionKey"`
	RealName       string    `json:"realName"`
	LuggageNum     int64     `json:"luggageNum"`
	ConsumptionFee float64   `json:"consumptionFee"`
	Luggage        []Luggage `json:"luggage"`
}
type Luggage struct {
	ID             uint    `json:"id"`
	OutTradeNo     string  `json:"outTradeNo"`
	GoodID         uint    `json:"goodId"`
	GoodName       string  `json:"goodName"`
	IpID           uint    `json:"ipId"`
	IpName         string  `json:"ipName"`
	SeriesID       uint    `json:"seriesId"`
	SeriesName     string  `json:"seriesName"`
	Pic            string  `json:"pic"`
	Price          float64 `json:"price"`
	PrizeIndexName string  `json:"prizeIndexName"`
	PrizeIndex     int     `json:"prizeIndex"`
	InLuggageTime  int64   `json:"inLuggageTime"`
	OutLuggageTime int64   `json:"outLuggageTime"`
	DeleStatus     int     `json:"deleStatus"`
}
type RespDelever struct {
	AllPages    float64      `json:"allPages"`
	Num         int          `json:"num"`
	OneDelevers []OneDelever `json:"oneDelevers"`
}
type RespDeleverCondition struct {
	AllPages    float64      `json:"allPages"`
	Num         int          `json:"num"`
	OneDelevers []OneDelever `json:"oneDelevers"`
}
type OneDelever struct {
	Price           float32         `json:"price"`
	DeleCompany     string          `json:"deleCompany"`
	DeleOrderId     uint            `json:"deleOrderId"`
	ActiveSureTime  int64           `json:"activeSureTime"`
	OrderId         uint            `json:"orderId"`
	CreateTime      string          `json:"createTime"`
	DeleverUserInfo DeleverUserInfo `json:"deleverUserInfo"`
	Num             int             `json:"num"`
	DeleverStatus   int             `json:"deleverStatus"`
	DeleverGoods    []DeleverGood   `json:"deleverGoods"`
}

type RespDeleverDetail struct {
	Price           float32         `json:"price"`
	DeleCompany     string          `json:"deleCompany"`
	DeleOrderId     uint            `json:"deleOrderId"`
	ActiveSureTime  int64           `json:"activeSureTime"`
	OrderId         uint            `json:"orderId"`
	CreateTime      string          `json:"createTime"`
	DeleverUserInfo DeleverUserInfo `json:"deleverUserInfo"`
	Num             int             `json:"num"`
	DeleverStatus   int             `json:"deleverStatus"`
	DeleverGoods    []DeleverGood   `json:"deleverGoods"`
	PayStyle        string          `json:"payStyle"`
}
type DeleverUserInfo struct {
	Address  string `json:"address"`
	UserName string `json:"userName"`
	Mobile   string `json:"mobile"`
	UserId   uint   `json:"userId"`
}
type DeleverGood struct {
	Pic            string  `json:"pic"`
	GoodName       string  `json:"goodName"`
	IpName         string  `json:"ipName"`
	SeriesName     string  `json:"seriesName"`
	PkgStatus      int     `json:"pkgStatus"`
	PrizeIndexName string  `json:"prizeIndexName"`
	GoodId         uint    `json:"goodId"`
	PrizeId        uint    `json:"prizeId"`
	DeleStatus     int     `json:"deleStatus"`
	Price          float64 `json:"price"`
}
type RespPageOfOrder struct {
	AllPages float64 `json:"allPages,omitempty"`
	Num      int     `json:"num"`
	Orders   []Order `json:"orders"`
}
type Order struct {
	ID           uint     `json:"-"`
	OutTradeNo   string   `json:"outTradeNo"`
	FanPic       string   `json:"fanPic"`
	Price        float64  `json:"price"`
	PrizeNum     int      `json:"prizeNum"`
	OrderType    string   `json:"orderType"`
	UserName     string   `json:"userName"`
	PayStyle     string   `json:"payStyle"`
	CreateTime   int64    `json:"createTime"`
	Status       string   `json:"status"`
	FanId        uint     `json:"-"`
	FanName      string   `json:"-"`
	BoxId        uint     `json:"-"`
	BoxIndex     int      `json:"-"`
	OpenId       string   `json:"-"`
	UserId       uint     `json:"userId"`
	Avatar       string   `json:"-"`
	UserMobile   string   `json:"userMobile"`
	FinishTime   int64    `json:"finishTime"`
	TansactionId string   `json:"tansactionId"`
	Appid        string   `json:"-"`
	Remark       string   `json:"remark"`
	Detail       string   `json:"detail"`
	Operator     string   `json:"operator"`
	Goods        []Goodxs `json:"goods"`
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
	FanId uint   `json:"fanId"`
	Pic   string `json:"pic"`
	Title string `json:"title"`
}
type RespGetBannerPic struct {
	Banners []Banner `json:"banners"`
}

type RespAddSecondTab struct {
}
type RespAddSecondTabSon struct {
}
type RespQuerySecondTab struct {
	TabSon1  string `json:"tabSon1"`
	TabSon2  string `json:"tabSon2"`
	TabSon3  string `json:"tabSon3"`
	TabSon4  string `json:"tabSon4"`
	TabSon5  string `json:"tabSon5"`
	TabSon6  string `json:"tabSon6"`
	TabSon7  string `json:"tabSon7"`
	TabSon8  string `json:"tabSon8"`
	TabSon9  string `json:"tabSon9"`
	TabSon10 string `json:"tabSon10"`
	IsHide1  bool   `json:"isHide1"`
	IsHide2  bool   `json:"isHide2"`
	IsHide3  bool   `json:"isHide3"`
	IsHide4  bool   `json:"isHide4"`
	IsHide5  bool   `json:"isHide5"`
	IsHide6  bool   `json:"isHide6"`
	IsHide7  bool   `json:"isHide7"`
	IsHide8  bool   `json:"isHide8"`
	IsHide9  bool   `json:"isHide9"`
	IsHide10 bool   `json:"isHide10"`
}

type RespQuerySecondSonTab struct {
	TabContent []TabContent `json:"tabContent"`
}
type TabContent struct {
	Id              uint   `json:"id"`
	TabTag          string `json:"tabTag"`
	TabSon          string `json:"tabSon"`
	ActiveBeginTime int64  `json:"activeBeginTime"`
	ActiveEndTime   int64  `json:"activeEndTime"`
	Title           string `json:"title"`
	FanId           uint   `json:"fanId"`
	FanTitle        string `json:"fanTitle"`
	AdvPic          string `json:"advPic"`
	Remark          string `json:"remark"`
	IsHide          bool   `json:"isHide"`
	RedirectUrl     string `json:"redirectUrl"`
	RedirectType    string `json:"redirectType"`
}

type RespShowOrHideSecondTab struct {
}
type RespModifyAndSaveSecondTab struct {
}
