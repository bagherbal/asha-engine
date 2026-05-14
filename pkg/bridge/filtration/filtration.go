// Package filtration implements Gate 107: finite filtration/order selector and
// monotone threshold predicate search.
//
// Gate 106 constructed a real finite shell projection family, but that family
// needed an arbitrary order on the threshold-open modes and composed as an
// idempotent meet semilattice rather than as additive/logarithmic RG flow.  This
// package performs the next honest search: can the existing finite data select a
// canonical filtration order and a monotone threshold predicate?
//
// The result is again deliberately strict.  Several compatible filtrations can
// be constructed: status-only, spectral-value ascending, spectral-value
// descending, and arbitrary shell order.  All preserve the finite inventory, and
// value-based orders are internally monotone.  But the current engine does not
// derive an orientation, a cutoff, a physical scale, or a beta-matching rule.
// Therefore the only invariant predicate is the already-known safe classifier:
// continuum candidates remain candidates, threshold-open modes stay open, and
// vacuum-frustration modes stay excluded.  No threshold correction or physical
// RG output is selected.
package filtration

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/shellfunctor"
	"github.com/bagherbal/asha-engine/pkg/bridge/thresholdactivation"
)

type ModeClass string

const (
	ContinuumMode ModeClass = "continuum-candidate"
	OpenMode      ModeClass = "threshold-open"
	VacuumMode    ModeClass = "vacuum-only"
)

type FiltrationMode struct {
	Name       string
	Kind       string
	Status     string
	RepStatus  string
	Class      ModeClass
	Value      float64
	CanCorrect bool
	Detail     string
}

type SelectorAttempt struct {
	Name                         string
	Symbol                       string
	Constructed                  bool
	PreorderConstructed          bool
	TotalOrderConstructed        bool
	MonotonePredicateConstructed bool
	OrientationDerived           bool
	CutoffDerived                bool
	PhysicalScaleDerived         bool
	DeltaBMatchingDerived        bool
	BetaCorrectionAllowed        bool
	RejectedAsPhysicalSelector   bool
	Detail                       string
}

type OrderingWitness struct {
	Name                     string
	OpenOrder                []string
	FirstOpen                string
	LastOpen                 string
	CompatibleWithFiniteData bool
	MonotoneInSpectralValue  bool
	OrientationDerived       bool
	Canonical                bool
	SelectsThresholdPhysics  bool
	Detail                   string
}

type PredicateWitness struct {
	Name                     string
	Formula                  string
	Direction                string
	OpenModesActivated       int
	CompatibleWithFiniteData bool
	Monotone                 bool
	CutoffDerived            bool
	ScaleDerived             bool
	CanCorrectBeta           bool
	DeltaBSelected           bool
	Detail                   string
}

type AntichainWitness struct {
	ClassName string
	Members   []string
	Reason    string
}

type Analysis struct {
	Shell shellfunctor.Analysis

	Modes          []FiltrationMode
	ModeCount      int
	ContinuumCount int
	OpenCount      int
	VacuumCount    int

	SelectorAttempts   []SelectorAttempt
	OrderingWitnesses  []OrderingWitness
	PredicateWitnesses []PredicateWitness
	Antichains         []AntichainWitness

	StatusPreorderConstructed          bool
	SpectralValueOrdersConstructed     bool
	ReverseOrderEquallyCompatible      bool
	NonUniqueFiltrationWitnessed       bool
	CanonicalTotalOrderDerived         bool
	CanonicalOrientationDerived        bool
	CanonicalCutoffDerived             bool
	MonotonePredicateFamilyConstructed bool
	DerivedActivationPredicate         bool
	DerivedDecouplingMatchingRule      bool
	ThresholdCorrectedBetaDerived      bool
	NativeFiniteRGFunctorDerived       bool

	InvariantSafePredicate string
	FirstAscendingOpen     string
	FirstDescendingOpen    string

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
		shell, err := shellfunctor.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(shell)
	})
	return defaultValue, defaultErr
}

func Build(shell shellfunctor.Analysis) (Analysis, error) {
	if !shell.NestedProjectionFamilyConstructed || !shell.IdempotentSemilatticeDerived {
		return Analysis{}, fmt.Errorf("Gate 107 requires Gate 106 finite projection semilattice")
	}
	if shell.AdditiveSemigroupDerived || shell.NativeFiniteRGFunctorDerived {
		return Analysis{}, fmt.Errorf("Gate 107 expects Gate 106 to leave additive/native RG unproved")
	}
	if shell.HiddenObservedInputUsed || shell.PhysicalWeakAngleDerived || shell.FineStructureDerived || shell.PhysicalMassesDerived {
		return Analysis{}, fmt.Errorf("Gate 107 refuses hidden observed physical input")
	}
	if len(shell.Coarse.Thresholds.Decisions) == 0 {
		return Analysis{}, fmt.Errorf("Gate 107 requires threshold activation decisions")
	}

	modes := buildModes(shell.Coarse.Thresholds.Decisions)
	continuum, open, vacuum := countModes(modes)
	if open == 0 {
		return Analysis{}, fmt.Errorf("Gate 107 requires at least one threshold-open mode to test filtration ambiguity")
	}

	orders := buildOrderingWitnesses(openModes(modes))
	predicates := buildPredicateWitnesses(modes, orders)
	attempts := buildSelectorAttempts(orders, predicates)
	antichains := buildAntichains(modes)

	firstAsc := ""
	firstDesc := ""
	for _, o := range orders {
		if o.Name == "spectral-value ascending order" {
			firstAsc = o.FirstOpen
		}
		if o.Name == "spectral-value descending order" {
			firstDesc = o.FirstOpen
		}
	}
	reverseCompatible := firstAsc != "" && firstDesc != "" && firstAsc != firstDesc

	truth := "Gate 107 constructs several finite filtrations on the shell carrier: a status preorder and value-monotone ascending/descending total orders.  They are all compatible with the current finite data, but no finite theorem selects the orientation, cutoff, physical scale, or beta-matching rule.  Therefore the only invariant predicate is the safe status predicate: continuum candidates remain candidates, threshold-open modes remain open, and vacuum-frustration modes remain excluded.  The residual physical-flow nullity is unchanged."

	return Analysis{
		Shell:                              shell,
		Modes:                              modes,
		ModeCount:                          len(modes),
		ContinuumCount:                     continuum,
		OpenCount:                          open,
		VacuumCount:                        vacuum,
		SelectorAttempts:                   attempts,
		OrderingWitnesses:                  orders,
		PredicateWitnesses:                 predicates,
		Antichains:                         antichains,
		StatusPreorderConstructed:          true,
		SpectralValueOrdersConstructed:     len(orders) >= 3,
		ReverseOrderEquallyCompatible:      reverseCompatible,
		NonUniqueFiltrationWitnessed:       reverseCompatible && len(orders) >= 3,
		CanonicalTotalOrderDerived:         false,
		CanonicalOrientationDerived:        false,
		CanonicalCutoffDerived:             false,
		MonotonePredicateFamilyConstructed: len(predicates) >= 3,
		DerivedActivationPredicate:         false,
		DerivedDecouplingMatchingRule:      false,
		ThresholdCorrectedBetaDerived:      false,
		NativeFiniteRGFunctorDerived:       false,
		InvariantSafePredicate:             "active(mode)=continuum-candidate; open(mode)=threshold-open; excluded(mode)=vacuum-frustration-only",
		FirstAscendingOpen:                 firstAsc,
		FirstDescendingOpen:                firstDesc,
		ResidualNullityBefore:              shell.ResidualNullityAfter,
		ResidualNullityAfter:               shell.ResidualNullityAfter,
		ResidualSymmetryBroken:             false,
		PhysicalWeakAngleDerived:           false,
		FineStructureDerived:               false,
		PhysicalMassesDerived:              false,
		HiddenObservedInputUsed:            false,
		TruthStatement:                     truth,
		RejectedClaims: []string{
			"a value-sorted list of finite threshold anchors is a physical activation order",
			"ascending and descending spectral filtrations can both be canonical without an orientation theorem",
			"a monotone predicate family P_tau becomes a threshold rule before tau is selected",
			"continuum-candidate scalar modes may correct beta coefficients without an individual decoupling/matching theorem",
			"the finite shell filtration reduces the absolute coupling/scale/threshold nullity",
		},
		RemainingUnknowns: []string{
			"U-27A-ORIENTATION: derive whether finite threshold activation follows increasing or decreasing spectral value",
			"U-27B-CUTOFF: derive a finite cutoff/index tau rather than choosing one",
			"U-27C-REPRESENTATION: assign gauge representations to B-gap and contact-overlap modes",
			"U-27D-DECOUPLING: derive a matching rule from activated modes to Δb_i",
			"U-27E-SCALE: derive the dimensional or logarithmic parameter that turns a finite filtration into RG running",
			"U-27F-FLOW: derive how the gauge-action prefactor changes under the filtration",
		},
		RecommendedNextGate: "Gate 108 — threshold representation completion / finite beta-matching tensor search",
	}, nil
}

func buildModes(decisions []thresholdactivation.ActivationDecision) []FiltrationMode {
	out := make([]FiltrationMode, 0, len(decisions))
	for _, d := range decisions {
		class := OpenMode
		switch d.Status {
		case thresholdactivation.ContinuumFieldCandidate:
			class = ContinuumMode
		case thresholdactivation.VacuumFrustrationOnly:
			class = VacuumMode
		default:
			class = OpenMode
		}
		out = append(out, FiltrationMode{
			Name:       d.Assignment.Candidate.Name,
			Kind:       string(d.Assignment.Candidate.Kind),
			Status:     string(d.Status),
			RepStatus:  string(d.Assignment.Status),
			Class:      class,
			Value:      d.Assignment.Candidate.Value,
			CanCorrect: d.CanCorrectBeta,
			Detail:     d.Reason,
		})
	}
	return out
}

func countModes(modes []FiltrationMode) (continuum, open, vacuum int) {
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

func openModes(modes []FiltrationMode) []FiltrationMode {
	out := make([]FiltrationMode, 0)
	for _, m := range modes {
		if m.Class == OpenMode {
			out = append(out, m)
		}
	}
	return out
}

func buildOrderingWitnesses(open []FiltrationMode) []OrderingWitness {
	status := append([]FiltrationMode(nil), open...)
	asc := append([]FiltrationMode(nil), open...)
	desc := append([]FiltrationMode(nil), open...)
	kindFirst := append([]FiltrationMode(nil), open...)

	sort.Slice(asc, func(i, j int) bool {
		if math.Abs(asc[i].Value-asc[j].Value) < 1e-12 {
			return asc[i].Name < asc[j].Name
		}
		return asc[i].Value < asc[j].Value
	})
	sort.Slice(desc, func(i, j int) bool {
		if math.Abs(desc[i].Value-desc[j].Value) < 1e-12 {
			return desc[i].Name < desc[j].Name
		}
		return desc[i].Value > desc[j].Value
	})
	sort.Slice(kindFirst, func(i, j int) bool {
		if kindFirst[i].Kind == kindFirst[j].Kind {
			if math.Abs(kindFirst[i].Value-kindFirst[j].Value) < 1e-12 {
				return kindFirst[i].Name < kindFirst[j].Name
			}
			return kindFirst[i].Value < kindFirst[j].Value
		}
		return kindFirst[i].Kind < kindFirst[j].Kind
	})

	return []OrderingWitness{
		{
			Name:                     "status/preexisting shell order",
			OpenOrder:                names(status),
			FirstOpen:                firstName(status),
			LastOpen:                 lastName(status),
			CompatibleWithFiniteData: true,
			MonotoneInSpectralValue:  false,
			OrientationDerived:       false,
			Canonical:                false,
			SelectsThresholdPhysics:  false,
			Detail:                   "inherits Gate 106 bookkeeping order; compatible but not a theorem of physical activation",
		},
		{
			Name:                     "spectral-value ascending order",
			OpenOrder:                names(asc),
			FirstOpen:                firstName(asc),
			LastOpen:                 lastName(asc),
			CompatibleWithFiniteData: true,
			MonotoneInSpectralValue:  true,
			OrientationDerived:       false,
			Canonical:                false,
			SelectsThresholdPhysics:  false,
			Detail:                   "orders open anchors by increasing dimensionless value; no theorem says lower value activates first",
		},
		{
			Name:                     "spectral-value descending order",
			OpenOrder:                names(desc),
			FirstOpen:                firstName(desc),
			LastOpen:                 lastName(desc),
			CompatibleWithFiniteData: true,
			MonotoneInSpectralValue:  true,
			OrientationDerived:       false,
			Canonical:                false,
			SelectsThresholdPhysics:  false,
			Detail:                   "orders open anchors by decreasing dimensionless value; equally compatible absent an orientation theorem",
		},
		{
			Name:                     "kind-then-value order",
			OpenOrder:                names(kindFirst),
			FirstOpen:                firstName(kindFirst),
			LastOpen:                 lastName(kindFirst),
			CompatibleWithFiniteData: true,
			MonotoneInSpectralValue:  false,
			OrientationDerived:       false,
			Canonical:                false,
			SelectsThresholdPhysics:  false,
			Detail:                   "orders by candidate kind before value; also compatible, proving value alone is not forced",
		},
	}
}

func buildPredicateWitnesses(modes []FiltrationMode, orders []OrderingWitness) []PredicateWitness {
	_, open, _ := countModes(modes)
	return []PredicateWitness{
		{
			Name:                     "safe status predicate",
			Formula:                  "P(mode)=continuum-candidate only; threshold-open remains open; vacuum-only excluded",
			Direction:                "status",
			OpenModesActivated:       0,
			CompatibleWithFiniteData: true,
			Monotone:                 true,
			CutoffDerived:            true,
			ScaleDerived:             false,
			CanCorrectBeta:           false,
			DeltaBSelected:           false,
			Detail:                   "this is invariant under all compatible open-mode orderings, but it activates no heavy thresholds",
		},
		{
			Name:                     "ascending value cut family",
			Formula:                  "P_tau(mode)=value(mode) <= tau",
			Direction:                "ascending",
			OpenModesActivated:       open,
			CompatibleWithFiniteData: true,
			Monotone:                 true,
			CutoffDerived:            false,
			ScaleDerived:             false,
			CanCorrectBeta:           false,
			DeltaBSelected:           false,
			Detail:                   "monotone for any chosen tau, but tau and physical orientation are not selected",
		},
		{
			Name:                     "descending value cut family",
			Formula:                  "P_tau(mode)=value(mode) >= tau",
			Direction:                "descending",
			OpenModesActivated:       open,
			CompatibleWithFiniteData: true,
			Monotone:                 true,
			CutoffDerived:            false,
			ScaleDerived:             false,
			CanCorrectBeta:           false,
			DeltaBSelected:           false,
			Detail:                   "equally monotone and equally unselected; chooses a different first open mode than ascending order",
		},
		{
			Name:                     "chosen shell-index cut",
			Formula:                  "P_n(mode)=shellRank(mode) <= n",
			Direction:                "Gate-106 shell rank",
			OpenModesActivated:       open,
			CompatibleWithFiniteData: true,
			Monotone:                 true,
			CutoffDerived:            false,
			ScaleDerived:             false,
			CanCorrectBeta:           false,
			DeltaBSelected:           false,
			Detail:                   "monotone after an arbitrary shell rank is chosen; does not derive the rank itself",
		},
	}
}

func buildSelectorAttempts(orders []OrderingWitness, predicates []PredicateWitness) []SelectorAttempt {
	return []SelectorAttempt{
		{
			Name:                         "status preorder selector",
			Symbol:                       "F_status",
			Constructed:                  true,
			PreorderConstructed:          true,
			TotalOrderConstructed:        false,
			MonotonePredicateConstructed: true,
			OrientationDerived:           true,
			CutoffDerived:                true,
			PhysicalScaleDerived:         false,
			DeltaBMatchingDerived:        false,
			BetaCorrectionAllowed:        false,
			RejectedAsPhysicalSelector:   true,
			Detail:                       "safe and invariant, but it leaves threshold-open modes open and generates no threshold corrections",
		},
		{
			Name:                         "spectral-value total-order selectors",
			Symbol:                       "F_value±",
			Constructed:                  len(orders) >= 3,
			PreorderConstructed:          true,
			TotalOrderConstructed:        true,
			MonotonePredicateConstructed: true,
			OrientationDerived:           false,
			CutoffDerived:                false,
			PhysicalScaleDerived:         false,
			DeltaBMatchingDerived:        false,
			BetaCorrectionAllowed:        false,
			RejectedAsPhysicalSelector:   true,
			Detail:                       "ascending and descending orders are both compatible; no finite orientation/cutoff theorem chooses between them",
		},
		{
			Name:                         "shell-index cut predicate",
			Symbol:                       "P_n",
			Constructed:                  len(predicates) > 0,
			PreorderConstructed:          true,
			TotalOrderConstructed:        true,
			MonotonePredicateConstructed: true,
			OrientationDerived:           false,
			CutoffDerived:                false,
			PhysicalScaleDerived:         false,
			DeltaBMatchingDerived:        false,
			BetaCorrectionAllowed:        false,
			RejectedAsPhysicalSelector:   true,
			Detail:                       "a cut is monotone once n is chosen, but n is not derived and no activated shell maps to Δb_i",
		},
		{
			Name:                         "physical threshold predicate",
			Symbol:                       "P_phys",
			Constructed:                  false,
			PreorderConstructed:          false,
			TotalOrderConstructed:        false,
			MonotonePredicateConstructed: false,
			OrientationDerived:           false,
			CutoffDerived:                false,
			PhysicalScaleDerived:         false,
			DeltaBMatchingDerived:        false,
			BetaCorrectionAllowed:        false,
			RejectedAsPhysicalSelector:   true,
			Detail:                       "requires orientation, cutoff, physical/log scale, representation completion, and decoupling/matching theorem",
		},
	}
}

func buildAntichains(modes []FiltrationMode) []AntichainWitness {
	open := make([]string, 0)
	for _, m := range modes {
		if m.Class == OpenMode {
			open = append(open, m.Name)
		}
	}
	return []AntichainWitness{{
		ClassName: "threshold-open antichain",
		Members:   open,
		Reason:    "status-level data classifies these modes as open but supplies no pairwise physical activation order",
	}}
}

func names(xs []FiltrationMode) []string {
	out := make([]string, 0, len(xs))
	for _, x := range xs {
		out = append(out, x.Name)
	}
	return out
}

func firstName(xs []FiltrationMode) string {
	if len(xs) == 0 {
		return ""
	}
	return xs[0].Name
}

func lastName(xs []FiltrationMode) string {
	if len(xs) == 0 {
		return ""
	}
	return xs[len(xs)-1].Name
}

func FormatModes(xs []FiltrationMode) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s [%s/%s, rep=%s, value=%.10f, beta=%t]", x.Name, x.Kind, x.Class, x.RepStatus, x.Value, x.CanCorrect))
	}
	return strings.Join(parts, "; ")
}

func FormatOrders(xs []OrderingWitness) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		head := x.OpenOrder
		if len(head) > 3 {
			head = head[:3]
		}
		parts = append(parts, fmt.Sprintf("%s first=%q last=%q monotone-value=%t canonical=%t order-head=%v", x.Name, x.FirstOpen, x.LastOpen, x.MonotoneInSpectralValue, x.Canonical, head))
	}
	return strings.Join(parts, "; ")
}

func FormatPredicates(xs []PredicateWitness) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s: %s open-activated=%d cutoff=%t scale=%t Δb=%t", x.Name, x.Formula, x.OpenModesActivated, x.CutoffDerived, x.ScaleDerived, x.DeltaBSelected))
	}
	return strings.Join(parts, "; ")
}

func FormatAttempts(xs []SelectorAttempt) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s:%s constructed=%t total=%t monotone=%t orientation=%t cutoff=%t scale=%t Δb=%t beta=%t", x.Symbol, x.Name, x.Constructed, x.TotalOrderConstructed, x.MonotonePredicateConstructed, x.OrientationDerived, x.CutoffDerived, x.PhysicalScaleDerived, x.DeltaBMatchingDerived, x.BetaCorrectionAllowed))
	}
	return strings.Join(parts, "; ")
}

func FormatAntichains(xs []AntichainWitness) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		members := x.Members
		if len(members) > 4 {
			members = members[:4]
		}
		parts = append(parts, fmt.Sprintf("%s size=%d head=%v", x.ClassName, len(x.Members), members))
	}
	return strings.Join(parts, "; ")
}

func Join(xs []string) string { return strings.Join(xs, "; ") }
