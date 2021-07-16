package jd

import (
	"context"
	"dragon_micro/utils"
	_ "dragon_micro/utils"
)

func (t *API) CreateData(ctx context.Context, args *CreateDataArgs, reply *Reply) error {
	jdApi := utils.NewJDAPI(args.AppId, args.EntryId, args.ApiKey)
	jdCallback := utils.JDAPICallback{}
	jdCallback.CreateData(jdApi, args.Data, args.IsStartWorkflow, args.IsStartTrigger, func(result map[string]interface{}, err error) {
		print(args.Data)
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
