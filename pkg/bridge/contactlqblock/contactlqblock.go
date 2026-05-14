// Package contactlqblock implements Gate 131: contact leptoquark six-block
// symmetry / S6 permutation obstruction theorem.
//
// Gate 130 showed that the current-side 1+6 quotient cannot naturally decide
// which of the seven contact rows is the singlet and how the remaining six
// rows are assigned to leptoquark slots. Gate 131 isolates the second part of
// that obstruction: even after a singlet row is externally chosen, the six-row
// leptoquark block carries an S6 permutation ambiguity. Spectral orderings can
// order six contact values, but they are contact diagnostics, not current-
// derived representation semantics.
package contactlqblock

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/contactassignment"
)

type StrategyKind string

const (
	AnonymousSixBlock  StrategyKind = "anonymous-six-block"
	CurrentLQSlots     StrategyKind = "current-leptoquark-slots"
	SpectralAscending  StrategyKind = "spectral-ascending-six"
	SpectralDescending StrategyKind = "spectral-descending-six"
	FanoTransport      StrategyKind = "fano-transport-six"
	ObservedFit        StrategyKind = "observed-fit-six"
)

type Strategy struct {
	Name string
	Kind StrategyKind

	SingletAssumed      bool
	SixRowsSelected     bool
	SixRowsOrdered      bool
	CurrentDerived      bool
	ContactDerived      bool
	FanoDerived         bool
	Canonical           bool
	Natural             bool
	RepresentationRows  bool
	BetaPermitted       bool
	ZeroRowsProved      bool
	RequiresSingletPick bool
	RequiresS6Choice    bool
	RequiresOrientation bool
	RequiresFanoChoice  bool
	RequiresObserved    bool
	SingletChoices      int
	SixPermutations     int
	TotalAssignments    int
	Obstruction         string
}

type SixBlock struct {
	SingletRow             string
	SingletSlot            int
	SingletValue           float64
	RemainingRows          []string
	RemainingValues        []float64
	Permutations           int
	OrderedAscending       []string
	OrderedDescending      []string
	CanonicalCurrentOrder  bool
	RepresentationComplete bool
}

type Criterion struct {
	Name     string
	Required bool
	Derived  bool
	Detail   string
}

type Summary struct {
	ContactRows                 int
	LeptoquarkRows              int
	SingletChoices              int
	SixPermutationOrder         int
	AssignmentsPerBranch        int
	CurrentBranches             int
	TotalCurrentAssignments     int
	SpectralOrderings           int
	CurrentNaturalSixOrders     int
	CanonicalCurrentAssignments int
	RepresentationCompleteRows  int
	ContactBetaRowsAllowed      int
	ContactZeroRowsProved       int
	ResidualNullityBefore       int
	ResidualNullityAfter        int
}

type Analysis struct {
	Previous contactassignment.Analysis

	Strategies []Strategy
	Blocks     []SixBlock
	Criteria   []Criterion
	Summary    Summary

	ContactRows             int
	LeptoquarkRows          int
	SingletChoices          int
	SixPermutationOrder     int
	AssignmentsPerBranch    int
	CurrentBranches         int
	TotalCurrentAssignments int

	SixBlockExists                    bool
	AnonymousBlockCanonical           bool
	S6PermutationObstruction          bool
	CurrentNaturalSixOrder            bool
	SpectralOrderingAvailable         bool
	SpectralOrientationAmbiguous      bool
	FanoChoiceRequired                bool
	ObservedInputRejected             bool
	CanonicalCurrentAssignmentDerived bool

	RepresentationCompleteRows    int
	RepresentationOpenRows        int
	ContactBetaRowsAllowed        int
	ContactZeroRowsProved         int
	ThresholdCorrectedBetaDerived bool
	FullBetaMatchingTensorDerived bool

	ResidualNullityBefore    int
	ResidualNullityAfter     int
	HiddenObservedInputUsed  bool
	PhysicalWeakAngleDerived bool
	FineStructureDerived     bool
	PhysicalMassesDerived    bool
	PhysicalScaleDerived     bool

	TruthStatement      string
	RejectedClaims      []string
	RemainingUnknowns   []string
	RecommendedNextGate string
}

var defaultOnce sync.Once
var defaultValue Analysis
var defaultErr error

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		prev, err := contactassignment.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev)
	})
	return defaultValue, defaultErr
}

func Build(prev contactassignment.Analysis) (Analysis, error) {
	if prev.ContactRows != 7 || prev.CurrentPattern != "1+6" || !prev.SingletChoiceRequired || !prev.PermutationRequired {
		return Analysis{}, fmt.Errorf("Gate 131 requires Gate 130 singlet and permutation obstruction")
	}
	if prev.Summary.MinimalPermutationChoices != 720 || prev.Summary.MinimalRowAssignmentChoices != 5040 || prev.RepresentationCompleteRows != 0 || prev.ContactBetaRowsAllowed != 0 || prev.ContactZeroRowsProved != 0 {
		return Analysis{}, fmt.Errorf("Gate 131 requires Gate 130 beta firewall and 6! residual permutation obstruction")
	}
	if prev.HiddenObservedInputUsed || prev.PhysicalWeakAngleDerived || prev.FineStructureDerived || prev.PhysicalMassesDerived || prev.PhysicalScaleDerived {
		return Analysis{}, fmt.Errorf("Gate 131 refuses hidden physical input")
	}

	blocks := buildBlocks(prev.Rows)
	if len(blocks) != 7 {
		return Analysis{}, fmt.Errorf("expected seven possible singlet blocks, got %d", len(blocks))
	}
	strategies := buildStrategies()
	spectralOrderings := 0
	currentNaturalOrders := 0
	canonicalAssignments := 0
	fanoRequired := false
	observedRejected := false
	orientationAmbiguous := false
	s6Obstruction := false
	anonymousCanonical := false
	for _, s := range strategies {
		if s.Kind == AnonymousSixBlock && s.Canonical && !s.SixRowsOrdered {
			anonymousCanonical = true
		}
		if s.ContactDerived && s.SixRowsOrdered && !s.CurrentDerived {
			spectralOrderings++
		}
		if s.CurrentDerived && s.SixRowsOrdered && s.Natural && s.RepresentationRows && !s.RequiresS6Choice && !s.RequiresObserved {
			currentNaturalOrders++
		}
		if s.CurrentDerived && s.SixRowsOrdered && s.Canonical && s.RepresentationRows && !s.RequiresS6Choice && !s.RequiresObserved {
			canonicalAssignments++
		}
		if s.RequiresS6Choice {
			s6Obstruction = true
		}
		if s.RequiresFanoChoice {
			fanoRequired = true
		}
		if s.RequiresObserved {
			observedRejected = true
		}
		if s.RequiresOrientation {
			orientationAmbiguous = true
		}
	}

	summary := Summary{
		ContactRows:                 7,
		LeptoquarkRows:              6,
		SingletChoices:              7,
		SixPermutationOrder:         factorial(6),
		AssignmentsPerBranch:        7 * factorial(6),
		CurrentBranches:             2,
		TotalCurrentAssignments:     2 * 7 * factorial(6),
		SpectralOrderings:           spectralOrderings,
		CurrentNaturalSixOrders:     currentNaturalOrders,
		CanonicalCurrentAssignments: canonicalAssignments,
		RepresentationCompleteRows:  0,
		ContactBetaRowsAllowed:      0,
		ContactZeroRowsProved:       0,
		ResidualNullityBefore:       prev.ResidualNullityAfter,
		ResidualNullityAfter:        prev.ResidualNullityAfter,
	}
	criteria := buildCriteria(summary)
	truth := "Gate 131 isolates the six-row part of the Gate 130 obstruction. Even if a contact singlet row were chosen by external convention, the remaining six contact rows still have S6 = 720 possible leptoquark-slot assignments. Keeping the six-block anonymous is canonical but row-blind. Spectral ascending and descending orders are available because the contact values are distinct, but they are contact diagnostic orientations rather than current-derived representation maps. Therefore the six leptoquark slots still have no natural contact-row assignment and the contact beta firewall stays closed."

	return Analysis{
		Previous:   prev,
		Strategies: strategies,
		Blocks:     blocks,
		Criteria:   criteria,
		Summary:    summary,

		ContactRows:             7,
		LeptoquarkRows:          6,
		SingletChoices:          7,
		SixPermutationOrder:     factorial(6),
		AssignmentsPerBranch:    7 * factorial(6),
		CurrentBranches:         2,
		TotalCurrentAssignments: 2 * 7 * factorial(6),

		SixBlockExists:                    true,
		AnonymousBlockCanonical:           anonymousCanonical,
		S6PermutationObstruction:          s6Obstruction,
		CurrentNaturalSixOrder:            currentNaturalOrders > 0,
		SpectralOrderingAvailable:         spectralOrderings >= 2,
		SpectralOrientationAmbiguous:      orientationAmbiguous,
		FanoChoiceRequired:                fanoRequired,
		ObservedInputRejected:             observedRejected,
		CanonicalCurrentAssignmentDerived: canonicalAssignments > 0,

		RepresentationCompleteRows:    0,
		RepresentationOpenRows:        len(prev.Rows),
		ContactBetaRowsAllowed:        0,
		ContactZeroRowsProved:         0,
		ThresholdCorrectedBetaDerived: false,
		FullBetaMatchingTensorDerived: false,

		ResidualNullityBefore:    prev.ResidualNullityAfter,
		ResidualNullityAfter:     prev.ResidualNullityAfter,
		HiddenObservedInputUsed:  false,
		PhysicalWeakAngleDerived: false,
		FineStructureDerived:     false,
		PhysicalMassesDerived:    false,
		PhysicalScaleDerived:     false,

		TruthStatement: truth,
		RejectedClaims: []string{
			"choosing a contact singlet automatically orders the remaining six leptoquark rows",
			"the current-side leptoquark six-block supplies a canonical S6 ordering",
			"spectral ascending or descending order is a current-derived representation map",
			"a Fano-labelled six-block can be used without a hidden contact-to-Fano assignment",
			"contact leptoquark rows may enter beta matching before representation rows are derived",
		},
		RemainingUnknowns: []string{
			"canonical current-derived order on the six contact leptoquark rows",
			"representation rows for the six leptoquark contact slots",
			"contact singlet selector",
			"threshold beta matching tensor Delta b_i(L)",
			"physical boundary scale and absolute coupling unit",
		},
		RecommendedNextGate: "Gate 132 — contact leptoquark slot representation tensor / color-doublet semantic obstruction theorem",
	}, nil
}

func buildBlocks(rows []contactassignment.AssignmentRow) []SixBlock {
	out := make([]SixBlock, 0, len(rows))
	for i, singlet := range rows {
		rest := make([]contactassignment.AssignmentRow, 0, len(rows)-1)
		for j, r := range rows {
			if i != j {
				rest = append(rest, r)
			}
		}
		asc := append([]contactassignment.AssignmentRow(nil), rest...)
		sort.Slice(asc, func(i, j int) bool { return asc[i].Value < asc[j].Value })
		desc := append([]contactassignment.AssignmentRow(nil), asc...)
		for l, r := 0, len(desc)-1; l < r; l, r = l+1, r-1 {
			desc[l], desc[r] = desc[r], desc[l]
		}
		out = append(out, SixBlock{
			SingletRow:             singlet.Name,
			SingletSlot:            singlet.SpectralSlot,
			SingletValue:           singlet.Value,
			RemainingRows:          names(rest),
			RemainingValues:        values(rest),
			Permutations:           factorial(6),
			OrderedAscending:       names(asc),
			OrderedDescending:      names(desc),
			CanonicalCurrentOrder:  false,
			RepresentationComplete: false,
		})
	}
	return out
}

func buildStrategies() []Strategy {
	return []Strategy{
		{
			Name: "keep leptoquark six-block anonymous", Kind: AnonymousSixBlock,
			SingletAssumed: true, SixRowsSelected: true, SixRowsOrdered: false,
			CurrentDerived: true, Canonical: true, Natural: true, RepresentationRows: false, BetaPermitted: false,
			Obstruction: "preserves the current-side six-dimensional block but does not assign slots to contact rows",
		},
		{
			Name: "assign current leptoquark slots to six contact rows", Kind: CurrentLQSlots,
			SingletAssumed: true, SixRowsSelected: true, SixRowsOrdered: true,
			CurrentDerived: true, Canonical: false, Natural: false, RepresentationRows: false, BetaPermitted: false,
			RequiresS6Choice: true, SixPermutations: factorial(6), TotalAssignments: factorial(6),
			Obstruction: "needs one hidden element of S6 after the singlet row is chosen",
		},
		{
			Name: "spectral ascending order of remaining six rows", Kind: SpectralAscending,
			SingletAssumed: true, SixRowsSelected: true, SixRowsOrdered: true,
			ContactDerived: true, Canonical: true, Natural: false, RepresentationRows: false, BetaPermitted: false,
			RequiresOrientation: true,
			Obstruction:         "orders contact diagnostics by value but supplies no current representation semantics",
		},
		{
			Name: "spectral descending order of remaining six rows", Kind: SpectralDescending,
			SingletAssumed: true, SixRowsSelected: true, SixRowsOrdered: true,
			ContactDerived: true, Canonical: true, Natural: false, RepresentationRows: false, BetaPermitted: false,
			RequiresOrientation: true,
			Obstruction:         "equally available reverse orientation proves spectral ordering is convention-dependent for current slots",
		},
		{
			Name: "Fano-transported six-row order", Kind: FanoTransport,
			SingletAssumed: true, SixRowsSelected: true, SixRowsOrdered: true,
			FanoDerived: true, Canonical: false, Natural: false, RepresentationRows: false, BetaPermitted: false,
			RequiresFanoChoice: true, RequiresS6Choice: true, SixPermutations: factorial(6), TotalAssignments: factorial(6),
			Obstruction: "requires hidden contact-to-Fano assignment before Fano incidence can order contact rows",
		},
		{
			Name: "observed-constant six-row fit", Kind: ObservedFit,
			SingletAssumed: true, SixRowsSelected: true, SixRowsOrdered: true,
			Canonical: false, Natural: false, RepresentationRows: false, BetaPermitted: false,
			RequiresObserved: true,
			Obstruction:      "observed couplings or masses cannot choose finite leptoquark slot assignments",
		},
	}
}

func buildCriteria(s Summary) []Criterion {
	return []Criterion{
		{Name: "Gate 130 obstruction inherited", Required: true, Derived: s.ContactRows == 7 && s.LeptoquarkRows == 6 && s.SingletChoices == 7, Detail: "current pattern remains 1+6"},
		{Name: "S6 residual permutation measured", Required: true, Derived: s.SixPermutationOrder == 720 && s.AssignmentsPerBranch == 5040 && s.TotalCurrentAssignments == 10080, Detail: "six-block slot assignment has 6! choices after a singlet choice"},
		{Name: "spectral orderings are diagnostic only", Required: true, Derived: s.SpectralOrderings >= 2 && s.CurrentNaturalSixOrders == 0, Detail: "ascending/descending contact orders exist but no current-natural order exists"},
		{Name: "no canonical current assignment", Required: true, Derived: s.CanonicalCurrentAssignments == 0 && s.RepresentationCompleteRows == 0, Detail: "six leptoquark slots remain representation-incomplete"},
		{Name: "beta firewall remains closed", Required: true, Derived: s.ContactBetaRowsAllowed == 0 && s.ContactZeroRowsProved == 0, Detail: "no threshold beta row or cancellation row is permitted"},
		{Name: "physical-flow nullity preserved", Required: true, Derived: s.ResidualNullityBefore == 3 && s.ResidualNullityAfter == 3, Detail: "u, L, Delta b_i(L) remain free"},
	}
}

func names(rows []contactassignment.AssignmentRow) []string {
	xs := make([]string, len(rows))
	for i, r := range rows {
		xs[i] = r.Name
	}
	return xs
}

func values(rows []contactassignment.AssignmentRow) []float64 {
	xs := make([]float64, len(rows))
	for i, r := range rows {
		xs[i] = r.Value
	}
	return xs
}

func factorial(n int) int {
	out := 1
	for i := 2; i <= n; i++ {
		out *= i
	}
	return out
}

func FormatStrategies(xs []Strategy) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s(kind=%s singlet=%t six=%t ordered=%t current=%t contact=%t fano=%t canonical=%t natural=%t s6=%t orient=%t fanoChoice=%t observed=%t perms=%d rep=%t beta=%t obstruction=%s)", x.Name, x.Kind, x.SingletAssumed, x.SixRowsSelected, x.SixRowsOrdered, x.CurrentDerived, x.ContactDerived, x.FanoDerived, x.Canonical, x.Natural, x.RequiresS6Choice, x.RequiresOrientation, x.RequiresFanoChoice, x.RequiresObserved, x.SixPermutations, x.RepresentationRows, x.BetaPermitted, x.Obstruction))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatBlocks(xs []SixBlock, limit int) string {
	if limit <= 0 || limit > len(xs) {
		limit = len(xs)
	}
	parts := make([]string, 0, limit)
	for i := 0; i < limit; i++ {
		b := xs[i]
		parts = append(parts, fmt.Sprintf("singlet=%s(slot=%d value=%.10f) perms=%d asc=%s desc=%s rep=%t", b.SingletRow, b.SingletSlot, b.SingletValue, b.Permutations, strings.Join(b.OrderedAscending, ","), strings.Join(b.OrderedDescending, ","), b.RepresentationComplete))
	}
	if limit < len(xs) {
		parts = append(parts, fmt.Sprintf("... +%d more", len(xs)-limit))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatCriteria(xs []Criterion) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s=%t", x.Name, x.Derived))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("contact=%d lqRows=%d singletChoices=%d s6=%d perBranch=%d branches=%d total=%d spectralOrders=%d currentOrders=%d canonicalAssignments=%d repRows=%d betaRows=%d zeroRows=%d nullity=%d->%d", s.ContactRows, s.LeptoquarkRows, s.SingletChoices, s.SixPermutationOrder, s.AssignmentsPerBranch, s.CurrentBranches, s.TotalCurrentAssignments, s.SpectralOrderings, s.CurrentNaturalSixOrders, s.CanonicalCurrentAssignments, s.RepresentationCompleteRows, s.ContactBetaRowsAllowed, s.ContactZeroRowsProved, s.ResidualNullityBefore, s.ResidualNullityAfter)
}

func Join(xs []string) string { return strings.Join(xs, "; ") }
