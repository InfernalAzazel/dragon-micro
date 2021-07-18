package crud


// 数据库操作
type API int

// 统一响应
type Reply struct {
	// 状态
	State string `json:"state"`
	// 数据
	Data interface{} `json:"data"`
	// 异常
	Err string `json:"err"`
}

// 表单字段查询接口
type RequestArgs struct {
	// 数据库驱动名称
	DriverName string `json:"driver_name"`
	// url
	DataSourceName string `json:"data_source_name"`
	// sql语句
	Sql string `json:"sql"`
	// 对应的数组
	Args []interface{} `json:"args"`
}