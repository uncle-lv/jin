package test

import (
	"fmt"
	"jin/log"
	"testing"
)

func TestSetLevel(t *testing.T) {
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

func TestSetLogFile(t *testing.T) {
	log.SetLogFile("log.log", "../")
	log.Error("===== user the log with default level =====")
	printLog()
	log.Error("===== user the log with InfoLevel level =====")
	log.SetLevel(log.InfoLevel)
	printLog()
	log.Error("===== user the log with WarnLevel level =====")
	log.SetLevel(log.WarnLevel)
	printLog()
	log.Error("===== user the log with ErrorLevel level =====")
	log.SetLevel(log.ErrorLevel)
	printLog()
	log.Error("===== user the log with Disabled level =====")
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
