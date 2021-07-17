package e_wchat

import (
	"context"
	"dragon_micro/utils"
)

func (t *API) UserList(ctx context.Context, args *UserListArgs, reply *Reply) error {
	api := utils.NewEWechatAPI(args.CorpId, args.CorpSecret)
	eWechatCallback := utils.EWechatCallback{}
	eWechatCallback.UserList(api, args.AccessToken, args.DepartmentId, args.FetchChild,func(result map[string]interface{}, err error) {
		if err != nil {
			reply.State = "fail"
			reply.Err = err.Error()
		} else {
			reply.State = "success"
			reply.Data = result
		}
	})

	return nil
}

