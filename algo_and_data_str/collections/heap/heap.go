package heap

type number interface {
	int8 | int16 | int32 | int64 | int
}
type heap[T any] struct {
	node  []T
	check func(a, b T) bool
}

func NewHeap[T any](f func(a, b T) bool) *heap[T] {
	return &heap[T]{make([]T, 1), f}
}
func NewMaxHeap[T number]() *heap[T] {
	return &heap[T]{make([]T, 1), func(x T, y T) bool { return x < y }}
}
func NewMinHeap[T number]() *heap[T] {
	return &heap[T]{make([]T, 1), func(x T, y T) bool { return x > y }}
}

// Add heap에 요소를 추가한다.
// num: 추가하고 싶은 값
func (s *heap[T]) Add(num T) {
	// 맨 마지막 인덱스에 요소 추가
	s.node = append(s.node, num)
	// 연산
	index := len(s.node) - 1
	for index > 1 {
		parent := index / 2

		if s.check(s.node[parent], s.node[index]) {
			// 스왑
			s.swap(parent, index)
			//temp := s.node[parent]
			//s.node[parent] = s.node[index]
			//s.node[index] = temp
		}
		index = parent
	}
}

// Poll 요소를 반환하고 삭제한다.
func (s *heap[T]) Poll() T {
	result := s.node[1]
	//맨마지막 요소 루트노드로
	s.node[1] = s.node[len(s.node)-1]
	// 크기 줄이기
	s.node = append(s.node[0 : len(s.node)-1])
	parent := 1

	for parent*2 < len(s.node) {
		//왼쪽부터
		min := s.node[parent*2]
		minIndex := parent * 2
		//오른쪽이 더크면 변경

		if parent*2+1 < len(s.node) && s.check(min, s.node[parent*2+1]) {
			min = s.node[parent*2+1]
			minIndex = parent*2 + 1
		}
		//부모가 더크면 종료
		if s.check(min, s.node[parent]) {
			break
		}
		//스왑
		s.swap(minIndex, parent)
		//temp := s.node[minIndex]
		//s.node[minIndex] = s.node[parent]
		//s.node[parent] = temp
		//밑으로
		parent = minIndex
	}
	return result
}

// Peek 최상 위 요소 반환
func (s *heap[T]) Peek() T {

	return s.node[1]

}

// Size heap 크기 반환
func (s *heap[T]) Size() int {
	return len(s.node) - 1
}

// Empty 공백 확인
func (s *heap[T]) Empty() bool {
	return len(s.node) == 1
}

func (s *heap[T]) swap(i, j int) {
	s.node[i], s.node[j] = s.node[j], s.node[i]
}
