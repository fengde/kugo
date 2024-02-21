package consts

import "github.com/zeromicro/go-zero/core/service"

const (
	// 开发环境
	ENV_DEV = service.DevMode
	// 测试环境 or qa
	ENV_TEST = service.TestMode
	// 回归测试环境
	ENV_RT = service.RtMode
	// 预发布环境 or staging
	ENV_PRE = service.PreMode
	// 线上环境 or online
	ENV_PRO = service.ProMode
)
