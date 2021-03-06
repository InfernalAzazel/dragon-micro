package utils

import (
	"encoding/json"
	"errors"
	"github.com/imroc/req"
	"github.com/spf13/cast"
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
func (t *JDAPICallback)sendRequest (api *jdAPIRequest, method string, requestUrl string, payload interface{},
	callback func(map[string]interface{}, error)) {

	var resp *req.Resp
	var err error
	var result map[string]interface{}
	r := req.New()

	header := req.Header{
		"Authorization": "Bearer " + api.apiKey,
		"Content-Type": "application/json;charset=utf-8",
	}

	if  method == "GET" {
		resp, err = r.Get(requestUrl, header, payload)
		if err != nil {
			callback(nil, err)
		}
	} else {
		// POST 请求
		resp, err = r.Post(requestUrl, header, payload)
		if err != nil {
			callback(nil, err)
		}

	}
	err = resp.ToJSON(&result)

	if resp.Response().StatusCode >= 400 {
		if result["code"].(float64) == 8303 && api.retryIfRateLimited {
			// 频率超限，1 毫秒后重试
			time.Sleep(1 * 1000 * 1000)
			t.sendRequest(api, method, requestUrl, payload, callback)
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
func (t *JDAPICallback) GetFormWidgets (api *jdAPIRequest, callback func([]map[string]interface{}, error)) {
	t.sendRequest(api, "POST", api.requestUrl.getWidgets, nil, func(result map[string]interface{}, err error) {
		if err != nil {
			callback(nil, err)
		} else {
			// 解析转 JSON 对象
			var m []map[string]interface{}
			b, err := json.Marshal(result["widgets"])
			if err != nil {
				callback(nil, err)
			}
			err = json.Unmarshal(b, &m)
			if err != nil {
				callback(nil, err)
			}
			callback(m, nil)
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
func (t *JDAPICallback) GetFormData (api *jdAPIRequest, limit int, fields []string, filter map[string]interface{}, dataId string, callback func([]map[string]interface{}, error)) {
	queryData := make(map[string]interface{})
	queryData["limit"] = limit
	queryData["fields"] = fields
	queryData["filter"] = filter
	if dataId != "" {
		queryData["data_id"] = dataId
	}
	t.sendRequest(api, "POST", api.requestUrl.getFormData, req.BodyJSON(queryData), func(result map[string]interface{}, err error) {
		if err != nil {
			callback(nil, err)
		} else {
			// 解析转 JSON 对象
			var m []map[string]interface{}
			b, err := json.Marshal(result["data"])
			if err != nil {
				callback(nil, err)
			}
			err = json.Unmarshal(b, &m)
			if err != nil {
				callback(nil, err)
			}
			callback(m , nil)
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
	t.sendRequest(api, "POST", api.requestUrl.retrieveData, req.BodyJSON(requestData), func(result map[string]interface{}, err error) {
		if err != nil {
			callback(nil, err)
		} else {
			callback(cast.ToStringMap(result["data"]), nil)
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
	t.sendRequest(api, "POST", api.requestUrl.updateData, req.BodyJSON(requestData), func(result map[string]interface{}, err error) {
		if err != nil {
			callback(nil, err)
		} else {
			callback(cast.ToStringMap(result["data"]), nil)
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
	t.sendRequest(api, "POST", api.requestUrl.createData,  req.BodyJSON(requestData), func(result map[string]interface{}, err error) {
		if err != nil {
			callback(nil, err)
		} else {
			callback(cast.ToStringMap(result["data"]), nil)
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
	t.sendRequest(api, "POST", api.requestUrl.deleteData,  req.BodyJSON(requestData), callback)
}

/**
 * 示例
 */

//funs main () {
	// 输入自己的key
//	appId := ""
//	entryId := ""
//	apiKey := ""
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
