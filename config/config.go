package config

import (
	_ "embed"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/goccy/go-yaml"
	"os"
	"sync"
)

type Config struct {
	Kitex struct {
		Service       string `yaml:"Service"`
		Address       string `yaml:"Address"`
		LogLevel      string `yaml:"LogLevel"`
		LogFileName   string `yaml:"LogFileName"`
		LogMaxSize    int    `yaml:"LogMaxSize"`
		LogMaxAge     int    `yaml:"LogMaxAge"`
		LogMaxBackups int    `yaml:"LogMaxBackups"`
	} `yaml:"Kitex"`
	Registry struct {
		RegistryAddress []string `yaml:"RegistryAddress"`
		UserName        string   `yaml:"Username"`
		PassWord        string   `yaml:"Password"`
	} `yaml:"Registry"`
	PostgresSQL struct {
		Host     string `yaml:"Host"`
		User     string `yaml:"User"`
		Password string `yaml:"Password"`
		DBName   string `yaml:"DBName"`
		Port     int    `yaml:"Port"`
		SSLMode  string `yaml:"SSLMode"`
	} `yaml:"PostgresSQL"`
}

var (
	instance *Config
	once     sync.Once
)

//go:embed config.yaml
var data []byte

func GetConfig() *Config {
	once.Do(func() {
		instance = loadConfig()
	})
	return instance
}

func LogLevel() klog.Level {
	level := GetConfig().Kitex.LogLevel
	switch level {
	case "trace":
		return klog.LevelTrace
	case "debug":
		return klog.LevelDebug
	case "info":
		return klog.LevelInfo
	case "notice":
		return klog.LevelNotice
	case "warn":
		return klog.LevelWarn
	case "error":
		return klog.LevelError
	case "fatal":
		return klog.LevelFatal
	default:
		return klog.LevelInfo
	}
}
func loadConfig() *Config {
	var conf Config

	err := yaml.Unmarshal(data, &conf)
	if err != nil {
		panic(err)
	}
	listenOnEnv := os.Getenv("LISTEN_ON")
	if listenOnEnv != "" {
		conf.Kitex.Address = listenOnEnv
	}

	// Registry
	registryAddrEnv := os.Getenv("REGISTRY_ADDR")
	if registryAddrEnv != "" {
		conf.Registry.RegistryAddress = []string{registryAddrEnv}
	}
	registryUserNameEnv := os.Getenv("REGISTRY_USERNAME")
	if registryUserNameEnv != "" {
		conf.Registry.UserName = ""
	}
	registryPasswordEnv := os.Getenv("REGISTRY_PASSWORD")
	if registryPasswordEnv != "" {
		conf.Registry.PassWord = registryPasswordEnv
	}

	return &conf
}
