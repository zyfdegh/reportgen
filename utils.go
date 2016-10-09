package main

// sort and get n biggest number, from smaller to bigger
func top(a []int, n int) (b []int) {
	if len(a) < n {
		return
	}
	b = make([]int, n)
	for i := 0; i < len(a); i++ {
		if a[i] < b[0] {
			continue
		}
		b[0] = a[i]
		for j := 0; j < n-1; j++ {
			if b[j] > b[j+1] {
				tmp := b[j]
				b[j] = b[j+1]
				b[j+1] = tmp
			} else {
				break
			}
		}
	}
	return
}

// input float time, return int time
// 9/24==0.375
// 10/24==0.4166666
// So, if input 0.4, will return 9
func hour(f float32) (h int) {
	return int(f * 24)
}
