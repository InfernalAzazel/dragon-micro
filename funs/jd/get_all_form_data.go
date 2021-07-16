package jd

import (
	"context"
	"dragon_micro/utils"
	_ "dragon_micro/utils"
)

func (t *API) GetAllFormData(ctx context.Context, args *GetAllFormDataArgs, reply *Reply) error {
	jdApi := utils.NewJDAPI(args.AppId, args.EntryId, args.ApiKey)
	jdCallback := utils.JDAPICallback{}
	jdCallback.GetAllFormData(jdApi, args.Fields, args.Filter, func(data []interface{}, err error) {
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
