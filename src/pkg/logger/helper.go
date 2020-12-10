package logger

import (
	"os"
)

type Helper struct {
	ILog
	fields map[string]interface{}
}

func NewHelper(log ILog) *Helper {
	return &Helper{ILog: log}
}

func (h *Helper) Info(args ...interface{}) {
	if !h.ILog.Options().Level.Enabled(InfoLevel) {
		return
	}
	h.ILog.Fields(h.fields).Log(InfoLevel, args...)
}

func (h *Helper) Infof(template string, args ...interface{}) {
	if !h.ILog.Options().Level.Enabled(InfoLevel) {
		return
	}
	h.ILog.Fields(h.fields).Logf(InfoLevel, template, args...)
}

func (h *Helper) Trace(args ...interface{}) {
	if !h.ILog.Options().Level.Enabled(TraceLevel) {
		return
	}
	h.ILog.Fields(h.fields).Log(TraceLevel, args...)
}

func (h *Helper) Tracef(template string, args ...interface{}) {
	if !h.ILog.Options().Level.Enabled(TraceLevel) {
		return
	}
	h.ILog.Fields(h.fields).Logf(TraceLevel, template, args...)
}

func (h *Helper) Debug(args ...interface{}) {
	if !h.ILog.Options().Level.Enabled(DebugLevel) {
		return
	}
	h.ILog.Fields(h.fields).Log(DebugLevel, args...)
}

func (h *Helper) Debugf(template string, args ...interface{}) {
	if !h.ILog.Options().Level.Enabled(DebugLevel) {
		return
	}
	h.ILog.Fields(h.fields).Logf(DebugLevel, template, args...)
}

func (h *Helper) Warn(args ...interface{}) {
	if !h.ILog.Options().Level.Enabled(WarnLevel) {
		return
	}
	h.ILog.Fields(h.fields).Log(WarnLevel, args...)
}

func (h *Helper) Warnf(template string, args ...interface{}) {
	if !h.ILog.Options().Level.Enabled(WarnLevel) {
		return
	}
	h.ILog.Fields(h.fields).Logf(WarnLevel, template, args...)
}

func (h *Helper) Error(args ...interface{}) {
	if !h.ILog.Options().Level.Enabled(ErrorLevel) {
		return
	}
	h.ILog.Fields(h.fields).Log(ErrorLevel, args...)
}

func (h *Helper) Errorf(template string, args ...interface{}) {
	if !h.ILog.Options().Level.Enabled(ErrorLevel) {
		return
	}
	h.ILog.Fields(h.fields).Logf(ErrorLevel, template, args...)
}

func (h *Helper) Fatal(args ...interface{}) {
	if !h.ILog.Options().Level.Enabled(FatalLevel) {
		return
	}
	h.ILog.Fields(h.fields).Log(FatalLevel, args...)
	os.Exit(1)
}

func (h *Helper) Fatalf(template string, args ...interface{}) {
	if !h.ILog.Options().Level.Enabled(FatalLevel) {
		return
	}
	h.ILog.Fields(h.fields).Logf(FatalLevel, template, args...)
	os.Exit(1)
}

func (h *Helper) WithError(err error) *Helper {
	fields := copyFields(h.fields)
	fields["error"] = err
	return &Helper{ILog: h.ILog, fields: fields}
}

func (h *Helper) WithFields(fields map[string]interface{}) *Helper {
	nfields := copyFields(fields)
	for k, v := range h.fields {
		nfields[k] = v
	}
	return &Helper{ILog: h.ILog, fields: nfields}
}
