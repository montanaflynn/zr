// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-02-01 10:15:28 2F9E4A                              [zr/int_tuple.go]
// -----------------------------------------------------------------------------

package zr

// # Types
//   IntTuple struct{ A, B int }
//   IntTuples []IntTuple
//
// # sort.Interface
//   (ob IntTuples) Len() int
//   (ob IntTuples) Less(i, j int) bool
//   (ob IntTuples) Swap(i, j int)
//
// # Stringer Interface
//   (ob IntTuple) String() string

import "fmt" // standard

// IntTuple holds a tuple (pair) of integer values.
type IntTuple struct{ A, B int }

// IntTuples holds multiple tuples (pairs) of integer values.
// It implements the sort.Inteface to make sorting easy.
type IntTuples []IntTuple

// -----------------------------------------------------------------------------
// # sort.Interface

// Len is the number of elements in the collection. (sort.Interface)
func (ob IntTuples) Len() int {
	return len(ob)
} //                                                                         Len

// Less reports whether the element with index
// 'i' should sort before element[j].
// (sort.Interface)
func (ob IntTuples) Less(i, j int) bool {
	var ti, tj = ob[i], ob[j]
	return ti.A < tj.A || (ti.A == tj.A && ti.B < tj.B)
} //                                                                        Less

// Swap swaps the elements with indexes i and j. (sort.Interface)
func (ob IntTuples) Swap(i, j int) {
	ob[i], ob[j] = ob[j], ob[i]
} //                                                                        Swap

// -----------------------------------------------------------------------------
// # Stringer Interface

// String returns a string representation of the IntTuple structure
// and implements the Stringer Interface.
func (ob IntTuple) String() string {
	return fmt.Sprintf("IntTuple{A:%d, B:%d}", ob.A, ob.B)
} //                                                                      String

//end
