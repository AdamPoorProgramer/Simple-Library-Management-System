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

func NewLogger(logLevel ...zapcore.Level) *zap.Logger {
	config, err0 := configs.LoadConfig()
	tempLogger := temporaryLogger()
	if err0 != nil {
		tempLogger.Error("Error loading config.", zap.Error(err0))
	}
	currentTime := time.Now().Format("2006-01-02_15-04-05")
	logFileName := filepath.Join(config.Logging.Path, "log_"+currentTime+".log")
	if len(logLevel) == 0 {
		configLogLevel, _ := zapcore.ParseLevel(strings.ToLower(config.Logging.Level))
		logLevel = append(logLevel, configLogLevel)
	}
	err1 := os.MkdirAll(config.Logging.Path, 0755)
	if err1 != nil {
		tempLogger.Error("Error occurred while creating the log folder.", zap.String("logDir", config.Logging.Path), zap.Error(err1))
	}

	file, err2 := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err2 != nil {
		tempLogger.Error("Error occurred while opening the file.", zap.String("logFileName", logFileName), zap.Error(err2))
	}

	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	fileEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())

	consoleSyncer := zapcore.AddSync(os.Stdout)
	fileSyncer := zapcore.AddSync(file)

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, consoleSyncer, logLevel[0]),
		zapcore.NewCore(fileEncoder, fileSyncer, logLevel[0]),
	)

	zapLogger := zap.New(core)
	defer tempLogger.Info("Logs will now be stored in the "+config.Logging.Path+" folder.", zap.String(logLevel[0].String(), logFileName))
	return zapLogger
}
