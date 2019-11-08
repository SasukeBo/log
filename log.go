package log

import (
	"fmt"
	color "gopkg.in/gookit/color.v1"
	"log"
)

var (
	info    = color.Notice.Render
	warn    = color.Warn.Render
	success = color.Success.Render
	err     = color.Danger.Render
)

// Info print info log with blue color, format like fmt.Printf
func Info(format interface{}, a ...interface{}) {
	logf(info("INFO   "), format, a...)
}

// Infoln print line of info log with blue color, format like fmt.Println
func Infoln(a ...interface{}) {
	logln(info("INFO   "), a...)
}

// Success print success log with green color, format like fmt.Printf
func Success(format interface{}, a ...interface{}) {
	logf(success("SUCCESS"), format, a...)
}

// Successln print line of success log with green color, format like fmt.Println
func Successln(a ...interface{}) {
	logln(success("SUCCESS"), a...)
}

// Error print error log with red color, format like fmt.Printf
func Error(format string, a ...interface{}) {
	logf(err("ERROR  "), format, a...)
}

// Errorln print line of error log with red color, format like fmt.Println
func Errorln(a ...interface{}) {
	logln(err("ERROR  "), a...)
}

// Warn print warning log with orange color, format like fmt.Printf
func Warn(format string, a ...interface{}) {
	logf(warn("WARN   "), format, a...)
}

// Warnln print line of warning log with orange color, format like fmt.Println
func Warnln(a ...interface{}) {
	logln(warn("WARN   "), a...)
}

func logf(highlight string, format interface{}, a ...interface{}) {
	log.Printf(fmt.Sprintf("%s  ▶ %v", highlight, format), a...)
}

func logln(highlight string, a ...interface{}) {
	v := append([]interface{}{highlight, " ▶"}, a...)
	log.Println(v...)
}
