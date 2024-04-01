package xorm_logger

import "go.uber.org/zap"

type zapOptions struct {
	*baseOptions

	sugared       bool
	logger        *zap.Logger
	sugaredLogger *zap.SugaredLogger
}

type ZapOption func(*zapOptions)

// WithZapLogger sets the logger to use. Global *zap.Logger will be used if the logger is nil.
func WithZapLogger(logger any) ZapOption {
	return func(options *zapOptions) {
		switch l := logger.(type) {
		case *zap.Logger:
			options.sugared = false
			options.logger = l
			options.sugaredLogger = nil
		case *zap.SugaredLogger:
			options.sugared = true
			options.logger = nil
			options.sugaredLogger = l
		default:
			panic("WithZapLogger: logger must be *zap.Logger or *zap.SugaredLogger. ")
		}
	}
}

// WithZapMsg sets the value of msg key. The default value is 'print sql'.
func WithZapMsg(msg string) ZapOption {
	return func(options *zapOptions) {
		options.msg = msg
	}
}

// WithZapDetail sets the key of sql value. The default value is 'detail'.
func WithZapDetail(detail string) ZapOption {
	return func(options *zapOptions) {
		options.detail = detail
	}
}

// WithZapShowSQL sets whether the sql will be printed or not. The default value is false.
func WithZapShowSQL(showSQL bool) ZapOption {
	return func(options *zapOptions) {
		options.showSQL = showSQL
	}
}
