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

func nextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
}
