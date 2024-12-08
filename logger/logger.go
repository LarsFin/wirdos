package logger

import (
	"fmt"
	"time"
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

func logPrefix(level lglvl) string {
	t := time.Now()
	return t.Format("2006-01-02 15:04:05") + " [" + level.String() + "] "
}

func InitLogger(logMethod string) error {
	var err error

	switch logMethod {
	case "file":
		logger, err = newFileLogger("wirdos.log")
	case "console":
		logger = &consoleLogger{}
	default:
		err = fmt.Errorf("unknown log method '%s'", logMethod)
	}

	if err != nil {
		return err
	}

	return nil
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
