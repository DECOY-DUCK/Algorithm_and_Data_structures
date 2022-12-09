package segementTree

type SegmentTree struct {
	arr  []int
	tree []int
}

func New(arr []int) *SegmentTree {
	return &SegmentTree{arr, make([]int, 4*len(arr))}
}

// build 세그먼트 트리를 채우는 함수
// start : 배열의 시작 인덱스, end : 배열의 마지막 인덱스
// node : 세그먼트 트리의 인덱스 (1부터 시작)
func (s *SegmentTree) build(start, end, node int) {
	// 리프 노드
	if start == end {
		s.tree[node] = s.arr[start]
		return
	}
	mid := (start + end) / 2
	// 좌측 노드와 우측 노드를 채워주면서 부모 노드의 값도 채워준다.
	s.build(start, mid, 2*node)
	s.build(mid+1, end, 2*node+1)
	s.tree[node] = s.tree[2*node] + s.tree[2*node+1]
	return
}

// query 구간 합을 구하는 함수
// start : 시작 인덱스, end : 마지막 인덱스
// left, right : 구간 합을 구하고자 하는 범위
func (s *SegmentTree) query(start, end, node, left, right int) int {
	if (left <= start) && (end <= right) {
		return s.tree[node]
	}
	if (right < start) || (end < left) {
		return 0
	}
	mid := (start + end) / 2
	leftSum := s.query(start, mid, 2*node, left, right)
	rightSum := s.query(mid+1, end, 2*node+1, left, right)
	return leftSum + rightSum
}

// update 특정 원소의 값을 수정하는 함수
// start : 시작 인덱스
// end : 마지막 인덱스
// node : 세그먼트 트리의 인덱스 (1부터 시작)
// i : 구간 합을 수정하고자 하는 노드
// diff : 원래값-수정 값
func (s *SegmentTree) update(start, end, node, i, diff int) {
	//범위에서 벗어남
	if (i < start) || (end < i) {
		return
	}
	s.tree[node] += diff
	if start != end {
		mid := (start + end) / 2
		s.update(start, mid, 2*node, i, diff)
		s.update(mid+1, end, 2*node+1, i, diff)
	}
	return
}
