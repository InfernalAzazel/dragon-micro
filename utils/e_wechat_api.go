package utils

import (
	"github.com/imroc/req"
	"time"
)


/**
 * 构造函数
 * @param corpId - 是 企业ID
 * @param corpSecret - 是 应用的凭证密钥
 */
func NewEWechatAPI(corpId string, corpSecret string) *eWechatAPIRequest {
	request := new(eWechatAPIRequest)
	request.retryIfRateLimited = false
	request.requestUrl.GetToken = EWechatBasic + "/gettoken?corpid="+ corpId + "&corpsecret=" + corpSecret
	request.requestUrl.getUserList = EWechatBasic + "/user/list?"
	request.requestUrl.getUserSimpleList = EWechatBasic + "/user/simplelist?"
	request.requestUrl.getDepartmentList = EWechatBasic + "/department/list?"

	return request

}
/**
 * 发送HTTP请求
 * @param ewAPIRequest - 是 请求实例化
 * @param method - 是 HTTP类型
 * @param requestUrl - 是 请求的url
 * @param payload - 是 请求数据
 * @param callback - 是 回调函数
 */
func (t *EWechatCallback) sendRequest (api *eWechatAPIRequest, method string, requestUrl string, payload interface{},
	callback func(map[string]interface{}, error)){

	var resp *req.Resp
	var err error
	var result map[string]interface{}
	r := req.New()

	if  method == "GET" {
		resp, err = r.Get(requestUrl, payload)
		if err != nil {
			callback(result, err)
		}
	} else {
		resp, err = r.Post(requestUrl, payload)
		if err != nil {
			callback(result, err)
		}

	}
	err = resp.ToJSON(&result)

	if api.retryIfRateLimited {
		// 频率超限，1s后重试
		time.Sleep(1 * 1000 * 1000 * 1000)
		t.sendRequest(api, "GET", requestUrl, payload, func(m map[string]interface{}, err error) {
			callback(result, err)
		})
	}else {
		callback(result, err)
	}

}

/**
 * 获取认证
 * @param ewAPIRequest - 是 请求实例化
 * @param callback - 是 回调函数
 */
func (t *EWechatCallback) GetToken(api *eWechatAPIRequest, callback func(map[string]interface{}, error))  {
	t.sendRequest(api, "GET",
		api.requestUrl.GetToken,
		nil,
		func(result map[string]interface{}, err error) {
			callback(result, err)
	})
}

/**
 * 获取成员详细信息
 * @param ewAPIRequest - 是 请求实例化
 * @param accessToken - 是 调用接口凭证
 * @param departmentId - 是 获取的部门id
 * @param fetchChild - 否 是否递归获取子部门下面的成员：1-递归获取，0-只获取本部门  可以不填
 * @param callback - 是 回调函数
 */
func (t *EWechatCallback) UserList(api *eWechatAPIRequest, accessToken string,departmentId int, fetchChild int, callback func(map[string]interface{}, error))  {
	t.sendRequest(api, "GET",
		api.requestUrl.getUserList,
		req.Param{
			"access_token": accessToken,
			"department_id": departmentId,
			"fetch_child": fetchChild,
		},
		func(result map[string]interface{}, err error) {
			callback(result, err)
		})
}

/**
 * 获取部门成员
 * @param ewAPIRequest - 是 请求实例化
 * @param accessToken - 是 调用接口凭证
 * @param departmentId - 是 获取的部门id
 * @param fetchChild - 否 是否递归获取子部门下面的成员：1-递归获取，0-只获取本部门
 * @param callback - 是 回调函数
 */
func (t *EWechatCallback) UserSimpleList(api *eWechatAPIRequest, accessToken string,departmentId int, fetchChild int, callback func(map[string]interface{}, error))  {
	t.sendRequest(api, "GET",
		api.requestUrl.getUserSimpleList,
		req.Param{
			"access_token": accessToken,
			"department_id": departmentId,
			"fetch_child": fetchChild,
		},
		func(result map[string]interface{}, err error) {
			callback(result, err)
		})
}

/**
 * 获取部门列表
 * @param ewAPIRequest - 是 请求实例化
 * @param accessToken - 是 调用接口凭证
 * @param id - 否 部门id。获取指定部门及其下的子部门（以及及子部门的子部门等等，递归）。 如果不填，默认获取全量组织架构
 * @param callback - 是 回调函数
 */
func (t *EWechatCallback) GetDepartmentList(api *eWechatAPIRequest, accessToken string, id int, callback func(map[string]interface{}, error))  {
	t.sendRequest(api, "GET",
		api.requestUrl.getDepartmentList,
		req.Param{
			"access_token": accessToken,
			"id": id,
		},
		func(result map[string]interface{}, err error) {
			callback(result, err)
		})
}