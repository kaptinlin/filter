package filter

func isFalsy(v any) bool {
	if v == nil {
		return true
	}
	b, ok := v.(bool)
	return ok && !b
}

func isTruthy(v any) bool {
	return !isFalsy(v)
}
