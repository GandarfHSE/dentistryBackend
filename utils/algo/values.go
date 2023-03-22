package algo

// https://cs.opensource.google/go/x/exp/+/9ff063c7:maps/maps.go;l=20
func Values[M ~map[K]V, K comparable, V any](m M) []V {
	r := make([]V, 0, len(m))
	for _, v := range m {
		r = append(r, v)
	}
	return r
}
