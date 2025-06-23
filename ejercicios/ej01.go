package ejercicios

import (
	"errors"
	"untref-ayp2/guia-abb/binarytree"
)

func SecondLargestElement(bt *binarytree.BinarySearchTree[int]) (int, error) {
	// Implementar
	if bt.IsEmpty() {
		var x int
		return x, errors.New("No hay valores")
	}
	return findSecLargeElem(bt.GetRoot())
}

func findSecLargeElem(n *binarytree.BinaryNode[int]) (int, error) {
	// Implementar
	if n == nil || n.Size() == 1 {
		return 0, errors.New("arbol con 1 o 0 elementos")
	}

	hijoDer := n.GetRight()
	hijoIzq := n.GetLeft()

	if hijoDer == nil && hijoIzq != nil {
		return hijoIzq.GetData(), nil
	}

	cant, err := findSecLargeElem(hijoDer)

	if err != nil {
		return n.GetData(), nil
	}
	return cant, nil
}
