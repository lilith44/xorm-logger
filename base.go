package xorm_logger

import "xorm.io/xorm/log"

type base struct {
	msg     string
	detail  string
	showSQL bool
	level   log.LogLevel
}

func (b *base) IsShowSQL() bool {
	return b.showSQL
}

func (b *base) ShowSQL(show ...bool) {
	if len(show) == 0 {
		b.showSQL = true
		return
	}
	b.showSQL = show[0]
}

type baseOptions struct {
	msg, detail string
	showSQL     bool
}
