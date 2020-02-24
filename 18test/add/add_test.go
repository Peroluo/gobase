package add

import "testing"

func TestAdd(t *testing.T) {
	got := Add(1, 3)
	want := 4
	if got != want {
		// 测试用例有问题
		t.Fatalf("失败了")
		t.Errorf("wangt:%v,got:%v \n", want, got)
	}
	// 执行go test
	// PASS
	// ok      github.com/gobase/18test/add    2.915s
}
