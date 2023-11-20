package node_test

import (
	"SDCS/node"
	"fmt"
	"testing"
)

/*
 * @brief: 测试节点功能
 */
func TestNode(t *testing.T) {
	n := node.NewNode(0, "9870")
	if ok := n.AddCache("key0", []string{"value1", "value2"}); ok == 0 {
		fmt.Println("add cache failed")
	}

	if n.Cache["key0"][0] == "value1" && n.Cache["key0"][1] == "value2" {
		fmt.Println("add cache success")
	} else {
		t.Errorf("添加内容错误")
	}

	if ok := n.AddCache("key0", []string{"value1", "value2"}); ok == 0 {
		fmt.Println("add cache failed")
	} else {
		t.Errorf("重复添加cache未能正确报错")
	}

	if ok := n.SetCache("key0", []string{"value3", "value4"}); ok == 1 {
		if n.Cache["key0"][0] == "value3" && n.Cache["key0"][1] == "value4" {
			fmt.Println("set cache success")
		} else {
			t.Errorf("修改内容错误")
		}
	} else {
		t.Errorf("修改cache未能成功")
	}

	if ok := n.GetCache("key0"); ok != nil {
		if ok[0] == "value3" && ok[1] == "value4" {
			fmt.Println("get cache success")
		} else {
			t.Errorf("获取内容错误")
		}
	} else {
		t.Errorf("获取cache未能成功")
	}

	if ok := n.DelCache("key1"); ok == 0 {
		fmt.Println("del cache failed")
	} else {
		t.Errorf("未能正确报错")
	}

	if ok := n.DelCache("key0"); ok == 1 {
		fmt.Println("已经正确删除")
		fmt.Println(n.Cache)
	} else {
		t.Errorf("未能正确删除")
	}
}
