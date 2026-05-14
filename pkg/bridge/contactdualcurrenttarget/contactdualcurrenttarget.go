// Package contactdualcurrenttarget implements Gate 125: contact dual-current
// target enlargement / seven-row carrier search.
//
// Gate 124 showed that the available source-current dual pairings are either
// canonical but row-blind, or row-distinguishing only as diagonal spectral
// diagnostics. Gate 125 asks the next structural question: can the current/dual
// target be enlarged so that the seven unresolved contact rows acquire a true
// seven-dimensional semantic carrier? The answer is still no. Several 7-row
// carriers can be written down (spectral R^7, anonymous contact R^7, Fano R^7),
// but none is a derived dual-current target with a source functional, local
// variables, gauge representation rows, mass activation, and decoupling rules.
package contactdualcurrenttarget

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/contactdualpairing"
)

type TargetKind string

const (
	UniformScalarTarget       TargetKind = "uniform-scalar-target"
	ContactEWBlockTarget      TargetKind = "contact-ew-block-target"
	PatiSalamCurrentTarget    TargetKind = "pati-salam-current-target"
	LeptoquarkSixTarget       TargetKind = "leptoquark-six-target"
	SpectralSevenTarget       TargetKind = "spectral-seven-target"
	FanoSevenTarget           TargetKind = "fano-seven-target"
	AnonymousSevenTarget      TargetKind = "anonymous-seven-target"
	ObservedFittedSevenTarget TargetKind = "observed-fitted-seven-target"
)

type TargetCandidate struct {
	Name string
	Kind TargetKind

	Dimension                 int
	Constructed               bool
	CanonicalUnderCurrentData bool
	SevenRowCarrier           bool
	RowDistinguishing         bool
	CurrentDerived            bool
	SourceFunctionalDerived   bool
	DualCurrentTargetDerived  bool
	RequiresHiddenChoice      bool
	HiddenChoices             int
	UsesObservedInput         bool
	RepresentationRowsDerived int
	BetaRowsPermitted         int
	ZeroRowsProved            int
	RejectedAsPremature       bool
	Obstruction               string
}

type CarrierRow struct {
	Name, ModeKind string
	Value          float64
	SpectralSlot   int

	InSpectralSevenCarrier bool
	InFanoSevenCarrier     bool
	InAnonymousCarrier     bool
	CurrentTargetRow       bool
	SourceFunctionalRow    bool
	RepresentationRow      bool
	CanEnterBetaTensor     bool
	ZeroRowProved          bool
	Status                 string
	Reason                 string
}

type TargetCriterion struct {
	Name     string
	Required bool
	Derived  bool
	Detail   string
}

type TargetSummary struct {
	ContactRows                 int
	DistinctSpectralRows        int
	CandidateSevenCarriers      int
	CanonicalSevenCarriers      int
	CurrentDerivedSevenCarriers int
	RepresentationCompleteRows  int
	ContactBetaRowsAllowed      int
	ContactZeroRowsProved       int
	HiddenChoices               int
	ResidualNullityBefore       int
	ResidualNullityAfter        int
}

type Analysis struct {
	Previous contactdualpairing.Analysis

	Rows       []CarrierRow
	Candidates []TargetCandidate
	Criteria   []TargetCriterion
	Summary    TargetSummary

	ContactRows              int
	OpenContactRowsInherited int
	OpenContactRowsAfter     int

	TargetEnlargementAttempted bool

	UniformTargetDimension int
	UniformTargetCanonical bool
	UniformTargetRowBlind  bool

	ContactEWTargetDimension int
	ContactEWTargetDerived   bool
	ContactEWTargetSevenRows bool

	PatiSalamTargetDimension int
	PatiSalamTargetDerived   bool
	PatiSalamTargetSevenRows bool

	LeptoquarkTargetDimension int
	LeptoquarkTargetDerived   bool
	LeptoquarkTargetSevenRows bool

	SpectralSevenTargetConstructed       bool
	SpectralSevenTargetCanonical         bool
	SpectralSevenTargetRowsDistinguished int
	SpectralSevenTargetCurrentDerived    bool
	SpectralSevenTargetSemantic          bool

	FanoSevenTargetConstructed      bool
	FanoSevenTargetCanonical        bool
	FanoSevenTargetRequiresChoice   bool
	FanoSevenTargetHiddenChoices    int
	FanoSevenTargetCurrentDerived   bool
	AnonymousSevenTargetConstructed bool
	AnonymousSevenTargetCanonical   bool
	AnonymousSevenTargetRowSemantic bool

	DualCurrentTargetDerived     bool
	SevenRowTargetNoGoDerived    bool
	NaturalSevenRowLabelsDerived bool

	RepresentationCompleteRows    int
	RepresentationOpenRows        int
	ContactBetaRowsAllowed        int
	ContactZeroRowsProved         int
	ThresholdCorrectedBetaDerived bool
	FullBetaMatchingTensorDerived bool

	ResidualNullityBefore   int
	ResidualNullityAfter    int
	ResidualSymmetryBroken  bool
	HiddenObservedInputUsed bool

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
		prev, err := contactdualpairing.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev, 1e-10)
	})
	return defaultValue, defaultErr
}

func Build(prev contactdualpairing.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if !prev.ContactDualPairingNoGoDerived || prev.ContactRows != 7 || prev.OpenContactRowsAfter != 7 || prev.ContactBetaRowsAllowed != 0 || prev.ContactZeroRowsProved != 0 {
		return Analysis{}, fmt.Errorf("Gate 125 requires Gate 124 dual-pairing no-go with seven open contact rows")
	}
	if !prev.SpectralPairingConstructed || prev.SpectralPairingRowsDistinguished != 7 || prev.SpectralPairingAddsSemantics || !prev.SpectralPairingDiagnosticOnly {
		return Analysis{}, fmt.Errorf("Gate 125 requires Gate 124 spectral pairing to be diagnostic-only")
	}
	if prev.CurrentToContactMapDerived || prev.CurrentFunctionalDerived || prev.SourceFunctionalDerived || prev.CurrentDualPairingDerived || prev.CurrentDualRowsDerived != 0 {
		return Analysis{}, fmt.Errorf("Gate 125 requires current-to-contact dual pairing to remain unselected")
	}
	if prev.ResidualNullityAfter != 3 || prev.HiddenObservedInputUsed || prev.PhysicalWeakAngleDerived || prev.FineStructureDerived || prev.PhysicalMassesDerived || prev.PhysicalScaleDerived {
		return Analysis{}, fmt.Errorf("Gate 125 refuses hidden physical input or changed residual nullity")
	}

	rows := buildRows(prev.Rows, eps)
	distinct := countDistinct(rows, eps)
	candidates := buildCandidates(prev, distinct)
	candidateSeven := 0
	canonicalSeven := 0
	currentDerivedSeven := 0
	for _, c := range candidates {
		if c.SevenRowCarrier {
			candidateSeven++
		}
		if c.SevenRowCarrier && c.CanonicalUnderCurrentData {
			canonicalSeven++
		}
		if c.SevenRowCarrier && c.CurrentDerived && c.DualCurrentTargetDerived {
			currentDerivedSeven++
		}
	}
	summary := TargetSummary{
		ContactRows:                 len(rows),
		DistinctSpectralRows:        distinct,
		CandidateSevenCarriers:      candidateSeven,
		CanonicalSevenCarriers:      canonicalSeven,
		CurrentDerivedSevenCarriers: currentDerivedSeven,
		RepresentationCompleteRows:  0,
		ContactBetaRowsAllowed:      0,
		ContactZeroRowsProved:       0,
		HiddenChoices:               prev.HiddenFanoChoices,
		ResidualNullityBefore:       prev.ResidualNullityAfter,
		ResidualNullityAfter:        prev.ResidualNullityAfter,
	}
	criteria := buildCriteria(summary)

	truth := "Gate 125 enlarges the possible current/dual target search to ask whether any current-side or source-side object supplies a genuine seven-row contact carrier. The uniform scalar target is canonical but one-dimensional and row-blind; the electroweak contact block target is derived but four-dimensional; the u(4)/Pati-Salam current target is sixteen-dimensional and lacks a selected projection to seven; the leptoquark sector is six-dimensional and not a contact-row target. Three seven-row carriers can be named: the spectral diagonal R^7, the anonymous contact R^7, and the Fano R^7. But the spectral and anonymous carriers are diagnostic/row-storage carriers rather than current-derived dual targets, while the Fano carrier requires one of 7! hidden contact-to-Fano assignments. Therefore no seven-row dual-current target, representation row, beta row, or zero-row cancellation is derived."

	return Analysis{
		Previous:   prev,
		Rows:       rows,
		Candidates: candidates,
		Criteria:   criteria,
		Summary:    summary,

		ContactRows:              len(rows),
		OpenContactRowsInherited: prev.OpenContactRowsAfter,
		OpenContactRowsAfter:     len(rows),

		TargetEnlargementAttempted: true,

		UniformTargetDimension: 1,
		UniformTargetCanonical: true,
		UniformTargetRowBlind:  true,

		ContactEWTargetDimension: 4,
		ContactEWTargetDerived:   true,
		ContactEWTargetSevenRows: false,

		PatiSalamTargetDimension: 16,
		PatiSalamTargetDerived:   true,
		PatiSalamTargetSevenRows: false,

		LeptoquarkTargetDimension: 6,
		LeptoquarkTargetDerived:   true,
		LeptoquarkTargetSevenRows: false,

		SpectralSevenTargetConstructed:       true,
		SpectralSevenTargetCanonical:         true,
		SpectralSevenTargetRowsDistinguished: distinct,
		SpectralSevenTargetCurrentDerived:    false,
		SpectralSevenTargetSemantic:          false,

		FanoSevenTargetConstructed:      true,
		FanoSevenTargetCanonical:        false,
		FanoSevenTargetRequiresChoice:   true,
		FanoSevenTargetHiddenChoices:    prev.HiddenFanoChoices,
		FanoSevenTargetCurrentDerived:   false,
		AnonymousSevenTargetConstructed: true,
		AnonymousSevenTargetCanonical:   true,
		AnonymousSevenTargetRowSemantic: false,

		DualCurrentTargetDerived:     false,
		SevenRowTargetNoGoDerived:    true,
		NaturalSevenRowLabelsDerived: false,

		RepresentationCompleteRows:    0,
		RepresentationOpenRows:        len(rows),
		ContactBetaRowsAllowed:        0,
		ContactZeroRowsProved:         0,
		ThresholdCorrectedBetaDerived: false,
		FullBetaMatchingTensorDerived: false,

		ResidualNullityBefore:   prev.ResidualNullityAfter,
		ResidualNullityAfter:    prev.ResidualNullityAfter,
		ResidualSymmetryBroken:  false,
		HiddenObservedInputUsed: false,

		PhysicalWeakAngleDerived: false,
		FineStructureDerived:     false,
		PhysicalMassesDerived:    false,
		PhysicalScaleDerived:     false,

		TruthStatement: truth,
		RejectedClaims: []string{
			"seven distinct contact spectral slots are physical current target rows",
			"Fano R^7 is a contact dual-current target without choosing a contact-to-Fano bijection",
			"u(4) or Pati-Salam current inventory naturally projects to the seven contact rows",
			"contact rows may enter threshold beta matching before representation and decoupling data are derived",
		},
		RemainingUnknowns: []string{
			"current-derived seven-row contact target",
			"source functional whose image is the seven contact-row carrier",
			"projection from u(4) or a dual carrier to contact rows",
			"representation, mass activation, and decoupling row for contact modes",
			"threshold-corrected beta tensor",
		},
		RecommendedNextGate: "Gate 126 — contact seven-row target projection / u(4)-to-contact quotient obstruction theorem",
	}, nil
}

func buildRows(in []contactdualpairing.DualPairRow, eps float64) []CarrierRow {
	rows := make([]CarrierRow, 0, len(in))
	vals := make([]float64, 0, len(in))
	for _, r := range in {
		vals = append(vals, r.Value)
	}
	sort.Float64s(vals)
	for _, r := range in {
		slot := sort.SearchFloat64s(vals, r.Value)
		if slot < len(vals) && math.Abs(vals[slot]-r.Value) <= eps {
			slot++
		} else {
			slot = 0
		}
		rows = append(rows, CarrierRow{
			Name:                   r.Name,
			ModeKind:               r.ModeKind,
			Value:                  r.Value,
			SpectralSlot:           slot,
			InSpectralSevenCarrier: true,
			InFanoSevenCarrier:     true,
			InAnonymousCarrier:     true,
			CurrentTargetRow:       false,
			SourceFunctionalRow:    false,
			RepresentationRow:      false,
			CanEnterBetaTensor:     false,
			ZeroRowProved:          false,
			Status:                 "seven-row-target-open",
			Reason:                 "this row has a spectral storage slot, but no current-derived target row, source functional, representation row, mass activation, or decoupling rule is derived",
		})
	}
	sort.SliceStable(rows, func(i, j int) bool {
		if math.Abs(rows[i].Value-rows[j].Value) < 1e-12 {
			return rows[i].Name < rows[j].Name
		}
		return rows[i].Value < rows[j].Value
	})
	return rows
}

func buildCandidates(prev contactdualpairing.Analysis, distinct int) []TargetCandidate {
	return []TargetCandidate{
		{
			Name: "uniform scalar/contact dual target", Kind: UniformScalarTarget,
			Dimension: 1, Constructed: true, CanonicalUnderCurrentData: true,
			SevenRowCarrier: false, RowDistinguishing: false, CurrentDerived: false,
			RejectedAsPremature: true, Obstruction: "canonical rank-one target is row-blind",
		},
		{
			Name: "derived contact electroweak block target", Kind: ContactEWBlockTarget,
			Dimension: 4, Constructed: true, CanonicalUnderCurrentData: true,
			SevenRowCarrier: false, RowDistinguishing: false, CurrentDerived: true,
			RejectedAsPremature: true, Obstruction: "derived contact target has four su(2)+u(1) seeds, not seven contact-row slots",
		},
		{
			Name: "typed u(4)/Pati-Salam current inventory", Kind: PatiSalamCurrentTarget,
			Dimension: 16, Constructed: true, CanonicalUnderCurrentData: true,
			SevenRowCarrier: false, RowDistinguishing: false, CurrentDerived: true,
			RejectedAsPremature: true, Obstruction: "sixteen current generators are typed, but no finite projection to seven contact rows is selected",
		},
		{
			Name: "leptoquark current subtarget", Kind: LeptoquarkSixTarget,
			Dimension: 6, Constructed: true, CanonicalUnderCurrentData: true,
			SevenRowCarrier: false, RowDistinguishing: false, CurrentDerived: true,
			RejectedAsPremature: true, Obstruction: "six-dimensional off-diagonal current sector is not the seven-row contact carrier",
		},
		{
			Name: "diagonal spectral seven-row carrier", Kind: SpectralSevenTarget,
			Dimension: 7, Constructed: distinct == 7, CanonicalUnderCurrentData: true,
			SevenRowCarrier: true, RowDistinguishing: distinct == 7, CurrentDerived: false,
			SourceFunctionalDerived: false, DualCurrentTargetDerived: false,
			RejectedAsPremature: true, Obstruction: "R^7 spectral slots distinguish rows diagnostically but are not current-derived source/dual target semantics",
		},
		{
			Name: "Fano seven-row carrier", Kind: FanoSevenTarget,
			Dimension: 7, Constructed: true, CanonicalUnderCurrentData: false,
			SevenRowCarrier: true, RowDistinguishing: true, CurrentDerived: false,
			RequiresHiddenChoice: true, HiddenChoices: prev.HiddenFanoChoices,
			RejectedAsPremature: true, Obstruction: "Fano R^7 needs one of 7! contact-to-Fano assignments before it can carry contact rows",
		},
		{
			Name: "anonymous contact seven-row storage carrier", Kind: AnonymousSevenTarget,
			Dimension: 7, Constructed: true, CanonicalUnderCurrentData: true,
			SevenRowCarrier: true, RowDistinguishing: false, CurrentDerived: false,
			RejectedAsPremature: true, Obstruction: "anonymous R^7 stores cardinality but supplies no row semantics or current functional",
		},
		{
			Name: "observed-fitted seven-row target", Kind: ObservedFittedSevenTarget,
			Dimension: 7, Constructed: false, CanonicalUnderCurrentData: false,
			SevenRowCarrier: true, UsesObservedInput: true,
			RejectedAsPremature: true, Obstruction: "observed constants, masses, or thresholds cannot define the finite contact target",
		},
	}
}

func buildCriteria(s TargetSummary) []TargetCriterion {
	return []TargetCriterion{
		{Name: "seven contact rows inherited", Required: true, Derived: s.ContactRows == 7, Detail: "Gate 124 leaves seven contact rows open"},
		{Name: "at least one seven-row carrier can be named", Required: true, Derived: s.CandidateSevenCarriers > 0, Detail: "spectral/Fano/anonymous R^7 carriers exist as candidates"},
		{Name: "current-derived seven-row target", Required: true, Derived: s.CurrentDerivedSevenCarriers > 0, Detail: "must be false: no current-derived R^7 target is selected"},
		{Name: "representation-complete rows", Required: true, Derived: s.RepresentationCompleteRows > 0, Detail: "must be false: no beta permission"},
		{Name: "residual nullity preserved", Required: true, Derived: s.ResidualNullityBefore == 3 && s.ResidualNullityAfter == 3, Detail: "physical-flow free variables remain sealed"},
	}
}

func countDistinct(rows []CarrierRow, eps float64) int {
	if len(rows) == 0 {
		return 0
	}
	vals := make([]float64, 0, len(rows))
	for _, r := range rows {
		vals = append(vals, r.Value)
	}
	sort.Float64s(vals)
	count := 1
	last := vals[0]
	for _, v := range vals[1:] {
		if math.Abs(v-last) > eps {
			count++
			last = v
		}
	}
	return count
}

func FormatRows(rows []CarrierRow, limit int) string {
	if limit <= 0 || limit > len(rows) {
		limit = len(rows)
	}
	parts := make([]string, 0, limit)
	for i := 0; i < limit; i++ {
		r := rows[i]
		parts = append(parts, fmt.Sprintf("%s[slot=%d value=%.10f current=%t rep=%t beta=%t]", r.Name, r.SpectralSlot, r.Value, r.CurrentTargetRow, r.RepresentationRow, r.CanEnterBetaTensor))
	}
	if limit < len(rows) {
		parts = append(parts, fmt.Sprintf("... +%d more", len(rows)-limit))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatCandidates(xs []TargetCandidate) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s(kind=%s dim=%d seven=%t canonical=%t current=%t hidden=%t beta=%d obstruction=%s)", x.Name, x.Kind, x.Dimension, x.SevenRowCarrier, x.CanonicalUnderCurrentData, x.CurrentDerived && x.DualCurrentTargetDerived, x.RequiresHiddenChoice || x.UsesObservedInput, x.BetaRowsPermitted, x.Obstruction))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatCriteria(xs []TargetCriterion) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s=%t", x.Name, x.Derived))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatSummary(s TargetSummary) string {
	return fmt.Sprintf("rows=%d distinct=%d candidateSeven=%d canonicalSeven=%d currentDerivedSeven=%d repRows=%d betaRows=%d zeroRows=%d hiddenChoices=%d nullity=%d->%d", s.ContactRows, s.DistinctSpectralRows, s.CandidateSevenCarriers, s.CanonicalSevenCarriers, s.CurrentDerivedSevenCarriers, s.RepresentationCompleteRows, s.ContactBetaRowsAllowed, s.ContactZeroRowsProved, s.HiddenChoices, s.ResidualNullityBefore, s.ResidualNullityAfter)
}

func Join(xs []string) string { return strings.Join(xs, "; ") }
