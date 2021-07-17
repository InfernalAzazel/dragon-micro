package main

import (
	"context"
	"dragon_micro/funs/jd"
	"dragon_micro/utils"
	"errors"
	"fmt"
	"github.com/smallnest/rpcx/protocol"
	"github.com/smallnest/rpcx/server"
)

func main()  {

	eWechatAPI := utils.NewEWechatAPI("ww9ec26301b4320ef9","IAFxK_Qalqg6DHVEXIpXN9_b42wVWBSjcFn9HV-Y1b0")
	eWechatCallback := utils.EWechatCallback{}
	eWechatCallback.GetToken(eWechatAPI, func(result map[string]interface{}, err error) {
		if err != nil {
			fmt.Println(err)
		}else {
			fmt.Println(result)
		}
	})

	s := server.NewServer()
	s.AuthFunc = auth
	s.RegisterName("JD_API", new(jd.API), "")
	s.Serve("tcp", ":8888")
}

// 认证
func auth(ctx context.Context, req *protocol.Message, token string) error {

	if token == "Bearer jAz2hk5pdOtqMwLz2ZkgP80Q" {
		return nil
	}

	return errors.New("invalid token")
}