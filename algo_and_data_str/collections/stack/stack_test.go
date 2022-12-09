package stack

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"reflect"
	"testing"
)

type TestSuite struct {
	suite.Suite
	s *stack[int]
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

// suite에서 각 테스트 실행 전에 실행된다
func (ts *TestSuite) SetupTest() {
	fmt.Println("SetupTest :: run setup test")
	ts.s = NewStack[int]()
}
func TestNewStack(t *testing.T) {
	s := NewStack[int]()

	if reflect.TypeOf(s) != reflect.TypeOf(&stack[int]{}) {
		t.Errorf("NewStack error")
	}
}

func TestStack_Push(t *testing.T) {
	s := NewStack[int]()
	s.Push(1)

	if len(s.nodes) != 1 {
		t.Errorf("push error")
	}
}

func TestStack_Empty(t *testing.T) {
	s := NewStack[int]()
	s.Push(1)

	if s.Empty() {
		t.Errorf("Empty error")
	}
}
func TestStack_Peek(t *testing.T) {
	s := NewStack[int]()
	s.Push(1)
	s.Push(2)

	if s.Peek() != 2 {
		t.Errorf("peek error")
	}
}
func TestStack_Pop(t *testing.T) {
	s := NewStack[int]()
	s.Push(1)
	result := s.Pop()
	if result != 1 {
		t.Errorf("Pop error")
	}
}
func TestStack_size(t *testing.T) {
	s := NewStack[int]()
	s.Push(1)
	s.Push(1)
	s.Push(1)
	s.Push(1)

	if s.Size() != 4 {
		t.Errorf("size error")
	}
}
