package json

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// 解析数据为JSON对象
func ToJSONObject(jsonData []byte) (*JSONObject, error) {
	m := make(map[string]interface{})

	err := json.Unmarshal(jsonData, &m)
	if err != nil {
		return nil, err
	}

	return parseJsonObject(m)
}

// 解析数据为JSON数组
func ToJSONArray(jsonData []byte) (*JSONArray, error) {
	var aa []interface{}

	err := json.Unmarshal(jsonData, &aa)
	if err != nil {
		return nil, err
	}

	return parseJsonArray(aa)
}

// 解析为JSON对象
func parseJsonObject(m map[string]interface{}) (*JSONObject, error) {
	jo := NewJSONObject()

	for k, v := range m {

		if v == nil {
			continue
		}

		vKind := reflect.TypeOf(v).Kind().String()

		if vKind == "map" {
			vJo, _ := parseJsonObject(v.(map[string]interface{}))
			jo.Put(k, vJo)
			continue
		}

		if vKind == "slice" {
			vJa, _ := parseJsonArray(v.([]interface{}))
			jo.Put(fmt.Sprint(k), vJa)
			continue
		}

		jo.Put(fmt.Sprint(k), v)
	}

	return jo, nil
}

// 解析为JSON数组
func parseJsonArray(aa []interface{}) (*JSONArray, error) {
	ja := NewJSONArray()

	for _, v := range aa {
		vKind := reflect.TypeOf(v).Kind().String()

		if vKind == "map" {
			vJo, _ := parseJsonObject(v.(map[string]interface{}))
			ja.Add(vJo)
			continue
		}

		if vKind == "slice" {
			vJa, _ := parseJsonArray(v.([]interface{}))
			ja.Add(vJa)
			continue
		}

		ja.Add(v)
	}

	return ja, nil
}
