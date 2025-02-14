package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

type SysInfo struct {
	CPUUsage    float64 `json:"cpu_usage"`
	MemoryTotal uint64  `json:"memory_total"`
	MemoryUsed  uint64  `json:"memory_used"`
	MemoryUsage float64 `json:"memory_usage"`
	DiskTotal   uint64  `json:"disk_total"`
	DiskUsed    uint64  `json:"disk_used"`
	DiskUsage   float64 `json:"disk_usage"`
}

type LogParser struct {
	Timestamp string `json:"timeStamp"`
	Level     string `json:"level"`
	Message   string `json:"message"`
}

type LogResponse struct {
	Error   uint64 `json:"errors"`
	Info    uint64 `json:"info"`
	Warning uint64 `json:"warning"`
}

func logParser(fileName string) (*LogResponse, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, errors.New("Error in opening file")
	}
	defer file.Close()

	logChecker := regexp.MustCompile(`(\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}) (INFO|WARN|ERROR) (.*)`)

	var logEntries []LogParser
	errorCount := 0
	warnCount := 0
	infoCount := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := logChecker.FindStringSubmatch(line)
		if len(matches) == 4 {
			entry := LogParser{
				Timestamp: matches[1],
				Level:     matches[2],
				Message:   matches[3],
			}
			logEntries = append(logEntries, entry)

			switch strings.ToUpper(entry.Level) {
			case "ERROR":
				errorCount++
			case "INFO":
				infoCount++
			case "WARNING":
				warnCount++
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &LogResponse{
		Error:   uint64(errorCount),
		Info:    uint64(infoCount),
		Warning: uint64(warnCount),
	}, nil
}

func getInfo() (*SysInfo, error) {
	cpuPercent, err := cpu.Percent(0, false)
	if err != nil {
		return nil, err
	}

	memStats, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	diskStats, err := disk.Usage("/")
	if err != nil {
		return nil, err
	}

	return &SysInfo{
		CPUUsage:    cpuPercent[0],
		MemoryTotal: memStats.Total / 1024 / 1024,
		MemoryUsed:  memStats.Used / 1024 / 1024,
		MemoryUsage: memStats.UsedPercent,
		DiskTotal:   diskStats.Total / 1024 / 1024 / 1024,
		DiskUsed:    diskStats.Used / 1024 / 1024 / 1024,
		DiskUsage:   diskStats.UsedPercent,
	}, nil
}

func home(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "home endpoint")
}

func sysInfo(w http.ResponseWriter, req *http.Request) {
	info, err := getInfo()
	if err != nil {
		http.Error(w, "Failed to get system info", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}

func logs(w http.ResponseWriter, req *http.Request) {
	logs, err := logParser("logger.log")
	if err != nil {
		http.Error(w, "Failed to get logs", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(logs)
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/api/sysInfo", sysInfo)
	http.HandleFunc("/api/logs", logs)

	http.ListenAndServe(":8080", nil)
}
