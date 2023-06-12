package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

type KV map[string]interface{}

func Init() {
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logrus.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	logrus.SetLevel(logrus.DebugLevel)
}
