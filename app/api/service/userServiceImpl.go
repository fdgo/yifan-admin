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
	"net/http"
	"strconv"
	"time"
	"yifan/app/api/param"
	"yifan/app/db"
	"yifan/configs"
	"yifan/pkg/define"
	"yifan/pkg/jwtex"
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
