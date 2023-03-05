/**
* @Description 启动文件
* @Author Mikael
* @Email 13247629622@163.com
* @Date 2021/1/18 12:05
* @Version 1.0
**/
package main

import (
	"bspick/cronjob/internal/config"
	"bspick/cronjob/internal/handler"
	"bspick/cronjob/internal/svc"
	"flag"
	"github.com/robfig/cron/v3"
	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/job.yaml", "the config file")

func main() {
	flag.Parse()

	//配置
	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	//注册job
	cn := cron.New()
	handler.RegisterJob(ctx, cn)

	cn.Start()
	select {}
}
