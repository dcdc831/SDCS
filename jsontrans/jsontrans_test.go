package jsontrans_test

import (
	"SDCS/jsontrans"
	"fmt"
	"testing"
)

func TestTrans(t *testing.T) {
	jstr := `{"tasks": ["task 1", "task 2", "task 3"]}`
	m := jsontrans.JsonToMap(jstr)
	if m["tasks"].([]interface{})[0] == "task 1" {
		fmt.Println("json to map success")
	} else {
		t.Errorf("json to map failed")
	}
	if m["tasks"].([]interface{})[1] == "task 2" {
		fmt.Println("json to map success")
	} else {
		t.Errorf("json to map failed")
	}
	if m["tasks"].([]interface{})[2] == "task 3" {
		fmt.Println("json to map success")
	} else {
		t.Errorf("json to map failed")
	}

	jstr_ret := jsontrans.MapToJson(m)
	fmt.Println(jstr)
	fmt.Println(jstr_ret)

}
