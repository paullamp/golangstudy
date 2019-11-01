package testing101

import (
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 3, 4, 5}
	expected := 13
	actual := Sum(numbers)
	if actual != expected {
		t.Error("测试失败")
	}
}
