package ejercicios

import (
	"errors"
	"untref-ayp2/guia-abb/binarytree"
	"untref-ayp2/guia-abb/stack"

	"golang.org/x/exp/constraints"
)

// BSTPreOrderIterator es un iterador para recorrer un BinarySearchTree en preorden.
type BSTPreOrderIterator[T constraints.Ordered] struct {
	internalStack *stack.Stack[*binarytree.BinaryNode[T]]
}

// NewBSTPreOrderIterator crea un nuevo BSTPreOrderIterator.
//
// Parámetros:
//   - `bst` un puntero a un BinarySearchTree.
//
// Retorna:
//   - un Iterator.
func NewBSTPreOrderIterator[T constraints.Ordered](bst *binarytree.BinarySearchTree[T]) Iterator[T] {
	// Implementar
	stackNodos := stack.NewStack[*binarytree.BinaryNode[T]]()
	if bst.IsEmpty() {
		return &BSTInOrderIterator[T]{internalStack: stackNodos}
	}
	recurPreorder(bst.GetRoot(), stackNodos)
	return &BSTInOrderIterator[T]{internalStack: stackNodos}
}

func recurPreorder[T constraints.Ordered](n *binarytree.BinaryNode[T], stack *stack.Stack[*binarytree.BinaryNode[T]]) {
	if n == nil {
		return
	}
	// Preorder: r i d
	// para un stack d i r

	recurPreorder(n.GetRight(), stack)
	recurPreorder(n.GetLeft(), stack)
	stack.Push(n)
}

// HasNext indica si hay un siguiente dato.
func (it *BSTPreOrderIterator[T]) HasNext() bool {
	// Implementar
	return !it.internalStack.IsEmpty()
}

// Next devuelve el siguiente dato, respetando el recorrido PreOrder.
func (it *BSTPreOrderIterator[T]) Next() (T, error) {
	// Implementar
	var zero T
	if !it.HasNext() {
		return zero, errors.New("árbol vacío")
	}
	top, _ := it.internalStack.Pop()

	return top.GetData(), nil
}
