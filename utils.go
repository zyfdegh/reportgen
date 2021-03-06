package main

import (
	"log"
	"regexp"
	"strconv"
)

// sort and get top n biggest number, from bigger to smaller
func top(a []CodeFreq, n int) (b []CodeFreq) {
	lenA := len(a)
	if lenA > n {
		b = make([]CodeFreq, n)
		for i := 0; i < lenA; i++ {
			if a[i].Freq > b[n-1].Freq {
				for j := 0; j < n-1; j++ {
					b[n-1-j].Freq = a[i].Freq
					b[n-1-j].Code = a[i].Code
					if b[n-2-j].Freq < b[n-1-j].Freq {
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

func scan(arr []CodeFreq, code int) (contain bool, index int) {
	for i, freq := range arr {
		if freq.Code == code {
			return true, i
		}
	}
	return false, -1
}

// input float time, return int time
// 9/24==0.375
// 10/24==0.4166666
// So, if input 0.4, will return 9
func hour(f float32) (h int) {
	return int(f * 24)
}

// match and return 6 number in string
func extractSixNum(str string) string {
	reg := regexp.MustCompile("[0-9]+")
	arr := reg.FindAllString(str, -1)
	for _, s := range arr {
		if len(s) == 6 {
			return s
		}
	}
	return ""
}

func parseInt(str string) (n int, err error) {
	n64, err := strconv.ParseInt(str, 10, 0)
	if err != nil {
		log.Printf("parse string to int error: %v\n", err)
		return
	}
	return int(n64), nil
}

func parseFloat(str string) (f float32, err error) {
	f64, err := strconv.ParseFloat(str, 32)
	if err != nil {
		log.Printf("parse string to float error:%v\n", err)
		return
	}
	return float32(f64), nil
}
