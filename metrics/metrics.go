package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	BiometricCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "biometric_events_total",
			Help: "Total de eventos biométricos por tipo",
		},
		[]string{"event_type"},
	)

	BiometricCounterByDate = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "biometric_events_by_date_total",
			Help: "Total de eventos biométricos por tipo, fecha y empleado",
		},
		[]string{"event_type", "date", "employee"},
	)

	LatencyGauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "esp32_latency_ms",
			Help: "Latencia del ESP32 en milisegundos",
		},
	)

	UptimeGauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "esp32_uptime_seconds",
			Help: "Uptime del ESP32 en segundos",
		},
	)

	// Total de marcaciones por empleado
	EmployeeMarkCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "employee_marks_total",
			Help: "Total de marcaciones por empleado",
		},
		[]string{"employee"},
	)

	// Hora de la última marcación por empleado y tipo de evento (epoch)
	LastMarkGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "employee_last_mark_epoch",
			Help: "Epoch de la última marcación por empleado y tipo de evento",
		},
		[]string{"employee", "event_type"},
	)
)

func init() {
	prometheus.MustRegister(BiometricCounter)
	prometheus.MustRegister(BiometricCounterByDate)
	prometheus.MustRegister(LatencyGauge)
	prometheus.MustRegister(UptimeGauge)
	prometheus.MustRegister(EmployeeMarkCounter)
	prometheus.MustRegister(LastMarkGauge)
}

func RecordBiometricEvent(eventType, employee, date string) {
	if employee == "" {
		employee = "unknown"
	}
	if date == "" {
		date = time.Now().Format("2006-01-02")
	}
	BiometricCounter.WithLabelValues(eventType).Inc()
	BiometricCounterByDate.WithLabelValues(eventType, date, employee).Inc()

	// Incrementa el contador total por empleado
	EmployeeMarkCounter.WithLabelValues(employee).Inc()

	// Actualiza la hora de la última marcación por empleado y tipo de evento (epoch)
	LastMarkGauge.WithLabelValues(employee, eventType).Set(float64(time.Now().Unix()))
}
