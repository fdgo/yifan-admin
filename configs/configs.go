package configs

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	defaultcfgfile = "configs/conf.json"
)

func GetConfig() *Configuration {
	return cfg
}

var cfg *Configuration

type ServerConfig struct {
	Port uint
}
type LogConfig struct {
	Level      string
	Name       string
	FileSize   int
	MaxBackups int
	MaxAge     int
}
type DBConfig struct {
	Ip              string
	Port            int
	User            string
	Password        string
	Db              string
	LogMode         string
	MaxOpenConn     int
	MaxIdleConn     int
	ConnMaxLifeTime int64
}
type CacheConfig struct {
	Ip       string
	Password string
}
type WxAppConfig struct {
	AppId     string
	AppSecret string
	Version   string
}

type JwtConfig struct {
	JwtSecret string
	Exptime   int
}

type Configuration struct {
	LogConfig LogConfig
	Server    ServerConfig
	DB        DBConfig
	Cache     CacheConfig
	WxApp     WxAppConfig
	Jwt       JwtConfig
}

func LoadDefaultConfig() error {
	return LoadConfig(defaultcfgfile)
}

func LoadConfig(cfgfile string) error {
	file, err := os.Open(cfgfile)
	if err != nil {
		fmt.Println("Open config file err:", err)
		return err
	}
	defer file.Close()
	cfg = &Configuration{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(cfg)
	if err != nil {
		fmt.Println("Decoder failed:", err)
		return err
	}
	return nil
}
func DoInit(c *Configuration) {
	cfg = c
}
