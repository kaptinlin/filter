package filter

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnique(t *testing.T) {
	tests := []struct {
		name      string
		input     interface{}
		want      []interface{}
		expectErr bool
	}{
		{
			name:      "Unique elements in a slice of int",
			input:     []interface{}{1, 2, 3, 2, 1},
			want:      []interface{}{1, 2, 3},
			expectErr: false,
		},
		{
			name:      "Unique elements in a slice of string",
			input:     []interface{}{"apple", "banana", "apple", "cherry"},
			want:      []interface{}{"apple", "banana", "cherry"},
			expectErr: false,
		},
		{
			name:      "Unique elements with mixed types",
			input:     []interface{}{1, "apple", 1, "banana", "apple"},
			want:      []interface{}{1, "apple", "banana"},
			expectErr: false,
		},
		{
			name:      "Empty slice",
			input:     []interface{}{},
			want:      []interface{}{},
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
		input     interface{}
		separator string
		want      string
		expectErr bool
	}{
		{
			name:      "Join with comma separator",
			input:     []interface{}{"apple", "banana", "cherry"},
			separator: ",",
			want:      "apple,banana,cherry",
			expectErr: false,
		},
		{
			name:      "Join with mixed types",
			input:     []interface{}{1, "apple", 2.5},
			separator: ":",
			want:      "1:apple:2.5",
			expectErr: false,
		},
		{
			name:      "Join with empty separator",
			input:     []interface{}{"apple", "banana", "cherry"},
			separator: "",
			want:      "",
			expectErr: true,
		},
		{
			name:      "Empty slice",
			input:     []interface{}{},
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
		input     interface{}
		index     int
		want      interface{}
		expectErr bool
	}{
		{
			name:      "Valid index in slice of int",
			input:     []interface{}{1, 2, 3},
			index:     1,
			want:      2,
			expectErr: false,
		},
		{
			name:      "Valid index in slice of string",
			input:     []interface{}{"apple", "banana", "cherry"},
			index:     2,
			want:      "cherry",
			expectErr: false,
		},
		{
			name:      "Index out of range (too high)",
			input:     []interface{}{"apple", "banana", "cherry"},
			index:     5,
			want:      nil,
			expectErr: true,
		},
		{
			name:      "Index out of range (negative)",
			input:     []interface{}{"apple", "banana", "cherry"},
			index:     -1,
			want:      nil,
			expectErr: true,
		},
		{
			name:      "Empty slice",
			input:     []interface{}{},
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
		input     interface{}
		want      interface{}
		expectErr bool
	}{
		{
			name:      "Non-empty slice of int",
			input:     []interface{}{1, 2, 3},
			want:      3,
			expectErr: false,
		},
		{
			name:      "Non-empty slice of string",
			input:     []interface{}{"apple", "banana", "cherry"},
			want:      "cherry",
			expectErr: false,
		},
		{
			name:      "Non-empty mixed slice",
			input:     []interface{}{1, "banana", 3.5},
			want:      3.5,
			expectErr: false,
		},
		{
			name:      "Empty slice",
			input:     []interface{}{},
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
		input     interface{}
		expectErr bool
	}{
		{
			name:      "Non-empty slice of int",
			input:     []interface{}{1, 2, 3},
			expectErr: false,
		},
		{
			name:      "Non-empty slice of string",
			input:     []interface{}{"apple", "banana", "cherry"},
			expectErr: false,
		},
		{
			name:      "Empty slice",
			input:     []interface{}{},
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
		input     interface{}
		want      []interface{}
		expectErr bool
	}{
		{
			name:      "Reverse a slice of int",
			input:     []interface{}{1, 2, 3, 4},
			want:      []interface{}{4, 3, 2, 1},
			expectErr: false,
		},
		{
			name:      "Reverse a slice of string",
			input:     []interface{}{"apple", "banana", "cherry"},
			want:      []interface{}{"cherry", "banana", "apple"},
			expectErr: false,
		},
		{
			name:      "Reverse a mixed type slice",
			input:     []interface{}{1, "banana", 3.5},
			want:      []interface{}{3.5, "banana", 1},
			expectErr: false,
		},
		{
			name:      "Empty slice",
			input:     []interface{}{},
			want:      []interface{}{},
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
		input     interface{}
		expectErr bool
	}{
		{
			name:      "Shuffle a slice of int",
			input:     []interface{}{1, 2, 3, 4},
			expectErr: false,
		},
		{
			name:      "Shuffle a slice of string",
			input:     []interface{}{"apple", "banana", "cherry"},
			expectErr: false,
		},
		{
			name:      "Shuffle a mixed type slice",
			input:     []interface{}{1, "banana", 3.5},
			expectErr: false,
		},
		{
			name:      "Empty slice",
			input:     []interface{}{},
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
		input     interface{}
		want      int
		expectErr bool
	}{
		{
			name:      "Size of a non-empty slice of int",
			input:     []interface{}{1, 2, 3, 4},
			want:      4,
			expectErr: false,
		},
		{
			name:      "Size of a non-empty slice of string",
			input:     []interface{}{"apple", "banana", "cherry"},
			want:      3,
			expectErr: false,
		},
		{
			name:      "Size of a mixed type slice",
			input:     []interface{}{1, "banana", 3.5},
			want:      3,
			expectErr: false,
		},
		{
			name:      "Size of an empty slice",
			input:     []interface{}{},
			want:      0,
			expectErr: false,
		},
		{
			name:      "Unsupported input type",
			input:     "not-a-slice",
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
		input     interface{}
		want      float64
		expectErr bool
	}{
		{
			name:      "Non-empty slice of float64",
			input:     []interface{}{1.2, 3.4, 5.6, 4.5},
			want:      5.6,
			expectErr: false,
		},
		{
			name:      "Slice with negative values",
			input:     []interface{}{-1.1, -3.3, -2.2},
			want:      -1.1,
			expectErr: false,
		},
		{
			name:      "Slice with a single element",
			input:     []interface{}{42.0},
			want:      42.0,
			expectErr: false,
		},
		{
			name:      "Empty slice",
			input:     []interface{}{},
			expectErr: true,
		},
		{
			name:      "Unsupported input type (not a slice)",
			input:     "not-a-slice",
			expectErr: true,
		},
		{
			name:      "Unsupported element type in slice",
			input:     []interface{}{1.1, "not-a-float", 3.3},
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
		input     interface{}
		want      float64
		expectErr bool
	}{
		{
			name:      "Non-empty slice of float64 with positive values",
			input:     []interface{}{2.2, 1.1, 3.3, 4.4},
			want:      1.1,
			expectErr: false,
		},
		{
			name:      "Slice with negative and positive values",
			input:     []interface{}{-1.2, 0.0, -2.2, 2.2},
			want:      -2.2,
			expectErr: false,
		},
		{
			name:      "Slice with a single element",
			input:     []interface{}{42.0},
			want:      42.0,
			expectErr: false,
		},
		{
			name:      "Empty slice",
			input:     []interface{}{},
			expectErr: true,
		},
		{
			name:      "Unsupported input type (not a slice)",
			input:     "not-a-slice",
			expectErr: true,
		},
		{
			name:      "Unsupported element type in slice",
			input:     []interface{}{1.1, "not-a-float", 3.3},
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
		input     interface{}
		want      float64
		expectErr bool
	}{
		{
			name:      "Slice of positive float64",
			input:     []interface{}{1.1, 2.2, 3.3},
			want:      6.6,
			expectErr: false,
		},
		{
			name:      "Slice of negative float64",
			input:     []interface{}{-1.1, -2.2, -3.3},
			want:      -6.6,
			expectErr: false,
		},
		{
			name:      "Mixed positive and negative float64",
			input:     []interface{}{-1.1, 2.2, -3.3, 4.4},
			want:      2.2,
			expectErr: false,
		},
		{
			name:      "Slice with a single element",
			input:     []interface{}{42.0},
			want:      42.0,
			expectErr: false,
		},
		{
			name:      "Empty slice",
			input:     []interface{}{},
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
			input:     []interface{}{1.1, "not-a-float", 3.3},
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
		input     interface{}
		want      float64
		expectErr bool
	}{
		{
			name:      "Average of positive float64",
			input:     []interface{}{1.0, 2.0, 3.0},
			want:      2.0,
			expectErr: false,
		},
		{
			name:      "Average of negative float64",
			input:     []interface{}{-1.0, -2.0, -3.0},
			want:      -2.0,
			expectErr: false,
		},
		{
			name:      "Mixed positive and negative float64",
			input:     []interface{}{-2.0, 2.0},
			want:      0.0,
			expectErr: false,
		},
		{
			name:      "Slice with a single element",
			input:     []interface{}{42.0},
			want:      42.0,
			expectErr: false,
		},
		{
			name:      "Empty slice",
			input:     []interface{}{},
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
			input:     []interface{}{1.1, "not-a-float", 3.3},
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
		input     interface{}
		key       string
		want      []interface{}
		expectErr bool
	}{
		{
			name: "All maps contain the key",
			input: []interface{}{
				map[interface{}]interface{}{"key1": "value1"},
				map[interface{}]interface{}{"key1": "value2"},
			},
			key:  "key1",
			want: []interface{}{"value1", "value2"},
		},
		{
			name: "Some maps do not contain the key",
			input: []interface{}{
				map[interface{}]interface{}{"key1": "value1"},
				map[interface{}]interface{}{"key2": "value2"},
			},
			key:  "key1",
			want: []interface{}{"value1", nil},
		},
		{
			name: "Key not present in any map",
			input: []interface{}{
				map[interface{}]interface{}{"key2": "value1"},
				map[interface{}]interface{}{"key3": "value2"},
			},
			key:  "key1",
			want: []interface{}{nil, nil},
		},
		{
			name:  "Empty slice",
			input: []interface{}{},
			key:   "key1",
			want:  []interface{}{},
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
			input: []interface{}{
				map[interface{}]interface{}{"parent": map[interface{}]interface{}{"child": "value1"}},
				map[interface{}]interface{}{"parent": map[interface{}]interface{}{"child": "value2"}},
			},
			key:  "parent.child",
			want: []interface{}{"value1", "value2"},
		},
		{
			name: "Multiple levels of nesting",
			input: []interface{}{
				map[interface{}]interface{}{"level1": map[interface{}]interface{}{"level2": map[interface{}]interface{}{"level3": "deepValue1"}}},
				map[interface{}]interface{}{"level1": map[interface{}]interface{}{"level2": map[interface{}]interface{}{"level3": "deepValue2"}}},
			},
			key:  "level1.level2.level3",
			want: []interface{}{"deepValue1", "deepValue2"},
		},
		{
			name: "Mixed nesting and types",
			input: []interface{}{
				map[interface{}]interface{}{"parent": []interface{}{"value1", map[interface{}]interface{}{"child": "nestedValue1"}}},
				map[interface{}]interface{}{"parent": map[interface{}]interface{}{"child": "nestedValue2"}},
			},
			key:  "parent.1.child",
			want: []interface{}{"nestedValue1", nil},
		},
		{
			name: "Value is a slice",
			input: []interface{}{
				map[interface{}]interface{}{"parent": []interface{}{"value1", "value2"}},
				map[interface{}]interface{}{"parent": []interface{}{"value3", "value4"}},
			},
			key:  "parent",
			want: []interface{}{[]interface{}{"value1", "value2"}, []interface{}{"value3", "value4"}},
		},
		{
			name: "Value is a map",
			input: []interface{}{
				map[interface{}]interface{}{"parent": map[interface{}]interface{}{"child": "value"}},
				map[interface{}]interface{}{"parent": map[interface{}]interface{}{"child": "value"}},
			},
			key:  "parent",
			want: []interface{}{map[interface{}]interface{}{"child": "value"}, map[interface{}]interface{}{"child": "value"}},
		},
		{
			name: "Intermediate key leads to a non-map and non-slice element",
			input: []interface{}{
				map[interface{}]interface{}{"parent": "notAMapOrSlice"},
			},
			key:  "parent.child",
			want: []interface{}{nil},
		},
		{
			name: "Key missing at intermediate level",
			input: []interface{}{
				map[interface{}]interface{}{"level1": map[interface{}]interface{}{}}, // Empty map at level1
			},
			key:  "level1.level2.missing",
			want: []interface{}{nil},
		},
		{
			name: "Index out of range in slice",
			input: []interface{}{
				map[interface{}]interface{}{"list": []interface{}{"value1", "value2"}},
			},
			key:  "list.2",
			want: []interface{}{nil}, // Index 2 is out of range
		},
		{
			name: "Incorrect index format in key",
			input: []interface{}{
				map[interface{}]interface{}{"list": []interface{}{"value1", "value2"}},
			},
			key:  "list.two",
			want: []interface{}{nil},
		},
		{
			name: "Nested key leading to array index",
			input: []interface{}{
				map[interface{}]interface{}{"parent": []interface{}{"value1", "value2"}},
			},
			key:  "parent.1",
			want: []interface{}{"value2"},
		},
		{
			name: "Deep nesting with arrays and maps",
			input: []interface{}{
				map[interface{}]interface{}{"level1": []interface{}{map[interface{}]interface{}{"level2": []interface{}{"value1"}}}},
			},
			key:  "level1.0.level2.0",
			want: []interface{}{"value1"},
		},
		{
			name: "Nesting with a non-existing final key",
			input: []interface{}{
				map[interface{}]interface{}{"exists": map[interface{}]interface{}{"doesNotExist": "value"}},
			},
			key:  "exists.notHere",
			want: []interface{}{nil},
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

// deepcopy creates a deep copy of the input slice to ensure the original slice is not modified.
func deepcopy(input interface{}) interface{} {
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
func elementInSlice(element interface{}, input interface{}) bool {
	slice, _ := toSlice(input)
	for _, elem := range slice {
		if reflect.DeepEqual(elem, element) {
			return true
		}
	}
	return false
}
