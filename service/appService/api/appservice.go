package main

import (
	"flag"
	"fmt"

	"github.com/Zecrey-Labs/zecrey-collisions/service/appService/api/internal/config"
	"github.com/Zecrey-Labs/zecrey-collisions/service/appService/api/internal/handler"
	cryptohandler "github.com/Zecrey-Labs/zecrey-collisions/service/appService/api/internal/logic/crypto/cryptoHandler"
	"github.com/Zecrey-Labs/zecrey-collisions/service/appService/api/internal/svc"
	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "/Users/gavin/Desktop/enc-collision-service/service/appService/api/etc/appservice-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	// init enc db
	err := cryptohandler.InitCrypto(ctx)
	if err != nil {
		logx.Error("[main] %s", err.Error())
	}

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
