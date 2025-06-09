package ejercicios

import (
	"errors"
	"untref-ayp2/guia-abb/binarytree"

	"golang.org/x/exp/constraints"
)

// BSTLevelOrderIterator es un iterador para recorrer un BinarySearchTree por niveles (BFS).
type BSTLevelOrderIterator[T constraints.Ordered] struct {
	internalQueue *Queue[*binarytree.BinaryNode[T]]
}

// NewBSTLevelOrderIterator crea un nuevo BSTLevelOrderIterator.
//
// Parámetros:
//   - `bst` un puntero a un BinarySearchTree.
//
// Retorna:
//   - un Iterator.
func NewBSTLevelOrderIterator[T constraints.Ordered](bst *binarytree.BinarySearchTree[T]) Iterator[T] {
	it := &BSTLevelOrderIterator[T]{
		internalQueue: NewQueue[*binarytree.BinaryNode[T]](),
	}

	if root := bst.GetRoot(); root != nil {
		it.internalQueue.Enqueue(root)
	}

	return it
}

// HasNext indica si hay un siguiente dato.
func (it *BSTLevelOrderIterator[T]) HasNext() bool {
	return !it.internalQueue.IsEmpty()
}

// Next devuelve el siguiente dato, respetando el recorrido por niveles.
func (it *BSTLevelOrderIterator[T]) Next() (T, error) {
	if !it.HasNext() {
		var zero T
		err := errors.New("árbol vacío")
		return zero, err
	}

	node, _ := it.internalQueue.Dequeue()

	if left := node.GetLeft(); left != nil {
		it.internalQueue.Enqueue(left)
	}
	if right := node.GetRight(); right != nil {
		it.internalQueue.Enqueue(right)
	}

	return node.GetData(), nil
}
