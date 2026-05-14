// Package contactfieldmap implements Gate 111: contact-overlap local field map /
// constraint-BRST classifier search.
//
// Gate 110 proved a useful but limited fact: the seven contact partial-overlap
// modes have positive finite overlap eigenvalues.  This gate asks whether those
// finite modes can already be lifted into local continuum variables, or instead
// classified as constrained/BRST/non-propagating modes.
//
// The result is conservative and falsifiable.  A local field map would require a
// spacetime support assignment, a Lorentz quadratic kinetic operator, a gauge
// representation, a pole/residue theorem, and an invertible or controlled
// finite-to-continuum map.  A BRST/constraint classification would require a
// constraint generator, ghost grading, nilpotent differential Q^2=0, and a
// cancellation/supertrace ledger.  The current finite data provides none of
// those structures for the seven contact partial-overlap modes, so their local
// field class remains open and threshold beta matching stays sealed.
package contactfieldmap

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/contactpropagator"
)

type FieldClass string

const (
	BaselineLocalSector      FieldClass = "baseline-local-sector"
	FiniteOverlapMapOpen     FieldClass = "finite-overlap-local-map-open"
	ExcludedFiniteDiagnostic FieldClass = "excluded-finite-diagnostic"
	ConstrainedVacuumLedger  FieldClass = "constrained-vacuum-ledger"
)

type EvidenceStatus string

const (
	EvidenceDerived   EvidenceStatus = "derived"
	EvidenceMissing   EvidenceStatus = "missing"
	EvidenceNotNeeded EvidenceStatus = "not-needed"
)

type FieldMapRow struct {
	Name, ModeKind string
	Class          FieldClass
	Value          float64

	FiniteOverlapPositive         bool
	LocalCoordinateDerived        bool
	SpacetimeSupportDerived       bool
	LorentzKineticOperatorDerived bool
	GaugeRepresentationDerived    bool
	CanonicalNormalizationDerived bool
	PoleResidueDerived            bool
	InvertibleFieldMapDerived     bool

	ConstraintGeneratorDerived    bool
	GhostGradingDerived           bool
	NilpotentBRSTDerived          bool
	BRSTPairingDerived            bool
	SupertraceCancellationDerived bool

	PhysicalLocalFieldDerived        bool
	ConstrainedNonPropagatingDerived bool
	RegulatorGhostDerived            bool
	VacuumFrustrationDerived         bool
	MayCorrectBeta                   bool
	Reason                           string
}

type LocalityCriterion struct {
	Name                     string
	RequiredForPhysicalField bool
	Derived                  bool
	Detail                   string
}

type BRSTCriterion struct {
	Name                 string
	RequiredForBRSTClass bool
	Derived              bool
	Detail               string
}

type BranchWitness struct {
	Name                     string
	CompatibleWithFiniteData bool
	RequiresExtraStructure   bool
	WouldCorrectBeta         bool
	Detail                   string
}

type Analysis struct {
	ContactPropagator contactpropagator.Analysis

	Rows             []FieldMapRow
	LocalityCriteria []LocalityCriterion
	BRSTCriteria     []BRSTCriterion
	BranchWitnesses  []BranchWitness

	TotalRows                 int
	ContactRows               int
	PositiveFiniteContactRows int
	BaselineRows              int
	ExcludedDiagnosticRows    int
	ConstrainedVacuumRows     int

	LocalFieldMapCandidateConstructed       bool
	LocalCoordinateDerived                  bool
	SpacetimeSupportDerived                 bool
	LorentzKineticOperatorDerived           bool
	GaugeRepresentationForContactDerived    bool
	CanonicalNormalizationForContactDerived bool
	PoleResidueForContactDerived            bool
	InvertibleFieldMapDerived               bool
	PhysicalLocalContactFieldsDerived       bool

	ConstraintClassifierConstructed      bool
	ConstraintGeneratorDerived           bool
	GhostGradingDerived                  bool
	NilpotentBRSTDerived                 bool
	BRSTPairingDerived                   bool
	SupertraceCancellationDerived        bool
	ConstrainedContactClassDerived       bool
	RegulatorGhostContactClassDerived    bool
	VacuumFrustrationContactClassDerived bool
	ContactFieldClassDerived             bool

	BetaCorrectionRowsAllowed           int
	ThresholdCorrectedBetaDerived       bool
	FullFiniteBetaMatchingTensorDerived bool

	ResidualNullityBefore  int
	ResidualNullityAfter   int
	ResidualSymmetryBroken bool

	PhysicalWeakAngleDerived bool
	FineStructureDerived     bool
	PhysicalMassesDerived    bool
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
		cp, err := contactpropagator.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(cp)
	})
	return defaultValue, defaultErr
}

func Build(cp contactpropagator.Analysis) (Analysis, error) {
	if cp.ResidualNullityAfter != 3 || cp.ThresholdCorrectedBetaDerived || cp.FullFiniteBetaMatchingTensorDerived {
		return Analysis{}, fmt.Errorf("Gate 111 requires Gate 110 sealed propagator state with residual nullity 3")
	}
	if cp.ContactRows != 7 || cp.PositiveFiniteContactRows != 7 || !cp.PositiveFiniteOverlapSpectrumDerived {
		return Analysis{}, fmt.Errorf("Gate 111 requires seven positive finite-overlap contact modes")
	}
	if cp.ContactPropagatorClassDerived || cp.PhysicalPositiveNormContactPropagatorDerived || cp.RegulatorGhostContactClassDerived || cp.ConstrainedContactClassDerived || cp.VacuumFrustrationContactClassDerived {
		return Analysis{}, fmt.Errorf("Gate 111 requires contact propagator class to remain open")
	}
	if cp.HiddenObservedInputUsed || cp.PhysicalWeakAngleDerived || cp.FineStructureDerived || cp.PhysicalMassesDerived {
		return Analysis{}, fmt.Errorf("Gate 111 refuses hidden observed physical input")
	}

	rows := buildRows(cp.Rows)
	locality := buildLocalityCriteria()
	brst := buildBRSTCriteria()
	branches := buildBranchWitnesses()
	counts := countRows(rows)

	truth := "Gate 111 constructs the local-field-map and constraint/BRST classifier tests for the seven contact partial-overlap modes. The modes have positive finite overlap data, but no spacetime support map, Lorentz kinetic operator, contact gauge-representation row, pole/residue theorem, constraint generator, ghost grading, nilpotent BRST operator, or cancellation ledger is derived. Therefore the modes remain finite-overlap local-map-open modes and may not be used as threshold beta rows."

	return Analysis{
		ContactPropagator:                       cp,
		Rows:                                    rows,
		LocalityCriteria:                        locality,
		BRSTCriteria:                            brst,
		BranchWitnesses:                         branches,
		TotalRows:                               counts.total,
		ContactRows:                             counts.contact,
		PositiveFiniteContactRows:               counts.positiveContact,
		BaselineRows:                            counts.baseline,
		ExcludedDiagnosticRows:                  counts.excluded,
		ConstrainedVacuumRows:                   counts.constrainedVacuum,
		LocalFieldMapCandidateConstructed:       true,
		LocalCoordinateDerived:                  false,
		SpacetimeSupportDerived:                 false,
		LorentzKineticOperatorDerived:           false,
		GaugeRepresentationForContactDerived:    false,
		CanonicalNormalizationForContactDerived: false,
		PoleResidueForContactDerived:            false,
		InvertibleFieldMapDerived:               false,
		PhysicalLocalContactFieldsDerived:       false,
		ConstraintClassifierConstructed:         true,
		ConstraintGeneratorDerived:              false,
		GhostGradingDerived:                     false,
		NilpotentBRSTDerived:                    false,
		BRSTPairingDerived:                      false,
		SupertraceCancellationDerived:           false,
		ConstrainedContactClassDerived:          false,
		RegulatorGhostContactClassDerived:       false,
		VacuumFrustrationContactClassDerived:    false,
		ContactFieldClassDerived:                false,
		BetaCorrectionRowsAllowed:               0,
		ThresholdCorrectedBetaDerived:           false,
		FullFiniteBetaMatchingTensorDerived:     false,
		ResidualNullityBefore:                   cp.ResidualNullityAfter,
		ResidualNullityAfter:                    cp.ResidualNullityAfter,
		ResidualSymmetryBroken:                  false,
		PhysicalWeakAngleDerived:                false,
		FineStructureDerived:                    false,
		PhysicalMassesDerived:                   false,
		HiddenObservedInputUsed:                 false,
		TruthStatement:                          truth,
		RejectedClaims: []string{
			"a positive contact-overlap eigenvalue is a local spacetime field",
			"finite overlap modes can be inserted as scalar doublet thresholds without a local field map",
			"absence of a negative finite eigenvalue derives a BRST/ghost regulator class",
			"a BRST interpretation can be claimed without Q, ghost grading, Q^2=0, and a cancellation ledger",
			"contact local-field classification derives alpha, physical thetaW, boundary scale, or masses",
		},
		RemainingUnknowns: []string{
			"U-31A-SPACETIME-SUPPORT: derive a map from each contact partial-overlap eigenmode to local spacetime support",
			"U-31B-LOCAL-KINETIC-ACTION: derive a Lorentzian quadratic action and canonical normalization for any lifted contact field",
			"U-31C-GAUGE-REPRESENTATION: derive SU(3)c×SU(2)L×U(1)Y representation rows for lifted contact modes",
			"U-31D-CONSTRAINT-BRST-COMPLEX: derive constraint generators, ghost grading, nilpotent Q, pairing, and supertrace cancellation if modes are nonphysical",
			"U-31E-THRESHOLD-PERMISSION: permit beta matching only after local field/BRST class, representation, activation, mass unit, and decoupling are all derived",
		},
		RecommendedNextGate: "Gate 112 — contact-overlap representation-or-constraint dichotomy / beta-permission firewall",
	}, nil
}

type rowCounts struct{ total, contact, positiveContact, baseline, excluded, constrainedVacuum int }

func buildRows(in []contactpropagator.PropagatorRow) []FieldMapRow {
	rows := make([]FieldMapRow, 0, len(in))
	for _, r := range in {
		row := FieldMapRow{Name: r.Name, ModeKind: r.ModeKind, Value: r.Value, Reason: r.Reason}
		switch r.Class {
		case contactpropagator.BaselinePropagatingSector:
			row.Class = BaselineLocalSector
			row.LocalCoordinateDerived = true
			row.SpacetimeSupportDerived = true
			row.LorentzKineticOperatorDerived = true
			row.GaugeRepresentationDerived = true
			row.CanonicalNormalizationDerived = true
			row.PoleResidueDerived = true
			row.InvertibleFieldMapDerived = true
			row.PhysicalLocalFieldDerived = true
			row.Reason = "baseline scalar/contact aggregate is the already-known continuum sector; Gate 111 audits only the seven partial-overlap modes"
		case contactpropagator.ConstrainedFiniteVacuum:
			row.Class = ConstrainedVacuumLedger
			row.ConstrainedNonPropagatingDerived = true
			row.Reason = "B-sector gap remains classified as a constrained finite vacuum-action ledger entry, not a contact local field"
		case contactpropagator.ExcludedFiniteInvariant:
			row.Class = ExcludedFiniteDiagnostic
			row.Reason = "finite diagnostic already excluded from individual local-field/threshold status"
		case contactpropagator.PositiveFiniteOverlapOpen:
			row.Class = FiniteOverlapMapOpen
			row.FiniteOverlapPositive = r.FiniteOverlapSign == contactpropagator.FinitePositive
			row.Reason = "positive finite overlap exists, but no local spacetime map or constraint/BRST complex is derived"
		default:
			row.Class = FiniteOverlapMapOpen
		}
		rows = append(rows, row)
	}
	sort.SliceStable(rows, func(i, j int) bool {
		if rows[i].Class == rows[j].Class {
			return rows[i].Name < rows[j].Name
		}
		return rows[i].Class < rows[j].Class
	})
	return rows
}

func buildLocalityCriteria() []LocalityCriterion {
	return []LocalityCriterion{
		{Name: "finite eigenmode coordinate", RequiredForPhysicalField: true, Derived: false, Detail: "contact overlap eigenvectors are finite K7 data but are not local spacetime field coordinates"},
		{Name: "spacetime support / locality map", RequiredForPhysicalField: true, Derived: false, Detail: "no map assigns each partial-overlap mode to local support or a continuum bundle section"},
		{Name: "Lorentz quadratic kinetic action", RequiredForPhysicalField: true, Derived: false, Detail: "no p²+m² or Lorentzian differential operator is derived for contact modes"},
		{Name: "gauge representation row", RequiredForPhysicalField: true, Derived: false, Detail: "no SU(3)c×SU(2)L×U(1)Y representation is selected for any partial-overlap mode"},
		{Name: "pole/residue theorem", RequiredForPhysicalField: true, Derived: false, Detail: "no pole denominator or positive residue theorem exists for these modes"},
		{Name: "mass unit and decoupling rule", RequiredForPhysicalField: true, Derived: false, Detail: "dimensionless finite overlap values are not physical threshold masses"},
	}
}

func buildBRSTCriteria() []BRSTCriterion {
	return []BRSTCriterion{
		{Name: "constraint generator", RequiredForBRSTClass: true, Derived: false, Detail: "no constraint operator removes the seven contact modes from propagation"},
		{Name: "ghost number grading", RequiredForBRSTClass: true, Derived: false, Detail: "no ghost/antighost grading is defined on the contact partial-overlap carrier"},
		{Name: "nilpotent BRST differential", RequiredForBRSTClass: true, Derived: false, Detail: "no Q with Q²=0 is constructed"},
		{Name: "BRST pair/doublet assignment", RequiredForBRSTClass: true, Derived: false, Detail: "no pairing or quartet cancellation is selected"},
		{Name: "supertrace/cancellation ledger", RequiredForBRSTClass: true, Derived: false, Detail: "no cancellation theorem shows the contact modes are regulators or gauge artifacts"},
	}
}

func buildBranchWitnesses() []BranchWitness {
	return []BranchWitness{
		{Name: "local physical scalar/contact fields", CompatibleWithFiniteData: true, RequiresExtraStructure: true, WouldCorrectBeta: true, Detail: "compatible only if locality, representation, pole/residue, activation, and decoupling are added"},
		{Name: "nonlocal finite overlap coordinates", CompatibleWithFiniteData: true, RequiresExtraStructure: true, WouldCorrectBeta: false, Detail: "compatible with current data, but does not produce continuum threshold rows"},
		{Name: "constrained non-propagating modes", CompatibleWithFiniteData: true, RequiresExtraStructure: true, WouldCorrectBeta: false, Detail: "compatible if a constraint algebra removes propagation"},
		{Name: "BRST/regulator modes", CompatibleWithFiniteData: true, RequiresExtraStructure: true, WouldCorrectBeta: false, Detail: "compatible only with ghost grading, nilpotent Q, and cancellation ledger"},
		{Name: "vacuum-frustration descendants", CompatibleWithFiniteData: true, RequiresExtraStructure: true, WouldCorrectBeta: false, Detail: "compatible, but not selected by the existing leakage invariant"},
	}
}

func countRows(rows []FieldMapRow) rowCounts {
	var c rowCounts
	c.total = len(rows)
	for _, r := range rows {
		switch r.Class {
		case BaselineLocalSector:
			c.baseline++
		case FiniteOverlapMapOpen:
			c.contact++
			if r.FiniteOverlapPositive {
				c.positiveContact++
			}
		case ExcludedFiniteDiagnostic:
			c.excluded++
		case ConstrainedVacuumLedger:
			c.constrainedVacuum++
		}
	}
	return c
}

func FormatRows(rows []FieldMapRow, max int) string {
	if max <= 0 || max > len(rows) {
		max = len(rows)
	}
	parts := make([]string, 0, max+1)
	for i := 0; i < max; i++ {
		r := rows[i]
		parts = append(parts, fmt.Sprintf("%s:%s:value=%.10f:local=%t:brst=%t", r.Name, r.Class, r.Value, r.PhysicalLocalFieldDerived, r.NilpotentBRSTDerived))
	}
	if max < len(rows) {
		parts = append(parts, fmt.Sprintf("... +%d more", len(rows)-max))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatLocality(xs []LocalityCriterion) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s(derived=%t)", x.Name, x.Derived))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatBRST(xs []BRSTCriterion) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s(derived=%t)", x.Name, x.Derived))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatBranches(xs []BranchWitness) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s(beta=%t,extra=%t)", x.Name, x.WouldCorrectBeta, x.RequiresExtraStructure))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func Join(xs []string) string { return strings.Join(xs, "; ") }
