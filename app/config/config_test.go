package config_test

import (
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/csd-world/csd_webstie_server_go/app/config"
)

/*
 * @Author: deylr1c
 * @Email: linyugang7295@gmail.com
 * @Description: 配置文件加载测试
 * @Date: 2024-09-28 00:25
 */
var configFile = flag.String("f", "../etc/server.yaml", "the config file")

func LoadConfig() (*config.Config, error) {
	// 加载并解析配置文件
	flag.Parse()
	// 检查配置文件是否存在，不存在则 panic
	if _, err := os.Stat(*configFile); os.IsNotExist(err) {
		panic(fmt.Sprintf("config file is not exit: %s", *configFile))
	}
	return config.Load(*configFile)
}
func TestLoadConfig(t *testing.T) {
	cfg, err := LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("无法解析配置文件: %v", err))
	}
	t.Logf("%v", cfg)
}
