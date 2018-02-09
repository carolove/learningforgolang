package server

import (
	"github.com/cuigh/auxo/config"
	"github.com/cuigh/auxo/net/rpc"
)

// 读取配置信息读取app name以及
// 第一 读取配置文件 惰性加载
// 第二 读取程序名注册
// 第三 生成Daemon类型的server
func Create() (*rpc.Server, error) {
	serverName := config.GetString("app")
}
