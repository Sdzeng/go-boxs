package logic

import (
	"context"

	"bspick/rpc/internal/svc"
	"bspick/rpc/types/bspick"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDataLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDataLogic {
	return &GetDataLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDataLogic) GetData(in *bspick.GetDataReq) (*bspick.GetDataResp, error) {
	// todo: add your logic here and delete this line

	return &bspick.GetDataResp{}, nil
}
