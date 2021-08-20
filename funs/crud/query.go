package crud

import (
	"context"
	"dragon-micro/utils"
)

func (t *API) Query(ctx context.Context, args *RequestArgs, reply *Reply) error {
	engineSQL, err := utils.NewEngineSQL(args.DriverName, args.DataSourceName)
	if err != nil {
		reply.State = "fail"
		reply.Err = err.Error()
	}else {
		mySqlCrudCallback := utils.SQLCrudCallback{}
		mySqlCrudCallback.Query(engineSQL, args.Sql, args.Args, func(result []map[string]interface{}, err error) {
			if err != nil {
				reply.State = "fail"
				reply.Err = err.Error()
			} else {
				reply.State = "success"
				reply.Data = result
			}
		})
	}
	return nil
}
