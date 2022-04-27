package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppend(t *testing.T) {
	var empty *List[int]
	l := empty.Append(1)
	assert.Equal(t, 1, l.val)

	l = &List[int]{val: 3}

	l.Append(5)
	l.Append(8)

	assert.Equal(t, 3, l.val)
	assert.Equal(t, 5, l.next.val)
	assert.Equal(t, 8, l.next.next.val)
}

func TestLength(t *testing.T) {
	var empty *List[int]
	assert.Equal(t, 0, empty.Length())

	l := List[int]{val: 3}

	assert.Equal(t, 1, l.Length())

	l.Append(5)
	assert.Equal(t, 2, l.Length())

	l.Append(5)
	l.Append(6)
	assert.Equal(t, 4, l.Length())
}

func TestString(t *testing.T) {
	var empty *List[int]
	assert.Equal(t, "[]", empty.String())

	l := List[int]{val: 3}

	assert.Equal(t, "[3]", l.String())

	l.Append(5)
	assert.Equal(t, "[3 5]", l.String())

	l.Append(5)
	l.Append(6)
	assert.Equal(t, "[3 5 5 6]", l.String())
}

func TestAt(t *testing.T) {
	var empty *List[int]
	_, err := empty.At(0)

	assert.Error(t, OutOfBoundsError{}, err)

	l := empty.Append(4)
	v, err := l.At(0)
	assert.Equal(t, 4, v)
	assert.Equal(t, nil, err)

	l.Append(5)
	l.Append(8)
	l.Append(10)

	v, _ = l.At(2)
	assert.Equal(t, 8, v)
}

func TestInsert(t *testing.T) {
	var empty *List[int]
	l, err := empty.Insert(0, 10)
	assert.Equal(t, 10, l.val)
	assert.Equal(t, nil, err)

	_, err = empty.Insert(1, 10)
	assert.Error(t, OutOfBoundsError{}, err)

	l.Insert(1, 20)
	assert.Equal(t, 20, l.next.val)
	l.Insert(1, 30)
	assert.Equal(t, 30, l.next.val)
	assert.Equal(t, 20, l.next.next.val)

	l, _ = l.Insert(0, 5)
	assert.Equal(t, 5, l.val)
	assert.Equal(t, 10, l.next.val)
	assert.Equal(t, 30, l.next.next.val)

	_, err = l.Insert(10, 100)
	assert.Error(t, OutOfBoundsError{}, err)
}
