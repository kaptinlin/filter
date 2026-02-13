package filter

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestExtractMap tests the Extract function with map data structures
func TestExtractMap(t *testing.T) {
	tests := []struct {
		name        string
		input       interface{}
		key         string
		expected    interface{}
		expectError bool
		errorType   error
	}{
		{
			name:     "single_level_map",
			input:    map[string]interface{}{"key1": "value1", "key2": 2},
			key:      "key1",
			expected: "value1",
		},
		{
			name:     "nested_map",
			input:    map[string]interface{}{"level1": map[string]interface{}{"level2": "value2"}},
			key:      "level1.level2",
			expected: "value2",
		},
		{
			name:     "nested_map_with_slice",
			input:    map[string]interface{}{"level1": []interface{}{1, 2, map[string]interface{}{"level3": "value3"}}},
			key:      "level1.2.level3",
			expected: "value3",
		},
		{
			name:        "key_not_found",
			input:       map[string]interface{}{"key1": "value1"},
			key:         "key2",
			expectError: true,
			errorType:   ErrKeyNotFound,
		},
		{
			name:        "invalid_nesting_key",
			input:       map[string]interface{}{"key1": "value1"},
			key:         "key1.level2",
			expectError: true,
			errorType:   ErrInvalidKeyType,
		},
		{
			name:     "empty_key_on_non_empty_map",
			input:    map[string]interface{}{"emptyKey": ""},
			key:      "emptyKey",
			expected: "",
		},
		{
			name:     "numeric_key_on_map",
			input:    map[string]interface{}{"1": "numeric key"},
			key:      "1",
			expected: "numeric key",
		},
		{
			name:        "nonexistent_nested_map_key",
			input:       map[string]interface{}{"level1": map[string]interface{}{"level2": map[string]interface{}{}}},
			key:         "level1.level2.nonexistent",
			expectError: true,
			errorType:   ErrKeyNotFound,
		},
		{
			name:     "complex_nested_structure_with_arrays",
			input:    map[string]interface{}{"array": []interface{}{[]interface{}{"nested array value"}}},
			key:      "array.0.0",
			expected: "nested array value",
		},
		{
			name:        "attempt_to_index_into_integer",
			input:       map[string]interface{}{"int": 42},
			key:         "int.0",
			expectError: true,
			errorType:   ErrInvalidKeyType,
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
				require.Equal(t, tt.expected, got)
			}
		})
	}
}

// TestExtractSlice tests the Extract function with slice/array data structures
func TestExtractSlice(t *testing.T) {
	tests := []struct {
		name        string
		input       interface{}
		key         string
		expected    interface{}
		expectError bool
		errorType   error
	}{
		{
			name:     "direct_access_to_array_item",
			input:    []interface{}{"first", "second"},
			key:      "0",
			expected: "first",
		},
		{
			name:     "extract_int_from_slice_by_index",
			input:    []interface{}{1, 2, 3},
			key:      "1",
			expected: 2,
		},
		{
			name:     "slice_with_mixed_types",
			input:    []interface{}{42, "string", map[string]interface{}{"key": "value"}},
			key:      "2.key",
			expected: "value",
		},
		{
			name:        "index_out_of_range",
			input:       []interface{}{"value1", "value2"},
			key:         "2",
			expectError: true,
			errorType:   ErrIndexOutOfRange,
		},
		{
			name:        "empty_key_on_slice",
			input:       []interface{}{0, 1, 2},
			key:         "",
			expectError: true,
			errorType:   ErrKeyNotFound,
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
				require.Equal(t, tt.expected, got)
			}
		})
	}
}

// TestExtractEdgeCases tests edge cases for the Extract function
func TestExtractEdgeCases(t *testing.T) {
	tests := []struct {
		name        string
		input       interface{}
		key         string
		expectError bool
		errorType   error
	}{
		{
			name:        "nil_input",
			input:       nil,
			key:         "key1",
			expectError: true,
			errorType:   ErrUnsupportedType,
		},
		{
			name:        "empty_key",
			input:       map[string]interface{}{"key": "value"},
			key:         "",
			expectError: true,
			errorType:   ErrKeyNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Extract(tt.input, tt.key)
			require.Error(t, err)
			if tt.errorType != nil {
				require.ErrorIs(t, err, tt.errorType)
			}
		})
	}
}

// TestExtractSimpleStruct tests basic struct field access
func TestExtractSimpleStruct(t *testing.T) {
	// Define a simple location structure
	type Location struct {
		Road     string    `json:"road"`
		District string    `json:"district"`
		PostCode int       `json:"post_code"`
		Created  time.Time `json:"created"`
	}

	now := time.Now()
	location := Location{
		Road:     "Broadway",
		District: "Manhattan",
		PostCode: 10001,
		Created:  now,
	}

	tests := []struct {
		name     string
		key      string
		expected interface{}
	}{
		{
			name:     "access_string_field",
			key:      "road",
			expected: "Broadway",
		},
		{
			name:     "access_another_string_field",
			key:      "district",
			expected: "Manhattan",
		},
		{
			name:     "access_integer_field",
			key:      "post_code",
			expected: 10001,
		},
		{
			name:     "access_time_field",
			key:      "created",
			expected: now,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Extract(location, tt.key)
			require.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// TestExtractNestedStruct tests nested struct field access
func TestExtractNestedStruct(t *testing.T) {
	type Address struct {
		Street  string `json:"street"`
		City    string `json:"city"`
		ZipCode int    `json:"zip_code"`
	}

	type Person struct {
		Name    string   `json:"name"`
		Age     int      `json:"age"`
		Address Address  `json:"address"`
		Tags    []string `json:"tags"`
	}

	person := Person{
		Name: "John Doe",
		Age:  30,
		Address: Address{
			Street:  "Main St",
			City:    "New York",
			ZipCode: 10001,
		},
		Tags: []string{"developer", "golang"},
	}

	tests := []struct {
		name     string
		key      string
		expected interface{}
	}{
		{
			name:     "access_basic_field",
			key:      "name",
			expected: "John Doe",
		},
		{
			name:     "access_numeric_field",
			key:      "age",
			expected: 30,
		},
		{
			name:     "access_nested_struct_field",
			key:      "address.street",
			expected: "Main St",
		},
		{
			name:     "access_deeply_nested_struct_field",
			key:      "address.city",
			expected: "New York",
		},
		{
			name:     "access_slice_element",
			key:      "tags.0",
			expected: "developer",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Extract(person, tt.key)
			require.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// TestExtractStructWithMap tests structures containing maps
func TestExtractStructWithMap(t *testing.T) {
	type Location struct {
		Name     string `json:"name"`
		PostCode int    `json:"post_code"`
	}

	type Organization struct {
		Name      string              `json:"name"`
		Locations map[string]Location `json:"locations"`
		Metadata  map[string]string   `json:"metadata"`
		Tags      map[string][]string `json:"tags"`
	}

	org := Organization{
		Name: "Acme Corp",
		Locations: map[string]Location{
			"hq": {
				Name:     "Headquarters",
				PostCode: 10001,
			},
			"branch": {
				Name:     "Branch Office",
				PostCode: 20001,
			},
		},
		Metadata: map[string]string{
			"founded": "1985",
			"revenue": "10M",
		},
		Tags: map[string][]string{
			"categories": {"tech", "software"},
			"markets":    {"US", "EU", "Asia"},
		},
	}

	tests := []struct {
		name     string
		key      string
		expected interface{}
	}{
		{
			name:     "access_basic_field",
			key:      "name",
			expected: "Acme Corp",
		},
		{
			name:     "access_map_with_struct_value",
			key:      "locations.hq.name",
			expected: "Headquarters",
		},
		{
			name:     "access_different_map_entry",
			key:      "locations.branch.post_code",
			expected: 20001,
		},
		{
			name:     "access_string_map_value",
			key:      "metadata.founded",
			expected: "1985",
		},
		{
			name:     "access_map_with_slice_value",
			key:      "tags.categories.1",
			expected: "software",
		},
		{
			name:     "access_another_slice_in_map",
			key:      "tags.markets.2",
			expected: "Asia",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Extract(org, tt.key)
			require.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// TestExtractStructWithPointers tests extracting from structures containing pointers
func TestExtractStructWithPointers(t *testing.T) {
	type Department struct {
		Name     string `json:"name"`
		Building string `json:"building"`
	}

	type Employee struct {
		Name       string      `json:"name"`
		Department *Department `json:"department"`
		Manager    *Employee   `json:"manager"`
	}

	engineering := &Department{
		Name:     "Engineering",
		Building: "Building A",
	}

	manager := &Employee{
		Name:       "Jane Smith",
		Department: engineering,
		Manager:    nil,
	}

	employee := Employee{
		Name:       "John Doe",
		Department: engineering,
		Manager:    manager,
	}

	tests := []struct {
		name     string
		key      string
		expected interface{}
	}{
		{
			name:     "access_basic_field",
			key:      "name",
			expected: "John Doe",
		},
		{
			name:     "access_pointer_to_struct",
			key:      "department.name",
			expected: "Engineering",
		},
		{
			name:     "access_nested_field_through_pointer",
			key:      "department.building",
			expected: "Building A",
		},
		{
			name:     "access_pointer_to_struct_with_pointer",
			key:      "manager.name",
			expected: "Jane Smith",
		},
		{
			name:     "access_nested_pointer_chain",
			key:      "manager.department.name",
			expected: "Engineering",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Extract(employee, tt.key)
			require.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// TestExtractComplexSelfReferentialStructures tests extraction from complex self-referential and deeply nested structures
func TestExtractComplexSelfReferentialStructures(t *testing.T) {
	// Define a simpler self-referential structure
	type Location struct {
		Street string `json:"street"`
		City   string `json:"city"`
	}

	type Department struct {
		Name     string            `json:"name"`
		Location Location          `json:"location"`
		Staff    map[string]string `json:"staff"` // Simplified to string values
	}

	type Organization struct {
		Title        string                   `json:"title"`
		Departments  []Department             `json:"departments"`
		Branches     map[string]Location      `json:"branches"`
		Parent       *Organization            `json:"parent"`
		Subsidiaries []*Organization          `json:"subsidiaries"`
		Partners     map[string]*Organization `json:"partners"`
	}

	// Create test data with simplified self-referential structures
	headquarter := &Organization{
		Title: "Headquarters",
		Departments: []Department{
			{
				Name: "Engineering",
				Location: Location{
					Street: "Main St",
					City:   "San Francisco",
				},
				Staff: map[string]string{
					"dev1": "John Doe",
					"dev2": "Jane Smith",
				},
			},
			{
				Name: "Marketing",
				Location: Location{
					Street: "Broadway",
					City:   "New York",
				},
				Staff: map[string]string{
					"marketing1": "Alice Johnson",
				},
			},
		},
		Branches: map[string]Location{
			"west": {
				Street: "Tech Drive",
				City:   "Seattle",
			},
			"east": {
				Street: "Innovation Blvd",
				City:   "Boston",
			},
		},
	}

	// Create subsidiary with reference to parent
	subsidiary := &Organization{
		Title:  "Subsidiary",
		Parent: headquarter,
		Departments: []Department{
			{
				Name: "Research",
				Location: Location{
					Street: "Science Park",
					City:   "Austin",
				},
				Staff: map[string]string{
					"researcher1": "Bob Williams",
				},
			},
		},
	}

	// Add subsidiary to parent
	headquarter.Subsidiaries = []*Organization{subsidiary}

	// Create partner with circular reference
	partner := &Organization{
		Title: "Partner Org",
		Partners: map[string]*Organization{
			"main": headquarter,
		},
	}

	// Add partner to headquarters
	headquarter.Partners = map[string]*Organization{
		"tech": partner,
	}

	// Define the test cases
	tests := []struct {
		name     string
		key      string
		expected interface{}
	}{
		{
			name:     "access_top_level_field",
			key:      "title",
			expected: "Headquarters",
		},
		{
			name:     "access_department_name",
			key:      "departments.0.name",
			expected: "Engineering",
		},
		{
			name:     "access_department_location",
			key:      "departments.0.location.city",
			expected: "San Francisco",
		},
		{
			name:     "access_staff_map",
			key:      "departments.0.staff.dev1",
			expected: "John Doe",
		},
		{
			name:     "access_branch_location",
			key:      "branches.west.city",
			expected: "Seattle",
		},
		{
			name:     "access_subsidiary",
			key:      "subsidiaries.0.title",
			expected: "Subsidiary",
		},
		{
			name:     "access_parent_reference",
			key:      "subsidiaries.0.parent.title",
			expected: "Headquarters",
		},
		{
			name:     "access_circular_reference",
			key:      "partners.tech.partners.main.title",
			expected: "Headquarters",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Extract(headquarter, tt.key)
			require.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}

	// Test error cases
	errorTests := []struct {
		name      string
		key       string
		errorType error
	}{
		{
			name:      "non_existent_field",
			key:       "non_existent_field",
			errorType: ErrKeyNotFound,
		},
		{
			name:      "index_out_of_range",
			key:       "departments.5.name",
			errorType: ErrIndexOutOfRange,
		},
		{
			name:      "invalid_path_through_primitive",
			key:       "title.something",
			errorType: ErrInvalidKeyType,
		},
		{
			name:      "key_not_found_in_map",
			key:       "branches.north.city",
			errorType: ErrKeyNotFound,
		},
	}

	for _, tt := range errorTests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Extract(headquarter, tt.key)
			require.Error(t, err)
			require.ErrorIs(t, err, tt.errorType)
		})
	}
}

// TestExtractUltraComplexStructures tests the Extract function on extremely complex nested data structures
func TestExtractUltraComplexStructures(t *testing.T) {
	// Define structure with ultra-complex nesting
	type Node struct {
		ID    int    `json:"id"`
		Label string `json:"label"`
	}

	type DeepStructure struct {
		Name      string                        `json:"name"`
		Nodes     map[string]*Node              `json:"nodes"`
		Matrix    [2][2]*Node                   `json:"matrix"`
		NestedMap map[string]map[string][]*Node `json:"nested_map"`
		GridMap   [2][2]map[string]*Node        `json:"grid_map"`
	}

	// Create test data with manageable nesting depth (max 4 levels)
	deep := DeepStructure{
		Name: "SimplifiedComplex",
		Nodes: map[string]*Node{
			"node1": {ID: 1, Label: "First Node"},
			"node2": {ID: 2, Label: "Second Node"},
			"node3": {ID: 3, Label: "Third Node"},
		},
		Matrix: [2][2]*Node{
			{
				{ID: 11, Label: "Matrix 0,0"},
				{ID: 12, Label: "Matrix 0,1"},
			},
			{
				{ID: 21, Label: "Matrix 1,0"},
				{ID: 22, Label: "Matrix 1,1"},
			},
		},
		NestedMap: map[string]map[string][]*Node{
			"region1": {
				"area1": {
					{ID: 101, Label: "Region1-Area1-Node1"},
					{ID: 102, Label: "Region1-Area1-Node2"},
				},
				"area2": {
					{ID: 103, Label: "Region1-Area2-Node1"},
				},
			},
			"region2": {
				"area3": {
					{ID: 201, Label: "Region2-Area3-Node1"},
					{ID: 202, Label: "Region2-Area3-Node2"},
				},
			},
		},
		GridMap: [2][2]map[string]*Node{
			{
				{"key1": {ID: 301, Label: "Grid 0,0 Key1"}},
				{"key2": {ID: 302, Label: "Grid 0,1 Key2"}},
			},
			{
				{"key3": {ID: 401, Label: "Grid 1,0 Key3"}},
				{"key4": {ID: 402, Label: "Grid 1,1 Key4"}},
			},
		},
	}

	// Define test cases for accessing nested structures
	tests := []struct {
		name     string
		key      string
		expected interface{}
	}{
		{
			name:     "access_name",
			key:      "name",
			expected: "SimplifiedComplex",
		},
		{
			name:     "access_map_pointer",
			key:      "nodes.node1.label",
			expected: "First Node",
		},
		{
			name:     "access_map_pointer_id",
			key:      "nodes.node3.id",
			expected: 3,
		},
		{
			name:     "access_matrix_element",
			key:      "matrix.0.1.label",
			expected: "Matrix 0,1",
		},
		{
			name:     "access_matrix_element_id",
			key:      "matrix.1.0.id",
			expected: 21,
		},
		{
			name:     "access_nested_map_slice",
			key:      "nested_map.region1.area1.0.label",
			expected: "Region1-Area1-Node1",
		},
		{
			name:     "access_nested_map_slice_another",
			key:      "nested_map.region2.area3.1.id",
			expected: 202,
		},
		{
			name:     "access_grid_map_element",
			key:      "grid_map.0.0.key1.label",
			expected: "Grid 0,0 Key1",
		},
		{
			name:     "access_grid_map_element_another",
			key:      "grid_map.1.1.key4.id",
			expected: 402,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Extract(deep, tt.key)
			require.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}

	// Test error cases
	errorTests := []struct {
		name      string
		key       string
		errorType error
	}{
		{
			name:      "non_existent_field",
			key:       "nonexistent",
			errorType: ErrKeyNotFound,
		},
		{
			name:      "index_out_of_range",
			key:       "matrix.3.0.id",
			errorType: ErrIndexOutOfRange,
		},
		{
			name:      "non_existent_key_in_map",
			key:       "nodes.nonexistent.id",
			errorType: ErrKeyNotFound,
		},
		{
			name:      "invalid_path_through_primitive",
			key:       "nodes.node1.id.something",
			errorType: ErrInvalidKeyType,
		},
	}

	for _, tt := range errorTests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Extract(deep, tt.key)
			require.Error(t, err)
			require.ErrorIs(t, err, tt.errorType)
		})
	}
}

// TestExtractErrorHandling tests specific error handling scenarios to improve coverage
func TestExtractErrorHandling(t *testing.T) {
	type NestedStruct struct {
		Value int `json:"value"`
	}

	type TestStruct struct {
		IntField     int            `json:"int_field"`
		StringField  string         `json:"string_field"`
		NestedField  NestedStruct   `json:"nested_field"`
		SliceField   []int          `json:"slice_field"`
		MapField     map[string]int `json:"map_field"`
		PointerField *NestedStruct  `json:"pointer_field"`
	}

	testData := TestStruct{
		IntField:     42,
		StringField:  "hello",
		NestedField:  NestedStruct{Value: 100},
		SliceField:   []int{1, 2, 3},
		MapField:     map[string]int{"key1": 10, "key2": 20},
		PointerField: &NestedStruct{Value: 200},
	}

	t.Run("ErrInvalidIndex", func(t *testing.T) {
		// Test invalid array index (non-numeric)
		_, err := Extract(testData, "slice_field.invalid")
		require.Error(t, err)
		require.ErrorIs(t, err, ErrInvalidKeyType)
	})

	t.Run("ErrInvalidPathStep", func(t *testing.T) {
		// This error would be triggered by jsonpointer for invalid path steps
		// Since jsonpointer handles most path validation, we test edge cases
		_, err := Extract(testData, "int_field.nested.deep")
		require.Error(t, err)
		require.ErrorIs(t, err, ErrInvalidKeyType)
	})

	t.Run("ErrNilPointer", func(t *testing.T) {
		// Test navigation through nil pointer
		testDataWithNil := TestStruct{
			IntField:     42,
			StringField:  "hello",
			PointerField: nil, // nil pointer
		}
		_, err := Extract(testDataWithNil, "pointer_field.value")
		require.Error(t, err)
		require.ErrorIs(t, err, ErrInvalidKeyType)
	})

	t.Run("ErrKeyNotFound_Map", func(t *testing.T) {
		// Test map key not found
		_, err := Extract(testData, "map_field.nonexistent")
		require.Error(t, err)
		require.ErrorIs(t, err, ErrKeyNotFound)
	})

	t.Run("ErrFieldNotFound_Struct", func(t *testing.T) {
		// Test struct field not found
		_, err := Extract(testData, "nonexistent_field")
		require.Error(t, err)
		require.ErrorIs(t, err, ErrKeyNotFound)
	})

	t.Run("ErrIndexOutOfBounds_Slice", func(t *testing.T) {
		// Test slice index out of bounds
		_, err := Extract(testData, "slice_field.10")
		require.Error(t, err)
		require.ErrorIs(t, err, ErrIndexOutOfRange)
	})

	t.Run("ErrIndexOutOfBounds_NegativeIndex", func(t *testing.T) {
		// Test negative array index - jsonpointer treats this as invalid index, not out of bounds
		_, err := Extract(testData, "slice_field.-1")
		require.Error(t, err)
		require.ErrorIs(t, err, ErrInvalidKeyType)
	})

	t.Run("Complex_Error_Path", func(t *testing.T) {
		// Test complex error path that triggers multiple error checks
		_, err := Extract(testData, "nested_field.value.invalid.deep.path")
		require.Error(t, err)
		require.ErrorIs(t, err, ErrInvalidKeyType)
	})
}

// TestExtractPointerHandling tests comprehensive pointer handling scenarios
func TestExtractPointerHandling(t *testing.T) {
	type Level3 struct {
		Value string `json:"value"`
	}

	type Level2 struct {
		Level3Ptr *Level3 `json:"level3_ptr"`
		Level3Val Level3  `json:"level3_val"`
	}

	type Level1 struct {
		Level2Ptr *Level2 `json:"level2_ptr"`
		Level2Val Level2  `json:"level2_val"`
	}

	type RootStruct struct {
		Level1Ptr    *Level1            `json:"level1_ptr"`
		Level1Val    Level1             `json:"level1_val"`
		PointerSlice []*Level3          `json:"pointer_slice"`
		PointerMap   map[string]*Level3 `json:"pointer_map"`
	}

	// Create test data with various pointer scenarios
	testData := RootStruct{
		Level1Ptr: &Level1{
			Level2Ptr: &Level2{
				Level3Ptr: &Level3{Value: "deep_pointer_value"},
				Level3Val: Level3{Value: "deep_value_through_pointer"},
			},
			Level2Val: Level2{
				Level3Ptr: &Level3{Value: "mixed_pointer_value"},
				Level3Val: Level3{Value: "mixed_value"},
			},
		},
		Level1Val: Level1{
			Level2Ptr: &Level2{
				Level3Ptr: &Level3{Value: "value_pointer_value"},
				Level3Val: Level3{Value: "all_values"},
			},
			Level2Val: Level2{
				Level3Ptr: nil, // Test nil pointer handling
				Level3Val: Level3{Value: "partial_nil"},
			},
		},
		PointerSlice: []*Level3{
			{Value: "slice_ptr_0"},
			{Value: "slice_ptr_1"},
			nil, // nil pointer in slice
		},
		PointerMap: map[string]*Level3{
			"key1": {Value: "map_ptr_value1"},
			"key2": {Value: "map_ptr_value2"},
			"key3": nil, // nil pointer in map
		},
	}

	tests := []struct {
		name     string
		key      string
		expected interface{}
	}{
		{
			name:     "deep_pointer_chain",
			key:      "level1_ptr.level2_ptr.level3_ptr.value",
			expected: "deep_pointer_value",
		},
		{
			name:     "mixed_pointer_value_access",
			key:      "level1_ptr.level2_ptr.level3_val.value",
			expected: "deep_value_through_pointer",
		},
		{
			name:     "value_with_pointer_access",
			key:      "level1_val.level2_ptr.level3_ptr.value",
			expected: "value_pointer_value",
		},
		{
			name:     "all_values_access",
			key:      "level1_val.level2_ptr.level3_val.value",
			expected: "all_values",
		},
		{
			name:     "pointer_slice_access",
			key:      "pointer_slice.0.value",
			expected: "slice_ptr_0",
		},
		{
			name:     "pointer_slice_second_element",
			key:      "pointer_slice.1.value",
			expected: "slice_ptr_1",
		},
		{
			name:     "pointer_map_access",
			key:      "pointer_map.key1.value",
			expected: "map_ptr_value1",
		},
		{
			name:     "pointer_map_second_key",
			key:      "pointer_map.key2.value",
			expected: "map_ptr_value2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Extract(testData, tt.key)
			require.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}

	// Test error cases with pointers
	errorTests := []struct {
		name      string
		key       string
		errorType error
	}{
		{
			name:      "nil_pointer_in_chain",
			key:       "level1_val.level2_val.level3_ptr.value",
			errorType: ErrInvalidKeyType, // jsonpointer.ErrNilPointer maps to ErrInvalidKeyType
		},
		{
			name:      "nil_pointer_in_slice",
			key:       "pointer_slice.2.value",
			errorType: ErrInvalidKeyType, // jsonpointer.ErrNilPointer maps to ErrInvalidKeyType
		},
		{
			name:      "nil_pointer_in_map",
			key:       "pointer_map.key3.value",
			errorType: ErrInvalidKeyType, // jsonpointer.ErrNilPointer maps to ErrInvalidKeyType
		},
		{
			name:      "out_of_bounds_pointer_slice",
			key:       "pointer_slice.10.value",
			errorType: ErrIndexOutOfRange,
		},
	}

	for _, tt := range errorTests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Extract(testData, tt.key)
			require.Error(t, err)
			require.ErrorIs(t, err, tt.errorType)
		})
	}
}

// TestExtractInterfaceHandling tests handling of interface{} types
func TestExtractInterfaceHandling(t *testing.T) {
	// Test with interface{} containing various types
	interfaceData := map[string]interface{}{
		"string_val": "hello",
		"int_val":    42,
		"float_val":  3.14,
		"bool_val":   true,
		"nil_val":    nil,
		"nested_map": map[string]interface{}{
			"inner_string": "inner_value",
			"inner_int":    100,
		},
		"nested_slice": []interface{}{
			"slice_string",
			123,
			map[string]interface{}{
				"slice_map_key": "slice_map_value",
			},
		},
		"mixed_types": []interface{}{
			"string",
			42,
			map[string]interface{}{
				"nested": "value",
			},
		},
	}

	tests := []struct {
		name     string
		key      string
		expected interface{}
	}{
		{
			name:     "interface_string",
			key:      "string_val",
			expected: "hello",
		},
		{
			name:     "interface_int",
			key:      "int_val",
			expected: 42,
		},
		{
			name:     "interface_nested_map",
			key:      "nested_map.inner_string",
			expected: "inner_value",
		},
		{
			name:     "interface_nested_slice",
			key:      "nested_slice.0",
			expected: "slice_string",
		},
		{
			name:     "interface_deep_nested",
			key:      "nested_slice.2.slice_map_key",
			expected: "slice_map_value",
		},
		{
			name:     "interface_mixed_types",
			key:      "mixed_types.2.nested",
			expected: "value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Extract(interfaceData, tt.key)
			require.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}

	// Test error cases with interfaces
	errorTests := []struct {
		name      string
		key       string
		errorType error
	}{
		{
			name:      "nil_interface_navigation",
			key:       "nil_val.something",
			errorType: ErrInvalidKeyType,
		},
		{
			name:      "primitive_interface_navigation",
			key:       "int_val.something",
			errorType: ErrInvalidKeyType,
		},
		{
			name:      "nonexistent_interface_key",
			key:       "nonexistent",
			errorType: ErrKeyNotFound,
		},
	}

	for _, tt := range errorTests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Extract(interfaceData, tt.key)
			require.Error(t, err)
			require.ErrorIs(t, err, tt.errorType)
		})
	}
}

// Benchmark tests for data extraction operations

func BenchmarkExtractMap(b *testing.B) {
	data := map[string]interface{}{
		"level1": map[string]interface{}{
			"level2": map[string]interface{}{
				"value": "deep",
			},
		},
	}
	b.ResetTimer()
	for b.Loop() {
		_, _ = Extract(data, "level1.level2.value")
	}
}

func BenchmarkExtractSlice(b *testing.B) {
	data := []interface{}{
		"first",
		"second",
		map[string]interface{}{"key": "value"},
	}
	b.ResetTimer()
	for b.Loop() {
		_, _ = Extract(data, "2.key")
	}
}

func BenchmarkExtractStruct(b *testing.B) {
	type Inner struct {
		Value string `json:"value"`
	}
	type Outer struct {
		Inner Inner `json:"inner"`
	}
	data := Outer{Inner: Inner{Value: "test"}}
	b.ResetTimer()
	for b.Loop() {
		_, _ = Extract(data, "inner.value")
	}
}
