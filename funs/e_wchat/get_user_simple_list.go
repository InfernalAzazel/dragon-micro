package e_wchat


import (
	"context"
	"dragon-micro/utils"
)

func (t *API) GetUserSimpleList(ctx context.Context, args *UserSimpleListArgs, reply *Reply) error {
	api := utils.NewEWechatAPI(args.CorpId, args.CorpSecret)
	eWechatCallback := utils.EWechatCallback{}
	eWechatCallback.GetUserSimpleList(api, args.AccessToken, args.DepartmentId, args.FetchChild,func(result map[string]interface{}, err error) {
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
