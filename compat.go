// Copyright 2013 Marc Weistroff. All rights reserved.
// Copyright 2016 Patryk Pomykalski. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package log

import (
	"fmt"
	"os"
)

var stdFormatter = &LineFormatter{LineFormat: "%datetime% %level_name%: %message%"}
var stdHandler = &writeCloserHandler{wc: os.Stderr, Handler: &Handler{Level: DEBUG, Formatter: stdFormatter}}
var DefaultLogger = &Logger{Name: "", handlers: []HandlerInterface{stdHandler}, processors: []Processor{}}

// SetDefaultLogger is not thread safe, use with caution
func SetDefaultLogger(l *Logger) {
	DefaultLogger = l
}

// Fatal is equivalent to a call to Print followed by a call to os.Exit(1)
func Fatal(v ...interface{}) {
	Print(v)
	os.Exit(1)
}

// Fatal is equivalent to a call to Printf followed by a call to os.Exit(1)
func Fatalf(format string, v ...interface{}) {
	Printf(format, v)
	os.Exit(1)
}

// Fatalln is equivalent to a call to Println followed by a call to os.Exit(1)
func Fatalln(v ...interface{}) {
	Println(v)
	os.Exit(1)
}

// Panic is equivalent to a call to Print followed by a call to panic
func Panic(v ...interface{}) {
	s := fmt.Sprint(v)
	Print(s)
	panic(s)
}

// Panicf is equivalent to a call to Printf followed by a call to panic
func Panicf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	Print(s)
	panic(s)
}

// Panicln is equivalent to a call to Println followed by a call to panic
func Panicln(v ...interface{}) {
	s := fmt.Sprintln(v...)
	Print(s)
	panic(s)
}

// Print calls Debug in an instance of Logger where the only handler outputs to Stderr
func Print(v ...interface{}) {
	DefaultLogger.Debug(v...)
}

// Printf calls Debug in an instance of Logger where the only handler outputs to Stderr
// Arguments are handled in the manner of fmt.Printf.
func Printf(format string, v ...interface{}) {
	DefaultLogger.Debug(fmt.Sprintf(format, v...))
}

// Arguments are handled in the manner of fmt.Println.
func Println(v ...interface{}) {
	DefaultLogger.Debug(fmt.Sprintln(v...))
}

// Error calls Error in the default instance of Logger
func Error(v ...interface{}) {
	DefaultLogger.Error(v...)
}

// Errorf calls Error in the default instance of Logger
// Arguments are handled in the manner of fmt.Printf.
func Errorf(format string, v ...interface{}) {
	DefaultLogger.Error(fmt.Sprintf(format, v...))
}

// Info calls Info in the default instance of Logger
func Info(v ...interface{}) {
	DefaultLogger.Info(v...)
}

// Infof calls Info in the default instance of Logger
// Arguments are handled in the manner of fmt.Printf.
func Infof(format string, v ...interface{}) {
	DefaultLogger.Info(fmt.Sprintf(format, v...))
}
