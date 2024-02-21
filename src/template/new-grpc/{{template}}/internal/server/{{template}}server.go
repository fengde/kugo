// Code generated by goctl. DO NOT EDIT.
// Source: {{template}}.proto

package server

import (
	"context"

	"{{template}}/{{template}}"
	"{{template}}/internal/logic"
	"{{template}}/internal/svc"
)

type {{Template}}Server struct {
	svcCtx *svc.ServiceContext
	{{template}}.Unimplemented{{Template}}Server
}

func New{{Template}}Server(svcCtx *svc.ServiceContext) *{{Template}}Server {
	return &{{Template}}Server{
		svcCtx: svcCtx,
	}
}

func (s *{{Template}}Server) Ping(ctx context.Context, in *{{template}}.Request) (*{{template}}.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}
