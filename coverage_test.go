package filter

import (
	"testing"
	"time"

	"github.com/dromara/carbon/v2"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func TestUniqueDeduplicatesNonComparableValues(t *testing.T) {
	t.Parallel()

	a := 1
	b := 1
	type sample struct {
		ID   int
		Name string
	}

	input := []any{
		true, true,
		float64(1.5), float64(1.5),
		int(3), int(3),
		"x", "x",
		[]any{"nested", 1}, []any{"nested", 1},
		map[string]any{"k": "v"}, map[string]any{"k": "v"},
		int64(7), int64(7),
		uint32(9), uint32(9),
		float32(2.5), float32(2.5),
		[2]int{1, 2}, [2]int{1, 2},
		map[int]string{1: "a"}, map[int]string{1: "a"},
		&a, &b,
		sample{ID: 1, Name: "alpha"}, sample{ID: 1, Name: "alpha"},
		nil, nil,
	}

	got, err := Unique(input)
	require.NoError(t, err)

	want := []any{
		true,
		float64(1.5),
		int(3),
		"x",
		[]any{"nested", 1},
		map[string]any{"k": "v"},
		int64(7),
		uint32(9),
		float32(2.5),
		[2]int{1, 2},
		map[int]string{1: "a"},
		&a,
		sample{ID: 1, Name: "alpha"},
		nil,
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Fatalf("Unique() mismatch (-want +got):\n%s", diff)
	}
}

func TestFirst(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   any
		want    any
		wantErr error
	}{
		{name: "returns first element", input: []any{"a", "b"}, want: "a"},
		{name: "returns empty slice error", input: []any{}, wantErr: ErrEmptySlice},
		{name: "returns not slice error", input: "not-a-slice", wantErr: ErrNotSlice},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got, err := First(tc.input)
			if tc.wantErr != nil {
				require.ErrorIs(t, err, tc.wantErr)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tc.want, got)
		})
	}
}

func TestCollectionFunctionsRejectNonSlices(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		call func() error
	}{
		{name: "sort natural", call: func() error { _, err := SortNatural("not-a-slice"); return err }},
		{name: "compact", call: func() error { _, err := Compact("not-a-slice"); return err }},
		{name: "where", call: func() error { _, err := Where("not-a-slice", "name"); return err }},
		{name: "reject", call: func() error { _, err := Reject("not-a-slice", "name"); return err }},
		{name: "find", call: func() error { _, err := Find("not-a-slice", "name", "x"); return err }},
		{name: "find index", call: func() error { _, err := FindIndex("not-a-slice", "name", "x"); return err }},
		{name: "has", call: func() error { _, err := Has("not-a-slice", "name"); return err }},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			require.ErrorIs(t, tc.call(), ErrNotSlice)
		})
	}
}

func TestCollectionFunctionsAdditionalBehaviors(t *testing.T) {
	t.Parallel()

	items := []any{
		map[string]any{"flag": true, "count": int64(10)},
		map[string]any{"flag": false, "count": 0},
		map[string]any{"flag": nil, "count": "10"},
		map[string]any{"name": "missing flag"},
	}

	gotWhere, err := Where(items, "flag")
	require.NoError(t, err)
	wantWhere := []any{
		map[string]any{"flag": true, "count": int64(10)},
	}
	if diff := cmp.Diff(wantWhere, gotWhere); diff != "" {
		t.Fatalf("Where() mismatch (-want +got):\n%s", diff)
	}

	gotReject, err := Reject(items, "flag")
	require.NoError(t, err)
	wantReject := []any{
		map[string]any{"flag": false, "count": 0},
		map[string]any{"flag": nil, "count": "10"},
		map[string]any{"name": "missing flag"},
	}
	if diff := cmp.Diff(wantReject, gotReject); diff != "" {
		t.Fatalf("Reject() mismatch (-want +got):\n%s", diff)
	}

	gotFind, err := Find(items, "count", "10")
	require.NoError(t, err)
	require.Equal(t, map[string]any{"flag": true, "count": int64(10)}, gotFind)

	gotIndex, err := FindIndex([]any{}, "count", 1)
	require.NoError(t, err)
	require.Equal(t, -1, gotIndex)

	gotHas, err := Has(items, "count", 10.0)
	require.NoError(t, err)
	require.True(t, gotHas)
}

func TestSortAdditionalCases(t *testing.T) {
	t.Parallel()

	t.Run("keeps missing keyed values stable and first", func(t *testing.T) {
		t.Parallel()

		input := []any{
			map[string]any{"name": "beta"},
			map[string]any{"id": 1},
			map[string]any{"name": "Alpha"},
		}

		got, err := Sort(input, "name")
		require.NoError(t, err)

		want := []any{
			map[string]any{"id": 1},
			map[string]any{"name": "Alpha"},
			map[string]any{"name": "beta"},
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Fatalf("Sort() mismatch (-want +got):\n%s", diff)
		}
	})
}

func TestSortNaturalAdditionalCases(t *testing.T) {
	t.Parallel()

	t.Run("sorts numerically when values are numbers", func(t *testing.T) {
		t.Parallel()

		got, err := SortNatural([]any{10, 2, 1})
		require.NoError(t, err)

		want := []any{1, 2, 10}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Fatalf("SortNatural() mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("keeps missing keyed values stable and first", func(t *testing.T) {
		t.Parallel()

		input := []any{
			map[string]any{"name": "beta"},
			map[string]any{"id": 1},
			map[string]any{"name": "Alpha"},
		}

		got, err := SortNatural(input, "name")
		require.NoError(t, err)

		want := []any{
			map[string]any{"id": 1},
			map[string]any{"name": "Alpha"},
			map[string]any{"name": "beta"},
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Fatalf("SortNatural() mismatch (-want +got):\n%s", diff)
		}
	})
}

func TestNumberSupportsAdditionalNumericTypes(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input any
		want  string
	}{
		{name: "int8", input: int8(12), want: "12.00"},
		{name: "int16", input: int16(12), want: "12.00"},
		{name: "int32", input: int32(12), want: "12.00"},
		{name: "int64", input: int64(12), want: "12.00"},
		{name: "uint", input: uint(12), want: "12.00"},
		{name: "uint8", input: uint8(12), want: "12.00"},
		{name: "uint16", input: uint16(12), want: "12.00"},
		{name: "uint32", input: uint32(12), want: "12.00"},
		{name: "uint64", input: uint64(12), want: "12.00"},
		{name: "float32", input: float32(12), want: "12.00"},
		{name: "numeric string", input: "12", want: "12.00"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got, err := Number(tc.input, "#,###.##")
			require.NoError(t, err)
			require.Equal(t, tc.want, got)
		})
	}

	_, err := Number(struct{}{}, "")
	require.ErrorIs(t, err, ErrNotNumeric)
}

func TestBytesAdditionalBehaviors(t *testing.T) {
	t.Parallel()

	got, err := Bytes("1024")
	require.NoError(t, err)
	require.Equal(t, "1.0 kB", got)

	_, err = Bytes(-1)
	require.ErrorIs(t, err, ErrNegativeValue)
}

func TestPascalize(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "empty string", input: "", want: ""},
		{name: "snake case with acronym", input: "profile_id", want: "ProfileID"},
		{name: "unicode words", input: "résumé operation", want: "RésuméOperation"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			require.Equal(t, tc.want, Pascalize(tc.input))
		})
	}
}

func TestDateAcceptsCarbonPointer(t *testing.T) {
	t.Parallel()

	input := carbon.CreateFromStdTime(time.Date(2024, time.March, 30, 12, 0, 0, 0, time.UTC))
	got, err := Date(input, "Y-m-d")
	require.NoError(t, err)
	require.Equal(t, "2024-03-30", got)
}
