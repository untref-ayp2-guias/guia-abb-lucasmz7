package ejercicios

import (
	"untref-ayp2/guia-abb/binarytree"
	"untref-ayp2/guia-abb/stack"

	"golang.org/x/exp/constraints"
)

// BSTPostOrderIterator es un iterador para recorrer un BinarySearchTree en postorden.
type BSTPostOrderIterator[T constraints.Ordered] struct {
	internalStack *stack.Stack[*binarytree.BinaryNode[T]]
}

// NewBSTPostOrderIterator crea un nuevo BSTPostOrderIterator.
//
// Parámetros:
//   - `bst` un puntero a un BinarySearchTree.
//
// Retorna:
//   - un Iterator.
func NewBSTPostOrderIterator[T constraints.Ordered](bst *binarytree.BinarySearchTree[T]) Iterator[T] {
	// Implementar
	return nil
}

// HasNext indica si hay un siguiente dato.
func (it *BSTPostOrderIterator[T]) HasNext() bool {
	// Implementar
	return false
}

// Next devuelve el siguiente dato, respetando el recorrido PostOrder.
func (it *BSTPostOrderIterator[T]) Next() (T, error) {
	// Implementar
	var zero T
	return zero, nil
}
