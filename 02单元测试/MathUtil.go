package mymath

/**/
func GetSum(n int) int {
	var sum = 0
	for i := 1;i<n+1;i++{
		sum += i
	}
	return sum
}

func GetSumRecursively(n int) int {
	if n == 1{
		return 1
	}
	return n + GetSumRecursively(n-1)
}
