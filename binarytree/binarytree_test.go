package binarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestINTERNALGetRoot(t *testing.T) {
	raiz := NewBinaryTree("algo")
	assert.Equal(t, "algo", raiz.GetRoot().GetData())
}

func TestINTERNALSizeHeightEmpty(t *testing.T) {
	mas := NewBinaryTree("+")
	menos := NewBinaryTree("-")
	mult := NewBinaryTree("*")
	a := NewBinaryTree("a")
	b := NewBinaryTree("b")
	c := NewBinaryTree("c")
	d := NewBinaryTree("d")
	menos.InsertLeft(b)
	menos.InsertRight(c)
	mult.InsertLeft(menos)
	mult.InsertRight(d)
	mas.InsertLeft(a)
	mas.InsertRight(mult)

	// Size del arbol(cantidad de nodos)
	assert.Equal(t, 7, mas.Size())

	// Height del arbol
	assert.Equal(t, 3, mas.Height())

	// isEmpty Test
	assert.False(t, mas.IsEmpty())

	// Vaciar el arbol
	mas.Clear()
	assert.True(t, mas.IsEmpty())
}

func TestINTERNALSizeHeightEmptyOnEmptyTree(t *testing.T) {
	tree := NewBinaryTree("+")
	tree.Clear()

	// Size del arbol(cantidad de nodos)
	assert.Equal(t, 0, tree.Size())

	// Height del arbol
	assert.Equal(t, -1, tree.Height())

	// isEmpty Test
	assert.True(t, tree.IsEmpty())
}

func TestINTERNALInertsOnEmptyTree(t *testing.T) {
	tree := NewBinaryTree("+")

	uno := NewBinaryTree("1")
	dos := NewBinaryTree("2")

	tree.Clear()
	tree.InsertLeft(uno)
	assert.False(t, tree.IsEmpty())

	tree.Clear()
	tree.InsertRight(dos)
	assert.False(t, tree.IsEmpty())
}
