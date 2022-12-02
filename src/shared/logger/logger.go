package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	loggerGorm "gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type BaseLogger struct {
	Logging         *zap.SugaredLogger
	DatabaseLogging loggerGorm.Interface
}

var Logger *BaseLogger = &BaseLogger{}

func GetNewLogger() *zap.SugaredLogger {
	atom := zap.NewAtomicLevel()
	atom.SetLevel(zap.InfoLevel) // level has been set
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./logs/all.log",
		MaxAge:     28, // days
		MaxSize:    20, // megabytes
		MaxBackups: 3,
	})
	zapNewConsole := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:   "message",
		LevelKey:     "level",
		TimeKey:      "time",
		CallerKey:    "caller",
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeLevel:  zapcore.CapitalLevelEncoder,
		EncodeTime:   SyslogTimeEncoder,
	})
	zapFile := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:   "message",
		LevelKey:     "level",
		TimeKey:      "time",
		CallerKey:    "caller",
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeLevel:  CustomLevelEncoder,
		EncodeTime:   SyslogTimeEncoder,
	})
	core := zapcore.NewTee(
		zapcore.NewCore(zapNewConsole,
			w,
			atom,
		),
		zapcore.NewCore(zapFile, zapcore.AddSync(os.Stdout), atom),
	)
	logger := zap.New(core, zap.AddCaller())
	sugar := logger.Sugar()
	defer func(sugar *zap.SugaredLogger) {
		err := sugar.Sync()
		if err != nil {
			return
		}
	}(sugar)

	wQuery := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./logs/query.log",
		MaxAge:     28, // days
		MaxSize:    20, // megabytes
		MaxBackups: 3,
	})

	databaseLogger := loggerGorm.New(
		log.New(wQuery, "\r\n", log.LstdFlags), // io writer
		loggerGorm.Config{
			SlowThreshold:             time.Second,     // Slow SQL threshold
			LogLevel:                  loggerGorm.Info, // Log level
			IgnoreRecordNotFoundError: true,
		},
	)
	createLogger := &BaseLogger{
		Logging:         sugar,
		DatabaseLogging: databaseLogger,
	}
	Logger = createLogger
	return sugar
}

func SyslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

func CustomLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	zapcore.CapitalColorLevelEncoder(level, enc)
}
func GetLogger() *BaseLogger {
	return Logger
}
