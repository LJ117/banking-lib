package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 小写私有变量，不做直接指定方便后期要换其他库。
var log *zap.Logger

// initialize the zap logger
func init() {
	var err error

	// 自定义zap配置:
	//	e.g. 自定义日志现实的时间戳
	config := zap.NewProductionConfig()
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	//encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	// 关闭 stack trace:
	//  直接设置该配置的 key == ""
	encoderConfig.StacktraceKey = ""
	config.EncoderConfig = encoderConfig

	// 使用自定义的配置的zap
	log, err = config.Build(zap.AddCallerSkip(1))

	//log, err = zap.NewProduction(
	//	// 这里是因为所有调用都是直接通过这个包内的方法调用的，
	//	// 因此，跳过这个包本身[ 1级 ]就是实际调用者所在的位置。
	//	// 可以精准捕获日志的位置了
	//	zap.AddCallerSkip(1),
	//
	//)

	if err != nil {
		panic(err)
	}
}

// Info 使用函数，使记录日志这件事本身和直接使用 Log 这个对象脱钩
func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	log.Error(message, fields...)
}
