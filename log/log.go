package log

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

var (
	multiWriter = io.MultiWriter(os.Stdout)

	errorLog = log.New(multiWriter, "[ERROR] ", log.LstdFlags|log.Lshortfile)
	warnLog  = log.New(multiWriter, "[WARNING] ", log.LstdFlags|log.Lshortfile)
	infoLog  = log.New(multiWriter, "[INFO] ", log.LstdFlags|log.Lshortfile)
	loggers  = []*log.Logger{errorLog, warnLog, infoLog}

	mu sync.Mutex
)

var (
	Error  = errorLog.Println
	Errorf = errorLog.Printf
	Warn   = warnLog.Println
	Warnf  = warnLog.Printf
	Info   = infoLog.Println
	Infof  = infoLog.Printf
)

const (
	InfoLevel = iota
	WarnLevel
	ErrorLevel
	Disabled
)

func SetLevel(level int) {
	mu.Lock()
	defer mu.Unlock()

	for _, logger := range loggers {
		logger.SetOutput(multiWriter)
	}

	if ErrorLevel < level {
		errorLog.SetOutput(ioutil.Discard)
	}

	if WarnLevel < level {
		warnLog.SetOutput(ioutil.Discard)
	}

	if InfoLevel < level {
		infoLog.SetOutput(ioutil.Discard)
	}
}

func SetLogFile(fileName string, dst string) {
	mu.Lock()
	defer mu.Unlock()

	logFile, err := os.OpenFile(dst+fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	multiWriter = io.MultiWriter(os.Stdout, logFile)
	for _, logger := range loggers {
		logger.SetOutput(multiWriter)
	}
}
