package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type ECGData struct {
	Timestamp  string  `json:"timestamp"`
	HeartRate  int     `json:"heart_rate"`
	RRInterval float64 `json:"rr_interval"`
}

func SimulateECGData() ECGData {
	return ECGData{
		Timestamp:  time.Now().UTC().Format(time.RFC3339),
		HeartRate:  rand.Intn(80) + 40,
		RRInterval: rand.Float64() + 0.5,
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading to WebSocket:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Client connected!")

	for {
		data := SimulateECGData()
		message, _ := json.Marshal(data)
		if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Println("Error sending message:", err)
			break
		}

		time.Sleep(1 * time.Second)
	}
}

func main() {
	http.HandleFunc("/ws", handleConnections)
	fmt.Println("ECG Data Sender running on ws://localhost:8080/ws")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
