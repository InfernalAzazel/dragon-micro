package main

import (
	"context"
	"dragon-micro/funs/crud"
	"dragon-micro/funs/e_wchat"
	"dragon-micro/funs/jd"
	"dragon-micro/funs/test"
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/smallnest/rpcx/protocol"
	"github.com/smallnest/rpcx/server"
	"os"
	"runtime"
)

func main()  {
	// 开启多核心支持
	runtime.GOMAXPROCS(runtime.NumCPU())

	customFormatter := new(logrus.TextFormatter)
	customFormatter.FullTimestamp = true                    // 显示完整时间
	customFormatter.TimestampFormat = "2006-01-02 15:04:05" // 时间格式
	customFormatter.DisableTimestamp = false                // 禁止显示时间
	customFormatter.DisableColors = false                   // 禁止颜色显示

	logrus.SetFormatter(customFormatter)
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)

	s := server.NewServer()
	s.AuthFunc = auth
	s.RegisterName("Arith", new(test.Arith), "")
	s.RegisterName("CRUD", new(crud.API), "")
	s.RegisterName("JD_API", new(jd.API), "")
	s.RegisterName("EWechat_API", new(e_wchat.API), "")
	s.Serve("tcp", ":8888")
}

// 认证
func auth(ctx context.Context, req *protocol.Message, token string) error {

	if token == "Bearer jAz2hk5pdOtqMwLz2ZkgP80Q" {
		return nil
	}

	return errors.New("invalid token")
}