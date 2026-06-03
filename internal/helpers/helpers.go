package helpers

func Map[T any, V any](arr []T, fn func(T) V) []V {
	newArr := make([]V, len(arr))

	for i, v := range arr {
		newArr[i] = fn(v)
	}

	return newArr
}

func StringFromPtr(str *string) string {
	if str == nil {
		return ""
	}

	return *str
}

func IntFromPtr(v *int) int {
	if v == nil {
		return 0
	}

	return *v
}

func IntToBool(v int) bool {
	switch v {
	case 0:
		return false
	case 1:
		return true
	default:
		panic("non boolean value")
	}
}
