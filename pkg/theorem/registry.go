package theorem

type Registry struct {
	theorems []Theorem
}

func NewRegistry(theorems ...Theorem) Registry {
	return Registry{theorems: append([]Theorem(nil), theorems...)}
}

func (r Registry) RunAll() []Result {
	results := make([]Result, 0, len(r.theorems))
	for _, t := range r.theorems {
		results = append(results, t.Run())
	}
	return results
}
