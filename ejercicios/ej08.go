package ejercicios

import (
	"errors"
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
	stack := stack.NewStack[*binarytree.BinaryNode[T]]()
	if bst.IsEmpty() {
		return &BSTLevelOrderIterator[T]{internalStack: stack}
	}

	var queue []*binarytree.BinaryNode[T]
	queue = append(queue, bst.GetRoot())

	var levelOrder []*binarytree.BinaryNode[T]

	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		levelOrder = append(levelOrder, n)

		if n.GetLeft() != nil {
			queue = append(queue, n.GetLeft())
		}
		if n.GetRight() != nil {
			queue = append(queue, n.GetRight())
		}
	}

	// Invertimos el orden para usarlo como pila (LIFO)
	for i := len(levelOrder) - 1; i >= 0; i-- {
		stack.Push(levelOrder[i])
	}

	return &BSTLevelOrderIterator[T]{internalStack: stack}
}

// HasNext indica si hay un siguiente dato.
func (it *BSTLevelOrderIterator[T]) HasNext() bool {
	// Implementar
	return !it.internalStack.IsEmpty()
}

// Next devuelve el siguiente dato, respetando el recorrido por niveles.
func (it *BSTLevelOrderIterator[T]) Next() (T, error) {
	// Implementar
	var zero T
	if !it.HasNext() {
		return zero, errors.New("árbol vacío")
	}
	nodo, _ := it.internalStack.Pop()
	return nodo.GetData(), nil
}
