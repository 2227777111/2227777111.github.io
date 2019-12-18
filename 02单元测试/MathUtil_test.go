package mymath

import "testing"

/*
测试用例

	//标记失败但继续运行该测试
	t.Fail()
	//标记失败并立刻终止该测试
	t.FailNow()

	//打印测试日志
	t.Log("劳资是日志")
	t.Logf("劳资是%s/n","错误原因")

	//t.Log("劳资是日志") + t.Fail()
	t.Error("劳资是日志")
	//t.Logf("劳资是%s/n","错误原因") + t.Fail()
	t.Errorf("劳资是%s/n","错误原因")

	//t.Log("劳资是日志") + t.FailNow()
	t.Fatal("劳资是日志")
	//t.Logf("劳资是%s/n","错误原因") + t.FailNow()
	t.Fatalf("劳资是%s/n","错误原因")

*/
func TestGetSum(t *testing.T) {
	sum := GetSum(10)
	if sum != 55{
		//打印日志，参数用法同fmt.Printf(...)
		t.Logf("期望%d,实际%d\n",55,sum)

		//宣告失败并立即终止
		t.FailNow()
	}
	t.Log("测试成功！")
}

func TestGetSumRecursively(t *testing.T) {
	sum := GetSumRecursively(10)
	if sum != 55{
		t.Fatalf("期望%d,实际%d\n",55,sum)
	}
	t.Log("测试成功！")
}