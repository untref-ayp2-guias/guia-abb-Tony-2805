package ejercicios

import (
	"errors"
	"untref-ayp2/guia-abb/binarytree"
)

func PredecesorInOrder(bt *binarytree.BinarySearchTree[int], k int) (int, error) {
	// Implementar
	if bt.IsEmpty() {
		return 0, errors.New("No hay predecesores")
	}
	current := bt.GetRoot()
	var posPredecesor *binarytree.BinaryNode[int]

	for current != nil {
		if k > current.GetData() {
			posPredecesor = current
			current = current.GetRight()
		} else if k < current.GetData() {
			current = current.GetLeft()

			// Si se encuentra un nodo con el valor, el predecesir es el maximo
		} else {
			if current.GetLeft() != nil {
				return max(current.GetLeft()).GetData(), nil
			}
			if posPredecesor != nil {
				return posPredecesor.GetData(), nil
			}
			return 0, errors.New("No hay predecesores menores que el mínimo")
		}
	}

	if posPredecesor != nil {
		return posPredecesor.GetData(), nil
	}

	return 0, errors.New("No hay predecesores menores que el mínimo")
}

func max(node *binarytree.BinaryNode[int]) *binarytree.BinaryNode[int] {
	for node.GetRight() != nil {
		node = node.GetRight()
	}
	return node
}
