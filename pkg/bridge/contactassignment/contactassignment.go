// Package contactassignment implements Gate 130: contact singlet/leptoquark
// assignment naturality / permutation obstruction theorem.
//
// Gate 129 proved that the current-side seven-dimensional quotients have a
// 1+6 sector pattern, while the contact carrier has seven singleton spectral
// rows. Gate 130 asks the sharper question: can the finite data naturally
// decide which contact row is the current-side singlet and which six rows are
// the leptoquark sector? The answer remains negative. Spectral conventions can
// pick a row, but they are contact diagnostics rather than current-derived
// natural assignments. Current-derived assignments still need a hidden singlet
// choice and, for row-level semantics, a hidden permutation of the six-row
// block.
package contactassignment

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/contactequivrefinement"
)

type AssignmentKind string

const (
	CurrentAnonymousOneSix AssignmentKind = "current-anonymous-1+6"
	CentralSingletChoice   AssignmentKind = "central-singlet-choice"
	BLSingletChoice        AssignmentKind = "bminusl-singlet-choice"
	SpectralMinimumChoice  AssignmentKind = "spectral-minimum-choice"
	SpectralMaximumChoice  AssignmentKind = "spectral-maximum-choice"
	SpectralMedianChoice   AssignmentKind = "spectral-median-choice"
	FanoFlagChoice         AssignmentKind = "fano-flag-choice"
	ObservedChoice         AssignmentKind = "observed-choice"
)

type AssignmentCandidate struct {
	Name string
	Kind AssignmentKind

	Branch              string
	Pattern             string
	SingletSelected     bool
	SixBlockSelected    bool
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
	RequiresPermutation bool
	RequiresOrientation bool
	RequiresFanoChoice  bool
	RequiresObserved    bool
	SingletChoices      int
	PermutationChoices  int
	TotalChoices        int
	Obstruction         string
}

type AssignmentRow struct {
	Name                string
	SpectralSlot        int
	Value               float64
	PossibleSinglet     bool
	AssignedSinglet     bool
	LeptoquarkCandidate bool
	AssignmentSelected  bool
	CurrentBranch       string
	RepresentationRow   bool
	CanEnterBetaTensor  bool
	Reason              string
}

type Criterion struct {
	Name     string
	Required bool
	Derived  bool
	Detail   string
}

type Summary struct {
	ContactRows                    int
	CurrentPattern                 string
	NaturalCurrentQuotients        int
	SpectralRowsDistinct           bool
	CurrentNaturalSingletSelectors int
	ContactDiagnosticSelectors     int
	CanonicalRowAssignments        int
	CurrentDerivedRowAssignments   int
	MinimalSingletChoices          int
	MinimalPermutationChoices      int
	MinimalRowAssignmentChoices    int
	TotalHiddenBranchChoices       int
	RepresentationCompleteRows     int
	ContactBetaRowsAllowed         int
	ContactZeroRowsProved          int
	ResidualNullityBefore          int
	ResidualNullityAfter           int
}

type Analysis struct {
	Previous contactequivrefinement.Analysis

	Candidates []AssignmentCandidate
	Rows       []AssignmentRow
	Criteria   []Criterion
	Summary    Summary

	ContactRows             int
	NaturalCurrentQuotients int
	CurrentPattern          string
	ContactSingletonPattern string
	SpectralRowsDistinct    bool
	MinSpectralRow          AssignmentRow
	MaxSpectralRow          AssignmentRow
	MedianSpectralRow       AssignmentRow

	SingletChoiceRequired        bool
	PermutationRequired          bool
	CurrentNaturalSelector       bool
	ContactDiagnosticSelectors   int
	CanonicalAssignmentDerived   bool
	CurrentDerivedAssignment     bool
	FanoChoiceRequired           bool
	ObservedInputRejected        bool
	ArbitraryOrientationRejected bool

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
		prev, err := contactequivrefinement.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev)
	})
	return defaultValue, defaultErr
}

func Build(prev contactequivrefinement.Analysis) (Analysis, error) {
	if !prev.SectorPatternMismatch || !prev.HiddenAssignmentRequired || prev.Summary.MinimalHiddenChoicesPerBranch != 5040 || prev.ContactRows != 7 {
		return Analysis{}, fmt.Errorf("Gate 130 requires Gate 129 sector-pattern mismatch and 5040-hidden-choice obstruction")
	}
	if prev.CanonicalRefinementDerived || prev.CurrentDerivedRefinement || prev.RepresentationCompleteRows != 0 || prev.ContactBetaRowsAllowed != 0 || prev.ContactZeroRowsProved != 0 {
		return Analysis{}, fmt.Errorf("Gate 130 requires Gate 129 beta firewall to remain closed")
	}
	if prev.HiddenObservedInputUsed || prev.PhysicalWeakAngleDerived || prev.FineStructureDerived || prev.PhysicalMassesDerived || prev.PhysicalScaleDerived {
		return Analysis{}, fmt.Errorf("Gate 130 refuses hidden physical input")
	}

	rows := buildRows(prev.Rows)
	if len(rows) != 7 {
		return Analysis{}, fmt.Errorf("expected seven contact rows, got %d", len(rows))
	}
	distinct := valuesDistinct(rows)
	minRow, maxRow, medianRow := spectralExtrema(rows)
	candidates := buildCandidates(distinct)

	currentNaturalSelectors := 0
	contactDiagnosticSelectors := 0
	canonicalAssignments := 0
	currentDerivedAssignments := 0
	minSingletChoices := 0
	minPermutationChoices := 0
	minRowChoices := 0
	totalHidden := 0
	singletRequired := false
	permRequired := false
	fanoRequired := false
	observedRejected := false
	orientationRejected := false
	for _, c := range candidates {
		if c.SingletSelected && c.CurrentDerived && c.Natural && c.RepresentationRows && !c.RequiresSingletPick && !c.RequiresPermutation && !c.RequiresObserved {
			currentNaturalSelectors++
		}
		if c.SingletSelected && c.ContactDerived && !c.CurrentDerived {
			contactDiagnosticSelectors++
		}
		if c.SingletSelected && c.SixRowsOrdered && c.Canonical && c.CurrentDerived && c.RepresentationRows && !c.RequiresSingletPick && !c.RequiresPermutation && !c.RequiresObserved {
			canonicalAssignments++
		}
		if c.SingletSelected && c.SixRowsOrdered && c.CurrentDerived && c.RepresentationRows && !c.RequiresSingletPick && !c.RequiresPermutation && !c.RequiresObserved {
			currentDerivedAssignments++
		}
		if c.RequiresSingletPick {
			singletRequired = true
			if c.SingletChoices > 0 && (minSingletChoices == 0 || c.SingletChoices < minSingletChoices) {
				minSingletChoices = c.SingletChoices
			}
		}
		if c.RequiresPermutation {
			permRequired = true
			if c.PermutationChoices > 0 && (minPermutationChoices == 0 || c.PermutationChoices < minPermutationChoices) {
				minPermutationChoices = c.PermutationChoices
			}
		}
		if c.RequiresSingletPick || c.RequiresPermutation || c.RequiresFanoChoice {
			if c.TotalChoices > 0 {
				totalHidden += c.TotalChoices
				if c.CurrentDerived || c.FanoDerived {
					if minRowChoices == 0 || c.TotalChoices < minRowChoices {
						minRowChoices = c.TotalChoices
					}
				}
			}
		}
		if c.RequiresFanoChoice {
			fanoRequired = true
		}
		if c.RequiresObserved {
			observedRejected = true
		}
		if c.RequiresOrientation {
			orientationRejected = true
		}
	}

	summary := Summary{
		ContactRows:                    7,
		CurrentPattern:                 "1+6",
		NaturalCurrentQuotients:        prev.NaturalCurrentQuotients,
		SpectralRowsDistinct:           distinct,
		CurrentNaturalSingletSelectors: currentNaturalSelectors,
		ContactDiagnosticSelectors:     contactDiagnosticSelectors,
		CanonicalRowAssignments:        canonicalAssignments,
		CurrentDerivedRowAssignments:   currentDerivedAssignments,
		MinimalSingletChoices:          minSingletChoices,
		MinimalPermutationChoices:      minPermutationChoices,
		MinimalRowAssignmentChoices:    minRowChoices,
		TotalHiddenBranchChoices:       totalHidden,
		RepresentationCompleteRows:     0,
		ContactBetaRowsAllowed:         0,
		ContactZeroRowsProved:          0,
		ResidualNullityBefore:          prev.ResidualNullityAfter,
		ResidualNullityAfter:           prev.ResidualNullityAfter,
	}
	criteria := buildCriteria(summary)
	truth := "Gate 130 tests the sharpest form of the 1+6 obstruction. A current-side quotient can name one singlet sector and one six-dimensional leptoquark sector, but it does not decide which of the seven contact spectral rows is the singlet. Choosing a row needs seven hidden choices; resolving the six remaining rows needs 6! more choices, so every row-level current-to-contact assignment still carries 7*6! = 5040 hidden choices per branch. Spectral minimum, maximum, or median conventions can pick contact rows diagnostically, but they are not current-derived naturality laws and do not supply gauge representation, mass activation, or beta-matching rows."

	return Analysis{
		Previous:   prev,
		Candidates: candidates,
		Rows:       rows,
		Criteria:   criteria,
		Summary:    summary,

		ContactRows:             7,
		NaturalCurrentQuotients: prev.NaturalCurrentQuotients,
		CurrentPattern:          "1+6",
		ContactSingletonPattern: "1+1+1+1+1+1+1",
		SpectralRowsDistinct:    distinct,
		MinSpectralRow:          minRow,
		MaxSpectralRow:          maxRow,
		MedianSpectralRow:       medianRow,

		SingletChoiceRequired:        singletRequired,
		PermutationRequired:          permRequired,
		CurrentNaturalSelector:       currentNaturalSelectors > 0,
		ContactDiagnosticSelectors:   contactDiagnosticSelectors,
		CanonicalAssignmentDerived:   canonicalAssignments > 0,
		CurrentDerivedAssignment:     currentDerivedAssignments > 0,
		FanoChoiceRequired:           fanoRequired,
		ObservedInputRejected:        observedRejected,
		ArbitraryOrientationRejected: orientationRejected,

		RepresentationCompleteRows:    0,
		RepresentationOpenRows:        len(rows),
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
			"the current-side singlet naturally selects one contact row",
			"the six leptoquark slots naturally order the remaining six contact rows",
			"spectral minimum, maximum, or median conventions are current-derived naturality laws",
			"the 1+6 current pattern is already a contact representation-row assignment",
			"Fano flags or observed constants may be used to choose the contact singlet",
		},
		RemainingUnknowns: []string{
			"canonical contact singlet selector",
			"canonical assignment of the six leptoquark slots to contact rows",
			"representation rows for the seven contact modes",
			"threshold beta matching tensor Delta b_i(L)",
			"physical boundary scale and absolute coupling unit",
		},
		RecommendedNextGate: "Gate 131 — contact leptoquark six-block symmetry / S6 permutation obstruction theorem",
	}, nil
}

func buildRows(in []contactequivrefinement.RefinedRow) []AssignmentRow {
	rows := make([]AssignmentRow, 0, len(in))
	for _, r := range in {
		rows = append(rows, AssignmentRow{
			Name:                r.Name,
			SpectralSlot:        r.SpectralSlot,
			Value:               r.Value,
			PossibleSinglet:     true,
			AssignedSinglet:     false,
			LeptoquarkCandidate: true,
			AssignmentSelected:  false,
			CurrentBranch:       "unselected-current-singlet/leptoquark-branch",
			RepresentationRow:   false,
			CanEnterBetaTensor:  false,
			Reason:              "no current-natural rule selects this row as the singlet or assigns it to one of six leptoquark slots",
		})
	}
	return rows
}

func buildCandidates(distinct bool) []AssignmentCandidate {
	return []AssignmentCandidate{
		{
			Name: "keep anonymous current 1+6 block", Kind: CurrentAnonymousOneSix,
			Branch: "both-current-quotient-branches", Pattern: "1+6", SingletSelected: false, SixBlockSelected: true, SixRowsOrdered: false,
			CurrentDerived: true, Canonical: true, Natural: true, RepresentationRows: false, BetaPermitted: false,
			Obstruction: "preserves the current-side sector split but does not select a contact singlet row",
		},
		{
			Name: "central current singlet assigned to a contact row", Kind: CentralSingletChoice,
			Branch: "central+leptoquark", Pattern: "1+6 -> 1+1+1+1+1+1+1", SingletSelected: true, SixBlockSelected: true, SixRowsOrdered: true,
			CurrentDerived: true, Canonical: false, Natural: false, RepresentationRows: false, BetaPermitted: false,
			RequiresSingletPick: true, RequiresPermutation: true, SingletChoices: 7, PermutationChoices: 720, TotalChoices: 5040,
			Obstruction: "needs one hidden contact-row choice for the central singlet and a hidden permutation of six leptoquark slots",
		},
		{
			Name: "B-L current singlet assigned to a contact row", Kind: BLSingletChoice,
			Branch: "B-L+leptoquark", Pattern: "1+6 -> 1+1+1+1+1+1+1", SingletSelected: true, SixBlockSelected: true, SixRowsOrdered: true,
			CurrentDerived: true, Canonical: false, Natural: false, RepresentationRows: false, BetaPermitted: false,
			RequiresSingletPick: true, RequiresPermutation: true, SingletChoices: 7, PermutationChoices: 720, TotalChoices: 5040,
			Obstruction: "coexists with the central branch and has the same hidden 7*6! assignment cost",
		},
		{
			Name: "minimum-overlap contact row as singlet", Kind: SpectralMinimumChoice,
			Branch: "contact-spectrum", Pattern: "min + remaining-six", SingletSelected: distinct, SixBlockSelected: true, SixRowsOrdered: false,
			ContactDerived: true, Canonical: distinct, Natural: false, RepresentationRows: false, BetaPermitted: false,
			RequiresOrientation: true, PermutationChoices: 720, TotalChoices: 720,
			Obstruction: "spectral-min convention is diagnostic and orientation-like; it is not current-derived and leaves six leptoquark slots unordered",
		},
		{
			Name: "maximum-overlap contact row as singlet", Kind: SpectralMaximumChoice,
			Branch: "contact-spectrum", Pattern: "max + remaining-six", SingletSelected: distinct, SixBlockSelected: true, SixRowsOrdered: false,
			ContactDerived: true, Canonical: distinct, Natural: false, RepresentationRows: false, BetaPermitted: false,
			RequiresOrientation: true, PermutationChoices: 720, TotalChoices: 720,
			Obstruction: "spectral-max convention is equally available, proving that spectral orientation is not a current-natural selector",
		},
		{
			Name: "median-overlap contact row as singlet", Kind: SpectralMedianChoice,
			Branch: "contact-spectrum", Pattern: "median + remaining-six", SingletSelected: distinct, SixBlockSelected: true, SixRowsOrdered: false,
			ContactDerived: true, Canonical: distinct, Natural: false, RepresentationRows: false, BetaPermitted: false,
			RequiresOrientation: true, PermutationChoices: 720, TotalChoices: 720,
			Obstruction: "median convention is another diagnostic row pick but supplies no current representation or beta row",
		},
		{
			Name: "Fano flag chooses singlet and leptoquark slots", Kind: FanoFlagChoice,
			Branch: "Fano-transport", Pattern: "flag + six complement", SingletSelected: true, SixBlockSelected: true, SixRowsOrdered: true,
			FanoDerived: true, Canonical: false, Natural: false, RepresentationRows: false, BetaPermitted: false,
			RequiresFanoChoice: true, RequiresSingletPick: true, RequiresPermutation: true, SingletChoices: 7, PermutationChoices: 720, TotalChoices: 5040,
			Obstruction: "requires selecting a contact-to-Fano bijection before a flag can label contact rows",
		},
		{
			Name: "observed-constant singlet assignment", Kind: ObservedChoice,
			Branch: "forbidden-physical-fit", Pattern: "fitted-row", SingletSelected: true, SixBlockSelected: true, SixRowsOrdered: true,
			Canonical: false, Natural: false, RepresentationRows: false, BetaPermitted: false,
			RequiresObserved: true,
			Obstruction:      "observed alpha/thetaW/masses cannot be used to select finite contact assignments",
		},
	}
}

func valuesDistinct(rows []AssignmentRow) bool {
	for i := range rows {
		for j := i + 1; j < len(rows); j++ {
			if math.Abs(rows[i].Value-rows[j].Value) < 1e-12 {
				return false
			}
		}
	}
	return true
}

func spectralExtrema(rows []AssignmentRow) (AssignmentRow, AssignmentRow, AssignmentRow) {
	minRow := rows[0]
	maxRow := rows[0]
	ordered := append([]AssignmentRow(nil), rows...)
	for _, r := range rows[1:] {
		if r.Value < minRow.Value {
			minRow = r
		}
		if r.Value > maxRow.Value {
			maxRow = r
		}
	}
	for i := 0; i < len(ordered); i++ {
		for j := i + 1; j < len(ordered); j++ {
			if ordered[j].Value < ordered[i].Value {
				ordered[i], ordered[j] = ordered[j], ordered[i]
			}
		}
	}
	return minRow, maxRow, ordered[len(ordered)/2]
}

func buildCriteria(s Summary) []Criterion {
	return []Criterion{
		{Name: "Gate 129 refinement obstruction inherited", Required: true, Derived: s.ContactRows == 7 && s.CurrentPattern == "1+6" && s.NaturalCurrentQuotients == 2, Detail: "two current quotient branches still coexist"},
		{Name: "contact spectral rows are distinguishable but not semantic", Required: true, Derived: s.SpectralRowsDistinct && s.ContactDiagnosticSelectors >= 3, Detail: "min/max/median contact selectors exist only as diagnostics"},
		{Name: "no current-natural singlet selector", Required: true, Derived: s.CurrentNaturalSingletSelectors == 0 && s.CanonicalRowAssignments == 0 && s.CurrentDerivedRowAssignments == 0, Detail: "current-side 1+6 data does not choose a contact row"},
		{Name: "hidden row assignment choices measured", Required: true, Derived: s.MinimalSingletChoices == 7 && s.MinimalPermutationChoices == 720 && s.MinimalRowAssignmentChoices == 5040 && s.TotalHiddenBranchChoices >= 10080, Detail: "row-level current assignments require 7*6! choices per branch; spectral conventions still leave 6!"},
		{Name: "beta firewall remains closed", Required: true, Derived: s.RepresentationCompleteRows == 0 && s.ContactBetaRowsAllowed == 0 && s.ContactZeroRowsProved == 0, Detail: "no contact threshold contribution is permitted"},
		{Name: "physical-flow nullity preserved", Required: true, Derived: s.ResidualNullityBefore == 3 && s.ResidualNullityAfter == 3, Detail: "u, L, and Delta b_i(L) remain free"},
	}
}

func FormatCandidates(xs []AssignmentCandidate) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s(kind=%s branch=%s pattern=%s singlet=%t six=%t ordered=%t current=%t contact=%t fano=%t canonical=%t natural=%t pick=%t perm=%t orient=%t fanoChoice=%t observed=%t choices=%d rep=%t beta=%t obstruction=%s)", x.Name, x.Kind, x.Branch, x.Pattern, x.SingletSelected, x.SixBlockSelected, x.SixRowsOrdered, x.CurrentDerived, x.ContactDerived, x.FanoDerived, x.Canonical, x.Natural, x.RequiresSingletPick, x.RequiresPermutation, x.RequiresOrientation, x.RequiresFanoChoice, x.RequiresObserved, x.TotalChoices, x.RepresentationRows, x.BetaPermitted, x.Obstruction))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatRows(rows []AssignmentRow, limit int) string {
	if limit <= 0 || limit > len(rows) {
		limit = len(rows)
	}
	parts := make([]string, 0, limit)
	for i := 0; i < limit; i++ {
		r := rows[i]
		parts = append(parts, fmt.Sprintf("%s[slot=%d value=%.10f possibleSinglet=%t assigned=%t lqCandidate=%t rep=%t beta=%t]", r.Name, r.SpectralSlot, r.Value, r.PossibleSinglet, r.AssignedSinglet, r.LeptoquarkCandidate, r.RepresentationRow, r.CanEnterBetaTensor))
	}
	if limit < len(rows) {
		parts = append(parts, fmt.Sprintf("... +%d more", len(rows)-limit))
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
	return fmt.Sprintf("contact=%d currentPattern=%s naturalQuotients=%d distinct=%t currentNaturalSelectors=%d diagnosticSelectors=%d canonicalAssignments=%d currentAssignments=%d minSingletChoices=%d minPermChoices=%d minRowChoices=%d totalHidden=%d repRows=%d betaRows=%d zeroRows=%d nullity=%d->%d", s.ContactRows, s.CurrentPattern, s.NaturalCurrentQuotients, s.SpectralRowsDistinct, s.CurrentNaturalSingletSelectors, s.ContactDiagnosticSelectors, s.CanonicalRowAssignments, s.CurrentDerivedRowAssignments, s.MinimalSingletChoices, s.MinimalPermutationChoices, s.MinimalRowAssignmentChoices, s.TotalHiddenBranchChoices, s.RepresentationCompleteRows, s.ContactBetaRowsAllowed, s.ContactZeroRowsProved, s.ResidualNullityBefore, s.ResidualNullityAfter)
}

func FormatExtrema(minRow, maxRow, medianRow AssignmentRow) string {
	return fmt.Sprintf("min=%s(slot=%d value=%.10f); median=%s(slot=%d value=%.10f); max=%s(slot=%d value=%.10f)", minRow.Name, minRow.SpectralSlot, minRow.Value, medianRow.Name, medianRow.SpectralSlot, medianRow.Value, maxRow.Name, maxRow.SpectralSlot, maxRow.Value)
}

func Join(xs []string) string { return strings.Join(xs, "; ") }
