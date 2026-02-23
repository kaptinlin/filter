package filter

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnique(t *testing.T) {
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
			got, err := Unique(tt.input)
			if tt.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.ElementsMatch(t, tt.want, got, "The returned slice should contain unique elements only")
			}
		})
	}
}

func TestJoin(t *testing.T) {
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
			want:      "",
			expectErr: true,
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
			got, err := Random(tt.input)
			if tt.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				// Verify the returned element is within the input slice.
				require.True(t, elementInSlice(got, tt.input), "The returned element should belong to the input slice")
			}
		})
	}
}

func TestReverse(t *testing.T) {
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
			got, err := Reverse(tt.input)
			if tt.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.want, got, "The reversed slice should match the expected output")
			}
		})
	}
}

func TestShuffle(t *testing.T) {
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
			original := deepcopy(tt.input)
			shuffled, err := Shuffle(tt.input)
			if tt.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.ElementsMatch(t, original, shuffled, "Shuffled slice should contain the same elements as the original")
			}
		})
	}
}

func TestSize(t *testing.T) {
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
			name:      "Unsupported input type (string)",
			input:     "hello",
			want:      0,
			expectErr: true,
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

func TestAverage(t *testing.T) {
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
			got, err := Map(tt.input, tt.key)
			if tt.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.want, got, "The extracted slice should match the expected output")
			}
		})
	}
}

func TestSort(t *testing.T) {
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
			got, err := Sort(tt.input, tt.key...)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.want, got)
			}
		})
	}
}

func TestSortNatural(t *testing.T) {
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
			got, err := SortNatural(tt.input, tt.key...)
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestCompact(t *testing.T) {
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
			got, err := Compact(tt.input, tt.key...)
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestConcat(t *testing.T) {
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
			got, err := Concat(tt.input, tt.other)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.want, got)
			}
		})
	}
}

func TestWhere(t *testing.T) {
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
			got, err := Where(tt.input, tt.key, tt.value...)
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestReject(t *testing.T) {
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
			got, err := Reject(tt.input, tt.key, tt.value...)
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestFind(t *testing.T) {
	products := []any{
		map[string]any{"handle": "shoes", "price": 50.0},
		map[string]any{"handle": "shirt", "price": 30.0},
		map[string]any{"handle": "pants", "price": 40.0},
	}

	tests := []struct {
		name  string
		input any
		key   string
		value any
		want  any
	}{
		{
			name:  "Find existing",
			input: products,
			key:   "handle",
			value: "shirt",
			want:  map[string]any{"handle": "shirt", "price": 30.0},
		},
		{
			name:  "Find not existing",
			input: products,
			key:   "handle",
			value: "hat",
			want:  nil,
		},
		{
			name:  "Find in empty slice",
			input: []any{},
			key:   "handle",
			value: "shoes",
			want:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Find(tt.input, tt.key, tt.value)
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestFindIndex(t *testing.T) {
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
			got, err := FindIndex(tt.input, tt.key, tt.value)
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestHas(t *testing.T) {
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
			got, err := Has(tt.input, tt.key, tt.value...)
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

// deepcopy creates a deep copy of the input slice to ensure the original slice is not modified.
func deepcopy(input any) any {
	inVal := reflect.ValueOf(input)
	if inVal.Kind() != reflect.Slice {
		return input
	}

	copyVal := reflect.MakeSlice(inVal.Type(), inVal.Len(), inVal.Cap())
	reflect.Copy(copyVal, inVal)
	return copyVal.Interface()
}

// elementInSlice checks if an element is present in the slice.
// This is a helper function for the Random test to verify the element belongs to the input slice.
func elementInSlice(element any, input any) bool {
	slice, _ := toSlice(input)
	for _, elem := range slice {
		if reflect.DeepEqual(elem, element) {
			return true
		}
	}
	return false
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
