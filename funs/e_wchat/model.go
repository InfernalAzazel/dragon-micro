package e_wchat

// 企业微信API
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
type GetTokenArgs struct {
	// 企业ID
	CorpId string `json:"corp_id"`
	// 应用的凭证密钥
	CorpSecret string `json:"corp_secret"`
}
