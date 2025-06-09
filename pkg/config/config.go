package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Config 전체 설정 구조체
type Config struct {
	App struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
		Port    int    `yaml:"port"`
	} `yaml:"app"`

	Logging struct {
		Level  string `yaml:"level"`
		Format string `yaml:"format"`
		Output string `yaml:"output"`
	} `yaml:"logging"`

	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Name     string `yaml:"name"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"database"`
}

// LoadConfig 설정 파일을 로드하는 함수
func LoadConfig(configPath string) (*Config, error) {
	// 설정 파일 읽기
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %v", err)
	}

	// YAML 파싱
	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("error parsing config file: %v", err)
	}

	return &config, nil
} 