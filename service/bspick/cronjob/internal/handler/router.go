/**
* @Description 注册job
* @Author Mikael
* @Email 13247629622@163.com
* @Date 2021/1/18 12:05
* @Version 1.0
**/
package handler

import (
	"bspick/cronjob/internal/logic"
	"bspick/cronjob/internal/svc"
	"context"
	"github.com/robfig/cron/v3"
)

func RegisterJob(serverCtx *svc.ServiceContext, cn *cron.Cron) {

	jobLogic := logic.NewJobLogic(context.Background(), serverCtx)
	cn.AddJob(jobLogic.Spec(), jobLogic)
}
