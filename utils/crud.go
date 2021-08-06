package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"github.com/xormplus/xorm"
)

func NewEngineSQL(driverName string, dataSourceName string)(*xorm.Engine, error) {
	// "root:3LZSsMW9EJOFtbcm2gJFySKE@tcp(127.0.0.1:3306)/test?charset=utf8"
	//sql_1_1 := "select * from myform"
	//sql_2 := "update myform set name = ? where id = ?"

	var engine *xorm.Engine
	var err error
	engine, err = xorm.NewEngine(driverName, dataSourceName)
	if err != nil {
		logrus.Error(err.Error())
	}
	return engine, err
}

func (t *SQLCrudCallback) Query(db *xorm.Engine, query interface{}, args []interface{}, callback func([]map[string]interface{}, error))  {
	var results []map[string]interface{}

	defer func(db *xorm.Engine) {
		err := db.Close()
		if err != nil {
			callback(nil, err)
		}
	}(db)
	if args == nil{
		results = db.SQL(query).Query().Result
	}else {
		results = db.SQL(query, args...).Query().Result
	}
	callback(results, nil)
}

func (t *SQLCrudCallback) Execute(db *xorm.Engine, query string, args []interface{},callback func(rowsAffected int64, err error)) {
	var rowsAffected int64
	var results sql.Result
	var err error
	defer func(db *xorm.Engine) {
		err = db.Close()
		if err != nil {
			callback(-1, err)
		}
	}(db)

	if args == nil{
		results, err = db.SQL(query).Execute()
	}else {
		results, err = db.SQL(query, args...).Execute()
	}

	if err != nil {
		callback(-1, err)
		return
	}
	rowsAffected,err= results.RowsAffected()
	if err != nil {
		callback(rowsAffected, err)
	}else {
		callback(rowsAffected, nil)
	}

}

