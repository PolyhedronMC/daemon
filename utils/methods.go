package utils

// PanicIf Panics if the parsed value is nil or false.
func PanicIf(x interface {}, err string) {
	switch x.(type) {
	case bool:
		if (x.(bool)) {
			panic(err)
		}
	default:
		if (x == nil) {
			panic(err)
		}
	}
}
