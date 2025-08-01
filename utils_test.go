package filter

import (
	"testing"
	"time"

	"github.com/dromara/carbon/v2"
	"github.com/stretchr/testify/assert"
)

func TestToCarbon(t *testing.T) {
	// Define test cases
	testCases := []struct {
		name    string
		input   interface{}
		wantErr bool
	}{
		{
			name:    "Valid carbon.Carbon input",
			input:   carbon.Now(),
			wantErr: false,
		},
		{
			name:    "Valid time.Time input",
			input:   time.Now(),
			wantErr: false,
		},
		{
			name:    "Valid string input",
			input:   "2023-03-28",
			wantErr: false,
		},
		{
			name:    "Invalid string input",
			input:   "invalid-date",
			wantErr: true,
		},
		{
			name:    "Unsupported type",
			input:   12345,
			wantErr: true,
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := toCarbon(tc.input)
			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
