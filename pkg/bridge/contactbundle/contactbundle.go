// Package contactbundle implements Gate 115: contact local-bundle obstruction /
// representation-row construction attempt.
//
// Gate 114 blocked the constraint/BRST shortcut: the zero differential leaves
// seven contact cohomology classes alive, while nontrivial differentials require
// extra choices.  This gate returns to the other honest branch.  It asks whether
// the seven positive finite-overlap contact modes can be lifted into local
// continuum bundle data with representation rows suitable for threshold beta
// matching.
//
// The answer is still no.  The finite contact carrier is real and positive, but
// a local bundle requires a base-space map, fiber, transition/cocycle law,
// sections, gauge representation, Lorentz kinetic operator, hypercharge/color
// assignment, mass activation unit, and decoupling rule.  None of those are
// selected by the current finite data for the contact partial-overlap modes, so
// no representation rows and no threshold beta corrections are permitted.
package contactbundle

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/contactcohomology"
)

type BundleAttemptKind string

const (
	FiniteLabelBundleAttempt   BundleAttemptKind = "finite-label-bundle"
	SpectralLineBundleAttempt  BundleAttemptKind = "spectral-line-bundle"
	SingletScalarAttempt       BundleAttemptKind = "singlet-scalar-representation"
	DoubletScalarAttempt       BundleAttemptKind = "doublet-scalar-representation"
	ContactDistributionAttempt BundleAttemptKind = "contact-distribution-lift"
)

type RepresentationStatus string

const (
	RepresentationOpen      RepresentationStatus = "representation-open"
	RepresentationComplete  RepresentationStatus = "representation-complete"
	RepresentationForbidden RepresentationStatus = "representation-forbidden"
)

type BundleRow struct {
	Name, ModeKind string
	Value          float64

	FiniteOverlapPositive bool
	SurvivesCohomology    bool

	BaseSpaceMapDerived           bool
	FiberDerived                  bool
	TransitionFunctionsDerived    bool
	SectionMapDerived             bool
	GaugeRepresentationDerived    bool
	HyperchargeDerived            bool
	ColorRepresentationDerived    bool
	LorentzKineticDerived         bool
	CanonicalNormalizationDerived bool
	MassActivationDerived         bool
	DecouplingRuleDerived         bool

	RepresentationStatus     RepresentationStatus
	RepresentationRowDerived bool
	CanEnterBetaTensor       bool
	Reason                   string
}

type BundleAttempt struct {
	Name string
	Kind BundleAttemptKind

	Constructed               bool
	UsesExtraConvention       bool
	CanonicalUnderCurrentData bool

	BaseSpaceMapDerived        bool
	FiberDerived               bool
	TransitionFunctionsDerived bool
	SectionMapDerived          bool
	GaugeRepresentationDerived bool
	LorentzKineticDerived      bool
	MassActivationDerived      bool
	DecouplingRuleDerived      bool

	RepresentationRowDerived bool
	BetaRowPermitted         bool
	RejectedAsPremature      bool
	MissingTerms             []string
	Detail                   string
}

type BundleCriterion struct {
	Name     string
	Required bool
	Derived  bool
	Detail   string
}

type Analysis struct {
	ContactCohomology contactcohomology.Analysis

	Rows           []BundleRow
	BundleAttempts []BundleAttempt
	Criteria       []BundleCriterion

	ContactRows               int
	PositiveFiniteContactRows int
	SurvivingCohomologyRows   int
	OpenContactRowsBefore     int
	OpenContactRowsAfter      int

	LocalBundleConstructionAttempted        bool
	FiniteCarrierAvailable                  bool
	BaseSpaceMapDerived                     bool
	FiberDerived                            bool
	TransitionFunctionsDerived              bool
	SectionMapDerived                       bool
	GaugeRepresentationForContactDerived    bool
	HyperchargeForContactDerived            bool
	ColorRepresentationForContactDerived    bool
	LorentzKineticForContactDerived         bool
	CanonicalNormalizationForContactDerived bool
	MassActivationForContactDerived         bool
	DecouplingRuleForContactDerived         bool

	AnyRepresentationAttemptConstructed bool
	AnyRepresentationAttemptCanonical   bool
	RepresentationCompleteRows          int
	RepresentationOpenRows              int
	ContactBetaRowsAllowed              int
	ContactZeroRowsProved               int
	ThresholdCorrectedBetaDerived       bool
	FullFiniteBetaMatchingTensorDerived bool

	LocalBundleObstructionDerived              bool
	RepresentationRowConstructionObstructed    bool
	ConstraintShortcutBlocked                  bool
	RepresentationOrConstraintDichotomyDerived bool
	BranchSelectorDerived                      bool

	ResidualNullityBefore  int
	ResidualNullityAfter   int
	ResidualSymmetryBroken bool

	PhysicalWeakAngleDerived bool
	FineStructureDerived     bool
	PhysicalMassesDerived    bool
	PhysicalScaleDerived     bool
	HiddenObservedInputUsed  bool

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
		cc, err := contactcohomology.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(cc)
	})
	return defaultValue, defaultErr
}

func Build(cc contactcohomology.Analysis) (Analysis, error) {
	if !cc.CohomologyObstructionDerived || !cc.NoCanonicalBRSTDifferentialUnderCurrentData {
		return Analysis{}, fmt.Errorf("Gate 115 requires Gate 114 cohomology obstruction")
	}
	if cc.ContactRows != 7 || cc.PositiveFiniteContactRows != 7 || cc.OpenContactRowsAfter != 7 {
		return Analysis{}, fmt.Errorf("Gate 115 requires seven unresolved positive contact cohomology rows")
	}
	if cc.ContactZeroRowsProved != 0 || cc.ContactBetaRowsAllowed != 0 || cc.ThresholdCorrectedBetaDerived || cc.FullFiniteBetaMatchingTensorDerived {
		return Analysis{}, fmt.Errorf("Gate 115 requires sealed contact beta rows")
	}
	if cc.ResidualNullityAfter != 3 || cc.HiddenObservedInputUsed || cc.PhysicalWeakAngleDerived || cc.FineStructureDerived || cc.PhysicalMassesDerived || cc.PhysicalScaleDerived {
		return Analysis{}, fmt.Errorf("Gate 115 refuses hidden physical input or changed residual nullity")
	}

	rows := buildRows(cc.Rows)
	attempts := buildBundleAttempts()
	criteria := buildCriteria()
	counts := countRows(rows)

	anyConstructed := false
	anyCanonical := false
	for _, a := range attempts {
		if a.Constructed {
			anyConstructed = true
		}
		if a.CanonicalUnderCurrentData && a.RepresentationRowDerived {
			anyCanonical = true
		}
	}

	truth := "Gate 115 attempts the local-bundle branch for the seven contact partial-overlap modes after the constraint/BRST shortcut was blocked. The finite positive contact carrier survives cohomology, but no base-space map, fiber/cocycle data, section map, contact gauge-representation row, Lorentz kinetic operator, mass activation unit, or decoupling rule is derived. Therefore the local bundle and representation-row construction is obstructed under current data; all seven contact modes remain representation-open and threshold beta matching stays sealed."

	return Analysis{
		ContactCohomology: cc,
		Rows:              rows,
		BundleAttempts:    attempts,
		Criteria:          criteria,

		ContactRows:               counts.contact,
		PositiveFiniteContactRows: counts.positiveContact,
		SurvivingCohomologyRows:   counts.survivesCohomology,
		OpenContactRowsBefore:     cc.OpenContactRowsAfter,
		OpenContactRowsAfter:      counts.open,

		LocalBundleConstructionAttempted:        true,
		FiniteCarrierAvailable:                  counts.contact == 7 && counts.positiveContact == 7 && counts.survivesCohomology == 7,
		BaseSpaceMapDerived:                     false,
		FiberDerived:                            false,
		TransitionFunctionsDerived:              false,
		SectionMapDerived:                       false,
		GaugeRepresentationForContactDerived:    false,
		HyperchargeForContactDerived:            false,
		ColorRepresentationForContactDerived:    false,
		LorentzKineticForContactDerived:         false,
		CanonicalNormalizationForContactDerived: false,
		MassActivationForContactDerived:         false,
		DecouplingRuleForContactDerived:         false,

		AnyRepresentationAttemptConstructed: anyConstructed,
		AnyRepresentationAttemptCanonical:   anyCanonical,
		RepresentationCompleteRows:          counts.complete,
		RepresentationOpenRows:              counts.open,
		ContactBetaRowsAllowed:              0,
		ContactZeroRowsProved:               0,
		ThresholdCorrectedBetaDerived:       false,
		FullFiniteBetaMatchingTensorDerived: false,

		LocalBundleObstructionDerived:              true,
		RepresentationRowConstructionObstructed:    true,
		ConstraintShortcutBlocked:                  cc.CohomologyObstructionDerived,
		RepresentationOrConstraintDichotomyDerived: false,
		BranchSelectorDerived:                      false,

		ResidualNullityBefore:  cc.ResidualNullityAfter,
		ResidualNullityAfter:   cc.ResidualNullityAfter,
		ResidualSymmetryBroken: false,

		PhysicalWeakAngleDerived: false,
		FineStructureDerived:     false,
		PhysicalMassesDerived:    false,
		PhysicalScaleDerived:     false,
		HiddenObservedInputUsed:  false,

		TruthStatement: truth,
		RejectedClaims: []string{
			"seven positive finite-overlap contact modes are automatically local continuum fields",
			"a finite label set is a spacetime bundle",
			"spectral ordering alone supplies transition functions or sections",
			"contact modes may be assigned singlet or doublet rows by convenience",
			"representation rows can enter beta matching without Lorentz kinetic, mass activation, and decoupling rules",
			"Gate 115 derives threshold-corrected beta coefficients, alpha, physical thetaW, M*, g_*, or masses",
		},
		RemainingUnknowns: []string{
			"U-35A-BASE: derive a base-space/support map for contact-overlap modes",
			"U-35B-FIBER: derive fiber type and transition/cocycle data",
			"U-35C-SECTION: derive local section variables from the finite contact carrier",
			"U-35D-REP: derive SU(3)c×SU(2)L×U(1)Y representation rows for contact modes",
			"U-35E-KINETIC: derive Lorentz-sign kinetic operators and residues for contact rows",
			"U-35F-MASS: derive activation/mass/decoupling rules before any Δb_i correction",
		},
		RecommendedNextGate: "Gate 116 — contact incidence/fiber functor search from Fano/contact geometry",
	}, nil
}

type rowCounts struct{ contact, positiveContact, survivesCohomology, open, complete int }

func buildRows(in []contactcohomology.ModeCohomologyRow) []BundleRow {
	rows := make([]BundleRow, 0, len(in))
	for _, r := range in {
		rows = append(rows, BundleRow{
			Name:                          r.Name,
			ModeKind:                      r.ModeKind,
			Value:                         r.Value,
			FiniteOverlapPositive:         r.FiniteOverlapPositive,
			SurvivesCohomology:            r.SurvivesZeroDifferentialCohomology,
			BaseSpaceMapDerived:           false,
			FiberDerived:                  false,
			TransitionFunctionsDerived:    false,
			SectionMapDerived:             false,
			GaugeRepresentationDerived:    false,
			HyperchargeDerived:            false,
			ColorRepresentationDerived:    false,
			LorentzKineticDerived:         false,
			CanonicalNormalizationDerived: false,
			MassActivationDerived:         false,
			DecouplingRuleDerived:         false,
			RepresentationStatus:          RepresentationOpen,
			RepresentationRowDerived:      false,
			CanEnterBetaTensor:            false,
			Reason:                        "finite positive contact cohomology class exists, but no local bundle, gauge representation, Lorentz kinetic row, mass activation, or decoupling rule is derived",
		})
	}
	sort.SliceStable(rows, func(i, j int) bool {
		if rows[i].Value == rows[j].Value {
			return rows[i].Name < rows[j].Name
		}
		return rows[i].Value < rows[j].Value
	})
	return rows
}

func buildBundleAttempts() []BundleAttempt {
	return []BundleAttempt{
		{
			Name:                      "finite seven-label contact carrier",
			Kind:                      FiniteLabelBundleAttempt,
			Constructed:               true,
			UsesExtraConvention:       false,
			CanonicalUnderCurrentData: true,
			RejectedAsPremature:       true,
			MissingTerms:              []string{"base-space support", "fiber type", "transition functions", "local sections", "Lorentz kinetic operator"},
			Detail:                    "the seven-mode carrier is finite and real, but a label carrier is not yet a local continuum bundle",
		},
		{
			Name:                      "spectral-value line-bundle attempt",
			Kind:                      SpectralLineBundleAttempt,
			Constructed:               true,
			UsesExtraConvention:       true,
			CanonicalUnderCurrentData: false,
			RejectedAsPremature:       true,
			MissingTerms:              []string{"orientation", "charts", "transition/cocycle law", "sections", "field normalization"},
			Detail:                    "ordering by overlap values can make lines, but it does not derive chart overlaps or a continuum base",
		},
		{
			Name:                      "assign contact modes as scalar singlets",
			Kind:                      SingletScalarAttempt,
			Constructed:               true,
			UsesExtraConvention:       true,
			CanonicalUnderCurrentData: false,
			RejectedAsPremature:       true,
			MissingTerms:              []string{"singlet representation theorem", "kinetic residue", "mass activation", "decoupling rule"},
			Detail:                    "singlet assignment is compatible with zero beta contribution but is not derived from finite contact data",
		},
		{
			Name:                      "assign contact modes as scalar doublets",
			Kind:                      DoubletScalarAttempt,
			Constructed:               true,
			UsesExtraConvention:       true,
			CanonicalUnderCurrentData: false,
			RejectedAsPremature:       true,
			MissingTerms:              []string{"SU(2)L fiber action", "hypercharge", "complex orientation", "mass activation", "decoupling rule"},
			Detail:                    "doublet assignment would change beta coefficients, so it is forbidden until derived",
		},
		{
			Name:                      "contact-distribution local lift",
			Kind:                      ContactDistributionAttempt,
			Constructed:               true,
			UsesExtraConvention:       true,
			CanonicalUnderCurrentData: false,
			RejectedAsPremature:       true,
			MissingTerms:              []string{"projection to spacetime", "locality theorem", "gauge representation", "propagator pole/residue", "threshold mass"},
			Detail:                    "contact geometry supplies a finite carrier, not a proven local spacetime field bundle for these partial-overlap modes",
		},
	}
}

func buildCriteria() []BundleCriterion {
	return []BundleCriterion{
		{Name: "finite positive contact carrier", Required: true, Derived: true, Detail: "seven positive finite-overlap modes survive Gate 114 cohomology"},
		{Name: "base-space/support map", Required: true, Derived: false, Detail: "no map from contact-overlap modes to spacetime support/charts is selected"},
		{Name: "fiber and transition/cocycle data", Required: true, Derived: false, Detail: "no local bundle fiber or transition functions are derived"},
		{Name: "local section variables", Required: true, Derived: false, Detail: "no finite-to-section map is available"},
		{Name: "gauge representation row", Required: true, Derived: false, Detail: "no SU(3)c×SU(2)L×U(1)Y row is derived for contact modes"},
		{Name: "Lorentz kinetic and residue", Required: true, Derived: false, Detail: "positive finite overlap is not a Lorentz-sign kinetic theorem"},
		{Name: "mass activation and decoupling", Required: true, Derived: false, Detail: "no threshold mass unit, activation predicate, or matching rule is selected"},
	}
}

func countRows(rows []BundleRow) rowCounts {
	c := rowCounts{contact: len(rows)}
	for _, r := range rows {
		if r.FiniteOverlapPositive {
			c.positiveContact++
		}
		if r.SurvivesCohomology {
			c.survivesCohomology++
		}
		if r.RepresentationStatus == RepresentationOpen {
			c.open++
		}
		if r.RepresentationStatus == RepresentationComplete {
			c.complete++
		}
	}
	return c
}

func FormatRows(rows []BundleRow, limit int) string {
	parts := make([]string, 0, len(rows))
	for i, r := range rows {
		if limit > 0 && i >= limit {
			parts = append(parts, fmt.Sprintf("... +%d", len(rows)-i))
			break
		}
		parts = append(parts, fmt.Sprintf("%s(local=%t,rep=%t,beta=%t)", r.Name, r.BaseSpaceMapDerived && r.SectionMapDerived, r.RepresentationRowDerived, r.CanEnterBetaTensor))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatAttempts(xs []BundleAttempt) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s(kind=%s,canonical=%t,row=%t,beta=%t)", x.Name, x.Kind, x.CanonicalUnderCurrentData, x.RepresentationRowDerived, x.BetaRowPermitted))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatCriteria(xs []BundleCriterion) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s=%t", x.Name, x.Derived))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func Join(xs []string) string { return strings.Join(xs, "; ") }
