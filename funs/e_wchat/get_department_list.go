package e_wchat

import (
	"context"
	"dragon-micro/utils"
)

func (t *API) GetDepartmentList(ctx context.Context, args *GetDepartmentListArgs, reply *Reply) error {
	eWechatAPI := utils.NewEWechatAPI(args.CorpId, args.CorpSecret)
	eWechatCallback := utils.EWechatCallback{}
	eWechatCallback.GetDepartmentList(eWechatAPI, args.AccessToken, args.Id, func(result map[string]interface{}, err error) {
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
