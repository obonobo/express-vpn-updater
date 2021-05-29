package logging

import (
	"log"
	"os"
)

const callingAPIFormat = "Calling API: %s\n"

var (
	loggerOut    = os.Stderr
	loggerPrefix = "<><><> "
	loggerFlags  = log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile
)

// A custom logger - enhanced with extra powers
type Logger struct {
	*log.Logger
}

// Creates a new Logger - enhanced with some custom methods
func New() *Logger {
	return &Logger{Logger: log.New(loggerOut, loggerPrefix, loggerFlags)}
}

// Logs a single API call along with some extra messages if you choose
func (l *Logger) LogApiCall(name string, additionalMessages ...interface{}) {
	l.Printf(callingAPIFormat, name)
	for _, msg := range additionalMessages {
		l.Println(msg)
	}
}
