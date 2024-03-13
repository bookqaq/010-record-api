package logger

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/bookqaq/010-record-api/config"
	"github.com/bookqaq/010-record-api/utils"
)

const (
	LOGLEVELINFO    = "info"
	LOGLEVELWARNING = "warning"
	LOGLEVELERROR   = "error"
)

var Info, Warning, Error, Debug *log.Logger

func New(openDebugLog bool, loglevel string) {
	writer := GetWriter()

	switch loglevel {
	case LOGLEVELINFO:
		Info = log.New(writer, "[Info] ", log.Ldate|log.Ltime|log.Lmsgprefix)
		fallthrough
	case LOGLEVELWARNING:
		Warning = log.New(writer, "[warning] ", log.Ldate|log.Ltime|log.Lmsgprefix)
		fallthrough
	case LOGLEVELERROR:
		Error = log.New(writer, "[error] ", log.Ldate|log.Ltime|log.Lmsgprefix|log.Lshortfile)
	default:
		utils.SimulatedPanic(fmt.Sprintf("invalid loglevel setting \"%s\", exiting...\n", loglevel))
	}

	// seperate debug log setting
	debugWriter := io.Discard
	if openDebugLog {
		debugWriter = writer
	}
	Debug = log.New(debugWriter, "[debug] ", log.Ldate|log.Ltime|log.Lmsgprefix|log.Lshortfile)
}

func GetWriter() io.Writer {
	fp, err := os.OpenFile(config.Config.LogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(fmt.Errorf("log file open failed: %w", err))
	}
	return io.MultiWriter(fp, os.Stdout)
}
