package logic

import (
	"context"
	"errors"

	"{{template}}/internal/svc"
	"{{template}}/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HealthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHealthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HealthLogic {
	return &HealthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HealthLogic) Httptemplate(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	return &types.Response{
		Message: "abc",
	}, errors.New("wait??")
}
