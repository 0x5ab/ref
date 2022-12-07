package ref

// V takes address of a value x.
func V[T any](x T) *T {
	return &x
}

// S turns a slice of elements to a slice of pointers to the elements
func S[T any](s []T) []*T {
	r := make([]*T, len(s))
	for i, x := range s {
		r[i] = &x
	}
	return r
}
