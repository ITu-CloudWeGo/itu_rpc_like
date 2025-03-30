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
		RegistryAddress []string `yaml:"registryAddress"`
		UserName        string   `yaml:"username"`
		PassWord        string   `yaml:"password"`
	} `yaml:"Registry"`
	PostgresSQL struct {
		Host     string `yaml:"host"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
		Port     int    `yaml:"port"`
	} `yaml:""`
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

	etcdHostEnv := os.Getenv("ETCD_HOST")
	if etcdHostEnv != "" {
		conf.Registry.RegistryAddress = etcdHostEnv
	}

	return &conf
}
