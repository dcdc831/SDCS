package jsontrans

import (
	"encoding/json"
	"fmt"
)

func JsonToMap(jstr string) map[string]interface{} {
	jsonMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(jstr), &jsonMap)
	if err != nil {
		fmt.Println("error:", err)
		return nil
	}
	return jsonMap
}

func MapToJson(m map[string]interface{}) string {
	jstr, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error:", err)
		// 返回一个空string
		return ""
	}
	return string(jstr)
}
