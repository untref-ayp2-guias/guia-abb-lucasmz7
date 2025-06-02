package binarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestINTERNALBSTNuevo(t *testing.T) {
	bstree := NewBinarySearchTree[int]()

	assert.Equal(t, 0, bstree.Size())
	assert.Nil(t, bstree.GetRoot())

	minValue, errMin := bstree.FindMin()
	assert.Zero(t, minValue)
	require.Error(t, errMin)

	maxValue, errMax := bstree.FindMax()
	assert.Zero(t, maxValue)
	require.Error(t, errMax)
}

func TestINTERNALBSTInsertarUnElemento(t *testing.T) {
	bstree := NewBinarySearchTree[int]()
	bstree.Insert(4)

	assert.Equal(t, 1, bstree.Size())
}

func TestINTERNALBSTObtenerRaiz(t *testing.T) {
	bstree := NewBinarySearchTree[int]()
	bstree.Insert(4)

	assert.Equal(t, 4, bstree.GetRoot().GetData())
}

func TestINTERNALBSTBuscarElementoExistente(t *testing.T) {
	bstree := NewBinarySearchTree[int]()
	bstree.Insert(4)
	bstree.Insert(1)
	bstree.Insert(2)
	bstree.Insert(5)
	bstree.Insert(3)

	assert.True(t, bstree.Search(2))
	assert.False(t, bstree.Search(6))
}

func TestINTERNALBSTBuscarElementoInexistente(t *testing.T) {
	bstree := NewBinarySearchTree[int]()
	bstree.Insert(4)
	bstree.Insert(1)
	bstree.Insert(2)
	bstree.Insert(5)
	bstree.Insert(3)

	assert.False(t, bstree.Search(6))
}

func TestINTERNALBSTBuscarMinMax(t *testing.T) {
	bstree := NewBinarySearchTree[int]()
	bstree.Insert(4)
	bstree.Insert(1)
	bstree.Insert(2)
	bstree.Insert(5)
	bstree.Insert(3)

	minValue, errMin := bstree.FindMin()
	assert.Equal(t, 1, minValue)
	require.NoError(t, errMin)

	maxValue, errMax := bstree.FindMax()
	assert.Equal(t, 5, maxValue)
	require.NoError(t, errMax)
}

func TestINTERNALBSTBorrarRaiz(t *testing.T) {
	bstree := NewBinarySearchTree[int]()
	bstree.Insert(4)
	bstree.Insert(1)
	bstree.Insert(2)
	bstree.Insert(5)
	bstree.Insert(3)

	assert.True(t, bstree.Search(4))
	assert.Equal(t, 5, bstree.Size())

	bstree.Remove(4)

	assert.False(t, bstree.Search(4))
	assert.Equal(t, 4, bstree.Size())
}

func TestINTERNALBSTBorrarTodosDeAUno(t *testing.T) {
	bstree := NewBinarySearchTree[int]()
	bstree.Insert(4)
	bstree.Insert(1)
	bstree.Insert(2)
	bstree.Insert(5)
	bstree.Insert(3)

	assert.Equal(t, 5, bstree.Size())

	bstree.Remove(1)
	assert.False(t, bstree.Search(1))
	assert.Equal(t, 4, bstree.Size())

	bstree.Remove(2)
	assert.False(t, bstree.Search(2))
	assert.Equal(t, 3, bstree.Size())

	bstree.Remove(3)
	assert.False(t, bstree.Search(3))
	assert.Equal(t, 2, bstree.Size())

	bstree.Remove(5)
	assert.False(t, bstree.Search(5))
	assert.Equal(t, 1, bstree.Size())

	bstree.Remove(4)
	assert.False(t, bstree.Search(4))
	assert.Equal(t, 0, bstree.Size())
}

func TestINTERNALBSTBorrarInexistente(t *testing.T) {
	bstree := NewBinarySearchTree[int]()

	assert.Equal(t, 0, bstree.Size())

	bstree.Remove(1)

	assert.Equal(t, 0, bstree.Size())
}

func TestINTERNALBSTBorrarUnicoElemento(t *testing.T) {
	bstree := NewBinarySearchTree[int]()
	bstree.Insert(4)

	assert.Equal(t, 1, bstree.Size())

	bstree.Remove(4)

	assert.Equal(t, 0, bstree.Size())
}

func TestINTERNALBSTBorrarRaizConUnHijo(t *testing.T) {
	bstree := NewBinarySearchTree[int]()
	bstree.Insert(4)
	bstree.Insert(3)

	assert.Equal(t, 2, bstree.Size())

	bstree.Remove(4)

	assert.False(t, bstree.Search(4))
	assert.Equal(t, 1, bstree.Size())
}

func TestINTERNALBSTClear(t *testing.T) {
	bstree := NewBinarySearchTree[int]()
	bstree.Insert(4)

	assert.Equal(t, 1, bstree.Size())

	bstree.Clear()

	assert.Equal(t, 0, bstree.Size())
	assert.Nil(t, bstree.GetRoot())
}

func TestINTERNALBSTIsEmpty(t *testing.T) {
	bstree := NewBinarySearchTree[int]()

	assert.True(t, bstree.IsEmpty())

	bstree.Insert(4)

	assert.False(t, bstree.IsEmpty())
}
