package backend

import (
	"io"
	"log"
	"os"
)

//ServerLogger  prints in various formats to various outputs.
type ServerLogger struct {
	status     *log.Logger
	debug      *log.Logger
	trace      *log.Logger
	warning    *log.Logger
	alertFatal *log.Logger
}

//NewServerLogger Create an instance of type ServerLogger.
func NewServerLogger(writer io.Writer) *ServerLogger {
	logfile, err := os.OpenFile("log-output.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file:", err)
	}

	return &ServerLogger{
		status:     log.New(writer, "[INFO] ", log.Ltime),
		debug:      log.New(writer, "[DEBUG] ", log.Ldate|log.Ltime),
		warning:    log.New(io.MultiWriter(logfile, writer), "[WARNING] ", log.Ldate|log.Ltime|log.Lshortfile),
		alertFatal: log.New(io.MultiWriter(logfile, writer), "[ALERT] ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

//Helper function, for printing strings and errors passed to logger methods.
func printLog(logger interface{}, msgs []interface{}, isFatal bool) {

	if l, ok := logger.(*log.Logger); ok {
		for _, m := range msgs {
			switch msg := m.(type) {
			case string: // print any strings
				if len(msg) > 0 {
					l.Println("[ ", msg, " ]")
				}
			case error: // print any errors
				if msg != nil {
					if isFatal {
						defer l.Fatal("[FATAL] ", msg)
					} else {
						l.Println("[ERROR] ", msg)
					}
				}
			}
		}
	}
}

//LogStatus Non-verbose.
func (sl *ServerLogger) LogStatus(msgs ...interface{}) {
	printLog(sl.status, msgs, false)
}

//LogDebug Semi-verbose.
func (sl *ServerLogger) LogDebug(msgs ...interface{}) {
	printLog(sl.debug, msgs, false)
}

//LogWarning Verbose.
func (sl *ServerLogger) LogWarning(msgs ...interface{}) {
	printLog(sl.warning, msgs, false)
}

//LogFatalAlert Fatal.
func (sl *ServerLogger) LogFatalAlert(msgs ...interface{}) {
	printLog(sl.alertFatal, msgs, true)
}
