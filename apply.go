package matf64

import "fmt"

/*
Apply applies a given transformer function to each element of a [][]float64 or
[]float64 modifying it in place.
*/
func Apply(m interface{}, f TransformerFn) {
	switch v := m.(type) {
	case []float64:
		ApplyVec(v, f)
	case [][]float64:
		ApplyMat(v, f)
	default:
		s := "In matf64.%s, expected []float64, or [][]float64 but received type: %T."
		s = fmt.Sprintf(s, "Apply()", v)
		panic(s)
	}
}

/*
ApplyVec applies a Transformer function to each element of a []float64
modifying it in place
*/
func ApplyVec(v []float64, f TransformerFn) {
	for i := range v {
		f(&v[i])
	}
}

/*
ApplyMat applies a Transformer function to each element of a [][]float64
modifying it in place
*/
func ApplyMat(v [][]float64, f TransformerFn) {
	for i := range v {
		for j := range v[i] {
			f(&v[i][j])
		}
	}
}
