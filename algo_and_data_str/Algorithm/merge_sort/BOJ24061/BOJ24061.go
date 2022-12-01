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
	count int
	n     int
)

func init() {
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
}

func main() {
	defer w.Flush()
	k := 0
	fmt.Fscanln(r, &n, &k)
	count = k
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &arr[i])
	}

	mergeSort(0, n-1)
	if count > 0 {
		fmt.Fprintln(w, -1)
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
		count--
		arr[i] = temp[t]
		i++
		t++
		if count == 0 {
			for m := 0; m < n; m++ {
				fmt.Fprintf(w, "%d ", arr[m])
			}
		}
	}
}
