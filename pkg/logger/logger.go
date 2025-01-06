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
	config, err := configs.LoadConfig()
	tempLogger := temporaryLogger()
	if err != nil {
		tempLogger.Error("Error loading config.", zap.Error(err))
		return nil, err
	}
	currentTime := time.Now().Format("2006-01-02_15-04-05")
	logFileName := filepath.Join(config.Logging.Path, "log_"+currentTime+".log")
	if len(logLevel) == 0 {
		configLogLevel, err := zapcore.ParseLevel(strings.ToLower(config.Logging.Level))
		if err != nil {
			return nil, err
		}
		logLevel = append(logLevel, configLogLevel)
	}
	err = os.MkdirAll(config.Logging.Path, 0755)
	if err != nil {
		tempLogger.Error("Error occurred while creating the log folder.", zap.String("logDir", config.Logging.Path), zap.Error(err))
		return nil, err
	}

	file, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		tempLogger.Error("Error occurred while opening the file.", zap.String("logFileName", logFileName), zap.Error(err))
		return nil, err
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
	defer tempLogger.Info("Logs will now be stored in the \"logs\" folder.", zap.String(logLevel[0].String(), logFileName))
	return &Logger{zapLogger}, nil
}
