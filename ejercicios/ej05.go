package ejercicios

import (
	"untref-ayp2/guia-abb/binarytree"
	"untref-ayp2/guia-abb/stack"

	"golang.org/x/exp/constraints"
)

// BSTInOrderIterator es un iterador para recorrer un BinarySearchTree.
type BSTInOrderIterator[T constraints.Ordered] struct {
	internalStack *stack.Stack[*binarytree.BinaryNode[T]]
}

// NewBSTInOrderIterator crea un nuevo BSTInOrderIterator.
//
// Parámetros:
//   - `bst` un puntero a un BinarySearchTree.
//
// Retorna:
//   - un Iterator.
func NewBSTInOrderIterator[T constraints.Ordered](bst *binarytree.BinarySearchTree[T]) Iterator[T] {
	// Implementar
	return nil
}

// HasNext indica si hay un siguiente dato.
func (it *BSTInOrderIterator[T]) HasNext() bool {
	// Implementar
	return false
}

// Next devuelve el siguiente dato, respetando el recorrido InOrder.
func (it *BSTInOrderIterator[T]) Next() (T, error) {
	// Implementar
	var zero T
	return zero, nil
}
