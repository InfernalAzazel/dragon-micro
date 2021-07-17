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

// 获取成员详细信息
type UserListArgs struct {
	// 企业ID
	CorpId string `json:"corp_id"`
	// 应用的凭证密钥
	CorpSecret string `json:"corp_secret"`
	// 接口调用凭证
	AccessToken string `json:"access_token"`
	// 部门ID
	DepartmentId int `json:"department_id"`
	// 是否递归获取子部门下面的成员：1-递归获取，0-只获取本部门
	FetchChild int `json:"fetch_child"`
}


// 获取成员
type UserSimpleListArgs struct {
	// 企业ID
	CorpId string `json:"corp_id"`
	// 应用的凭证密钥
	CorpSecret string `json:"corp_secret"`
	// 接口调用凭证
	AccessToken string `json:"access_token"`
	// 部门ID
	DepartmentId int `json:"department_id"`
	// 是否递归获取子部门下面的成员：1-递归获取，0-只获取本部门
	FetchChild int `json:"fetch_child"`

}

// 获取部门列表
type GetDepartmentListArgs struct {
	// 企业ID
	CorpId string `json:"corp_id"`
	// 应用的凭证密钥
	CorpSecret string `json:"corp_secret"`
	// 接口调用凭证
	AccessToken string `json:"access_token"`
	// 部门ID
	Id  int `json:"id"`
}
