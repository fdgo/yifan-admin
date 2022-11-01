package service

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
	"time"
	"yifan/app/api/param"
	"yifan/app/db"
	"yifan/configs"
	"yifan/pkg/define"
	"yifan/pkg/jwtex"
	timex "yifan/pkg/times"
)

func Dncrypt(rawData, key, iv string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(rawData)
	key_b, err_1 := base64.StdEncoding.DecodeString(key)
	iv_b, _ := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return "", err
	}
	if err_1 != nil {
		return "", err_1
	}
	dnData, err := AesCBCDncrypt(data, key_b, iv_b)
	if err != nil {
		return "", err
	}
	return string(dnData), nil
}

// 解密
func AesCBCDncrypt(encryptData, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	blockSize := block.BlockSize()
	if len(encryptData) < blockSize {
		panic("ciphertext too short")
	}
	if len(encryptData)%blockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(encryptData, encryptData)
	// 解填充
	encryptData = PKCS7UnPadding(encryptData)
	return encryptData, nil
}

//去除填充
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
func (s *UserServiceImpl) IsNew(req param.ReqIsNew) (bool, error) {
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=" +
		configs.GetConfig().WxApp.AppId + "&secret=" +
		configs.GetConfig().WxApp.AppSecret + "&js_code=" +
		req.Code + "&grant_type=authorization_code"
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	body, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		return false, errors.New("服务正忙...")
	}
	userBase := map[string]string{}
	json.Unmarshal(body, &userBase)
	OpenId := userBase["openid"]
	//到库里检查是否老用户
	var user db.User
	result := s.db.GetDb().Where("open_id=?", OpenId).First(&user)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return false, errors.New("服务正忙...")
	}
	if result.RowsAffected == 0 { //新用户
		return true, nil
	}
	return false, nil
}
func (s *UserServiceImpl) GetOpenId(req param.ReqGetOpenId) (param.RespGetOpenId, error) {
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=" +
		configs.GetConfig().WxApp.AppId + "&secret=" +
		configs.GetConfig().WxApp.AppSecret + "&js_code=" +
		req.Code + "&grant_type=authorization_code"
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	body, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		return param.RespGetOpenId{}, err1
	}
	if bytes.Contains(body, []byte("errcode")) {
		ater := define.AccessTokenErrorResponse{}
		err := json.Unmarshal(body, &ater)
		if err != nil {
			return param.RespGetOpenId{}, errors.New("获取open_id,session_key失败...")
		}
		return param.RespGetOpenId{}, errors.New(ater.ErrMsg)
	}
	userBase := map[string]string{}
	json.Unmarshal(body, &userBase)
	src, err2 := Dncrypt(req.EncryptedData, userBase["session_key"], req.Iv)
	if err2 != nil {
		return param.RespGetOpenId{}, errors.New("iv,encryptedData有误...")
	}
	OpenId := userBase["openid"]
	SessionKey := userBase["session_key"]
	//到库里检查是否老用户
	var user db.User
	result := s.db.GetDb().Where("open_id=?", OpenId).First(&user)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return param.RespGetOpenId{}, errors.New("服务正忙...")
	}
	if result.RowsAffected == 0 { //新用户
		var userAdv = map[string]interface{}{}
		json.Unmarshal([]byte(src), &userAdv)
		id := define.GetRandUserId()
		token, _ := s.GenToken(id)
		nickname := userAdv["nickName"].(string)
		ava := userAdv["avatarUrl"].(string)
		s.db.GetDb().Create(&db.User{
			ID:         id,
			NickName:   nickname,
			AvatarUrl:  ava,
			OpenId:     OpenId,
			AppId:      configs.GetConfig().WxApp.AppId,
			SessionKey: SessionKey,
		})
		return param.RespGetOpenId{
			JwtToken: "Bearer " + token,
			UserId:   id,
			NickName: nickname,
			Avatar:   ava,
		}, nil
	}
	token, _ := s.GenToken(user.ID)
	return param.RespGetOpenId{
		JwtToken: "Bearer " + token,
		UserId:   user.ID,
		NickName: user.NickName,
		Avatar:   user.AvatarUrl,
	}, nil
}

//type UserInfo struct {
//	NickName  string `json:"nickName"`
//	Gender    int    `json:"gender"`
//	Language  string `json:"language"`
//	City      string `json:"city"`
//	Province  string `json:"province"`
//	Country   string `json:"country"`
//	UnionId   string `json:"unionId"`
//	AvatarURL string `json:"avatarUrl"`
//	Watermark struct {
//		Timestamp int    `json:"timestamp"`
//		Appid     string `json:"appid"`
//	} `json:"watermark"`
//}
func (s *UserServiceImpl) GenToken(user_id uint) (string, int64) {
	exp := time.Now().Unix() + int64(configs.GetConfig().Jwt.Exptime)
	claims := jwt.StandardClaims{Id: strconv.Itoa(int(user_id)), ExpiresAt: exp}
	jwtoken := &jwtex.JwtToken{SigningKey: []byte(configs.GetConfig().Jwt.JwtSecret)}
	token, _ := jwtoken.CreateToken(claims)
	return token, exp
}

func (s *UserServiceImpl) Delever(req param.ReqDelever) (param.RespDelever, error) {
	resp := param.RespDelever{}
	tx := s.db.GetDb().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	total := int64(0)
	err := tx.Model(&db.OrderDeliver{}).Count(&total).Error
	if err != nil {
		tx.Rollback()
		return param.RespDelever{}, errors.New("服务正忙...")
	}
	var deles []db.OrderDeliver
	result := tx.Find(&deles)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		tx.Rollback()
		return param.RespDelever{}, errors.New("服务正忙...")
	}
	for _, dele := range deles {
		var oneDel param.OneDelever
		var details []*db.OrderDeliverDetail
		result = tx.Where("out_trade_no=?", dele.OutTradeNo).Find(&details)
		if result.Error != nil {
			tx.Rollback()
			return param.RespDelever{}, result.Error
		}
		for _, oneDelDetail := range details {
			oneDel.DeleverGoods = append(oneDel.DeleverGoods, param.DeleverGood{
				Pic:            oneDelDetail.Pic,
				GoodName:       oneDelDetail.GoodName,
				IpName:         oneDelDetail.IpName,
				SeriesName:     oneDelDetail.SeriesName,
				PkgStatus:      oneDelDetail.PkgStatus,
				PrizeIndexName: oneDelDetail.PrizeIndexName,
			})
		}
		address := db.Address{}
		result = tx.Where("id=?", dele.AddressId).First(&address)
		if result.Error != nil {
			tx.Rollback()
			return param.RespDelever{}, result.Error
		}
		oneDel.DeleCompany = dele.DeleCompany
		oneDel.DeleOrderId = dele.DeleOrderId
		oneDel.ActiveSureTime = dele.ActiveSureTime
		oneDel.Num = dele.PrizeNum
		oneDel.Price = dele.Price
		oneDel.DeleverUserInfo = param.DeleverUserInfo{
			Address:  address.Detail,
			UserName: dele.UserName,
			UserId:   dele.UserId,
			Mobile:   dele.UserMobile,
		}
		oneDel.DeleverStatus = dele.DeleverStatus
		oneDel.OrderId = dele.ID
		oneDel.CreateTime = timex.TimeInt64ToTimeString(timex.TimetToInt64(dele.CreatedAt))
		resp.OneDelevers = append(resp.OneDelevers, oneDel)
	}
	resp.Num = len(resp.OneDelevers)
	resp.Num = len(resp.OneDelevers)
	resp.AllPages = math.Ceil(float64(total) / float64(req.PageSize))
	tx.Commit()
	return resp, nil
}
func (s *UserServiceImpl) OneDel(dele db.OrderDeliver) (param.OneDelever, error) {
	DB := s.db.GetDb()
	var oneDel param.OneDelever
	var details []*db.OrderDeliverDetail
	result := DB.Where("out_trade_no=?", dele.OutTradeNo).Find(&details)
	if result.Error != nil {
		return param.OneDelever{}, result.Error
	}
	for _, oneDelDetail := range details {
		oneDel.DeleverGoods = append(oneDel.DeleverGoods, param.DeleverGood{
			Pic:            oneDelDetail.Pic,
			GoodName:       oneDelDetail.GoodName,
			IpName:         oneDelDetail.IpName,
			SeriesName:     oneDelDetail.SeriesName,
			PkgStatus:      oneDelDetail.PkgStatus,
			PrizeIndexName: oneDelDetail.PrizeIndexName,
		})
	}
	address := db.Address{}
	result = DB.Where("id=?", dele.AddressId).First(&address)
	if result.Error != nil {
		return param.OneDelever{}, result.Error
	}
	oneDel.DeleCompany = dele.DeleCompany
	oneDel.DeleOrderId = dele.DeleOrderId
	oneDel.ActiveSureTime = dele.ActiveSureTime
	oneDel.Num = dele.PrizeNum
	oneDel.Price = dele.Price
	oneDel.DeleverUserInfo = param.DeleverUserInfo{
		Address:  address.Detail,
		UserName: dele.UserName,
		Mobile:   dele.UserMobile,
		UserId:   dele.UserId,
	}
	oneDel.DeleverStatus = dele.DeleverStatus
	oneDel.OrderId = dele.ID
	oneDel.CreateTime = timex.TimeInt64ToTimeString(timex.TimetToInt64(dele.CreatedAt))
	return oneDel, nil
}

func (s *UserServiceImpl) SetDelId(req param.ReqSetDelId) error {
	err := s.db.GetDb().Model(&db.OrderDeliver{}).Where("id=?", req.Id).Update("dele_order_id", req.DeleOrderId).Update("dele_company", req.DeleCompany).Error
	return err
}

func (s *UserServiceImpl) DeleverCondition(req param.ReqDeleverCondition) (param.RespDeleverCondition, error) {
	if req.GoodId != 0 {
		var (
			DB              = s.db.GetDb()
			total           = int64(0)
			resp            = param.RespDeleverCondition{}
			orderDelDetails []db.OrderDeliverDetail
		)
		DB.Model(&db.OrderDeliverDetail{}).Where("good_id=?", req.GoodId).Find(&orderDelDetails)
		for _, oneOdd := range orderDelDetails {
			var orderDels []db.OrderDeliver
			if req.DeleverStatus != 0 {
				DB = DB.Model(&db.OrderDeliver{}).Where("out_trade_no=?", oneOdd.OutTradeNo).Where("delever_status=?", req.DeleverStatus)
			}
			if req.OrderId != 0 {
				DB = DB.Model(&db.OrderDeliver{}).Where("out_trade_no=?", oneOdd.OutTradeNo).Where("id=?", req.OrderId)
			}
			if req.DeleOrderId != 0 {
				DB = DB.Model(&db.OrderDeliver{}).Where("out_trade_no=?", oneOdd.OutTradeNo).Where("dele_order_id=?", req.UserId)
			}
			if req.Mobile != "" {
				DB = DB.Model(&db.OrderDeliver{}).Where("out_trade_no=?", oneOdd.OutTradeNo).Where("user_mobile=?", req.Mobile)
			}
			if req.UserId != 0 {
				DB = DB.Model(&db.OrderDeliver{}).Where("out_trade_no=?", oneOdd.OutTradeNo).Where("user_id=?", req.UserId)
			}
			if len(req.TimeRange) == 2 {
				DB = DB.Model(&db.OrderDeliver{}).Where("out_trade_no=?", oneOdd.OutTradeNo).Where("created_at Between ? and ?",
					time.Unix(req.TimeRange[0], 0).Format("2006-01-02 15:04:05"),
					time.Unix(req.TimeRange[1], 0).Format("2006-01-02 15:04:05"))
			}
			result := DB.Find(&orderDels)
			if result.Error != nil {
				return param.RespDeleverCondition{}, result.Error
			}
			if len(orderDels) == 0 {
				continue
			}
			var oneDel param.OneDelever
			oneDel.DeleverGoods = append(oneDel.DeleverGoods, param.DeleverGood{
				Pic:            oneOdd.Pic,
				GoodName:       oneOdd.GoodName,
				IpName:         oneOdd.IpName,
				SeriesName:     oneOdd.SeriesName,
				PkgStatus:      oneOdd.PkgStatus,
				PrizeIndexName: oneOdd.PrizeIndexName,
			})
			for _, dele := range orderDels {
				address := db.Address{}
				result = DB.Where("id=?", dele.AddressId).First(&address)
				if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
					return param.RespDeleverCondition{}, result.Error
				}
				oneDel.DeleCompany = dele.DeleCompany
				oneDel.DeleOrderId = dele.DeleOrderId
				oneDel.ActiveSureTime = dele.ActiveSureTime
				oneDel.Num = dele.PrizeNum
				oneDel.Price = dele.Price
				oneDel.DeleverUserInfo = param.DeleverUserInfo{
					Address:  address.Detail,
					UserName: dele.UserName,
					Mobile:   dele.UserMobile,
					UserId:   dele.UserId,
				}
				oneDel.DeleverStatus = dele.DeleverStatus
				oneDel.OrderId = dele.ID
				oneDel.CreateTime = timex.TimeInt64ToTimeString(timex.TimetToInt64(dele.CreatedAt))
			}
			resp.OneDelevers = append(resp.OneDelevers, oneDel)
			total = int64(len(orderDels))
		}
		resp.Num = len(resp.OneDelevers)
		resp.AllPages = math.Ceil(float64(total) / float64(req.PageSize))
		return resp, nil
	} else {
		DB := s.db.GetDb()
		if req.DeleverStatus != 0 {
			DB = DB.Model(&db.OrderDeliver{}).Where("delever_status=?", req.DeleverStatus)
		}
		if req.OrderId != 0 {
			DB = DB.Model(&db.OrderDeliver{}).Where("id=?", req.OrderId)
		}
		if req.DeleOrderId != 0 {
			DB = DB.Model(&db.OrderDeliver{}).Where("dele_order_id=?", req.UserId)
		}
		if req.Mobile != "" {
			DB = DB.Model(&db.OrderDeliver{}).Where("user_mobile=?", req.Mobile)
		}
		if req.UserId != 0 {
			DB = DB.Model(&db.OrderDeliver{}).Where("user_id=?", req.UserId)
		}
		if len(req.TimeRange) == 2 {
			DB = DB.Model(&db.OrderDeliver{}).Where("created_at Between ? and ?",
				time.Unix(req.TimeRange[0], 0).Format("2006-01-02 15:04:05"),
				time.Unix(req.TimeRange[1], 0).Format("2006-01-02 15:04:05"))
		}
		var resp param.RespDeleverCondition
		total := int64(0)
		err := DB.Model(&db.OrderDeliver{}).Count(&total).Error
		if err != nil {
			return param.RespDeleverCondition{}, err
		}
		var orderDelivers []db.OrderDeliver
		if err := DB.Limit(int(req.PageSize)).Offset(int((req.PageIndex - 1) * req.PageSize)).Order("created_at desc").Find(&orderDelivers).Error; err != nil {
			return param.RespDeleverCondition{}, errors.New("服务正忙...")
		}
		for _, oneOd := range orderDelivers {
			oneDer, err := s.OneDel(oneOd)
			if err != nil {
				return param.RespDeleverCondition{}, err
			}
			resp.OneDelevers = append(resp.OneDelevers, oneDer)
		}
		resp.Num = len(resp.OneDelevers)
		resp.AllPages = math.Ceil(float64(total) / float64(req.PageSize))
		return resp, nil
	}
}
func (s *UserServiceImpl) DeleverDetail(req param.ReqDeleverDetail) (param.RespDeleverDetail, error) {
	resp := param.RespDeleverDetail{}
	tx := s.db.GetDb().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	var dele db.OrderDeliver
	result := tx.Where("id=?", req.Id).First(&dele)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		tx.Rollback()
		return param.RespDeleverDetail{}, errors.New("服务正忙...")
	}
	var details []*db.OrderDeliverDetail
	result = tx.Where("out_trade_no=?", dele.OutTradeNo).Find(&details)
	if result.Error != nil {
		tx.Rollback()
		return param.RespDeleverDetail{}, result.Error
	}
	for _, oneDelDetail := range details {
		resp.DeleverGoods = append(resp.DeleverGoods, param.DeleverGood{
			Pic:            oneDelDetail.Pic,
			GoodName:       oneDelDetail.GoodName,
			IpName:         oneDelDetail.IpName,
			SeriesName:     oneDelDetail.SeriesName,
			PkgStatus:      oneDelDetail.PkgStatus,
			PrizeIndexName: oneDelDetail.PrizeIndexName,
			GoodId:         oneDelDetail.GoodId,
		})
	}
	address := db.Address{}
	result = tx.Where("id=?", dele.AddressId).First(&address)
	if result.Error != nil {
		tx.Rollback()
		return param.RespDeleverDetail{}, result.Error
	}
	resp.DeleCompany = dele.DeleCompany
	resp.DeleOrderId = dele.DeleOrderId
	resp.ActiveSureTime = dele.ActiveSureTime
	resp.Num = dele.PrizeNum
	resp.Price = dele.Price
	resp.PayStyle = dele.PayStyle
	resp.DeleverUserInfo = param.DeleverUserInfo{
		Address:  address.Detail,
		UserName: dele.UserName,
		Mobile:   dele.UserMobile,
		UserId:   dele.UserId,
	}
	resp.DeleverStatus = dele.DeleverStatus
	resp.OrderId = dele.ID
	resp.CreateTime = timex.TimeInt64ToTimeString(timex.TimetToInt64(dele.CreatedAt))
	tx.Commit()
	return resp, nil
}

func (s *UserServiceImpl) UserList(req param.ReqUserList) (param.RespUserList, error) {
	DB := s.db.GetDb()
	total := int64(0)
	err := DB.Model(&db.User{}).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return param.RespUserList{}, errors.New("服务正忙......")
	}
	users := []*db.User{}
	if err = DB.Limit(int(req.PageSize)).Offset(int((req.PageIndex - 1) * req.PageSize)).Order("created_at desc").Find(&users).Error; err != nil {
		return param.RespUserList{}, errors.New("服务正忙...")
	}
	var resp param.RespUserList
	for _, oneUser := range users {
		lugs := []db.Luggage{}
		result := DB.Model(&db.Luggage{}).Where("user_id=?", oneUser.ID).Find(&lugs)
		if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
			return param.RespUserList{}, errors.New("服务正忙...")
		}
		if len(lugs) == 0 {
			continue
		}
		var u param.User
		for _, one := range lugs {
			var deleStatus int
			deleStatus = one.DeleverStatus
			u.ConsumptionFee += one.Price
			orderDetail := db.OrderDeliverDetail{}
			err = DB.Where("luggage_id=?", one.ID).First(&orderDetail).Error
			if err == gorm.ErrRecordNotFound {
				continue
			}
			if err != nil {
				return param.RespUserList{}, errors.New("服务正忙...")
			}
			deleStatus = orderDetail.DeleverStatus
			u.Luggage = append(u.Luggage, param.Luggage{
				ID:             one.ID,
				OutTradeNo:     one.OutTradeNo,
				GoodID:         one.GoodID,
				GoodName:       one.GoodName,
				IpID:           one.IpID,
				IpName:         one.IpName,
				SeriesID:       one.SeriesID,
				SeriesName:     one.SeriesName,
				Pic:            one.Pic,
				Price:          one.Price,
				PrizeIndexName: one.PrizeIndexName,
				PrizeIndex:     one.PrizeIndex,
				InLuggageTime:  timex.TimetToInt64(one.CreatedAt),
				OutLuggageTime: timex.TimetToInt64(orderDetail.CreatedAt),
				DeleStatus:     deleStatus,
			})
		}
		u.LuggageNum = int64(len(lugs))
		u.UserId = oneUser.ID
		u.Mobile = oneUser.Mobile
		u.NickName = oneUser.NickName
		u.Avatar = oneUser.AvatarUrl
		u.CreatTime = oneUser.CreatedAt.Format("2006-01-02 15:04:05")
		resp.Users = append(resp.Users, u)
	}
	resp.Num = len(resp.Users)
	resp.AllPages = math.Ceil(float64(total) / float64(req.PageSize))
	return resp, nil
}
func (s *UserServiceImpl) UserListCondition(req param.ReqUserListCondition) (param.RespUserListCondition, error) {
	DB := s.db.GetDb()
	if req.UserId != 0 {
		DB = DB.Where("id=?", req.UserId)
	}
	if req.NickName != "" {
		DB = DB.Where("nick_name=?", req.NickName)
	}
	if req.Mobile != "" {
		DB = DB.Where("mobile=?", req.Mobile)
	}
	total := int64(0)
	err := DB.Model(&db.User{}).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return param.RespUserListCondition{}, errors.New("服务正忙......")
	}
	users := []*db.User{}
	if err = DB.Limit(int(req.PageSize)).Offset(int((req.PageIndex - 1) * req.PageSize)).Order("created_at desc").Find(&users).Error; err != nil {
		return param.RespUserListCondition{}, errors.New("服务正忙...")
	}
	var resp param.RespUserListCondition
	for _, oneUser := range users {
		lugs := []db.Luggage{}
		s.db.GetDb().Model(&db.Luggage{}).Where("user_id=?", oneUser.ID).Find(&lugs)
		var u param.User
		for _, one := range lugs {
			var deleStatus int
			deleStatus = one.DeleverStatus
			u.ConsumptionFee += one.Price
			orderDetail := db.OrderDeliverDetail{}
			err = s.db.GetDb().Where("luggage_id=?", one.ID).First(&orderDetail).Error
			if err == gorm.ErrRecordNotFound {
				continue
			}
			if err != nil {
				return param.RespUserListCondition{}, errors.New("服务正忙...")
			}
			deleStatus = orderDetail.DeleverStatus
			u.Luggage = append(u.Luggage, param.Luggage{
				ID:             one.ID,
				OutTradeNo:     one.OutTradeNo,
				GoodID:         one.GoodID,
				GoodName:       one.GoodName,
				IpID:           one.IpID,
				IpName:         one.IpName,
				SeriesID:       one.SeriesID,
				SeriesName:     one.SeriesName,
				Pic:            one.Pic,
				Price:          one.Price,
				PrizeIndexName: one.PrizeIndexName,
				PrizeIndex:     one.PrizeIndex,
				InLuggageTime:  timex.TimetToInt64(one.CreatedAt),
				OutLuggageTime: timex.TimetToInt64(orderDetail.CreatedAt),
				DeleStatus:     deleStatus,
			})
		}
		u.LuggageNum = int64(len(lugs))
		u.UserId = oneUser.ID
		u.Mobile = oneUser.Mobile
		u.NickName = oneUser.NickName
		u.Avatar = oneUser.AvatarUrl
		u.CreatTime = oneUser.CreatedAt.Format("2006-01-02 15:04:05")
		resp.Users = append(resp.Users, u)
	}
	resp.Num = len(resp.Users)
	resp.AllPages = math.Ceil(float64(total) / float64(req.PageSize))
	return resp, nil
}
