package log

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"
)

var (
	Mtx = &sync.RWMutex{}
)

type LogEntry struct {
	Timestamp time.Time `json:"timestamp"`
	Function  string    `json:"function"`
	Message   string    `json:"message"`
	Error     string    `json:"error,omitempty"`
}

// LogW writes a log entry to a JSON file
func LogW(fn string, message string, err error) {
	logEntry := LogEntry{
		Timestamp: time.Now(),
		Function:  fn,
		Message:   message,
	}

	if err != nil {
		logEntry.Error = err.Error()
	}

	file, fileErr := os.OpenFile("resources/logs/logs.art", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if fileErr != nil {
		fmt.Printf("Error opening log file: %v\n", fileErr)
		return
	}
	defer file.Close()

	Mtx.Lock()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if encodeErr := encoder.Encode(logEntry); encodeErr != nil {
		fmt.Printf("Error encoding log entry: %v\n", encodeErr)
	}
	Mtx.Unlock()
}

// LogR reads logs from the JSON file based on the function name
// if fn is defined as ALL, then return all logs
func LogR(fn string) []string {
	file, fileErr := os.Open("resources/logs/logs.art")
	if fileErr != nil {
		fmt.Printf("Error opening log file: %v\n", fileErr)
		return nil
	}
	defer file.Close()

	var logs []LogEntry

	Mtx.Lock()
	decoder := json.NewDecoder(file)
	for {
		var log LogEntry
		if err := decoder.Decode(&log); err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Printf("Error decoding log entry: %v\n", err)
			return nil
		}
		logs = append(logs, log)
	}
	Mtx.Unlock()

	// filter the logs by function name
	var filteredLogs []string
	for _, log := range logs {
		if fn == "ALL" || fn == log.Function {
			logStr := fmt.Sprintf("[%s] %s: %s", log.Timestamp.Format(time.RFC3339), log.Function, log.Message)
			if log.Error != "" {
				logStr += fmt.Sprintf(" (Error: %s)", log.Error)
			}
			filteredLogs = append(filteredLogs, logStr)
		}
	}

	return filteredLogs
}
