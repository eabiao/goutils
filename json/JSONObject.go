package json

import (
	"fmt"
	"strconv"
)

// JSON对象
type JSONObject struct {
	m map[string]interface{}
}

func NewJSONObject() *JSONObject {
	return &JSONObject{make(map[string]interface{})}
}

func (j *JSONObject) Size() int {
	return len(j.m)
}

func (j *JSONObject) ContainsKey(key string) bool {
	_, exist := j.m[key]
	return exist
}

func (j *JSONObject) Put(key string, value interface{}) {
	j.m[key] = value
}

func (j *JSONObject) Get(key string) interface{} {
	value, _ := j.m[key]
	return value
}

func (j *JSONObject) GetJSONObject(key string) *JSONObject {
	return j.Get(key).(*JSONObject)
}

func (j *JSONObject) GetJSONArray(key string) *JSONArray {
	return j.Get(key).(*JSONArray)
}

func (j *JSONObject) GetBool(key string) bool {
	str := j.GetString(key)
	if str == "" {
		return false
	}
	value, _ := strconv.ParseBool(str)
	return value
}

func (j *JSONObject) GetString(key string) string {
	value := j.Get(key)
	if value == nil {
		return ""
	}
	return fmt.Sprintf("%v", value)
}

func (j *JSONObject) GetInt(key string) int {
	str := j.GetString(key)
	if str == "" {
		return 0
	}
	value, _ := strconv.Atoi(str)
	return value
}

func (j *JSONObject) GetFloat(key string) float64 {
	str := j.GetString(key)
	if str == "" {
		return 0
	}
	value, _ := strconv.ParseFloat(str, 64)
	return value
}

func (j *JSONObject) KeySet() []string {
	var set []string
	for k := range j.m {
		set = append(set, k)
	}
	return set
}
