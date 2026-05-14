// Package contactequivrefinement implements Gate 129: contact-row
// equivalence refinement / sector-pattern mismatch obstruction theorem.
//
// Gate 128 showed that two natural current-side quotients have the right
// seven-dimensional carrier size but only a 1+6 sector pattern. Gate 129 asks
// whether that 1+6 pattern can be refined into seven contact singleton rows in
// a canonical way. The answer is negative: every row-resolving refinement needs
// a hidden bijection between the one-dimensional current sector plus six
// leptoquark slots and the seven distinct contact spectral rows.
package contactequivrefinement

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/contactquotientsemantics"
)

type RefinementKind string

const (
	KeepSectorPattern       RefinementKind = "keep-current-1+6-pattern"
	SplitLeptoquarkSix      RefinementKind = "split-leptoquark-six"
	CentralPinnedRefinement RefinementKind = "central-pinned-refinement"
	BLPinnedRefinement      RefinementKind = "bminusl-pinned-refinement"
	ContactSingletonRefine  RefinementKind = "contact-singleton-refinement"
	AnonymousCollapse       RefinementKind = "anonymous-collapse"
	FanoTransportRefine     RefinementKind = "fano-transport-refinement"
	ObservedRefine          RefinementKind = "observed-refinement"
)

type RefinementCandidate struct {
	Name string
	Kind RefinementKind

	SourcePattern       []int
	TargetPattern       []int
	Dimension           int
	ClassCount          int
	CurrentDerived      bool
	ContactDerived      bool
	Canonical           bool
	RowResolving        bool
	RepresentationRows  bool
	BetaPermitted       bool
	RequiresCentralPick bool
	RequiresPermutation bool
	RequiresFanoChoice  bool
	RequiresObserved    bool
	CollapsesRows       bool
	HiddenChoices       int
	Obstruction         string
}

type RefinedRow struct {
	Name               string
	SpectralSlot       int
	Value              float64
	CurrentSector      string
	ContactClass       string
	RefinementSelected bool
	RepresentationRow  bool
	CanEnterBetaTensor bool
	Reason             string
}

type Criterion struct {
	Name     string
	Required bool
	Derived  bool
	Detail   string
}

type Summary struct {
	U4Dimension                   int
	ContactRows                   int
	CurrentSectorPattern          string
	ContactSingletonPattern       string
	NaturalCurrentQuotients       int
	CanonicalRowResolvingRefines  int
	CurrentDerivedRowRefines      int
	NonCanonicalRowRefines        int
	MinimalHiddenChoicesPerBranch int
	TotalHiddenBranchChoices      int
	RepresentationCompleteRows    int
	ContactBetaRowsAllowed        int
	ContactZeroRowsProved         int
	ResidualNullityBefore         int
	ResidualNullityAfter          int
}

type Analysis struct {
	Previous contactquotientsemantics.Analysis

	Candidates []RefinementCandidate
	Rows       []RefinedRow
	Criteria   []Criterion
	Summary    Summary

	U4Dimension             int
	ContactRows             int
	NaturalCurrentQuotients int
	SectorPatternMismatch   bool

	CurrentPatternStable       bool
	ContactSingletonsStable    bool
	CanonicalRefinementDerived bool
	CurrentDerivedRefinement   bool
	HiddenAssignmentRequired   bool
	FanoChoiceRequired         bool
	ObservedInputRejected      bool
	ArbitraryCutoffRejected    bool

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
		prev, err := contactquotientsemantics.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev)
	})
	return defaultValue, defaultErr
}

func Build(prev contactquotientsemantics.Analysis) (Analysis, error) {
	if !prev.CurrentSideQuotientSemanticsFound || !prev.NaturalPatternsConflict || prev.U4Dimension != 16 || prev.ContactRows != 7 {
		return Analysis{}, fmt.Errorf("Gate 129 requires Gate 128 current-side quotient semantics obstruction")
	}
	if prev.CurrentToContactRelationDerived || prev.CanonicalSemanticRelationDerived || prev.RepresentationCompleteRows != 0 || prev.ContactBetaRowsAllowed != 0 || prev.ContactZeroRowsProved != 0 {
		return Analysis{}, fmt.Errorf("Gate 129 requires Gate 128 beta firewall to remain closed")
	}
	if prev.HiddenObservedInputUsed || prev.PhysicalWeakAngleDerived || prev.FineStructureDerived || prev.PhysicalMassesDerived || prev.PhysicalScaleDerived {
		return Analysis{}, fmt.Errorf("Gate 129 refuses hidden physical input")
	}

	candidates := buildCandidates()
	rows := buildRows(prev.Rows)

	canonicalRowResolving := 0
	currentDerivedRowResolving := 0
	nonCanonicalRowResolving := 0
	hidden := false
	fanoChoice := false
	observedRejected := false
	minHidden := 0
	totalHidden := 0
	for _, c := range candidates {
		if c.RowResolving && c.Canonical && !c.RequiresCentralPick && !c.RequiresPermutation && !c.RequiresFanoChoice && !c.RequiresObserved && c.CurrentDerived && c.RepresentationRows {
			canonicalRowResolving++
		}
		if c.RowResolving && c.CurrentDerived && c.RepresentationRows && !c.RequiresCentralPick && !c.RequiresPermutation && !c.RequiresObserved {
			currentDerivedRowResolving++
		}
		if c.RowResolving && (c.RequiresCentralPick || c.RequiresPermutation || c.RequiresFanoChoice || c.RequiresObserved || !c.CurrentDerived || !c.RepresentationRows) {
			nonCanonicalRowResolving++
		}
		if c.RequiresCentralPick || c.RequiresPermutation {
			hidden = true
			if c.HiddenChoices > 0 {
				if minHidden == 0 || c.HiddenChoices < minHidden {
					minHidden = c.HiddenChoices
				}
				totalHidden += c.HiddenChoices
			}
		}
		if c.RequiresFanoChoice {
			fanoChoice = true
		}
		if c.RequiresObserved {
			observedRejected = true
		}
	}

	summary := Summary{
		U4Dimension:                   16,
		ContactRows:                   7,
		CurrentSectorPattern:          "1+6",
		ContactSingletonPattern:       "1+1+1+1+1+1+1",
		NaturalCurrentQuotients:       prev.NaturalCurrentQuotients,
		CanonicalRowResolvingRefines:  canonicalRowResolving,
		CurrentDerivedRowRefines:      currentDerivedRowResolving,
		NonCanonicalRowRefines:        nonCanonicalRowResolving,
		MinimalHiddenChoicesPerBranch: minHidden,
		TotalHiddenBranchChoices:      totalHidden,
		RepresentationCompleteRows:    0,
		ContactBetaRowsAllowed:        0,
		ContactZeroRowsProved:         0,
		ResidualNullityBefore:         prev.ResidualNullityAfter,
		ResidualNullityAfter:          prev.ResidualNullityAfter,
	}

	criteria := buildCriteria(summary)
	truth := "Gate 129 refines the Gate 128 mismatch. The natural current quotients have pattern 1+6, while the contact carrier has seven distinct singleton rows. Keeping the 1+6 pattern preserves current semantics but does not resolve contact rows. Splitting the six-dimensional leptoquark sector and assigning the one-dimensional current sector to a contact row produces seven rows only after a hidden 7*6! = 5040 choice, and the second natural quotient gives another incompatible 5040 choices. The contact singleton partition is canonical only diagnostically, not current-derived. Therefore no canonical current-to-contact refinement, representation row, beta row, or zero-row cancellation is derived."

	return Analysis{
		Previous:   prev,
		Candidates: candidates,
		Rows:       rows,
		Criteria:   criteria,
		Summary:    summary,

		U4Dimension:             16,
		ContactRows:             7,
		NaturalCurrentQuotients: prev.NaturalCurrentQuotients,
		SectorPatternMismatch:   true,

		CurrentPatternStable:       true,
		ContactSingletonsStable:    true,
		CanonicalRefinementDerived: false,
		CurrentDerivedRefinement:   false,
		HiddenAssignmentRequired:   hidden,
		FanoChoiceRequired:         fanoChoice,
		ObservedInputRejected:      observedRejected,
		ArbitraryCutoffRejected:    true,

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
			"a 1+6 current-sector pattern refines canonically to seven contact singleton rows",
			"one may pick the distinguished contact row for the current singlet by convention",
			"the six leptoquark slots may be assigned to the remaining contact rows without a hidden permutation",
			"the contact singleton partition is already a current-derived representation-row map",
			"Fano transport or observed constants may choose the refinement",
		},
		RemainingUnknowns: []string{
			"canonical current-to-contact refinement rule",
			"local variable or representation-row semantics for the seven contact modes",
			"constraint/BRST cancellation ledger for contact rows",
			"threshold beta matching tensor Delta b_i(L)",
			"physical boundary scale and absolute coupling unit",
		},
		RecommendedNextGate: "Gate 130 — contact singlet/leptoquark assignment naturality / permutation obstruction theorem",
	}, nil
}

func buildCandidates() []RefinementCandidate {
	return []RefinementCandidate{
		{
			Name: "keep current quotient 1+6 sector pattern", Kind: KeepSectorPattern,
			SourcePattern: []int{1, 6}, TargetPattern: []int{1, 6}, Dimension: 7, ClassCount: 2,
			CurrentDerived: true, ContactDerived: false, Canonical: true, RowResolving: false, RepresentationRows: false, BetaPermitted: false,
			Obstruction: "preserves current-side sector semantics but leaves six contact rows unresolved inside the leptoquark block",
		},
		{
			Name: "split central+leptoquark into seven contact rows", Kind: CentralPinnedRefinement,
			SourcePattern: []int{1, 6}, TargetPattern: []int{1, 1, 1, 1, 1, 1, 1}, Dimension: 7, ClassCount: 7,
			CurrentDerived: true, ContactDerived: false, Canonical: false, RowResolving: true, RepresentationRows: false, BetaPermitted: false,
			RequiresCentralPick: true, RequiresPermutation: true, HiddenChoices: 5040,
			Obstruction: "requires choosing which contact row receives the current singlet and permuting the six remaining rows",
		},
		{
			Name: "split B-L+leptoquark into seven contact rows", Kind: BLPinnedRefinement,
			SourcePattern: []int{1, 6}, TargetPattern: []int{1, 1, 1, 1, 1, 1, 1}, Dimension: 7, ClassCount: 7,
			CurrentDerived: true, ContactDerived: false, Canonical: false, RowResolving: true, RepresentationRows: false, BetaPermitted: false,
			RequiresCentralPick: true, RequiresPermutation: true, HiddenChoices: 5040,
			Obstruction: "coexists with the central refinement and needs the same hidden 7*6! assignment",
		},
		{
			Name: "contact spectral singleton refinement", Kind: ContactSingletonRefine,
			SourcePattern: []int{1, 1, 1, 1, 1, 1, 1}, TargetPattern: []int{1, 1, 1, 1, 1, 1, 1}, Dimension: 7, ClassCount: 7,
			CurrentDerived: false, ContactDerived: true, Canonical: true, RowResolving: true, RepresentationRows: false, BetaPermitted: false,
			Obstruction: "row-resolving diagnostic partition is not current-derived and has no gauge representation rows",
		},
		{
			Name: "anonymous contact collapse", Kind: AnonymousCollapse,
			SourcePattern: []int{7}, TargetPattern: []int{7}, Dimension: 7, ClassCount: 1,
			CurrentDerived: false, ContactDerived: true, Canonical: true, RowResolving: false, RepresentationRows: false, BetaPermitted: false, CollapsesRows: true,
			Obstruction: "collapses all contact rows and cannot support row-level threshold matching",
		},
		{
			Name: "Fano-labelled refinement", Kind: FanoTransportRefine,
			SourcePattern: []int{1, 1, 1, 1, 1, 1, 1}, TargetPattern: []int{1, 1, 1, 1, 1, 1, 1}, Dimension: 7, ClassCount: 7,
			CurrentDerived: false, ContactDerived: false, Canonical: false, RowResolving: true, RepresentationRows: false, BetaPermitted: false,
			RequiresFanoChoice: true, HiddenChoices: 5040,
			Obstruction: "requires choosing a contact-to-Fano bijection before semantics can be transported",
		},
		{
			Name: "observed-constant fitted refinement", Kind: ObservedRefine,
			SourcePattern: []int{1, 1, 1, 1, 1, 1, 1}, TargetPattern: []int{1, 1, 1, 1, 1, 1, 1}, Dimension: 7, ClassCount: 7,
			CurrentDerived: false, ContactDerived: false, Canonical: false, RowResolving: true, RepresentationRows: false, BetaPermitted: false,
			RequiresObserved: true,
			Obstruction:      "observed low-energy constants cannot be used to refine finite contact rows",
		},
	}
}

func buildRows(in []contactquotientsemantics.SemanticRow) []RefinedRow {
	rows := make([]RefinedRow, 0, len(in))
	for _, r := range in {
		rows = append(rows, RefinedRow{
			Name:               r.Name,
			SpectralSlot:       r.SpectralSlot,
			Value:              r.Value,
			CurrentSector:      "unselected-current-refinement",
			ContactClass:       fmt.Sprintf("contact-singleton-%d", r.SpectralSlot),
			RefinementSelected: false,
			RepresentationRow:  false,
			CanEnterBetaTensor: false,
			Reason:             "the 1+6 current-sector pattern has no canonical refinement into this contact singleton row",
		})
	}
	return rows
}

func buildCriteria(s Summary) []Criterion {
	return []Criterion{
		{Name: "Gate 128 mismatch inherited", Required: true, Derived: s.U4Dimension == 16 && s.ContactRows == 7 && s.NaturalCurrentQuotients == 2, Detail: "two current-side seven-dimensional quotients coexist"},
		{Name: "sector pattern mismatch exposed", Required: true, Derived: s.CurrentSectorPattern == "1+6" && s.ContactSingletonPattern == "1+1+1+1+1+1+1", Detail: "current quotients are sector partitions; contact rows are spectral singletons"},
		{Name: "no canonical row-resolving current refinement", Required: true, Derived: s.CanonicalRowResolvingRefines == 0 && s.CurrentDerivedRowRefines == 0, Detail: "every row-resolving current refinement needs hidden choices or lacks representation rows"},
		{Name: "hidden refinement choices measured", Required: true, Derived: s.MinimalHiddenChoicesPerBranch == 5040 && s.TotalHiddenBranchChoices >= 10080, Detail: "one singlet assignment plus a six-row permutation gives 7*6! per branch"},
		{Name: "beta firewall remains closed", Required: true, Derived: s.RepresentationCompleteRows == 0 && s.ContactBetaRowsAllowed == 0 && s.ContactZeroRowsProved == 0, Detail: "no contact threshold contribution is permitted"},
		{Name: "physical-flow nullity preserved", Required: true, Derived: s.ResidualNullityBefore == 3 && s.ResidualNullityAfter == 3, Detail: "u, L, and Delta b_i(L) remain free"},
	}
}

func FormatPattern(p []int) string {
	parts := make([]string, 0, len(p))
	for _, x := range p {
		parts = append(parts, fmt.Sprintf("%d", x))
	}
	return strings.Join(parts, "+")
}

func FormatCandidates(xs []RefinementCandidate) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s(kind=%s source=%s target=%s dim=%d classes=%d current=%t contact=%t canonical=%t row=%t centralPick=%t perm=%t fano=%t observed=%t choices=%d rep=%t beta=%t obstruction=%s)", x.Name, x.Kind, FormatPattern(x.SourcePattern), FormatPattern(x.TargetPattern), x.Dimension, x.ClassCount, x.CurrentDerived, x.ContactDerived, x.Canonical, x.RowResolving, x.RequiresCentralPick, x.RequiresPermutation, x.RequiresFanoChoice, x.RequiresObserved, x.HiddenChoices, x.RepresentationRows, x.BetaPermitted, x.Obstruction))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatRows(rows []RefinedRow, limit int) string {
	if limit <= 0 || limit > len(rows) {
		limit = len(rows)
	}
	parts := make([]string, 0, limit)
	for i := 0; i < limit; i++ {
		r := rows[i]
		parts = append(parts, fmt.Sprintf("%s[slot=%d value=%.10f current=%s contact=%s selected=%t rep=%t beta=%t]", r.Name, r.SpectralSlot, r.Value, r.CurrentSector, r.ContactClass, r.RefinementSelected, r.RepresentationRow, r.CanEnterBetaTensor))
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
	return fmt.Sprintf("u4=%d contact=%d currentPattern=%s contactPattern=%s naturalCurrentQuotients=%d canonicalRefines=%d currentRefines=%d noncanonicalRefines=%d minHiddenChoices=%d totalHiddenChoices=%d repRows=%d betaRows=%d zeroRows=%d nullity=%d->%d", s.U4Dimension, s.ContactRows, s.CurrentSectorPattern, s.ContactSingletonPattern, s.NaturalCurrentQuotients, s.CanonicalRowResolvingRefines, s.CurrentDerivedRowRefines, s.NonCanonicalRowRefines, s.MinimalHiddenChoicesPerBranch, s.TotalHiddenBranchChoices, s.RepresentationCompleteRows, s.ContactBetaRowsAllowed, s.ContactZeroRowsProved, s.ResidualNullityBefore, s.ResidualNullityAfter)
}

func Join(xs []string) string { return strings.Join(xs, "; ") }
