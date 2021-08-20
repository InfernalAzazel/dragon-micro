package jd

import (
	"context"
	"dragon-micro/utils"
	_ "dragon-micro/utils"
)

func (t *API) UpdateData(ctx context.Context, args *UpdateDataArgs, reply *Reply) error {
	jdApi := utils.NewJDAPI(args.AppId, args.EntryId, args.ApiKey)
	jdCallback := utils.JDAPICallback{}
	jdCallback.UpdateData(jdApi, args.DataId, args.Data, args.IsStartTrigger, func(result map[string]interface{}, err error) {
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
