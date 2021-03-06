package goutils

import (
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func InstallShutdownHandler(shutdownFunc func()) {
	// install shutdown handler
	killChan := make(chan os.Signal, 1)
	signal.Notify(killChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-killChan
		log.Infof("received shutdown signal: %s", sig)
		shutdownFunc() // run shutdown function
	}()
}
