package goRest

import (
	log "github.com/sirupsen/logrus"
	"os"
)

type Logger struct{}

func (l *Logger) InitLog() {
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	log.SetLevel(log.TraceLevel)
}
