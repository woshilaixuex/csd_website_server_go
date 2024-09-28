package engine_test

import (
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/csd-world/csd_webstie_server_go/app/config"
	models "github.com/csd-world/csd_webstie_server_go/internal/models/engine"
)

/*
 * @Author: deylr1c
 * @Email: linyugang7295@gmail.com
 * @Description: 引擎测试
 * @Date: 2024-09-28 02:04
 */
var configFile = flag.String("f", "../../../app/etc/server.yaml", "the config file")

func TestInitDB(t *testing.T) {
	flag.Parse()
	if _, err := os.Stat(*configFile); os.IsNotExist(err) {
		panic(fmt.Sprintf("config file is not exit: %s", *configFile))
	}
	cfg := config.MustLoad(*configFile)
	models.InitDB(cfg)
}
