package goutils

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func InitLogger(dbg *bool, defaultLogger bool) *log.Logger {

	var logger = log.StandardLogger()

	if !defaultLogger {
		logger = log.New()
	}
	logger.SetFormatter(&log.TextFormatter{
		//DisableColors: true,
		TimestampFormat: "060102 15:04:05",
		DisableSorting:  false,
		FullTimestamp:   true,
	})
	logger.SetOutput(os.Stdout)
	if *dbg {
		logger.SetLevel(log.DebugLevel)
		logger.Debugln("debug is set")
	} else {
		logger.SetLevel(log.InfoLevel)
	}
	return logger
}
