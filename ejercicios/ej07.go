package ejercicios

import (
	"errors"
	"untref-ayp2/guia-abb/binarytree"
	"untref-ayp2/guia-abb/stack"

	"golang.org/x/exp/constraints"
)

// BSTPostOrderIterator es un iterador para recorrer un BinarySearchTree en postorden.
type BSTPostOrderIterator[T constraints.Ordered] struct {
	internalStack *stack.Stack[*binarytree.BinaryNode[T]]
}

// NewBSTPostOrderIterator crea un nuevo BSTPostOrderIterator.
//
// Parámetros:
//   - `bst` un puntero a un BinarySearchTree.
//
// Retorna:
//   - un Iterator.
func NewBSTPostOrderIterator[T constraints.Ordered](bst *binarytree.BinarySearchTree[T]) Iterator[T] {
	// Implementar
	stack := stack.NewStack[*binarytree.BinaryNode[T]]()
	if bst.IsEmpty() {
		return &BSTPostOrderIterator[T]{internalStack: stack}
	}
	postorden(bst.GetRoot(), stack)

	return &BSTPostOrderIterator[T]{internalStack: stack}
}

func postorden[T constraints.Ordered](n *binarytree.BinaryNode[T], stack *stack.Stack[*binarytree.BinaryNode[T]]) {
	if n == nil {
		return
	}
	// postOrden: izq -> der -> raiz
	//Para un stack LIFO: raiz -> der -> izq
	stack.Push(n)
	postorden(n.GetRight(), stack)
	postorden(n.GetLeft(), stack)
}

// HasNext indica si hay un siguiente dato.
func (it *BSTPostOrderIterator[T]) HasNext() bool {
	// Implementar
	return !it.internalStack.IsEmpty()
}

// Next devuelve el siguiente dato, respetando el recorrido PostOrder.
func (it *BSTPostOrderIterator[T]) Next() (T, error) {
	// Implementar
	var zero T
	if !it.HasNext() {
		return zero, errors.New("árbol vacío")
	}
	nodo, _ := it.internalStack.Pop()
	return nodo.GetData(), nil
}
