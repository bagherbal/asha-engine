// Package contactu4projection implements Gate 126: contact seven-row target
// projection / u(4)-to-contact quotient obstruction theorem.
//
// Gate 125 showed that the project can name seven-row contact carriers
// (spectral R^7, anonymous contact R^7, Fano R^7), but none is a derived
// dual-current target. Gate 126 tests the most tempting repair: project or
// quotient the typed sixteen-dimensional u(4)/Pati-Salam current carrier down
// to seven contact rows. The result is again an obstruction. Rank-seven linear
// maps 16 -> 7 exist abstractly, and several dimension-seven sector sums can be
// named, but no finite action, source functional, representation rule, quotient
// relation, or naturality condition selects one as the contact-row target.
package contactu4projection

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/contactdualcurrenttarget"
)

type ProjectionKind string

const (
	ArbitraryRankSevenProjection ProjectionKind = "arbitrary-rank-seven-projection"
	SectorSumCentralLepto        ProjectionKind = "sector-sum-central-plus-leptoquark"
	SectorSumBLLepto             ProjectionKind = "sector-sum-bminusl-plus-leptoquark"
	ColorQuotientProjection      ProjectionKind = "color-su3-eight-to-seven-quotient"
	ContactEWPlusThreeProjection ProjectionKind = "contact-ew-four-plus-three-extension"
	SpectralSevenProjection      ProjectionKind = "spectral-seven-identification"
	FanoSevenProjection          ProjectionKind = "fano-seven-identification"
	ObservedFittedProjection     ProjectionKind = "observed-fitted-projection"
)

type Sector struct {
	Name      string
	Dimension int
	Role      string
}

type ProjectionCandidate struct {
	Name string
	Kind ProjectionKind

	SourceDimension          int
	TargetDimension          int
	KernelDimension          int
	Constructed              bool
	RankSeven                bool
	Canonical                bool
	CurrentDerived           bool
	QuotientDerived          bool
	SectorNatural            bool
	RequiresChoice           bool
	UsesObserved             bool
	WrongSemantics           bool
	ContinuousFreeParameters int
	RepresentationRows       int
	BetaRowsAllowed          int
	ZeroRowsProved           int
	Obstruction              string
}

type ProjectionRow struct {
	Name                string
	SpectralSlot        int
	Value               float64
	U4ProjectionRow     bool
	SelectedQuotientRow bool
	RepresentationRow   bool
	CanEnterBetaTensor  bool
	Reason              string
}

type ProjectionCriterion struct {
	Name     string
	Required bool
	Derived  bool
	Detail   string
}

type ProjectionSummary struct {
	ContactRows                    int
	U4Dimension                    int
	TargetDimension                int
	AbstractRankSevenMapsExist     bool
	CanonicalRankSevenMaps         int
	SectorDimensionSevenCandidates int
	CurrentDerivedSevenTargets     int
	RepresentationCompleteRows     int
	ContactBetaRowsAllowed         int
	ContactZeroRowsProved          int
	ResidualNullityBefore          int
	ResidualNullityAfter           int
}

type Analysis struct {
	Previous contactdualcurrenttarget.Analysis

	Sectors    []Sector
	Rows       []ProjectionRow
	Candidates []ProjectionCandidate
	Criteria   []ProjectionCriterion
	Summary    ProjectionSummary

	ContactRows              int
	OpenContactRowsInherited int
	OpenContactRowsAfter     int

	U4CurrentDimension                 int
	U4SectorDimensions                 []int
	U4DecompositionCanonical           bool
	TargetContactRows                  int
	RankSevenLinearMapsExist           bool
	GenericKernelDimension             int
	CanonicalProjectionCount           int
	ContinuousProjectionFamily         bool
	ContinuousProjectionFreeParameters int

	CentralPlusLeptoDimension              int
	BLPlusLeptoDimension                   int
	DimensionSevenSectorSums               int
	DimensionSevenSectorSumsCanonical      bool
	DimensionSevenSectorSumsWrongSemantics bool

	ColorEightToSevenQuotientDerived bool
	ContactEWFourPlusThreeDerived    bool
	SpectralSevenIsU4Quotient        bool
	FanoSevenIsU4Quotient            bool
	ObservedProjectionRejected       bool

	U4ToContactProjectionDerived     bool
	U4ToContactQuotientDerived       bool
	NaturalSevenRowProjectionDerived bool
	SevenRowProjectionNoGoDerived    bool

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
		prev, err := contactdualcurrenttarget.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev)
	})
	return defaultValue, defaultErr
}

func Build(prev contactdualcurrenttarget.Analysis) (Analysis, error) {
	if !prev.SevenRowTargetNoGoDerived || prev.ContactRows != 7 || prev.OpenContactRowsAfter != 7 || prev.ContactBetaRowsAllowed != 0 || prev.ContactZeroRowsProved != 0 {
		return Analysis{}, fmt.Errorf("Gate 126 requires Gate 125 seven-row target no-go with seven open contact rows")
	}
	if prev.PatiSalamTargetDimension != 16 || !prev.PatiSalamTargetDerived || prev.PatiSalamTargetSevenRows {
		return Analysis{}, fmt.Errorf("Gate 126 requires the typed 16-dimensional u(4)/Pati-Salam current inventory and no selected seven-row target")
	}
	if prev.DualCurrentTargetDerived || prev.Summary.CurrentDerivedSevenCarriers != 0 || prev.RepresentationCompleteRows != 0 || prev.ResidualNullityAfter != 3 {
		return Analysis{}, fmt.Errorf("Gate 126 requires no current-derived seven-row target and unchanged physical-flow nullity")
	}
	if prev.HiddenObservedInputUsed || prev.PhysicalWeakAngleDerived || prev.FineStructureDerived || prev.PhysicalMassesDerived || prev.PhysicalScaleDerived {
		return Analysis{}, fmt.Errorf("Gate 126 refuses hidden physical input")
	}

	sectors := []Sector{
		{Name: "central-current", Dimension: 1, Role: "u(1) central current slot"},
		{Name: "color-su3", Dimension: 8, Role: "su(3)c/color current sector"},
		{Name: "b-minus-l", Dimension: 1, Role: "B-L abelian current slot"},
		{Name: "leptoquark", Dimension: 6, Role: "off-diagonal Pati-Salam/leptoquark current sector"},
	}
	rows := buildRows(prev.Rows)
	candidates := buildCandidates()
	canon := 0
	sectorSeven := 0
	currentDerived := 0
	for _, c := range candidates {
		if c.RankSeven && c.Canonical && c.CurrentDerived && c.QuotientDerived {
			canon++
		}
		if c.RankSeven && c.SectorNatural && c.TargetDimension == 7 {
			sectorSeven++
		}
		if c.RankSeven && c.CurrentDerived && c.QuotientDerived && !c.WrongSemantics {
			currentDerived++
		}
	}

	summary := ProjectionSummary{
		ContactRows:                    len(rows),
		U4Dimension:                    16,
		TargetDimension:                7,
		AbstractRankSevenMapsExist:     true,
		CanonicalRankSevenMaps:         canon,
		SectorDimensionSevenCandidates: sectorSeven,
		CurrentDerivedSevenTargets:     currentDerived,
		RepresentationCompleteRows:     0,
		ContactBetaRowsAllowed:         0,
		ContactZeroRowsProved:          0,
		ResidualNullityBefore:          prev.ResidualNullityAfter,
		ResidualNullityAfter:           prev.ResidualNullityAfter,
	}
	criteria := buildCriteria(summary)

	truth := "Gate 126 tests whether the sixteen-dimensional u(4)/Pati-Salam current carrier can be projected or quotiented into the seven unresolved contact rows. A generic rank-seven linear map from u(4) to a contact R^7 exists abstractly and has a nine-dimensional kernel, but the project contains no finite action, quotient relation, naturality condition, source functional, or representation rule selecting such a map. Dimension-seven sector sums such as central+leptoquark or B-L+leptoquark can be named, but they are not contact-row semantics and are not canonical target projections. Color su(3) has dimension eight, and quotienting it to seven would require choosing a generator or direction to remove. Therefore no u(4)->contact seven-row projection, representation row, beta row, or zero-row cancellation is derived."

	return Analysis{
		Previous:   prev,
		Sectors:    sectors,
		Rows:       rows,
		Candidates: candidates,
		Criteria:   criteria,
		Summary:    summary,

		ContactRows:              len(rows),
		OpenContactRowsInherited: prev.OpenContactRowsAfter,
		OpenContactRowsAfter:     len(rows),

		U4CurrentDimension:                 16,
		U4SectorDimensions:                 []int{1, 8, 1, 6},
		U4DecompositionCanonical:           true,
		TargetContactRows:                  7,
		RankSevenLinearMapsExist:           true,
		GenericKernelDimension:             9,
		CanonicalProjectionCount:           canon,
		ContinuousProjectionFamily:         true,
		ContinuousProjectionFreeParameters: 63,

		CentralPlusLeptoDimension:              7,
		BLPlusLeptoDimension:                   7,
		DimensionSevenSectorSums:               sectorSeven,
		DimensionSevenSectorSumsCanonical:      false,
		DimensionSevenSectorSumsWrongSemantics: true,

		ColorEightToSevenQuotientDerived: false,
		ContactEWFourPlusThreeDerived:    false,
		SpectralSevenIsU4Quotient:        false,
		FanoSevenIsU4Quotient:            false,
		ObservedProjectionRejected:       true,

		U4ToContactProjectionDerived:     false,
		U4ToContactQuotientDerived:       false,
		NaturalSevenRowProjectionDerived: false,
		SevenRowProjectionNoGoDerived:    true,

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
			"an arbitrary 16->7 rank-seven map is a derived contact projection",
			"central+leptoquark or B-L+leptoquark sector sums are contact representation rows",
			"color su(3) can be canonically quotiented from eight to seven contact modes",
			"spectral or Fano seven-row carriers are u(4)-derived quotient targets",
			"contact rows may enter threshold beta matching before a selected projection and representation rows are derived",
		},
		RemainingUnknowns: []string{
			"finite source/action selecting a u(4)->contact projection",
			"canonical quotient relation or equivalence relation on u(4) currents",
			"contact representation rows and mass activation rules",
			"decoupling tensor for threshold beta matching",
			"physical boundary scale and absolute coupling unit",
		},
		RecommendedNextGate: "Gate 127 — u(4) projection kernel / canonical quotient relation search",
	}, nil
}

func buildRows(in []contactdualcurrenttarget.CarrierRow) []ProjectionRow {
	rows := make([]ProjectionRow, 0, len(in))
	for _, r := range in {
		rows = append(rows, ProjectionRow{
			Name:                r.Name,
			SpectralSlot:        r.SpectralSlot,
			Value:               r.Value,
			U4ProjectionRow:     false,
			SelectedQuotientRow: false,
			RepresentationRow:   false,
			CanEnterBetaTensor:  false,
			Reason:              "row remains a contact spectral slot; no selected u(4) projection, quotient relation, representation row, mass activation, or decoupling rule is derived",
		})
	}
	return rows
}

func buildCandidates() []ProjectionCandidate {
	return []ProjectionCandidate{
		{
			Name: "generic rank-seven linear map u(4)->R^7_contact", Kind: ArbitraryRankSevenProjection,
			SourceDimension: 16, TargetDimension: 7, KernelDimension: 9,
			Constructed: true, RankSeven: true, Canonical: false, CurrentDerived: false, QuotientDerived: false,
			RequiresChoice: true, ContinuousFreeParameters: 63,
			Obstruction: "rank-seven maps exist, but choosing one is a continuous coefficient choice not selected by finite data",
		},
		{
			Name: "central current plus leptoquark sector", Kind: SectorSumCentralLepto,
			SourceDimension: 16, TargetDimension: 7, KernelDimension: 9,
			Constructed: true, RankSeven: true, Canonical: false, CurrentDerived: true, QuotientDerived: false, SectorNatural: true, WrongSemantics: true,
			Obstruction: "dimension 1+6=7, but this is a Pati-Salam sector sum, not contact-row semantics",
		},
		{
			Name: "B-L current plus leptoquark sector", Kind: SectorSumBLLepto,
			SourceDimension: 16, TargetDimension: 7, KernelDimension: 9,
			Constructed: true, RankSeven: true, Canonical: false, CurrentDerived: true, QuotientDerived: false, SectorNatural: true, WrongSemantics: true,
			Obstruction: "dimension 1+6=7, but it selects a matter-current subspace rather than seven contact modes",
		},
		{
			Name: "color su(3) eight-to-seven quotient", Kind: ColorQuotientProjection,
			SourceDimension: 8, TargetDimension: 7, KernelDimension: 1,
			Constructed: false, RankSeven: false, Canonical: false, CurrentDerived: true, QuotientDerived: false, RequiresChoice: true, WrongSemantics: true,
			Obstruction: "quotienting su(3) from eight to seven requires choosing a generator or direction to remove",
		},
		{
			Name: "contact EW four plus arbitrary three extension", Kind: ContactEWPlusThreeProjection,
			SourceDimension: 16, TargetDimension: 7, KernelDimension: 9,
			Constructed: false, RankSeven: false, Canonical: false, CurrentDerived: false, QuotientDerived: false, RequiresChoice: true,
			Obstruction: "the derived contact block is four-dimensional; no canonical three extra current rows are selected",
		},
		{
			Name: "spectral R^7 as u(4) quotient", Kind: SpectralSevenProjection,
			SourceDimension: 16, TargetDimension: 7, KernelDimension: 9,
			Constructed: true, RankSeven: true, Canonical: true, CurrentDerived: false, QuotientDerived: false,
			Obstruction: "spectral R^7 is contact diagnostic storage, not a quotient of u(4) currents",
		},
		{
			Name: "Fano R^7 as u(4) quotient", Kind: FanoSevenProjection,
			SourceDimension: 16, TargetDimension: 7, KernelDimension: 9,
			Constructed: true, RankSeven: true, Canonical: false, CurrentDerived: false, QuotientDerived: false, RequiresChoice: true,
			Obstruction: "Fano R^7 requires contact-to-Fano assignment and is not derived from u(4) current quotienting",
		},
		{
			Name: "observed-fitted u(4)->contact projection", Kind: ObservedFittedProjection,
			SourceDimension: 16, TargetDimension: 7, KernelDimension: 9,
			Constructed: false, RankSeven: true, Canonical: false, CurrentDerived: false, QuotientDerived: false, UsesObserved: true,
			Obstruction: "observed constants or masses cannot select the finite projection",
		},
	}
}

func buildCriteria(s ProjectionSummary) []ProjectionCriterion {
	return []ProjectionCriterion{
		{Name: "u(4) current inventory available", Required: true, Derived: s.U4Dimension == 16, Detail: "typed current carrier has 1+8+1+6 dimensions"},
		{Name: "abstract rank-seven maps exist", Required: true, Derived: s.AbstractRankSevenMapsExist, Detail: "linear algebra permits 16->7 maps with 9D kernels"},
		{Name: "canonical current-derived seven-row projection", Required: true, Derived: s.CanonicalRankSevenMaps > 0 && s.CurrentDerivedSevenTargets > 0, Detail: "must be false under current finite data"},
		{Name: "representation-complete contact rows", Required: true, Derived: s.RepresentationCompleteRows > 0, Detail: "must be false: beta firewall stays closed"},
		{Name: "residual physical-flow nullity preserved", Required: true, Derived: s.ResidualNullityBefore == 3 && s.ResidualNullityAfter == 3, Detail: "u, L, and Delta b_i(L) remain free"},
	}
}

func FormatSectors(xs []Sector) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s(dim=%d role=%s)", x.Name, x.Dimension, x.Role))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatRows(rows []ProjectionRow, limit int) string {
	if limit <= 0 || limit > len(rows) {
		limit = len(rows)
	}
	parts := make([]string, 0, limit)
	for i := 0; i < limit; i++ {
		r := rows[i]
		parts = append(parts, fmt.Sprintf("%s[slot=%d value=%.10f u4=%t rep=%t beta=%t]", r.Name, r.SpectralSlot, r.Value, r.U4ProjectionRow, r.RepresentationRow, r.CanEnterBetaTensor))
	}
	if limit < len(rows) {
		parts = append(parts, fmt.Sprintf("... +%d more", len(rows)-limit))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatCandidates(xs []ProjectionCandidate) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s(kind=%s source=%d target=%d kernel=%d rank7=%t canonical=%t current=%t quotient=%t choice=%t observed=%t obstruction=%s)", x.Name, x.Kind, x.SourceDimension, x.TargetDimension, x.KernelDimension, x.RankSeven, x.Canonical, x.CurrentDerived, x.QuotientDerived, x.RequiresChoice, x.UsesObserved, x.Obstruction))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatCriteria(xs []ProjectionCriterion) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s=%t", x.Name, x.Derived))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatSummary(s ProjectionSummary) string {
	return fmt.Sprintf("contactRows=%d u4=%d target=%d abstractMaps=%t canonicalMaps=%d sectorSeven=%d currentDerivedSeven=%d repRows=%d betaRows=%d zeroRows=%d nullity=%d->%d", s.ContactRows, s.U4Dimension, s.TargetDimension, s.AbstractRankSevenMapsExist, s.CanonicalRankSevenMaps, s.SectorDimensionSevenCandidates, s.CurrentDerivedSevenTargets, s.RepresentationCompleteRows, s.ContactBetaRowsAllowed, s.ContactZeroRowsProved, s.ResidualNullityBefore, s.ResidualNullityAfter)
}

func Join(xs []string) string { return strings.Join(xs, "; ") }
