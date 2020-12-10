// Package log provides debug logging
package logger

import (
	"encoding/json"
	"fmt"
	"time"
)


var (
	// Default buffer size if any
	DefaultSize = 256
	// Default formatter
	DefaultFormat = TextFormat
)

// Tlog is debug log interface for reading and writing logs
type Tlog interface {
	// Read reads log entries from the logger
	Read(...ReadTlogOption) ([]Record, error)
	// Write writes records to log
	Write(Record) error
	// Stream log records
	Stream() (Stream, error)
}

// Record is log record entry
type Record struct {
	// Timestamp of logged event
	Timestamp time.Time `json:"timestamp"`
	// Metadata to enrich log record
	Metadata map[string]string `json:"metadata"`
	// Value contains log entry
	Message interface{} `json:"message"`
}

// Stream returns a log stream
type Stream interface {
	Chan() <-chan Record
	Stop() error
}

// Format is a function which formats the output
type FormatFunc func(Record) string

// TextFormat returns text format
func TextFormat(r Record) string {
	t := r.Timestamp.Format("2006-01-02 15:04:05")
	return fmt.Sprintf("%s %v ", t, r.Message)
}

// JSONFormat is a json Format func
func JSONFormat(r Record) string {
	b, _ := json.Marshal(r)
	return string(b) + " "
}



// TlogOption used by the logger
type TlogOption func(*TlogOptions)

// TlogOptions are logger options
type TlogOptions struct {
	// Name of the log
	Name string
	// Size is the size of ring buffer
	Size int
	// Format specifies the output format
	Format FormatFunc
}

// Name of the log
func Name(n string) TlogOption {
	return func(o *TlogOptions) {
		o.Name = n
	}
}

// Size sets the size of the ring buffer
func Size(s int) TlogOption {
	return func(o *TlogOptions) {
		o.Size = s
	}
}

func Format(f FormatFunc) TlogOption {
	return func(o *TlogOptions) {
		o.Format = f
	}
}

// DefaultOptions returns default options
func DefaultTlogOptions() TlogOptions {
	return TlogOptions{
		Size: DefaultSize,
	}
}

// ReadOptions for querying the logs
type ReadTlogOptions struct {
	// Since what time in past to return the logs
	Since time.Time
	// Count specifies number of logs to return
	Count int
	// Stream requests continuous log stream
	Stream bool
}

// ReadOption used for reading the logs
type ReadTlogOption func(*ReadTlogOptions)

// Since sets the time since which to return the log records
func Since(s time.Time) ReadTlogOption {
	return func(o *ReadTlogOptions) {
		o.Since = s
	}
}

// Count sets the number of log records to return
func Count(c int) ReadTlogOption {
	return func(o *ReadTlogOptions) {
		o.Count = c
	}
}
