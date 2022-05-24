package container

// Slice is a container type for type T. The main motivation of this type is the
// abstraction of the pop, push, shift, and unshift operations.
type Slice[T any] []T

// Pop removes the last element from s and returns that element.
// Pop returns the zero value of T if s is empty.
func (s *Slice[T]) Pop() (tn T) {
	l := len(*s)
	if l != 0 {
		tn, *s = (*s)[l-1], (*s)[:l-1]
	}
	return
}

// Push adds one or more elements to the end of s.
// Push returns the length of s.
func (s *Slice[T]) Push(tt ...T) int {
	*s = append(*s, tt...)
	return len(*s)
}

// Shift removes the first element from s and returns that removed element.
// Shift returns the zero element of T if s is empty.
func (s *Slice[T]) Shift() (tn T) {
	if len(*s) != 0 {
		tn, *s = (*s)[0], (*s)[1:]
	}
	return
}

// Unshift adds one or more elements to the beginning of s.
// Unshift returns the length of s.
func (s *Slice[T]) Unshift(tt ...T) int {
	*s = append(tt, *s...)
	return len(*s)
}
