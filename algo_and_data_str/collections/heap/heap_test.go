package heap

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"testing"
)

type TestSuite struct {
	suite.Suite
	h *heap
}

func TestRunSuit(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

// 함수가 실행할때마다
func (ts *TestSuite) SetupTest() {
	ts.h = NewHeap()
}

func (ts *TestSuite) TestHeap_Add() {
	ts.h.Add(1)
	ts.h.Add(2)
	ts.h.Add(3)

	assert.Equal(ts.T(), len(ts.h.node)-1, 3)
}
func (ts *TestSuite) TestHeap_Poll() {
	ts.h.Add(1)
	ts.h.Add(3)
	ts.h.Add(10)
	result := ts.h.Poll()

	assert.Equal(ts.T(), result, 10)
}

func (ts *TestSuite) TestHeap_Peek() {
	ts.h.Add(1)
	ts.h.Add(100)
	result := ts.h.Peek()

	assert.Equal(ts.T(), result, 100, "Peek error")
	assert.Equal(ts.T(), len(ts.h.node)-1, 2, "size error")
}

func (ts *TestSuite) TestHeap_Size() {
	ts.h.Add(1)
	ts.h.Add(1)
	ts.h.Add(1)
	ts.h.Add(1)
	assert.Equal(ts.T(), ts.h.Size(), 4, "Size error")
}
func (ts *TestSuite) TestHeap_Empty() {
	ts.h.Add(1)
	assert.Equal(ts.T(), ts.h.Empty(), false)
}
