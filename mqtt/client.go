package mqtt

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"iot-backend/metrics"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
)

var db *sql.DB

func StartClient(database *sql.DB) {
	db = database
	_ = godotenv.Load()

	host := os.Getenv("MQTT_HOST")
	port := os.Getenv("MQTT_PORT")
	if host == "" {
		host = "mqtt-broker"
	}
	if port == "" {
		port = "1883"
	}
	broker := fmt.Sprintf("tcp://%s:%s", host, port)

	opts := mqtt.NewClientOptions().AddBroker(broker).SetClientID("iot-backend")
	opts.OnConnect = func(c mqtt.Client) {
		log.Println("[MQTT] Conectado")
		c.Subscribe("esp32/metrics", 0, onMetrics)
		c.Subscribe("iot/biometric", 0, onBiometric)
	}

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Error conectando MQTT: %v", token.Error())
	}
}

func onMetrics(client mqtt.Client, msg mqtt.Message) {
	payload := string(msg.Payload())
	log.Printf("[MQTT] Metrics Topic: %s | Payload: %s", msg.Topic(), payload)

	parts := strings.Split(payload, ";")
	for _, part := range parts {
		if strings.Contains(part, "=") {
			kv := strings.Split(part, "=")
			key := strings.TrimSpace(kv[0])
			val := strings.TrimSpace(kv[1])

			switch key {
			case "latencia":
				val = strings.TrimSuffix(val, "ms")
				if latency, err := strconv.ParseFloat(val, 64); err == nil {
					metrics.LatencyGauge.Set(latency)
				}
			case "uptime":
				if uptime, err := strconv.Atoi(val); err == nil {
					metrics.UptimeGauge.Set(float64(uptime))
				}
			}
		}
	}
}

func onBiometric(client mqtt.Client, msg mqtt.Message) {
	payload := string(msg.Payload())
	log.Printf("[MQTT] Biometric Topic: %s | Payload: %s", msg.Topic(), payload)

	eventType := "unknown"
	employee := "desconocido"
	eventDate := ""

	parts := strings.Split(payload, ";")
	if len(parts) == 2 {
		if strings.Contains(parts[1], "=") {
			employee = parts[0]
			eventType = strings.Split(parts[1], "=")[0]
			timePart := strings.Split(parts[1], "=")[1]
			// eventDate será YYYY-MM-DD
			eventDate = getTodayFromTimeString(timePart)
		} else {
			employee = parts[0]
			eventType = parts[1]
			eventDate = getTodayFromTimeString("")
		}
	}
	
	// Save to Prometheus metrics (existing functionality)
	metrics.RecordBiometricEvent(eventType, employee, eventDate)
	
	// Save to database (NEW functionality for persistence)
	if db != nil {
		err := saveBiometricEventToDB(employee, eventType, eventDate, payload)
		if err != nil {
			log.Printf("[ERROR] Failed to save biometric event to database: %v", err)
		} else {
			log.Printf("[DB] Saved biometric event: %s - %s - %s", employee, eventType, eventDate)
		}
	}
}

// Extrae la fecha de hoy en formato YYYY-MM-DD (puedes mejorar para extraer del payload si lo envías)
func getTodayFromTimeString(_ string) string {
	return time.Now().Format("2006-01-02")
}

// saveBiometricEventToDB saves the biometric event to the database for persistence
func saveBiometricEventToDB(employee, eventType, eventDate, rawPayload string) error {
	query := `
		INSERT INTO attendance (employee_name, event_type, event_date, raw_payload, timestamp) 
		VALUES (?, ?, ?, ?, CURRENT_TIMESTAMP)
	`
	
	_, err := db.Exec(query, employee, eventType, eventDate, rawPayload)
	return err
}
