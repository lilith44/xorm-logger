package xorm_logger

import (
	"log/slog"
)

type slogOptions struct {
	*baseOptions

	logger *slog.Logger
}

type SlogOption func(*slogOptions)

// WithSlogLogger sets the logger to use. Default *slog.Logger will be used if the logger is nil.
func WithSlogLogger(logger *slog.Logger) SlogOption {
	return func(options *slogOptions) {
		options.logger = logger
	}
}

// WithSlogMsg sets the value of msg key. The default value is 'print sql'.
func WithSlogMsg(msg string) SlogOption {
	return func(options *slogOptions) {
		options.msg = msg
	}
}

// WithSlogDetail sets the key of sql value. The default value is 'detail'.
func WithSlogDetail(detail string) SlogOption {
	return func(options *slogOptions) {
		options.detail = detail
	}
}

// WithSlogShowSQL sets whether the sql will be printed or not. The default value is false.
func WithSlogShowSQL(showSQL bool) SlogOption {
	return func(options *slogOptions) {
		options.showSQL = showSQL
	}
}
