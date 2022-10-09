package db

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"testing"
)

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

type Configuration struct {
	LogConfig LogConfig
	Server    ServerConfig
	DB        DBConfig
}

func DeleteModel() {
	dir, _ := ioutil.ReadDir("./model")
	for _, d := range dir {
		os.RemoveAll(path.Join([]string{"model", d.Name()}...))
	}
}

var t1 = fmt.Sprintf("DROP TABLE IF EXISTS wcd_style")
var t2 = fmt.Sprintf("DROP TABLE IF EXISTS wcd_goods")
var t3 = fmt.Sprintf("DROP TABLE IF EXISTS wcd_style_pic")
var t4 = fmt.Sprintf("DROP TABLE IF EXISTS wcd_user")

func TestCreate(t *testing.T) {
	file, _ := os.Open("../../configs/conf.json")
	defer file.Close()
	cfg = &Configuration{}
	decoder := json.NewDecoder(file)
	decoder.Decode(cfg)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.DB.User, cfg.DB.Password, cfg.DB.Ip,
		cfg.DB.Port, cfg.DB.Db)
	mysqlConfig := mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         64,
		SkipInitializeWithVersion: false,
	}
	db, _ := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	DeleteTable(db)
	DeleteModel()
	RegisterTables(db)
	CreateModel()
}
func DeleteTable(db *gorm.DB) {
	db.Exec(t1)
	db.Exec(t2)
	db.Exec(t3)
	db.Exec(t4)
}
func CreateModel() {
	x := exec.Command("cmd.exe", "/c", "start "+"E:\\wcd\\app\\db\\gormt.exe")
	x.Run()
}
