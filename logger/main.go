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

const (
	NUMLOGLEVELINFO = iota
	NUMLOGLEVELWARNING
	NUMLOGLEVELERROR
)

var Info, Warning, Error, Debug *log.Logger

func New(openDebugLog bool, loglevel string) {
	writer := GetWriter()

	// decide loglevel
	loglevelValue := NUMLOGLEVELWARNING
	switch loglevel {
	case LOGLEVELINFO:
		loglevelValue = NUMLOGLEVELINFO
	case LOGLEVELWARNING:
		loglevelValue = NUMLOGLEVELWARNING
	case LOGLEVELERROR:
		loglevelValue = NUMLOGLEVELERROR
	default:
		utils.SimulatedPanic(fmt.Sprintf("invalid loglevel setting \"%s\", exiting...\n", loglevel))
	}

	// make sure writer and logger never be nil
	// bad code, but cant find a better way to set loglevel
	infoWriter, warningWriter, errorWriter := io.Discard, io.Discard, io.Discard
	if loglevelValue <= NUMLOGLEVELINFO {
		infoWriter = writer
	}
	if loglevelValue <= NUMLOGLEVELWARNING {
		warningWriter = writer
	}
	if loglevelValue <= NUMLOGLEVELERROR {
		errorWriter = writer
	}

	Info = log.New(infoWriter, "[Info] ", log.Ldate|log.Ltime|log.Lmsgprefix)
	Warning = log.New(warningWriter, "[warning] ", log.Ldate|log.Ltime|log.Lmsgprefix)
	Error = log.New(errorWriter, "[error] ", log.Ldate|log.Ltime|log.Lmsgprefix|log.Lshortfile)

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
