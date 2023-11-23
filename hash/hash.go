package hash

import (
	"github.com/serialx/hashring"
)

const (
	port0 = "9527"
	port1 = "9528"
	port2 = "9529"
)

func GetCacheNode(key string) (string, bool) {
	ring := hashring.New([]string{})
	ring = ring.AddNode(port0)
	ring = ring.AddNode(port1)
	ring = ring.AddNode(port2)

	node, ok := ring.GetNode(key)
	if !ok {
		return "", false
	}
	return node, true
}
