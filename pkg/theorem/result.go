package theorem

import (
	"fmt"
	"strings"
)

type Layer string

const (
	LayerAlgebra    Layer = "ALGEBRA"
	LayerGeometry   Layer = "GEOMETRY"
	LayerDynamics   Layer = "DYNAMICS"
	LayerGauge      Layer = "GAUGE"
	LayerMatter     Layer = "MATTER"
	LayerBridge     Layer = "BRIDGE"
	LayerValidation Layer = "VALIDATION"
)

type Check struct {
	Name   string
	Passed bool
	Detail string
}

type Result struct {
	ID     string
	Name   string
	Layer  Layer
	Status Status
	Checks []Check
	Notes  []string
}

func (r Result) Passed() bool {
	for _, c := range r.Checks {
		if !c.Passed {
			return false
		}
	}
	return true
}

func (r Result) Summary() string {
	state := "PASS"
	if !r.Passed() {
		state = "FAIL"
	}
	return fmt.Sprintf("%s [%s] %s: %s", state, r.Status, r.ID, r.Name)
}

func (r Result) Details() string {
	var b strings.Builder
	b.WriteString(r.Summary())
	for _, c := range r.Checks {
		mark := "✓"
		if !c.Passed {
			mark = "✗"
		}
		b.WriteString(fmt.Sprintf("\n  %s %s — %s", mark, c.Name, c.Detail))
	}
	for _, n := range r.Notes {
		b.WriteString(fmt.Sprintf("\n  note: %s", n))
	}
	return b.String()
}
