package binarytree

import (
	"golang.org/x/exp/constraints"
)

// BinaryNode implementa el tipo BinaryNode con dos campos de tipo
// *BinaryNode que son punteros a los hijos izquierdo y derecho, tambien de tipo BinaryNode,
// y un tercer campo de tipo T generico pero Ordered, por compatibilidad con BinarySerchTree.
type BinaryNode[T constraints.Ordered] struct {
	left  *BinaryNode[T]
	right *BinaryNode[T]
	data  T
}

// NewBinaryNode crea un nuevo BinaryNode.
//
// Uso:
//
//	d := binarytree.NewBinarNode[int](data, hIzq, hDer)
//
// Parámetros:
//   - 'data' : el dato que guarda el nodo de tipo T
//   - 'left' : el nodo que será asignado como hijo izquierdo
//   - 'right' : el nodo que será asignado como hijo derecho
//
// Retorna:
//   - un puntero a un nuevo BinaryNode.
func NewBinaryNode[T constraints.Ordered](data T, left *BinaryNode[T], right *BinaryNode[T]) *BinaryNode[T] {
	return &BinaryNode[T]{left: left, right: right, data: data}
}

// GetData retorna el dato guardado en el nodo de tipo T.
//
// Retorna:
//   - el dato guardado en el nodo de tipo T.
func (n *BinaryNode[T]) GetData() T {
	return n.data
}

// GetLeft retorna el hijo izquierdo del nodo.
//
// Retorna:
//   - un puntero al hijo izquierdo del nodo.
func (n *BinaryNode[T]) GetLeft() *BinaryNode[T] {
	return n.left
}

// GetRight retorna el hijo derecho del nodo.
//
// Retorna:
//   - un puntero al hijo derecho del nodo.
func (n *BinaryNode[T]) GetRight() *BinaryNode[T] {
	return n.right
}

// Size evalúa el tamaño del árbol.
//
// Retorna:
//   - la cantidad de nodos del árbol
func (n *BinaryNode[T]) Size() int {
	if n == nil {
		return 0
	}
	return 1 + n.left.Size() + n.right.Size()
}

// Height retorna la altura del árbol.
//
// Retorna:
//   - la altura del árbol.
func (n *BinaryNode[T]) Height() int {
	if n == nil {
		return -1
	}
	leftHeight := -1
	rightHeight := -1
	if n.left != nil {
		leftHeight = n.left.Height()
	}
	if n.right != nil {
		rightHeight = n.right.Height()
	}
	return 1 + max(leftHeight, rightHeight)
}
