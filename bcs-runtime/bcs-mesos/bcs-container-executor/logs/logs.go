

// Package logs xxx
package logs

import (
	"log"
	"os"
	"sync"
)

var stdLogger *log.Logger
var errLogger *log.Logger
var once sync.Once

func init() {
	once.Do(func() {
		// init conLogger only once
		errLogger = log.New(os.Stderr, "", log.LstdFlags)
		stdLogger = log.New(os.Stdout, "", log.LstdFlags)
	})
}

// Fatal is equivalent to Print() followed by a call to os.Exit(1).
func Fatal(v ...interface{}) {
	errLogger.Fatal(v...)
}

// Fatalf is equivalent to Printf() followed by a call to os.Exit(1).
func Fatalf(format string, v ...interface{}) {
	errLogger.Fatalf(format, v...)
}

// Fatalln is equivalent to Println() followed by a call to os.Exit(1).
func Fatalln(v ...interface{}) {
	errLogger.Fatalln(v...)
}

// Info is equivalent to Print()
func Info(v ...interface{}) {
	stdLogger.Print(v...)
}

// Infof is equivalent to Printf()
func Infof(format string, v ...interface{}) {
	stdLogger.Printf(format, v...)
}

// Infoln is equivalent to Println()
func Infoln(v ...interface{}) {
	stdLogger.Println(v...)
}

// Error is equivalent to Print()
func Error(v ...interface{}) {
	stdLogger.Print(v...)
}

// Errorf is equivalent to Printf()
func Errorf(format string, v ...interface{}) {
	stdLogger.Printf(format, v...)
}

// Errorln is equivalent to Println()
func Errorln(v ...interface{}) {
	stdLogger.Println(v...)
}
