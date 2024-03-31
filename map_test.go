package filter

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExtract(t *testing.T) {
	tests := []struct {
		name        string
		input       interface{}
		key         string
		expected    interface{}
		expectError bool
		errorType   error
	}{
		{
			name:     "extract from single level map",
			input:    map[string]interface{}{"key1": "value1", "key2": 2},
			key:      "key1",
			expected: "value1",
		},
		{
			name:     "extract from nested map",
			input:    map[string]interface{}{"level1": map[string]interface{}{"level2": "value2"}},
			key:      "level1.level2",
			expected: "value2",
		},
		{
			name:     "extract from nested map with slice",
			input:    map[string]interface{}{"level1": []interface{}{1, 2, map[string]interface{}{"level3": "value3"}}},
			key:      "level1.2.level3",
			expected: "value3",
		},
		{
			name:        "key not found",
			input:       map[string]interface{}{"key1": "value1"},
			key:         "key2",
			expected:    nil,
			expectError: true,
			errorType:   ErrKeyNotFound,
		},
		{
			name:        "index out of range in slice",
			input:       []interface{}{"value1", "value2"},
			key:         "2",
			expected:    nil,
			expectError: true,
			errorType:   ErrIndexOutOfRange,
		},
		{
			name:        "invalid key type",
			input:       map[string]interface{}{"key1": "value1"},
			key:         "key1.level2",
			expected:    nil,
			expectError: true,
			errorType:   ErrInvalidKeyType,
		},
		{
			name:        "nil input",
			input:       nil,
			key:         "key1",
			expected:    nil,
			expectError: true,
			errorType:   ErrUnsupportedType,
		},
		{
			name:     "extract int from slice by index",
			input:    []interface{}{1, 2, 3},
			key:      "1",
			expected: 2,
		},
		{
			name:     "extract from deeply nested structure",
			input:    map[string]interface{}{"a": map[string]interface{}{"b": []interface{}{map[string]interface{}{"c": "deep value"}}}},
			key:      "a.b.0.c",
			expected: "deep value",
		},
		{
			name:     "empty key on non-empty map",
			input:    map[string]interface{}{"emptyKey": ""},
			key:      "emptyKey",
			expected: "",
		},
		{
			name:        "empty key on non-empty slice",
			input:       []interface{}{0, 1, 2},
			key:         "",
			expected:    nil,
			expectError: true,
			errorType:   ErrKeyNotFound,
		},
		{
			name:     "numeric key on map",
			input:    map[string]interface{}{"1": "numeric key"},
			key:      "1",
			expected: "numeric key",
		},
		{
			name:        "nonexistent nested map key",
			input:       map[string]interface{}{"level1": map[string]interface{}{"level2": map[string]interface{}{}}},
			key:         "level1.level2.nonexistent",
			expected:    nil,
			expectError: true,
			errorType:   ErrKeyNotFound,
		},
		{
			name:     "complex nested structure with arrays",
			input:    map[string]interface{}{"array": []interface{}{[]interface{}{"nested array value"}}},
			key:      "array.0.0",
			expected: "nested array value",
		},
		{
			name:        "attempt to index into integer",
			input:       map[string]interface{}{"int": 42},
			key:         "int.0",
			expected:    nil,
			expectError: true,
			errorType:   ErrInvalidKeyType,
		},
		{
			name:     "direct access to array item",
			input:    []interface{}{"first", "second"},
			key:      "0",
			expected: "first",
		},
		{
			name:     "slice with mixed types",
			input:    []interface{}{42, "string", map[string]interface{}{"key": "value"}},
			key:      "2.key",
			expected: "value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Extract(tt.input, tt.key)
			if tt.expectError {
				require.Error(t, err)
				if tt.errorType != nil {
					require.ErrorIs(t, err, tt.errorType)
				}
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expected, got, "The expected and actual value should match.")
			}
		})
	}
}
