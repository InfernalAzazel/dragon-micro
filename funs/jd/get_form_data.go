package jd

import (
	"context"
	"dragon_micro/utils"
	_ "dragon_micro/utils"
)

func (t *API) GetFormData(ctx context.Context, args *GetFormDataArgs, reply *Reply) error {
	jdApi := utils.NewJDAPI(args.AppId, args.EntryId, args.ApiKey)
	jdCallback := utils.JDAPICallback{}
	jdCallback.GetFormData(jdApi, args.Limit, args.Fields, args.Filter, args.DataId, func(data []interface{}, err error) {
		if err != nil {
			reply.State = "fail"
			reply.Err = err.Error()
		} else {
			reply.State = "success"
			reply.Data = data
		}
	})

	return nil
}
