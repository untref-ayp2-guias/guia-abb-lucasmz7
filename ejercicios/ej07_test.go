package ejercicios

import (
	"testing"
	"untref-ayp2/guia-abb/binarytree"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPostOrderIteratorEmpty(t *testing.T) {
	bst := binarytree.NewBinarySearchTree[int]()
	it := NewBSTPostOrderIterator(bst)
	require.NotNil(t, it)
	assert.False(t, it.HasNext())
	_, err := it.Next()
	assert.EqualError(t, err, "árbol vacío")
}

func TestPostOrderIteratorSingle(t *testing.T) {
	bst := binarytree.NewBinarySearchTree[int]()
	bst.Insert(10)
	it := NewBSTPostOrderIterator(bst)
	require.NotNil(t, it)
	assert.True(t, it.HasNext())
	val, err := it.Next()
	assert.NoError(t, err)
	assert.Equal(t, 10, val)
	assert.False(t, it.HasNext())
}

func TestPostOrderIteratorMultiple(t *testing.T) {
	bst := binarytree.NewBinarySearchTree[int]()
	values := []int{15, 10, 20, 8, 12, 16, 25}
	for _, v := range values {
		bst.Insert(v)
	}
	it := NewBSTPostOrderIterator(bst)
	require.NotNil(t, it)
	var result []int
	for it.HasNext() {
		v, err := it.Next()
		assert.NoError(t, err)
		result = append(result, v)
	}
	assert.Equal(t, []int{8, 12, 10, 16, 25, 20, 15}, result)
}

func TestPostOrderIteratorLeftSkewed(t *testing.T) {
	bst := binarytree.NewBinarySearchTree[int]()
	for _, v := range []int{3, 2, 1} {
		bst.Insert(v)
	}
	it := NewBSTPostOrderIterator(bst)
	require.NotNil(t, it)
	var result []int
	for it.HasNext() {
		v, err := it.Next()
		assert.NoError(t, err)
		result = append(result, v)
	}
	assert.Equal(t, []int{1, 2, 3}, result)
}

func TestPostOrderIteratorRightSkewed(t *testing.T) {
	bst := binarytree.NewBinarySearchTree[int]()
	for _, v := range []int{1, 2, 3} {
		bst.Insert(v)
	}
	it := NewBSTPostOrderIterator(bst)
	require.NotNil(t, it)
	var result []int
	for it.HasNext() {
		v, err := it.Next()
		assert.NoError(t, err)
		result = append(result, v)
	}
	assert.Equal(t, []int{3, 2, 1}, result)
}
