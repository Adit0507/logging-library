package pocketlog_test

import (
	pocketlog "logging/pocketLog"
	"os"
)

func ExampleLogger_Debugf() {
	debugLogger := pocketlog.New(pocketlog.LevelDebug, os.Stdout)
	debugLogger.Debugf("Hello, %s", "world")
}