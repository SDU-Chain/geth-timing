package log2

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"geth-timing/go-logging"
)

var logger = logging.MustGetLogger("timing")
var format = logging.MustStringFormatter(
	`%{message}`,
)

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

	backend := logging.NewLogBackend(file, "", 0)
	formatter := logging.NewBackendFormatter(backend, format)
	logging.SetBackend(formatter)

	_ = Record(map[string]interface{}{"Message": "Timing log initialized.", "Type": "Message"})
	//_, err = file.WriteString("Timing log initialized.\n")
	//if err != nil {
	//	_, _ = fmt.Fprintln(os.Stderr, "Failed to write log file for timing.")
	//	_, _ = fmt.Fprintln(os.Stderr, err)
	//	os.Exit(1)
	//}
	//err = file.Sync()
	//if err != nil {
	//	_, _ = fmt.Fprintln(os.Stderr, "Failed to write log file for timing.")
	//	_, _ = fmt.Fprintln(os.Stderr, err)
	//	os.Exit(1)
	//}
}

func Record(timingLog map[string]interface{}) error {
	timingLog["Timestamp"] = time.Now().UnixNano()

	b, err := json.Marshal(timingLog)
	if err != nil {
		return err
	}

	logger.Info(string(b))
	return nil
	//b = append(b, '\n')
	//
	//_, err = file.Write(b)
	//if err != nil {
	//	return err
	//}
	//
	//err = file.Sync()
	//return err
}
