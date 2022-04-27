package linkedlist

import "fmt"

// Would the implementation be better with a head pointer?

type List[T any] struct {
	next *List[T]
	val  T
}

type OutOfBoundsError struct {
	index int
	len   int
}

// TODO extract traversal into higher order function

func (e OutOfBoundsError) Error() string {
	return fmt.Sprintf("Index %d out of bounds for list of length %d", e.index, e.len)
}

func (l *List[T]) Append(v T) *List[T] {
	if l == nil {
		return &List[T]{val: v}
	}

	curr := l
	for ; curr.next != nil; curr = curr.next {
	}

	curr.next = &List[T]{val: v}
	return l
}

func (l *List[T]) Length() int {
	if l == nil {
		return 0
	}

	len := 1
	for curr := l; curr.next != nil; curr = curr.next {
		len++
	}

	return len
}

func (l *List[T]) String() string {
	// TODO user a more efficient builder
	s := "["

	for curr := l; curr != nil; curr = curr.next {
		s += fmt.Sprint(curr.val)

		if curr.next != nil {
			s += " "
		}
	}

	s += "]"
	return s
}

func (l *List[T]) Insert(index int, v T) (*List[T], error) {
	new := List[T]{val: v}

	if index == 0 {
		new.next = l
		return &new, nil
	}

	var prev *List[T]
	curr := l

	for i := 1; i <= index; i++ {
		if curr == nil {
			var zero *List[T]
			return zero, OutOfBoundsError{index, i - 1}
		}
		prev = curr
		curr = curr.next
	}

	new.next = curr
	prev.next = &new

	return l, nil
}

func (l *List[T]) At(index int) (T, error) {
	if l == nil {
		var zero T
		return zero, OutOfBoundsError{index, 0}
	}

	v := l.val
	curr := l

	for i := 1; i <= index; i++ {
		if curr.next == nil {
			var zero T
			return zero, OutOfBoundsError{index, i}
		}
		curr = curr.next
		v = curr.val
	}

	return v, nil
}
