package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type LogEntry struct {
	Timestamp string
	Level     string
	Message   string
}

func main() {
	file, err := os.Open("logfile.log")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	logPattern := regexp.MustCompile(`(\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}) (INFO|WARN|ERROR) (.*)`)

	var logEntries []LogEntry
	errorCount := 0
	warnCount := 0
	infoCount := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := logPattern.FindStringSubmatch(line)
		if len(matches) == 4 {
			entry := LogEntry{
				Timestamp: matches[1],
				Level:     matches[2],
				Message:   matches[3],
			}
			logEntries = append(logEntries, entry)

			switch strings.ToUpper(entry.Level) {
			case "ERROR":
				errorCount++
			case "WARN":
				warnCount++
			case "INFO":
				infoCount++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("📊 Log Summary:")
	fmt.Printf("✅ INFO Logs: %d\n", infoCount)
	fmt.Printf("⚠️  WARN Logs: %d\n", warnCount)
	fmt.Printf("❌ ERROR Logs: %d\n", errorCount)
	fmt.Println("--------------------")

	for _, entry := range logEntries {
		fmt.Printf("[%s] [%s] %s\n", entry.Timestamp, entry.Level, entry.Message)
	}
}
