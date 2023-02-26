package logic

import (
	"context"

	"demo1/great/internal/svc"
	"demo1/great/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GreatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGreatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GreatLogic {
	return &GreatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GreatLogic) Great(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return &types.Response{
		Message: req.Name,
	}, nil
}
