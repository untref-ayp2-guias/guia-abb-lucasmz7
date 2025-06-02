package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestINTERNALNewStack(t *testing.T) {
	s := NewStack[int]()
	assert.NotNil(t, s)
}

func TestINTERNALStackPush(t *testing.T) {
	s := NewStack[int]()

	s.Push(1)
	assert.False(t, s.IsEmpty())
}

func TestINTERNALStackTop(t *testing.T) {
	s := NewStack[int]()

	s.Push(1)
	v, err := s.Top()
	assert.Equal(t, 1, v)
	assert.NoError(t, err)
}

func TestINTERNALStackTopWhenEmpty(t *testing.T) {
	s := NewStack[int]()

	_, err := s.Top()
	assert.EqualError(t, err, "pila vacía")
}

func TestINTERNALStackPop(t *testing.T) {
	s := NewStack[int]()

	s.Push(1)

	v, err := s.Pop()
	assert.Equal(t, 1, v)
	assert.NoError(t, err)
}

func TestINTERNALStackPopWhenEmpty(t *testing.T) {
	s := NewStack[int]()

	_, err := s.Pop()
	assert.EqualError(t, err, "pila vacía")
}

func TestINTERNALStackIsEmpty(t *testing.T) {
	s := NewStack[int]()
	assert.True(t, s.IsEmpty())

	s.Push(1)
	assert.False(t, s.IsEmpty())
}
