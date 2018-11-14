package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"time"
)

var Logger *zap.SugaredLogger
var logLevel zap.AtomicLevel

const (
	DEBUG_LEVEL = zap.DebugLevel
	INFO_LEVEL  = zap.InfoLevel
	WARN_LEVEL  = zap.WarnLevel
	ERR_LEVEL   = zap.ErrorLevel
	FATAL_LEVEL = zap.FatalLevel
)

type Config struct {
	Console bool
}

func NewEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		CallerKey:      "C",
		MessageKey:     "M",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}
func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}
func DefaultWrite() *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   "log/app.log",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
	}
}

//DefaultLogger 默认日志配置（onlyfile）
func FLogger() {
	SetupLogger(DefaultWrite())
}

//FCLogger 控制台和文件日志
func FCLogger() {
	SetupLogger(DefaultWrite(), os.Stdout)
}

//CLogger 控制台日志
func CLogger() {
	SetupLogger(os.Stdout)
}

//SetLogLevel 设置日志级别
func SetLogLevel(level zapcore.Level) {
	logLevel.SetLevel(level)
}

// 自定义日志输出
func SetupLogger(ws ...io.Writer) {
	outs := make([]zapcore.WriteSyncer, 0)
	logLevel = zap.NewAtomicLevel()
	for _, w := range ws {
		outs = append(outs, zapcore.AddSync(w))
	}
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(NewEncoderConfig()),
		zapcore.NewMultiWriteSyncer(outs...),
		logLevel,
	)
	logger := zap.New(core, zap.AddCaller())
	Logger = logger.Sugar()
}
