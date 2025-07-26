CREATE TABLE IF NOT EXISTS users (username TEXT PRIMARY KEY, password TEXT);
INSERT OR IGNORE INTO users (username, password) VALUES ('admin', 'admin');
INSERT OR IGNORE INTO users (username, password) VALUES ('user', 'admin');
