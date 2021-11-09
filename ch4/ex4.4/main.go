func rotate(nums []int, k int) {
	// we want to just store a temp value to update each value one-by-one
	// however if len(s) % n = 0 then we will have to do it in d (divisor) sets
	d := gcd(len(nums), k)
	// loop through each of the d sets
	for i := 0; i < d; i++ {
		temp := nums[i+k]
		// for each set loop through the index of each j in the dth set
		for j := i + k; j != i; j = (j + k) % len(nums) {
			// swap temp and n-along
			nums[(j+k)%len(nums)], temp = temp, nums[(j+k)%len(nums)]
		}
		nums[(i+k)%len(nums)] = temp
	}
}

func gcd(a, b int) int {
	// 12 % 15 = 12
	// 15 % 12 = 3
	// 12 % 3 = 0
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}