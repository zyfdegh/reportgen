package main

// sort and get top n unique biggest number, from bigger to smaller
func top(a []int, n int) (b []int) {
	lenA := len(a)
	if lenA > n {
		b = make([]int, n)
		for i := 0; i < lenA; i++ {
			if a[i] > b[n-1] && !contain(b, a[i]) {
				for j := 0; j < n-1; j++ {
					b[n-1-j] = a[i]
					if b[n-2-j] < b[n-1-j] {
						tmp := b[n-2-j]
						b[n-2-j] = b[n-1-j]
						b[n-1-j] = tmp
					} else {
						break
					}
				}
			}
		}
	}
	return
}

// return true if a in arr
func contain(arr []int, a int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == a {
			return true
		}
	}
	return false
}

// input float time, return int time
// 9/24==0.375
// 10/24==0.4166666
// So, if input 0.4, will return 9
func hour(f float32) (h int) {
	return int(f * 24)
}
