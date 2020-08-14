package json

import (
	"fmt"
	"strconv"
)

// JSON数组
type JSONArray struct {
	a []interface{}
}

func NewJSONArray() *JSONArray {
	return &JSONArray{}
}

func (ja *JSONArray) Size() int {
	return len(ja.a)
}

func (ja *JSONArray) Add(e interface{}) {
	ja.a = append(ja.a, e)
}

func (ja *JSONArray) GetJSONObject(index int) *JSONObject {
	return ja.a[index].(*JSONObject)
}

func (ja *JSONArray) GetJSONArray(index int) *JSONArray {
	return ja.a[index].(*JSONArray)
}

func (ja *JSONArray) GetBool(index int) bool {
	str := ja.GetString(index)
	if str == "" {
		return false
	}
	value, _ := strconv.ParseBool(str)
	return value
}

func (ja *JSONArray) GetString(index int) string {
	value := ja.a[index]
	if value == nil {
		return ""
	}
	return fmt.Sprintf("%v", value)
}

func (ja *JSONArray) GetInt(index int) int {
	str := ja.GetString(index)
	if str == "" {
		return 0
	}
	value, _ := strconv.Atoi(str)
	return value
}

func (ja *JSONArray) GetFloat(index int) float64 {
	str := ja.GetString(index)
	if str == "" {
		return 0
	}
	value, _ := strconv.ParseFloat(str, 64)
	return value
}

func (ja *JSONArray) GetStringArray() []string {
	var aa []string
	for _, v := range ja.a {
		aa = append(aa, v.(string))
	}
	return aa
}

func (ja *JSONArray) GetIntArray() []int {
	var aa []int
	for _, v := range ja.a {
		aa = append(aa, v.(int))
	}
	return aa
}
