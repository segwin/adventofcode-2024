package day5

type PageRule struct {
	Before int
	After  int
}

// AggregateRules collects the given rules into an easily-queried map.
// Each entry in the map defines the set of pages that must come after the keyed page.
func AggregateRules(rules []PageRule) map[int][]int {
	out := map[int][]int{}
	for _, r := range rules {
		out[r.Before] = append(out[r.Before], r.After)
	}
	return out
}
