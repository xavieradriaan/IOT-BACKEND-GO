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
	go mqtt.StartClient()

	// Inicializa la base de datos SQLite
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

	// Crea la tabla y usuario admin si no existe
	if !dbExists {
		_, err = db.Exec(`CREATE TABLE users (username TEXT PRIMARY KEY, password TEXT);`)
		if err != nil {
			log.Fatalf("Error creando tabla: %v", err)
		}
		_, err = db.Exec(`INSERT INTO users (username, password) VALUES (?, ?)`, "admin", "admin")
		if err != nil {
			log.Fatalf("Error insertando usuario admin: %v", err)
		}
	}

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

	// Ruta principal
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Servidor backend activo. Las métricas están en /metrics.")
	})

	r.Run(":8000")
}
