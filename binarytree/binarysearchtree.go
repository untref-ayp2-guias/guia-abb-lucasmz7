package binarytree

import (
	"errors"

	"golang.org/x/exp/constraints"
)

type BinarySearchTree[T constraints.Ordered] struct {
	root *BinaryNode[T]
}

// NewBinarySearchTree crea un nuevo BinarySearchTree de tipo Ordered.
//
// Uso:
//
//	bst := tree.NewBinarySearchTree[int]()
//
// Retorna:
//   - un puntero a un nuevo BinarySearchTree.
func NewBinarySearchTree[T constraints.Ordered]() *BinarySearchTree[T] {
	return &BinarySearchTree[T]{root: nil}
}

// GetRoot obtiene el nodo raíz del árbol.
//
// Uso:
//
//	bst := tree.NewBinarySearchTree[int]()
//	bst.GetRoot()
//
// Retorna:
//   - un puntero al nodo raíz del árbol.
func (bst *BinarySearchTree[T]) GetRoot() *BinaryNode[T] {
	return bst.root
}

// Insert inserta un nuevo nodo en el árbol.
//
// Uso:
//
//	bst := tree.NewBinarySearchTree[int]()
//	bst.Insert(4)
//
// Parámetros:
//   - `k` un valor de tipo T.
func (bst *BinarySearchTree[T]) Insert(k T) {
	bst.root = bst.insertByNode(bst.root, k)
}

// insertByNode inserta un nuevo nodo en el árbol.
//
// Parámetros:
//   - `node` un puntero a un BinaryNode.
//   - `k` un valor de tipo T.
//
// Retorna:
//   - un puntero al nodo raíz del árbol.
func (bst *BinarySearchTree[T]) insertByNode(node *BinaryNode[T], k T) *BinaryNode[T] {
	if node == nil {
		return NewBinaryNode(k, nil, nil)
	}
	if k < node.data {
		node.left = bst.insertByNode(node.left, k)
	} else if k > node.data {
		node.right = bst.insertByNode(node.right, k)
	}
	return node
}

// Search busca un nodo en el árbol.
//
// Uso:
//
//	bst := tree.NewBinarySearchTree[int]()
//	bst.Insert(4)
//	bst.Insert(2)
//	bst.Insert(7)
//	node := bst.Search(2)
//
// Parámetros:
//   - `k` un valor de tipo T.
//
// Retorna:
//   - un puntero al nodo encontrado.
func (bst *BinarySearchTree[T]) Search(k T) bool {
	node := bst.root
	for node != nil {
		switch {
		case k < node.data:
			node = node.left
		case k > node.data:
			node = node.right
		default:
			return true
		}
	}
	return false
}

// FindMin busca el nodo con el valor mínimo en el árbol.
//
// Uso:
//
//	bst := tree.NewBinarySearchTree[int]()
//	bst.Insert(4)
//	bst.Insert(2)
//	bst.Insert(7)
//	node := bst.FindMin()
//
// Retorna:
//   - un puntero al nodo con el valor mínimo.
func (bst *BinarySearchTree[T]) FindMin() (T, error) {
	if bst.root == nil {
		var nullElement T

		return nullElement, errors.New("árbol vacío")
	}
	nextLeft := bst.root
	for nextLeft.left != nil {
		nextLeft = nextLeft.left
	}

	return nextLeft.GetData(), nil
}

func (bst *BinarySearchTree[T]) FindMax() (T, error) {
	if bst.root == nil {
		var nullElement T

		return nullElement, errors.New("árbol vacío")
	}
	nextRight := bst.root
	for nextRight.right != nil {
		nextRight = nextRight.right
	}

	return nextRight.GetData(), nil
}

// Remove elimina un nodo del árbol.
//
// Uso:
//
//	bst := tree.NewBinarySearchTree[int]()
//	bst.Insert(4)
//	bst.Insert(2)
//	bst.Insert(7)
//	bst.Remove(2)
//
// Parámetros:
//   - `k` un valor de tipo T.
func (bst *BinarySearchTree[T]) Remove(k T) {
	bst.root = bst.removeByNode(bst.root, k)
}

// removeByNode elimina un nodo del árbol.
//
// Parámetros:
//   - `root` un puntero a un BinaryNode.
//   - `k` un valor de tipo T.
//
// Retorna:
//   - un puntero al nodo raíz del árbol.
func (bst *BinarySearchTree[T]) removeByNode(root *BinaryNode[T], k T) *BinaryNode[T] {
	if root == nil {
		return root
	}
	switch {
	case k > root.data:
		root.right = bst.removeByNode(root.right, k)
	case k < root.data:
		root.left = bst.removeByNode(root.left, k)
	default:
		if root.left == nil {
			return root.right
		}
		if root.right == nil {
			return root.left
		}
		temp := root.left
		for temp.right != nil {
			temp = temp.right
		}
		root.data = temp.data
		root.left = bst.removeByNode(root.left, temp.data)
	}

	return root
}

// Clear limpia el árbol poniendo la raíz en nil.
//
// Uso:
//
//	bst := binarytree.NewBinarySearchTree[int]()
//	bst.Insert(4)
//	bst.Clear()
func (bst *BinarySearchTree[T]) Clear() {
	bst.root = nil
}

// IsEmpty evalúa si el árbol está vacío.
//
// Uso:
//
//	bst := tree.NewBinarySearchTree[int]()
//	bst.IsEmpty()
//
// Retorna:
//   - `true` si el árbol está vacío; `false` si no lo está.
func (bst *BinarySearchTree[T]) IsEmpty() bool {
	return bst.Size() == 0
}

// Size devuelve la cantidad de nodos en el árbol.
//
// Uso:
//
//	bst := tree.NewBinarySearchTree[int]()
//	bst.Insert(4)
//	bst.Insert(2)
//	bst.Insert(7)
//	fmt.Println(bst.Size())
//
// Retorna:
//   - la cantidad de nodos en el árbol.
func (bst *BinarySearchTree[T]) Size() int {
	return size(bst.root)
}

// size devuelve la cantidad de nodos en el árbol local, recursivamente.
//
// Parámetros:
//   - `node` un puntero a un BinaryNode.
//
// Retorna:
//   - la cantidad de nodos en el árbol local.
func size[T constraints.Ordered](node *BinaryNode[T]) int {
	return node.Size()
}
