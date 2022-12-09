package heap

type heap struct {
	node []int
}

func NewHeap() *heap {
	return &heap{[]int{0}}
}

// Add heap에 요소를 추가한다.
// num: 추가하고 싶은 값
func (s *heap) Add(num int) {
	// 맨 마지막 인덱스에 요소 추가
	s.node = append(s.node, num)
	// 연산
	index := len(s.node) - 1
	for index > 1 {
		parent := index / 2
		// 스왑
		if s.node[parent] < s.node[index] {
			temp := s.node[parent]
			s.node[parent] = s.node[index]
			s.node[index] = temp
		}
		index = parent
	}
}

// Poll 요소를 반환하고 삭제한다.
func (s *heap) Poll() int {
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
		if parent*2+1 < len(s.node) && min > s.node[parent*2+1] {
			min = s.node[parent*2+1]
			minIndex = parent * 2
		}
		//부모가 더크면 종료
		if min < s.node[parent*2] {
			break
		}
		//스왑
		temp := s.node[minIndex]
		s.node[minIndex] = s.node[parent]
		s.node[parent] = temp
		//밑으로
		parent = minIndex
	}
	return result
}
func (s *heap) Peek() int {
	if len(s.node) > 0 {
		return s.node[1]
	} else {
		return -1
	}
}
func (s *heap) Size() int {
	return len(s.node) - 1
}
func (s *heap) Empty() bool {
	return len(s.node) == 1
}
