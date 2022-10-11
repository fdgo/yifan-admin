package db

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gl "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
	"yifan/configs"
)

type Repo interface {
	GetDb() *gorm.DB
	Close()
	IsNotFound(errs ...error) bool
	RecordNotFound() bool
}
type dbRepo struct {
	Db *gorm.DB
}

func New() (Repo, error) {
	db, err := GormMysql()
	if err != nil {
		return nil, err
	}
	RegisterTables(db)
	return &dbRepo{
		Db: db,
	}, nil
}
func (d *dbRepo) GetDb() *gorm.DB {
	return d.Db
}
func (d *dbRepo) Close() {

}
func (i *dbRepo) IsNotFound(errs ...error) bool {
	if len(errs) > 0 {
		for _, err := range errs {
			if err == gorm.ErrRecordNotFound {
				return true
			}
		}
	}
	return i.RecordNotFound()
}
func (i *dbRepo) RecordNotFound() bool {
	return !errors.Is(i.Db.Error, gorm.ErrRecordNotFound)
}

func GormMysql() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		configs.GetConfig().DB.User, configs.GetConfig().DB.Password, configs.GetConfig().DB.Ip,
		configs.GetConfig().DB.Port, configs.GetConfig().DB.Db)
	mysqlConfig := mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		SkipInitializeWithVersion: false,
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			TablePrefix:   "yf_",
		},
		Logger: gl.New(NewWriter(log.New(os.Stdout, "\r\n", log.LstdFlags)), gl.Config{
			SlowThreshold: 200 * time.Millisecond,
			LogLevel:      gl.Info,
			Colorful:      true,
		}),
	}); err != nil {
		return nil, err
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(100)
		sqlDB.SetMaxOpenConns(100)
		return db, nil
	}
}
func RegisterTables(db *gorm.DB) {
	db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4 COMMENT='奖品表'").AutoMigrate(&Prize{})
	db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4 COMMENT='IP表'").AutoMigrate(&Ip{})
	db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4 COMMENT='系列表'").AutoMigrate(&Series{})
	db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4 COMMENT='蕃表'").AutoMigrate(&Fan{})
	db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4 COMMENT='商品表'").AutoMigrate(&Goods{})
	db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4 COMMENT='抽奖记录表'").AutoMigrate(&RecordPrize{})
	db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4 COMMENT='箱表'").AutoMigrate(&Box{})
	db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4 COMMENT='订单表'").AutoMigrate(&Order{})
	db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4 COMMENT='用户表'").AutoMigrate(&User{})
}
