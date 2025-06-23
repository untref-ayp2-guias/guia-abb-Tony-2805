package ejercicios

import (
	"untref-ayp2/guia-abb/binarytree"
)

func IsBst(bt *binarytree.BinaryTree[int]) bool {
	// Implementar
	if bt.IsEmpty() {
		return true
	}
	//return isBstByNode(bt.GetRoot())
	return IsBstByNode2(bt.GetRoot(), nil, nil)
}

func isBstByNode(n *binarytree.BinaryNode[int]) bool {
	// Implementar
	if n.Size() == 1 || n == nil {
		return true
	}

	isBst := true
	der, izq := n.GetRight(), n.GetLeft()

	if der != nil && der.GetData() < n.GetData() {
		isBst = false
	}

	if izq != nil && izq.GetData() > n.GetData() {
		isBst = false
	}

	return isBst && isBstByNode(der) && isBstByNode(izq)

}

// isBstByNode valida que todos los nodos respeten los limites [min, max].
func IsBstByNode2(n *binarytree.BinaryNode[int], min *int, max *int) bool {
	if n == nil {
		return true
	}

	val := n.GetData()

	if min != nil && val < *min {
		return false
	}
	if max != nil && val > *max {
		return false
	}

	return IsBstByNode2(n.GetLeft(), min, &val) && IsBstByNode2(n.GetRight(), &val, max)
}
