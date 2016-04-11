// Copyright 2013 Marc Weistroff. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package log

import (
	"os"
	"runtime"
)

// RuntimeProcessor adds some information about the current go runtime to a log record
var RuntimeProcessor = NewProcessor(runtimeProcessor)

// RuntimeProcessor adds hostname to a log record
var HostnameProcessor = NewProcessor(hostnameProcessor())

// A processor transforms a log records in whatever way it wants.
// It is usefull to add extra information to a log record
type Processor interface {
	Process(*Record)
}

type processor struct {
	process func(*Record)
}

// NewProcessor wraps a function to a Processor
func NewProcessor(f func(*Record)) Processor {
	return &processor{process: f}
}

func (p *processor) Process(r *Record) {
	p.process(r)
}

func runtimeProcessor(r *Record) {
	r.Extra["go.num_cpu"] = runtime.NumCPU()
	r.Extra["go.version"] = runtime.Version()
	r.Extra["go.num_goroutines"] = runtime.NumGoroutine()
}

func hostnameProcessor() func(*Record) {
	host, _ := os.Hostname()
	return func(r *Record) {
		r.Extra["host"] = host
	}
}
