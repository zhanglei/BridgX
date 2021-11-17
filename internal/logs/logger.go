package logs

import (
	"fmt"
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.SugaredLogger

func Init() {
	Logger = createZapLogger().Sugar()
}

func createZapLogger() *zap.Logger {
	encoderConfig := zap.NewProductionEncoderConfig()
	var recordTimeFormat = "2006-01-02 15:04:05.000"
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format(recordTimeFormat))
	}
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.TimeKey = "created_at"
	var encoder = zapcore.NewConsoleEncoder(encoderConfig)
	serviceName := os.Getenv("ServiceName")
	if serviceName == "" {
		serviceName = "app"
	}
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fmt.Sprintf("/tmp/bridgx/logs/%v.log", serviceName),
		MaxSize:    200,
		MaxBackups: 7,
		MaxAge:     7,
		Compress:   false,
	}

	writer := zapcore.NewMultiWriteSyncer(zapcore.AddSync(lumberJackLogger), zapcore.AddSync(os.Stdout))
	zapCore := zapcore.NewCore(encoder, writer, zap.DebugLevel)
	return zap.New(zapCore, zap.AddCaller())
}
