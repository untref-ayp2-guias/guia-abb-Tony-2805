package ejercicios

import (
	"errors"
	"untref-ayp2/guia-abb/binarytree"
	"untref-ayp2/guia-abb/stack"

	"golang.org/x/exp/constraints"
)

// BSTInOrderIterator es un iterador para recorrer un BinarySearchTree.
type BSTInOrderIterator[T constraints.Ordered] struct {
	internalStack *stack.Stack[*binarytree.BinaryNode[T]]
}

// NewBSTInOrderIterator crea un nuevo BSTInOrderIterator.
//
// Parámetros:
//   - `bst` un puntero a un BinarySearchTree.
//
// Retorna:
//   - un Iterator.
func NewBSTInOrderIterator[T constraints.Ordered](bst *binarytree.BinarySearchTree[T]) Iterator[T] {
	// Implementar
	stackNodos := stack.NewStack[*binarytree.BinaryNode[T]]()
	if bst.IsEmpty() {
		return &BSTInOrderIterator[T]{internalStack: stackNodos}
	}
	recurInorder(bst.GetRoot(), stackNodos)
	return &BSTInOrderIterator[T]{internalStack: stackNodos}
}
func recurInorder[T constraints.Ordered](n *binarytree.BinaryNode[T], stack *stack.Stack[*binarytree.BinaryNode[T]]) {
	if n == nil {
		return
	}
	//al ser un stack Last In First Out, hay que cargar al en el sentido contrario los datos
	// no izq -> r -> der
	// der -> r -> izq
	recurInorder(n.GetRight(), stack)
	stack.Push(n)
	recurInorder(n.GetLeft(), stack)
}

// HasNext indica si hay un siguiente dato.
func (it *BSTInOrderIterator[T]) HasNext() bool {
	// Implementar
	return !it.internalStack.IsEmpty()
}

// Next devuelve el siguiente dato, respetando el recorrido InOrder.
func (it *BSTInOrderIterator[T]) Next() (T, error) {
	// Implementar
	var zero T
	if !it.HasNext() {
		return zero, errors.New("árbol vacío")
	}
	top, _ := it.internalStack.Pop()

	return top.GetData(), nil
}
