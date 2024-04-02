package xorm_logger

import (
	"fmt"
	"log/slog"

	"xorm.io/xorm/log"
)

var slogLogLevelMapping = map[slog.Level]log.LogLevel{
	slog.LevelDebug: log.LOG_DEBUG,
	slog.LevelInfo:  log.LOG_INFO,
	slog.LevelWarn:  log.LOG_WARNING,
	slog.LevelError: log.LOG_ERR,
}

type SlogLogger struct {
	*base

	logger *slog.Logger
}

func NewSlogLogger(options ...SlogOption) *SlogLogger {
	slogOptions := &slogOptions{
		baseOptions: &baseOptions{
			msg:    "print sql",
			detail: "detail",
		},
	}
	for _, option := range options {
		option(slogOptions)
	}
	if slogOptions.logger == nil {
		slogOptions.logger = slog.Default()
	}

	logger := &SlogLogger{
		base: &base{
			msg:     slogOptions.msg,
			detail:  slogOptions.detail,
			showSQL: slogOptions.showSQL,
		},
		logger: slogOptions.logger,
	}
	logger.level = slogLogLevelMapping[logger.minLevel()]
	return logger
}

func (s *SlogLogger) Debug(v ...any) {
	if s.level < log.LOG_DEBUG {
		return
	}
	s.logger.Debug(s.msg, s.detail, fmt.Sprint(v...))
}

func (s *SlogLogger) Debugf(format string, v ...any) {
	if s.level < log.LOG_DEBUG {
		return
	}
	s.logger.Debug(s.msg, s.detail, fmt.Sprintf(format, v...))
}

func (s *SlogLogger) Error(v ...any) {
	if s.level < log.LOG_ERR {
		return
	}
	s.logger.Error(s.msg, s.detail, fmt.Sprint(v...))
}

func (s *SlogLogger) Errorf(format string, v ...any) {
	if s.level < log.LOG_ERR {
		return
	}
	s.logger.Error(s.msg, s.detail, fmt.Sprintf(format, v...))
}

func (s *SlogLogger) Info(v ...any) {
	if s.level < log.LOG_INFO {
		return
	}
	s.logger.Info(s.msg, s.detail, fmt.Sprint(v...))
}

func (s *SlogLogger) Infof(format string, v ...any) {
	if s.level < log.LOG_INFO {
		return
	}
	s.logger.Info(s.msg, s.detail, fmt.Sprintf(format, v...))
}

func (s *SlogLogger) Warn(v ...any) {
	if s.level < log.LOG_WARNING {
		return
	}
	s.logger.Warn(s.msg, s.detail, fmt.Sprint(v...))
}

func (s *SlogLogger) Warnf(format string, v ...any) {
	if s.level < log.LOG_WARNING {
		return
	}
	s.logger.Warn(s.msg, s.detail, fmt.Sprintf(format, v...))
}

func (s *SlogLogger) Level() log.LogLevel {
	return s.level
}

func (s *SlogLogger) SetLevel(level log.LogLevel) {
	s.level = level
}

func (s *SlogLogger) minLevel() slog.Level {
	levels := []slog.Level{slog.LevelError, slog.LevelWarn, slog.LevelInfo, slog.LevelDebug}
	for _, level := range levels {
		if s.logger.Enabled(nil, level) {
			return level
		}
	}
	return slog.LevelDebug
}
