package param

import (
	"yifan/app/db"
)

type ReqGetUser struct {
	Id uint `uri:"id" binding:"required" example:"123456789" ` //用户的id()
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
	PageSize  int32 `uri:"pageSize"`
	PageIndex int32 `uri:"pageIndex"`
	StyleId   uint  `uri:"styleId"`
	PkgStatus int8  `uri:"pkgStatus"`
}

type ReqIsNeedToJoinQueue struct {
}

//
type ReqUpLoadGoods struct {
	UpLoadGoods []UpLoadGoods `json:"goods"`
}
type UpLoadGoods struct {
	GoodsName       string  `json:"goodsName" binding:"required"` //商品名
	IpName          string  `json:"ipName" binding:"required"`
	SeriesName      string  `json:"seriesName" binding:"required"`
	Pic             string  `json:"pic" binding:"required"`             //图片
	Price           float64 `json:"price" binding:"required"`           //售价
	SingleOrMuti    int     `json:"singleOrMuti"`                       //单个还是多个
	MultiIds        []int   `json:"multiIds" `                          //多个商品id
	PkgStatus       int8    `json:"pkgStatus" binding:"required"`       //打包状态
	Introduce       string  `json:"introduce" binding:"required"`       //商品简介
	Integral        int32   `json:"integral" binding:"required"`        //积分
	Status          string  `json:"status" binding:"required"`          //
	ActiveBeginTime int64   `json:"activeBeginTime" binding:"required"` //活动开始时间
	ActiveEndTime   int64   `json:"activeEndTime" binding:"required""`
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
	Status       string  `json:"status" binding:"required"`    //现货状态
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
	Status          int         `json:"status"`
	WhoUpdate       string      `json:"whoUpdate"`
	ActiveBeginTime int64       `json:"activeBeginTime"`
	ActiveEndTime   int64       `json:"activeEndTime"`
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
	ReqAddSeries []ReqAddSeries `json:"reqAddSeries"`
}
type ReqAddSeries struct {
	Name       string `json:"name" binding:"required,lte=32,gte=1"`
	CreateName string `json:"createName" binding:"required,lte=32,gte=1"` //创建人
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
	GoodId         uint        `json:"goodId"`         //奖品id
	GoodName       string      `json:"goodName"`       //奖品名
	PrizeNum       int32       `json:"prizeNum"`       //某一个种类奖品数量
	PrizeIndex     int32       `json:"prizeIndex"`     //某一个种类奖品在箱子中的序号
	PrizeIndexName string      `json:"prizeIndexName"` //A赏,B赏
	Position       []int       `json:"position"`       //A赏,B赏 位置
	PrizeStyle     string      `json:"prizeStyle"`     //抽取方式
	Price          float64     `json:"price"`          //某种商品价格
	IpId           uint        `json:"ipId"`           //该奖品所属IP
	IpName         string      `json:"ipName"`         //该奖品所属IP的名字
	Remark         string      `json:"remark"`
	SeriId         uint        `json:"seriesId"`     //该奖品所属系列
	SeriName       string      `json:"seriesName"`   //该奖品所属系列名字
	Pic            string      `json:"pic"`          //奖品图片
	PkgStatus      int         `json:"pkgStatus"`    //品相状态
	SingleOrMuti   int         `json:"singleOrMuti"` //单一商品填1, 有n个组合就写n
	MultiIds       db.GormList `json:"multiIds"`     //商品id组合,单一商品[435], n个商品[34,456,234,...]
}
type ReqPageOfPosition struct {
}
type ReqPageOfPositionCondition struct {
	FanId          uint    `json:"fanId"`
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
	OldPrizeIndex     int    `json:"oldPrizeIndex"`
	NewGoodId         uint   `json:"newGoodId"`         //商品id
	NewGoodName       string `json:"newGoodName"`       //商品名
	NewPrizeNum       int32  `json:"newPrizeNum"`       //某一个种类奖品数量
	NewPrizeIndex     int32  `json:"newPrizeIndex"`     //某一个种类奖品在箱子中的序号
	NewPrizeIndexName string `json:"newPrizeIndexName"` //A赏,B赏
	NewPrizePosition  string `json:"newPrizePosition"`  //A赏,B赏 位置
	NewPrizeStyle     string `json:"newPrizeStyle"`     //抽取方式
	NewPkgStatus      int    `json:"newPkgStatus"`      //品相状态
	NewSingleOrMuti   int    `json:"newSingleOrMuti"`   //单一商品填1, 有n个组合就写n
	NewMultiIds       string `json:"newMultiIds"`       //商品id组合,单一商品[435], n个商品[34,456,234,...]
}
type ReqDeleteBoxGoods struct {
	FanId  uint `json:"fanId"`
	GoodId int  `json:"goodId"`
}
type ReqAddFan struct {
	FanName       string  `json:"fanName"`       //蕃的名字
	FanPrice      float64 `json:"fanPrice"`      //蕃的价格
	Status        int     `json:"status"`        //蕃的状态  1.上架  2.下架
	OnActiveTime  int64   `json:"onActiveTime"`  //活动开始时间
	OffActiveTime int64   `json:"offActiveTime"` //活动结束时间
	WhoCreated    string  `json:"whoCreated"`
	Pic           string  `json:"pic"` //蕃的图片
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
	FanID           uint     `json:"fanId"`
	FanName         string   `json:"fanName"`
	FanType         string   `json:"fanType"`
	Rule            string   `json:"rule"`
	Title           string   `json:"title"`
	FanPrice        float64  `json:"fanPrice"`
	Status          int      `json:"status"`
	ActiveBeginTime int64    `json:"activeBeginTime"`
	ActiveEndTime   int64    `json:"activeEndTime"`
	DetailPic       string   `json:"detailPic"`
	SharePic        string   `json:"sharePic"`
	WhoUpdate       string   `json:"whoUpdate"`
	TotalBoxNum     int      `json:"totalBoxNum"`
	Price           float64  `json:"price"`  //售卖价格
	Prizes          []PrizeX `json:"prizes"` //每个箱的所有奖品
}
type PrizeX struct {
	PrizeId        uint        `json:"prizeId"`   //奖品id
	PrizeName      string      `json:"prizeName"` //奖品名
	GoodId         uint        `json:"goodId"`
	GoodName       string      `json:"goodName"`
	PrizeNum       int32       `json:"prizeNum"`       //某一个种类奖品数量
	PrizeIndex     int32       `json:"prizeIndex"`     //某一个种类奖品在箱子中的序号
	PrizeIndexName string      `json:"prizeIndexName"` //A赏,B赏
	Position       []int       `json:"position"`       //A赏,B赏 位置
	PrizeStyle     string      `json:"prizeStyle"`     //抽取方式
	Price          float64     `json:"price"`          //某种商品价格
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
