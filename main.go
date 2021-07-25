package main

import (
	"context"
	"dragon_micro/funs/crud"
	"dragon_micro/funs/e_wchat"
	"dragon_micro/funs/jd"
	"dragon_micro/funs/test"
	"errors"
	"github.com/smallnest/rpcx/protocol"
	"github.com/smallnest/rpcx/server"
	"runtime"
)

func main()  {

	// 开启多核心支持
	runtime.GOMAXPROCS(runtime.NumCPU())
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