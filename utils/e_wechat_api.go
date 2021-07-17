package utils

import (
	"github.com/imroc/req"

)


/**
 * 构造函数
 * @param corpId - 企业ID
 * @param corpSecret - 应用的凭证密钥
 */
func NewEWechatAPI(corpId string, corpSecret string) *eWechatAPIRequest {
	request := new(eWechatAPIRequest)
	request.requestUrl.GetToken = EWechatBasic + "/gettoken?corpid="+ corpId + "&corpsecret=" + corpSecret
	request.requestUrl.getUserList = EWechatBasic + "/user/list?"
	request.requestUrl.getUserSimpleList = EWechatBasic + "/user/simplelist?"
	request.requestUrl.getDepartmentList = EWechatBasic + "/department/list?"

	return request

}

func (t *EWechatCallback) sendRequest (ewAPIRequest *eWechatAPIRequest, method string, requestUrl string, payload interface{},
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
		resp, err = r.Post(requestUrl, req.BodyJSON(payload))
		if err != nil {
			callback(result, err)
		}

	}
	err = resp.ToJSON(&result)

	if ewAPIRequest.retryIfRateLimited {
		callback(result, err)
	}else {
		callback(result, err)
	}
}

func (t *EWechatCallback) GetToken(api *eWechatAPIRequest, callback func(map[string]interface{}, error))  {
	t.sendRequest(api, "GET",
		api.requestUrl.GetToken,
		nil,
		func(result map[string]interface{}, err error) {
			callback(result, err)
	})
}
