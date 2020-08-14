package test

import (
	"fmt"
	"github.com/eabiao/goutils/json"
	"log"
	"testing"
)

func TestJSONObject(t *testing.T) {
	var jsonText = `{
	  "name": "aa",
	  "age": 18,
      "isBoy": true
	}`

	jsonObj, err := json.ToJSONObject([]byte(jsonText))
	if err != nil {
		log.Println(err)
	}
	fmt.Println("name", jsonObj.GetString("name"))
	fmt.Println("age", jsonObj.GetInt("age"))
	fmt.Println("isBoy", jsonObj.GetBool("isBoy"))
}

func TestJSONArray1(t *testing.T) {
	var jsonText = `[
	  {
		"name": "aa",
		"age": 18,
		"isBoy": true
	  },
	  {
		"name": "bb",
		"age": 16,
		"isBoy": false
	  }
	]`

	jsonArray, err := json.ToJSONArray([]byte(jsonText))
	if err != nil {
		log.Println(err)
	}

	for i := 0; i < jsonArray.Size(); i++ {
		jsonObj := jsonArray.GetJSONObject(i)
		fmt.Println("name", jsonObj.GetString("name"))
		fmt.Println("age", jsonObj.GetInt("age"))
		fmt.Println("isBoy", jsonObj.GetBool("isBoy"))
	}
}

func TestJSONArray2(t *testing.T) {
	var jsonText = `[
	  1,
	  "aa",
	  "1.123",
	  true,
	  [1, 2, 3, 4]
	]`

	jsonArray, err := json.ToJSONArray([]byte(jsonText))
	if err != nil {
		log.Println(err)
	}

	fmt.Println(jsonArray.GetInt(0))
	fmt.Println(jsonArray.GetString(1))
	fmt.Println(jsonArray.GetFloat(2))
	fmt.Println(jsonArray.GetBool(3))
	fmt.Println(jsonArray.GetJSONArray(4).GetInt(0))
}
