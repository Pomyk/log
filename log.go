// Copyright 2013 Marc Weistroff. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package log

import (
	"time"
)

type Severity int

const (
	// From /usr/include/sys/syslog.h. and RFC5424
	EMERGENCY Severity = iota // Emergency: system is unusable
	ALERT                     // Alert: action must be taken immediately
	CRITICAL                  // Critical: critical conditions
	ERROR                     // Error: error conditions
	WARNING                   // Warning: warning conditions
	NOTICE                    // Notice: normal but significant condition
	INFO                      // Informational: informational messages
	DEBUG                     // Debug: debug-level messages
)

// Textual translation of severities
var Severities = map[Severity]string{
	DEBUG:     "DEBUG",
	INFO:      "INFO",
	NOTICE:    "NOTICE",
	WARNING:   "WARNING",
	ERROR:     "ERROR",
	CRITICAL:  "CRITICAL",
	ALERT:     "ALERT",
	EMERGENCY: "EMERGENCY"}

// A record is a log message at a given time
type Record struct {
	Message   string                 `json:"message"`           // Text message of the log
	Formatted string                 `json:"-"`                 // Formatted version of the log (once all processors and formatters have done their jobs)
	Level     Severity               `json:"level"`             // Severity level
	LevelName string                 `json:"level_name"`        // Severity name
	Channel   string                 `json:"channel"`           // Logger's name
	Time      time.Time              `json:"@timestamp"`        // Creation date
	Context   interface{}            `json:"context,omitempty"` // Context set by logger's caller
	Extra     map[string]interface{} `json:"extra,omitempty"`   // Extra values that can be added by Processors
}

func newRecord(level Severity, channel, message string, context interface{}) *Record {
	return &Record{
		Message:   message,
		Level:     level,
		LevelName: Severities[level],
		Channel:   channel,
		Time:      time.Now(),
		Context:   context,
		Extra:     make(map[string]interface{})}
}
