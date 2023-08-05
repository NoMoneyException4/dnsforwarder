package main

import (
	"os"

	"github.com/op/go-logging"
)

var (
	//Logger Global logger
	Logger *logging.Logger
	format = logging.MustStringFormatter(
		`%{color}%{time:2006/01/02 15:04:05}:%{longfunc} %{level}: %{message} %{color:reset}`,
	)
	fileFormat = logging.MustStringFormatter(
		`%{longfunc} %{level} %{time:2006/01/02 15:04:05}: %{message}`,
	)
)

// InitLogger Init Logger
func InitLogger() {
	Logger = logging.MustGetLogger("dnsforwarder")
	var loggers []logging.Backend

	if Conf.Loggers.Console.Enable {
		consoleBackend := logging.NewLogBackend(os.Stderr, "", 0)
		consoleFormatter := logging.NewBackendFormatter(consoleBackend, format)
		consoleBackendLeveled := logging.AddModuleLevel(consoleFormatter)
		level, err := logging.LogLevel(Conf.Loggers.Console.Level)
		if err != nil {
			panic(err)
		}
		consoleBackendLeveled.SetLevel(level, "")
		loggers = append(loggers, consoleBackendLeveled)
	}

	if Conf.Loggers.File.Enable {
		logfile, err := os.OpenFile(Conf.Loggers.File.Path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
		if err != nil {
			panic(err)
		}
		fileBackend := logging.NewLogBackend(logfile, "", 0)
		fileFormatter := logging.NewBackendFormatter(fileBackend, fileFormat)
		fileBackendLeveld := logging.AddModuleLevel(fileFormatter)
		level, err := logging.LogLevel(Conf.Loggers.File.Level)
		if err != nil {
			panic(err)
		}
		fileBackendLeveld.SetLevel(level, "")
		loggers = append(loggers, fileBackendLeveld)
	}

	logging.SetBackend(loggers...)
}
