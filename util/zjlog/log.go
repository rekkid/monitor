package zjlog

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var levelMap = map[string]zapcore.Level{
	"DEBUG":  zap.DebugLevel,
	"INFO":   zap.InfoLevel,
	"WARN":   zap.WarnLevel,
	"ERROR":  zap.ErrorLevel,
	"DPANIC": zap.DPanicLevel,
	"PANIC":  zap.PanicLevel,
	"FATAL":  zap.FatalLevel,
}

type Log struct {
	level   string
	isDebug bool
	*zap.SugaredLogger
	logger *zap.Logger
}

func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + t.Format("2006-01-02 15:04:05.000") + "]")
}

func NewLogger(level string, isDebug bool, file string) (*Log, error) {
	encoderCfg := zapcore.EncoderConfig{
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

	zapLevel, ok := levelMap[level]
	if !ok {
		zapLevel = zap.ErrorLevel
	}

	currLevel := zap.NewAtomicLevelAt(zapLevel)

	filename := file + "_log_" + time.Now().Format("2006-01-02") + ".log"
	customCfg := zap.Config{
		Level:            currLevel,
		Development:      true,
		Encoding:         "console",
		EncoderConfig:    encoderCfg,
		OutputPaths:      []string{"stdout", filename},
		ErrorOutputPaths: []string{"stderr"},
		DisableCaller:    false,
	}

	logger, err := customCfg.Build()
	if err != nil {
		fmt.Println("init logger error: ", err)
		return nil, err
	}
	//customCfg.
	return &Log{level, isDebug, logger.Sugar(), logger}, nil
}

func (l *Log) GetLogLevel() zapcore.Level {
	return levelMap[l.level]
}

func (l *Log) GetLogger() *zap.Logger {
	return l.logger
}
