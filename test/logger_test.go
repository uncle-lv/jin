package test

import (
	"fmt"
	"jin/log"
	"testing"
)

func TestLogger(t *testing.T) {
	fmt.Println("===== user the log with default level =====")
	printLog()
	fmt.Println("===== set the log to InfoLevel =====")
	log.SetLevel(log.InfoLevel)
	printLog()
	fmt.Println("===== set the log to WarnLevel =====")
	log.SetLevel(log.WarnLevel)
	printLog()
	fmt.Println("===== set the log to ErrorLevel =====")
	log.SetLevel(log.ErrorLevel)
	printLog()
	fmt.Println("===== set the log to Disabled =====")
	log.SetLevel(log.Disabled)
	printLog()
}

func printLog() {
	log.Info("This is a Info log")
	log.Warn("This is a Warn log")
	log.Error("This is a Error log")
	log.Infof("This is a %s log", "Infof")
	log.Warnf("This is a %s log", "Warnf")
	log.Errorf("This is a %s log", "Errorf")
}
