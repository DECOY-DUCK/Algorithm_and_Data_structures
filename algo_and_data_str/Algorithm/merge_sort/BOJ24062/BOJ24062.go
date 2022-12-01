package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	r     *bufio.Reader
	w     *bufio.Writer
	arr   = [500_001]int{0}
	temp  = [500_001]int{0}
	check = [500_001]int{0}
	count = 0
	n     int
)

func init() {
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
}

func main() {
	defer w.Flush()

	fmt.Fscanln(r, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &arr[i])
	}

	for i := 0; i < n; i++ {
		fmt.Fscan(r, &check[i])
	}
	mergeSort(0, n)
	if count == n-1 {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}

}

func mergeSort(p, r int) {
	if p < r {
		q := (p + r) / 2
		mergeSort(p, q)
		mergeSort(q+1, r)
		merge(p, q, r)
	}
}

func merge(p, q, r int) {
	i, j, t := p, q+1, 1
	for i <= q && j <= r {
		if arr[i] <= arr[j] {
			temp[t] = arr[i]
			t++
			i++
		} else {
			temp[t] = arr[j]
			t++
			j++
		}
	}
	for i <= q {
		temp[t] = arr[i]
		t++
		i++
	}
	for j <= r {
		temp[t] = arr[j]
		t++
		j++
	}
	i, t = p, 1
	for i <= r {
		arr[i] = temp[t]
		i++
		t++
	}
	if count < n-1 {
		for m := count; m < n; m++ {
			if arr[m] == check[m] {
				count = m
			} else {
				break
			}
		}
	} else {
		return
	}

}
