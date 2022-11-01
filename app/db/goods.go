package db

import (
	"gorm.io/gorm"
	"time"
)

type Fan struct {
	ID              uint   `gorm:"primarykey" json:"id"`
	Title           string `gorm:"unique;comment:蕃名" json:"title"`
	FanIndex        int    `json:"fanIndex"`
	Status          int    `gorm:"comment:是否上架1.手动上架 2.自动上架 3.未上架 4.手动下架 5.自动下架" json:"status"`
	Boxs            []*Box
	Price           float32        `gorm:"comment:蕃价格" json:"price"`
	SharePic        string         `gorm:"comment:共享图;type:varchar(128);not null" json:"sharePic"`
	DetailPic       string         `gorm:"comment:详细图;type:varchar(128);not null" json:"detailPic"`
	Rule            string         `gorm:"comment:规则;type:varchar(128);not null" json:"rule"`
	ActiveBeginTime int64          `gorm:"comment:活动开始时间" json:"activeBeginTime"`
	ActiveEndTime   int64          `gorm:"comment:活动结束时间" json:"activeEndTime"`
	CreatedAt       time.Time      `json:"created_time"`
	UpdatedAt       time.Time      `json:"updated_time"`
	WhoUpdate       string         `gorm:"comment:更新人" json:"whoUpdate"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}
type Ip struct {
	ID         uint   `gorm:"primarykey" json:"id"`
	Name       string `gorm:"comment:ip名" json:"name"`
	CreateName string `gorm:"comment:创建人" json:"createName"`
	CreateTime int64  `gorm:"comment:创建时间" json:"createTime"`
	Series     []*Series
	CreatedAt  time.Time      `json:"created_time"`
	UpdatedAt  time.Time      `json:"updated_time"`
	WhoUpdate  string         `gorm:"comment:更新人" json:"whoUpdate,omitempty"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}
type Series struct {
	ID         uint   `gorm:"primarykey" json:"id"`
	Name       string `gorm:"comment:系列名" json:"name"`
	CreateName string `gorm:"comment:创建人" json:"createName"`
	IpID       *uint
	IpName     string
	Goods      []*Goods
	CreatedAt  time.Time      `json:"created_time"`
	UpdatedAt  time.Time      `json:"updated_time"`
	WhoUpdate  string         `gorm:"comment:更新人" json:"whoUpdate,omitempty"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}
type Goods struct {
	ID           uint           `gorm:"primarykey" json:"id"`                             // 商品 id
	IpID         *uint          `gorm:"comment:所属ip的id" json:"ipId"`                      // IP的id
	IpName       string         `gorm:"comment:ip名字" json:"ipName"`                       //  IP名字
	SeriesID     *uint          `gorm:"comment:所属系列id" json:"seriesId"`                   //  系列id
	SeriesName   string         `gorm:"comment:所属系列id" json:"seriesName"`                 //  系列名字
	Name         string         `gorm:"comment:所属系列id" json:"name"`                       // 商品名
	PkgStatus    int8           `gorm:"comment:打包状态" json:"pkgStatus"`                    // 打包状态
	PreStore     string         `gorm:"comment:1. 预售  2.现货" json:"preStore"`              // 预存
	Integral     int32          `gorm:"comment:积分" json:"integral"`                       //   积分
	Pic          string         `gorm:"comment:图片;type:varchar(128);not null" json:"pic"` // 图片
	Price        float64        `gorm:"comment:建议售价" json:"price"`                        // 价格
	SingleOrMuti int            `gorm:"comment:组合" json:"singleOrMuti"`                   //组合个数
	MultiIds     GormList       `gorm:"type:varchar(128);not null"`                       //组合内部商品id
	Introduce    string         `gorm:"comment:介绍" json:"introduce"`                      //介绍
	WhoUpdate    string         `gorm:"comment:更新人" json:"whoUpdate"`                     //谁更新了商品
	CreatedAt    time.Time      `json:"created_time"`
	UpdatedAt    time.Time      `json:"updated_time"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}
type Box struct {
	ID            uint    `gorm:"primarykey" json:"id"`
	Price         float64 `gorm:"comment:箱子价格" json:"price"` //售卖价格
	FanName       string  `gorm:"comment:该箱属于蕃名" json:"fanName"`
	BoxIndex      int32   `gorm:"comment:箱子在番中的次序(第3/10箱)" json:"boxIndex"`
	PriczeNum     int32   `gorm:"comment:箱子开始多少商品" json:"priczeNum"`
	PriczeLeftNum int32   `gorm:"comment:箱子还剩多少个商品(剩余300/400个)" json:"priczeLeftNum"`
	Status        int     `gorm:"comment:箱子状态(1.上架有奖品.2未上架.3上架无商品)" json:"status"`
	FanID         *uint
	Prizes        []*Prize
	CreatedAt     time.Time      `json:"created_time"`
	UpdatedAt     time.Time      `json:"updated_time"`
	WhoUpdate     string         `gorm:"comment:更新人" json:"whoUpdate,omitempty"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}

type Prize struct {
	ID                uint     `gorm:"primarykey" json:"id"`
	GoodID            uint     `gorm:"comment:真实商品id;index" json:"goodId"`
	GoodName          string   `gorm:"comment:商品名字" json:"goodName"`
	FanId             uint     `gorm:"comment:所属蕃的id;index" json:"fanId"`
	FanName           string   `gorm:"comment:所属蕃的名字" json:"fanName"`
	IpID              uint     `gorm:"comment:所属ip的id" json:"ipId"`
	IpName            string   `gorm:"comment:所属ip的名字" json:"ipName"`
	SeriesID          uint     `gorm:"comment:所属系列id" json:"seriesId"`
	SeriesName        string   `gorm:"comment:所属系列id"  json:"seriesName"`
	Position          GormList `gorm:"type:varchar(128);not null"`
	BoxID             *uint
	BoxIndex          int
	Pic               string         `gorm:"comment:图片;type:varchar(128);not null" json:"pic"`
	Price             float64        `gorm:"comment:建议售价" json:"price"`
	PrizeNum          int32          `gorm:"comment:该类奖品总数" json:"priczeNum"`
	PriczeLeftNum     int32          `gorm:"comment:该类奖品剩余(3/10)" json:"priczeLeftNum"`
	PrizeIndex        int32          `gorm:"comment:赏的次序" json:"prizeIndex"`
	PrizeIndexName    string         `gorm:"comment:A赏,B赏..." json:"prizeIndexName"`
	PrizeRate         string         `gorm:"comment:中奖率" json:"prizeRate"`
	Remark            string         `gorm:"comment:备注" json:"remark"`
	SingleOrMuti      int            `json:"singleOrMuti"`
	MultiIds          GormList       `gorm:"type:varchar(128);not null"`
	PreStore          int            `gorm:"comment:1.预售， 2.库存;index" json:"preStore"`
	SoldStatus        int            `gorm:"comment:是否售罄1.奖品售罄,2.奖品未售罄;index" json:"soldStatus"`
	TimeForSoldStatus string         `gorm:"comment:预售时间" json:"timeForSoldStatus,omitempty"`
	PkgStatus         int            `gorm:"comment:打包状态;index" json:"pkgStatus"`
	WhoUpdate         string         `gorm:"comment:更新人" json:"whoUpdate,omitempty"`
	CreatedAt         time.Time      `json:"created_time"`
	UpdatedAt         time.Time      `json:"updated_time"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"-"`
}
type Order struct {
	ID            uint           `gorm:"primarykey" json:"id"`
	OutTradeNo    string         `gorm:"comment:订单编号;index" json:"orderId"`
	PrizeNum      int            `gorm:"comment:奖品数量（抽奖次数，不包括特殊赏）" json:"prizeNum"`
	Price         float64        `gorm:"comment:金额" json:"price"`
	UserName      string         `gorm:"comment:用户名" json:"userName"`
	FanId         uint           `gorm:"comment:蕃id" json:"fanId"`
	FanName       string         `gorm:"comment:蕃名" json:"fanName"`
	BoxId         uint           `gorm:"comment:箱id" json:"boxId"`
	BoxIndex      int            `gorm:"comment:箱名顺序" json:"boxIndex"`
	OpenId        string         `gorm:"comment:用户openid " json:"openId"`
	UserId        uint           `gorm:"comment:用户id" json:"userId"`
	Avatar        string         `gorm:"comment:头像" json:"avatar"`
	FanPic        string         `gorm:"comment:蕃图片" json:"fanPic"`
	UserMobile    string         `gorm:"comment:用户手机号" json:"userMobile"`
	TransactionId string         `gorm:"comment:交易流水号" json:"transactionId"`
	Appid         string         `gorm:"comment:appid" json:"appid"`
	FinishTime    int64          `gorm:"comment:完成时间" json:"finishTime"`
	OrderType     string         `gorm:"comment:订单类型" json:"orderType"`
	PayStyle      string         `gorm:"comment:支付方式 " json:"payStyle"`
	Status        string         `gorm:"comment:订单状态" json:"Status"`
	Remark        string         `gorm:"comment:备注" json:"remark"`
	Detail        string         `gorm:"comment:详情" json:"detail"`
	Operator      string         `gorm:"comment:操作" json:"operator"`
	CreatedAt     time.Time      `json:"created_time"`
	UpdatedAt     time.Time      `json:"updated_time"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}

type OrderDeliver struct {
	ID             uint           `gorm:"primarykey;comment:提货单号;" json:"id"`
	DeleCompany    string         `gorm:"comment:快递公司地址;index" json:"deleCompany"`
	DeleOrderId    uint           `gorm:"comment:快递单号;index" json:"deleOrderId"`
	ActiveSureTime int64          `gorm:"comment:自动确认时间;index" json:"activeSureTime"`
	AddressId      uint           `gorm:"comment:地址id;index" json:"addressId"`
	OutTradeNo     string         `gorm:"comment:订单编号;index" json:"orderId"`
	PrizeNum       int            `gorm:"comment:奖品数量（抽奖次数，不包括特殊赏）" json:"prizeNum"`
	Price          float32        `gorm:"comment:金额" json:"price"`
	UserName       string         `gorm:"comment:用户名" json:"userName"`
	OpenId         string         `gorm:"comment:用户openid " json:"openId"`
	UserId         uint           `gorm:"comment:用户id;index" json:"userId"`
	Avatar         string         `gorm:"comment:头像" json:"avatar"`
	UserMobile     string         `gorm:"comment:用户手机号" json:"userMobile"`
	TransactionId  string         `gorm:"comment:交易流水号" json:"transactionId"`
	Appid          string         `gorm:"comment:appid" json:"appid"`
	FinishTime     int64          `gorm:"comment:完成时间" json:"finishTime"`
	OrderType      string         `gorm:"comment:订单类型" json:"orderType"`
	PayStyle       string         `gorm:"comment:支付方式 " json:"payStyle"`
	DeleverStatus  int            `gorm:"comment:1.待提货,2.待发货,3.已经收货;index" json:"deleverStatus"`
	Status         string         `gorm:"comment:支付状态;index" json:"status"`
	Remark         string         `gorm:"comment:备注" json:"remark"`
	Detail         string         `gorm:"comment:详情" json:"detail"`
	Operator       string         `gorm:"comment:操作" json:"operator"`
	CreatedAt      time.Time      `json:"created_time"`
	UpdatedAt      time.Time      `json:"updated_time"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
}
type OrderDeliverDetail struct {
	ID             uint           `gorm:"primarykey" json:"id"`
	AddressId      uint           `gorm:"comment:地址id;index" json:"addressId"`
	LuggageId      uint           `gorm:"comment:提货柜商品id;index" json:"luggageId"`
	ActiveSureTime int64          `gorm:"comment:自动确认时间;index" json:"activeSureTime"`
	OutTradeNo     string         `gorm:"comment:订单编号;index" json:"orderId"`
	UserName       string         `gorm:"comment:用户名" json:"userName"`
	FanId          uint           `gorm:"comment:蕃id" json:"fanId"`
	GoodId         uint           `gorm:"comment:商品id" json:"goodId"`
	GoodName       string         `gorm:"comment:商品名" json:"goodName"`
	IpName         string         `gorm:"comment:IP名称" json:"ipName"`
	SeriesName     string         `gorm:"comment:系列名称" json:"seriesName"`
	PkgStatus      int            `gorm:"comment:打包状态" json:"pkgStatus"`
	PrizeIndexName string         `gorm:"comment:商品等级名" json:"prizeIndexName"`
	FanName        string         `gorm:"comment:蕃名" json:"fanName"`
	BoxId          uint           `gorm:"comment:箱id" json:"boxId"`
	BoxIndex       int            `gorm:"comment:箱名顺序" json:"boxIndex"`
	OpenId         string         `gorm:"comment:用户openid " json:"openId"`
	UserId         uint           `gorm:"comment:用户id;index" json:"userId"`
	Avatar         string         `gorm:"comment:头像" json:"avatar"`
	Pic            string         `gorm:"comment:图片" json:"pic"`
	UserMobile     string         `gorm:"comment:用户手机号" json:"userMobile"`
	TransactionId  string         `gorm:"comment:交易流水号" json:"transactionId"`
	Appid          string         `gorm:"comment:appid" json:"appid"`
	FinishTime     int64          `gorm:"comment:完成时间" json:"finishTime"`
	OrderType      string         `gorm:"comment:订单类型" json:"orderType"`
	PayStyle       string         `gorm:"comment:支付方式 " json:"payStyle"`
	Price          float64        `gorm:"comment:价格 " json:"price"`
	DeleverStatus  int            `gorm:"comment:1.待提货,2.待发货,3.已经收货;index" json:"deleverStatus"`
	Status         string         `gorm:"comment:支付状态;index" json:"status"`
	Remark         string         `gorm:"comment:备注" json:"remark"`
	Detail         string         `gorm:"comment:详情" json:"detail"`
	Operator       string         `gorm:"comment:操作" json:"operator"`
	CreatedAt      time.Time      `json:"created_time"`
	UpdatedAt      time.Time      `json:"updated_time"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
}
type User struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	NickName    string         `gorm:"comment:昵称" json:"nickName"` //昵称
	Mobile      string         `gorm:"comment:手机号" json:"mobile"`  //手机号
	AppId       string         `gorm:"comment:appId" json:"appId"`
	OpenId      string         `gorm:"comment:openId;unique;" json:"openId"`
	SessionKey  string         `gorm:"comment:sessionKey" json:"sessionKey"`
	RealName    string         `gorm:"comment:真实姓名" json:"realName"` //真实姓名
	LoginSalt   string         `gorm:"comment:登录盐" json:"loginSalt"` //
	LoginHash   string         `gorm:"comment:登录hash" json:"loginHash"`
	WxCode      string         `gorm:"comment:微信号" json:"wxCode"`            // 用户微信号
	Gender      int            `gorm:"comment:性别" json:"gender"`             // 性别
	AvatarUrl   string         `gorm:"comment:头像" json:"avatarUrl"`          // 头像
	WxAppOpenid string         `gorm:"comment:小程序唯一code" json:"wxAppOpenid"` // 小程序唯一code
	WxPubOpenid string         `gorm:"comment:公众号唯一id"  json:"wxPubOpenid"`  // 微信公众号openid
	Unionid     string         `gorm:"comment:开发平台id" json:"unionid"`        // 微信开放平台唯一id
	CountryCode string         `gorm:"comment:区号" json:"countryCode"`        // 区号
	InviteId    int            `gorm:"comment:邀请人" json:"inviteId"`          // 邀请人id(父级)
	InviteCode  string         `gorm:"comment:邀请码" json:"inviteCode"`        // 邀请码
	InviteTime  int64          `gorm:"comment:邀请时间" json:"inviteTime"`       // 邀请时间
	CreatedAt   time.Time      `json:"created_time"`
	UpdatedAt   time.Time      `json:"updated_time"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
type Luggage struct {
	ID                uint           `gorm:"primarykey" json:"id"`
	OutTradeNo        string         `gorm:"comment:番id" json:"outTradeNo"`
	UserId            uint           `gorm:"index;comment:用户id" json:"userId"`
	UserName          string         `gorm:"comment:用户昵称" json:"userName"`
	GoodID            uint           `gorm:"comment:真实商品id;index" json:"goodId"`
	GoodName          string         `gorm:"comment:商品名字" json:"goodName"`
	Avatar            string         `gorm:"comment:头像" json:"avatar"`
	FanId             uint           `gorm:"comment:所属蕃的id;index" json:"fanId"`
	FanName           string         `gorm:"comment:所属蕃的名字" json:"fanName"`
	BoxId             uint           `gorm:"comment:箱子Id;index" json:"boxId"`
	BoxIndex          int            `gorm:"comment:箱子序号;index" json:"boxIndex"`
	IpID              uint           `gorm:"comment:所属ip的id" json:"ipId"`
	IpName            string         `gorm:"comment:所属ip的名字" json:"ipName"`
	Position          GormList       `gorm:"comment:位置;type:varchar(128)" json:"position"`
	SeriesID          uint           `gorm:"comment:所属系列id" json:"seriesId"`
	SeriesName        string         `gorm:"comment:所属系列id"  json:"seriesName"`
	Payments          float64        `gorm:"comment:支付额度" json:"payments"`
	Pic               string         `gorm:"comment:图片;type:varchar(128);not null" json:"pic"`
	Price             float64        `gorm:"comment:建议售价" json:"price"`
	PrizeIndexName    string         `gorm:"comment:A赏,B赏..." json:"prizeIndexName"`
	PrizeIndex        int            `gorm:"comment:A,B的index..." json:"prizeIndex"`
	Remark            string         `gorm:"comment:备注" json:"remark"`
	SingleOrMuti      int            `json:"singleOrMuti"`
	MultiIds          GormList       `gorm:"type:varchar(128);not null"`
	PreStore          int            `gorm:"comment:1.预售 2.现货" json:"preStore"`
	DeleverStatus     int            `gorm:"index;comment:1.待提货,2.待发货,3.已经收货" json:"deleverStatus"`
	TimeForSoldStatus string         `gorm:"comment:预售时间" json:"timeForSoldStatus,omitempty"`
	PkgStatus         int            `gorm:"comment:打包状态" json:"pkgStatus"`
	CreatedAt         time.Time      `json:"created_time"`
	UpdatedAt         time.Time      `json:"updated_time"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"-"`
}

type Sure struct {
	ID         uint           `gorm:"primarykey" json:"id"`
	FanId      uint           `gorm:"comment:番的id" json:"fanId"`
	FanTitle   string         `gorm:"comment:番的标题;"  json:"fanTitle"`
	BoxId      uint           `gorm:"comment:箱子的id;" json:"boxId"`
	PrizeIndex GormList       `gorm:"type:varchar(5120);not null"`
	CreatedAt  time.Time      `json:"created_time"`
	UpdatedAt  time.Time      `json:"updated_time"`
	WhoUpdate  string         `gorm:"comment:更新人" json:"whoUpdate,omitempty"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}
type Left struct {
	ID         uint           `gorm:"primarykey" json:"id"`
	FanId      uint           `gorm:"comment:番的id;" json:"fanId"`
	FanTitle   string         `gorm:"comment:番的标题;"  json:"fanTitle"`
	BoxId      uint           `gorm:"comment:箱子的id;" json:"boxId"`
	PrizeIndex GormList       `gorm:"type:varchar(5120);not null"`
	CreatedAt  time.Time      `json:"created_time"`
	UpdatedAt  time.Time      `json:"updated_time"`
	WhoUpdate  string         `gorm:"comment:更新人" json:"whoUpdate,omitempty"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}
type Target struct {
	ID         uint           `gorm:"primarykey" json:"id"`
	FanId      uint           `gorm:"comment:番的id;" json:"fanId"`
	FanTitle   string         `gorm:"comment:番的标题;"  json:"fanTitle"`
	BoxId      uint           `gorm:"comment:箱子的id;" json:"boxId"`
	PrizeIndex GormList       `gorm:"type:varchar(5120);not null"`
	CreatedAt  time.Time      `json:"created_time"`
	UpdatedAt  time.Time      `json:"updated_time"`
	WhoUpdate  string         `gorm:"comment:更新人" json:"whoUpdate,omitempty"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}
type FirstPrize struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	FanId       uint           `gorm:"comment:番的id;uniqueIndex:udx_name" json:"fanId"`
	FanTitle    string         `gorm:"comment:番的标题;"  json:"fanTitle"`
	BoxId       uint           `gorm:"comment:箱子的id;uniqueIndex:udx_name" json:"boxId"`
	Pos         GormList       `gorm:"comment:位置;type:varchar(64)" json:"pos"`
	PrizeIndexs GormList       `gorm:"comment:商品顺序;type:varchar(64);" json:"prizeIndexs"`
	CreatedAt   time.Time      `json:"created_time"`
	UpdatedAt   time.Time      `json:"updated_time"`
	WhoUpdate   string         `gorm:"comment:更新人" json:"whoUpdate,omitempty"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
type LastPrize struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	FanId       uint           `gorm:"comment:番的id;uniqueIndex:udx_name" json:"fanId"`
	FanTitle    string         `gorm:"comment:番的标题;"  json:"fanTitle"`
	BoxId       uint           `gorm:"comment:箱子的id;uniqueIndex:udx_name" json:"boxId"`
	Pos         GormList       `gorm:"comment:位置;type:varchar(64)" json:"pos"`
	PrizeIndexs GormList       `gorm:"comment:商品顺序;type:varchar(64);" json:"prizeIndexs"`
	CreatedAt   time.Time      `json:"created_time"`
	UpdatedAt   time.Time      `json:"updated_time"`
	WhoUpdate   string         `gorm:"comment:更新人" json:"whoUpdate,omitempty"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
type GlobalPrize struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	FanId       uint           `gorm:"comment:番的id;uniqueIndex:udx_name" json:"fanId"`
	FanTitle    string         `gorm:"comment:番的标题;"  json:"fanTitle"`
	BoxId       uint           `gorm:"comment:箱子的id;uniqueIndex:udx_name" json:"boxId"`
	Pos         GormList       `gorm:"comment:位置;type:varchar(64)" json:"pos"`
	PrizeIndexs GormList       `gorm:"comment:商品顺序;type:varchar(64);" json:"prizeIndexs"`
	CreatedAt   time.Time      `json:"created_time"`
	UpdatedAt   time.Time      `json:"updated_time"`
	WhoUpdate   string         `gorm:"comment:更新人" json:"whoUpdate,omitempty"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

type SpecialPrizeRecord struct {
	ID                uint           `gorm:"primarykey" json:"id"`
	FanId             uint           `gorm:"comment:番的id" json:"fanId"`
	FanTitle          string         `gorm:"comment:番的标题"  json:"fanTitle"`
	BoxId             uint           `gorm:"comment:箱子的id" json:"boxId"`
	Times             int            `gorm:"comment:抽奖次数记录" json:"times"`
	FirstPrizeRecord  GormList       `gorm:"comment:奖品序号;type:varchar(128)" json:"firstPrizeRecord"`
	LastPrizeRecord   GormList       `gorm:"comment:奖品序号;type:varchar(128)" json:"lastPrizeRecord""`
	GlobalPrizeRecord GormList       `gorm:"comment:奖品序号;type:varchar(128)" json:"globalPrizeRecord"`
	CreatedAt         time.Time      `json:"created_time"`
	UpdatedAt         time.Time      `json:"updated_time"`
	WhoUpdate         string         `gorm:"comment:更新人" json:"whoUpdate,omitempty"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"-"`
}

type Integral struct {
	UserId    uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_time"`
	UpdatedAt time.Time      `json:"updated_time"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type IntegralRecord struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	UserId    uint           `gorm:"comment:用户id" json:"user_id"`
	CreatedAt time.Time      `json:"created_time"`
	UpdatedAt time.Time      `json:"updated_time"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type Adver struct {
	ID              uint           `gorm:"primarykey" json:"id"`
	Title           string         `gorm:"comment:title" json:"title"`
	AdvPic          string         `gorm:"comment:图片" json:"advPic"`
	FanId           uint           `gorm:"comment:蕃id" json:"fanId"`
	FanTitle        string         `gorm:"comment:蕃的名字" json:"fanTitle"`
	RedirectUrl     string         `gorm:"comment:轮播图对应跳转连接" json:"redirectUrl"`
	RedirectType    string         `gorm:"comment:轮播图对应跳转连接类型" json:"redirectType"`
	TipsAfterBanner string         `gorm:"comment:轮播图下面的文字提示" json:"tips_afterBanner"`
	ActiveBeginTime int64          `gorm:"comment:轮播图上架时间" json:"activeBeginTime"`
	ActiveEndTime   int64          `gorm:"comment:轮播图下架时间" json:"activeEndTime"`
	IsHide          bool           `gorm:"comment:是否隐藏" json:"is_hide"`
	Remark          string         `gorm:"comment:备注" json:"remark"`
	CreatedAt       time.Time      `json:"created_time"`
	UpdatedAt       time.Time      `json:"updated_time"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}
type AdverTab struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	TabTag    string         `gorm:"comment:推荐;index" json:"tabTag"`
	TabSon1   string         `gorm:"comment:son1;" json:"tabSon1"`
	TabSon2   string         `gorm:"comment:son2;" json:"tabSon2"`
	TabSon3   string         `gorm:"comment:son3;" json:"tabSon3"`
	TabSon4   string         `gorm:"comment:son4;" json:"tabSon4"`
	TabSon5   string         `gorm:"comment:son5;" json:"tabSon5"`
	TabSon6   string         `gorm:"comment:son6;" json:"tabSon6"`
	TabSon7   string         `gorm:"comment:son7;" json:"tabSon7"`
	TabSon8   string         `gorm:"comment:son8;" json:"tabSon8"`
	TabSon9   string         `gorm:"comment:son9;" json:"tabSon9"`
	TabSon10  string         `gorm:"comment:son10;" json:"tabSon10"`
	IsHide1   bool           `gorm:"comment:hide1;" json:"isHide1"`
	IsHide2   bool           `gorm:"comment:hide2;" json:"isHide2"`
	IsHide3   bool           `gorm:"comment:hide3;" json:"isHide3"`
	IsHide4   bool           `gorm:"comment:hide4;" json:"isHide4"`
	IsHide5   bool           `gorm:"comment:hide5;" json:"isHide5"`
	IsHide6   bool           `gorm:"comment:hide6;" json:"isHide6"`
	IsHide7   bool           `gorm:"comment:hide7;" json:"isHide7"`
	IsHide8   bool           `gorm:"comment:hide8;" json:"isHide8"`
	IsHide9   bool           `gorm:"comment:hide9;" json:"isHide9"`
	IsHide10  bool           `gorm:"comment:hide10;" json:"isHide10"`
	CreatedAt time.Time      `json:"created_time"`
	UpdatedAt time.Time      `json:"updated_time"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
type AdverContent struct {
	ID              uint           `gorm:"primarykey" json:"id"`
	TabTag          string         `gorm:"comment:推荐;index" json:"tabTag"`
	TabSon          string         `gorm:"comment:推荐下面的二级tab;type:varchar(128);not null" json:"tabSon"`
	ActiveBeginTime int64          `gorm:"comment:广告开始时间;" json:"activeBeginTime"`
	ActiveEndTime   int64          `gorm:"comment:广告结束时间;" json:"activeEndTime"`
	Title           string         `gorm:"comment:广告标题;type:varchar(128);not null" json:"title"`
	FanId           uint           `gorm:"comment:蕃id;type:varchar(128);not null" json:"fanId"`
	FanTitle        string         `gorm:"comment:蕃的标题;type:varchar(128);not null" json:"fanTitle"`
	AdvPic          string         `gorm:"comment:广告图片;type:varchar(128);not null" json:"advPic"`
	Remark          string         `gorm:"comment:备注;type:varchar(128);not null" json:"remark"`
	IsHide          bool           `gorm:"comment:是否隐藏;" json:"isHide"`
	RedirectUrl     string         `gorm:"comment:H5跳转;type:varchar(128);not null" json:"redirectUrl"`
	RedirectType    string         `gorm:"comment:广告类型;type:varchar(128);not null" json:"redirectType"`
	CreatedAt       time.Time      `json:"created_time"`
	UpdatedAt       time.Time      `json:"updated_time"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}
type Address struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	UserName  string         `gorm:"comment:用户" json:"userName"`
	Post      string         `gorm:"comment:邮编" json:"post"`
	UserId    uint           `gorm:"comment:用户;index" json:"userId"`
	Mobile    string         `gorm:"comment:手机号" json:"mobile"`
	Province  string         `gorm:"comment:省份" json:"province"`
	City      string         `gorm:"comment:城市" json:"city"`
	Country   string         `gorm:"comment:县区" json:"country"`
	Detail    string         `gorm:"comment:详细地址;unique" json:"detail"`
	IsDefault bool           `gorm:"comment:是否默认" json:"isDefault"`
	CreatedAt time.Time      `json:"created_time"`
	UpdatedAt time.Time      `json:"updated_time"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
type Config struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	Fee         float64        `gorm:"comment:运费" json:"fee"`
	FeeLimitNum int            `gorm:"comment:运费数量限制" json:"feeLimitNum"`
	CreatedAt   time.Time      `json:"created_time"`
	UpdatedAt   time.Time      `json:"updated_time"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

//type Urls struct {
//	ID        string `gorm:"primarykey" json:"id"`
//	FanId     uint
//	FanName   string
//	FanUrl    string
//	BoxId     uint
//	GoodsId   uint
//	GoodsName string
//	GoodsUrl  string
//	CreatedAt time.Time      `json:"created_time"`
//	UpdatedAt time.Time      `json:"updated_time"`
//	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
//}
