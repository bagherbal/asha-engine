package clifford

import "fmt"

type Signature struct {
	Positive int
	Negative int
}

func (s Signature) Dimension() int {
	return s.Positive + s.Negative
}

func (s Signature) Validate() error {
	if s.Positive < 0 || s.Negative < 0 {
		return fmt.Errorf("signature entries must be non-negative: %+v", s)
	}
	if s.Dimension() == 0 {
		return fmt.Errorf("signature dimension must be positive")
	}
	return nil
}

func (s Signature) DiagonalMetric() []int {
	metric := make([]int, 0, s.Dimension())
	for i := 0; i < s.Positive; i++ {
		metric = append(metric, +1)
	}
	for i := 0; i < s.Negative; i++ {
		metric = append(metric, -1)
	}
	return metric
}
