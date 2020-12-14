package mycasbin

import (
	"github.com/panda8z/istone/src/pkg/logger"
	"sync/atomic"

)

type CasbinLogger struct {
	enable int32
}

func (l *CasbinLogger) EnableLog(enable bool) {
	i := 0
	if enable {
		i = 1
	}
	atomic.StoreInt32(&(l.enable), int32(i))
}

func (l *CasbinLogger) IsEnabled() bool {
	return atomic.LoadInt32(&(l.enable)) != 0
}

func (l *CasbinLogger) Print(v ...interface{}) {
	if l.IsEnabled() {
		logger.DefaultLogger.Log(logger.InfoLevel, v...)
	}
}

func (l *CasbinLogger) Printf(format string, v ...interface{}) {
	if l.IsEnabled() {
		logger.DefaultLogger.Logf(logger.InfoLevel, format, v...)
	}
}
