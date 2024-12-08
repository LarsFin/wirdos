package logger

import (
	"fmt"
	"os"
)

type fileLogger struct {
	fileName string
}

func (l *fileLogger) debug(msg string) {
	l.write(logPrefix(debug) + msg)
}

func (l *fileLogger) info(msg string) {
	l.write(logPrefix(info) + msg)
}

func (l *fileLogger) warn(msg string) {
	l.write(logPrefix(warn) + msg)
}

func (l *fileLogger) err(err error) {
	l.write(logPrefix(erro) + fmt.Sprintf("%s", err))
}

// don't panic when an error is thrown while trying to log an error, just log the
// error message via the console instead
func (l *fileLogger) write(text string) {
	f, err := os.OpenFile(l.fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		l.logFormattedError(text, err)
		return
	}

	if _, err = f.WriteString(fmt.Sprintf("%s\n", text)); err != nil {
		l.logFormattedError(text, err)
	}

	f.Close()
}

func (l *fileLogger) logFormattedError(msg string, err error) {
	fmt.Printf("failed to log to file: %s, log message: %s, error: %s\n", l.fileName, msg, err)
}

func newFileLogger(fileName string) (*fileLogger, error) {
	return &fileLogger{fileName: fileName}, nil
}
