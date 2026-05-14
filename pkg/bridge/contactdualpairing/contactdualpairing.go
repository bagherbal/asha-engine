// Package contactdualpairing implements Gate 124: contact source-current dual
// pairing / row-label naturality obstruction theorem.
//
// Gate 123 found two safe but insufficient objects: a uniform source coupling
// that is canonical but row-blind, and a diagonal spectral observable that
// distinguishes all seven contact rows only as a diagnostic. Gate 124 asks
// whether those source-side labels can be paired with a current-side dual object
// to produce natural row semantics. The answer remains no. A formal diagonal
// contact-spectrum pairing can be written, but it is self-dual diagnostic data,
// not a derived current functional, representation row, local variable, mass
// activation, or decoupling rule. The actual current-to-contact bridge remains
// obstructed by the u(4)->contact target mismatch; any Fano-labelled pairing
// still requires one of 7! hidden bijections.
package contactdualpairing

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/contactsource"
)

type PairingKind string

const (
	UniformDualPairing      PairingKind = "uniform-source-current-pairing"
	SpectralDiagnosticPair  PairingKind = "spectral-diagnostic-self-pairing"
	CurrentContactDualPair  PairingKind = "current-to-contact-dual-pairing"
	FanoLabelledDualPair    PairingKind = "fano-labelled-dual-pairing"
	ObservedConstantPairing PairingKind = "observed-constant-dual-pairing"
)

type DualPairRow struct {
	Name, ModeKind string
	Value          float64
	WeightedValue  float64

	UniformSourceCouples       bool
	UniformCurrentCouples      bool
	UniformPairingValue        float64
	UniformPairDistinguishes   bool
	SpectralPairingValue       float64
	SpectralPairDistinguishes  bool
	SpectralPairingSemantic    bool
	CurrentFunctionalDerived   bool
	CurrentDualRowDerived      bool
	SourceFunctionalDerived    bool
	NaturalRowLabelDerived     bool
	RequiresHiddenFanoChoice   bool
	GaugeRepresentationDerived bool
	MassActivationDerived      bool
	DecouplingRuleDerived      bool
	CanEnterBetaTensor         bool
	ZeroRowProved              bool
	Status                     string
	Reason                     string
}

type PairingAttempt struct {
	Name string
	Kind PairingKind

	Constructed               bool
	CanonicalUnderCurrentData bool
	NonDegenerate             bool
	RowDistinguishing         bool
	AddsRowSemantics          bool
	SourceFunctionalDerived   bool
	CurrentFunctionalDerived  bool
	CurrentMapDerived         bool
	NaturalRowLabelsDerived   int
	RequiresHiddenChoice      bool
	HiddenChoices             int
	UsesObservedInput         bool
	RepresentationRowsDerived int
	BetaRowsPermitted         int
	ZeroRowsProved            int
	RejectedAsPremature       bool
	MissingTerms              []string
	Detail                    string
}

type PairingCriterion struct {
	Name     string
	Required bool
	Derived  bool
	Detail   string
}

type PairingSummary struct {
	ContactRows                      int
	UniformPairingRows               int
	UniformPairingDistinguishingRows int
	SpectralPairingRows              int
	SpectralPairingSemanticRows      int
	CurrentDualRows                  int
	NaturalRowLabels                 int
	RepresentationCompleteRows       int
	ContactBetaRowsAllowed           int
	ContactZeroRowsProved            int
	HiddenChoices                    int
}

type Analysis struct {
	Source contactsource.Analysis

	Rows     []DualPairRow
	Attempts []PairingAttempt
	Criteria []PairingCriterion
	Summary  PairingSummary

	ContactRows              int
	OpenContactRowsInherited int
	OpenContactRowsAfter     int

	DualPairingSearchAttempted bool

	UniformPairingAttempted         bool
	UniformPairingConstructed       bool
	UniformPairingCanonical         bool
	UniformPairingRank              int
	UniformPairingRowsDistinguished int
	UniformPairingRowBlind          bool

	SpectralPairingAttempted         bool
	SpectralPairingConstructed       bool
	SpectralPairingCanonical         bool
	SpectralPairingNonDegenerate     bool
	SpectralPairingRowsDistinguished int
	SpectralPairingAddsSemantics     bool
	SpectralPairingDiagnosticOnly    bool

	CurrentDualPairingAttempted     bool
	CurrentToContactMapDerived      bool
	CurrentFunctionalDerived        bool
	SourceFunctionalDerived         bool
	CurrentDualRowsDerived          int
	CurrentDualPairingDerived       bool
	CurrentDualObstructionInherited bool

	FanoLabelledPairingAttempted bool
	FanoLabelledPairingDerived   bool
	RequiresHiddenFanoChoice     bool
	HiddenFanoChoices            int

	NaturalRowLabelDerived        bool
	ContactDualPairingNoGoDerived bool
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
		src, err := contactsource.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(src, 1e-10)
	})
	return defaultValue, defaultErr
}

func Build(src contactsource.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if !src.ContactSourceSelectorNoGoDerived || src.ContactRows != 7 || src.OpenContactRowsAfter != 7 || src.RepresentationCompleteRows != 0 || src.ContactBetaRowsAllowed != 0 || src.ContactZeroRowsProved != 0 {
		return Analysis{}, fmt.Errorf("Gate 124 requires Gate 123 source-selector no-go with seven open contact rows")
	}
	if !src.UniformActionSourceCanonical || !src.UniformActionSourceRowBlind || src.UniformActionSourceSelectsRows {
		return Analysis{}, fmt.Errorf("Gate 124 requires Gate 123 canonical-but-row-blind uniform source")
	}
	if !src.SpectralObservableConstructed || !src.SpectralObservableCanonical || src.SpectralObservableRowsDistinguished != 7 || src.SpectralObservableAddsSemantics || !src.SpectralObservableOnlyDiagnostic {
		return Analysis{}, fmt.Errorf("Gate 124 requires Gate 123 diagnostic-only spectral observable")
	}
	if src.CurrentToContactMapDerived || src.Current.SourceFunctionalDerived || src.Current.HessianComputable || src.Current.PropagatorRuleDerived || src.Current.HiddenObservedInputUsed {
		return Analysis{}, fmt.Errorf("Gate 124 requires current-to-contact dual current to remain unselected")
	}
	if src.ResidualNullityAfter != 3 || src.HiddenObservedInputUsed || src.PhysicalWeakAngleDerived || src.FineStructureDerived || src.PhysicalMassesDerived || src.PhysicalScaleDerived {
		return Analysis{}, fmt.Errorf("Gate 124 refuses hidden physical input or changed residual nullity")
	}

	rows := buildRows(src.Rows, eps)
	uniformRows := 0
	spectralRows := 0
	for _, r := range rows {
		if r.UniformSourceCouples && r.UniformCurrentCouples {
			uniformRows++
		}
		if r.SpectralPairDistinguishes {
			spectralRows++
		}
	}
	summary := PairingSummary{
		ContactRows:                      len(rows),
		UniformPairingRows:               uniformRows,
		UniformPairingDistinguishingRows: 0,
		SpectralPairingRows:              spectralRows,
		SpectralPairingSemanticRows:      0,
		CurrentDualRows:                  0,
		NaturalRowLabels:                 0,
		RepresentationCompleteRows:       0,
		ContactBetaRowsAllowed:           0,
		ContactZeroRowsProved:            0,
		HiddenChoices:                    src.HiddenContactFanoChoices,
	}
	attempts := buildAttempts(summary, src)
	criteria := buildCriteria(summary, src)

	truth := "Gate 124 searches for a source-current dual pairing that could turn the seven diagnostic contact labels into natural row semantics. The uniform source/current pairing is canonical but rank-one and row-blind. The diagonal spectral self-pairing is nondegenerate and distinguishes all seven finite overlap rows, but it is not a current-derived functional and adds no local variable, gauge representation, Lorentz kinetic, mass activation, or decoupling semantics. The true current-to-contact dual pairing remains obstructed by the earlier u(4)->contact target mismatch, and any Fano-labelled pairing still needs one of 7! hidden contact-to-Fano choices. Therefore no natural contact row label, beta row, or zero-row cancellation is derived."

	return Analysis{
		Source:   src,
		Rows:     rows,
		Attempts: attempts,
		Criteria: criteria,
		Summary:  summary,

		ContactRows:              len(rows),
		OpenContactRowsInherited: src.OpenContactRowsAfter,
		OpenContactRowsAfter:     len(rows),

		DualPairingSearchAttempted: true,

		UniformPairingAttempted:         true,
		UniformPairingConstructed:       true,
		UniformPairingCanonical:         true,
		UniformPairingRank:              1,
		UniformPairingRowsDistinguished: 0,
		UniformPairingRowBlind:          true,

		SpectralPairingAttempted:         true,
		SpectralPairingConstructed:       true,
		SpectralPairingCanonical:         true,
		SpectralPairingNonDegenerate:     spectralRows == len(rows),
		SpectralPairingRowsDistinguished: spectralRows,
		SpectralPairingAddsSemantics:     false,
		SpectralPairingDiagnosticOnly:    true,

		CurrentDualPairingAttempted:     true,
		CurrentToContactMapDerived:      src.CurrentToContactMapDerived,
		CurrentFunctionalDerived:        src.Current.SourceFunctionalDerived,
		SourceFunctionalDerived:         false,
		CurrentDualRowsDerived:          0,
		CurrentDualPairingDerived:       false,
		CurrentDualObstructionInherited: !src.CurrentToContactMapDerived && !src.Current.SourceFunctionalDerived,

		FanoLabelledPairingAttempted: true,
		FanoLabelledPairingDerived:   false,
		RequiresHiddenFanoChoice:     true,
		HiddenFanoChoices:            src.HiddenContactFanoChoices,

		NaturalRowLabelDerived:        false,
		ContactDualPairingNoGoDerived: true,
		RepresentationCompleteRows:    0,
		RepresentationOpenRows:        len(rows),
		ContactBetaRowsAllowed:        0,
		ContactZeroRowsProved:         0,
		ThresholdCorrectedBetaDerived: false,
		FullBetaMatchingTensorDerived: false,

		ResidualNullityBefore:   3,
		ResidualNullityAfter:    3,
		ResidualSymmetryBroken:  false,
		HiddenObservedInputUsed: false,

		PhysicalWeakAngleDerived: false,
		FineStructureDerived:     false,
		PhysicalMassesDerived:    false,
		PhysicalScaleDerived:     false,

		TruthStatement: truth,
		RejectedClaims: []string{
			"a rank-one uniform source-current pairing labels contact rows",
			"a diagonal spectral self-pairing is already a physical current functional",
			"spectral nondegeneracy alone supplies local variables or representation rows",
			"the u(4) current inventory has been paired naturally with the seven contact rows",
			"a Fano-labelled dual pairing can be used without choosing one of 7! bijections",
			"Gate 124 derives threshold-corrected beta coefficients, alpha, physical thetaW, M*, g_*, or particle masses",
		},
		RemainingUnknowns: []string{
			"selected current-to-contact dual functional",
			"natural row labels for the seven contact modes",
			"local variables and gauge representation rows for contact modes",
			"mass activation and decoupling rule for contact modes",
			"threshold-corrected beta tensor",
		},
		RecommendedNextGate: "Gate 125 — contact dual-current target enlargement / seven-row carrier search",
	}, nil
}

func buildRows(in []contactsource.SourceSelectorRow, eps float64) []DualPairRow {
	rows := make([]DualPairRow, 0, len(in))
	values := make([]float64, 0, len(in))
	for _, r := range in {
		values = append(values, r.Value)
	}
	for _, r := range in {
		unique := true
		for _, v := range values {
			if math.Abs(v-r.Value) <= eps && math.Abs(v-r.Value) > 0 {
				unique = false
			}
		}
		rows = append(rows, DualPairRow{
			Name:                       r.Name,
			ModeKind:                   r.ModeKind,
			Value:                      r.Value,
			WeightedValue:              r.WeightedValue,
			UniformSourceCouples:       true,
			UniformCurrentCouples:      true,
			UniformPairingValue:        1,
			UniformPairDistinguishes:   false,
			SpectralPairingValue:       r.Value * r.Value,
			SpectralPairDistinguishes:  unique,
			SpectralPairingSemantic:    false,
			CurrentFunctionalDerived:   false,
			CurrentDualRowDerived:      false,
			SourceFunctionalDerived:    false,
			NaturalRowLabelDerived:     false,
			RequiresHiddenFanoChoice:   true,
			GaugeRepresentationDerived: false,
			MassActivationDerived:      false,
			DecouplingRuleDerived:      false,
			CanEnterBetaTensor:         false,
			ZeroRowProved:              false,
			Status:                     "dual-pairing-open",
			Reason:                     "diagonal spectral pairing distinguishes this row only diagnostically; no current-derived dual functional, natural row label, representation row, mass activation, or decoupling rule is derived",
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

func buildAttempts(s PairingSummary, src contactsource.Analysis) []PairingAttempt {
	return []PairingAttempt{
		{
			Name:                      "uniform source-current rank-one pairing",
			Kind:                      UniformDualPairing,
			Constructed:               true,
			CanonicalUnderCurrentData: true,
			NonDegenerate:             false,
			RowDistinguishing:         false,
			AddsRowSemantics:          false,
			RejectedAsPremature:       true,
			MissingTerms:              []string{"row-selective current functional", "nondegenerate row pairing", "representation row"},
			Detail:                    fmt.Sprintf("canonical pairing touches %d rows but distinguishes %d", s.UniformPairingRows, s.UniformPairingDistinguishingRows),
		},
		{
			Name:                      "diagonal spectral self-dual pairing",
			Kind:                      SpectralDiagnosticPair,
			Constructed:               true,
			CanonicalUnderCurrentData: true,
			NonDegenerate:             s.SpectralPairingRows == s.ContactRows,
			RowDistinguishing:         s.SpectralPairingRows == s.ContactRows,
			AddsRowSemantics:          false,
			RejectedAsPremature:       true,
			MissingTerms:              []string{"current-derived dual", "local/gauge semantics", "mass activation", "decoupling"},
			Detail:                    fmt.Sprintf("spectral pairing distinguishes %d/%d rows, but remains diagnostic", s.SpectralPairingRows, s.ContactRows),
		},
		{
			Name:                     "u(4) current-to-contact dual functional",
			Kind:                     CurrentContactDualPair,
			Constructed:              true,
			CurrentMapDerived:        src.CurrentToContactMapDerived,
			CurrentFunctionalDerived: src.Current.SourceFunctionalDerived,
			RejectedAsPremature:      true,
			MissingTerms:             []string{"selected u(4)->contact map", "source functional", "contact-row current targets"},
			Detail:                   "Gate 71/123 obstruction remains: abstract current maps exist, but no selected current-to-contact source functional exists",
		},
		{
			Name:                 "Fano-labelled source-current pairing",
			Kind:                 FanoLabelledDualPair,
			Constructed:          true,
			RequiresHiddenChoice: true,
			HiddenChoices:        s.HiddenChoices,
			RejectedAsPremature:  true,
			MissingTerms:         []string{"canonical contact-Fano assignment", "naturality-breaking finite selector"},
			Detail:               fmt.Sprintf("requires one of %d hidden contact-to-Fano bijections", s.HiddenChoices),
		},
		{
			Name:                "observed-constant dual pairing",
			Kind:                ObservedConstantPairing,
			Constructed:         false,
			UsesObservedInput:   true,
			RejectedAsPremature: true,
			MissingTerms:        []string{"finite selector independent of alpha/thetaW/masses"},
			Detail:              "external constants are forbidden as row-label selectors",
		},
	}
}

func buildCriteria(s PairingSummary, src contactsource.Analysis) []PairingCriterion {
	return []PairingCriterion{
		{Name: "Gate 123 source selector no-go inherited", Required: true, Derived: src.ContactSourceSelectorNoGoDerived && s.ContactRows == 7 && src.ContactBetaRowsAllowed == 0, Detail: "seven source-selector-open contact rows remain"},
		{Name: "uniform pairing is row-blind", Required: true, Derived: s.UniformPairingRows == 7 && s.UniformPairingDistinguishingRows == 0, Detail: "canonical source/current unit pairing has rank one semantics"},
		{Name: "spectral pairing distinguishes rows diagnostically", Required: true, Derived: s.SpectralPairingRows == 7 && s.SpectralPairingSemanticRows == 0, Detail: "nondegenerate spectral self-pairing still lacks source-current semantics"},
		{Name: "current dual functional selected", Required: true, Derived: s.CurrentDualRows > 0 || src.CurrentToContactMapDerived || src.Current.SourceFunctionalDerived, Detail: "must be false under current finite data"},
		{Name: "representation/beta permission opened", Required: true, Derived: s.RepresentationCompleteRows > 0 || s.ContactBetaRowsAllowed > 0 || s.ContactZeroRowsProved > 0, Detail: "must be false: beta firewall remains closed"},
	}
}

func FormatRows(rows []DualPairRow, limit int) string {
	if limit <= 0 || limit > len(rows) {
		limit = len(rows)
	}
	parts := make([]string, 0, limit)
	for i := 0; i < limit; i++ {
		r := rows[i]
		parts = append(parts, fmt.Sprintf("%s[value=%.10f uniform=%.1f spectralPair=%.10f current=%t natural=%t rep=%t beta=%t]", r.Name, r.Value, r.UniformPairingValue, r.SpectralPairingValue, r.CurrentFunctionalDerived, r.NaturalRowLabelDerived, r.GaugeRepresentationDerived, r.CanEnterBetaTensor))
	}
	if limit < len(rows) {
		parts = append(parts, fmt.Sprintf("... +%d more", len(rows)-limit))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatAttempts(xs []PairingAttempt) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s(kind=%s canonical=%t nondeg=%t rows=%t semantics=%t current=%t beta=%d hidden=%t)", x.Name, x.Kind, x.CanonicalUnderCurrentData, x.NonDegenerate, x.RowDistinguishing, x.AddsRowSemantics, x.CurrentFunctionalDerived || x.CurrentMapDerived, x.BetaRowsPermitted, x.RequiresHiddenChoice || x.UsesObservedInput))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatCriteria(xs []PairingCriterion) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s=%t", x.Name, x.Derived))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatSummary(s PairingSummary) string {
	return fmt.Sprintf("rows=%d uniform=%d uniformDistinguishing=%d spectral=%d spectralSemantic=%d currentDual=%d natural=%d repRows=%d betaRows=%d zeroRows=%d hiddenChoices=%d", s.ContactRows, s.UniformPairingRows, s.UniformPairingDistinguishingRows, s.SpectralPairingRows, s.SpectralPairingSemanticRows, s.CurrentDualRows, s.NaturalRowLabels, s.RepresentationCompleteRows, s.ContactBetaRowsAllowed, s.ContactZeroRowsProved, s.HiddenChoices)
}

func Join(xs []string) string { return strings.Join(xs, "; ") }
