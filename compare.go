package filter

import (
	"cmp"
	"fmt"
	"reflect"
	"slices"
	"strings"
)

// compareValues sorts numbers numerically, everything else as strings.
func compareValues(a, b any) int {
	return compareValuesBy(a, b, func(s string) string { return s })
}

// compareValuesNatural compares case-insensitively after numeric coercion.
func compareValuesNatural(a, b any) int {
	return compareValuesBy(a, b, strings.ToLower)
}

func compareValuesBy(a, b any, normalize func(string) string) int {
	switch {
	case a == nil && b == nil:
		return 0
	case a == nil:
		return -1
	case b == nil:
		return 1
	}
	fa, errA := toFloat64(a)
	fb, errB := toFloat64(b)
	if errA == nil && errB == nil {
		return cmp.Compare(fa, fb)
	}
	return cmp.Compare(normalize(fmt.Sprint(a)), normalize(fmt.Sprint(b)))
}

// valuesEqual checks equality with cross-type numeric coercion.
func valuesEqual(a, b any) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if reflect.TypeOf(a).Comparable() && reflect.TypeOf(b).Comparable() && a == b {
		return true
	}
	fa, errA := toFloat64(a)
	fb, errB := toFloat64(b)
	if errA == nil && errB == nil {
		return fa == fb
	}
	return reflect.DeepEqual(a, b)
}

func containsValue(values []any, target any) bool {
	return slices.ContainsFunc(values, func(v any) bool {
		return valuesEqual(v, target)
	})
}
