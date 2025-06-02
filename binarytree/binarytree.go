// Package binarytree provee una implementación de árboles.
package binarytree

import (
	"golang.org/x/exp/constraints"
)

type BinaryTree[T constraints.Ordered] struct {
	root *BinaryNode[T]
}

// NewBinaryTree crea un nuevo BinaryTree a partir de un dato de tipo T.
//
// Uso:
//
//	bt := binarytree.NewBinaryTree[int](data)
//
// Parámetros:
//   - 'data' : el dato de tipo T que guarda el nodo raíz
//
// Retorna:
//   - un puntero a un nuevo BinaryTree.
func NewBinaryTree[T constraints.Ordered](data T) *BinaryTree[T] {
	node := NewBinaryNode(data, nil, nil)

	return &BinaryTree[T]{root: node}
}

// GetRoot obtiene el nodo raíz del árbol.
//
// Uso:
//
//	bt := binarytree.NewBinaryTree[int](data)
//	bt.GetRoot()
//
// Retorna:
//   - un puntero al nodo raíz del árbol.
func (t *BinaryTree[T]) GetRoot() *BinaryNode[T] {
	return t.root
}

// InsertLeft inserta del lado izquierdo de la raíz, el árbol que se pasa por parámetro
//
// Uso:
//
//	bt := binarytree.NewBinaryTree[int](data)
//	bt2 := binarytree.NewBinaryTree[int](data2)
//	bt.InsertLeft(bt2)
//
// Parámetros:
//   - `bt` un puntero a un BinaryTree.
func (t *BinaryTree[T]) InsertLeft(bt *BinaryTree[T]) {
	if t.root == nil {
		t.root = bt.root
	} else {
		t.root.left = bt.root
	}
}

// InsertRight inserta del lado derecho de la raíz, el árbol que se pasa por parámetro
//
// Uso:
//
//	bt := binarytree.NewBinaryTree[int](data)
//	bt2 := binarytree.NewBinaryTree[int](data2)
//	bt.InsertRight(bt2)
//
// Parámetros:
//   - `bt` un puntero a un BinaryTree.
func (t *BinaryTree[T]) InsertRight(bt *BinaryTree[T]) {
	if t.root == nil {
		t.root = bt.root
	} else {
		t.root.right = bt.root
	}
}

// Clear limpia el árbol poniendo la raíz en nil.
//
// Uso:
//
//	bt := binarytree.NewBinaryTree[int](data)
//	bt.Clear()
func (t *BinaryTree[T]) Clear() {
	t.root = nil
}

// IsEmpty evalúa si el árbol está vacío.
//
// Uso:
//
//	bt := binarytree.NewBinaryTree[int](data)
//	bt.IsEmpty()
//
// Retorna:
//   - `true` si el árbol está vacío; `false` si no lo está.
func (t *BinaryTree[T]) IsEmpty() bool {
	return t.root == nil
}

// Size retorna la cantidad de nodos del árbol.
//
// Uso:
//
//	bt := binarytree.NewBinaryTree[int](data)
//	bt.Size()
//
// Retorna:
//   - la cantidad de nodos del árbol.
func (t *BinaryTree[T]) Size() int {
	return t.root.Size()
}

// Height etorna la altura del árbol, o sea, la distancia desde la raíz al nodo más profundo.
//
// Uso:
//
//	bt := binarytree.NewBinaryTree[int](data)
//	bt.Height()
//
// Retorna:
//   - la altura del árbol.
func (t *BinaryTree[T]) Height() int {
	return t.root.Height()
}
