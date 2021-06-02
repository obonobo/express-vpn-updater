package logging

import (
	"log"
	"os"

	"github.com/obonobo/express-vpn-updater/server/app/util"
)

const (
	callingAPIFormat = "Calling API: %s\n"
	insideFunction   = "Inside function: %s\n"
)

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

// Logs the request and response objects of a request handler
func (l *Logger) LogRequestAndResponse(request util.Request, handler func(util.Request) util.Response) util.Response {
	l.Println("REQUEST:", request)
	resp := handler(request)
	l.Println("RESPONSE:", resp)
	return resp
}

// Logs a message to indicate what function we are inside of
func (l *Logger) Inside(functionName string) {
	l.Printf(insideFunction, functionName)
}
