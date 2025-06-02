package ejercicios

import (
	"untref-ayp2/guia-abb/binarytree"

	"golang.org/x/exp/constraints"
)

type TreeSet[T constraints.Ordered] struct {
	set *binarytree.BinarySearchTree[T]
}

func NewTreeSet[T constraints.Ordered](elements ...T) *TreeSet[T] {
	// Implementar
	return nil
}

func (ts *TreeSet[T]) Add(elements ...T) {
	// Implementar
}

func (ts *TreeSet[T]) Size() int {
	// Implementar
	return 0
}

func (ts *TreeSet[T]) Contains(element T) bool {
	// Implementar
	return false
}

func (ts *TreeSet[T]) Remove(element T) {
	// Implementar
}

func (ts *TreeSet[T]) Values() []T {
	// Implementar
	return nil
}

func (ts *TreeSet[T]) String() string {
	// Implementar
	return ""
}
