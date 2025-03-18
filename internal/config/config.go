package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var MonicaConfig *Config

// Config 存储应用配置
type Config struct {
	MonicaCookie string
	BearerToken  string
	Port         int // 新增端口配置
}

// LoadConfig 从环境变量加载配置
func LoadConfig() *Config {
	// 尝试加载 .env 文件，但不强制要求文件存在
	_ = godotenv.Load()

	// 获取端口配置，默认为8080
	port := 8080
	if portStr := os.Getenv("PORT"); portStr != "" {
		if p, err := strconv.Atoi(portStr); err == nil {
			port = p
		}
	}

	MonicaConfig = &Config{
		MonicaCookie: os.Getenv("MONICA_COOKIE"),
		BearerToken:  os.Getenv("BEARER_TOKEN"),
		Port:         port,
	}
	return MonicaConfig
}
