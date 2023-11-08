package main

import (
	"SDCS/node"
	"fmt"
)

func main() {
	n := node.NewNode(0, "9870")
	if ok := n.AddCache("key0", []string{"value1", "value2"}); ok == 0 {
		fmt.Println("add cache failed")
	}
	if ok := n.AddCache("key0", []string{"value1", "value2"}); ok == 0 {
		fmt.Println("add cache failed")
	}
	fmt.Println(n.Cache)
}
