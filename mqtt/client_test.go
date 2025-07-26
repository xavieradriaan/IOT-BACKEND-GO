package mqtt

import (
	"strings"
	"testing"
)

func TestBiometricParsing(t *testing.T) {
	tests := []struct {
		payload       string
		wantEmployee  string
		wantEventType string
	}{
		{"juan.perez;entrada=22:17:39", "juan.perez", "entrada"},
		{"empleadoB;salida", "empleadoB", "salida"},
		{"malformado", "desconocido", "unknown"},
	}

	for _, tt := range tests {
		employee, eventType := parseBiometricPayload(tt.payload)
		if employee != tt.wantEmployee || eventType != tt.wantEventType {
			t.Errorf("Payload %q -> got (%q,%q); want (%q,%q)",
				tt.payload, employee, eventType, tt.wantEmployee, tt.wantEventType)
		}
	}
}

// función interna para testear parsing sin necesidad de MQTT
func parseBiometricPayload(payload string) (string, string) {
	eventType := "unknown"
	employee := "desconocido"

	parts := strings.Split(payload, ";")
	if len(parts) == 2 {
		if strings.Contains(parts[1], "=") {
			employee = parts[0]
			eventType = strings.Split(parts[1], "=")[0]
		} else {
			employee = parts[0]
			eventType = parts[1]
		}
	}
	return employee, eventType
}
