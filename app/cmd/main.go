package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/csd-world/csd_webstie_server_go/app/config"
	"github.com/csd-world/csd_webstie_server_go/internal/handlers"
	"github.com/csd-world/csd_webstie_server_go/internal/models/engine"
	"github.com/csd-world/csd_webstie_server_go/internal/services"
	"github.com/csd-world/csd_webstie_server_go/pkg"
	"github.com/csd-world/csd_webstie_server_go/routes"
	"github.com/gin-gonic/gin"
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
	// 初始化数据库
	engine := engine.InitDB(cfg)
	// 协程开启飞书服务监听（启动失败会报错，不会panic）
	resultCh := pkg.NewMssChan(30)
	if cfg.FeiShuServer.Open {
		var once sync.Once
		defer func() {
			if resultCh.CheckIsOpen() {
				resultCh.Close()
			}
		}()
		once.Do(func() {
			resultCh.Open()
		})
		// 启动飞书监听服务
		go func() {
			services.FeiShuServiceLisen(cfg, resultCh.C) // 传入只读管道
		}()
	}
	// 初始化服务
	InitSever := func() {
		handler := handlers.NewHandler(engine, resultCh)
		r := gin.Default()
		routes.RegisterRoutes(r, handler)
		r.Run(":" + strconv.FormatInt(cfg.Server.Port, 10))
	}
	InitSever()
}
