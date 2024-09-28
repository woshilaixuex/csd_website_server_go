package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

/*
 * @Author: deylr1c
 * @Email: linyugang7295@gmail.com
 * @Description: 配置信息
 * @Date: 2024-09-27 23:42
 */
type Config struct {
	Server struct {
		Port int64 `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		User         string `yaml:"user"`
		PassWord     string `yaml:"password"`
		Host         string `yaml:"host"`
		Port         int32  `yaml:"port"`
		DatabaseName string `yaml:"database_name"`
		TableName    string `yaml:"table_name"`
	} `yaml:"database"`
	FeiShuServer struct {
		Open      bool   `yaml:"open"`
		AppId     string `yaml:"app_id"`
		TableId   string `yaml:"table_id"`
		AppToken  string `yaml:"app_token"`
		AppSecret string `yaml:"app_secret"`
	} `yaml:"feishu_server"`
}

// 解析配置
func Load(path string) (*Config, error) {
	var cfg Config
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %w", err)
	}
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("解析 YAML 文件失败: %w", err)
	}
	// 默认数据库名称和表名
	if cfg.Database.DatabaseName == "" {
		cfg.Database.DatabaseName = "csd_website_server"
	}
	if cfg.Database.TableName == "" {
		cfg.Database.TableName = "enroll_table"
	}
	return &cfg, nil
}

// 强制要求配置格式，否则panic
func MustLoad(path string) *Config {
	cfg, err := Load(path)
	if err != nil {
		panic(fmt.Sprintf("无法解析配置文件: %v", err))
	}
	return cfg
}
