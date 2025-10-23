package helper

// Get the pointer of `v`.
func Ptr[T any](v T) *T {
	return &v
}
