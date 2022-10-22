package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func setlevel(level string) *zap.Config {
	switch level {
	case "info":
		return &zap.Config{
			Encoding:         "json",
			Level:            zap.NewAtomicLevelAt(zapcore.InfoLevel),
			OutputPaths:      []string{"stderr"},
			ErrorOutputPaths: []string{"stderr"},
			EncoderConfig: zapcore.EncoderConfig{
				MessageKey:  "message",
				LevelKey:    "level",
				EncodeLevel: zapcore.LowercaseLevelEncoder,
				TimeKey:     "ts",
				EncodeTime:  zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"),
			},
		}
	default:
		return &zap.Config{
			Encoding:         "json",
			Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
			OutputPaths:      []string{"stderr"},
			ErrorOutputPaths: []string{"stderr"},
			EncoderConfig: zapcore.EncoderConfig{
				MessageKey:   "message",
				LevelKey:     "level",
				EncodeLevel:  zapcore.LowercaseLevelEncoder,
				TimeKey:      "ts",
				EncodeTime:   zapcore.EpochMillisTimeEncoder,
				CallerKey:    "caller",
				EncodeCaller: zapcore.ShortCallerEncoder,
			},
		}
	}
}

func Logger(level string) *zap.SugaredLogger {
	cfg := setlevel(level)
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	return logger.Sugar()
}
