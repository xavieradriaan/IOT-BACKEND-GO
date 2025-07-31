package main

import (
	"database/sql"
	"iot-backend/mqtt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Inicializa la base de datos SQLite primero
	dbPath := "users.db"
	dbExists := false
	if _, err := os.Stat(dbPath); err == nil {
		dbExists = true
	}
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("Error abriendo la base de datos: %v", err)
	}
	defer db.Close()

	// Ejecuta el script de inicialización si la base de datos no existía
	if !dbExists {
		initSQL, err := os.ReadFile("init_users.sql")
		if err != nil {
			log.Fatalf("Error leyendo init_users.sql: %v", err)
		}
		
		// Ejecuta todo el script SQL
		_, err = db.Exec(string(initSQL))
		if err != nil {
			log.Fatalf("Error ejecutando init_users.sql: %v", err)
		}
		log.Println("Database initialized with users and attendance tables")
	}

	// Inicia el cliente MQTT con la conexión a la base de datos
	go mqtt.StartClient(db)

	r := gin.Default()
	// Configuración CORS para permitir peticiones desde el frontend
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Login con SQLite
	r.POST("/login", func(c *gin.Context) {
		var json struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(400, gin.H{"error": "Datos inválidos"})
			return
		}
		var dbPassword string
		err := db.QueryRow("SELECT password FROM users WHERE username = ?", json.Username).Scan(&dbPassword)
		if err != nil || dbPassword != json.Password {
			c.JSON(401, gin.H{"error": "Credenciales incorrectas"})
			return
		}
		c.JSON(200, gin.H{"token": "demo-token"})
	})

	// Ruta pública /metrics para Prometheus
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// API endpoints para datos de asistencia
	r.GET("/api/attendance", func(c *gin.Context) {
		var attendances []map[string]interface{}
		
		query := `
			SELECT id, employee_name, event_type, event_date, timestamp, device_id, raw_payload 
			FROM attendance 
			ORDER BY timestamp DESC 
			LIMIT 100
		`
		
		rows, err := db.Query(query)
		if err != nil {
			c.JSON(500, gin.H{"error": "Error querying attendance data"})
			return
		}
		defer rows.Close()
		
		for rows.Next() {
			var id int
			var employeeName, eventType, eventDate, timestamp, deviceId, rawPayload string
			
			err := rows.Scan(&id, &employeeName, &eventType, &eventDate, &timestamp, &deviceId, &rawPayload)
			if err != nil {
				continue
			}
			
			attendances = append(attendances, map[string]interface{}{
				"id":            id,
				"employee_name": employeeName,
				"event_type":    eventType,
				"event_date":    eventDate,
				"timestamp":     timestamp,
				"device_id":     deviceId,
				"raw_payload":   rawPayload,
			})
		}
		
		c.JSON(200, gin.H{
			"data":  attendances,
			"count": len(attendances),
		})
	})
	
	// Endpoint para estadísticas de asistencia
	r.GET("/api/attendance/stats", func(c *gin.Context) {
		var stats map[string]interface{} = make(map[string]interface{})
		
		// Total de registros
		var totalCount int
		db.QueryRow("SELECT COUNT(*) FROM attendance").Scan(&totalCount)
		
		// Registros de hoy
		var todayCount int
		db.QueryRow("SELECT COUNT(*) FROM attendance WHERE event_date = date('now')").Scan(&todayCount)
		
		// Empleados únicos
		var uniqueEmployees int
		db.QueryRow("SELECT COUNT(DISTINCT employee_name) FROM attendance").Scan(&uniqueEmployees)
		
		stats["total_records"] = totalCount
		stats["today_records"] = todayCount
		stats["unique_employees"] = uniqueEmployees
		
		c.JSON(200, stats)
	})

	// Ruta principal
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Servidor backend activo. Las métricas están en /metrics.")
	})

	r.Run(":8000")
}
