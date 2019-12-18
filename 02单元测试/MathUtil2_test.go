package mymath

import (
	"testing"
)

func TestGetSumManyTimes(t *testing.T) {
	answerMap := make(map[int]int)
	answerMap[10] = 55
	answerMap[5] = 15

	for n,answer := range answerMap{
		sum := GetSum(n)
		if sum != answer {
			//t.Log + t.FailNow()
			//打印日志 + 宣告失败并立即停止
			t.Fatalf("劳资是一个致命错误,1到%d的连续和应是%d，实际得到%d\n",n,answer,sum)
		}
	}
	t.Log("测试成功！")
}
