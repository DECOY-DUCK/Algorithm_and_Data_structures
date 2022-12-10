package heap

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"testing"
)

type TestSuite struct {
	suite.Suite
	h *heap[int]
	p *heap[int]
}

func TestRunSuit(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

// 함수가 실행할때마다
func (ts *TestSuite) SetupTest() {
	ts.h = NewMaxHeap[int]()
	ts.p = NewMinHeap[int]()
}

func (ts *TestSuite) TestHeap_Add() {
	ts.h.Add(1)
	ts.h.Add(2)
	ts.h.Add(3)

	assert.Equal(ts.T(), 3, len(ts.h.node)-1, "maxHeap add error")

	ts.p.Add(1)
	ts.p.Add(2)
	ts.p.Add(3)

	assert.Equal(ts.T(), 3, len(ts.p.node)-1, "minHeap add error")

}
func (ts *TestSuite) TestHeap_Poll() {
	ts.h.Add(1)
	ts.h.Add(3)
	ts.h.Add(10)
	result := ts.h.Poll()

	assert.Equal(ts.T(), 10, result)

	ts.p.Add(1)
	ts.p.Add(3)
	ts.p.Add(10)
	result = ts.p.Poll()

	assert.Equal(ts.T(), 1, result)
}

func (ts *TestSuite) TestHeap_Peek() {
	ts.h.Add(1)
	ts.h.Add(100)
	result := ts.h.Peek()

	assert.Equal(ts.T(), 100, result, "Peek error")
	assert.Equal(ts.T(), 2, len(ts.h.node)-1, "size error")
}

func (ts *TestSuite) TestHeap_Size() {
	ts.h.Add(1)
	ts.h.Add(1)
	ts.h.Add(1)
	ts.h.Add(1)
	assert.Equal(ts.T(), 4, ts.h.Size(), "Size error")
}

func (ts *TestSuite) TestHeap_Empty() {
	ts.h.Add(1)
	assert.Equal(ts.T(), false, ts.h.Empty())
}
