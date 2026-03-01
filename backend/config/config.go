package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Elasticsearch ElasticsearchConfig `yaml:"elasticsearch"`
}

type ServerConfig struct {
	Addr string `yaml:"addr"`
}

type ElasticsearchConfig struct {
	Host string `yaml:"host"`
}

func Load() (*Config, error) {
	// 读取配置文件
	configPath := filepath.Join("config", "config.yaml")
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	// 解析 YAML
	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
