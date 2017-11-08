/*
Package matf64 implements a large set of functions which act on two dimensional
slices of float64.

Many of the functions in this library expect either a float64, a []float64,
or [][]float64, and do "the right thing" based on what is passed. For example,
consider the function

	matf64.Mul(m, n)

In this function, m is a [][]float64, where as n could be
a float64, []float64, or [][]float64. This allows the same function to be called
for wide range of situations.

all the functions in this library act on Go primitive types, which allows the code
to be easily modified to serve in different situations, and to easily integrate with
existing code bases

*/
package matf64

import (
	"fmt"
	"math/rand"
)

/*
New is a utility function to create [][]float64s. New is a variadic function,
expecting 0, 1 or 2 ints, with differing behavior as follows:

	m := matf64.New(x)

will return a x by x (square) [][]float64. Alternatively

	m := matf64.New(x, y)

is a [][]float64 with x rows and y columns.

*/
func New(dims ...int) [][]float64 {
	var m [][]float64
	switch len(dims) {
	case 1:
		r := dims[0]
		m = make([][]float64, r)
		for i := range m {
			m[i] = make([]float64, r)
		}
	case 2:
		r := dims[0]
		c := dims[1]
		m = make([][]float64, r)
		for i := range m {
			m[i] = make([]float64, c)
		}
	default:
		s := "In matf64.%s expected 1 or 2 arguments, but recieved %d"
		s = fmt.Sprintf(s, "New()", len(dims))
		panic(s)
	}
	return m
}

/*
I returns a square [][]float64 with all elements along the diagonal equal to
1.0, and 0.0 elsewhere. This is the identity matrix.
*/
func I(x int) [][]float64 {
	m := New(x)
	for i := range m {
		m[i][i] = 1.0
	}
	return m
}

/*
RandMat creates a x by y [][]float64 with the entries set to random numbers in the
range [0, 1) (including 0, but excluding 1).

A n by n [][]float64 can be created if one int is passed to this constructor, where
a n by m matrix is created when two ints are passed.
*/
func RandMat(x, y int, args ...float64) [][]float64 {
	m := New(y, y)
	var from float64
	var to float64

	switch len(args) {
	case 0:
		to = 1
	case 1:
		to = args[0]
	case 2:
		from = args[0]
		to = args[1]
	default:
		s := "In matf64.%s expected 0-2 float64s for the range, but recieved %d"
		s = fmt.Sprintf(s, "RandMat()", len(args))
		panic(s)
	}
	for i := range m {
		for j := range m[i] {
			m[i][j] = rand.Float64()*(to-from) + from
		}
	}
	return m
}

/*
RandVec returns a []float64 with the entries set to random number in the
range [0, 1)
*/
func RandVec(size int, args ...float64) []float64 {
	v := make([]float64, size)
	var from float64
	var to float64

	switch len(args) {
	case 0:
		to = 1
	case 1:
		to = args[0]
	case 2:
		from = args[0]
		to = args[1]
	default:
		s := "In matf64.%s expected 0-2 float64s for the range, but recieved %d"
		s = fmt.Sprintf(s, "RandVec()", len(args))
		panic(s)
	}
	for i := range v {
		v[i] = rand.Float64()*(to-from) + from
	}
	return v
}

/*
Flatten constructs a []float64 from a [][]float64 where each row's head is
appended to the previous row's tail. For example:

[[1, 2, 3]
 [4, 5, 6]
 [7, 8, 9]]

becomes

[1 2 3 4 5 6 7 8 9]

*/
func Flatten(m [][]float64) []float64 {
	var n []float64
	for i := range m {
		n = append(n, m[i]...)
	}
	return n
}

/*
Col returns a column from a [][]float64. For example:

	fmt.Println(m) // [[1.0, 2.3], [3.4, 1.7]]
	matf64.Col(m, 0) // [1.0, 3.4]

Col also accepts negative indices. For example:

	matf64.Col(m, -1) // [2.3, 1.7]

The original [][]float64 is not mutated in this function.
*/
func Col(m [][]float64, x int) []float64 {
	v := make([]float64, len(m))
	if x >= 0 {
		for i := range m {
			v[i] = m[i][x]
		}
	} else {
		for i := range m {
			v[i] = m[i][len(m[0])+x]
		}
	}
	return v
}

/*
Row returns a row from a [][]float64. For example:

	fmt.Println(m) // [[1.0, 2.3], [3.4, 1.7]]
	matf64.Row(m, 0) // [1.0, 2.3]

Row also accepts negative indices. For example:

	matf64.Row(m, -1) // [3.4, 1.7]

The original [][]float64 is not mutated in this function.
*/
func Row(m [][]float64, x int) []float64 {
	v := make([]float64, len(m[0]))
	if x >= 0 {
		copy(v, m[x])
	} else {
		copy(v, m[len(m)+x])
	}
	return v
}

/*
Equal checks to see if two [][]float64s are equal. That mean that the two slices
have the same number of rows, same number of columns, and have the same float64
in each entry at a given set of indices.
*/
func Equal(m, n [][]float64) bool {
	if len(m) != len(n) {
		return false
	}
	for i := range m {
		if len(m[i]) != len(n[i]) {
			return false
		}
	}
	for i := range m {
		for j := range m[i] {
			if m[i][j] != n[i][j] {
				return false
			}
		}
	}
	return true
}

/*
Clone returns a duplicate of a [][]float64. The returned duplicate is "deep",
meaning that the object can be manipulated without effecting the original.
*/
func Clone(m [][]float64) [][]float64 {
	n := make([][]float64, len(m))
	for i := range m {
		n[i] = make([]float64, len(m[i]))
		for j := range m[i] {
			n[i][j] = m[i][j]
		}
	}
	return n
}

/*
T returns the transpose of the original [][]float64. The transpose of a [][]float64
is defined in the usual manner, where every value at row x, and column y is
placed at row y, and column x. The number of rows and column of the transpose
of a slice are equal to the number of columns and rows of the original slice,
respectively. This method creates a new [][]float64, and the original is
left intact. The passed [][]float64 is assumed to be non-jagged.
*/
func T(m [][]float64) [][]float64 {
	n := New(len(m[0]), len(m))
	for i := range m {
		for j := range m[i] {
			n[j][i] = m[i][j]
		}
	}
	return n
}

/*
All checks if a supplied function is true for all elements of a mat object.
The supplied function is to have the signature of FilterFn.
For instance, consider

	positive := func(i *float64) bool {
		return *i > 0.0
	}

Then calling

	matf64.All(m, positive)

will return true if and only if all elements in m are positive.
*/
func All(m [][]float64, f FilterFn) bool {
	for i := range m {
		for j := range m[i] {
			if !f(&m[i][j]) {
				return false
			}
		}
	}
	return true
}

/*
Any checks if a supplied function is true for at least one elements of
a [][]float64. The supplied function must have the signature of FilterFn.
For instance,

	positive := func(i *float64) bool {
		return *i > 0.0
	}

Then calling

	matf64.Any(m, positive)

would be true if at least one element of the m is positive.
*/
func Any(m [][]float64, f FilterFn) bool {
	for i := range m {
		for j := range m[i] {
			if f(&m[i][j]) {
				return true
			}
		}
	}
	return false
}

/*
Sum returns the sum of all elements in a [][]float64. For example:

	m := matf64.New(10, 5)
	matf64.SetAllTo(m, 1.0)
	x := matf64.Sum(m) // x is 50.0

It is also possible for this function to return the sum of a specific row
or column in a [][]float64, by passing two additional integers to it: The
first integer must be either 0 for picking a row, or 1 for picking a column.
The second integer determines the specific row or column for which the sum is
desired. This function allow the index to be negative. For example, the sum
of the last row of a [][]float64 is given by:

	matf64.Sum(m, 0, -1)

where as the sum of the first column is given by:

	matf64.Sum(m, 1, 0)

The original [][]float64 is not mutated in this function.
*/
func Sum(m [][]float64, args ...int) float64 {
	sum := 0.0
	switch len(args) {
	case 0:
		for i := range m {
			for j := range m[i] {
				sum += m[i][j]
			}
		}
	case 2:
		switch args[0] {
		case 0:
			x := args[1]
			if x >= 0 {
				for i := range m[x] {
					sum += m[x][i]
				}
			} else {
				for i := range m[len(m)+x] {
					sum += m[len(m)+x][i]
				}
			}
		case 1:
			x := args[1]
			if x >= 0 {
				for i := range m {
					sum += m[i][x]
				}
			} else {
				for i := range m {
					sum += m[i][len(m[0])+x]
				}
			}
		default:
			s := "In matf64.%s the first argument after the [][]float64 determine the axis.\n"
			s += "It must be 0 for row, or 1 for column. but %d was passed."
			s = fmt.Sprintf(s, "Sum()", args[0])
			panic(s)
		}
	default:
		s := "In matf64.%s expected 0 or 2 arguments after the [][]float64 \n"
		s += "but received %d"
		s = fmt.Sprintf(s, "Sum()", len(args))
		panic(s)
	}
	return sum
}

/*
Prod returns the product of all elements in a [][]float64. For example:

	m := matf64.New(2, 2)
	matf64.SetAllTo(m, 2.0)
	x := matf64.Prod(m) // x is 16.0

It is also possible for this function to return the Product of a specific row
or column in a [][]float64, by passing two additional integers to it: The
first integer must be either 0 for picking a row, or 1 for picking a column.
The second integer determines the specific row or column for which the product is
desired. This function allow the index to be negative. For example, the product
of the last row of a [][]float64 is given by:

	matf64.Prod(m, 0, -1)

where as the product of the first column is given by:

	matf64.Prod(m, 1, 0)

The original [][]float64 is not mutated in this function.
*/
func Prod(m [][]float64, args ...int) float64 {
	prod := 1.0
	switch len(args) {
	case 0:
		for i := range m {
			for j := range m[i] {
				prod *= m[i][j]
			}
		}
	case 2:
		switch args[0] {
		case 0:
			x := args[1]
			if x >= 0 {
				for i := range m[x] {
					prod *= m[x][i]
				}
			} else {
				for i := range m[len(m)+x] {
					prod *= m[len(m)+x][i]
				}
			}
		case 1:
			x := args[1]
			if x >= 0 {
				for i := range m {
					prod *= m[i][x]
				}
			} else {
				for i := range m {
					prod *= m[i][len(m[0])+x]
				}
			}
		default:
			s := "In matf64.%s the first argument after the [][]float64 determines the axis.\n"
			s += "It must be 0 for row, or 1 for column. but %d was passed."
			s = fmt.Sprintf(s, "Prod()", args[0])
			panic(s)

		}
	default:
		s := "In matf64.%s expected 0 or 2 arguments after the [][]float64 \n"
		s += "but received %d"
		s = fmt.Sprintf(s, "Prod()", len(args))
		panic(s)

	}
	return prod
}

/*
Avg returns the average value of all the elements in a [][]float64. For
example:

	m := matf64.New(12, 13)
	matf64.SetAllTo(m, 1.0)
	x := matf64.Avg(m) // x is 1.0

It's also possible to return the average of a specific row or column in
a [][]float64, by passing two additional integers to it: The first integer
must be either 0 for picking a row, or 1 for picking a column. The second
integer determines the specific row or column for which the average is desired.
This function allow the index to be negative. For example, the average of the
last row of a [][]float64 is given by:

	matf64.Avg(m, 0, -1)

where as the sum of the first column is given by:

	matf64.Avg(m, 1, 0)

The original [][]float64 is not mutated in this function.
*/
func Avg(m [][]float64, args ...int) float64 {
	avg := 0.0
	sum := 0.0
	numItems := 0
	switch len(args) {
	case 0:
		for i := range m {
			for j := range m[i] {
				sum += m[i][j]
				numItems++
			}
		}
	case 2:
		switch args[0] {
		case 0:
			x := args[1]
			if x >= 0 {
				for i := range m[x] {
					sum += m[x][i]
				}
				numItems = len(m[x])
			} else {
				for i := range m[len(m)+x] {
					sum += m[len(m)+x][i]
				}
				numItems = len(m[len(m)+x])
			}
		case 1:
			x := args[1]
			if x >= 0 {
				for i := range m {
					sum += m[i][x]
				}
			} else {
				for i := range m {
					sum += m[i][len(m[0])+x]
				}
			}
			numItems = len(m)
		default:
			s := "In matf64.%s the first argument after the [][]float64 determines the axis.\n"
			s += "It must be 0 for row, or 1 for column, but %d was passed."
			s = fmt.Sprintf(s, "Avg()", args[0])
			panic(s)
		}
	default:
		s := "In matf64.%s expected 0 or 2 arguments after the [][]float64 \n"
		s += "but received %d"
		s = fmt.Sprintf(s, "Avg()", len(args))
		panic(s)
	}
	avg = sum / float64(numItems)
	return avg
}

/*
Dot is the matrix product of two [][]float64. In essence, this means that
each row of the first [][]float64 is multiplied by each column of the
second [][]float64, which creates the first row of the result.

For the sake of simplicity, it is assumed that both passed [][]float64s are
non-jagged, meaning that each row has the same number of entries as any
other row in both [][]float64, and each column has the same number of entries
as any other column in both [][]float64s passed to this function.

The regular rules of a dot product hold: for any two [][]float64s passed to
this function, the number of columns of the first must be equal to the number
of rows of the second. The resulting [][]float64 has the same number of rows
as the first [][]float64 and the same number of columns as the second.
*/
func Dot(m, n [][]float64) [][]float64 {
	res := New(len(m), len(n[0]))
	for i := range m {
		for j := range n[0] {
			for k := range m[i] {
				res[i][j] += m[i][k] * n[k][j]
			}
		}
	}
	return res
}

/*
AppendCol returns a copy of a passed [][]float64, with the second argument, a
[]float64, appended to its right side. For example, consider:

	m := matf64.New(2, 2) // [[0.0, 0.0], [0.0, 0.0]]
	v := []float64{1.0, 2.0}
	matf64.AppendCol(m, v) // [[0.0, 0.0, 1.0], [0.0, 0.0, 2.0]]

The passed arguments are not mutated by this function.
*/
func AppendCol(m [][]float64, v []float64) {
	for i := range v {
		m[i] = append(m[i], v[i])
	}
}
