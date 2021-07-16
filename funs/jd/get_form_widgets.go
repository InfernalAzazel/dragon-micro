package jd

import (
	"context"
	"dragon_micro/utils"
	_ "dragon_micro/utils"
)





func (t *API) GetFormWidgets(ctx context.Context, args *GetFormWidgetsArgs, reply *Reply) error {
	jdApi := utils.NewJDAPI(args.AppId, args.EntryId, args.ApiKey)
	jdCallback := utils.JDAPICallback{}
	jdCallback.GetFormWidgets(jdApi, func(widgets []interface{}, err error) {
		if err != nil {
			reply.State = "fail"
			reply.Err = err.Error()
		} else {
			reply.State = "success"
			reply.Data = widgets
		}
	})

	return nil
}
