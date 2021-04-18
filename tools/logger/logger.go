package logger

import (
	"log"
	"os"
)

var filePath = "../araneus.log"

type Logger struct {
	InfoLogger *log.Logger
	ErrorLogger *log.Logger
}

func NewLogger(file *os.File) Logger {
	return Logger{
		InfoLogger:  log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		ErrorLogger: log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func OpenLogFile() (*os.File, error) {
	return os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
}
