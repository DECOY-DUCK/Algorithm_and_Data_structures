package stack

type Stack[T any] struct {
	nodes []T
}

// NewStack stack stack 생성
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{nodes: make([]T, 0)}
}

// Push 요소 추가
func (s *Stack[T]) Push(item T) {
	s.nodes = append(s.nodes, item)
}

// Empty 비어 있는지 확인
func (s *Stack[T]) Empty() bool {
	return len(s.nodes) == 0
}

// IsNotEmpty 저장된 요소가 있는지 확인
func (s *Stack[T]) IsNotEmpty() bool {
	return len(s.nodes) != 0
}

// Pop 요소 삭제
func (s *Stack[T]) Pop() T {
	result := s.nodes[len(s.nodes)-1]
	s.nodes = s.nodes[:len(s.nodes)-1]
	return result
}

// 맨 위 요소 확인
func (s *Stack[T]) peek() T {
	return s.nodes[len(s.nodes)-1]
}
