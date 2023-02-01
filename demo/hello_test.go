package demo

import (
	"testing"
)

func TestHello(t *testing.T) {
	result := GetHello("山田さん")
	expect := "こんにちは、山田さん！"
	if result != expect {
		t.Error("\n実際：", result, "\n理想：", expect)
	}
	t.Log("TestHello終了")
}