package ejercicios

import (
	"fmt"
	"untref-ayp2/guia-abb/binarytree"

	"golang.org/x/exp/constraints"
)

type TreeSet[T constraints.Ordered] struct {
	set *binarytree.BinarySearchTree[T]
}

func NewTreeSet[T constraints.Ordered](elements ...T) *TreeSet[T] {
	// Implementar
	abb := binarytree.NewBinarySearchTree[T]()
	for _, actual := range elements {
		abb.Insert(actual)
	}
	return &TreeSet[T]{set: abb}
}

func (ts *TreeSet[T]) Add(elements ...T) {
	// Implementar
	for _, actual := range elements {
		ts.set.Insert(actual)
	}
}

func (ts *TreeSet[T]) Size() int {
	// Implementar
	return ts.set.Size()
}

func (ts *TreeSet[T]) Contains(element T) bool {
	// Implementar
	return ts.set.Search(element)
}

func (ts *TreeSet[T]) Remove(element T) {
	// Implementar
	ts.set.Remove(element)
}

func (ts *TreeSet[T]) Values() []T {
	// Implementar
	var val []T
	if ts.Size() == 0 {
		return val
	}
	root := ts.set.GetRoot()

	cola := []*binarytree.BinaryNode[T]{root}

	for len(cola) > 0 {

		actual := cola[0]
		cola = cola[1:]

		val = append(val, actual.GetData())

		hijIzq := actual.GetLeft()
		hijoDer := actual.GetRight()

		if hijIzq != nil {
			cola = append(cola, hijIzq)
		}

		if hijoDer != nil {
			cola = append(cola, hijoDer)
		}
	}

	return val
}

func (ts *TreeSet[T]) String() string {
	// Implementar
	sting := "Set: {"
	if ts.Size() == 0 {
		return sting + "}"
	}
	root := ts.set.GetRoot()

	cola := []*binarytree.BinaryNode[T]{root}

	for len(cola) > 0 {
		actual := cola[0]
		cola = cola[1:]

		der := actual.GetRight()
		izq := actual.GetLeft()

		if izq != nil {
			cola = append(cola, izq)
		}
		if der != nil {
			cola = append(cola, der)
		}

		sting += fmt.Sprintf("%v", actual.GetData())
		if len(cola) != 0 {
			sting += " "
		}
	}
	return sting + "}"
}
