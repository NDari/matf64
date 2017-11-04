package matf64

import "fmt"

/*
Set sets all elements of a [][]float64 or []float64 to a given value,
modifying it in place
*/
func Set(m interface{}, val float64) {
	switch v := m.(type) {
	case []float64:
		SetVec(v, val)
	case [][]float64:
		SetMat(v, val)
	default:
		s := "In matf64.%s, expected []float64, or [][]float64 but received type: %T."
		s = fmt.Sprintf(s, "Set()", v)
		panic(s)
	}
}

/*
SetVec sets all elements of a []float64 or []float64 to a given value,
modifying it in place
*/
func SetVec(v []float64, val float64) {
	for i := range v {
		v[i] = val
	}
}

/*
SetMat sets all elements of a [][]float64 to a given value, modifying it in place
*/
func SetMat(m [][]float64, val float64) {
	for i := range m {
		for j := range m[i] {
			m[i][j] = val
		}
	}
}
