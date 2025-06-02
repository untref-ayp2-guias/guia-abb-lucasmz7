package ejercicios

import (
	"errors"
	"untref-ayp2/guia-abb/binarytree"
)

func SecondLargestElement(bt *binarytree.BinarySearchTree[int]) (int, error) {
	n, err := bt.FindMax()
	if err != nil {
		erro := errors.New("No hay valores")
		return 0, erro
	}

	bt.Remove(n)

	x, err2 := bt.FindMax()

	if err2 != nil {
		erro := errors.New("No hay valores")
		return 0, erro
	}

	bt.Insert(n)

	return x, nil
}

/*
func SecondLargestElement2(bt *binarytree.BinarySearchTree[int]) (int, error) {
	if bt == nil || bt.GetRoot() == nil {
		erro := errors.New("No hay valores")
		return 0, erro
	}
	x, _ := bt.FindMin()
	return findSecLargeElem(x, x, bt.GetRoot())
}

func findSecLargeElem(max1, max2 int, n *binarytree.BinaryNode[int]) (int, error) {
	if n == nil {
		err := errors.New("error")
		return 0, err
	}

	if n.GetData() > max1 {
		max1 = n.GetData()
	}
	if n.GetData() > max2 && n.GetData() <= max1 {
		max2 = n.GetData()
	}
	findSecLargeElem(max1, max2, n.GetRight())
	findSecLargeElem(max1, max2, n.GetLeft())

	return max2, nil
}
*/
