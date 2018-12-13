package zjlog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"testing"
)

type Person struct {
	name string
	age  int
}

func TestLog_Debug(t *testing.T) {
	log, err := NewLogger("DEBUG", true, "log/log_test.log")
	if err != nil {
		panic(err)
	}
	//zap.AddCallerSkip(1)
	log.Debug("This is debug ", "ZJ")
	log.Info("This is Ino f ", 3434)
	log.Warn("This is Warn ", Person{"MJ", 20})
	log.Error("This is error %v", 4.356)
}

func TestLog_GetLogLevel(t *testing.T) {
	log, err := NewLogger("ERROR", true, "log/log_test3.log")
	if err != nil {
		panic(err)
	}
	if log.GetLogLevel() != zapcore.ErrorLevel {
		t.Errorf("GetLogLevel failed, should be %v, get %v", zapcore.ErrorLevel, log.GetLogLevel())
	}
}

func BenchmarkLog_Debug(b *testing.B) {
	log, _ := NewLogger("DEBUG", true, "log/log_test2.log")
	for i := 0; i < b.N; i++ {
		log.Debug("This is debug", "ZJ")
	}
}

func BenchmarkLog_levelInfo_Debug(b *testing.B) {
	log, _ := NewLogger("Info", true, "log/log_test2.log")
	for i := 0; i < b.N; i++ {
		log.Debug("This is debug", "ZJ")
	}
}

func BenchmarkLog_Debug_Logger(b *testing.B) {
	log, _ := NewLogger("DEBUG", true, "log/log_test3.log")
	for i := 0; i < b.N; i++ {
		log.GetLogger().Debug("This is logger debug", zap.String("name", "ZJ"))
	}
}
func BenchmarkLog_levelInfo_Debug2(b *testing.B) {
	log, _ := NewLogger("Info", true, "log/log_test2.log")
	for i := 0; i < b.N; i++ {
		log.GetLogger().Debug("This is debug", zap.String("name", "ZJ"))
	}
}
