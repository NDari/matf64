package matf64

// TransformerFn is a function which takes a pointer to a float64 and returns
// nothing. This function type is used to transform [][]float64s in place.
type TransformerFn func(*float64)

/*
FilterFn is a type of a function that takes a pointer to a float64 and returns a bool.
These functions are used to check for the truthyness of a condition on elements of
[][]float64s.
*/
type FilterFn func(*float64) bool

/*
BinaryFn is a function that takes two float64 pointers and returns nothing.
*/
type BinaryFn func(*float64, *float64)

/*
ReducerFn are functions which aggregate data, such as summing all the
values in a [][]float64 or finding the average.
*/
type ReducerFn func([][]float64) float64

/*
NewReducer will generate a function which can be used as a customer reducer. For
example:

	sum := matf64.NewReducer(0, func(i *float64, j *float64) {
		*i += *j
	})

	m := matf64.New(4)
	matf64.SetAllTo(m, 2.0)
	s := sum(m) // s is 32.0
*/
func NewReducer(initialValue float64, f BinaryFn) ReducerFn {
	return func(m [][]float64) float64 {
		for i := range m {
			for j := range m[i] {
				f(&initialValue, &m[i][j])
			}
		}
		return initialValue
	}
}
