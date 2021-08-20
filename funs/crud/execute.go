package crud

import (
	"context"
	"dragon-micro/utils"
)

func (t *API) Execute(ctx context.Context, args *RequestArgs, reply *Reply) error {
	mysql, err := utils.NewEngineSQL(args.DriverName, args.DataSourceName)
	if err != nil {
		reply.State = "fail"
		reply.Err = err.Error()
	}else {
		mySqlCrudCallback := utils.SQLCrudCallback{}
		mySqlCrudCallback.Execute(mysql, args.Sql, args.Args, func(result int64, err error) {
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
