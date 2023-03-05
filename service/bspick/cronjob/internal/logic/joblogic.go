package logic

import (
	"bspick/cronjob/internal/svc"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/threading"
	"time"
)

type JobLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewJobLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JobLogic {
	return &JobLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}
func (l *JobLogic) Spec() string {
	spec := "@every 1s"
	return spec
}
func (l *JobLogic) Run() {

	threading.GoSafe(func() {
		l.work()
	})
}

func (l *JobLogic) work() {
	time.Sleep(5 * time.Second)
	l.Logger.Infof("start  work")
}
