package logger

import (
	"LIBRARY-API-SERVER/configs"
	"os"
	"path/filepath"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func temporaryLogger() *zap.Logger {
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
		zapcore.AddSync(os.Stdout),
		zapcore.DebugLevel,
	)
	return zap.New(core)
}

type Logger struct {
	*zap.Logger
}

func NewLogger(logLevel ...zapcore.Level) (*Logger, error) {
	tempLogger := temporaryLogger()
	if len(logLevel) == 0 {
		temp, err := configs.LoadConfig()
		if err != nil {
			return nil, err
		}
		configLogLevel, err := zapcore.ParseLevel(strings.ToLower(temp.Logging.Level))
		if err != nil {
			return nil, err
		}
		logLevel = append(logLevel, configLogLevel)
	}
	// ایجاد پوشه لاگ اگر وجود ندارد
	logDir := "./logs"
	err := os.MkdirAll(logDir, 0755)
	if err != nil {
		tempLogger.Error("Error occurred while creating the log folder.", zap.String("logDir", logDir), zap.Error(err))
		return nil, err
	}

	// ساخت نام فایل با تاریخ/زمان فعلی
	currentTime := time.Now().Format("2006-01-02_15-04-05")
	logFileName := filepath.Join(logDir, "log_"+currentTime+".log")

	// باز کردن فایل لاگ
	file, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		tempLogger.Error("Error occurred while opening the file.", zap.String("logFileName", logFileName), zap.Error(err))
		return nil, err
	}

	// تنظیم خروجی برای کنسول و فایل
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	fileEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())

	consoleSyncer := zapcore.AddSync(os.Stdout)
	fileSyncer := zapcore.AddSync(file)

	// ترکیب خروجی‌ها
	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, consoleSyncer, logLevel[0]),
		zapcore.NewCore(fileEncoder, fileSyncer, logLevel[0]),
	)

	// ساخت Logger
	zapLogger := zap.New(core)
	defer tempLogger.Info("Logs will now be stored in the \"logs\" folder.", zap.String(logLevel[0].String(), logFileName))
	return &Logger{zapLogger}, nil
}
