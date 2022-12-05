package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	r    = bufio.NewReader(os.Stdin)
	w    = bufio.NewWriter(os.Stdout)
	tree []int
	arr  []int
)

func main() {
	defer w.Flush()
	var n, m, k int
	fmt.Fscanln(r, &n, &m, &k)

	arr = make([]int, n)
	tree = make([]int, n*4)

	for i := range arr {
		fmt.Fscanln(r, &arr[i])
	}

	makeTree(0, n-1, 1)
	//쿼리
	for i := 0; i < m+k; i++ {
		var a, b, c int
		fmt.Fscanln(r, &a, &b, &c)

		switch a {
		//변경
		case 1:
			update(0, n-1, 1, b-1, c-arr[b-1])
			arr[b-1] = c
		case 2:
			value := query(0, n-1, 1, b-1, c-1)
			fmt.Fprintln(w, value)
		}
	}
}

func makeTree(start, end, node int) int {
	if start == end {
		tree[node] = arr[start]
		return tree[node]
	}
	mid := (start + end) / 2
	tree[node] = makeTree(start, mid, node*2) + makeTree(mid+1, end, node*2+1)
	return tree[node]
}

func update(start, end, node, what, value int) {
	//범위 밖
	if what < start || what > end {
		return
	}
	tree[node] += value
	if start == end {
		return
	}
	mid := (start + end) / 2
	update(start, mid, node*2, what, value)
	update(mid+1, end, node*2+1, what, value)
}

func query(start, end, node, left, right int) int {
	//구간 밖인경우
	if left > end || right < start {
		return 0
	}
	if left <= start && right >= end {
		return tree[node]
	}
	mid := (start + end) / 2
	return query(start, mid, node*2, left, right) + query(mid+1, end, node*2+1, left, right)
}
