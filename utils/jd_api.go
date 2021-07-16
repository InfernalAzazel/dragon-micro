package utils

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

/**
 * 构造函数
 * @param appId - 应用id
 * @param entryId - 表单id
 * @param apiKey - 接口密钥
 */
func NewJDAPI (appId string, entryId string, apiKey string) *jdAPIRequest {
	request := new(jdAPIRequest)
	// 对应请求的url
	request.requestUrl.getWidgets = JDBasic + "/api/v1/app/" + appId + "/entry/" + entryId + "/widgets"
	request.requestUrl.getFormData = JDBasic + "/api/v2/app/" + appId + "/entry/" + entryId + "/data"
	request.requestUrl.retrieveData = JDBasic + "/api/v2/app/" + appId + "/entry/" + entryId + "/data_retrieve"
	request.requestUrl.createData = JDBasic + "/api/v3/app/" + appId + "/entry/" + entryId + "/data_create"
	request.requestUrl.updateData = JDBasic + "/api/v3/app/" + appId + "/entry/" + entryId + "/data_update"
	request.requestUrl.deleteData = JDBasic + "/api/v1/app/" + appId + "/entry/" + entryId + "/data_delete"
	request.retryIfRateLimited = true
	request.apiKey = apiKey
	return request
}

/**
 * 发送HTTP请求
 * @param method - HTTP动词
 * @param header - HTTP Header信息
 * @param requestUrl - 请求的url
 * @param data - 请求数据
 * @param callback - 回调函数
 */
func sendRequest (api *jdAPIRequest, method string, requestUrl string, data map[string]interface{},
	callback func(map[string]interface{}, error)) {
	method = strings.ToUpper(method)
	var resp *http.Response

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	if method == "GET" {
		// GET请求
		u, _ := url.Parse(requestUrl)
		q := u.Query()
		for k, v := range data {
			value, _ := json.Marshal(v)
			q.Set(k, string(value[:]))
		}
		u.RawQuery = q.Encode()
		client := &http.Client{ Transport: tr }
		req, _ := http.NewRequest("GET", requestUrl, nil)
		req.Header.Set("Authorization", "Bearer " + api.apiKey)
		resp, _ = client.Do(req)
	} else {
		// POST请求
		serialData, _ := json.Marshal(data)
		body := ioutil.NopCloser(bytes.NewBuffer(serialData))
		client := &http.Client{ Transport: tr }
		req, _ := http.NewRequest("POST", requestUrl, body)
		req.Header.Set("Content-Type", "application/json;charset=utf-8")
		req.Header.Set("Authorization", "Bearer " + api.apiKey)
		resp, _ = client.Do(req)
	}
	resData, _ := ioutil.ReadAll(resp.Body)
	var result map[string]interface{}
	json.Unmarshal(resData, &result)
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		if result["code"].(float64) == 8303 && api.retryIfRateLimited {
			// 频率超限，1s后重试
			time.Sleep(1 * 1000 * 1000 * 1000)
			sendRequest(api, method, requestUrl, data, callback)
		} else {
			code, _ := json.Marshal(result["code"])
			msg, _ := json.Marshal(result["msg"])
			callback(nil, errors.New("请求错误 Error Code: " + string(code[:]) + " Error Msg: " + string(msg[:])))
		}
	} else {
		callback(result, nil)
	}

}



/**
 * 获取表单字段
 * @param callback - 回调函数
 */
func (JDAPICallback) GetFormWidgets (api *jdAPIRequest, callback func([]interface{}, error)) {
	sendRequest(api, "POST", api.requestUrl.getWidgets, map[string]interface{}{}, func(result map[string]interface{}, err error) {
		if err != nil {
			callback(nil, err)
		} else {
			callback(result["widgets"].([]interface{}), nil)
		}
	})
}

/**
 * 根据条件获取表单数据
 * @param limit - 查询的数据条数
 * @param fields - 查询的字段列表
 * @param filter - 过滤配置
 * @param dataId - 上一次查询数据结果的最后一条数据的id
 * @param callback - 回调函数
 */
func (t *JDAPICallback) GetFormData (api *jdAPIRequest, limit int, fields []string, filter map[string]interface{}, dataId string, callback func([]interface{}, error)) {
	queryData := make(map[string]interface{})
	queryData["limit"] = limit
	queryData["fields"] = fields
	queryData["filter"] = filter
	if dataId != "" {
		queryData["data_id"] = dataId
	}
	sendRequest(api, "POST", api.requestUrl.getFormData, queryData, func(result map[string]interface{}, err error) {
		if err != nil {
			callback(nil, err)
		} else {
			callback(result["data"].([]interface{}), nil)
		}
	})
}

/**
 * 获取满足条件的所有表单数据
 * @param limit - 查询的数据条数
 * @param fields - 查询的字段列表
 * @param filter - 过滤配置
 * @param dataId - 上一次查询数据结果的最后一条数据的id
 * @param callback - 回调函数
 */
func (t *JDAPICallback) GetAllFormData (api *jdAPIRequest, fields []string, filter map[string]interface{}, callback func([]interface{}, error)) {
	// 递归获取所有的数据
	t.getNextPageData(api, []interface{}{}, 100, fields, filter, "", callback)
}

/**
 * 获取下一页的数据，主要用来取所有的表单数据
 * @param formData - 当前已经获取的数据
 * @param limit - 查询的数据条数
 * @param fields - 查询的字段列表
 * @param filter - 过滤配置
 * @param dataId - 上一次查询数据结果的最后一条数据的id
 * @param callback - 回调函数
 */
func (t *JDAPICallback) getNextPageData (api *jdAPIRequest, formData []interface{}, limit int, fields []string, filter map[string]interface{}, dataId string,
	callback func([]interface{}, error))  {
	// 获取单页数据
	t.GetFormData(api, limit, fields, filter, dataId, func(data []interface{}, err error) {
		if err != nil {
			callback(nil, err)
		} else {
			if data != nil && len(data) != 0 {
				// 返回的数据非空
				formData = append(formData, data...)
				// 取出最后一条数据
				lastData := data[len(data)-1].(map[string]interface{})
				// 递归取下一页的数据
				t.getNextPageData(api, formData, limit, fields, filter, lastData["_id"].(string), callback)
			} else {
				// 没有更多的数据
				callback(formData, nil)
			}
		}
	})
}

/**
 * 查询单条数据
 * @param dataId - 数据id
 * @param callback - 回调函数
 */
func (t *JDAPICallback) GetRetrieveData (api *jdAPIRequest, dataId string, callback func(map[string]interface{}, error)) {
	requestData := map[string]interface{}{
		"data_id": dataId,
	}
	sendRequest(api, "POST", api.requestUrl.retrieveData, requestData, func(result map[string]interface{}, err error) {
		if err != nil {
			callback(nil, err)
		} else {
			callback(result["data"].(map[string]interface{}), nil)
		}
	})
}

/**
 * 更新单条数据
 * @param dataId - 数据id
 * @param data - 更新的内容
 * @param isStartTrigger - Bool		是否触发智能助手	false
 * @param callback - 回调函数
 */
func (t *JDAPICallback) UpdateData (api *jdAPIRequest, dataId string, data map[string]interface{}, isStartTrigger bool, callback func(map[string]interface{}, error)) {
	requestData := map[string]interface{}{
		"data_id": dataId,
		"data": data,
		"is_start_trigger": isStartTrigger,
	}
	sendRequest(api, "POST", api.requestUrl.updateData, requestData, func(result map[string]interface{}, err error) {
		if err != nil {
			callback(nil, err)
		} else {
			callback(result["data"].(map[string]interface{}), nil)
		}
	})
}

/**
 * 创建单条数据
 * @param data - 数据内容
 * @param isStartWorkflow - Bool	是否发起流程（仅流程表单有效）	false
 * @param isStartTrigger - Bool		是否触发智能助手	false
 * @param callback - 回调函数
 */
func (t *JDAPICallback) CreateData (api *jdAPIRequest, data map[string]interface{}, isStartWorkflow bool,isStartTrigger bool, callback func(map[string]interface{}, error)) {
	requestData := map[string]interface{}{
		"data": data,
		"is_start_workflow":isStartWorkflow,
		"is_start_trigger": isStartTrigger,
	}
	sendRequest(api, "POST", api.requestUrl.createData, requestData, func(result map[string]interface{}, err error) {
		if err != nil {
			callback(nil, err)
		} else {
			callback(result["data"].(map[string]interface{}), nil)
		}
	})
}

/**
 * 删除单条数据
 * @param dataId - 数据id
 * @param isStartTrigger - Bool		是否触发智能助手	false
 * @param callback - 回调函数
 */
func (t *JDAPICallback) DeleteData (api *jdAPIRequest, dataId string, isStartTrigger bool, callback func(map[string]interface{}, error)) {
	requestData := map[string]interface{}{
		"data_id": dataId,
		"is_start_trigger": isStartTrigger,
	}
	sendRequest(api, "POST", api.requestUrl.deleteData, requestData, callback)
}

/**
 * 示例
 */

//funs main () {
//	appId := "5b1747e93b708d0a80667400"
//	entryId := "5b1749ae3b708d0a80667408"
//	apiKey := "CTRP5jibfk7qnnsGLCCcmgnBG6axdHiX"
//	api := NewAPIRequest(appId, entryId, apiKey)
//
//	// 获取表单字段
//	getFormWidgets(api, funs(widgets []interface{}, err error) {
//		if err != nil {
//			fmt.Println(err)
//		} else {
//			fmt.Println("表单字段：")
//			fmt.Println(widgets)
//		}
//	})
//
//	// 按条件查询数据
//	filter := map[string]interface{}{
//		"rel": "and",
//		"cond": []interface{}{
//			map[string]interface{}{
//				"field": "_widget_1528252846720",
//				"type": "text",
//				"method": "empty",
//			},
//		},
//	}
//	getFormData(api, 10, []string{ "_widget_1528252846720", "_widget_1528252846801" }, filter, "", funs(data []interface{}, err error) {
//		if err != nil {
//			fmt.Println(err)
//		} else {
//			fmt.Println("按条件查询表单数据：")
//			fmt.Println(data)
//		}
//	})
//
//	// 获取全部数据
//	getAllFormData(api, nil, nil, funs(data []interface{}, err error) {
//		if err != nil {
//			fmt.Println(err)
//		} else {
//			fmt.Println("表单全部数据：")
//			fmt.Println(data)
//		}
//	})
//
//	// 创建单条数据
//	data := map[string]interface{}{
//		// 单行文本
//		"_widget_1528252846720": map[string]interface{}{
//			"value": "123",
//		},
//		// 子表单
//		"_widget_1528252846801": map[string]interface{}{
//			"value": []interface{}{
//				map[string]interface{}{
//					"_widget_1528252846952": map[string]interface{}{
//						"value": "123",
//					},
//				},
//			},
//		},
//		// 数字
//		"_widget_1528252847027": map[string]interface{}{
//			"value": 123,
//		},
//		// 地址
//		"_widget_1528252846785": map[string]interface{}{
//			"value": map[string]interface{}{
//				"province": "江苏省",
//				"city": "无锡市",
//				"district": "南长区",
//				"detail": "清名桥街道",
//			},
//		},
//		// 多行文本
//		"_widget_1528252846748": map[string]interface{}{
//			"value": "123123",
//		},
//	}
//	var newData map[string]interface{}
//	createData(api, data, funs(data map[string]interface{}, err error) {
//		if err != nil {
//			fmt.Println(err)
//		} else {
//			newData = data
//			fmt.Println("创建单条数据：")
//			fmt.Println(data)
//		}
//	})
//
//	// 更新单条数据
//	update := map[string]interface{}{
//		// 单行文本
//		"_widget_1528252846720": map[string]interface{}{
//			"value": "12345",
//		},
//		// 子表单
//		"_widget_1528252846801": map[string]interface{}{
//			"value": []interface{}{
//				map[string]interface{}{
//					"_widget_1528252846952": map[string]interface{}{
//						"value": "12345",
//					},
//				},
//			},
//		},
//		// 数字
//		"_widget_1528252847027": map[string]interface{}{
//			"value": 12345,
//		},
//	}
//	updateData(api, newData["_id"].(string), update, funs(result map[string]interface{}, err error) {
//		if err != nil {
//			fmt.Println(err)
//		} else {
//			fmt.Println("更新单条数据：")
//			fmt.Println(result)
//		}
//	})
//
//	// 查询单条数据
//	retrieveData(api, newData["_id"].(string), funs(data map[string]interface{}, err error) {
//		if err != nil {
//			fmt.Println(err)
//		} else {
//			fmt.Println("查询单条数据：")
//			fmt.Println(data)
//		}
//	})
//
//	// 删除单条数据
//	deleteData(api, newData["_id"].(string), funs(result map[string]interface{}, err error) {
//		if err != nil {
//			fmt.Println(err)
//		} else {
//			fmt.Println("删除单条数据：")
//			fmt.Println(result)
//		}
//	})
//}