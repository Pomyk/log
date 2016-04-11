// Copyright 2013 Marc Weistroff. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package log

import "fmt"

// A logger will log records transformed by the default processors to a collection of handlers
type Logger struct {
	Name       string
	handlers   []HandlerInterface
	processors []Processor
}

// Instanciates a new logger with specified name, handlers and processors
func NewLogger(name string) *Logger {
	return &Logger{Name: name, handlers: []HandlerInterface{}, processors: []Processor{}}
}

// Push a handler to the handlers stack
func (l *Logger) PushHandler(h HandlerInterface) {
	l.handlers = append(l.handlers, h)
}

// Pop a handler from the handlers stack
func (l *Logger) PopHandler() {
	if len(l.handlers) > 0 {
		l.handlers = l.handlers[0 : len(l.handlers)-1]
		return
	}

	panic("Handlers stack is empty")
}

// Push a processor to the processor stack
func (l *Logger) PushProcessor(p Processor) {
	l.processors = append(l.processors, p)
}

// Pop a processor from the processor stack
func (l *Logger) PopProcessor() {
	if len(l.processors) > 0 {
		l.processors = l.processors[0 : len(l.processors)-1]
		return
	}

	panic("Processors stack is empty")
}

func AddRecord(level Severity, message string, context interface{}) {
	DefaultLogger.AddRecord(level, message, context)
}

// Log string with specified severity
func (l *Logger) AddRecord(level Severity, message string, context interface{}) {
	if !l.S(level) {
		return
	}

	r := newRecord(level, l.Name, message, context)

	for k := range l.processors {
		l.processors[k].Process(r)
	}

	for k := range l.handlers {
		if l.handlers[k].S(level) {
			l.handlers[k].Handle(*r)
		}
	}
}

// Log string with the DEBUG level
func (l *Logger) AddDebug(message string, context interface{}) {
	l.AddRecord(DEBUG, message, context)
}

// Log string with the INFO level
func (l *Logger) AddInfo(message string, context interface{}) {
	l.AddRecord(INFO, message, context)
}

// Log string with the NOTICE level
func (l *Logger) AddNotice(message string, context interface{}) {
	l.AddRecord(NOTICE, message, context)
}

// Log string with the WARNING level
func (l *Logger) AddWarning(message string, context interface{}) {
	l.AddRecord(WARNING, message, context)
}

// Log string with the ERROR level
func (l *Logger) AddError(message string, context interface{}) {
	l.AddRecord(ERROR, message, context)
}

// Log string with the CRITICAL level
func (l *Logger) AddCritical(message string, context interface{}) {
	l.AddRecord(CRITICAL, message, context)
}

// Log string with the ALERT level
func (l *Logger) AddAlert(message string, context interface{}) {
	l.AddRecord(ALERT, message, context)
}

// Log string with the EMERGENCY level
func (l *Logger) AddEmergency(message string, context interface{}) {
	l.AddRecord(EMERGENCY, message, context)
}

// Log parameters with the DEBUG level
func (l *Logger) Debug(v ...interface{}) {
	l.AddDebug(fmt.Sprint(v...), nil)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.AddDebug(fmt.Sprintf(format, v...), nil)
}

func (l *Logger) Debugln(v ...interface{}) {
	l.AddDebug(fmt.Sprintln(v...), nil)
}

// Log parameters with the INFO level
func (l *Logger) Info(v ...interface{}) {
	l.AddInfo(fmt.Sprint(v...), nil)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.AddInfo(fmt.Sprintf(format, v...), nil)
}

func (l *Logger) Infoln(v ...interface{}) {
	l.AddInfo(fmt.Sprintln(v...), nil)
}

// Log parameters with the NOTICE level
func (l *Logger) Notice(v ...interface{}) {
	l.AddNotice(fmt.Sprint(v...), nil)
}

func (l *Logger) Noticef(format string, v ...interface{}) {
	l.AddNotice(fmt.Sprintf(format, v...), nil)
}

func (l *Logger) Noticeln(v ...interface{}) {
	l.AddNotice(fmt.Sprintln(v...), nil)
}

// Log parameters with the WARNING level
func (l *Logger) Warning(v ...interface{}) {
	l.AddWarning(fmt.Sprint(v...), nil)
}

func (l *Logger) Warningf(format string, v ...interface{}) {
	l.AddWarning(fmt.Sprintf(format, v...), nil)
}

func (l *Logger) Warningln(v ...interface{}) {
	l.AddWarning(fmt.Sprintln(v...), nil)
}

// Log parameters with the ERROR level
func (l *Logger) Error(v ...interface{}) {
	l.AddError(fmt.Sprint(v...), nil)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.AddError(fmt.Sprintf(format, v...), nil)
}

func (l *Logger) Errorln(v ...interface{}) {
	l.AddError(fmt.Sprintln(v...), nil)
}

// Log parameters with the CRITICAL level
func (l *Logger) Critical(v ...interface{}) {
	l.AddCritical(fmt.Sprint(v...), nil)
}

func (l *Logger) Criticalf(format string, v ...interface{}) {
	l.AddCritical(fmt.Sprintf(format, v...), nil)
}

func (l *Logger) Criticalln(v ...interface{}) {
	l.AddCritical(fmt.Sprintln(v...), nil)
}

// Log parameters with the ALERT level
func (l *Logger) Alert(v ...interface{}) {
	l.AddAlert(fmt.Sprint(v...), nil)
}

func (l *Logger) Alertf(format string, v ...interface{}) {
	l.AddAlert(fmt.Sprintf(format, v...), nil)
}

func (l *Logger) Alertln(v ...interface{}) {
	l.AddAlert(fmt.Sprintln(v...), nil)
}

// Log parameters with the EMERGENCY level
func (l *Logger) Emergency(v ...interface{}) {
	l.AddEmergency(fmt.Sprint(v...), nil)
}

func (l *Logger) Emergencyf(format string, v ...interface{}) {
	l.AddEmergency(fmt.Sprintf(format, v...), nil)
}

func (l *Logger) Emergencyln(v ...interface{}) {
	l.AddEmergency(fmt.Sprintln(v...), nil)
}

// Returns true if a Handler can handle this severity level
func (l *Logger) S(level Severity) bool {
	for k := range l.handlers {
		if l.handlers[k].S(level) {
			return true
		}
	}

	return false
}
