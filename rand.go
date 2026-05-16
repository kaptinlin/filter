package filter

import "math/rand/v2"

// SeededRand returns a deterministic *rand.Rand for callers that need
// reproducible RandomWithRand or ShuffleWithRand output.
//
//	r := filter.SeededRand(1, 2)
//	got, _ := filter.ShuffleWithRand(r, []int{1, 2, 3, 4})
//	// got is the same on every run.
//
// The returned *rand.Rand follows math/rand/v2's usual rule: it is intended
// for one goroutine at a time. Use the package-level Random and Shuffle
// helpers when callers only need non-deterministic, concurrent-safe behavior.
func SeededRand(s1, s2 uint64) *rand.Rand {
	// #nosec G404 -- deterministic non-cryptographic randomness is the point of this helper.
	return rand.New(rand.NewPCG(s1, s2))
}
