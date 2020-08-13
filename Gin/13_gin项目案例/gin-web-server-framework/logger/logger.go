package logger

import (
	"os"

	"github.com/op/go-logging"
	"gopkg.in/natefinch/lumberjack.v2"
)

var myLog = logging.MustGetLogger("-")

// Info logs ...
var (
	Info    = myLog.Info
	Notice  = myLog.Notice
	Debug   = myLog.Debug
	Warning = myLog.Warning
	Error   = myLog.Error

	Infof    = myLog.Infof
	Noticef  = myLog.Noticef
	Debugf   = myLog.Debugf
	Warningf = myLog.Warningf
	Errorf   = myLog.Errorf
)

func init() {

	// Example format string. Everything except the message has a custom color
	// which is dependent on the log level. Many fields have a custom output
	// formatting too, eg. the time returns the hour down to the milli second.
	format := logging.MustStringFormatter(
		`%{color}%{time} [%{level}] %{module} %{shortfile} > %{color:reset} %{message}`,
	)
	/* TODO bug: Aliyun log print in a wrong path */

	// For demo purposes, create two backend for os.Stderr.
	console := logging.NewLogBackend(os.Stderr, "", 0)
	fileBackend := logging.NewLogBackend(&lumberjack.Logger{
		Filename: "./tmp/logs/gin.log",
		MaxSize:  2,    // megabytes
		Compress: true, // disabled by default
	}, "", 0)

	// For messages written to console we want to add some additional
	// information to the output, including the used log level and the name of
	// the function.
	consoleFormatter := logging.NewBackendFormatter(console, format)
	fileOutputFormatter := logging.NewBackendFormatter(fileBackend, format)

	// Set the backends to be used.
	logging.SetBackend(consoleFormatter, fileOutputFormatter)
}
