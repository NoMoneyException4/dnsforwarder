package logger

import (
	"os"

	"github.com/codebear4/dnsforwarder/conf"

	"github.com/op/go-logging"
)

var (
	Logger *logging.Logger
	format = logging.MustStringFormatter(
		`%{color} %{longfunc} %{level} %{time:2006-01-02 15:04:05}:%{color:reset} %{message}`,
	)
	fileFormat = logging.MustStringFormatter(
		`%{longfunc} %{level} %{time:2006-01-02 15:04:05}: %{message}`,
	)
)

func InitLogger() {
	Logger = logging.MustGetLogger("dnsforwarder")
	loggers := make([]logging.Backend, 0)

	if conf.Conf.Loggers.Console.Enable {
		consoleBackend := logging.NewLogBackend(os.Stderr, "DnsForwarder", 0)
		consoleFormatter := logging.NewBackendFormatter(consoleBackend, format)
		loggers = append(loggers, consoleFormatter)
	}

	if conf.Conf.Loggers.File.Enable {
		logfile, err := os.OpenFile(conf.Conf.Loggers.File.Path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
		if err != nil {
			panic(err)
		}
		fileBackend := logging.NewLogBackend(logfile, "", 0)
		fileFormatter := logging.NewBackendFormatter(fileBackend, fileFormat)
		fileBackendLeveld := logging.AddModuleLevel(fileFormatter)
		level, err := logging.LogLevel(conf.Conf.Loggers.File.Level)
		if err != nil {
			panic(err)
		}
		fileBackendLeveld.SetLevel(level, "")
		loggers = append(loggers, fileBackendLeveld)
	}

	logging.SetBackend(loggers...)
}
