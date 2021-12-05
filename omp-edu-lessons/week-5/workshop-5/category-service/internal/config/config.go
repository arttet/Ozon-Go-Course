package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Build information -ldflags .
var (
	version    string = "dev"
	commitHash string = "-"
)

var cfg *Config

func GetConfigInstance() Config {
	if cfg != nil {
		return *cfg
	}

	return Config{}
}

// Grpc - contains parameter address grpc.
type Grpc struct {
	Port              int    `yaml:"port"`
	MaxConnectionIdle int64  `yaml:"maxConnectionIdle"`
	Timeout           int64  `yaml:"timeout"`
	MaxConnectionAge  int64  `yaml:"maxConnectionAge"`
	Host              string `yaml:"host"`
}

// Gateway - contains parameters for grpc-gateway port
type Gateway struct {
	Port               int      `yaml:"port"`
	Host               string   `yaml:"host"`
	AllowedCORSOrigins []string `yaml:"allowedCorsOrigins"`
}

// Swagger - contains parameters for swagger port
type Swagger struct {
	Port     int    `yaml:"port"`
	Host     string `yaml:"host"`
	Filepath string `yaml:"filepath"`
}

type Telemetry struct {
	GraylogPath string `yaml:"graylogPath"`
}

// Project - contains all parameters project information.
type Project struct {
	Debug       bool   `yaml:"debug"`
	Name        string `yaml:"name"`
	Environment string `yaml:"environment"`
	ServiceName string `yaml:"serviceName"`
	Version     string
	CommitHash  string
}

type Database struct {
	DSN string `yaml:"dsn"`
}

// Config - contains all configuration parameters in config package.
type Config struct {
	Project   Project   `yaml:"project"`
	Grpc      Grpc      `yaml:"grpc"`
	Gateway   Gateway   `yaml:"gateway"`
	Swagger   Swagger   `yaml:"swagger"`
	Database  Database  `yaml:"database"`
	Telemetry Telemetry `yaml:"telemetry"`
}

// ReadConfigYML - read configurations from file and init instance Config.
func ReadConfigYML(configYML string) error {
	if cfg != nil {
		return nil
	}

	file, err := os.Open(configYML)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return err
	}

	cfg.Project.Version = version
	cfg.Project.CommitHash = commitHash

	return nil
}
