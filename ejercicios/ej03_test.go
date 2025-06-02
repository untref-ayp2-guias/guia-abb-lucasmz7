package ejercicios

import (
	"testing"

	"untref-ayp2/guia-abb/binarytree"

	"github.com/stretchr/testify/assert"
)

//	   4
//	  / \
//	 2   5
//	/ \   \
// 1   3   6
//        / \
//       8   7

func TestNoIsBst(t *testing.T) {
	t1 := binarytree.NewBinaryTree(1)
	t3 := binarytree.NewBinaryTree(3)
	t2 := binarytree.NewBinaryTree(2)
	t5 := binarytree.NewBinaryTree(5)
	t4 := binarytree.NewBinaryTree(4)
	t6 := binarytree.NewBinaryTree(6)
	t7 := binarytree.NewBinaryTree(7)
	t8 := binarytree.NewBinaryTree(8)
	t2.InsertLeft(t1)
	t2.InsertRight(t3)
	t4.InsertLeft(t2)
	t4.InsertRight(t5)
	t6.InsertRight(t7)
	t6.InsertLeft(t8)
	t5.InsertRight(t6)
	tree := t4
	assert.False(t, IsBst(tree))
}

//	   4
//	  / \
//	 2   15
//	/ \    \
// 1   3    17
//           \
//           18

func TestSiIsBst(t *testing.T) {
	t1 := binarytree.NewBinaryTree(1)
	t3 := binarytree.NewBinaryTree(3)
	t2 := binarytree.NewBinaryTree(2)
	t15 := binarytree.NewBinaryTree(15)
	t4 := binarytree.NewBinaryTree(4)
	t17 := binarytree.NewBinaryTree(17)
	t18 := binarytree.NewBinaryTree(18)
	t2.InsertLeft(t1)
	t2.InsertRight(t3)
	t4.InsertLeft(t2)
	t4.InsertRight(t15)
	t15.InsertRight(t17)
	t17.InsertRight(t18)
	tree := t4
	assert.True(t, IsBst(tree))
}

func TestNilIsBst(t *testing.T) {
	t1 := binarytree.NewBinaryTree(1)
	t1.Clear()
	tree := t1
	assert.True(t, IsBst(tree))
}

func TestNodoIsBst(t *testing.T) {
	t1 := binarytree.NewBinaryTree(1)
	tree := t1
	assert.True(t, IsBst(tree))
}

//	   4
//	  / \
//	 2   15
//	/ \    \
// 14  3    17
//           \
//           18

func TestMalMinIsBst(t *testing.T) {
	t1 := binarytree.NewBinaryTree(14)
	t3 := binarytree.NewBinaryTree(3)
	t2 := binarytree.NewBinaryTree(2)
	t15 := binarytree.NewBinaryTree(15)
	t4 := binarytree.NewBinaryTree(4)
	t17 := binarytree.NewBinaryTree(17)
	t18 := binarytree.NewBinaryTree(18)
	t2.InsertLeft(t1)
	t2.InsertRight(t3)
	t4.InsertLeft(t2)
	t4.InsertRight(t15)
	t15.InsertRight(t17)
	t17.InsertRight(t18)
	tree := t4
	assert.False(t, IsBst(tree))
}

//	   4
//	  / \
//	 2   15
//	/ \    \
// 1   3    17
//           \
//           -18

func TestMalMaxSiIsBst(t *testing.T) {
	t1 := binarytree.NewBinaryTree(1)
	t3 := binarytree.NewBinaryTree(3)
	t2 := binarytree.NewBinaryTree(2)
	t15 := binarytree.NewBinaryTree(15)
	t4 := binarytree.NewBinaryTree(4)
	t17 := binarytree.NewBinaryTree(17)
	t18 := binarytree.NewBinaryTree(-18)
	t2.InsertLeft(t1)
	t2.InsertRight(t3)
	t4.InsertLeft(t2)
	t4.InsertRight(t15)
	t15.InsertRight(t17)
	t17.InsertRight(t18)
	tree := t4
	assert.False(t, IsBst(tree))
}
