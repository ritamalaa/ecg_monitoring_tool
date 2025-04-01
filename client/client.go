package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/websocket"
)

type ECGData struct {
	Timestamp  string  `json:"timestamp"`
	HeartRate  int     `json:"heart_rate"`
	RRInterval float64 `json:"rr_interval"`
}

var logFile *os.File

func logToFile(alert string) {
	if logFile == nil {
		var err error
		logFile, err = os.OpenFile("ecg_alerts.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println("Error opening log file:", err)
			return
		}
	}
	logEntry := fmt.Sprintf("%s - %s\n", time.Now().Format("2006-01-02 15:04:05"), alert)
	if _, err := logFile.WriteString(logEntry); err != nil {
		log.Println("Error writing to log file:", err)
	}
}

func processECGData(data ECGData) {
	alert := ""
	if data.HeartRate > 100 {
		alert = fmt.Sprintf("ALERT: Tachycardia detected! Heart rate: %d BPM", data.HeartRate)
	} else if data.HeartRate < 60 {
		alert = fmt.Sprintf("ALERT: Bradycardia detected! Heart rate: %d BPM", data.HeartRate)
	}
	if data.RRInterval < 0.6 || data.RRInterval > 1.2 {
		alert = fmt.Sprintf("ALERT: Possible Arrhythmia detected! RR Interval: %.2f", data.RRInterval)
	}

	if alert != "" {
		log.Println(alert)
		logToFile(alert)
	}
}

func connectWebSocket() {
	url := "ws://localhost:8080/ws"
	for {
		conn, _, err := websocket.DefaultDialer.Dial(url, nil)
		if err != nil {
			log.Println("WebSocket connection failed:", err)
			time.Sleep(2 * time.Second)
			continue
		}

		log.Println("Connected to ECG Data Sender!")
		go func() {
			for {
				_, message, err := conn.ReadMessage()
				if err != nil {
					log.Println("Error reading message:", err)
					conn.Close()
					break
				}

				var data ECGData
				if err := json.Unmarshal(message, &data); err != nil {
					log.Println("Error decoding JSON:", err)
					continue
				}

				processECGData(data)
			}
		}()

		time.Sleep(2 * time.Second)
	}
}

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go connectWebSocket()

	defer func() {
		if logFile != nil {
			logFile.Close()
			log.Println("Log file closed.")
		}
	}()

	<-sigChan
	log.Println("Shutting down ECG receiver...")
}
