package stack

type stack[T any] struct {
	nodes []T
	index int
}

// NewStack stack stack 생성
func NewStack[T any]() *stack[T] {
	return &stack[T]{nodes: make([]T, 0)}
}

// Push 요소 추가
func (s *stack[T]) Push(item T) {
	s.nodes = append(s.nodes, item)
	s.index++
}

// Empty 비어 있는지 확인
func (s *stack[T]) Empty() bool {
	return len(s.nodes) == 0
}

// IsNotEmpty 저장된 요소가 있는지 확인
func (s *stack[T]) IsNotEmpty() bool {
	return !s.Empty()
}

// Pop 요소 삭제
func (s *stack[T]) Pop() T {
	result := s.nodes[len(s.nodes)-1]
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.index--
	return result
}

// Peek 맨 위 요소 확인
func (s *stack[T]) Peek() T {
	return s.nodes[len(s.nodes)-1]
}

// Size 크기
func (s *stack[T]) Size() int {
	return s.index
}
