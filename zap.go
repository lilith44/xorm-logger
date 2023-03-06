package xorm_logger

import (
	"go.uber.org/zap"
	"xorm.io/xorm/log"
)

type ZapLogger struct {
	logger *zap.SugaredLogger

	level   log.LogLevel
	showSQL bool
}

func NewZapLogger(logger *zap.SugaredLogger) *ZapLogger {
	return &ZapLogger{
		logger:  logger,
		level:   log.LOG_DEBUG,
		showSQL: true,
	}
}

func (zl *ZapLogger) Debug(v ...any) {
	if zl.level <= log.LOG_DEBUG {
		zl.logger.Debug(v...)
	}
}

func (zl *ZapLogger) Debugf(format string, v ...any) {
	if zl.level <= log.LOG_DEBUG {
		zl.logger.Debugf(format, v...)
	}
}

func (zl *ZapLogger) Error(v ...any) {
	if zl.level <= log.LOG_ERR {
		zl.logger.Error(v...)
	}
}

func (zl *ZapLogger) Errorf(format string, v ...any) {
	if zl.level <= log.LOG_ERR {
		zl.logger.Errorf(format, v...)
	}
}

func (zl *ZapLogger) Info(v ...any) {
	if zl.level <= log.LOG_INFO {
		zl.logger.Info(v...)
	}
}

func (zl *ZapLogger) Infof(format string, v ...any) {
	if zl.level <= log.LOG_INFO {
		zl.logger.Infof(format, v...)
	}
}

func (zl *ZapLogger) Warn(v ...any) {
	if zl.level <= log.LOG_WARNING {
		zl.logger.Warn(v...)
	}
}

func (zl *ZapLogger) Warnf(format string, v ...any) {
	if zl.level <= log.LOG_WARNING {
		zl.logger.Warnf(format, v...)
	}
}

func (zl *ZapLogger) Level() log.LogLevel {
	return zl.level
}

func (zl *ZapLogger) SetLevel(l log.LogLevel) {
	zl.level = l
}

func (zl *ZapLogger) ShowSQL(show ...bool) {
	if len(show) == 0 {
		zl.showSQL = true
		return
	}
	zl.showSQL = show[0]
}

func (zl *ZapLogger) IsShowSQL() bool {
	return zl.showSQL
}
