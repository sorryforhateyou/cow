package main

// This logging trick is learnt from a post by Rob Pike
// https://groups.google.com/d/msg/golang-nuts/gU7oQGoCkmg/j3nNxuS2O_sJ

import (
	"github.com/golang/glog"
	//"github.com/cyfdecyf/color"
)

type infoLogging bool
type debugLogging bool
type errorLogging bool
type requestLogging bool
type responseLogging bool

var (
	info   infoLogging
	debug  debugLogging
	errl   errorLogging
	dbgRq  requestLogging
	dbgRep responseLogging

	verbose  bool
	colorize bool
)

func (d infoLogging) Printf(format string, args ...interface{}) {
	glog.Infof(format, args...)
	/*
		if d {
			log.Printf(format, args...)
		}
	*/
}

func (d infoLogging) Println(args ...interface{}) {
	glog.Infoln(args...)
	/*
		if d {
			log.Println(args...)
		}
	*/
}

func (d debugLogging) Printf(format string, args ...interface{}) {
	glog.Warningf(format, args...)
	/*
		if d {
			debugLog.Printf(format, args...)
		}
	*/
}

func (d debugLogging) Println(args ...interface{}) {
	glog.Infoln(args...)
	/*
		if d {
			debugLog.Println(args...)
		}
	*/
}

func (d errorLogging) Printf(format string, args ...interface{}) {
	glog.Errorf(format, args...)
	/*
		if d {
			errorLog.Printf(format, args...)
		}
	*/
}

func (d errorLogging) Println(args ...interface{}) {
	glog.Errorln(args...)
	/*
		if d {
			errorLog.Println(args...)
		}
	*/
}

func (d requestLogging) Printf(format string, args ...interface{}) {
	glog.Infof(format, args...)
	/*
		if d {
			requestLog.Printf(format, args...)
		}
	*/
}

func (d responseLogging) Printf(format string, args ...interface{}) {
	glog.Infof(format, args...)
	/*
		if d {
			responseLog.Printf(format, args...)
		}
	*/
}

func Fatal(args ...interface{}) {
	glog.Fatal(args...)
	/*
		fmt.Println(args...)
		os.Exit(1)
	*/
}

func Fatalf(format string, args ...interface{}) {
	glog.Fatalf(format, args...)
	/*
		fmt.Printf(format, args...)
		os.Exit(1)
	*/
}
