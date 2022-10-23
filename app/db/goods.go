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
	ID              uint           `gorm:"primarykey" json:"id"`
	IpID            *uint          `gorm:"comment:所属ip的id" json:"ipId"`
	IpName          string         `gorm:"comment:ip名字" json:"ipName"`
	SeriesID        *uint          `gorm:"comment:所属系列id" json:"seriesId"`
	SeriesName      string         `json:"seriesName" json:"seriesName"`
	Pic             string         `gorm:"comment:图片;type:varchar(128);not null" json:"pic"`
	Price           float64        `gorm:"comment:建议售价" json:"price"`
	Name            string         `gorm:"comment:名字" json:"name"`
	SingleOrMuti    int            `json:"singleOrMuti"`
	MultiIds        GormList       `gorm:"type:varchar(128);not null"`
	PkgStatus       int8           `gorm:"comment:打包状态" json:"pkgStatus"`
	Introduce       string         `gorm:"comment:介绍" json:"introduce"`
	Integral        int32          `gorm:"comment:积分" json:"integral"`
	SoldStatus      string         `gorm:"comment:1.预售 2.现货" json:"soldStatus"`
	ActiveBeginTime int64          `gorm:"comment:活动开始时间" json:"activeBeginTime"`
	ActiveEndTime   int64          `gorm:"comment:活动结束时间" json:"activeEndTime"`
	CreatedAt       time.Time      `json:"created_time"`
	UpdatedAt       time.Time      `json:"updated_time"`
	WhoUpdate       string         `gorm:"comment:更新人" json:"whoUpdate,omitempty"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}
type Box struct {
	ID            uint    `gorm:"primarykey" json:"id"`
	Price         float64 `gorm:"comment:箱子价格" json:"price"` //售卖价格
	FanId         uint    `gorm:"comment:该箱属于蕃id" json:"fanId"`
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
	Status            int            `gorm:"comment:上下架" json:"status"`
	SoldStatus        int            `gorm:"comment:是否售罄1.奖品售罄,2.奖品未售罄" json:"soldStatus"`
	TimeForSoldStatus string         `gorm:"comment:预售时间" json:"timeForSoldStatus,omitempty"`
	PkgStatus         int            `gorm:"comment:打包状态" json:"pkgStatus"`
	WhoUpdate         string         `gorm:"comment:更新人" json:"whoUpdate,omitempty"`
	CreatedAt         time.Time      `json:"created_time"`
	UpdatedAt         time.Time      `json:"updated_time"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"-"`
}
type Order struct {
	ID              uint    `gorm:"primarykey" json:"id"`
	OutTradeNo      string  `gorm:"unique;comment:番id" json:"outTradeNo"`
	Times           int     `json:"times"`
	FanId           uint    `gorm:"comment:番id" json:"fanId"`
	FanName         string  `gorm:"comment:番名" json:"fanName"`
	BoxId           uint    `gorm:"comment:箱id" json:"boxId"`
	BoxIndex        int     `gorm:"comment:箱index" json:"boxIndex"`
	PrizeIndex      uint    `gorm:"comment:奖品id" json:"prizeIndex"`
	PrizeName       string  `gorm:"comment:奖品名字" json:"prizeName"`
	OpenId          string  `gorm:"comment:openid" json:"openId"`
	PrizeIndexName  string  `gorm:"comment:A赏,B赏..." json:"prizeIndexName"`
	UserId          uint    `gorm:"comment:用户id" json:"userId"`
	UserName        string  `gorm:"comment:用户名" json:"userName"`
	Avatar          string  `gorm:"comment:头像" json:"avatar"`
	Position        string  `gorm:"comment:特殊赏的位置" json:"position"`
	Pic             string  `gorm:"comment:图片" json:"pic"`
	PrizeTimes      int     `gorm:"comment:抽奖次数" json:"prizeTimes"`
	UserMobile      string  `gorm:"comment:用户手机号" json:"userMobile"`
	PrepayId        *string `gorm:"index;comment:预支付交易会话标识" json:"prepayId"`
	Appid           *string `gorm:"comment:应用ID" json:"appid"`
	TimeStamp       *string `gorm:"comment:时间戳" json:"timeStamp"`
	NonceStr        *string `gorm:"comment:随机字符串" json:"nonceStr"`
	Package         *string `gorm:"comment:订单详情扩展字符串" json:"package"`
	SignType        *string `gorm:"comment:签名方式" json:"signType"`
	PaySign         *string `gorm:"comment:签名" json:"paySign"`
	Payments        float64 `gorm:"comment:支付额度" json:"payments"`
	OrderRemark     string  `gorm:"comment:订单备注" json:"orderRemark"`
	OrderType       string  `gorm:"comment:订单类型" json:"orderType"`
	PayStyle        string  `gorm:"comment:支付类型" json:"payStyle"`
	CreateTime      int64   `gorm:"comment:创建时间" json:"createTime"`
	PayStatus       string  `gorm:"comment:支付状态" json:"payStatus"`
	Remark          string  `gorm:"comment:备注" json:"remark"`
	Operator        string  `gorm:"comment:操作" json:"operator"`
	FirstLastGlobal int
	CreatedAt       time.Time      `json:"created_time"`
	UpdatedAt       time.Time      `json:"updated_time"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
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
	OutTradeNo        string         `gorm:"unique;comment:番id" json:"outTradeNo"`
	UserId            uint           `gorm:"index;comment:用户id" json:"userId"`
	UserName          string         `gorm:"comment:用户昵称" json:"userName"`
	GoodID            uint           `gorm:"comment:真实商品id;index" json:"goodId"`
	GoodName          string         `gorm:"comment:商品名字" json:"goodName"`
	FanId             uint           `gorm:"comment:所属蕃的id;index" json:"fanId"`
	FanName           string         `gorm:"comment:所属蕃的名字" json:"fanName"`
	IpID              uint           `gorm:"comment:所属ip的id" json:"ipId"`
	IpName            string         `gorm:"comment:所属ip的名字" json:"ipName"`
	SeriesID          uint           `gorm:"comment:所属系列id" json:"seriesId"`
	SeriesName        string         `gorm:"comment:所属系列id"  json:"seriesName"`
	Pic               string         `gorm:"comment:图片;type:varchar(128);not null" json:"pic"`
	Price             float64        `gorm:"comment:建议售价" json:"price"`
	PrizeIndexName    string         `gorm:"comment:A赏,B赏..." json:"prizeIndexName"`
	Remark            string         `gorm:"comment:备注" json:"remark"`
	SingleOrMuti      int            `json:"singleOrMuti"`
	MultiIds          GormList       `gorm:"type:varchar(128);not null"`
	Status            int            `gorm:"comment:上下架" json:"status"`
	SoldStatus        int            `gorm:"index;comment:是否售罄1.奖品售罄,2.奖品未售罄" json:"soldStatus"`
	TimeForSoldStatus string         `gorm:"comment:预售时间" json:"timeForSoldStatus,omitempty"`
	PkgStatus         int            `gorm:"comment:打包状态" json:"pkgStatus"`
	CreatedAt         time.Time      `json:"created_time"`
	UpdatedAt         time.Time      `json:"updated_time"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"-"`
}

type Sure struct {
	ID         uint           `gorm:"primarykey" json:"id"`
	FanId      uint           `gorm:"comment:番的id;uniqueIndex:udx_name" json:"fanId"`
	FanTitle   string         `gorm:"comment:番的标题;"  json:"fanTitle"`
	BoxId      uint           `gorm:"comment:箱子的id;uniqueIndex:udx_name" json:"boxId"`
	PrizeIndex GormList       `gorm:"type:varchar(1024);not null"`
	CreatedAt  time.Time      `json:"created_time"`
	UpdatedAt  time.Time      `json:"updated_time"`
	WhoUpdate  string         `gorm:"comment:更新人" json:"whoUpdate,omitempty"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}
type Left struct {
	ID         uint           `gorm:"primarykey" json:"id"`
	FanId      uint           `gorm:"comment:番的id;uniqueIndex:udx_name" json:"fanId"`
	FanTitle   string         `gorm:"comment:番的标题;"  json:"fanTitle"`
	BoxId      uint           `gorm:"comment:箱子的id;uniqueIndex:udx_name" json:"boxId"`
	PrizeIndex GormList       `gorm:"type:varchar(10240);not null"`
	CreatedAt  time.Time      `json:"created_time"`
	UpdatedAt  time.Time      `json:"updated_time"`
	WhoUpdate  string         `gorm:"comment:更新人" json:"whoUpdate,omitempty"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}
type Target struct {
	ID         uint           `gorm:"primarykey" json:"id"`
	FanId      uint           `gorm:"comment:番的id;uniqueIndex:udx_name" json:"fanId"`
	FanTitle   string         `gorm:"comment:番的标题;"  json:"fanTitle"`
	BoxId      uint           `gorm:"comment:箱子的id;uniqueIndex:udx_name" json:"boxId"`
	PrizeIndex GormList       `gorm:"type:varchar(10240);not null"`
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
	ID                  uint           `gorm:"primarykey" json:"id"`
	FanId               uint           `gorm:"comment:番的id;uniqueIndex:udx_name" json:"fanId"`
	FanTitle            string         `gorm:"comment:番的标题;"  json:"fanTitle"`
	BoxId               uint           `gorm:"comment:箱子的id;uniqueIndex:udx_name" json:"boxId"`
	Times               int            `gorm:"comment:抽奖次数记录;index" json:"times"`
	FirstPrizePosition  GormList       `gorm:"comment:奖品序号;type:varchar(128);index" json:"firstPrizePosition"`
	LastPrizePosition   GormList       `gorm:"comment:奖品序号;type:varchar(128);index" json:"lastPrizePosition"`
	GlobalPrizePosition GormList       `gorm:"comment:奖品序号;type:varchar(128);index" json:"globalPrizePosition"`
	CreatedAt           time.Time      `json:"created_time"`
	UpdatedAt           time.Time      `json:"updated_time"`
	WhoUpdate           string         `gorm:"comment:更新人" json:"whoUpdate,omitempty"`
	DeletedAt           gorm.DeletedAt `gorm:"index" json:"-"`
}
