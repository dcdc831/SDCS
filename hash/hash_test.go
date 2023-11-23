package hash_test

import (
	"SDCS/hash"
	"testing"
)

func TestHashGetCacheNode(t *testing.T) {
	if node, ok := hash.GetCacheNode("keyTest"); ok == true {
		t.Log(node)
	} else {
		t.Error("GetCacheNode failed")
	}
}
