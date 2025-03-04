package config

import (
	"fmt"
	"os"
	"time"
)

const (
	ENV_DB_CONNECTION          = "DB_CONNECTION"
	ENV_JWT_SECRET             = "JWT_SECRET"
	ENV_JWT_ISSUER             = "JWT_ISSUER"
	ENV_SERVER_HOST            = "SERVER_HOST"
	ENV_SERVER_PORT            = "SERVER_PORT"
	CONF_DEFAULT_HOST          = "0.0.0.0"
	CONF_DEFAULT_PORT          = "3000"
	CONF_DEFAULT_JWT_EXPIRY    = time.Hour
	CONF_AUTH_USER_CONTEXT_KEY = "auth-user"
)

type Config struct {
	Database *DatabaseConfig
	Jwt      *JwtConfig
	Auth     *AuthConfig
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
	authConfig := NewAuthConfig()

	config := Config{
		Database: dbConfig,
		Jwt:      jwtConfig,
		Auth:     authConfig,
		Server:   serverConfig,
	}

	return &config, err
}

type DatabaseConfig struct {
	ConnectionURI string
}

func NewDatabaseConfig() (*DatabaseConfig, error) {
	connectionURI := os.Getenv(ENV_DB_CONNECTION)
	if connectionURI == "" {
		return nil, fmt.Errorf("%s is not set", ENV_DB_CONNECTION)
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
	secret := os.Getenv(ENV_JWT_SECRET)
	if secret == "" {
		return nil, fmt.Errorf("%s is not set", ENV_JWT_SECRET)
	}

	issuer := os.Getenv(ENV_JWT_ISSUER)
	if issuer == "" {
		return nil, fmt.Errorf("%s is not set", ENV_JWT_ISSUER)
	}

	config := &JwtConfig{
		Secret: secret,
		Issuer: issuer,
		Expiry: CONF_DEFAULT_JWT_EXPIRY,
	}

	return config, nil
}

type ServerConfig struct {
	Host string
	Port string
}

func NewServerConfig() *ServerConfig {
	host := os.Getenv(ENV_SERVER_HOST)
	if host == "" {
		host = CONF_DEFAULT_HOST
	}

	port := os.Getenv(ENV_SERVER_PORT)
	if port == "" {
		port = CONF_DEFAULT_PORT
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

type AuthConfig struct {
	AuthUserContextKey string
}

func NewAuthConfig() *AuthConfig {
	return &AuthConfig{
		AuthUserContextKey: CONF_AUTH_USER_CONTEXT_KEY,
	}
}
