package jd

// 简道云API
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
type GetFormWidgetsArgs struct {
	// 应用ID
	AppId string `json:"app_id"`
	// 表单ID
	EntryId string `json:"entry_id"`
	// API密钥
	ApiKey string `json:"api_key"`
}

// 查询多条数据接口
type GetFormDataArgs struct {
	// 应用ID
	AppId string `json:"app_id"`
	// 表单ID
	EntryId string `json:"entry_id"`
	// API密钥
	ApiKey string `json:"api_key"`
	// 上一次查询数据结果的最后一条数据的ID，没有则留空
	DataId string `json:"data_id"`
	// 需要查询的数据字段
	Fields []string `json:"fields"`
	// 数据筛选器
	Filter map[string]interface{} `json:"filter"`
	// 查询的数据条数，1~100，默认10
	Limit int `json:"limit"`
}

// 查询单条数据接口
type GetRetrieveData struct {
	// 应用ID
	AppId string `json:"app_id"`
	// 表单ID
	EntryId string `json:"entry_id"`
	// API密钥
	ApiKey string `json:"api_key"`
	// 数据ID
	DataId string `json:"data_id"`
}

// 查询全部数据接口
type GetAllFormDataArgs struct {
	// 应用ID
	AppId string `json:"app_id"`
	// 表单ID
	EntryId string `json:"entry_id"`
	// API密钥
	ApiKey string `json:"api_key"`
	// 需要查询的数据字段
	Fields []string `json:"fields"`
	// 数据筛选器
	Filter map[string]interface{} `json:"filter"`
}

// 创建表单接口
type CreateDataArgs struct {
	// 应用ID
	AppId string `json:"app_id"`
	// 表单ID
	EntryId string `json:"entry_id"`
	// API密钥
	ApiKey string `json:"api_key"`
	// 数据
	Data map[string]interface{} `json:"data"`
	// Bool	是否发起流程（仅流程表单有效）
	IsStartWorkflow bool `json:"is_start_workflow"`
	// 是否触发智能助手
	IsStartTrigger bool `json:"is_start_trigger"`
}

// 删除数据接口
type DeleteDataArgs struct {
	// 应用ID
	AppId string `json:"app_id"`
	// 表单ID
	EntryId string `json:"entry_id"`
	// API密钥
	ApiKey string `json:"api_key"`
	// 数据ID
	DataId string `json:"data_id"`
	// 是否触发智能助手
	IsStartTrigger bool `json:"is_start_trigger"`
}

// 更新数据接口
type UpdateDataArgs struct {
	// 应用ID
	AppId string `json:"app_id"`
	// 表单ID
	EntryId string `json:"entry_id"`
	// API密钥
	ApiKey string `json:"api_key"`
	// 数据ID
	DataId string `json:"data_id"`
	// 数据
	Data map[string]interface{} `json:"data"`
	// 是否触发智能助手
	IsStartTrigger bool `json:"is_start_trigger"`
}