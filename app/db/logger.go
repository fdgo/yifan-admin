package db

import (
	"fmt"
	gl "gorm.io/gorm/logger"
	"yifan/pkg/logger"
)

type writer struct {
	gl.Writer
}

// NewWriter writer 构造函数
// Author [SliverHorn](https://github.com/SliverHorn)
func NewWriter(w gl.Writer) *writer {
	return &writer{Writer: w}
}

// Printf 格式化打印日志
// Author [SliverHorn](https://github.com/SliverHorn)
func (w *writer) Printf(message string, data ...interface{}) {
	logger.Info(fmt.Sprintf(message+"\n", data...))
}
