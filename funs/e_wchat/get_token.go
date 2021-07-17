package e_wchat

import (
	"context"
	"dragon_micro/utils"
)

func (t *API) GetToken(ctx context.Context, args *GetTokenArgs, reply *Reply) error {
	eWechatAPI := utils.NewEWechatAPI(args.CorpId, args.CorpSecret)
	eWechatCallback := utils.EWechatCallback{}
	eWechatCallback.GetToken(eWechatAPI, func(result map[string]interface{}, err error) {
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
