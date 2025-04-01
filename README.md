# ECG WebSocket Simulator

This project consists of a WebSocket-based ECG data simulator and receiver. The server generates simulated ECG data and sends it to clients via WebSockets. The client receives the data, processes it, and logs any detected anomalies.

## Features
- **WebSocket Server**: Sends simulated ECG data at regular intervals.
- **WebSocket Client**: Connects to the server, processes the ECG data, and logs alerts for abnormal conditions.
- **Real-time Alerts**: Detects Tachycardia, Bradycardia, and Arrhythmia.
- **Logging**: Alerts are saved to a file (`ecg_alerts.log`).

## How to Run  
Follow these steps to run the ECG WebSocket Simulator:

### **1. Start the WebSocket Server**  
Run the following command to start the server:
```bash
go run ./server/server.go
```
The server will start and listen on ws://localhost:8080/ws.
It will generate simulated ECG data every second and send it to connected clients.

### **2. Start the WebSocket Client**  
In a separate terminal, run:
```bash
go run client/client.go
```
The client will connect to the server and begin receiving ECG data.
If anomalies are detected (e.g., Tachycardia, Bradycardia, Arrhythmia), they will be logged in the console and ecg_alerts.log.

### **3. Test the ECG Processing Logic (Optional)**  

To run unit tests for ECG data processing:

```bash
go test -v ./client/
```
This will execute ecg_test.go and verify that the ECG processing logic works correctly.




