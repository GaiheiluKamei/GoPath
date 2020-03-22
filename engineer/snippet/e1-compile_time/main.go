package metrics

import "fmt"

// Compile-time checks for ensuring a type implements a particular interface.
var (
	// Works but allocates a dummy Foo instance on the heap.
	_ fmt.Stringer = &Foo{}

	// Preferred way that does not allocate anything on the heap.
	_ fmt.Stringer = (*Foo)(nil)
)

type Foo struct{}

func (*Foo) String() string { return "Foo"}


/*
	The preceding code snippet outlines two fairly common approaches to achieve
this compile-time check by defining a pair of global variables that use the reserved
blank identifier as a hint to the compiler that they are not actually used.
*/