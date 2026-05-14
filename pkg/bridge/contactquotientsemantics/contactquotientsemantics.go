// Package contactquotientsemantics implements Gate 128: current-side sector
// quotient semantics / contact-row equivalence relation search.
//
// Gate 127 exposed natural nine-dimensional U(4) kernels whose seven-
// dimensional quotients are current-side sector targets. Gate 128 asks whether
// those quotient targets can be interpreted as contact-row equivalence
// relations. The answer is negative: the current-side quotients have typed
// sector semantics such as 1+6, while the contact carrier has seven distinct
// spectral singleton rows. Converting one into the other requires a hidden
// choice of contact-row assignment or collapses the row-level data needed for
// beta matching.
package contactquotientsemantics

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/contactu4kernel"
)

type EquivalenceKind string

const (
	CurrentCentralLeptoquark EquivalenceKind = "current-central-leptoquark"
	CurrentBLLeptoquark      EquivalenceKind = "current-bminusl-leptoquark"
	ContactSpectralSingleton EquivalenceKind = "contact-spectral-singleton"
	ContactAnonymousOrbit    EquivalenceKind = "contact-anonymous-orbit"
	ContactFanoTransport     EquivalenceKind = "contact-fano-transport"
	SpectralThresholdCut     EquivalenceKind = "spectral-threshold-cut"
	ObservedFittedRelation   EquivalenceKind = "observed-fitted-relation"
)

type SectorQuotient struct {
	Name string
	Kind EquivalenceKind

	SourceKernel       string
	Dimension          int
	ClassCount         int
	ClassPattern       []int
	CurrentDerived     bool
	ContactDerived     bool
	Canonical          bool
	UniformRows        bool
	ContactSemantic    bool
	RequiresAssignment bool
	RequiresCutoff     bool
	UsesObserved       bool
	RepresentationRows bool
	BetaPermitted      bool
	Obstruction        string
}

type ContactEquivalence struct {
	Name string
	Kind EquivalenceKind

	Classes            int
	ClassPattern       []int
	Canonical          bool
	RowPreserving      bool
	CurrentDerived     bool
	ContactSemantic    bool
	RequiresAssignment bool
	RequiresCutoff     bool
	CollapsesRows      bool
	RepresentationRows bool
	BetaPermitted      bool
	Obstruction        string
}

type SemanticRow struct {
	Name               string
	SpectralSlot       int
	Value              float64
	ContactClass       string
	CurrentClass       string
	RelationSelected   bool
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
	U4Dimension                int
	ContactRows                int
	NaturalCurrentQuotients    int
	CurrentQuotientPatterns    []string
	CanonicalContactRelations  int
	CurrentContactRelations    int
	RepresentationCompleteRows int
	ContactBetaRowsAllowed     int
	ContactZeroRowsProved      int
	ResidualNullityBefore      int
	ResidualNullityAfter       int
}

type Analysis struct {
	Previous contactu4kernel.Analysis

	SectorQuotients  []SectorQuotient
	ContactRelations []ContactEquivalence
	Rows             []SemanticRow
	Criteria         []Criterion
	Summary          Summary

	U4Dimension             int
	ContactRows             int
	NaturalCurrentQuotients int
	NaturalPatternsConflict bool

	CurrentSideQuotientSemanticsFound bool
	ContactRowEquivalenceFound        bool
	CanonicalSemanticRelationDerived  bool
	CurrentToContactRelationDerived   bool
	RowPreservingRelationDerived      bool
	HiddenAssignmentRequired          bool
	ArbitraryCutoffRequired           bool

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
		prev, err := contactu4kernel.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev)
	})
	return defaultValue, defaultErr
}

func Build(prev contactu4kernel.Analysis) (Analysis, error) {
	if !prev.KernelProjectionNoGoDerived || !prev.CurrentSideQuotientOnly || prev.U4Dimension != 16 || prev.TargetContactRows != 7 || prev.RequiredKernelDimension != 9 {
		return Analysis{}, fmt.Errorf("Gate 128 requires Gate 127 current-side quotient-only obstruction")
	}
	if prev.CanonicalKernelDerived || prev.CanonicalQuotientRelation || prev.ContactSemanticKernelDerived || prev.RepresentationCompleteRows != 0 || prev.ContactBetaRowsAllowed != 0 || prev.ContactZeroRowsProved != 0 {
		return Analysis{}, fmt.Errorf("Gate 128 requires the Gate 127 contact beta firewall to remain closed")
	}
	if prev.HiddenObservedInputUsed || prev.PhysicalWeakAngleDerived || prev.FineStructureDerived || prev.PhysicalMassesDerived || prev.PhysicalScaleDerived {
		return Analysis{}, fmt.Errorf("Gate 128 refuses hidden physical input")
	}

	sectorQuotients := buildSectorQuotients()
	contactRelations := buildContactRelations()
	rows := buildRows(prev.Rows)

	naturalQuotients := 0
	patterns := []string{}
	for _, q := range sectorQuotients {
		if q.CurrentDerived && q.Dimension == 7 && !q.RequiresAssignment && !q.RequiresCutoff && !q.UsesObserved {
			naturalQuotients++
			patterns = append(patterns, FormatPattern(q.ClassPattern))
		}
	}

	canonicalContactRelations := 0
	currentContactRelations := 0
	rowPreserving := false
	hiddenAssignment := false
	cutoff := false
	for _, r := range contactRelations {
		if r.Canonical && r.RowPreserving && !r.CurrentDerived {
			canonicalContactRelations++
		}
		if r.CurrentDerived && r.ContactSemantic && r.RepresentationRows {
			currentContactRelations++
		}
		if r.RowPreserving {
			rowPreserving = true
		}
		if r.RequiresAssignment {
			hiddenAssignment = true
		}
		if r.RequiresCutoff {
			cutoff = true
		}
	}
	for _, q := range sectorQuotients {
		if q.RequiresAssignment {
			hiddenAssignment = true
		}
		if q.RequiresCutoff {
			cutoff = true
		}
	}

	summary := Summary{
		U4Dimension:                16,
		ContactRows:                7,
		NaturalCurrentQuotients:    naturalQuotients,
		CurrentQuotientPatterns:    patterns,
		CanonicalContactRelations:  canonicalContactRelations,
		CurrentContactRelations:    currentContactRelations,
		RepresentationCompleteRows: 0,
		ContactBetaRowsAllowed:     0,
		ContactZeroRowsProved:      0,
		ResidualNullityBefore:      prev.ResidualNullityAfter,
		ResidualNullityAfter:       prev.ResidualNullityAfter,
	}
	criteria := buildCriteria(summary)

	truth := "Gate 128 tests whether the current-side sector quotients exposed by Gate 127 can be promoted to contact-row equivalence relations. Two natural seven-dimensional quotients exist, but their semantics are sector patterns 1+6 rather than seven contact spectral singleton rows. The canonical contact relation preserves seven singletons, but it is not current-derived; the anonymous one-orbit relation is canonical only by forgetting row data. Fano transport and spectral cutoff relations require hidden assignments or arbitrary thresholds. Therefore no current-to-contact equivalence relation, representation row, beta row, or contact zero-row cancellation is derived."

	return Analysis{
		Previous:         prev,
		SectorQuotients:  sectorQuotients,
		ContactRelations: contactRelations,
		Rows:             rows,
		Criteria:         criteria,
		Summary:          summary,

		U4Dimension:             16,
		ContactRows:             7,
		NaturalCurrentQuotients: naturalQuotients,
		NaturalPatternsConflict: naturalQuotients >= 2,

		CurrentSideQuotientSemanticsFound: naturalQuotients >= 2,
		ContactRowEquivalenceFound:        canonicalContactRelations > 0,
		CanonicalSemanticRelationDerived:  false,
		CurrentToContactRelationDerived:   false,
		RowPreservingRelationDerived:      rowPreserving,
		HiddenAssignmentRequired:          hiddenAssignment,
		ArbitraryCutoffRequired:           cutoff,

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
			"a current quotient with dimension seven is automatically a contact-row equivalence relation",
			"central+leptoquark or B-L+leptoquark semantics label the seven contact rows",
			"the singleton contact partition is a current-derived representation map",
			"an anonymous seven-mode orbit preserves the row-level data needed for beta matching",
			"spectral cutoff or observed constants may select contact threshold rows",
		},
		RemainingUnknowns: []string{
			"current-to-contact equivalence relation with row semantics",
			"canonical local or constraint semantic attached to each contact row",
			"representation-complete contact rows",
			"mass activation and decoupling rules for contact thresholds",
			"physical boundary scale and absolute coupling unit",
		},
		RecommendedNextGate: "Gate 129 — contact-row equivalence refinement / sector-pattern mismatch obstruction theorem",
	}, nil
}

func buildRows(in []contactu4kernel.KernelRow) []SemanticRow {
	rows := make([]SemanticRow, 0, len(in))
	for _, r := range in {
		rows = append(rows, SemanticRow{
			Name:               r.Name,
			SpectralSlot:       r.SpectralSlot,
			Value:              r.Value,
			ContactClass:       fmt.Sprintf("contact-singleton-%d", r.SpectralSlot),
			CurrentClass:       "unassigned-current-class",
			RelationSelected:   false,
			RepresentationRow:  false,
			CanEnterBetaTensor: false,
			Reason:             "no current-side sector quotient has been promoted to a contact-row semantic equivalence relation",
		})
	}
	return rows
}

func buildSectorQuotients() []SectorQuotient {
	return []SectorQuotient{
		{
			Name: "u(4)/(color su(3)+B-L) = central + leptoquark", Kind: CurrentCentralLeptoquark,
			SourceKernel: "color su(3)+B-L", Dimension: 7, ClassCount: 2, ClassPattern: []int{1, 6},
			CurrentDerived: true, ContactDerived: false, Canonical: false, UniformRows: false, ContactSemantic: false, RepresentationRows: false, BetaPermitted: false,
			Obstruction: "the quotient is a typed current-sector target with pattern 1+6, not seven contact row semantics",
		},
		{
			Name: "u(4)/(central+color su(3)) = B-L + leptoquark", Kind: CurrentBLLeptoquark,
			SourceKernel: "central+color su(3)", Dimension: 7, ClassCount: 2, ClassPattern: []int{1, 6},
			CurrentDerived: true, ContactDerived: false, Canonical: false, UniformRows: false, ContactSemantic: false, RepresentationRows: false, BetaPermitted: false,
			Obstruction: "another typed current-sector target with pattern 1+6 coexists, so no canonical contact semantics is selected",
		},
		{
			Name: "contact spectral singleton partition", Kind: ContactSpectralSingleton,
			SourceKernel: "contact spectrum", Dimension: 7, ClassCount: 7, ClassPattern: []int{1, 1, 1, 1, 1, 1, 1},
			CurrentDerived: false, ContactDerived: true, Canonical: true, UniformRows: true, ContactSemantic: false, RepresentationRows: false, BetaPermitted: false,
			Obstruction: "it preserves contact rows but is not a current-derived representation or threshold relation",
		},
		{
			Name: "anonymous all-contact orbit", Kind: ContactAnonymousOrbit,
			SourceKernel: "forgotten row labels", Dimension: 7, ClassCount: 1, ClassPattern: []int{7},
			CurrentDerived: false, ContactDerived: true, Canonical: true, UniformRows: false, ContactSemantic: false, RepresentationRows: false, BetaPermitted: false,
			Obstruction: "it restores symmetry only by collapsing the row-level spectral data",
		},
		{
			Name: "Fano-transported contact relation", Kind: ContactFanoTransport,
			SourceKernel: "Fano assignment", Dimension: 7, ClassCount: 7, ClassPattern: []int{1, 1, 1, 1, 1, 1, 1},
			CurrentDerived: false, ContactDerived: false, Canonical: false, UniformRows: true, ContactSemantic: false, RequiresAssignment: true, RepresentationRows: false, BetaPermitted: false,
			Obstruction: "requires choosing one of 7! contact-to-Fano assignments",
		},
		{
			Name: "observed-fitted contact relation", Kind: ObservedFittedRelation,
			SourceKernel: "observed constants", Dimension: 7, ClassCount: 7, ClassPattern: []int{1, 1, 1, 1, 1, 1, 1},
			CurrentDerived: false, ContactDerived: false, Canonical: false, UniformRows: true, ContactSemantic: false, UsesObserved: true, RepresentationRows: false, BetaPermitted: false,
			Obstruction: "observed low-energy physics cannot select finite contact semantics",
		},
	}
}

func buildContactRelations() []ContactEquivalence {
	return []ContactEquivalence{
		{
			Name: "weighted contact singleton equivalence", Kind: ContactSpectralSingleton,
			Classes: 7, ClassPattern: []int{1, 1, 1, 1, 1, 1, 1},
			Canonical: true, RowPreserving: true, CurrentDerived: false, ContactSemantic: false, RepresentationRows: false, BetaPermitted: false,
			Obstruction: "row-preserving diagnostic partition lacks gauge representation, locality, mass activation, and current semantics",
		},
		{
			Name: "anonymous contact orbit equivalence", Kind: ContactAnonymousOrbit,
			Classes: 1, ClassPattern: []int{7},
			Canonical: true, RowPreserving: false, CurrentDerived: false, ContactSemantic: false, CollapsesRows: true, RepresentationRows: false, BetaPermitted: false,
			Obstruction: "collapses all seven rows and cannot produce row-level beta data",
		},
		{
			Name: "central/leptoquark current pullback", Kind: CurrentCentralLeptoquark,
			Classes: 2, ClassPattern: []int{1, 6},
			Canonical: false, RowPreserving: false, CurrentDerived: true, ContactSemantic: false, RequiresAssignment: true, RepresentationRows: false, BetaPermitted: false,
			Obstruction: "needs a map from the 1+6 current pattern to the seven contact rows",
		},
		{
			Name: "B-L/leptoquark current pullback", Kind: CurrentBLLeptoquark,
			Classes: 2, ClassPattern: []int{1, 6},
			Canonical: false, RowPreserving: false, CurrentDerived: true, ContactSemantic: false, RequiresAssignment: true, RepresentationRows: false, BetaPermitted: false,
			Obstruction: "coexists with the other 1+6 pullback and still lacks contact row semantics",
		},
		{
			Name: "spectral cutoff equivalence", Kind: SpectralThresholdCut,
			Classes: 2, ClassPattern: []int{3, 4},
			Canonical: false, RowPreserving: false, CurrentDerived: false, ContactSemantic: false, RequiresCutoff: true, RepresentationRows: false, BetaPermitted: false,
			Obstruction: "requires an arbitrary threshold/cutoff on the contact spectrum",
		},
		{
			Name: "Fano-labelled equivalence", Kind: ContactFanoTransport,
			Classes: 7, ClassPattern: []int{1, 1, 1, 1, 1, 1, 1},
			Canonical: false, RowPreserving: true, CurrentDerived: false, ContactSemantic: false, RequiresAssignment: true, RepresentationRows: false, BetaPermitted: false,
			Obstruction: "requires a hidden contact-to-Fano bijection",
		},
	}
}

func buildCriteria(s Summary) []Criterion {
	return []Criterion{
		{Name: "Gate 127 quotient-only obstruction inherited", Required: true, Derived: s.U4Dimension == 16 && s.ContactRows == 7, Detail: "u(4) current carrier remains sixteen-dimensional; contact target has seven rows"},
		{Name: "natural current-side seven-dimensional quotients exposed", Required: true, Derived: s.NaturalCurrentQuotients == 2, Detail: "two 1+6 current-sector quotients coexist"},
		{Name: "canonical contact equivalence exists only diagnostically", Required: true, Derived: s.CanonicalContactRelations == 1, Detail: "the singleton contact relation is row-preserving but not current-derived"},
		{Name: "current-derived contact semantic relation absent", Required: true, Derived: s.CurrentContactRelations == 0, Detail: "no current quotient gives representation-complete contact rows"},
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

func FormatSectorQuotients(xs []SectorQuotient) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s(kind=%s dim=%d classes=%d pattern=%s current=%t contact=%t canonical=%t assignment=%t observed=%t rep=%t beta=%t obstruction=%s)", x.Name, x.Kind, x.Dimension, x.ClassCount, FormatPattern(x.ClassPattern), x.CurrentDerived, x.ContactDerived, x.Canonical, x.RequiresAssignment, x.UsesObserved, x.RepresentationRows, x.BetaPermitted, x.Obstruction))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatContactRelations(xs []ContactEquivalence) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s(kind=%s classes=%d pattern=%s canonical=%t rowPreserving=%t current=%t assignment=%t cutoff=%t collapse=%t rep=%t beta=%t obstruction=%s)", x.Name, x.Kind, x.Classes, FormatPattern(x.ClassPattern), x.Canonical, x.RowPreserving, x.CurrentDerived, x.RequiresAssignment, x.RequiresCutoff, x.CollapsesRows, x.RepresentationRows, x.BetaPermitted, x.Obstruction))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatRows(rows []SemanticRow, limit int) string {
	if limit <= 0 || limit > len(rows) {
		limit = len(rows)
	}
	parts := make([]string, 0, limit)
	for i := 0; i < limit; i++ {
		r := rows[i]
		parts = append(parts, fmt.Sprintf("%s[slot=%d value=%.10f contact=%s current=%s selected=%t rep=%t beta=%t]", r.Name, r.SpectralSlot, r.Value, r.ContactClass, r.CurrentClass, r.RelationSelected, r.RepresentationRow, r.CanEnterBetaTensor))
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
	return fmt.Sprintf("u4=%d contact=%d naturalCurrentQuotients=%d patterns=%s canonicalContactRelations=%d currentContactRelations=%d repRows=%d betaRows=%d zeroRows=%d nullity=%d->%d", s.U4Dimension, s.ContactRows, s.NaturalCurrentQuotients, strings.Join(s.CurrentQuotientPatterns, ","), s.CanonicalContactRelations, s.CurrentContactRelations, s.RepresentationCompleteRows, s.ContactBetaRowsAllowed, s.ContactZeroRowsProved, s.ResidualNullityBefore, s.ResidualNullityAfter)
}

func Join(xs []string) string { return strings.Join(xs, "; ") }
