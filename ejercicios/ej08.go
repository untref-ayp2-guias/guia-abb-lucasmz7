package ejercicios

import (
	"untref-ayp2/guia-abb/binarytree"
	"untref-ayp2/guia-abb/stack"

	"golang.org/x/exp/constraints"
)

// BSTLevelOrderIterator es un iterador para recorrer un BinarySearchTree por niveles (BFS).
type BSTLevelOrderIterator[T constraints.Ordered] struct {
	internalStack *stack.Stack[*binarytree.BinaryNode[T]]
}

// NewBSTLevelOrderIterator crea un nuevo BSTLevelOrderIterator.
//
// Parámetros:
//   - `bst` un puntero a un BinarySearchTree.
//
// Retorna:
//   - un Iterator.
func NewBSTLevelOrderIterator[T constraints.Ordered](bst *binarytree.BinarySearchTree[T]) Iterator[T] {
	// Implementar
	return nil
}

// HasNext indica si hay un siguiente dato.
func (it *BSTLevelOrderIterator[T]) HasNext() bool {
	// Implementar
	return false
}

// Next devuelve el siguiente dato, respetando el recorrido por niveles.
func (it *BSTLevelOrderIterator[T]) Next() (T, error) {
	// Implementar
	var zero T
	return zero, nil
}
