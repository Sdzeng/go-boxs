package main

import (
	"flag"
	"fmt"
	"os"

	"demo1/great/internal/config"
	"demo1/great/internal/handler"
	"demo1/great/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "/great/etc/great-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	configFile = absPath(configFile)
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

func absPath(conf *string) *string {
	path, _ := os.Getwd()
	result := path + *conf
	return &result
}
