package logger

import (
	"github.com/panda8z/istone/src/pkg/tools"
	"github.com/panda8z/istone/src/pkg/writer"
	"path/filepath"
)

// SetupLogger 日志
func SetupLogger(path string, subPath string) ILog {
	var setLogger ILog
	fullPath := filepath.Join(path, subPath)
	if !tools.PathExist(fullPath) {
		err := tools.PathCreate(fullPath)
		if err != nil {
			Fatal("create dir error: %s", err.Error())
		}
	}
	output, err := writer.NewFileWriter(fullPath, "log")
	if err != nil {
		Fatal("%s logger setup error: %s", subPath, err.Error())
	}
	setLogger = NewHelper(NewLogger(WithOutput(output)))
	return setLogger
}
