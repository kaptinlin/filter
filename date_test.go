package filter

import (
	"testing"
	"time"

	"github.com/dromara/carbon/v2"
	"github.com/stretchr/testify/require"
)

func TestDate(t *testing.T) {
	// Set up some reference times for testing
	now := time.Now()
	carbonNow := carbon.Now()

	// Define test cases
	tests := []struct {
		name      string
		input     interface{}
		format    string
		want      string
		expectErr bool
	}{
		{
			name:      "carbon.Carbon input with custom format",
			input:     carbonNow,
			format:    "Y-m-d",
			want:      carbonNow.Format("Y-m-d"),
			expectErr: false,
		},
		{
			name:      "time.Time input with empty format (default)",
			input:     now,
			format:    "",
			want:      carbon.CreateFromStdTime(now).ToDateTimeString(),
			expectErr: false,
		},
		{
			name:      "string input valid date with custom format",
			input:     "2024-03-30",
			format:    "d/m/Y",
			want:      "30/03/2024",
			expectErr: false,
		},
		{
			name:      "string input invalid date",
			input:     "not-a-date",
			format:    "Y-m-d",
			expectErr: true,
		},
		{
			name:      "unsupported input type",
			input:     12345,
			format:    "Y-m-d",
			expectErr: true,
		},
		{
			name:      "string input with month number",
			input:     "2024-01-01",
			format:    "m",
			want:      "01",
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Date(tt.input, tt.format)
			if tt.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.want, got)
			}
		})
	}
}

func TestDay(t *testing.T) {
	// Define test cases
	tests := []struct {
		name      string
		input     interface{}
		wantDay   int
		expectErr bool
	}{
		{
			name:      "carbon.Carbon input",
			input:     carbon.Now().SetDay(15),
			wantDay:   15,
			expectErr: false,
		},
		{
			name:      "time.Time input",
			input:     time.Date(2024, time.March, 10, 0, 0, 0, 0, time.UTC),
			wantDay:   10,
			expectErr: false,
		},
		{
			name:      "string input valid date",
			input:     "2024-03-20",
			wantDay:   20,
			expectErr: false,
		},
		{
			name:      "string input invalid date",
			input:     "not-a-date",
			expectErr: true,
		},
		{
			name:      "unsupported input type",
			input:     12345,
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDay, err := Day(tt.input)
			if tt.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.wantDay, gotDay)
			}
		})
	}
}

func TestMonth(t *testing.T) {
	// Define test cases
	tests := []struct {
		name      string
		input     interface{}
		wantMonth int
		expectErr bool
	}{
		{
			name:      "carbon.Carbon input for January",
			input:     carbon.Now().SetMonth(int(time.January)),
			wantMonth: 1,
			expectErr: false,
		},
		{
			name:      "time.Time input for December",
			input:     time.Date(2024, time.December, 1, 0, 0, 0, 0, time.UTC),
			wantMonth: 12,
			expectErr: false,
		},
		{
			name:      "string input valid date for June",
			input:     "2024-06-15",
			wantMonth: 6,
			expectErr: false,
		},
		{
			name:      "string input invalid date",
			input:     "this-is-not-a-date",
			expectErr: true,
		},
		{
			name:      "unsupported input type",
			input:     []int{2024, 4},
			expectErr: true,
		},
		{
			name:      "string input for February in leap year",
			input:     "2024-02-29",
			wantMonth: 2,
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMonth, err := Month(tt.input)
			if tt.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.wantMonth, gotMonth, "Expected and actual month should match")
			}
		})
	}
}

func TestMonthFull(t *testing.T) {
	// Define test cases
	tests := []struct {
		name          string
		input         interface{}
		expectedMonth string
		expectErr     bool
	}{
		{
			name:          "carbon.Carbon input for January",
			input:         carbon.Now().SetMonth(int(time.January)),
			expectedMonth: "January",
			expectErr:     false,
		},
		{
			name:          "time.Time input for December",
			input:         time.Date(2024, time.December, 1, 0, 0, 0, 0, time.UTC),
			expectedMonth: "December",
			expectErr:     false,
		},
		{
			name:          "string input valid date for June",
			input:         "2024-06-15",
			expectedMonth: "June",
			expectErr:     false,
		},
		{
			name:      "string input invalid date",
			input:     "not-a-real-date",
			expectErr: true,
		},
		{
			name:      "unsupported input type",
			input:     123456,
			expectErr: true,
		},
		{
			name:          "string input for February in a leap year",
			input:         "2024-02-29",
			expectedMonth: "February",
			expectErr:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMonth, err := MonthFull(tt.input)
			if tt.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expectedMonth, gotMonth, "The expected and actual full month name should match.")
			}
		})
	}
}

func TestYear(t *testing.T) {
	// Define test cases
	tests := []struct {
		name         string
		input        interface{}
		expectedYear int
		expectErr    bool
	}{
		{
			name:         "carbon.Carbon input for current year",
			input:        carbon.Now(),
			expectedYear: time.Now().Year(),
			expectErr:    false,
		},
		{
			name:         "time.Time input for specific year",
			input:        time.Date(1999, time.December, 31, 0, 0, 0, 0, time.UTC),
			expectedYear: 1999,
			expectErr:    false,
		},
		{
			name:         "string input valid date for year",
			input:        "2024-06-15",
			expectedYear: 2024,
			expectErr:    false,
		},
		{
			name:      "string input invalid date",
			input:     "not-a-valid-date",
			expectErr: true,
		},
		{
			name:      "unsupported input type",
			input:     []int{2024},
			expectErr: true,
		},
		{
			name:         "string input for leap year",
			input:        "2024-02-29",
			expectedYear: 2024,
			expectErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotYear, err := Year(tt.input)
			if tt.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expectedYear, gotYear, "The expected and actual year should match.")
			}
		})
	}
}

func TestWeek(t *testing.T) {
	// Define test cases
	tests := []struct {
		name         string
		input        interface{}
		expectedWeek int
		expectErr    bool
	}{
		{
			name:         "carbon.Carbon input for first week of year",
			input:        carbon.Now().SetDate(2024, 1, 3),
			expectedWeek: 1,
			expectErr:    false,
		},
		{
			name:         "time.Time input for last week of year",
			input:        time.Date(2022, 12, 28, 0, 0, 0, 0, time.UTC), // Late December can be part of the first week of the next year
			expectedWeek: 52,
			expectErr:    false,
		},
		{
			name:         "string input valid date mid-year",
			input:        "2024-06-15",
			expectedWeek: 24,
			expectErr:    false,
		},
		{
			name:      "string input invalid date",
			input:     "this-is-not-a-date",
			expectErr: true,
		},
		{
			name:      "unsupported input type",
			input:     98765,
			expectErr: true,
		},
		{
			name:         "string input for leap year end",
			input:        "2024-12-31",
			expectedWeek: 1, // Depending on the year, Dec 31 can be part of the first week of the next year
			expectErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotWeek, err := Week(tt.input)
			if tt.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expectedWeek, gotWeek, "The expected and actual week number should match.")
			}
		})
	}
}

func TestWeekday(t *testing.T) {
	tests := []struct {
		name            string
		input           interface{}
		expectedWeekday string
		expectErr       bool
	}{
		{
			name:            "carbon.Carbon input for a known weekday",
			input:           carbon.Now().SetDate(2024, 4, 5), // Adjusted to a specific date known to be a Wednesday
			expectedWeekday: "Friday",
			expectErr:       false,
		},
		{
			name:            "time.Time input for a known weekday",
			input:           time.Date(2024, time.April, 6, 0, 0, 0, 0, time.UTC), // Known to be a Thursday
			expectedWeekday: "Saturday",
			expectErr:       false,
		},
		{
			name:            "string input valid date for a known weekday",
			input:           "2024-04-07", // Known to be a Friday
			expectedWeekday: "Sunday",
			expectErr:       false,
		},
		{
			name:      "string input invalid date",
			input:     "this-is-not-a-date",
			expectErr: true,
		},
		{
			name:      "unsupported input type",
			input:     123456,
			expectErr: true,
		},
		{
			name:            "carbon.Carbon input for another known weekday",
			input:           carbon.Now().SetDate(2024, 4, 8), // Adjusted to a specific date known to be a Saturday
			expectedWeekday: "Monday",
			expectErr:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotWeekday, err := Weekday(tt.input)
			if tt.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expectedWeekday, gotWeekday, "The expected and actual weekday should match.")
			}
		})
	}
}

func TestTimeAgo(t *testing.T) {
	// Set a fixed test time
	fixedNow := carbon.Parse("2023-08-05 13:14:15")
	carbon.SetTestNow(fixedNow)
	defer carbon.ClearTestNow() // Clean up after test

	tests := []struct {
		name           string
		input          interface{}
		expectedOutput string
		expectErr      bool
	}{
		{
			name:           "One hour ago",
			input:          fixedNow.Copy().SubHours(1),
			expectedOutput: "1 hour ago",
			expectErr:      false,
		},
		{
			name:           "One day ago",
			input:          fixedNow.Copy().SubDays(1),
			expectedOutput: "1 day ago",
			expectErr:      false,
		},
		{
			name:           "One week ago",
			input:          fixedNow.Copy().SubWeeks(1),
			expectedOutput: "1 week ago",
			expectErr:      false,
		},
		{
			name:      "Unsupported type",
			input:     123,
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOutput, err := TimeAgo(tt.input)
			if tt.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				// Due to potential localization and precision, check if gotOutput contains expectedOutput
				require.Contains(t, gotOutput, tt.expectedOutput, "The output should contain the expected time difference.")
			}
		})
	}
}
