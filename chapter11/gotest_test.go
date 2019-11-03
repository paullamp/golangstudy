package main

import (
	"testing"
)

func Test_Division(t *testing.T) {
	if i, e := Division(6, 2); i != 3 || e != nil {
		t.Error("Test failed")
	} else {
		t.Log("Test passed")
	}
}
