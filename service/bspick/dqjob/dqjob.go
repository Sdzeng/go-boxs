/**
* @Description 启动文件
* @Author Mikael
* @Email 13247629622@163.com
* @Date 2021/1/18 12:05
* @Version 1.0
**/
package main

import (
	"bspick/dqjob/internal/config"
	"bspick/dqjob/internal/handler"
	"bspick/dqjob/internal/svc"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var configFile = flag.String("f", "bspick/cronjob/etc/job.yaml", "the config file")

func main() {
	flag.Parse()

	//配置
	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	//注册job
	group := service.NewServiceGroup()
	handler.RegisterJob(ctx, group)

	//捕捉信号
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		for {
			s := <-ch
			logx.Infof("get a signal %s", s.String())
			switch s {
			case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
				fmt.Printf("stop group")
				group.Stop()
				logx.Info("job exit")
				time.Sleep(time.Second)
				return
			case syscall.SIGHUP:
			default:
				return
			}
		}
	}()

	group.Start()
}
