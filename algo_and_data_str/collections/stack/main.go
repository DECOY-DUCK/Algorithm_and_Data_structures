package main

import (
	"bufio"
	"fmt"
	"os"
)

type Stack[T any] struct {
	nodes []T
	index int
}

// NewStack stack stack 생성
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{nodes: make([]T, 0)}
}

// Push 요소 추가
func (s *Stack[T]) Push(item T) {
	s.nodes = append(s.nodes, item)
	s.index++
}

// Empty 비어 있는지 확인
func (s *Stack[T]) Empty() bool {
	return len(s.nodes) == 0
}

// IsNotEmpty 저장된 요소가 있는지 확인
func (s *Stack[T]) IsNotEmpty() bool {
	return !s.Empty()
}

// Pop 요소 삭제
func (s *Stack[T]) Pop() T {
	result := s.nodes[len(s.nodes)-1]
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.index--
	return result
}

// 맨 위 요소 확인
func (s *Stack[T]) peek() T {
	return s.nodes[len(s.nodes)-1]
}

// 사이즈 확인
func (s *Stack[T]) size() int {
	return s.index
}

// push X: 정수 X를 스택에 넣는 연산이다.
// pop: 스택에서 가장 위에 있는 정수를 빼고, 그 수를 출력한다. 만약 스택에 들어있는 정수가 없는 경우에는 -1을 출력한다.
// size: 스택에 들어있는 정수의 개수를 출력한다.
// empty: 스택이 비어있으면 1, 아니면 0을 출력한다.
// top: 스택의 가장 위에 있는 정수를 출력한다. 만약 스택에 들어있는 정수가 없는 경우에는 -1을 출력한다.
var (
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

func main() {
	defer w.Flush()
	var n int
	var stack = NewStack[int]()
	fmt.Fscanln(r, &n)
	for i := 0; i < n; i++ {
		var command string
		fmt.Fscan(r, &command)
		switch command {
		case "push":
			var num int
			fmt.Fscan(r, &num)
			stack.Push(num)
		case "pop":
			if stack.Empty() {
				fmt.Fprintln(w, -1)
			} else {
				fmt.Fprintln(w, stack.Pop())
			}
		case "size":
			fmt.Fprintln(w, stack.size())
		case "empty":
			if stack.Empty() {
				fmt.Fprintln(w, 1)
			} else {
				fmt.Fprintln(w, 0)
			}
		case "top":
			if stack.Empty() {
				fmt.Fprintln(w, -1)
			} else {
				fmt.Fprintln(w, stack.peek())
			}
		}

	}
}
