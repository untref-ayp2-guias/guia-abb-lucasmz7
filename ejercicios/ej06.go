package ejercicios

import (
	"untref-ayp2/guia-abb/binarytree"
	"untref-ayp2/guia-abb/stack"

	"golang.org/x/exp/constraints"
)

// BSTPreOrderIterator es un iterador para recorrer un BinarySearchTree en preorden.
type BSTPreOrderIterator[T constraints.Ordered] struct {
	internalStack *stack.Stack[*binarytree.BinaryNode[T]]
}

// NewBSTPreOrderIterator crea un nuevo BSTPreOrderIterator.
//
// Parámetros:
//   - `bst` un puntero a un BinarySearchTree.
//
// Retorna:
//   - un Iterator.
func NewBSTPreOrderIterator[T constraints.Ordered](bst *binarytree.BinarySearchTree[T]) Iterator[T] {
	// Implementar
	return nil
}

// HasNext indica si hay un siguiente dato.
func (it *BSTPreOrderIterator[T]) HasNext() bool {
	// Implementar
	return false
}

// Next devuelve el siguiente dato, respetando el recorrido PreOrder.
func (it *BSTPreOrderIterator[T]) Next() (T, error) {
	// Implementar
	var zero T
	return zero, nil
}
