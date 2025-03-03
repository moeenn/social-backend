package config

import (
	"fmt"
	"os"
	"time"
)

type Config struct {
	Database *DatabaseConfig
	Jwt      *JwtConfig
	Server   *ServerConfig
}

func NewConfig() (*Config, error) {
	dbConfig, err := NewDatabaseConfig()
	if err != nil {
		return nil, err
	}

	jwtConfig, err := NewJwtConfig()
	if err != nil {
		return nil, err
	}

	serverConfig := NewServerConfig()

	config := Config{
		Database: dbConfig,
		Jwt:      jwtConfig,
		Server:   serverConfig,
	}

	return &config, err
}

type DatabaseConfig struct {
	ConnectionURI string
}

func NewDatabaseConfig() (*DatabaseConfig, error) {
	connectionURI := os.Getenv("DB_CONNECTION")
	if connectionURI == "" {
		return nil, fmt.Errorf("DB_CONNECTION is not set")
	}

	return &DatabaseConfig{
		ConnectionURI: connectionURI,
	}, nil
}

type JwtConfig struct {
	Secret string
	Issuer string
	Expiry time.Duration
}

func NewJwtConfig() (*JwtConfig, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return nil, fmt.Errorf("JWT_SECRET is not set")
	}

	issuer := os.Getenv("JWT_ISSUER")
	if issuer == "" {
		return nil, fmt.Errorf("JWT_ISSUER is not set")
	}

	config := &JwtConfig{
		Secret: secret,
		Issuer: issuer,
		Expiry: time.Hour,
	}

	return config, nil
}

type ServerConfig struct {
	Host string
	Port string
}

func NewServerConfig() *ServerConfig {
	host := os.Getenv("SERVER_HOST")
	if host == "" {
		host = "0.0.0.0"
	}

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "3000"
	}

	config := ServerConfig{
		Host: host,
		Port: port,
	}

	return &config
}

func (c *ServerConfig) Address() string {
	return c.Host + ":" + c.Port
}
