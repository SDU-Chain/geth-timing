package log2

import (
	"encoding/json"
	"fmt"
	kitlog "github.com/go-kit/kit/log"
	"io"
	"os"
	"time"
)

//var logger kitlog.Logger
var syncWriter io.Writer

func InitOutputFile(outputFile string) {

	if outputFile == "" {
		_, _ = fmt.Fprintln(os.Stderr, "You must specify the log file, e.g. \"geth --timing.output=/path/to/file.txt\"")
		os.Exit(1)
	}

	var err error
	file, err := os.OpenFile(outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Failed to create log file for timing.")
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	syncWriter = kitlog.NewSyncWriter(file)

	_ = Record(map[string]interface{}{"Message": "Timing log initialized.", "Type": "Message"})
}

func Record(timingLog map[string]interface{}) error {
	timingLog["Timestamp"] = time.Now().UnixNano()

	b, err := json.Marshal(timingLog)
	if err != nil {
		return err
	}

	b = append(b, '\n')

	_, err = syncWriter.Write(b)
	if err != nil {
		return err
	}

	return nil
}
