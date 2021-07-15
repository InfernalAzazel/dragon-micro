package main

import (
	"dragon_micro/funs/jd"
	"github.com/smallnest/rpcx/server"
)

func main()  {
	s := server.NewServer()
	s.RegisterName("GetFormWidgets", new(jd.GetFormWidgets), "")
	s.Serve("tcp", ":8972")
}