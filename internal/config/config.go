package config

import (
    "os"
    "strconv"
)

type Config struct {
    Server   ServerConfig
    Database DatabaseConfig
    JWT      JWTConfig
    Redis    RedisConfig
}

type ServerConfig struct {
    Port string
    Host string
}

type DatabaseConfig struct {
    Host     string
    Port     int
    User     string
    Password string
    Name     string
    SSLMode  string
}

type JWTConfig struct {
    Secret     string
    Expiration string
}

type RedisConfig struct {
    URL string
}

func Load() (*Config, error) {
    port, _ := strconv.Atoi(getEnv("DB_PORT", "5432"))
    
    return &Config{
        Server: ServerConfig{
            Port: getEnv("PORT", "4000"),
            Host: getEnv("HOST", "localhost"),
        },
        Database: DatabaseConfig{
            Host:     getEnv("DB_HOST", "localhost"),
            Port:     port,
            User:     getEnv("DB_USER", "postgres"),
            Password: getEnv("DB_PASSWORD", ""),
            Name:     getEnv("DB_NAME", "ctrix_social"),
            SSLMode:  getEnv("DB_SSLMODE", "disable"),
        },
        JWT: JWTConfig{
            Secret:     getEnv("JWT_SECRET", ""),
            Expiration: getEnv("JWT_EXPIRATION", "24h"),
        },
        Redis: RedisConfig{
            URL: getEnv("REDIS_URL", "redis://localhost:6379"),
        },
    }, nil
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}
