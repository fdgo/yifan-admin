package param

import (
	"yifan/app/db"
)

type ReqGetUser struct {
	Id uint `json:"id" binding:"required" example:"123456789" ` //用户的id()
}

type ReqAddRankingQuery struct {
	Id     uint                   `json:"id" binding:"required" example:"123456789" ` //用户的id()
	BoxId  uint                   `json:"boxId"`
	Params map[string]interface{} `json:"params"` //修改用户属性()
}

type ReqMakeSureBuy struct {
	DiffTime uint `json:"diffTime"`
	BoxId    uint `json:"boxId"`
}
type ReqGetAllGoods struct {
	PageSize  int32 `json:"pageSize"`
	PageIndex int32 `json:"pageIndex"`
	StyleId   uint  `json:"styleId"`
	PkgStatus int8  `json:"pkgStatus"`
}

type ReqIsNeedToJoinQueue struct {
}

//
type ReqUpLoadGoods struct {
	UpLoadGoods []UpLoadGoods `json:"goods"`
}
type UpLoadGoods struct {
	IpName       string  `json:"ipName" binding:"required"`
	SeriesName   string  `json:"seriesName" binding:"required"`
	GoodsName    string  `json:"goodsName" binding:"required"` //商品名
	PkgStatus    int8    `json:"pkgStatus" binding:"required"` //打包状态
	PreStore     string  `json:"preStore" binding:"required"`
	Integral     int32   `json:"integral" binding:"required"` //积分
	Pic          string  `json:"pic" binding:"required"`      //图片
	Price        float64 `json:"price" binding:"required"`    //售价
	CreateTime   string  `json:"createTime" binding:"required"`
	SingleOrMuti int     `json:"singleOrMuti"`
	MultiIds     []int   `json:"multiIds"`
	Introduce    string  `json:"introduce"`
}

type ReqSearchGoods struct {
	Search string
}
type ReqAddGoods struct {
	GoodsName    string  `json:"goodsName" binding:"required"` //商品名
	IpId         *uint   `json:"ipId" binding:"required"`      //所属IP
	IpName       string  `json:"ipName" binding:"required"`
	SeriesId     *uint   `json:"seriesId" binding:"required"` //所属系列
	SeriesName   string  `json:"seriesName" binding:"required"`
	Pic          string  `json:"pic" binding:"required"`       //图片
	Price        float64 `json:"price" binding:"required"`     //售价
	SingleOrMuti int     `json:"singleOrMuti"`                 //单个还是多个
	MultiIds     []int   `json:"multiIds" `                    //多个商品id
	PkgStatus    int8    `json:"pkgStatus" binding:"required"` //打包状态
	Integral     int32   `json:"integral"`                     //积分
	PreStore     string  `json:"preStore" binding:"required"`  //现货状态
}
type ReqBoxBindGoods struct {
}
type ReqDeleteGoods struct {
	GoodsId *uint `json:"goodsId" binding:"required,gt=0"`
	Num     *uint `json:"num" binding:"required,gt=0"`
}
type ReqQueryGoods struct {
	PageSize  int32 `json:"pageSize" binding:"required"`
	PageIndex int32 `json:"pageIndex" binding:"required"`
}
type ReqModifyGoods struct {
	GoodsId         *uint       `json:"goodsId" binding:"required,gt=0"`
	IpID            *uint       `json:"ipId"`
	SeriesID        *uint       `json:"seriesId"`
	Pic             string      `json:"pic"`
	Price           float64     `json:"price"`
	Name            string      `json:"name"`
	SingleOrMuti    int         `json:"singleOrMuti"`
	MultiIds        db.GormList `json:"multiIds"`
	PkgStatus       int8        `json:"pkgStatus"`
	Introduce       string      `json:"introduce"`
	Integral        int32       `json:"integral"`
	PreStore        int         `json:"preStore"`
	WhoUpdate       string      `json:"whoUpdate"`
	ActiveBeginTime int64       `json:"activeBeginTime"`
	ActiveEndTime   int64       `json:"activeEndTime"`
}

type ReqGlobalConfig struct {
	Fee         float64 `json:"fee"`
	FeeLimitNum int     `json:"feeLimitNum"`
}

//
type ReqUpLoadIPs struct {
	ReqAddIP []ReqAddIP `json:"ips"`
}
type ReqSearchIP struct {
	Search string `json:"search"`
}
type ReqAddIP struct {
	Name       string `json:"name" binding:"required,lte=32,gte=1"`       //IP 名字
	CreateName string `json:"createName" binding:"required,lte=32,gte=1"` //创建人
	CreateTime int64  `json:"createTime"`
}
type ReqDeleteIP struct {
	Id *int `json:"id" binding:"required"` //IP id
}
type ReqQueryIP struct {
	PageSize  int32 `json:"pageSize" binding:"required"`
	PageIndex int32 `json:"pageIndex" binding:"required"`
}
type ReqModifyIP struct {
	Id         *int   `json:"id" binding:"required"`                      //IP id
	Name       string `json:"name"`                                       //IP 名字
	CreateName string `json:"createName" binding:"required,lte=32,gte=1"` //创建人
}

//
type ReqUpLoadSeries struct {
	ReqAddSeries []ReqAddSeries `json:"series"`
}
type ReqAddSeries struct {
	Name       string `json:"name" binding:"required,lte=32,gte=1"`
	CreateName string `json:"createName" binding:"required,lte=32,gte=1"` //创建人
	CreateTime int64  `json:"createTime"`
	IpId       *uint  `json:"ipId" binding:"required"`
	IpName     string `json:"ipName" binding:"required"`
}

type ReqSearchSeries struct {
	Search string `json:"search"`
}
type ReqDeleteSeries struct {
	SerId *uint `json:"serId" binding:"required"`
}
type ReqQuerySeries struct {
	PageSize  int32 `json:"pageSize" binding:"required"`
	PageIndex int32 `json:"pageIndex" binding:"required"`
}
type ReqModifySeries struct {
	Id         *int   `json:"id" binding:"required"` //系列 id
	Name       string `json:"name" binding:"required,lte=32,gte=1"`
	CreateName string `json:"createName" binding:"required,lte=32,gte=1"` //创建人
	IpId       uint   `json:"ipId"`
}

//
type ReqAddBox struct {
	Type            string  `json:"type"`            //蕃的类型
	FanPrice        float64 `json:"fanPrice"`        //蕃的价格
	ActiveBeginTime int64   `json:"activeBeginTime"` //上架时间
	ActiveEndTime   int64   `json:"activeEndTime"`   //下架时间
	BoxNum          int     `json:"boxNum"`          //一个蕃的箱数
	Rule            string  `json:"rule"`            //活动规则
	Title           string  `json:"title"`           //活动标题
	DetailPic       string  `json:"detailPic"`       //详细图片
	SharePic        string  `json:"sharePic"`        //分享图片
	Boxes           Box     `json:"box"`             //所有箱数的数组
}
type Box struct {
	Status int     `json:"status"` //蕃的状态  1.上架  2.下架
	Prizes []Prize `json:"prizes"` //每个箱的所有奖品
}

type Prize struct {
	GoodId            uint        `json:"goodId"`         //奖品id
	GoodName          string      `json:"goodName"`       //奖品名
	PrizeNum          int32       `json:"prizeNum"`       //某一个种类奖品数量
	PrizeIndex        int32       `json:"prizeIndex"`     //某一个种类奖品在箱子中的序号
	PrizeIndexName    string      `json:"prizeIndexName"` //A赏,B赏
	Position          []int       `json:"position"`       //A赏,B赏 位置
	PrizeStyle        string      `json:"prizeStyle"`     //抽取方式
	Price             float64     `json:"price"`          //某种商品价格
	IpId              uint        `json:"ipId"`           //该奖品所属IP
	IpName            string      `json:"ipName"`         //该奖品所属IP的名字
	Remark            string      `json:"remark"`
	SeriId            uint        `json:"seriesId"`   //该奖品所属系列
	SeriName          string      `json:"seriesName"` //该奖品所属系列名字
	Pic               string      `json:"pic"`        //奖品图片
	PkgStatus         int         `json:"pkgStatus"`  //品相状态
	PreStore          int         `json:"preStore"`
	SoldStatus        int         `json:"soldStatus"`        //是否售罄1.奖品售罄,2.奖品未售罄
	TimeForSoldStatus string      `json:"timeForSoldStatus"` //预售时间
	SingleOrMuti      int         `json:"singleOrMuti"`      //单一商品填1, 有n个组合就写n
	MultiIds          db.GormList `json:"multiIds"`          //商品id组合,单一商品[435], n个商品[34,456,234,...]
}
type ReqPageOfPosition struct {
	FanIndex  int     `json:"fanIndex"`
	TimeRange []int64 `json:"timeRange"`
	Status    int     `json:"status"`
}
type ReqPageOfPositionCondition struct {
	FanIndex       uint    `json:"fanIndex"`
	BoxIndex       int     `json:"boxIndex"`
	PrizeIndexName string  `json:"prizeIndexName"`
	PrizeName      string  `json:"prizeName"`
	TimeRange      []int64 `json:"timeRange"`
	Status         int     `json:"status"`
}

type ReqSetNormalPrizePosition struct {
	Fanex []Fanex `json:"fanex"`
}
type Fanex struct {
	FanId uint    `json:"fanId"`
	Boxex []Boxex `json:"boxex"`
}
type Boxex struct {
	BoxId  uint     `json:"boxId"`
	Prizex []Prizex `json:"prizex"`
}
type Prizex struct {
	PrizeIndexName string `json:"prizeIndexName"`
	PrizeIndex     int    `json:"prizeIndex"`
	Num            int    `json:"num"`
	Pos            []int  `json:"pos"`
}

type ReqDeleteBox struct {
}
type ReqQueryBoxAllPrize struct {
	FanId    uint `json:"fanId"`
	BoxId    uint `json:"boxId"`
	BoxIndex int  `json:"boxIndex"`
	Tag      int  `json:"tag"`
}
type ReqQueryBoxAllRecord struct {
	FanId     uint  `json:"fanId"`
	BoxId     uint  `json:"boxId"`
	BoxIndex  int   `json:"boxIndex"`
	Tag       int   `json:"tag"`
	PageSize  int32 `json:"pageSize" binding:"required"`
	PageIndex int32 `json:"pageIndex" binding:"required"`
}
type ReqQueryBoxLeftPrize struct {
	FanId    uint `json:"fanId"`
	BoxId    uint `json:"boxId"`
	BoxIndex int  `json:"boxIndex"`
	Tag      int  `json:"tag"`
}
type ReqModifyBox struct {
	BoxId uint `json:"boxId"`
}
type ReqModifyBoxStatus struct {
	BoxId  uint `json:"boxId"`
	Status int  `json:"status"`
}
type ReqQueryGoodsForBox struct {
	PageSize  int32 `json:"pageSize" binding:"required"`
	PageIndex int32 `json:"pageIndex" binding:"required"`
}
type ReqGoodsToBePrize struct {
}

type ReqModifyBoxGoods struct {
	FanId             uint   `json:"fanId"`
	OldGoodId         uint   `json:"oldGoodId"`
	NewRemark         string `json:"newRemark"`
	NewGoodId         uint   `json:"newGoodId"`         //商品id
	NewGoodName       string `json:"newGoodName"`       //商品名
	NewPrizeNum       int32  `json:"newPrizeNum"`       //某一个种类奖品数量
	NewPrizeIndex     int32  `json:"newPrizeIndex"`     //某一个种类奖品在箱子中的序号
	NewPrizeIndexName string `json:"newPrizeIndexName"` //A赏,B赏
	NewPrizePosition  string `json:"newPrizePosition"`  //A赏,B赏 位置
	NewPkgStatus      int    `json:"newPkgStatus"`      //品相状态
}
type ReqDeleteBoxGoods struct {
	FanId  uint `json:"fanId"`
	GoodId int  `json:"goodId"`
}

//type ReqAddFan struct {
//	FanName       string  `json:"fanName"`       //蕃的名字
//	FanPrice      float64 `json:"fanPrice"`      //蕃的价格
//	Status        int     `json:"status"`        //蕃的状态  1.上架  2.下架
//	OnActiveTime  int64   `json:"onActiveTime"`  //活动开始时间
//	OffActiveTime int64   `json:"offActiveTime"` //活动结束时间
//	WhoCreated    string  `json:"whoCreated"`
//	Pic           string  `json:"pic"` //蕃的图片
//}
type ReqQueryFanStatus struct {
}
type ReqQueryFanStatusCondition struct {
	FanId uint `json:"fanId"`
}
type ReqQueryFan struct {
	PageSize  int32 `json:"pageSize" binding:"required"`
	PageIndex int32 `json:"pageIndex" binding:"required"`
}

type ReqModifyFanStatus struct {
	FanId  uint `json:"fanId" binding:"required"`
	Status int  `json:"status" binding:"required"`
}

type ReqModifyFan struct {
	FanId uint `json:"fanId" binding:"required"`
}

type ReqModifySaveFan struct {
	FanID           uint    `json:"fanId"`
	Type            string  `json:"type"`            //蕃的类型
	FanPrice        float64 `json:"fanPrice"`        //蕃的价格
	ActiveBeginTime int64   `json:"activeBeginTime"` //上架时间
	ActiveEndTime   int64   `json:"activeEndTime"`   //下架时间
	BoxNum          int     `json:"boxNum"`          //一个蕃的箱数
	Rule            string  `json:"rule"`            //活动规则
	Title           string  `json:"title"`           //活动标题
	DetailPic       string  `json:"detailPic"`       //详细图片
	SharePic        string  `json:"sharePic"`        //分享图片
	Boxes           Box     `json:"box"`             //所有箱数的数组
}
type PrizeX struct {
	GoodId         uint        `json:"goodId"`
	GoodName       string      `json:"goodName"`
	PrizeNum       int32       `json:"prizeNum"`       //某一个种类奖品数量
	PrizeIndex     int32       `json:"prizeIndex"`     //某一个种类奖品在箱子中的序号
	PrizeIndexName string      `json:"prizeIndexName"` //A赏,B赏
	Position       []int       `json:"position"`       //A赏,B赏 位置
	IpId           uint        `json:"ipId"`           //该奖品所属IP
	IpName         string      `json:"ipName"`         //该奖品所属IP的名字
	Remark         string      `json:"remark"`
	SeriId         uint        `json:"seriesId"`     //该奖品所属系列
	SeriName       string      `json:"seriesName"`   //该奖品所属系列名字
	Pic            string      `json:"pic"`          //奖品图片
	PkgStatus      int         `json:"pkgStatus"`    //品相状态
	SingleOrMuti   int         `json:"singleOrMuti"` //单一商品填1, 有n个组合就写n
	MultiIds       db.GormList `json:"multiIds"`     //商品id组合,单一
}
type ReqEnterFan struct {
	FanId uint `json:"fanId" binding:"required"`
}
type ReqQueryPrizePostion struct {
	FanId          uint   `json:"fanId" binding:"required"`
	PrizeIndexName string `json:"prizeIndexName" binding:"required"` //  first赏 last赏  全局赏  A赏  B赏...
	PrizeName      string `json:"prizeName" binding:"required"`
	TimeRange      string `json:"timeRange" binding:"required"`
	Status         int    `json:"status" binding:"required"`
}
type ReqModifyGoodsPosition struct {
	FanId          uint   `json:"fanId" binding:"required"`
	BoxId          uint   `json:"boxId" binding:"required"`
	PrizeIndexName string `json:"prizeIndexName" binding:"required"`
	PrizeName      string `json:"prizeName" binding:"required"`
	TimeRange      string `json:"timeRange" binding:"required"`
	Status         int    `json:"status" binding:"required"`
}
type ReqBuy struct {
	FanId uint `json:"fanId" binding:"required"`
	BoxId uint `json:"boxId" binding:"required"`
	Times int  `json:"times" binding:"required"`
}

type ReqBuySure struct {
	FanId   uint   `json:"fanId" binding:"required"`
	FanName string `json:"fanName" binding:"required"`
	BoxId   uint   `json:"boxId" binding:"required"`
	Times   int    `json:"times" binding:"required"`
}

type ReqBuyQuery struct {
	FanId uint `json:"fanId" binding:"required"`
	BoxId uint `json:"boxId" binding:"required"`
}

type ReqIsNew struct {
	Code string `json:"code"`
}
type ReqGetOpenId struct {
	EncryptedData string `json:"encryptedData"`
	Code          string `json:"code"`
	Iv            string `json:"iv"`
}

type ReqDelever struct {
	PageSize  int `json:"pageSize"`
	PageIndex int `json:"pageIndex"`
}
type ReqDeleverCondition struct {
	PageSize      int     `json:"pageSize"`
	PageIndex     int     `json:"pageIndex"`
	DeleverStatus int     `json:"deleverStatus"`
	OrderId       uint    `json:"orderId"`
	Mobile        string  `json:"mobile"`
	UserId        uint    `json:"userId"`
	GoodId        uint    `json:"goodId"`
	DeleOrderId   uint    `json:"deleOrderId"`
	TimeRange     []int64 `json:"timeRange"`
}
type ReqDeleverDetail struct {
	Id uint `json:"id"`
}

type ReqSetDelId struct {
	Id            uint   `json:"id"`
	DeleOrderId   uint   `json:"deleOrderId"`
	DeleCompany   string `json:"deleCompany"`
	DeleverStatus int    `json:"deleverStatus"`
}
type ReqUserList struct {
	//User      uint  `json:"user"`
	PageSize  int32 `json:"pageSize"`
	PageIndex int32 `json:"pageIndex"`
}
type ReqUserListCondition struct {
	UserId    uint   `json:"userId"`
	Mobile    string `json:"mobile"`
	NickName  string `json:"nickName"`
	PageSize  int    `json:"pageSize"`
	PageIndex int    `json:"pageIndex"`
}
type ReqAddRemark struct {
	OrderId string
	Remark  string
}
type ReqPageOfOrder struct {
	PageSize  int32 `json:"pageSize"`
	PageIndex int32 `json:"pageIndex"`
}

type ReqPageOfOrderCondition struct {
	PageSize    int32  `json:"pageSize"`
	PageIndex   int32  `json:"pageIndex"`
	OrderId     string `json:"outTradeNo"`
	Mobile      string `json:"mobile"`
	UserId      uint   `json:"userId"`
	OrderStatus string `json:"orderStatus"`
	PayStyle    string `json:"payStyle"`
}

type ReqPageOfOrderDetail struct {
	OrderId string `json:"outTradeNo"`
}

type ReqActiveByMan struct {
	TabName    string `json:"tabName"`
	TabNameSon string `json:"tabNameSon"`
}
type ReqSingleClick struct {
}
type ReqDelBannerPic struct {
	Id uint `json:"id"`
}
type ReqSetBannerPic struct {
	Banners []Banner `json:"banners"`
}
type Banner struct {
	Id              uint   `json:"id"`
	Title           string `json:"title"`
	FanId           uint   `json:"fanId"`
	FanTitle        string `json:"fanTitle"`
	AdvPic          string `json:"advPic"`
	Remark          string `json:"remark"`
	RedirectUrl     string `json:"redirectUrl"`
	RedirectType    string `json:"redirectType"`
	IsHide          bool   `json:"isHide"`
	ActiveBeginTime int64  `json:"activeBeginTime"`
	ActiveEndTime   int64  `json:"activeEndTime"`
}
type ReqGetBannerPic struct {
}
type ReqAddSecondTab struct {
	TabTag   string `json:"tabTag"`
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
type ReqAddSecondTabSon struct {
	TabTag          string `json:"tabTag"`
	TabSon          string `json:"tabSon"`
	ActiveBeginTime int64  `json:"activeBeginTime"`
	ActiveEndTime   int64  `json:"activeEndTime"`
	Title           string `json:"title"`
	FanId           uint   `json:"fanId"`
	FanTitle        string `json:"fanTitle"`
	IsHide          bool   `json:"isHide"`
	AdvPic          string `json:"advPic"`
	Remark          string `json:"remark"`
	RedirectUrl     string `json:"redirectUrl"`
	RedirectType    string `json:"redirectType"`
}

type ReqQuerySecondTab struct {
	TabTag string `json:"tabTag"`
}

type ReqQuerySecondSonTab struct {
	TabTag string `json:"tabTag"`
	TabSon string `json:"tabSon"`
}
type ReqShowOrHideBanner struct {
	BannerId int  `json:"bannerId"`
	IsHide   bool `json:"isHide"`
}
type ReqShowOrHideSecondTab struct {
	TabTag   string `json:"tabTag"`
	IsHide1  int    `json:"isHide1"`
	IsHide2  int    `json:"isHide2"`
	IsHide3  int    `json:"isHide3"`
	IsHide4  int    `json:"isHide4"`
	IsHide5  int    `json:"isHide5"`
	IsHide6  int    `json:"isHide6"`
	IsHide7  int    `json:"isHide7"`
	IsHide8  int    `json:"isHide8"`
	IsHide9  int    `json:"isHide9"`
	IsHide10 int    `json:"isHide10"`
}
type ReqShowOrHideSecondTabSon struct {
	TabTag string `json:"tabTag"`
	TabSon string `json:"tabSon"`
	Id     int    `json:"id"`
	IsHide bool   `json:"isHide"`
}
type ReqDeleteTabSon struct {
	Id int `json:"id"`
}
type ReqModifyAndSaveSecondTab struct {
}
