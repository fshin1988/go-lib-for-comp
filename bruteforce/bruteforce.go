package bruteforce

func bitSearch(n int) {
	for bit := 0; bit < 1<<uint(n); bit++ {
		arr := make([]bool, n)
		for i := 0; i < n; i++ {
			if 1<<uint(i)&bit > 0 {
				arr[i] = true
			}
		}
		// Do something by using arr
	}
}

func permutations(nums []int, n int) {
	var arr []int
	seen := make(map[int]bool)

	var search func(next int)
	search = func(next int) {
		arr = append(arr, next)
		seen[next] = true
		if len(arr) >= n {
			// Do something by using arr
			return
		}
		for _, next := range nums {
			if seen[next] {
				continue
			}
			search(next)
			arr = arr[:len(arr)-1]
			seen[next] = false
		}
	}
	for _, v := range nums {
		search(v)
		arr = arr[:len(arr)-1]
		seen[v] = false
	}
}
