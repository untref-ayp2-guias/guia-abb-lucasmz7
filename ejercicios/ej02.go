package ejercicios

import (
	"errors"
	"untref-ayp2/guia-abb/binarytree"
)

func PredecesorInOrder(bt *binarytree.BinarySearchTree[int], k int) (int, error) {
	if bt == nil || bt.GetRoot() == nil {
		return 0, errors.New("no hay predecesores")
	}
	return search(bt.GetRoot(), k, -1)
}

func search(n *binarytree.BinaryNode[int], k int, lastRight int) (int, error) {
	if n == nil {
		return 0, errors.New("no hay predecesores menores que el mínimo")
	}

	// Si encontramos el valor k
	if n.GetData() == k {
		// Si tiene subárbol izquierdo, el predecesor es el máximo de ese subárbol
		if n.GetLeft() != nil {
			current := n.GetLeft()
			for current.GetRight() != nil {
				current = current.GetRight()
			}
			return current.GetData(), nil
		}
		// Si no tiene subárbol izquierdo, el predecesor es el último padre del que vinimos por la derecha
		if lastRight != -1 {
			return lastRight, nil
		}
		return 0, errors.New("no hay predecesores menores que el mínimo")
	}

	// Si k es menor que el nodo actual, buscar en el subárbol izquierdo
	if k < n.GetData() {
		return search(n.GetLeft(), k, lastRight)
	}
	// Si k es mayor que el nodo actual, buscar en el subárbol derecho
	// y actualizar lastRight porque estamos yendo a la derecha
	return search(n.GetRight(), k, n.GetData())
}
