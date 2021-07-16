package jd

import (
	"context"
	"dragon_micro/utils"
	_ "dragon_micro/utils"
)

func (t *API) DeleteData(ctx context.Context, args *DeleteDataArgs, reply *Reply) error {
	jdApi := utils.NewJDAPI(args.AppId, args.EntryId, args.ApiKey)
	jdCallback := utils.JDAPICallback{}
	jdCallback.DeleteData(jdApi, args.DataId, args.IsStartTrigger, func(result map[string]interface{}, err error) {
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
