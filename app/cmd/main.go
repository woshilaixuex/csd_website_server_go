package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/csd-world/csd_webstie_server_go/app/config"
	"github.com/csd-world/csd_webstie_server_go/internal/models"
	"github.com/zeromicro/go-zero/core/logx"
)

/*
 * @Author: deylr1c
 * @Email: linyugang7295@gmail.com
 * @Description: 主函数程序入口
 * @Date: 2024-09-27 23:14
 */
var configFile = flag.String("f", "../etc/server.yaml", "the config file")

func main() {
	// 加载并解析配置文件
	flag.Parse()
	if _, err := os.Stat(*configFile); os.IsNotExist(err) {
		panic(fmt.Sprintf("config file is not exit: %s", *configFile))
	}
	cfg := config.MustLoad(*configFile)
	logx.Info("config is init %v", *cfg)
	// 初始化数据库
	models.InitDB(cfg)
}
