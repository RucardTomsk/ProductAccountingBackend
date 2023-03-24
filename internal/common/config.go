package common

import (
	"github.com/gin-contrib/cors"
	"net/http"
	"time"
)

// ServerConfig configures gin server.
type ServerConfig struct {
	Host string
	Port string

	GinMode string

	Limits     []string // TODO(a.s.zorkin): not implemented
	Operations map[string]string
}

// DatabaseConfig stores DB credentials.
type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

// KafkaReaderConfig stores credentials and config for kafka.Reader.
type KafkaReaderConfig struct {
	Auth struct {
		Username string
		Password string
	}
	Brokers []string
	GroupID string
	Topic   string
}

// KafkaWriterConfig stores credentials and config for kafka.Writer.
type KafkaWriterConfig struct {
	Auth struct {
		Username string
		Password string
	}
	Brokers []string
	Topic   string
}

// MinioConfig is used to connect to minio (s3).
type MinioConfig struct {
	UseMocks  bool
	Endpoint  string
	AccessKey string
	SecretKey string
	Token     string
	UseSSL    bool
}

const (
	defaultHost     = "localhost:8080"
	defaultBasePath = "/api/v1"
)

var defaultSchemes = []string{"http"}

// SwaggerConfig configures swaggo/swag.
type SwaggerConfig struct {
	Title       string
	Description string
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
}

// Neo4jConfig stores DB credentials.
type Neo4jConfig struct {
	Host         string
	Port         string
	User         string
	Password     string
	DataBaseName string
}

// NewSwaggerConfig returns *SwaggerConfig with preconfigured fields.
func NewSwaggerConfig(title, description, version string) *SwaggerConfig {
	return &SwaggerConfig{
		Title:       title,
		Description: description,
		Version:     version,
		Host:        defaultHost,
		BasePath:    defaultBasePath,
		Schemes:     defaultSchemes,
	}
}

// DefaultCorsConfig returns cors.Config with very permissive policy.
func DefaultCorsConfig() cors.Config {
	return cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodOptions},
		AllowHeaders:  []string{"Origin", "X-Requested-With", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        12 * time.Hour,
	}
}

// DataProcessingConfig configures default sort, order and pagination parameters.
type DataProcessingConfig struct {
	DefaultSortField string
	DefaultSortOrder string
	DefaultLimit     int
}

// NewDataProcessingConfig returns *DataProcessingConfig with preconfigured fields.
func NewDataProcessingConfig(
	defaultSortField string,
	defaultSortOrder string,
	defaultLimit int,
) *DataProcessingConfig {
	return &DataProcessingConfig{
		DefaultSortField: defaultSortField,
		DefaultSortOrder: defaultSortOrder,
		DefaultLimit:     defaultLimit,
	}
}

// RunnersConfig is a common configuration for runners and executors.
// It contains paths to common files (sources, tests, checkers, etc).
type RunnersConfig struct {
	UseMocks bool

	TestsTempFolder       string
	SourcesTempFolder     string
	ConfigsTempFolder     string
	CheckersTempFolder    string
	CompilationTempFolder string
	OutputsTempFolder     string

	Docker DockerConfig
}

// DockerConfig configures docker client.
type DockerConfig struct {
	Version string
}
