package goutils

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"runtime"
)

func InitLogger(dbg *bool, defaultLogger bool) *log.Logger {
	var logger = log.StandardLogger()

	if !defaultLogger {
		logger = log.New()
	}

	log.SetReportCaller(true)

	logger.SetFormatter(&log.TextFormatter{
		//DisableColors: true,
		TimestampFormat: "060102 15:04:05",
		DisableSorting:  false,
		FullTimestamp:   true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return "", fmt.Sprintf(" %s:%d", filename, f.Line)
		},

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


