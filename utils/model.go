package utils

const (

	JDBasic= "https://www.jiandaoyun.com"
	EWechatBasic = "https://qyapi.weixin.qq.com/cgi-bin"

)

// sql
type SQLCrudCallback struct {
}

// 简道云API回调
type JDAPICallback struct {}

// 简道云API URL请求
type jdAPIRequest struct {
	// 对应表单API请求的url
	requestUrl struct{
		getWidgets string
		getFormData string
		retrieveData string
		createData string
		updateData string
		deleteData string
	}
	// 频率超限后请求是否重试
	retryIfRateLimited bool
	apiKey string
}

// 企业微信API回调
type EWechatCallback struct {}
// 企业微信API URL请求
type eWechatAPIRequest struct {
	// 对应企业微信API请求的url
	requestUrl struct{
		// 获取access_token
		GetToken string
		// 获取成员
		getUserList string
		// 获取成员详细信息
		getUserSimpleList string
		// 获取部门
		getDepartmentList string

	}
	// 频率超限后请求是否重试
	retryIfRateLimited bool
	apiKey string
}

