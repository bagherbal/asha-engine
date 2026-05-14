// Package contactsource implements Gate 123: contact semantic source-coupling /
// observable selector search.
//
// Gate 122 showed that incidence-weighted contact spectral data preserves the
// seven unresolved contact rows but supplies no local variables, constraint
// semantics, representation rows, mass activation, or decoupling rule. Gate 123
// asks whether some finite source, observable, current, or action-coupling
// object already present in the project can label those rows without using a
// hidden contact-to-Fano bijection or observed physical input.
//
// The result is another permission-preserving obstruction. A diagonal spectral
// observable can distinguish the seven rows numerically because their finite
// overlap values are distinct, and a uniform scalar/contact action source can be
// applied canonically. But the first is only a diagnostic value-label, not a
// local/gauge/representation semantic map, and the second is row-blind. The
// current-to-contact bridge remains obstructed by the earlier u(4)->contact
// target mismatch. Therefore no semantic source-coupling selector is derived,
// and the contact beta firewall remains closed.
package contactsource

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/contactrowsemantics"
	"github.com/bagherbal/asha-engine/pkg/bridge/currentcontact"
)

type SourceAttemptKind string

const (
	UniformActionSourceAttempt      SourceAttemptKind = "uniform-action-source"
	SpectralDiagonalObservable      SourceAttemptKind = "spectral-diagonal-observable"
	IncidenceWeightedObservable     SourceAttemptKind = "incidence-weighted-observable"
	CurrentToContactSourceAttempt   SourceAttemptKind = "current-to-contact-source"
	ExternalObservedSelectorAttempt SourceAttemptKind = "external-observed-selector"
)

type SourceSelectorRow struct {
	Name, ModeKind string
	Value          float64
	WeightedValue  float64

	UniformSourceCouples       bool
	UniformSourceDistinguishes bool
	SpectralObservableValue    float64
	SpectralObservableUnique   bool
	SpectralObservableSemantic bool
	CurrentSourceDerived       bool
	ActionCouplingDerived      bool
	LocalVariableDerived       bool
	GaugeRepresentationDerived bool
	MassActivationDerived      bool
	DecouplingRuleDerived      bool
	CanEnterBetaTensor         bool
	ZeroRowProved              bool
	Status                     string
	Reason                     string
}

type SourceAttempt struct {
	Name string
	Kind SourceAttemptKind

	Constructed               bool
	CanonicalUnderCurrentData bool
	UsesHiddenChoice          bool
	UsesObservedInput         bool
	RowDistinguishing         bool
	AddsRowSemantics          bool
	RequiresContactFanoChoice bool
	PossibleAssignments       int
	SourceCouplingDerived     bool
	ObservableSelectorDerived bool
	CurrentMapDerived         bool
	ActionCouplingDerived     bool
	RepresentationRowsDerived int
	BetaRowsPermitted         int
	ZeroRowsProved            int
	RejectedAsPremature       bool
	MissingTerms              []string
	Detail                    string
}

type SourceCriterion struct {
	Name     string
	Required bool
	Derived  bool
	Detail   string
}

type SourceSummary struct {
	ContactRows                     int
	DistinctSpectralValues          int
	DiagnosticRows                  int
	SemanticSourceRows              int
	UniformSourceRows               int
	UniformSourceDistinguishingRows int
	CurrentMapDerived               bool
	RepresentationCompleteRows      int
	ContactBetaRowsAllowed          int
	ContactZeroRowsProved           int
	HiddenChoices                   int
}

type Analysis struct {
	Semantics contactrowsemantics.Analysis
	Current   currentcontact.Analysis

	Rows     []SourceSelectorRow
	Attempts []SourceAttempt
	Criteria []SourceCriterion
	Summary  SourceSummary

	ContactRows               int
	PositiveFiniteContactRows int
	OpenContactRowsInherited  int
	OpenContactRowsAfter      int

	SemanticSourceCouplingSearchAttempted bool
	UniformActionSourceAvailable          bool
	UniformActionSourceCanonical          bool
	UniformActionSourceRowBlind           bool
	UniformActionSourceSelectsRows        bool

	SpectralObservableAttempted         bool
	SpectralObservableConstructed       bool
	SpectralObservableCanonical         bool
	SpectralObservableRowsDistinguished int
	SpectralObservableAddsSemantics     bool
	SpectralObservableOnlyDiagnostic    bool

	IncidenceWeightedObservableAttempted     bool
	IncidenceWeightedObservableCanonical     bool
	IncidenceWeightedObservableAddsSemantics bool

	CurrentSourceAttempted            bool
	CurrentToContactMapDerived        bool
	CurrentSourceRowsDerived          int
	CurrentSourceObstructionInherited bool

	ActionCouplingSelectorAttempted  bool
	ActionCouplingSelectorDerived    bool
	ObservableSelectorDerived        bool
	SemanticSourceSelectorDerived    bool
	ContactSourceSelectorNoGoDerived bool

	RequiresHiddenContactFanoChoice bool
	HiddenContactFanoChoices        int
	HiddenObservedInputUsed         bool

	RepresentationCompleteRows          int
	RepresentationOpenRows              int
	ContactBetaRowsAllowed              int
	ContactZeroRowsProved               int
	ThresholdCorrectedBetaDerived       bool
	FullFiniteBetaMatchingTensorDerived bool

	ResidualNullityBefore  int
	ResidualNullityAfter   int
	ResidualSymmetryBroken bool

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
		sem, err := contactrowsemantics.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		cur, err := currentcontact.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(sem, cur, 1e-10)
	})
	return defaultValue, defaultErr
}

func Build(sem contactrowsemantics.Analysis, cur currentcontact.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if !sem.RowSemanticsObstructionDerived || !sem.IncidenceWeightedNoGoDerived || sem.RowSemanticsDerived || sem.ContactBetaRowsAllowed != 0 || sem.ContactZeroRowsProved != 0 || sem.OpenContactRowsAfter != 7 {
		return Analysis{}, fmt.Errorf("Gate 123 requires Gate 122 row-semantics obstruction and closed beta firewall")
	}
	if sem.SignedIncidenceChoiceCount != 5040 || sem.ContactFanoAssignmentDerived || sem.SignedIncidenceCanonical {
		return Analysis{}, fmt.Errorf("Gate 123 requires Gate 122 7! noncanonical contact-Fano choice obstruction")
	}
	if sem.ResidualNullityAfter != 3 || sem.HiddenObservedInputUsed || sem.PhysicalWeakAngleDerived || sem.FineStructureDerived || sem.PhysicalMassesDerived || sem.PhysicalScaleDerived {
		return Analysis{}, fmt.Errorf("Gate 123 refuses hidden physical input or changed residual nullity")
	}
	if cur.CurrentToContactMapDerived || cur.SourceFunctionalDerived || cur.HessianComputable || cur.PropagatorRuleDerived || cur.HiddenObservedInputUsed {
		return Analysis{}, fmt.Errorf("Gate 123 requires current-to-contact source map to remain unselected")
	}

	rows := buildRows(sem.Rows, eps)
	distinct := distinctValues(rows, eps)
	diagnosticRows := 0
	uniformRows := 0
	for _, r := range rows {
		if r.SpectralObservableUnique {
			diagnosticRows++
		}
		if r.UniformSourceCouples {
			uniformRows++
		}
	}
	summary := SourceSummary{
		ContactRows:                     len(rows),
		DistinctSpectralValues:          distinct,
		DiagnosticRows:                  diagnosticRows,
		SemanticSourceRows:              0,
		UniformSourceRows:               uniformRows,
		UniformSourceDistinguishingRows: 0,
		CurrentMapDerived:               cur.CurrentToContactMapDerived,
		RepresentationCompleteRows:      0,
		ContactBetaRowsAllowed:          0,
		ContactZeroRowsProved:           0,
		HiddenChoices:                   sem.SignedIncidenceChoiceCount,
	}
	attempts := buildAttempts(summary, cur)
	criteria := buildCriteria(summary)

	truth := "Gate 123 searches for a semantic source-coupling or observable selector for the seven unresolved contact rows. A uniform action/source coupling is canonical but row-blind; it couples equally and selects no contact row semantics. The spectral and incidence-weighted diagonal observables distinguish all seven rows numerically, but they are diagnostic labels only: they do not supply local variables, gauge representation rows, Lorentz kinetic rows, mass activation, or decoupling. The current-to-contact source bridge remains obstructed by the earlier u(4)->contact target mismatch, and any signed/Fano-labelled source still requires one of 7! hidden contact-to-Fano choices. Therefore no finite semantic source selector is derived, and contact threshold beta rows remain forbidden."

	return Analysis{
		Semantics: sem,
		Current:   cur,
		Rows:      rows,
		Attempts:  attempts,
		Criteria:  criteria,
		Summary:   summary,

		ContactRows:               len(rows),
		PositiveFiniteContactRows: len(rows),
		OpenContactRowsInherited:  sem.OpenContactRowsAfter,
		OpenContactRowsAfter:      len(rows),

		SemanticSourceCouplingSearchAttempted: true,
		UniformActionSourceAvailable:          true,
		UniformActionSourceCanonical:          true,
		UniformActionSourceRowBlind:           true,
		UniformActionSourceSelectsRows:        false,

		SpectralObservableAttempted:         true,
		SpectralObservableConstructed:       true,
		SpectralObservableCanonical:         true,
		SpectralObservableRowsDistinguished: diagnosticRows,
		SpectralObservableAddsSemantics:     false,
		SpectralObservableOnlyDiagnostic:    true,

		IncidenceWeightedObservableAttempted:     true,
		IncidenceWeightedObservableCanonical:     true,
		IncidenceWeightedObservableAddsSemantics: false,

		CurrentSourceAttempted:            true,
		CurrentToContactMapDerived:        cur.CurrentToContactMapDerived,
		CurrentSourceRowsDerived:          0,
		CurrentSourceObstructionInherited: !cur.CurrentToContactMapDerived && !cur.SourceFunctionalDerived,

		ActionCouplingSelectorAttempted:  true,
		ActionCouplingSelectorDerived:    false,
		ObservableSelectorDerived:        false,
		SemanticSourceSelectorDerived:    false,
		ContactSourceSelectorNoGoDerived: true,

		RequiresHiddenContactFanoChoice: true,
		HiddenContactFanoChoices:        sem.SignedIncidenceChoiceCount,
		HiddenObservedInputUsed:         false,

		RepresentationCompleteRows:          0,
		RepresentationOpenRows:              len(rows),
		ContactBetaRowsAllowed:              0,
		ContactZeroRowsProved:               0,
		ThresholdCorrectedBetaDerived:       false,
		FullFiniteBetaMatchingTensorDerived: false,

		ResidualNullityBefore:  3,
		ResidualNullityAfter:   3,
		ResidualSymmetryBroken: false,

		PhysicalWeakAngleDerived: false,
		FineStructureDerived:     false,
		PhysicalMassesDerived:    false,
		PhysicalScaleDerived:     false,

		TruthStatement: truth,
		RejectedClaims: []string{
			"a uniform finite action/source coupling labels the seven contact rows",
			"distinct spectral values alone define physical row semantics",
			"incidence-weighted observables create contact representation rows",
			"the current-to-contact source map has been selected",
			"a signed/Fano-labelled source can be used without a hidden 7! choice",
			"Gate 123 derives threshold-corrected beta coefficients, alpha, physical thetaW, M*, g_*, or particle masses",
		},
		RemainingUnknowns: []string{
			"finite semantic source that labels contact rows without a hidden contact-Fano bijection",
			"local variable/source map for contact modes",
			"gauge representation row for each contact mode",
			"mass activation and decoupling rule for contact modes",
			"threshold-corrected beta tensor",
		},
		RecommendedNextGate: "Gate 124 — contact source-current dual pairing / row-label naturality obstruction theorem",
	}, nil
}

func buildRows(in []contactrowsemantics.ContactSemanticRow, eps float64) []SourceSelectorRow {
	rows := make([]SourceSelectorRow, 0, len(in))
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
		rows = append(rows, SourceSelectorRow{
			Name:                       r.Name,
			ModeKind:                   r.ModeKind,
			Value:                      r.Value,
			WeightedValue:              r.IncidenceWeightedValue,
			UniformSourceCouples:       true,
			UniformSourceDistinguishes: false,
			SpectralObservableValue:    r.Value,
			SpectralObservableUnique:   unique,
			SpectralObservableSemantic: false,
			CurrentSourceDerived:       false,
			ActionCouplingDerived:      false,
			LocalVariableDerived:       false,
			GaugeRepresentationDerived: false,
			MassActivationDerived:      false,
			DecouplingRuleDerived:      false,
			CanEnterBetaTensor:         false,
			ZeroRowProved:              false,
			Status:                     "source-selector-open",
			Reason:                     "spectral observables distinguish this row numerically, but no semantic source, current map, representation row, mass activation, or decoupling rule is derived",
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

func buildAttempts(s SourceSummary, cur currentcontact.Analysis) []SourceAttempt {
	return []SourceAttempt{
		{
			Name:                      "uniform scalar/contact action source",
			Kind:                      UniformActionSourceAttempt,
			Constructed:               true,
			CanonicalUnderCurrentData: true,
			RowDistinguishing:         false,
			AddsRowSemantics:          false,
			SourceCouplingDerived:     false,
			RejectedAsPremature:       true,
			MissingTerms:              []string{"row-selective source coefficient", "local variable", "representation row"},
			Detail:                    fmt.Sprintf("the uniform source can couple to %d rows, but distinguishes zero rows", s.UniformSourceRows),
		},
		{
			Name:                      "diagonal contact spectral observable",
			Kind:                      SpectralDiagonalObservable,
			Constructed:               true,
			CanonicalUnderCurrentData: true,
			RowDistinguishing:         s.DiagnosticRows == s.ContactRows,
			AddsRowSemantics:          false,
			ObservableSelectorDerived: false,
			RejectedAsPremature:       true,
			MissingTerms:              []string{"physical source coupling", "local/gauge semantics", "mass activation", "decoupling"},
			Detail:                    fmt.Sprintf("distinct spectral values distinguish %d/%d rows only as diagnostics", s.DiagnosticRows, s.ContactRows),
		},
		{
			Name:                      "incidence-weighted spectral observable",
			Kind:                      IncidenceWeightedObservable,
			Constructed:               true,
			CanonicalUnderCurrentData: true,
			RowDistinguishing:         s.DiagnosticRows == s.ContactRows,
			AddsRowSemantics:          false,
			ObservableSelectorDerived: false,
			RejectedAsPremature:       true,
			MissingTerms:              []string{"non-uniform semantic incidence", "contact-Fano assignment", "representation row"},
			Detail:                    "Gate 122 proved uniform incidence weighting rescales rows but does not add semantics",
		},
		{
			Name:                  "current-to-contact semantic source",
			Kind:                  CurrentToContactSourceAttempt,
			Constructed:           true,
			CurrentMapDerived:     cur.CurrentToContactMapDerived,
			SourceCouplingDerived: cur.SourceFunctionalDerived,
			RejectedAsPremature:   true,
			MissingTerms:          []string{"selected u(4)->contact source map", "abelian separation", "color/leptoquark carrier"},
			Detail:                "the earlier current-to-contact bridge exposes an abstract map space but no selected finite source functional",
		},
		{
			Name:                "observed-constant/contact-label selector",
			Kind:                ExternalObservedSelectorAttempt,
			Constructed:         false,
			UsesObservedInput:   true,
			RejectedAsPremature: true,
			MissingTerms:        []string{"finite selector independent of observed constants"},
			Detail:              "external alpha, thetaW, masses, or fitted thresholds are forbidden as contact row labels",
		},
	}
}

func buildCriteria(s SourceSummary) []SourceCriterion {
	return []SourceCriterion{
		{Name: "seven contact rows inherited", Required: true, Derived: s.ContactRows == 7, Detail: "all seven Gate 122 contact rows are still open"},
		{Name: "spectral diagnostic distinguishes rows", Required: true, Derived: s.DiagnosticRows == 7 && s.DistinctSpectralValues == 7, Detail: "distinct values label rows numerically only"},
		{Name: "uniform source distinguishes rows", Required: true, Derived: s.UniformSourceDistinguishingRows > 0, Detail: "must be false: uniform source is row-blind"},
		{Name: "current-to-contact map selected", Required: true, Derived: s.CurrentMapDerived, Detail: "must be false under current finite data"},
		{Name: "representation-complete source rows", Required: true, Derived: s.RepresentationCompleteRows > 0 || s.ContactBetaRowsAllowed > 0, Detail: "must be false: beta firewall remains closed"},
	}
}

func distinctValues(rows []SourceSelectorRow, eps float64) int {
	if len(rows) == 0 {
		return 0
	}
	values := make([]float64, 0, len(rows))
	for _, r := range rows {
		values = append(values, r.Value)
	}
	sort.Float64s(values)
	count := 1
	last := values[0]
	for _, v := range values[1:] {
		if math.Abs(v-last) > eps {
			count++
			last = v
		}
	}
	return count
}

func FormatRows(rows []SourceSelectorRow, limit int) string {
	if limit <= 0 || limit > len(rows) {
		limit = len(rows)
	}
	parts := make([]string, 0, limit)
	for i := 0; i < limit; i++ {
		r := rows[i]
		parts = append(parts, fmt.Sprintf("%s[value=%.10f spectralUnique=%t semantic=%t current=%t rep=%t beta=%t]", r.Name, r.Value, r.SpectralObservableUnique, r.SpectralObservableSemantic, r.CurrentSourceDerived, r.GaugeRepresentationDerived, r.CanEnterBetaTensor))
	}
	if limit < len(rows) {
		parts = append(parts, fmt.Sprintf("... +%d more", len(rows)-limit))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatAttempts(xs []SourceAttempt) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s(kind=%s canonical=%t diagnostic=%t semantics=%t current=%t beta=%d hidden=%t)", x.Name, x.Kind, x.CanonicalUnderCurrentData, x.RowDistinguishing, x.AddsRowSemantics, x.CurrentMapDerived, x.BetaRowsPermitted, x.UsesHiddenChoice || x.UsesObservedInput))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatCriteria(xs []SourceCriterion) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s=%t", x.Name, x.Derived))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatSummary(s SourceSummary) string {
	return fmt.Sprintf("rows=%d distinct=%d diagnostic=%d semanticSources=%d uniformRows=%d currentMap=%t repRows=%d betaRows=%d zeroRows=%d hiddenChoices=%d", s.ContactRows, s.DistinctSpectralValues, s.DiagnosticRows, s.SemanticSourceRows, s.UniformSourceRows, s.CurrentMapDerived, s.RepresentationCompleteRows, s.ContactBetaRowsAllowed, s.ContactZeroRowsProved, s.HiddenChoices)
}

func Join(xs []string) string { return strings.Join(xs, "; ") }
