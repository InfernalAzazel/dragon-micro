package utils




/**
 * 构造函数
 * @param corpId - 企业ID
 * @param corpSecret - 应用的凭证密钥
 */
func NewEWechat(qwAPIRequest *eWechatAPIRequest, corpId string, corpSecret string)  {
	qwAPIRequest.requestUrl.getDepartmentList = EWechatBasic + "/department/list?"
	qwAPIRequest.requestUrl.getUserSimpleList = EWechatBasic + "/user/simplelist?"
	qwAPIRequest.requestUrl.getUserList = EWechatBasic + "/user/list?"

}