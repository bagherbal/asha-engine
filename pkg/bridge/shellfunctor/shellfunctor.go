// Package shellfunctor implements Gate 106: the finite shell functor / semigroup
// construction attempt.
//
// Gate 105 proved that the existing projection, quotient, spectral-ordering,
// action-weight, beta-vector, and threshold-classifier data do not yet define a
// native finite RG operator.  This package performs the next constructive
// attempt: build the most honest finite shell family that the current data allow
// and test its composition law.
//
// The result is useful but negative.  A nested family of finite shell
// projections exists and composes cleanly as an idempotent meet/semilattice:
// C_a o C_b = C_min(a,b) after ordering the current mode inventory.  That is a
// real functorial projection system, but it is not the additive/logarithmic RG
// semigroup needed for physical running.  It also needs an arbitrary shell
// ordering for the threshold-open modes and therefore cannot activate thresholds
// or reduce the Gate 104/105 residual nullity.
package shellfunctor

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/coarsegrain"
	"github.com/bagherbal/asha-engine/pkg/bridge/thresholdactivation"
)

type ModeClass string

const (
	ContinuumMode ModeClass = "continuum-candidate"
	OpenMode      ModeClass = "threshold-open"
	VacuumMode    ModeClass = "vacuum-only"
)

type Mode struct {
	Name        string
	Kind        string
	Status      string
	Class       ModeClass
	ShellRank   int
	InCarrier   bool
	CanActivate bool
	Detail      string
}

type ShellProjection struct {
	Name         string
	Parameter    int
	Keeps        []bool
	KeptCount    int
	RemovedCount int
	Detail       string
}

type FunctorAttempt struct {
	Name                     string
	Symbol                   string
	Constructed              bool
	EndofunctorOnFiniteModes bool
	Nontrivial               bool
	ClosureUnderComposition  bool
	IdentityExists           bool
	Associative              bool
	Idempotent               bool
	AdditiveLaw              bool
	CanonicalShellParameter  bool
	ThresholdActivationRule  bool
	AbsoluteScaleRule        bool
	RejectedAsNativeRG       bool
	Detail                   string
}

type CompositionWitness struct {
	Left       int
	Right      int
	Composed   int
	Expected   int
	Closed     bool
	AdditiveOK bool
	Detail     string
}

type ScheduleWitness struct {
	Name               string
	ShellOrder         []string
	CompatibleWithData bool
	SelectsPhysics     bool
	Detail             string
}

type Analysis struct {
	Coarse coarsegrain.Analysis

	Modes          []Mode
	ModeCount      int
	ContinuumCount int
	OpenCount      int
	VacuumCount    int

	Projections       []ShellProjection
	ProjectionCount   int
	FunctorAttempts   []FunctorAttempt
	CompositionTable  []CompositionWitness
	ScheduleWitnesses []ScheduleWitness

	NestedProjectionFamilyConstructed bool
	IdentityProjectionExists          bool
	CompositionClosed                 bool
	AssociativityVerified             bool
	IdempotentSemilatticeDerived      bool
	AdditiveSemigroupDerived          bool
	NontrivialAdditiveCounterexample  bool

	CanonicalShellOrderingDerived       bool
	CanonicalScaleLogParameterDerived   bool
	ThresholdActivationPredicateDerived bool
	DecouplingMatchingRuleDerived       bool
	AbsoluteCouplingFlowDerived         bool
	NativeFiniteRGFunctorDerived        bool

	ResidualNullityBefore  int
	ResidualNullityAfter   int
	ResidualSymmetryBroken bool

	PhysicalWeakAngleDerived bool
	FineStructureDerived     bool
	PhysicalMassesDerived    bool
	HiddenObservedInputUsed  bool

	TruthStatement      string
	RejectedClaims      []string
	RemainingUnknowns   []string
	RecommendedNextGate string
}

var (
	defaultOnce  sync.Once
	defaultValue Analysis
	defaultErr   error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		coarse, err := coarsegrain.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(coarse)
	})
	return defaultValue, defaultErr
}

func Build(coarse coarsegrain.Analysis) (Analysis, error) {
	if !coarse.FiniteBoundarySeedInherited || coarse.BoundaryKY <= 0 {
		return Analysis{}, fmt.Errorf("Gate 106 requires Gate 105 with inherited finite boundary seed")
	}
	if coarse.NativeCoarseGrainingFound || coarse.SemigroupLawDerived || coarse.ScaleLogParameterDerived {
		return Analysis{}, fmt.Errorf("Gate 106 expects Gate 105 to leave native RG/coarse-graining open")
	}
	if coarse.HiddenObservedInputUsed || coarse.PhysicalWeakAngleDerived || coarse.FineStructureDerived || coarse.PhysicalMassesDerived {
		return Analysis{}, fmt.Errorf("Gate 106 refuses hidden observed physical input")
	}
	if len(coarse.Thresholds.Decisions) == 0 {
		return Analysis{}, fmt.Errorf("Gate 106 requires threshold decisions as finite shell carrier")
	}

	modes := buildModes(coarse.Thresholds.Decisions)
	continuum, open, vacuum := countModes(modes)
	projections := buildNestedProjections(modes)
	table := buildCompositionTable(projections)

	closure := true
	additive := true
	counterexample := false
	for _, w := range table {
		if !w.Closed {
			closure = false
		}
		if !w.AdditiveOK {
			additive = false
			if w.Left > 0 && w.Right > 0 && w.Left+w.Right < len(projections) {
				counterexample = true
			}
		}
	}
	identity := len(projections) > 0 && projections[len(projections)-1].KeptCount >= continuum+open
	associative := verifyAssociative(projections)
	idempotent := verifyIdempotent(table)

	attempts := []FunctorAttempt{
		{
			Name:                     "identity shell functor",
			Symbol:                   "Id",
			Constructed:              true,
			EndofunctorOnFiniteModes: true,
			Nontrivial:               false,
			ClosureUnderComposition:  true,
			IdentityExists:           true,
			Associative:              true,
			Idempotent:               true,
			AdditiveLaw:              true,
			CanonicalShellParameter:  true,
			RejectedAsNativeRG:       true,
			Detail:                   "identity is a valid endofunctor and semigroup unit, but it performs no coarse-graining and creates no running or thresholds",
		},
		{
			Name:                     "status projection functor",
			Symbol:                   "P_cont",
			Constructed:              true,
			EndofunctorOnFiniteModes: true,
			Nontrivial:               continuum < continuum+open,
			ClosureUnderComposition:  true,
			IdentityExists:           true,
			Associative:              true,
			Idempotent:               true,
			AdditiveLaw:              false,
			ThresholdActivationRule:  false,
			RejectedAsNativeRG:       true,
			Detail:                   "projects onto currently continuum-candidate modes; repeats idempotently but cannot decide open thresholds",
		},
		{
			Name:                     "nested finite shell projection family",
			Symbol:                   "C_n",
			Constructed:              true,
			EndofunctorOnFiniteModes: true,
			Nontrivial:               len(projections) >= 3,
			ClosureUnderComposition:  closure,
			IdentityExists:           identity,
			Associative:              associative,
			Idempotent:               idempotent,
			AdditiveLaw:              additive,
			CanonicalShellParameter:  false,
			ThresholdActivationRule:  false,
			AbsoluteScaleRule:        false,
			RejectedAsNativeRG:       true,
			Detail:                   "composes as C_a∘C_b=C_min(a,b) for the nested projection convention; this is an idempotent semilattice, not the additive/logarithmic RG law C_s∘C_t=C_{s+t}",
		},
		{
			Name:                     "hypothetical additive finite RG step",
			Symbol:                   "R_n",
			Constructed:              false,
			EndofunctorOnFiniteModes: false,
			Nontrivial:               false,
			ClosureUnderComposition:  false,
			IdentityExists:           false,
			Associative:              false,
			Idempotent:               false,
			AdditiveLaw:              false,
			CanonicalShellParameter:  false,
			ThresholdActivationRule:  false,
			AbsoluteScaleRule:        false,
			RejectedAsNativeRG:       true,
			Detail:                   "no current finite operator defines an additive shell index, beta derivative, or absolute scale variable",
		},
	}

	schedules := buildScheduleWitnesses(modes)

	truth := "Gate 106 constructs the strongest finite shell object currently available: a nested family of projections on the threshold/mode carrier.  The family is real and composable, but its law is idempotent semilattice composition C_a∘C_b=C_min(a,b), not additive/logarithmic RG flow.  Since the shell order for open modes is not canonically selected and no scale, activation predicate, or matching rule is derived, the residual physical-flow nullity remains unchanged."

	return Analysis{
		Coarse:                              coarse,
		Modes:                               modes,
		ModeCount:                           len(modes),
		ContinuumCount:                      continuum,
		OpenCount:                           open,
		VacuumCount:                         vacuum,
		Projections:                         projections,
		ProjectionCount:                     len(projections),
		FunctorAttempts:                     attempts,
		CompositionTable:                    table,
		ScheduleWitnesses:                   schedules,
		NestedProjectionFamilyConstructed:   len(projections) >= 2,
		IdentityProjectionExists:            identity,
		CompositionClosed:                   closure,
		AssociativityVerified:               associative,
		IdempotentSemilatticeDerived:        idempotent && closure && !additive,
		AdditiveSemigroupDerived:            additive,
		NontrivialAdditiveCounterexample:    counterexample,
		CanonicalShellOrderingDerived:       false,
		CanonicalScaleLogParameterDerived:   false,
		ThresholdActivationPredicateDerived: false,
		DecouplingMatchingRuleDerived:       false,
		AbsoluteCouplingFlowDerived:         false,
		NativeFiniteRGFunctorDerived:        false,
		ResidualNullityBefore:               coarse.ResidualNullityAfter,
		ResidualNullityAfter:                coarse.ResidualNullityAfter,
		ResidualSymmetryBroken:              false,
		PhysicalWeakAngleDerived:            false,
		FineStructureDerived:                false,
		PhysicalMassesDerived:               false,
		HiddenObservedInputUsed:             false,
		TruthStatement:                      truth,
		RejectedClaims: []string{
			"a nested projection semilattice is already a logarithmic RG semigroup",
			"an arbitrary ordering of threshold-open modes is a derived shell parameter",
			"C_a∘C_b=C_min(a,b) can be renamed into C_{a+b} without changing the mathematics",
			"continuum-active scalar candidates determine heavy-threshold matching by themselves",
			"finite shell projections reduce the absolute coupling/scale/threshold nullity",
		},
		RemainingUnknowns: []string{
			"U-26A-ORDERING: derive a canonical order or filtration on threshold-open finite modes",
			"U-26B-ADDITIVE-LAW: construct a non-idempotent finite map with C_s∘C_t=C_{s+t} or prove impossible",
			"U-26C-SCALE-FUNCTOR: derive the variable replacing L=ln(M*/μ) from finite data",
			"U-26D-ACTIVATION: derive a threshold predicate from the finite shell functor rather than from choice",
			"U-26E-MATCHING: map activated shells to Δb_i contributions without observed masses",
			"U-26F-FLOW-UNIT: derive whether the gauge-action prefactor runs under the shell functor",
		},
		RecommendedNextGate: "Gate 107 — finite filtration/order selector and monotone threshold predicate search",
	}, nil
}

func buildModes(decisions []thresholdactivation.ActivationDecision) []Mode {
	modes := make([]Mode, 0, len(decisions))
	shellRank := 0
	for _, d := range decisions {
		class := OpenMode
		rank := shellRank
		inCarrier := true
		canActivate := false
		switch d.Status {
		case thresholdactivation.ContinuumFieldCandidate:
			class = ContinuumMode
			rank = 0
			canActivate = true
		case thresholdactivation.VacuumFrustrationOnly:
			class = VacuumMode
			rank = math.MaxInt / 4
			inCarrier = false
		default:
			class = OpenMode
			shellRank++
			rank = shellRank
		}
		modes = append(modes, Mode{
			Name:        d.Assignment.Candidate.Name,
			Kind:        string(d.Assignment.Candidate.Kind),
			Status:      string(d.Status),
			Class:       class,
			ShellRank:   rank,
			InCarrier:   inCarrier,
			CanActivate: canActivate && d.CanCorrectBeta,
			Detail:      d.Reason,
		})
	}
	return modes
}

func countModes(modes []Mode) (continuum, open, vacuum int) {
	for _, m := range modes {
		switch m.Class {
		case ContinuumMode:
			continuum++
		case OpenMode:
			open++
		case VacuumMode:
			vacuum++
		}
	}
	return continuum, open, vacuum
}

func buildNestedProjections(modes []Mode) []ShellProjection {
	maxRank := 0
	for _, m := range modes {
		if m.Class == OpenMode && m.ShellRank > maxRank {
			maxRank = m.ShellRank
		}
	}
	projections := make([]ShellProjection, 0, maxRank+1)
	for n := 0; n <= maxRank; n++ {
		keeps := make([]bool, len(modes))
		kept := 0
		for i, m := range modes {
			keep := false
			switch m.Class {
			case ContinuumMode:
				keep = true
			case OpenMode:
				keep = m.ShellRank <= n
			case VacuumMode:
				keep = false
			}
			keeps[i] = keep
			if keep {
				kept++
			}
		}
		projections = append(projections, ShellProjection{
			Name:         fmt.Sprintf("C_%d", n),
			Parameter:    n,
			Keeps:        keeps,
			KeptCount:    kept,
			RemovedCount: len(modes) - kept,
			Detail:       fmt.Sprintf("keeps continuum modes plus threshold-open modes with arbitrary shell rank <= %d; vacuum-only modes remain excluded", n),
		})
	}
	return projections
}

func buildCompositionTable(projections []ShellProjection) []CompositionWitness {
	out := make([]CompositionWitness, 0, len(projections)*len(projections))
	for _, a := range projections {
		for _, b := range projections {
			composed := compose(a, b)
			expectedParam := a.Parameter
			if b.Parameter < expectedParam {
				expectedParam = b.Parameter
			}
			expected := projections[expectedParam]
			closed := sameKept(composed, expected.Keeps)
			additiveOK := false
			if a.Parameter+b.Parameter < len(projections) {
				additiveOK = sameKept(composed, projections[a.Parameter+b.Parameter].Keeps)
			} else if a.Parameter == 0 || b.Parameter == 0 {
				// For the projection convention, C_0 is not an additive identity; this branch stays false
				// except when the actual kept sets coincide with a saturated candidate.
				additiveOK = sameKept(composed, projections[int(math.Min(float64(a.Parameter+b.Parameter), float64(len(projections)-1)))].Keeps)
			}
			out = append(out, CompositionWitness{
				Left:       a.Parameter,
				Right:      b.Parameter,
				Composed:   expectedParam,
				Expected:   expectedParam,
				Closed:     closed,
				AdditiveOK: additiveOK,
				Detail:     fmt.Sprintf("C_%d∘C_%d = C_%d under nested projection/intersection; additive target would be C_%d when defined", a.Parameter, b.Parameter, expectedParam, a.Parameter+b.Parameter),
			})
		}
	}
	return out
}

func compose(a, b ShellProjection) []bool {
	n := len(a.Keeps)
	if len(b.Keeps) < n {
		n = len(b.Keeps)
	}
	out := make([]bool, n)
	for i := 0; i < n; i++ {
		out[i] = a.Keeps[i] && b.Keeps[i]
	}
	return out
}

func sameKept(a, b []bool) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func verifyAssociative(projections []ShellProjection) bool {
	for _, a := range projections {
		for _, b := range projections {
			for _, c := range projections {
				left := composeRaw(compose(a, b), c.Keeps)
				right := composeRaw(a.Keeps, compose(b, c))
				if !sameKept(left, right) {
					return false
				}
			}
		}
	}
	return len(projections) > 0
}

func verifyIdempotent(table []CompositionWitness) bool {
	for _, w := range table {
		if w.Left == w.Right && (!w.Closed || w.Composed != w.Left) {
			return false
		}
	}
	return len(table) > 0
}

func composeRaw(a, b []bool) []bool {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	out := make([]bool, n)
	for i := 0; i < n; i++ {
		out[i] = a[i] && b[i]
	}
	return out
}

func buildScheduleWitnesses(modes []Mode) []ScheduleWitness {
	openNames := make([]string, 0)
	reverseOpenNames := make([]string, 0)
	for _, m := range modes {
		if m.Class == OpenMode {
			openNames = append(openNames, m.Name)
		}
	}
	for i := len(openNames) - 1; i >= 0; i-- {
		reverseOpenNames = append(reverseOpenNames, openNames[i])
	}
	return []ScheduleWitness{
		{
			Name:               "status-first schedule",
			ShellOrder:         openNames,
			CompatibleWithData: true,
			SelectsPhysics:     false,
			Detail:             "uses the current threshold inventory order; compatible as bookkeeping but not derived as physical activation order",
		},
		{
			Name:               "reverse-open schedule",
			ShellOrder:         reverseOpenNames,
			CompatibleWithData: true,
			SelectsPhysics:     false,
			Detail:             "the reverse order is equally compatible with current finite data, proving non-uniqueness of shell ordering",
		},
	}
}

func FormatModes(xs []Mode) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		rank := fmt.Sprintf("%d", x.ShellRank)
		if x.Class == VacuumMode {
			rank = "excluded"
		}
		parts = append(parts, fmt.Sprintf("%s [%s/%s, rank=%s, carrier=%t]", x.Name, x.Kind, x.Class, rank, x.InCarrier))
	}
	return strings.Join(parts, "; ")
}

func FormatProjections(xs []ShellProjection) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s: kept=%d removed=%d", x.Name, x.KeptCount, x.RemovedCount))
	}
	return strings.Join(parts, "; ")
}

func FormatFunctorAttempts(xs []FunctorAttempt) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		verdict := "rejected"
		if x.Constructed && x.AdditiveLaw && x.Nontrivial && x.CanonicalShellParameter && x.ThresholdActivationRule && x.AbsoluteScaleRule {
			verdict = "native-rg"
		}
		parts = append(parts, fmt.Sprintf("%s:%s constructed=%t nontrivial=%t closed=%t additive=%t canonical-shell=%t threshold=%t scale=%t %s", x.Symbol, x.Name, x.Constructed, x.Nontrivial, x.ClosureUnderComposition, x.AdditiveLaw, x.CanonicalShellParameter, x.ThresholdActivationRule, x.AbsoluteScaleRule, verdict))
	}
	return strings.Join(parts, "; ")
}

func FormatCompositionWitnesses(xs []CompositionWitness) string {
	samples := make([]string, 0)
	for _, x := range xs {
		if x.Left == 1 && x.Right == 2 || x.Left == 2 && x.Right == 3 || x.Left == x.Right && x.Left <= 2 {
			samples = append(samples, fmt.Sprintf("C_%d∘C_%d=C_%d closed=%t additiveOK=%t", x.Left, x.Right, x.Composed, x.Closed, x.AdditiveOK))
		}
	}
	if len(samples) == 0 {
		return "no composition samples"
	}
	return strings.Join(samples, "; ")
}

func FormatSchedules(xs []ScheduleWitness) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		head := x.ShellOrder
		if len(head) > 3 {
			head = head[:3]
		}
		parts = append(parts, fmt.Sprintf("%s compatible=%t selects-physics=%t first=%v", x.Name, x.CompatibleWithData, x.SelectsPhysics, head))
	}
	return strings.Join(parts, "; ")
}

func Join(xs []string) string { return strings.Join(xs, "; ") }
