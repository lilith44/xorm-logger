package xorm_logger

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"xorm.io/xorm/log"
)

var zapLogLevelMapping = map[zapcore.Level]log.LogLevel{
	zapcore.DebugLevel: log.LOG_DEBUG,
	zapcore.InfoLevel:  log.LOG_INFO,
	zapcore.WarnLevel:  log.LOG_WARNING,
	zapcore.ErrorLevel: log.LOG_ERR,
}

type ZapLogger struct {
	*base

	sugared       bool
	logger        *zap.Logger
	sugaredLogger *zap.SugaredLogger
}

func NewZapLogger(options ...ZapOption) *ZapLogger {
	zapOptions := &zapOptions{
		baseOptions: &baseOptions{
			msg:    "print sql",
			detail: "detail",
		},
	}
	for _, option := range options {
		option(zapOptions)
	}
	if zapOptions.logger == nil && zapOptions.sugaredLogger == nil {
		zapOptions.logger = zap.L()
	}

	return &ZapLogger{
		base: &base{
			msg:     zapOptions.msg,
			detail:  zapOptions.detail,
			showSQL: zapOptions.showSQL,
		},
		sugared:       zapOptions.sugared,
		logger:        zapOptions.logger,
		sugaredLogger: zapOptions.sugaredLogger,
	}
}

func (z *ZapLogger) Debug(v ...any) {
	if z.sugared {
		z.sugaredLogger.Debug(v...)
	} else {
		z.logger.Debug(z.msg, zap.String(z.detail, fmt.Sprint(v...)))
	}
}

func (z *ZapLogger) Debugf(format string, v ...any) {
	if z.sugared {
		z.sugaredLogger.Debugf(format, v...)
	} else {
		z.logger.Debug(z.msg, zap.String(z.detail, fmt.Sprintf(format, v...)))
	}
}

func (z *ZapLogger) Error(v ...any) {
	if z.sugared {
		z.sugaredLogger.Error(v...)
	} else {
		z.logger.Error(z.msg, zap.String(z.detail, fmt.Sprint(v...)))
	}
}

func (z *ZapLogger) Errorf(format string, v ...any) {
	if z.sugared {
		z.sugaredLogger.Errorf(format, v...)
	} else {
		z.logger.Error(z.msg, zap.String(z.detail, fmt.Sprintf(format, v...)))
	}
}

func (z *ZapLogger) Info(v ...any) {
	if z.sugared {
		z.sugaredLogger.Info(v...)
	} else {
		z.logger.Info(z.msg, zap.String(z.detail, fmt.Sprint(v...)))
	}
}

func (z *ZapLogger) Infof(format string, v ...any) {
	if z.sugared {
		z.sugaredLogger.Infof(format, v...)
	} else {
		z.logger.Info(z.msg, zap.String(z.detail, fmt.Sprintf(format, v...)))
	}
}

func (z *ZapLogger) Warn(v ...any) {
	if z.sugared {
		z.sugaredLogger.Warn(v...)
	} else {
		z.logger.Warn(z.msg, zap.String(z.detail, fmt.Sprint(v...)))
	}
}

func (z *ZapLogger) Warnf(format string, v ...any) {
	if z.sugared {
		z.sugaredLogger.Warnf(format, v...)
	} else {
		z.logger.Warn(z.msg, zap.String(z.detail, fmt.Sprintf(format, v...)))
	}
}

func (z *ZapLogger) Level() log.LogLevel {
	if z.sugared {
		return parseLogLevelFromZapLogLevel(z.sugaredLogger.Level())
	}
	return parseLogLevelFromZapLogLevel(z.logger.Level())
}

func (z *ZapLogger) SetLevel(_ log.LogLevel) {

}

func parseLogLevelFromZapLogLevel(l zapcore.Level) log.LogLevel {
	if level, ok := zapLogLevelMapping[l]; ok {
		return level
	}
	return log.LOG_UNKNOWN
}
