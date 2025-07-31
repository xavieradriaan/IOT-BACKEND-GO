CREATE TABLE IF NOT EXISTS users (username TEXT PRIMARY KEY, password TEXT);
INSERT OR IGNORE INTO users (username, password) VALUES ('admin', 'admin');
INSERT OR IGNORE INTO users (username, password) VALUES ('user', 'admin');

-- Attendance/Biometric Events Table
CREATE TABLE IF NOT EXISTS attendance (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    employee_name TEXT NOT NULL,
    event_type TEXT NOT NULL,
    event_date DATE NOT NULL,
    timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
    device_id TEXT DEFAULT 'ESP32',
    raw_payload TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
