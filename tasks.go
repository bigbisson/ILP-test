package main

import (
	"errors"
	"sort"
	"strconv"
	"strings"
)

// Task1 - GO1
func GO1(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	hasPos, hasNeg := false, false
	for _, v := range arr {
		if v > 0 {
			hasPos = true
		} else if v < 0 {
			hasNeg = true
		}
		if hasPos && hasNeg {
			return 0
		}
	}
	if hasPos && !hasNeg {
		set := make(map[int]bool)
		for _, v := range arr {
			if v > 0 {
				set[v] = true
			}
		}
		for i := 1; ; i++ {
			if !set[i] {
				return i
			}
		}
	}
	set := make(map[int]bool)
	for _, v := range arr {
		if v < 0 {
			set[-v] = true
		}
	}
	for i := 1; ; i++ {
		if !set[i] {
			return -i
		}
	}
}

// Task2 - GO2
func GO2(exclude []int, nums ...int) (int, error) {
	excl := make(map[rune]bool)
	for _, d := range exclude {
		if d >= 0 && d <= 9 {
			excl[rune('0'+d)] = true
		}
	}
	var b strings.Builder
	for _, n := range nums {
		b.WriteString(strconv.Itoa(n))
	}
	s := b.String()
	rs := reverseString(s)
	var out strings.Builder
	for _, ch := range rs {
		if !excl[ch] {
			out.WriteRune(ch)
		}
	}
	outStr := out.String()
	if outStr == "" {
		return 0, errors.New("not found")
	}
	val, err := strconv.Atoi(outStr)
	if err != nil {
		return 0, err
	}
	return val, nil
}

func reverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// Task3 - GO3
func GO3(numsA, numsB []int) []int {
	setB := make(map[int]bool)
	for _, v := range numsB {
		setB[v] = true
	}
	setRes := make(map[int]bool)
	for _, v := range numsA {
		if setB[v] {
			setRes[v] = true
		}
	}
	res := make([]int, 0, len(setRes))
	for k := range setRes {
		res = append(res, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(res)))
	return res
}

// Task4 - GO4
func GO4(num int) map[int]int {
	res := make(map[int]int)
	s := strconv.Itoa(num)
	for _, ch := range s {
		d := int(ch - '0')
		res[d]++
	}
	return res
}

// Task5 - GO5 (concurrent multiplications)
func GO5(nums ...int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	pairs := (n + 1) / 2
	ch := make(chan int, pairs)

	for i := 0; i < n; i += 2 {
		i := i
		go func() {
			if i+1 < n {
				ch <- nums[i] * nums[i+1]
			} else {
				ch <- nums[i] * nums[i]
			}
		}()
	}
	sum := 0
	for i := 0; i < pairs; i++ {
		sum += <-ch
	}
	return sum
}
