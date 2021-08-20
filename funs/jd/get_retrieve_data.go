
package jd

import (
	"context"
	"dragon-micro/utils"
	_ "dragon-micro/utils"
)

func (t *API) GetRetrieveData(ctx context.Context, args *GetRetrieveData, reply *Reply) error {
	jdApi := utils.NewJDAPI(args.AppId, args.EntryId, args.ApiKey)
	jdCallback := utils.JDAPICallback{}
	jdCallback.GetRetrieveData(jdApi, args.DataId, func(data map[string]interface{}, err error) {
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