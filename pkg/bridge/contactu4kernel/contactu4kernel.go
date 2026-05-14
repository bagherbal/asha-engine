// Package contactu4kernel implements Gate 127: u(4) projection kernel /
// canonical quotient relation search.
//
// Gate 126 showed that rank-seven maps from the sixteen-dimensional
// u(4)/Pati-Salam current carrier to the seven contact rows exist abstractly,
// but no finite projection is selected. Gate 127 searches one layer deeper:
// perhaps a nine-dimensional kernel or quotient relation is already canonical.
// The result is again an obstruction. Natural nine-dimensional sector kernels
// can be named, but they correspond to current-side matter-sector quotients
// rather than contact-row semantics, and there are multiple incompatible
// natural kernels. A generic kernel is a point in Gr(9,16), with a 63-parameter
// choice. No finite action, source, relation, or symmetry criterion selects one.
package contactu4kernel

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/contactu4projection"
)

type KernelKind string

const (
	GenericGrassmannKernel   KernelKind = "generic-grassmann-kernel"
	ColorBLKernel            KernelKind = "color-plus-bminusl-kernel"
	CentralColorKernel       KernelKind = "central-plus-color-kernel"
	ColorGeneratorKernel     KernelKind = "single-color-generator-kernel"
	ContactEWKernel          KernelKind = "contact-ew-kernel"
	SpectralOrthogonalKernel KernelKind = "spectral-orthogonal-kernel"
	FittedKernel             KernelKind = "observed-fitted-kernel"
)

type Sector struct {
	Name      string
	Dimension int
	Role      string
}

type KernelCandidate struct {
	Name string
	Kind KernelKind

	SourceDimension          int
	KernelDimension          int
	QuotientDimension        int
	Constructed              bool
	DimensionCorrect         bool
	Canonical                bool
	CurrentDerived           bool
	ContactDerived           bool
	QuotientRelationDerived  bool
	RequiresChoice           bool
	UsesObserved             bool
	WrongSemantics           bool
	ContinuousFreeParameters int
	IncompatibleWith         string
	Obstruction              string
}

type QuotientRelationCandidate struct {
	Name string

	EquivalenceClasses     int
	KernelDimension        int
	Canonical              bool
	ContactSemantic        bool
	RepresentationComplete bool
	BetaPermitted          bool
	Obstruction            string
}

type KernelRow struct {
	Name               string
	SpectralSlot       int
	Value              float64
	KernelSelected     bool
	QuotientClass      string
	RepresentationRow  bool
	CanEnterBetaTensor bool
	Reason             string
}

type KernelCriterion struct {
	Name     string
	Required bool
	Derived  bool
	Detail   string
}

type KernelSummary struct {
	U4Dimension                 int
	TargetDimension             int
	RequiredKernelDimension     int
	GenericKernelsExist         bool
	GenericKernelFreeParameters int
	NaturalDimensionNineKernels int
	CanonicalContactKernels     int
	QuotientRelationsDerived    int
	RepresentationCompleteRows  int
	ContactBetaRowsAllowed      int
	ContactZeroRowsProved       int
	ResidualNullityBefore       int
	ResidualNullityAfter        int
}

type Analysis struct {
	Previous contactu4projection.Analysis

	Sectors            []Sector
	KernelCandidates   []KernelCandidate
	QuotientCandidates []QuotientRelationCandidate
	Rows               []KernelRow
	Criteria           []KernelCriterion
	Summary            KernelSummary

	U4Dimension              int
	TargetContactRows        int
	RequiredKernelDimension  int
	GenericKernelsExist      bool
	GrassmannKernelDimension int

	NaturalNineDimensionalKernels int
	SectorKernelAmbiguity         bool
	ColorBLKernelDimension        int
	CentralColorKernelDimension   int
	SectorKernelsWrongSemantics   bool

	CanonicalKernelDerived       bool
	CanonicalQuotientRelation    bool
	KernelProjectionNoGoDerived  bool
	ContactSemanticKernelDerived bool
	CurrentSideQuotientOnly      bool

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
		prev, err := contactu4projection.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev)
	})
	return defaultValue, defaultErr
}

func Build(prev contactu4projection.Analysis) (Analysis, error) {
	if !prev.SevenRowProjectionNoGoDerived || prev.U4CurrentDimension != 16 || prev.TargetContactRows != 7 || !prev.RankSevenLinearMapsExist || prev.GenericKernelDimension != 9 {
		return Analysis{}, fmt.Errorf("Gate 127 requires Gate 126 rank-seven u(4)->contact projection no-go")
	}
	if prev.U4ToContactProjectionDerived || prev.U4ToContactQuotientDerived || prev.NaturalSevenRowProjectionDerived || prev.RepresentationCompleteRows != 0 || prev.ContactBetaRowsAllowed != 0 || prev.ContactZeroRowsProved != 0 {
		return Analysis{}, fmt.Errorf("Gate 127 requires the Gate 126 contact beta firewall to remain closed")
	}
	if prev.HiddenObservedInputUsed || prev.PhysicalWeakAngleDerived || prev.FineStructureDerived || prev.PhysicalMassesDerived || prev.PhysicalScaleDerived {
		return Analysis{}, fmt.Errorf("Gate 127 refuses hidden physical input")
	}

	sectors := []Sector{
		{Name: "central-current", Dimension: 1, Role: "u(1) central current slot"},
		{Name: "color-su3", Dimension: 8, Role: "su(3)c/color current sector"},
		{Name: "b-minus-l", Dimension: 1, Role: "B-L abelian current slot"},
		{Name: "leptoquark", Dimension: 6, Role: "off-diagonal Pati-Salam/leptoquark current sector"},
	}
	kernels := buildKernelCandidates()
	quotientCandidates := buildQuotientCandidates()
	rows := buildRows(prev.Rows)

	naturalKernels := 0
	canonicalContactKernels := 0
	quotientRelations := 0
	for _, k := range kernels {
		if k.DimensionCorrect && k.CurrentDerived && !k.RequiresChoice {
			naturalKernels++
		}
		if k.DimensionCorrect && k.Canonical && k.ContactDerived && k.QuotientRelationDerived && !k.WrongSemantics {
			canonicalContactKernels++
		}
	}
	for _, q := range quotientCandidates {
		if q.Canonical && q.ContactSemantic && q.RepresentationComplete {
			quotientRelations++
		}
	}

	summary := KernelSummary{
		U4Dimension:                 16,
		TargetDimension:             7,
		RequiredKernelDimension:     9,
		GenericKernelsExist:         true,
		GenericKernelFreeParameters: 63,
		NaturalDimensionNineKernels: naturalKernels,
		CanonicalContactKernels:     canonicalContactKernels,
		QuotientRelationsDerived:    quotientRelations,
		RepresentationCompleteRows:  0,
		ContactBetaRowsAllowed:      0,
		ContactZeroRowsProved:       0,
		ResidualNullityBefore:       prev.ResidualNullityAfter,
		ResidualNullityAfter:        prev.ResidualNullityAfter,
	}
	criteria := buildCriteria(summary)

	truth := "Gate 127 searches beneath the abstract u(4)->contact projection problem for a canonical nine-dimensional kernel or quotient relation. Generic nine-dimensional kernels in a sixteen-dimensional carrier exist as a 63-parameter Grassmannian choice. Two sector-natural nine-dimensional kernels can be named: color+B-L, whose quotient is central+leptoquark, and central+color, whose quotient is B-L+leptoquark. Both are current-side Pati-Salam sector quotients, not contact-row semantics, and their coexistence proves non-uniqueness rather than selection. No contact-derived equivalence relation, representation row, source functional, finite action, or symmetry criterion selects a kernel. Therefore no u(4)->contact quotient, contact beta row, or zero-row cancellation is derived."

	return Analysis{
		Previous:           prev,
		Sectors:            sectors,
		KernelCandidates:   kernels,
		QuotientCandidates: quotientCandidates,
		Rows:               rows,
		Criteria:           criteria,
		Summary:            summary,

		U4Dimension:              16,
		TargetContactRows:        7,
		RequiredKernelDimension:  9,
		GenericKernelsExist:      true,
		GrassmannKernelDimension: 63,

		NaturalNineDimensionalKernels: naturalKernels,
		SectorKernelAmbiguity:         naturalKernels >= 2,
		ColorBLKernelDimension:        9,
		CentralColorKernelDimension:   9,
		SectorKernelsWrongSemantics:   true,

		CanonicalKernelDerived:       false,
		CanonicalQuotientRelation:    false,
		KernelProjectionNoGoDerived:  true,
		ContactSemanticKernelDerived: false,
		CurrentSideQuotientOnly:      true,

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
			"a generic nine-dimensional kernel in u(4) is a derived contact quotient relation",
			"color+B-L or central+color kernels select contact-row semantics",
			"coexisting natural sector kernels remove projection ambiguity",
			"a quotient target of dimension seven is enough to permit contact beta rows",
			"observed constants may select the u(4)->contact kernel",
		},
		RemainingUnknowns: []string{
			"finite source/action selecting a contact-semantic u(4) kernel",
			"canonical equivalence relation relating current sectors to contact rows",
			"representation-complete contact rows",
			"mass activation and decoupling rules for any contact threshold",
			"physical boundary scale and absolute coupling unit",
		},
		RecommendedNextGate: "Gate 128 — current-side sector quotient semantics / contact-row equivalence relation search",
	}, nil
}

func buildRows(in []contactu4projection.ProjectionRow) []KernelRow {
	rows := make([]KernelRow, 0, len(in))
	for _, r := range in {
		rows = append(rows, KernelRow{
			Name:               r.Name,
			SpectralSlot:       r.SpectralSlot,
			Value:              r.Value,
			KernelSelected:     false,
			QuotientClass:      "open-contact-row",
			RepresentationRow:  false,
			CanEnterBetaTensor: false,
			Reason:             "no canonical u(4) kernel or quotient relation assigns this contact row a current-derived representation semantic",
		})
	}
	return rows
}

func buildKernelCandidates() []KernelCandidate {
	return []KernelCandidate{
		{
			Name: "generic nine-dimensional kernel in u(4)", Kind: GenericGrassmannKernel,
			SourceDimension: 16, KernelDimension: 9, QuotientDimension: 7,
			Constructed: true, DimensionCorrect: true, Canonical: false, CurrentDerived: false, ContactDerived: false, QuotientRelationDerived: false,
			RequiresChoice: true, ContinuousFreeParameters: 63,
			Obstruction: "a point of Gr(9,16) must be chosen; no finite selector chooses it",
		},
		{
			Name: "kernel = color su(3) plus B-L", Kind: ColorBLKernel,
			SourceDimension: 16, KernelDimension: 9, QuotientDimension: 7,
			Constructed: true, DimensionCorrect: true, Canonical: false, CurrentDerived: true, ContactDerived: false, QuotientRelationDerived: true,
			WrongSemantics: true, IncompatibleWith: "central-plus-color kernel",
			Obstruction: "quotient is central+leptoquark, a current-side sector sum rather than seven contact rows",
		},
		{
			Name: "kernel = central plus color su(3)", Kind: CentralColorKernel,
			SourceDimension: 16, KernelDimension: 9, QuotientDimension: 7,
			Constructed: true, DimensionCorrect: true, Canonical: false, CurrentDerived: true, ContactDerived: false, QuotientRelationDerived: true,
			WrongSemantics: true, IncompatibleWith: "color-plus-B-L kernel",
			Obstruction: "quotient is B-L+leptoquark, another current-side sector sum rather than contact semantics",
		},
		{
			Name: "single color-generator kernel for su(3)->seven", Kind: ColorGeneratorKernel,
			SourceDimension: 8, KernelDimension: 1, QuotientDimension: 7,
			Constructed: false, DimensionCorrect: false, Canonical: false, CurrentDerived: true, ContactDerived: false, QuotientRelationDerived: false,
			RequiresChoice: true, WrongSemantics: true,
			Obstruction: "requires choosing one color generator/direction and is not a u(4) nine-dimensional kernel",
		},
		{
			Name: "contact electroweak four-row kernel complement", Kind: ContactEWKernel,
			SourceDimension: 16, KernelDimension: 12, QuotientDimension: 4,
			Constructed: true, DimensionCorrect: false, Canonical: true, CurrentDerived: false, ContactDerived: true, QuotientRelationDerived: false,
			Obstruction: "the derived contact/electroweak block is four-dimensional, not a seven-row contact quotient",
		},
		{
			Name: "orthogonal complement of spectral contact R^7", Kind: SpectralOrthogonalKernel,
			SourceDimension: 16, KernelDimension: 9, QuotientDimension: 7,
			Constructed: false, DimensionCorrect: true, Canonical: false, CurrentDerived: false, ContactDerived: false, QuotientRelationDerived: false,
			RequiresChoice: true,
			Obstruction:    "requires an embedding of spectral R^7 into u(4), which is exactly the missing map",
		},
		{
			Name: "observed-fitted kernel", Kind: FittedKernel,
			SourceDimension: 16, KernelDimension: 9, QuotientDimension: 7,
			Constructed: false, DimensionCorrect: true, Canonical: false, CurrentDerived: false, ContactDerived: false, QuotientRelationDerived: false,
			UsesObserved: true,
			Obstruction:  "physical constants or masses cannot be used to choose the finite kernel",
		},
	}
}

func buildQuotientCandidates() []QuotientRelationCandidate {
	return []QuotientRelationCandidate{
		{
			Name: "sector quotient u(4)/(color+B-L)", EquivalenceClasses: 7, KernelDimension: 9,
			Canonical: false, ContactSemantic: false, RepresentationComplete: false, BetaPermitted: false,
			Obstruction: "target is central+leptoquark current sector, not contact representation rows",
		},
		{
			Name: "sector quotient u(4)/(central+color)", EquivalenceClasses: 7, KernelDimension: 9,
			Canonical: false, ContactSemantic: false, RepresentationComplete: false, BetaPermitted: false,
			Obstruction: "target is B-L+leptoquark current sector, not contact representation rows",
		},
		{
			Name: "contact spectral equivalence relation", EquivalenceClasses: 7, KernelDimension: 9,
			Canonical: false, ContactSemantic: false, RepresentationComplete: false, BetaPermitted: false,
			Obstruction: "would require an already selected u(4)->contact embedding to define the kernel",
		},
		{
			Name: "observed threshold equivalence relation", EquivalenceClasses: 7, KernelDimension: 9,
			Canonical: false, ContactSemantic: false, RepresentationComplete: false, BetaPermitted: false,
			Obstruction: "forbidden because it imports observed low-energy physics",
		},
	}
}

func buildCriteria(s KernelSummary) []KernelCriterion {
	return []KernelCriterion{
		{Name: "Gate 126 projection no-go inherited", Required: true, Derived: s.U4Dimension == 16 && s.TargetDimension == 7 && s.RequiredKernelDimension == 9, Detail: "rank-seven quotient would need a nine-dimensional kernel"},
		{Name: "generic kernels exist", Required: true, Derived: s.GenericKernelsExist && s.GenericKernelFreeParameters == 63, Detail: "Gr(9,16) has 9*(16-9)=63 dimensions"},
		{Name: "natural sector kernels exposed", Required: true, Derived: s.NaturalDimensionNineKernels >= 2, Detail: "multiple current-side nine-dimensional kernels exist"},
		{Name: "canonical contact kernel selected", Required: true, Derived: s.CanonicalContactKernels > 0, Detail: "must remain false without contact semantics"},
		{Name: "beta firewall remains closed", Required: true, Derived: s.RepresentationCompleteRows == 0 && s.ContactBetaRowsAllowed == 0 && s.ContactZeroRowsProved == 0, Detail: "no contact threshold contribution is permitted"},
		{Name: "physical-flow nullity preserved", Required: true, Derived: s.ResidualNullityBefore == 3 && s.ResidualNullityAfter == 3, Detail: "u, L, and Delta b_i(L) remain free"},
	}
}

func FormatSectors(xs []Sector) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s(dim=%d role=%s)", x.Name, x.Dimension, x.Role))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatKernels(xs []KernelCandidate) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s(kind=%s kernel=%d quotient=%d dimensionOK=%t canonical=%t current=%t contact=%t quotientRel=%t choice=%t observed=%t wrongSemantics=%t obstruction=%s)", x.Name, x.Kind, x.KernelDimension, x.QuotientDimension, x.DimensionCorrect, x.Canonical, x.CurrentDerived, x.ContactDerived, x.QuotientRelationDerived, x.RequiresChoice, x.UsesObserved, x.WrongSemantics, x.Obstruction))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatQuotients(xs []QuotientRelationCandidate) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s(classes=%d kernel=%d canonical=%t contact=%t rep=%t beta=%t obstruction=%s)", x.Name, x.EquivalenceClasses, x.KernelDimension, x.Canonical, x.ContactSemantic, x.RepresentationComplete, x.BetaPermitted, x.Obstruction))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatRows(rows []KernelRow, limit int) string {
	if limit <= 0 || limit > len(rows) {
		limit = len(rows)
	}
	parts := make([]string, 0, limit)
	for i := 0; i < limit; i++ {
		r := rows[i]
		parts = append(parts, fmt.Sprintf("%s[slot=%d value=%.10f kernel=%t rep=%t beta=%t]", r.Name, r.SpectralSlot, r.Value, r.KernelSelected, r.RepresentationRow, r.CanEnterBetaTensor))
	}
	if limit < len(rows) {
		parts = append(parts, fmt.Sprintf("... +%d more", len(rows)-limit))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatCriteria(xs []KernelCriterion) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s=%t", x.Name, x.Derived))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatSummary(s KernelSummary) string {
	return fmt.Sprintf("u4=%d target=%d kernel=%d generic=%t free=%d natural9=%d canonicalContact=%d quotientRelations=%d repRows=%d betaRows=%d zeroRows=%d nullity=%d->%d", s.U4Dimension, s.TargetDimension, s.RequiredKernelDimension, s.GenericKernelsExist, s.GenericKernelFreeParameters, s.NaturalDimensionNineKernels, s.CanonicalContactKernels, s.QuotientRelationsDerived, s.RepresentationCompleteRows, s.ContactBetaRowsAllowed, s.ContactZeroRowsProved, s.ResidualNullityBefore, s.ResidualNullityAfter)
}

func Join(xs []string) string { return strings.Join(xs, "; ") }
