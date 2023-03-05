/**
* @Description 注册job
* @Author Mikael
* @Email 13247629622@163.com
* @Date 2021/1/18 12:05
* @Version 1.0
**/
package handler

import (
	"bspick/dqjob/internal/logic"
	"bspick/dqjob/internal/svc"
	"context"
	"github.com/zeromicro/go-zero/core/service"
)

func RegisterJob(serverCtx *svc.ServiceContext, group *service.ServiceGroup) {

	group.Add(logic.NewProducerLogic(context.Background(), serverCtx))
	group.Add(logic.NewConsumerLogic(context.Background(), serverCtx))

}
