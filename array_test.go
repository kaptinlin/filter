package filter

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/require"
)

func TestUnique(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     any
		want      []any
		expectErr bool
	}{
		{
			name:      "Unique elements in a slice of int",
			input:     []any{1, 2, 3, 2, 1},
			want:      []any{1, 2, 3},
			expectErr: false,
		},
		{
			name:      "Unique elements in a slice of string",
			input:     []any{"apple", "banana", "apple", "cherry"},
			want:      []any{"apple", "banana", "cherry"},
			expectErr: false,
		},
		{
			name:      "Unique elements with mixed types",
			input:     []any{1, "apple", 1, "banana", "apple"},
			want:      []any{1, "apple", "banana"},
			expectErr: false,
		},
		{
			name:      "Empty slice",
			input:     []any{},
			want:      []any{},
			expectErr: false,
		},
		{
			name:      "Unsupported input type",
			input:     "not-a-slice",
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := Unique(tt.input)
			if tt.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Fatalf("Unique() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestUnique_NonComparableValues(t *testing.T) {
	t.Parallel()

	input := []any{
		map[string]any{"name": "shirt", "meta": []any{"blue", 42}},
		map[string]any{"name": "shirt", "meta": []any{"blue", 42}},
	}

	_, err := Unique(input)
	require.ErrorIs(t, err, ErrInvalidInput)
}

func TestUniqueBy(t *testing.T) {
	t.Parallel()

	products := []any{
		map[string]any{"handle": "shirt", "variant": "red"},
		map[string]any{"handle": "shoe", "variant": "blue"},
		map[string]any{"handle": "shirt", "variant": "green"},
	}

	got, err := UniqueBy(products, "handle")
	require.NoError(t, err)
	want := []any{
		map[string]any{"handle": "shirt", "variant": "red"},
		map[string]any{"handle": "shoe", "variant": "blue"},
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Fatalf("UniqueBy() mismatch (-want +got):\n%s", diff)
	}
}

func TestUniqueByNestedKey(t *testing.T) {
	t.Parallel()

	input := []any{
		map[string]any{"user": map[string]any{"id": 1}, "name": "Alice"},
		map[string]any{"user": map[string]any{"id": 2}, "name": "Bob"},
		map[string]any{"user": map[string]any{"id": 1}, "name": "Alicia"},
	}

	got, err := UniqueBy(input, "user.id")
	require.NoError(t, err)
	want := []any{
		map[string]any{"user": map[string]any{"id": 1}, "name": "Alice"},
		map[string]any{"user": map[string]any{"id": 2}, "name": "Bob"},
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Fatalf("UniqueBy() mismatch (-want +got):\n%s", diff)
	}
}

func TestUniqueByCrossTypeNumericKey(t *testing.T) {
	t.Parallel()

	oneInt := map[string]any{"id": 1, "name": "int"}
	oneString := map[string]any{"id": "1", "name": "string"}
	two := map[string]any{"id": 2, "name": "two"}
	input := []any{oneInt, oneString, two}

	got, err := UniqueBy(input, "id")
	require.NoError(t, err)
	want := []any{oneInt, two}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Fatalf("UniqueBy() mismatch (-want +got):\n%s", diff)
	}
}

func TestUniqueByNonComparableKey(t *testing.T) {
	t.Parallel()

	input := []any{
		map[string]any{"tags": []any{"sale", "new"}, "name": "Shoes"},
		map[string]any{"tags": []any{"clearance"}, "name": "Hat"},
		map[string]any{"tags": []any{"sale", "new"}, "name": "Shirt"},
	}

	got, err := UniqueBy(input, "tags")
	require.NoError(t, err)
	want := []any{
		map[string]any{"tags": []any{"sale", "new"}, "name": "Shoes"},
		map[string]any{"tags": []any{"clearance"}, "name": "Hat"},
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Fatalf("UniqueBy() mismatch (-want +got):\n%s", diff)
	}
}

func TestUniqueByMissingKey(t *testing.T) {
	t.Parallel()

	_, err := UniqueBy([]any{map[string]any{"name": "Shoes"}}, "handle")
	require.ErrorIs(t, err, ErrNotFound)
}

func TestJoin(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     any
		separator string
		want      string
		expectErr bool
	}{
		{
			name:      "Join with comma separator",
			input:     []any{"apple", "banana", "cherry"},
			separator: ",",
			want:      "apple,banana,cherry",
			expectErr: false,
		},
		{
			name:      "Join with mixed types",
			input:     []any{1, "apple", 2.5},
			separator: ":",
			want:      "1:apple:2.5",
			expectErr: false,
		},
		{
			name:      "Join with empty separator",
			input:     []any{"apple", "banana", "cherry"},
			separator: "",
			want:      "applebananacherry",
			expectErr: false,
		},
		{
			name:      "Empty slice",
			input:     []any{},
			separator: ",",
			want:      "",
			expectErr: false,
		},
		{
			name:      "Unsupported input type",
			input:     "not-a-slice",
			separator: ",",
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := Join(tt.input, tt.separator)
			if tt.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.want, got, "The joined string should match the expected output")
			}
		})
	}
}

func TestIndex(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     any
		index     int
		want      any
		expectErr bool
	}{
		{
			name:      "Valid index in slice of int",
			input:     []any{1, 2, 3},
			index:     1,
			want:      2,
			expectErr: false,
		},
		{
			name:      "Valid index in slice of string",
			input:     []any{"apple", "banana", "cherry"},
			index:     2,
			want:      "cherry",
			expectErr: false,
		},
		{
			name:      "Index out of range (too high)",
			input:     []any{"apple", "banana", "cherry"},
			index:     5,
			want:      nil,
			expectErr: true,
		},
		{
			name:      "Index out of range (negative)",
			input:     []any{"apple", "banana", "cherry"},
			index:     -1,
			want:      nil,
			expectErr: true,
		},
		{
			name:      "Empty slice",
			input:     []any{},
			index:     0,
			want:      nil,
			expectErr: true,
		},
		{
			name:      "Unsupported input type",
			input:     "not-a-slice",
			index:     0,
			want:      nil,
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := Index(tt.input, tt.index)
			if tt.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.want, got, "The element at the specified index should match the expected value")
			}
		})
	}
}

func TestLast(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     any
		want      any
		expectErr bool
	}{
		{
			name:      "Non-empty slice of int",
			input:     []any{1, 2, 3},
			want:      3,
			expectErr: false,
		},
		{
			name:      "Non-empty slice of string",
			input:     []any{"apple", "banana", "cherry"},
			want:      "cherry",
			expectErr: false,
		},
		{
			name:      "Non-empty mixed slice",
			input:     []any{1, "banana", 3.5},
			want:      3.5,
			expectErr: false,
		},
		{
			name:      "Empty slice",
			input:     []any{},
			want:      nil,
			expectErr: true,
		},
		{
			name:      "Unsupported input type",
			input:     "not-a-slice",
			want:      nil,
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := Last(tt.input)
			if tt.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.want, got, "The last element should match the expected value")
			}
		})
	}
}

func TestRandom(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     any
		expectErr bool
	}{
		{
			name:      "Non-empty slice of int",
			input:     []any{1, 2, 3},
			expectErr: false,
		},
		{
			name:      "Non-empty slice of string",
			input:     []any{"apple", "banana", "cherry"},
			expectErr: false,
		},
		{
			name:      "Empty slice",
			input:     []any{},
			expectErr: true,
		},
		{
			name:      "Unsupported input type",
			input:     "not-a-slice",
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := Random(tt.input)
			if tt.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Contains(t, tt.input, got, "The returned element should belong to the input slice")
			}
		})
	}
}

func TestRandomWithRandDeterministic(t *testing.T) {
	t.Parallel()

	input := []any{"red", "green", "blue", "yellow"}
	a, err := RandomWithRand(SeededRand(1, 2), input)
	require.NoError(t, err)
	b, err := RandomWithRand(SeededRand(1, 2), input)
	require.NoError(t, err)

	require.Equal(t, a, b)
	require.Contains(t, input, a)
}

func TestRandomWithRandRejectsNilRand(t *testing.T) {
	t.Parallel()

	_, err := RandomWithRand(nil, []any{1})
	require.ErrorIs(t, err, ErrInvalidInput)
}

func TestReverse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     any
		want      []any
		expectErr bool
	}{
		{
			name:      "Reverse a slice of int",
			input:     []any{1, 2, 3, 4},
			want:      []any{4, 3, 2, 1},
			expectErr: false,
		},
		{
			name:      "Reverse a slice of string",
			input:     []any{"apple", "banana", "cherry"},
			want:      []any{"cherry", "banana", "apple"},
			expectErr: false,
		},
		{
			name:      "Reverse a mixed type slice",
			input:     []any{1, "banana", 3.5},
			want:      []any{3.5, "banana", 1},
			expectErr: false,
		},
		{
			name:      "Empty slice",
			input:     []any{},
			want:      []any{},
			expectErr: false,
		},
		{
			name:      "Unsupported input type",
			input:     "not-a-slice",
			want:      nil,
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := Reverse(tt.input)
			if tt.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Fatalf("Reverse() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestShuffle(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     any
		expectErr bool
	}{
		{
			name:      "Shuffle a slice of int",
			input:     []any{1, 2, 3, 4},
			expectErr: false,
		},
		{
			name:      "Shuffle a slice of string",
			input:     []any{"apple", "banana", "cherry"},
			expectErr: false,
		},
		{
			name:      "Shuffle a mixed type slice",
			input:     []any{1, "banana", 3.5},
			expectErr: false,
		},
		{
			name:      "Empty slice",
			input:     []any{},
			expectErr: false,
		},
		{
			name:      "Unsupported input type",
			input:     "not-a-slice",
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			shuffled, err := Shuffle(tt.input)
			if tt.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				if diff := cmp.Diff(tt.input, shuffled, cmpopts.SortSlices(func(a, b any) bool {
					return fmt.Sprint(a) < fmt.Sprint(b)
				})); diff != "" {
					t.Fatalf("Shuffle() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestShuffleWithRandDeterministic(t *testing.T) {
	t.Parallel()

	input := []any{1, 2, 3, 4, 5}
	a, err := ShuffleWithRand(SeededRand(1, 2), input)
	require.NoError(t, err)
	b, err := ShuffleWithRand(SeededRand(1, 2), input)
	require.NoError(t, err)

	if diff := cmp.Diff(a, b); diff != "" {
		t.Fatalf("ShuffleWithRand() mismatch (-a +b):\n%s", diff)
	}
	if diff := cmp.Diff(input, a, cmpopts.SortSlices(func(a, b any) bool {
		return fmt.Sprint(a) < fmt.Sprint(b)
	})); diff != "" {
		t.Fatalf("ShuffleWithRand() changed elements (-want +got):\n%s", diff)
	}
}

func TestShuffleWithRandRejectsNilRand(t *testing.T) {
	t.Parallel()

	_, err := ShuffleWithRand(nil, []any{1})
	require.ErrorIs(t, err, ErrInvalidInput)
}

func TestSize(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     any
		want      int
		expectErr bool
	}{
		{
			name:      "Size of a non-empty slice of int",
			input:     []any{1, 2, 3, 4},
			want:      4,
			expectErr: false,
		},
		{
			name:      "Size of a non-empty slice of string",
			input:     []any{"apple", "banana", "cherry"},
			want:      3,
			expectErr: false,
		},
		{
			name:      "Size of a mixed type slice",
			input:     []any{1, "banana", 3.5},
			want:      3,
			expectErr: false,
		},
		{
			name:      "Size of an empty slice",
			input:     []any{},
			want:      0,
			expectErr: false,
		},
		{
			name:      "Size of a map with string keys",
			input:     map[string]any{"a": 1, "b": 2, "c": 3},
			want:      3,
			expectErr: false,
		},
		{
			name:      "Size of an empty map",
			input:     map[string]any{},
			want:      0,
			expectErr: false,
		},
		{
			name:      "Size of an array",
			input:     [3]int{1, 2, 3},
			want:      3,
			expectErr: false,
		},
		{
			name:      "Size of UTF-8 string",
			input:     "a界b",
			want:      3,
			expectErr: false,
		},
		{
			name:      "Unsupported input type (number)",
			input:     42,
			want:      0,
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := Size(tt.input)
			if tt.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.want, got, "The size should match the expected value")
			}
		})
	}
}

func TestMax(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     any
		want      float64
		expectErr bool
	}{
		{
			name:      "Non-empty slice of float64",
			input:     []any{1.2, 3.4, 5.6, 4.5},
			want:      5.6,
			expectErr: false,
		},
		{
			name:      "Slice with negative values",
			input:     []any{-1.1, -3.3, -2.2},
			want:      -1.1,
			expectErr: false,
		},
		{
			name:      "Slice with a single element",
			input:     []any{42.0},
			want:      42.0,
			expectErr: false,
		},
		{
			name:      "Empty slice",
			input:     []any{},
			expectErr: true,
		},
		{
			name:      "Unsupported input type (not a slice)",
			input:     "not-a-slice",
			expectErr: true,
		},
		{
			name:      "Unsupported element type in slice",
			input:     []any{1.1, "not-a-float", 3.3},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := Max(tt.input)
			if tt.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.want, got, "The maximum value should match the expected value")
			}
		})
	}
}

func TestMin(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     any
		want      float64
		expectErr bool
	}{
		{
			name:      "Non-empty slice of float64 with positive values",
			input:     []any{2.2, 1.1, 3.3, 4.4},
			want:      1.1,
			expectErr: false,
		},
		{
			name:      "Slice with negative and positive values",
			input:     []any{-1.2, 0.0, -2.2, 2.2},
			want:      -2.2,
			expectErr: false,
		},
		{
			name:      "Slice with a single element",
			input:     []any{42.0},
			want:      42.0,
			expectErr: false,
		},
		{
			name:      "Empty slice",
			input:     []any{},
			expectErr: true,
		},
		{
			name:      "Unsupported input type (not a slice)",
			input:     "not-a-slice",
			expectErr: true,
		},
		{
			name:      "Unsupported element type in slice",
			input:     []any{1.1, "not-a-float", 3.3},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := Min(tt.input)
			if tt.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.want, got, "The minimum value should match the expected value")
			}
		})
	}
}

func TestSum(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     any
		want      float64
		expectErr bool
	}{
		{
			name:      "Slice of positive float64",
			input:     []any{1.1, 2.2, 3.3},
			want:      6.6,
			expectErr: false,
		},
		{
			name:      "Slice of negative float64",
			input:     []any{-1.1, -2.2, -3.3},
			want:      -6.6,
			expectErr: false,
		},
		{
			name:      "Mixed positive and negative float64",
			input:     []any{-1.1, 2.2, -3.3, 4.4},
			want:      2.2,
			expectErr: false,
		},
		{
			name:      "Slice with a single element",
			input:     []any{42.0},
			want:      42.0,
			expectErr: false,
		},
		{
			name:      "Empty slice",
			input:     []any{},
			want:      0.0,
			expectErr: false,
		},
		{
			name:      "Unsupported input type (not a slice)",
			input:     "not-a-slice",
			want:      0,
			expectErr: true,
		},
		{
			name:      "Unsupported element type in slice",
			input:     []any{1.1, "not-a-float", 3.3},
			want:      0,
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := Sum(tt.input)
			if tt.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				if tt.want == 0.0 {
					require.Equal(t, tt.want, got, "The sum of the slice should exactly match the expected value")
				} else {
					// Use a tolerance for comparing floating-point numbers
					tolerance := 1e-9 // Adjust the tolerance based on the precision you need
					require.InEpsilon(t, tt.want, got, tolerance, "The sum of the slice should match the expected value within a tolerance")
				}
			}
		})
	}
}

func TestNumericAggregatesCoerceStrings(t *testing.T) {
	t.Parallel()

	input := []any{"1.5", 2, int64(3)}

	maximum, err := Max(input)
	require.NoError(t, err)
	require.Equal(t, 3.0, maximum)

	minimum, err := Min(input)
	require.NoError(t, err)
	require.Equal(t, 1.5, minimum)

	sum, err := Sum(input)
	require.NoError(t, err)
	require.Equal(t, 6.5, sum)

	average, err := Average(input)
	require.NoError(t, err)
	require.InEpsilon(t, 6.5/3, average, 1e-9)
}

func TestSumBy(t *testing.T) {
	t.Parallel()

	products := []any{
		map[string]any{"title": "Shoes", "price": 50},
		map[string]any{"title": "Shirt", "price": "30.5"},
		map[string]any{"title": "Hat", "price": 10.25},
	}

	got, err := SumBy(products, "price")
	require.NoError(t, err)
	require.InEpsilon(t, 90.75, got, 1e-9)
}

func TestSumByEmpty(t *testing.T) {
	t.Parallel()

	got, err := SumBy([]any{}, "price")
	require.NoError(t, err)
	require.Equal(t, 0.0, got)
}

func TestSumByMissingKey(t *testing.T) {
	t.Parallel()

	_, err := SumBy([]any{map[string]any{"title": "Shoes"}}, "price")
	require.ErrorIs(t, err, ErrNotFound)
}

func TestSumByNonNumeric(t *testing.T) {
	t.Parallel()

	_, err := SumBy([]any{map[string]any{"price": "free"}}, "price")
	require.ErrorIs(t, err, ErrFormat)
}

func TestCollectionMissingPolicies(t *testing.T) {
	t.Parallel()

	active := map[string]any{"name": "Ada", "active": true, "rank": 2, "price": 10}
	inactive := map[string]any{"name": "Bob", "active": false, "rank": 1, "price": 20}
	missing := map[string]any{"title": "Unknown"}
	records := []map[string]any{active, inactive, missing}

	t.Run("map substitutes nil", func(t *testing.T) {
		t.Parallel()

		got, err := Map(records, "name")
		require.NoError(t, err)
		if diff := cmp.Diff([]any{"Ada", "Bob", nil}, got); diff != "" {
			t.Fatalf("Map() mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("sort treats missing as nil", func(t *testing.T) {
		t.Parallel()

		got, err := Sort(records, "rank")
		require.NoError(t, err)
		want := []any{missing, inactive, active}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Fatalf("Sort() mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("sort natural treats missing as nil", func(t *testing.T) {
		t.Parallel()

		got, err := SortNatural(records, "name")
		require.NoError(t, err)
		want := []any{missing, active, inactive}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Fatalf("SortNatural() mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("compact skips missing", func(t *testing.T) {
		t.Parallel()

		got, err := Compact(records, "name")
		require.NoError(t, err)
		want := []any{active, inactive}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Fatalf("Compact() mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("where treats missing as no match", func(t *testing.T) {
		t.Parallel()

		got, err := Where(records, "active")
		require.NoError(t, err)
		want := []any{active}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Fatalf("Where() mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("reject keeps missing because it did not match", func(t *testing.T) {
		t.Parallel()

		got, err := Reject(records, "active")
		require.NoError(t, err)
		want := []any{inactive, missing}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Fatalf("Reject() mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("find skips missing and reports not found", func(t *testing.T) {
		t.Parallel()

		got, err := Find(records, "name", "Ada")
		require.NoError(t, err)
		if diff := cmp.Diff(active, got); diff != "" {
			t.Fatalf("Find() mismatch (-want +got):\n%s", diff)
		}

		_, err = Find(records, "missing", "value")
		require.ErrorIs(t, err, ErrNotFound)
	})

	t.Run("find index skips missing", func(t *testing.T) {
		t.Parallel()

		got, err := FindIndex(records, "name", "Bob")
		require.NoError(t, err)
		require.Equal(t, 1, got)

		got, err = FindIndex(records, "missing", "value")
		require.NoError(t, err)
		require.Equal(t, -1, got)
	})

	t.Run("has treats missing as no match", func(t *testing.T) {
		t.Parallel()

		got, err := Has(records, "name", "Bob")
		require.NoError(t, err)
		require.True(t, got)

		got, err = Has(records, "missing")
		require.NoError(t, err)
		require.False(t, got)
	})

	t.Run("unique by fails on missing", func(t *testing.T) {
		t.Parallel()

		_, err := UniqueBy(records, "name")
		require.ErrorIs(t, err, ErrNotFound)
	})

	t.Run("sum by fails on missing", func(t *testing.T) {
		t.Parallel()

		_, err := SumBy(records, "price")
		require.ErrorIs(t, err, ErrNotFound)
	})
}

func TestAverage(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     any
		want      float64
		expectErr bool
	}{
		{
			name:      "Average of positive float64",
			input:     []any{1.0, 2.0, 3.0},
			want:      2.0,
			expectErr: false,
		},
		{
			name:      "Average of negative float64",
			input:     []any{-1.0, -2.0, -3.0},
			want:      -2.0,
			expectErr: false,
		},
		{
			name:      "Mixed positive and negative float64",
			input:     []any{-2.0, 2.0},
			want:      0.0,
			expectErr: false,
		},
		{
			name:      "Slice with a single element",
			input:     []any{42.0},
			want:      42.0,
			expectErr: false,
		},
		{
			name:      "Typed int slice",
			input:     []int{1, 2, 3},
			want:      2.0,
			expectErr: false,
		},
		{
			name:      "Empty slice",
			input:     []any{},
			want:      0.0,
			expectErr: true,
		},
		{
			name:      "Unsupported input type (not a slice)",
			input:     "not-a-slice",
			want:      0,
			expectErr: true,
		},
		{
			name:      "Unsupported element type in slice",
			input:     []any{1.1, "not-a-float", 3.3},
			want:      0,
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := Average(tt.input)
			if tt.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.want, got, "The average of the slice should match the expected value")
			}
		})
	}
}

func TestMap(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     any
		key       string
		want      []any
		expectErr bool
	}{
		{
			name: "All maps contain the key",
			input: []any{
				map[any]any{"key1": "value1"},
				map[any]any{"key1": "value2"},
			},
			key:  "key1",
			want: []any{"value1", "value2"},
		},
		{
			name: "Some maps do not contain the key",
			input: []any{
				map[any]any{"key1": "value1"},
				map[any]any{"key2": "value2"},
			},
			key:  "key1",
			want: []any{"value1", nil},
		},
		{
			name: "Key not present in any map",
			input: []any{
				map[any]any{"key2": "value1"},
				map[any]any{"key3": "value2"},
			},
			key:  "key1",
			want: []any{nil, nil},
		},
		{
			name:  "Empty slice",
			input: []any{},
			key:   "key1",
			want:  []any{},
		},
		{
			name:      "Unsupported input type",
			input:     "not-a-slice",
			key:       "key1",
			want:      nil,
			expectErr: true,
		},
		{
			name: "Single level of nesting",
			input: []any{
				map[any]any{"parent": map[any]any{"child": "value1"}},
				map[any]any{"parent": map[any]any{"child": "value2"}},
			},
			key:  "parent.child",
			want: []any{"value1", "value2"},
		},
		{
			name: "Multiple levels of nesting",
			input: []any{
				map[any]any{"level1": map[any]any{"level2": map[any]any{"level3": "deepValue1"}}},
				map[any]any{"level1": map[any]any{"level2": map[any]any{"level3": "deepValue2"}}},
			},
			key:  "level1.level2.level3",
			want: []any{"deepValue1", "deepValue2"},
		},
		{
			name: "Mixed nesting and types",
			input: []any{
				map[any]any{"parent": []any{"value1", map[any]any{"child": "nestedValue1"}}},
				map[any]any{"parent": map[any]any{"child": "nestedValue2"}},
			},
			key:  "parent.1.child",
			want: []any{"nestedValue1", nil},
		},
		{
			name: "Value is a slice",
			input: []any{
				map[any]any{"parent": []any{"value1", "value2"}},
				map[any]any{"parent": []any{"value3", "value4"}},
			},
			key:  "parent",
			want: []any{[]any{"value1", "value2"}, []any{"value3", "value4"}},
		},
		{
			name: "Value is a map",
			input: []any{
				map[any]any{"parent": map[any]any{"child": "value"}},
				map[any]any{"parent": map[any]any{"child": "value"}},
			},
			key:  "parent",
			want: []any{map[any]any{"child": "value"}, map[any]any{"child": "value"}},
		},
		{
			name: "Intermediate key leads to a non-map and non-slice element",
			input: []any{
				map[any]any{"parent": "notAMapOrSlice"},
			},
			key:  "parent.child",
			want: []any{nil},
		},
		{
			name: "Key missing at intermediate level",
			input: []any{
				map[any]any{"level1": map[any]any{}}, // Empty map at level1
			},
			key:  "level1.level2.missing",
			want: []any{nil},
		},
		{
			name: "Index out of range in slice",
			input: []any{
				map[any]any{"list": []any{"value1", "value2"}},
			},
			key:  "list.2",
			want: []any{nil}, // Index 2 is out of range
		},
		{
			name: "Incorrect index format in key",
			input: []any{
				map[any]any{"list": []any{"value1", "value2"}},
			},
			key:  "list.two",
			want: []any{nil},
		},
		{
			name: "Nested key leading to array index",
			input: []any{
				map[any]any{"parent": []any{"value1", "value2"}},
			},
			key:  "parent.1",
			want: []any{"value2"},
		},
		{
			name: "Deep nesting with arrays and maps",
			input: []any{
				map[any]any{"level1": []any{map[any]any{"level2": []any{"value1"}}}},
			},
			key:  "level1.0.level2.0",
			want: []any{"value1"},
		},
		{
			name: "Nesting with a non-existing final key",
			input: []any{
				map[any]any{"exists": map[any]any{"doesNotExist": "value"}},
			},
			key:  "exists.notHere",
			want: []any{nil},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := Map(tt.input, tt.key)
			if tt.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Fatalf("Map() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestSort(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   any
		key     []string
		want    []any
		wantErr bool
	}{
		{
			name:  "Sort strings",
			input: []any{"banana", "apple", "cherry"},
			want:  []any{"apple", "banana", "cherry"},
		},
		{
			name:  "Sort numbers",
			input: []any{3.0, 1.0, 2.0},
			want:  []any{1.0, 2.0, 3.0},
		},
		{
			name:  "Sort cross-type numeric values",
			input: []any{"10", 2, int64(1), "3.5"},
			want:  []any{int64(1), 2, "3.5", "10"},
		},
		{
			name: "Sort by key",
			input: []any{
				map[string]any{"name": "Charlie", "age": 30},
				map[string]any{"name": "Alice", "age": 25},
				map[string]any{"name": "Bob", "age": 35},
			},
			key: []string{"name"},
			want: []any{
				map[string]any{"name": "Alice", "age": 25},
				map[string]any{"name": "Bob", "age": 35},
				map[string]any{"name": "Charlie", "age": 30},
			},
		},
		{
			name: "Sort by numeric key",
			input: []any{
				map[string]any{"name": "C", "price": 30.0},
				map[string]any{"name": "A", "price": 10.0},
				map[string]any{"name": "B", "price": 20.0},
			},
			key: []string{"price"},
			want: []any{
				map[string]any{"name": "A", "price": 10.0},
				map[string]any{"name": "B", "price": 20.0},
				map[string]any{"name": "C", "price": 30.0},
			},
		},
		{
			name:  "Sort with nil elements",
			input: []any{nil, "b", "a", nil},
			want:  []any{nil, nil, "a", "b"},
		},
		{
			name:  "Empty slice",
			input: []any{},
			want:  []any{},
		},
		{
			name:    "Not a slice",
			input:   "not-a-slice",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := Sort(tt.input, tt.key...)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Fatalf("result mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestSortNatural(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input any
		key   []string
		want  []any
	}{
		{
			name:  "Case insensitive sort",
			input: []any{"Banana", "apple", "Cherry"},
			want:  []any{"apple", "Banana", "Cherry"},
		},
		{
			name:  "All same case",
			input: []any{"c", "a", "b"},
			want:  []any{"a", "b", "c"},
		},
		{
			name:  "Numeric strings sort by value",
			input: []any{"10", "2", "1"},
			want:  []any{"1", "2", "10"},
		},
		{
			name: "Sort by key case insensitive",
			input: []any{
				map[string]any{"name": "charlie"},
				map[string]any{"name": "Alice"},
				map[string]any{"name": "Bob"},
			},
			key: []string{"name"},
			want: []any{
				map[string]any{"name": "Alice"},
				map[string]any{"name": "Bob"},
				map[string]any{"name": "charlie"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := SortNatural(tt.input, tt.key...)
			require.NoError(t, err)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Fatalf("result mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestCompact(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input any
		key   []string
		want  []any
	}{
		{
			name:  "Remove nils",
			input: []any{"a", nil, "b", nil, "c"},
			want:  []any{"a", "b", "c"},
		},
		{
			name:  "No nils",
			input: []any{"a", "b", "c"},
			want:  []any{"a", "b", "c"},
		},
		{
			name:  "All nils",
			input: []any{nil, nil, nil},
			want:  []any{},
		},
		{
			name:  "Empty slice",
			input: []any{},
			want:  []any{},
		},
		{
			name: "With key",
			input: []any{
				map[string]any{"name": "Alice"},
				map[string]any{"other": "Bob"},
				map[string]any{"name": "Charlie"},
			},
			key: []string{"name"},
			want: []any{
				map[string]any{"name": "Alice"},
				map[string]any{"name": "Charlie"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := Compact(tt.input, tt.key...)
			require.NoError(t, err)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Fatalf("result mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestConcat(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   any
		other   any
		want    []any
		wantErr bool
	}{
		{
			name:  "Two string slices",
			input: []any{"a", "b"},
			other: []any{"c", "d"},
			want:  []any{"a", "b", "c", "d"},
		},
		{
			name:  "First empty",
			input: []any{},
			other: []any{"a", "b"},
			want:  []any{"a", "b"},
		},
		{
			name:  "Second empty",
			input: []any{"a", "b"},
			other: []any{},
			want:  []any{"a", "b"},
		},
		{
			name:  "Both empty",
			input: []any{},
			other: []any{},
			want:  []any{},
		},
		{
			name:    "First not a slice",
			input:   "not-a-slice",
			other:   []any{"a"},
			wantErr: true,
		},
		{
			name:    "Second not a slice",
			input:   []any{"a"},
			other:   "not-a-slice",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := Concat(tt.input, tt.other)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Fatalf("result mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestWhere(t *testing.T) {
	t.Parallel()

	products := []any{
		map[string]any{"name": "Shoes", "available": true, "price": 50.0},
		map[string]any{"name": "Shirt", "available": false, "price": 30.0},
		map[string]any{"name": "Pants", "available": true, "price": 40.0},
	}

	tests := []struct {
		name  string
		input any
		key   string
		value []any
		want  []any
	}{
		{
			name:  "Where with value",
			input: products,
			key:   "available",
			value: []any{true},
			want: []any{
				map[string]any{"name": "Shoes", "available": true, "price": 50.0},
				map[string]any{"name": "Pants", "available": true, "price": 40.0},
			},
		},
		{
			name:  "Where truthy (no value)",
			input: products,
			key:   "available",
			want: []any{
				map[string]any{"name": "Shoes", "available": true, "price": 50.0},
				map[string]any{"name": "Pants", "available": true, "price": 40.0},
			},
		},
		{
			name:  "Where with string value",
			input: products,
			key:   "name",
			value: []any{"Shirt"},
			want: []any{
				map[string]any{"name": "Shirt", "available": false, "price": 30.0},
			},
		},
		{
			name:  "Where no match",
			input: products,
			key:   "name",
			value: []any{"Hat"},
			want:  []any{},
		},
		{
			name:  "Empty slice",
			input: []any{},
			key:   "name",
			want:  []any{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := Where(tt.input, tt.key, tt.value...)
			require.NoError(t, err)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Fatalf("result mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestWhereWithNonComparableValue(t *testing.T) {
	t.Parallel()

	products := []any{
		map[string]any{"name": "Shoes", "tags": []any{"sale", "new"}},
		map[string]any{"name": "Shirt", "tags": []any{"clearance"}},
	}

	got, err := Where(products, "tags", []any{"sale", "new"})
	require.NoError(t, err)
	want := []any{
		map[string]any{"name": "Shoes", "tags": []any{"sale", "new"}},
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Fatalf("Where() mismatch (-want +got):\n%s", diff)
	}
}

func TestCollectionNumericEquality(t *testing.T) {
	t.Parallel()

	oneInt := map[string]any{"id": 1, "name": "int"}
	oneString := map[string]any{"id": "1", "name": "string"}
	two := map[string]any{"id": 2, "name": "two"}
	records := []any{oneInt, oneString, two}

	t.Run("where matches numeric strings and numbers", func(t *testing.T) {
		t.Parallel()

		got, err := Where(records, "id", 1.0)
		require.NoError(t, err)
		want := []any{oneInt, oneString}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Fatalf("Where() mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("find index matches numeric strings and numbers", func(t *testing.T) {
		t.Parallel()

		got, err := FindIndex(records, "id", "1.0")
		require.NoError(t, err)
		require.Equal(t, 0, got)
	})

	t.Run("has matches numeric strings and numbers", func(t *testing.T) {
		t.Parallel()

		got, err := Has(records, "id", "1")
		require.NoError(t, err)
		require.True(t, got)
	})
}

func TestPredicateTruthiness(t *testing.T) {
	t.Parallel()

	truthyEmptyString := map[string]any{"name": "empty string", "value": ""}
	truthyZero := map[string]any{"name": "zero", "value": 0}
	truthyEmptySlice := map[string]any{"name": "empty slice", "value": []any{}}
	truthyEmptyMap := map[string]any{"name": "empty map", "value": map[string]any{}}
	falsyNil := map[string]any{"name": "nil", "value": nil}
	falsyFalse := map[string]any{"name": "false", "value": false}
	records := []map[string]any{
		truthyEmptyString,
		truthyZero,
		truthyEmptySlice,
		truthyEmptyMap,
		falsyNil,
		falsyFalse,
	}
	truthy := []any{truthyEmptyString, truthyZero, truthyEmptySlice, truthyEmptyMap}
	falsy := []any{falsyNil, falsyFalse}

	t.Run("where keeps every value except nil and false", func(t *testing.T) {
		t.Parallel()

		got, err := Where(records, "value")
		require.NoError(t, err)
		if diff := cmp.Diff(truthy, got); diff != "" {
			t.Fatalf("Where() mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("reject keeps nil and false", func(t *testing.T) {
		t.Parallel()

		got, err := Reject(records, "value")
		require.NoError(t, err)
		if diff := cmp.Diff(falsy, got); diff != "" {
			t.Fatalf("Reject() mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("has reports truthy value exists", func(t *testing.T) {
		t.Parallel()

		got, err := Has(records, "value")
		require.NoError(t, err)
		require.True(t, got)
	})
}

func TestReject(t *testing.T) {
	t.Parallel()

	products := []any{
		map[string]any{"name": "Shoes", "available": true},
		map[string]any{"name": "Shirt", "available": false},
		map[string]any{"name": "Pants", "available": true},
	}

	tests := []struct {
		name  string
		input any
		key   string
		value []any
		want  []any
	}{
		{
			name:  "Reject with value",
			input: products,
			key:   "available",
			value: []any{false},
			want: []any{
				map[string]any{"name": "Shoes", "available": true},
				map[string]any{"name": "Pants", "available": true},
			},
		},
		{
			name:  "Reject truthy (no value)",
			input: products,
			key:   "available",
			want: []any{
				map[string]any{"name": "Shirt", "available": false},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := Reject(tt.input, tt.key, tt.value...)
			require.NoError(t, err)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Fatalf("result mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFind(t *testing.T) {
	t.Parallel()

	products := []any{
		map[string]any{"handle": "shoes", "price": 50.0},
		map[string]any{"handle": "shirt", "price": 30.0},
		map[string]any{"handle": "pants", "price": 40.0},
	}

	tests := []struct {
		name    string
		input   any
		key     string
		value   any
		want    any
		wantErr error
	}{
		{
			name:  "Find existing",
			input: products,
			key:   "handle",
			value: "shirt",
			want:  map[string]any{"handle": "shirt", "price": 30.0},
		},
		{
			name:    "Find not existing",
			input:   products,
			key:     "handle",
			value:   "hat",
			wantErr: ErrNotFound,
		},
		{
			name:    "Find in empty slice",
			input:   []any{},
			key:     "handle",
			value:   "shoes",
			wantErr: ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := Find(tt.input, tt.key, tt.value)
			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				return
			}
			require.NoError(t, err)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Fatalf("result mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFindIndex(t *testing.T) {
	t.Parallel()

	products := []any{
		map[string]any{"handle": "shoes"},
		map[string]any{"handle": "shirt"},
		map[string]any{"handle": "pants"},
	}

	tests := []struct {
		name  string
		input any
		key   string
		value any
		want  int
	}{
		{
			name:  "Find index existing",
			input: products,
			key:   "handle",
			value: "shirt",
			want:  1,
		},
		{
			name:  "Find index not existing",
			input: products,
			key:   "handle",
			value: "hat",
			want:  -1,
		},
		{
			name:  "Find index first",
			input: products,
			key:   "handle",
			value: "shoes",
			want:  0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := FindIndex(tt.input, tt.key, tt.value)
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestHas(t *testing.T) {
	t.Parallel()

	products := []any{
		map[string]any{"name": "Shoes", "available": true},
		map[string]any{"name": "Shirt", "available": false},
	}

	tests := []struct {
		name  string
		input any
		key   string
		value []any
		want  bool
	}{
		{
			name:  "Has truthy property",
			input: products,
			key:   "available",
			want:  true,
		},
		{
			name:  "Has with value match",
			input: products,
			key:   "name",
			value: []any{"Shoes"},
			want:  true,
		},
		{
			name:  "Has with value no match",
			input: products,
			key:   "name",
			value: []any{"Hat"},
			want:  false,
		},
		{
			name:  "Has missing key",
			input: products,
			key:   "color",
			want:  false,
		},
		{
			name:  "Empty slice",
			input: []any{},
			key:   "name",
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := Has(tt.input, tt.key, tt.value...)
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

// Benchmark tests for array operations

func BenchmarkUnique(b *testing.B) {
	input := make([]any, 1000)
	for i := range 1000 {
		input[i] = i % 100 // 100 unique values
	}
	b.ResetTimer()
	for b.Loop() {
		_, _ = Unique(input)
	}
}

func BenchmarkShuffle(b *testing.B) {
	input := make([]any, 1000)
	for i := range 1000 {
		input[i] = i
	}
	b.ResetTimer()
	for b.Loop() {
		_, _ = Shuffle(input)
	}
}

func BenchmarkReverse(b *testing.B) {
	input := make([]any, 1000)
	for i := range 1000 {
		input[i] = i
	}
	b.ResetTimer()
	for b.Loop() {
		_, _ = Reverse(input)
	}
}

func BenchmarkJoin(b *testing.B) {
	input := make([]any, 100)
	for i := range 100 {
		input[i] = i
	}
	b.ResetTimer()
	for b.Loop() {
		_, _ = Join(input, ",")
	}
}

func BenchmarkSum(b *testing.B) {
	input := make([]any, 1000)
	for i := range 1000 {
		input[i] = float64(i)
	}
	b.ResetTimer()
	for b.Loop() {
		_, _ = Sum(input)
	}
}

func BenchmarkMax(b *testing.B) {
	input := make([]any, 1000)
	for i := range 1000 {
		input[i] = float64(i)
	}
	b.ResetTimer()
	for b.Loop() {
		_, _ = Max(input)
	}
}

func BenchmarkAverage(b *testing.B) {
	input := make([]any, 1000)
	for i := range 1000 {
		input[i] = float64(i)
	}
	b.ResetTimer()
	for b.Loop() {
		_, _ = Average(input)
	}
}
