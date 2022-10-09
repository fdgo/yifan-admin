package controller

import (
	"time"

	"yifan/app/api/service"
	"yifan/app/db"
	"yifan/configs"
	"yifan/pkg/logger"
)

var h *handler

var start, end time.Time

const uid = 1384711165950967808

func init() {
	err := configs.LoadConfig("../../../configs/conf.json")
	if err != nil {
		panic(err)
	}
	db, err := db.New()
	if err != nil {
		logger.Fatal("new db err", err)
	}
	h = &handler{
		userService: service.NewUserService(db),
	}

	tTimeStr := "2021-09-23 17:00:00"
	start, err = time.Parse("2006-01-02 15:04:05", tTimeStr)
	if err != nil {
		logger.Fatal("time,", err)
	}
	tTimeStr = "2021-10-27 9:00:00"
	end, err = time.Parse("2006-01-02 15:04:05", tTimeStr)
	if err != nil {
		logger.Fatal("time,", err)
	}
}
