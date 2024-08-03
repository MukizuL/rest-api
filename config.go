package main

import (
	"bytes"
	"log"
	"os"
)

type Config struct {
	Port       string
	DBUser     string
	DBPassword string
	DBHost	   string
	DBName     string
	DBSSLMode  string
	JWTSecret  string
}

//Initialize Config with env variables.
//If variable is not present, uses fallback
func NewConfig() Config {
	return Config{
		Port: getEnv("PORT", "5432"),
		DBUser: getEnv("DB_User", "postgres"),
		DBPassword: getEnv("DB_Password", ""),
		DBHost: getEnv("DB_Host", "localhost"),
		DBName: getEnv("DB_Name", "rest-api"),
		DBSSLMode: getEnv("DB_SSLMode", "require"),
		JWTSecret: getEnv("JWT_Secret", "jwterror"),
	}
}

//Gets env variable via os.LookupEnv
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

//Formats Config to usable DSN string
func (cfg *Config) FormatDSN() string {
	var buf bytes.Buffer

	if len(cfg.DBUser) > 0 {
		buf.WriteString("user=")
		buf.WriteString(cfg.DBUser)
	}

	if len(cfg.DBPassword) > 0 {
		buf.WriteByte(' ')

		buf.WriteString("password=")
		buf.WriteString(cfg.DBPassword)
	}

	if len(cfg.DBName) > 0 {
		buf.WriteByte(' ')

		buf.WriteString("dbname=")
		buf.WriteString(cfg.DBName)
	}

	if len(cfg.DBSSLMode) > 0 {
		buf.WriteByte(' ')

		buf.WriteString("sslmode=")
		switch(cfg.DBSSLMode) {
		case "disable", "require ", "verify-ca", "verify-full": buf.WriteString(cfg.DBSSLMode)
		default: log.Println("Wrong sslmode, defaulting to 'require'")
		}
		
	}
	
	return buf.String()
} 