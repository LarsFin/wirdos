package logger

import (
	"fmt"

	"github.com/fatih/color"
)

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
