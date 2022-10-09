package logger

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.SugaredLogger

// log level, dynamic adjustment
var dynamicLevel zap.AtomicLevel

var levelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func getLoggerLevel(lvl string) zapcore.Level {
	if level, ok := levelMap[lvl]; ok {
		return level
	}
	return zapcore.InfoLevel
}

// timeEncoder setting time print format
func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

// Init ,input: log file name, file max size, max files number, max saved date
func Init(filename string, maxsize, maxbackups, maxage int) {
	//init log level, default info
	dynamicLevel = zap.NewAtomicLevel()

	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   filename,   // log file name
		MaxSize:    maxsize,    // Maximum per file size (MB)
		MaxBackups: maxbackups, // Maximum number of old log files retained
		MaxAge:     maxage,     // Saved maximum date
	})

	jsoncoder := zap.NewProductionEncoderConfig()
	jsoncoder.EncodeTime = timeEncoder
	jsoncoder.EncodeLevel = zapcore.CapitalLevelEncoder //log level upper

	core := zapcore.NewTee(
		// print to Console
		zapcore.NewCore(zapcore.NewConsoleEncoder(jsoncoder), os.Stdout, dynamicLevel),
		// saved to log file (json)
		zapcore.NewCore(zapcore.NewJSONEncoder(jsoncoder), w, dynamicLevel),
	)
	logger = zap.New(core, zap.AddStacktrace(zap.ErrorLevel)).Sugar()
}

// SetLevel , set level, dynamic adjustment
func SetLevel(lvl string) {
	dynamicLevel.SetLevel(getLoggerLevel(lvl))
}

// Close , asynchronous write to log file if closed
func Close() {
	logger.Sync()
}
