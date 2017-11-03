package matf64

import "fmt"

/*
Mult multiples all elements of a [][]float64 by the passed value. The passed value can be
a float64, []float64, or a [][]float64. Note that due to this type switching, this
function is slower than the direct calls to MultScalar, MultVec, and MultMat.

When the passed value is a float64, then each element of the [][]float64 are multiplied
by the passed value, modifying it in place.

If the passed value is a []float64, then each row of the [][]float64 is elementally
multiplied by the corresponding entry in the passed 1D slice.

Finally, if the passed value is a [][]float64, then matf64.Mul() takes each element of the
first [][]float64 passed to it, and multiples that element by the corresponding element
in the second [][]float64 passed to this function.
*/
func Mult(m [][]float64, val interface{}) {
	switch v := val.(type) {
	case float64:
		MultScalar(m, v)
	case []float64:
		MultVec(m, v)
	case [][]float64:
		MultMat(m, v)
	default:
		s := "In matf64.%s, expected float64, []float64, or [][]float64 for the second\n"
		s += "argument, but received argument of type: %T."
		s = fmt.Sprintf(s, "Mult()", v)
		panic(s)
	}
}

/*
MultScalar multiplies all elements of a [][]float64 by a passed float64 in place.
*/
func MultScalar(m [][]float64, v float64) {
	for i := range m {
		for j := range m[i] {
			m[i][j] *= v
		}
	}
}

/*
MultVec multiples the elements of each row of a [][]float64 by a []float64, modifying
the [][]float64 in place.
*/
func MultVec(m [][]float64, v []float64) {
	for i := range m {
		for j := range v {
			m[i][j] *= v[j]
		}
	}
}

/*
MultMat multiplies element-wise a [][]float64 by another, modifying
the first in place.
*/
func MultMat(m, v [][]float64) {
	for i := range m {
		for j := range m[i] {
			m[i][j] *= v[i][j]
		}
	}

}

/*
add adds to all elements of a [][]float64 a passed value. The passed value can be
a float64, []float64, or a [][]float64. Note that due to this type switching, this
function is slower than the direct calls to AddScalar, AddVec, and AddMat.

When the passed value is a float64, then each element of the [][]float64 are increased
by the passed value, modifying it in place.

If the passed value is a []float64, then each row of the [][]float64 is elementally
increased by the corresponding entry in the passed 1D slice.

Finally, if the passed value is a [][]float64, then matf64.Add() takes each element of the
first [][]float64 , and add the corresponding element of the second [][]float64 to it.
*/
func Add(m [][]float64, val interface{}) {
	switch v := val.(type) {
	case float64:
		AddScalar(m, v)
	case []float64:
		AddVec(m, v)
	case [][]float64:
		AddMat(m, v)
	default:
		s := "In matf64.%s, expected float64, []float64, or [][]float64 for the second\n"
		s += "argument, but received argument of type: %T."
		s = fmt.Sprintf(s, "Add()", v)
		panic(s)
	}
}

/*
MultScalar increases all elements of a [][]float64 by a passed float64 in place.
*/
func AddScalar(m [][]float64, v float64) {
	for i := range m {
		for j := range m[i] {
			m[i][j] += v
		}
	}
}

/*
AddVec increases the elements of each row of a [][]float64 by a []float64, modifying
the [][]float64 in place.
*/
func AddVec(m [][]float64, v []float64) {
	for i := range m {
		for j := range v {
			m[i][j] += v[j]
		}
	}
}

/*
AddMat is an element-wise addition of a [][]float64 by another, modifying
the first in place.
*/
func AddMat(m, v [][]float64) {
	for i := range m {
		for j := range m[i] {
			m[i][j] += v[i][j]
		}
	}

}

/*
Sub subtracts from all elements of a [][]float64 a passed value. The passed value can be
a float64, []float64, or a [][]float64. Note that due to this type switching, this
function is slower than the direct calls to SubScalar, SubVec, and SubMat.

When the passed value is a float64, then each element of the [][]float64 are decreased
by the passed value, modifying it in place.

If the passed value is a []float64, then each row of the [][]float64 is elementally
decreased by the corresponding entry in the passed 1D slice.

Finally, if the passed value is a [][]float64, then matf64.Sub() takes each element of the
first [][]float64 , and subtracts the corresponding element of the second [][]float64 from
it
*/
func Sub(m [][]float64, val interface{}) {
	switch v := val.(type) {
	case float64:
		SubScalar(m, v)
	case []float64:
		SubVec(m, v)
	case [][]float64:
		SubMat(m, v)
	default:
		s := "In matf64.%s, expected float64, []float64, or [][]float64 for the second\n"
		s += "argument, but received argument of type: %T."
		s = fmt.Sprintf(s, "Sub()", v)
		panic(s)
	}
}

/*
SubScalar decreases all elements of a [][]float64 by a passed float64 in place.
*/
func SubScalar(m [][]float64, v float64) {
	for i := range m {
		for j := range m[i] {
			m[i][j] -= v
		}
	}
}

/*
AddVec increases the elements of each row of a [][]float64 by a []float64, modifying
the [][]float64 in place.
*/
func SubVec(m [][]float64, v []float64) {
	for i := range m {
		for j := range v {
			m[i][j] -= v[j]
		}
	}
}

/*
AddMat is an element-wise addition of a [][]float64 by another, modifying
the first in place.
*/
func SubMat(m, v [][]float64) {
	for i := range m {
		for j := range m[i] {
			m[i][j] -= v[i][j]
		}
	}

}

/*
Sub subtracts from all elements of a [][]float64 a passed value. The passed value can be
a float64, []float64, or a [][]float64. Note that due to this type switching, this
function is slower than the direct calls to SubScalar, SubVec, and SubMat.

When the passed value is a float64, then each element of the [][]float64 are decreased
by the passed value, modifying it in place.

If the passed value is a []float64, then each row of the [][]float64 is elementally
decreased by the corresponding entry in the passed 1D slice.

Finally, if the passed value is a [][]float64, then matf64.Sub() takes each element of the
first [][]float64 , and subtracts the corresponding element of the second [][]float64 from
it
*/
func Div(m [][]float64, val interface{}) {
	switch v := val.(type) {
	case float64:
		DivScalar(m, v)
	case []float64:
		DivVec(m, v)
	case [][]float64:
		DivMat(m, v)
	default:
		s := "In matf64.%s, expected float64, []float64, or [][]float64 for the second\n"
		s += "argument, but received argument of type: %T."
		s = fmt.Sprintf(s, "Div()", v)
		panic(s)
	}
}

/*
SubScalar decreases all elements of a [][]float64 by a passed float64 in place.
*/
func DivScalar(m [][]float64, v float64) {
	for i := range m {
		for j := range m[i] {
			m[i][j] /= v
		}
	}
}

/*
AddVec increases the elements of each row of a [][]float64 by a []float64, modifying
the [][]float64 in place.
*/
func DivVec(m [][]float64, v []float64) {
	for i := range m {
		for j := range v {
			m[i][j] /= v[j]
		}
	}
}

/*
AddMat is an element-wise addition of a [][]float64 by another, modifying
the first in place.
*/
func DivMat(m, v [][]float64) {
	for i := range m {
		for j := range m[i] {
			m[i][j] /= v[i][j]
		}
	}

}
