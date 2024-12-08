package logger

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

type lglvl int

const (
	debug lglvl = iota
	info
	warn
	erro
)

func (l lglvl) String() string {
	switch l {
	case debug:
		return "DBG"
	case info:
		return "INF"
	case warn:
		return "WRN"
	case erro:
		return "ERR"
	}
	return "UNKNOWN"
}

type lgr interface {
	debug(string)
	info(string)
	warn(string)
	err(error)
}

var logger lgr
var level lglvl

type consoleLogger struct {}

func (l *consoleLogger) debug(msg string) {
	color.RGB(25, 25, 25).Println(logPrefix(debug) + msg)
}

func (l *consoleLogger) info(msg string) {
	color.White(logPrefix(info) + msg)
}

func (l *consoleLogger) warn(msg string) {
	color.Yellow(logPrefix(warn) + msg)
}

func (l *consoleLogger) err(err error) {
	color.Red(logPrefix(erro) + fmt.Sprintf("%s", err))
}

func logPrefix(level lglvl) string {
	t := time.Now()
	return t.Format("2006-01-02 15:04:05") + " [" + level.String() + "] "
}

// TODO: should determine type and level of logger based on config
func InitLogger() {
	logger = &consoleLogger{}
}

func Debug(msg string) {
	if level > debug {
		return
	}
	logger.debug(msg)
}

func Info(msg string) {
	if level > info {
		return
	}
	logger.info(msg)
}

func Warn(msg string) {
	if level > warn {
		return
	}
	logger.warn(msg)
}

func Error(err error) {
	logger.err(err)
}
