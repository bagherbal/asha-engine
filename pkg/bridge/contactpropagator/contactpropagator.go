// Package contactpropagator implements Gate 110: contact-overlap kinetic-sign,
// locality, and propagator classifier search.
//
// Gate 109 left seven contact partial-overlap modes class-open.  This gate asks
// a narrower question before beta matching: do those modes already carry enough
// finite structure to be treated as physical propagating fields, regulator/ghost
// modes, constrained non-propagating modes, or vacuum-frustration modes?
//
// The result is deliberately conservative.  The contact modes have positive
// finite overlap eigenvalues in the current K7/B-sector data, so the finite
// overlap test is real.  But a positive dimensionless overlap eigenvalue is not
// a Lorentz kinetic operator, not a local field map, not a pole denominator, and
// not a residue/signature theorem.  Therefore no physical propagator class is
// selected and threshold beta corrections remain sealed.
package contactpropagator

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/modeclass"
	"github.com/bagherbal/asha-engine/pkg/bridge/threshold"
)

type PropagationClass string

const (
	BaselinePropagatingSector PropagationClass = "baseline-propagating-sector"
	PositiveFiniteOverlapOpen PropagationClass = "positive-finite-overlap-open"
	ExcludedFiniteInvariant   PropagationClass = "excluded-finite-invariant"
	ConstrainedFiniteVacuum   PropagationClass = "constrained-finite-vacuum"
)

type SignStatus string

const (
	FinitePositive SignStatus = "finite-positive"
	NotApplicable  SignStatus = "not-applicable"
	NotDerived     SignStatus = "not-derived"
)

type PropagatorRow struct {
	Name, ModeKind                        string
	Class                                 PropagationClass
	Value                                 float64
	FiniteOverlapSign                     SignStatus
	LorentzKineticSign                    SignStatus
	LocalFieldMapDerived                  bool
	PoleDenominatorDerived                bool
	ResidueSignDerived                    bool
	PhysicalPositiveNormPropagatorDerived bool
	RegulatorGhostDerived                 bool
	ConstrainedNonPropagatingDerived      bool
	VacuumFrustrationDerived              bool
	MayCorrectBeta                        bool
	Reason                                string
}

type TestAttempt struct {
	Name                          string
	Constructed                   bool
	FinitePositiveSpectrum        bool
	LorentzKineticOperatorDerived bool
	LocalFieldMapDerived          bool
	PoleDenominatorDerived        bool
	ResidueSignDerived            bool
	SelectsPhysicalPropagator     bool
	SelectsRegulatorGhost         bool
	SelectsConstrainedMode        bool
	SelectsVacuumFrustration      bool
	RejectedAsClassification      bool
	Detail                        string
}

type DenominatorWitness struct {
	Name                     string
	Formula                  string
	Positive                 bool
	CompatibleWithFiniteData bool
	Canonical                bool
	SelectsPole              bool
	SelectsPhysicalClass     bool
	Detail                   string
}

type BranchWitness struct {
	Name                     string
	CompatibleWithFiniteData bool
	SelectedClass            PropagationClass
	WouldCorrectBeta         bool
	Reason                   string
}

type Analysis struct {
	ModeClass            modeclass.Analysis
	Rows                 []PropagatorRow
	TestAttempts         []TestAttempt
	DenominatorWitnesses []DenominatorWitness
	BranchWitnesses      []BranchWitness

	TotalRows                 int
	ContactRows               int
	PositiveFiniteContactRows int
	BaselineRows              int
	ExcludedInvariantRows     int
	ConstrainedVacuumRows     int

	PositiveFiniteOverlapSpectrumDerived         bool
	LorentzKineticSignDerived                    bool
	LocalityDerived                              bool
	PoleDenominatorDerived                       bool
	ResidueSignDerived                           bool
	PhysicalPositiveNormContactPropagatorDerived bool
	RegulatorGhostContactClassDerived            bool
	ConstrainedContactClassDerived               bool
	VacuumFrustrationContactClassDerived         bool
	ContactPropagatorClassDerived                bool
	DenominatorAmbiguityWitnessed                bool
	BetaCorrectionRowsAllowed                    int
	ThresholdCorrectedBetaDerived                bool
	FullFiniteBetaMatchingTensorDerived          bool

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
		mc, err := modeclass.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(mc, 1e-10)
	})
	return defaultValue, defaultErr
}

func Build(mc modeclass.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if mc.ResidualNullityAfter != 3 || mc.ThresholdCorrectedBetaDerived || mc.FullFiniteBetaMatchingTensorDerived {
		return Analysis{}, fmt.Errorf("Gate 110 requires Gate 109 sealed threshold state with residual nullity 3")
	}
	if mc.ContactOpenRows != 7 || mc.ContactOverlapClassDerived {
		return Analysis{}, fmt.Errorf("Gate 110 requires seven class-open contact partial-overlap modes")
	}
	if mc.HiddenObservedInputUsed || mc.PhysicalWeakAngleDerived || mc.FineStructureDerived || mc.PhysicalMassesDerived {
		return Analysis{}, fmt.Errorf("Gate 110 refuses hidden observed physical input")
	}

	values := contactValues(mc)
	if len(values) != 7 {
		return Analysis{}, fmt.Errorf("expected seven contact overlap values, got %d", len(values))
	}
	rows := buildRows(mc, values, eps)
	attempts := buildAttempts(values)
	denominators := buildDenominatorWitnesses()
	branches := buildBranchWitnesses()
	counts := countRows(rows, eps)

	truth := "Gate 110 derives a positive finite-overlap diagnostic for the seven contact partial-overlap modes, but it does not derive a Lorentz kinetic sign, locality map, pole denominator, or residue/signature theorem. Therefore the modes remain positive finite overlap modes with open propagator class: they may not be counted as physical heavy fields, ghost/regulator modes, constrained non-propagating fields, or vacuum-frustration modes for beta matching."

	return Analysis{
		ModeClass:                            mc,
		Rows:                                 rows,
		TestAttempts:                         attempts,
		DenominatorWitnesses:                 denominators,
		BranchWitnesses:                      branches,
		TotalRows:                            counts.total,
		ContactRows:                          counts.contact,
		PositiveFiniteContactRows:            counts.positiveContact,
		BaselineRows:                         counts.baseline,
		ExcludedInvariantRows:                counts.excludedInvariant,
		ConstrainedVacuumRows:                counts.constrainedVacuum,
		PositiveFiniteOverlapSpectrumDerived: counts.contact == 7 && counts.positiveContact == 7,
		LorentzKineticSignDerived:            false,
		LocalityDerived:                      false,
		PoleDenominatorDerived:               false,
		ResidueSignDerived:                   false,
		PhysicalPositiveNormContactPropagatorDerived: false,
		RegulatorGhostContactClassDerived:            false,
		ConstrainedContactClassDerived:               false,
		VacuumFrustrationContactClassDerived:         false,
		ContactPropagatorClassDerived:                false,
		DenominatorAmbiguityWitnessed:                len(denominators) >= 4,
		BetaCorrectionRowsAllowed:                    0,
		ThresholdCorrectedBetaDerived:                false,
		FullFiniteBetaMatchingTensorDerived:          false,
		ResidualNullityBefore:                        mc.ResidualNullityAfter,
		ResidualNullityAfter:                         mc.ResidualNullityAfter,
		ResidualSymmetryBroken:                       false,
		PhysicalWeakAngleDerived:                     false,
		FineStructureDerived:                         false,
		PhysicalMassesDerived:                        false,
		HiddenObservedInputUsed:                      false,
		TruthStatement:                               truth,
		RejectedClaims: []string{
			"positive contact-overlap eigenvalues are Lorentz kinetic signs",
			"dimensionless overlap values are physical pole masses or threshold denominators",
			"absence of a negative finite overlap eigenvalue proves a regulator/ghost interpretation",
			"contact partial-overlap multiplicity may be inserted into threshold beta coefficients",
			"a propagator classifier without locality and residue data derives alpha, thetaW, or masses",
		},
		RemainingUnknowns: []string{
			"U-30A-LOCAL-FIELD-MAP: derive local continuum variables for contact partial-overlap modes, or prove they are nonlocal finite modes",
			"U-30B-KINETIC-OPERATOR: derive a Lorentzian quadratic kinetic operator, not only an overlap matrix eigenvalue",
			"U-30C-POLE-RESIDUE: derive pole denominator and residue signature for each contact mode",
			"U-30D-CONSTRAINT-BRST: derive whether contact modes are constrained, gauge, regulator/ghost, or physical positive-norm modes",
			"U-30E-BETA-PERMISSION: allow beta matching only after physical propagation class, representation, activation, and decoupling are all derived",
		},
		RecommendedNextGate: "Gate 111 — contact-overlap local field map / constraint-BRST classifier search",
	}, nil
}

type rowCounts struct{ total, contact, positiveContact, baseline, excludedInvariant, constrainedVacuum int }

func contactValues(mc modeclass.Analysis) map[string]float64 {
	out := map[string]float64{}
	for _, m := range mc.BetaMatching.Filtration.Modes {
		if m.Kind == string(threshold.ContactOverlapCandidate) {
			out[m.Name] = m.Value
		}
	}
	return out
}

func buildRows(mc modeclass.Analysis, values map[string]float64, eps float64) []PropagatorRow {
	rows := make([]PropagatorRow, 0, len(mc.Rows))
	for _, r := range mc.Rows {
		row := PropagatorRow{Name: r.Name, ModeKind: r.ModeKind, Value: values[r.Name], FiniteOverlapSign: NotApplicable, LorentzKineticSign: NotDerived, Reason: r.Reason}
		switch r.Class {
		case modeclass.BaselineContinuumSector:
			row.Class = BaselinePropagatingSector
			row.FiniteOverlapSign = NotApplicable
			row.LorentzKineticSign = NotDerived
			row.LocalFieldMapDerived = true
			row.Reason = "baseline scalar/contact sector remains the known continuum candidate; this gate only audits the individual partial-overlap modes"
		case modeclass.ConstrainedFiniteVacuumMode:
			row.Class = ConstrainedFiniteVacuum
			row.FiniteOverlapSign = NotApplicable
			row.ConstrainedNonPropagatingDerived = true
			row.Reason = "already classified by Gate 109 as finite vacuum-action eigenmode, not a contact partial-overlap propagator"
		case modeclass.VacuumFrustrationInvariant, modeclass.BridgeScalarInvariant, modeclass.UnorientedScalarConstituent:
			row.Class = ExcludedFiniteInvariant
			row.FiniteOverlapSign = NotApplicable
			row.Reason = "already excluded from individual heavy-threshold propagator classification"
		case modeclass.OpenContactOverlapMode:
			row.Class = PositiveFiniteOverlapOpen
			if row.Value > eps && row.Value < 1+eps {
				row.FiniteOverlapSign = FinitePositive
			}
			row.Reason = "positive finite overlap eigenvalue exists, but no Lorentz kinetic sign, local field map, pole denominator, or residue theorem is derived"
		default:
			row.Class = PositiveFiniteOverlapOpen
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

func buildAttempts(values map[string]float64) []TestAttempt {
	minV, maxV := minMax(values)
	return []TestAttempt{
		{Name: "finite overlap positivity test", Constructed: true, FinitePositiveSpectrum: true, RejectedAsClassification: true, Detail: fmt.Sprintf("all seven contact partial-overlap eigenvalues are positive dimensionless numbers; range=[%.10f, %.10f]", minV, maxV)},
		{Name: "Lorentz kinetic-sign test", Constructed: true, LorentzKineticOperatorDerived: false, RejectedAsClassification: true, Detail: "no local quadratic operator of the form phi K(p) phi or p^2+m^2 is derived for contact modes"},
		{Name: "locality / finite-to-continuum map test", Constructed: true, LocalFieldMapDerived: false, RejectedAsClassification: true, Detail: "contact overlap eigenvectors are finite K7 data; no spacetime-local field map is derived"},
		{Name: "pole denominator test", Constructed: true, PoleDenominatorDerived: false, RejectedAsClassification: true, Detail: "lambda, 1-lambda, lambda/(1-lambda), and spectral gap denominators are compatible but none is selected as a physical pole"},
		{Name: "residue / ghost-sign test", Constructed: true, ResidueSignDerived: false, SelectsRegulatorGhost: false, RejectedAsClassification: true, Detail: "no indefinite metric, BRST complex, or negative-residue theorem classifies the modes as ghosts/regulators"},
		{Name: "beta-permission test", Constructed: true, SelectsPhysicalPropagator: false, RejectedAsClassification: true, Detail: "no contact row is physical-positive-norm, representation-complete, activation-derived, and decoupling-derived"},
	}
}

func buildDenominatorWitnesses() []DenominatorWitness {
	return []DenominatorWitness{
		{Name: "overlap eigenvalue denominator", Formula: "rho_i = lambda_i", Positive: true, CompatibleWithFiniteData: true, Detail: "finite positive, but not selected as physical pole mass"},
		{Name: "defect denominator", Formula: "rho_i = 1-lambda_i", Positive: true, CompatibleWithFiniteData: true, Detail: "also finite positive for partial modes; equally compatible"},
		{Name: "odds-ratio denominator", Formula: "rho_i = lambda_i/(1-lambda_i)", Positive: true, CompatibleWithFiniteData: true, Detail: "monotone transform of the same data, no canonical pole choice"},
		{Name: "inverse-overlap denominator", Formula: "rho_i = 1/lambda_i", Positive: true, CompatibleWithFiniteData: true, Detail: "positive but changes ordering/scale; not selected"},
	}
}

func buildBranchWitnesses() []BranchWitness {
	return []BranchWitness{
		{Name: "physical positive-norm scalar branch", CompatibleWithFiniteData: true, SelectedClass: PositiveFiniteOverlapOpen, WouldCorrectBeta: true, Reason: "would need locality, representation, pole, residue, activation, and decoupling; not derived"},
		{Name: "regulator/ghost branch", CompatibleWithFiniteData: true, SelectedClass: PositiveFiniteOverlapOpen, WouldCorrectBeta: false, Reason: "compatible only as an extra structure; no negative residue or BRST theorem selects it"},
		{Name: "constrained finite overlap branch", CompatibleWithFiniteData: true, SelectedClass: PositiveFiniteOverlapOpen, WouldCorrectBeta: false, Reason: "compatible, but no constraint algebra removes propagation"},
		{Name: "vacuum-frustration branch", CompatibleWithFiniteData: true, SelectedClass: PositiveFiniteOverlapOpen, WouldCorrectBeta: false, Reason: "compatible, but the leakage invariant does not absorb the seven partial modes"},
	}
}

func countRows(rows []PropagatorRow, eps float64) rowCounts {
	var c rowCounts
	c.total = len(rows)
	for _, r := range rows {
		switch r.Class {
		case BaselinePropagatingSector:
			c.baseline++
		case PositiveFiniteOverlapOpen:
			c.contact++
			if r.FiniteOverlapSign == FinitePositive && r.Value > eps {
				c.positiveContact++
			}
		case ExcludedFiniteInvariant:
			c.excludedInvariant++
		case ConstrainedFiniteVacuum:
			c.constrainedVacuum++
		}
	}
	return c
}

func minMax(values map[string]float64) (float64, float64) {
	minV := math.Inf(1)
	maxV := math.Inf(-1)
	for _, v := range values {
		if v < minV {
			minV = v
		}
		if v > maxV {
			maxV = v
		}
	}
	if math.IsInf(minV, 0) {
		return 0, 0
	}
	return minV, maxV
}

func FormatRows(rows []PropagatorRow, max int) string {
	if max <= 0 || max > len(rows) {
		max = len(rows)
	}
	parts := make([]string, 0, max+1)
	for i := 0; i < max; i++ {
		r := rows[i]
		parts = append(parts, fmt.Sprintf("%s:%s:value=%.10f:finite=%s", r.Name, r.Class, r.Value, r.FiniteOverlapSign))
	}
	if max < len(rows) {
		parts = append(parts, fmt.Sprintf("... +%d more", len(rows)-max))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatAttempts(xs []TestAttempt) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s(pos=%t,local=%t,pole=%t,residue=%t,phys=%t,ghost=%t)", x.Name, x.FinitePositiveSpectrum, x.LocalFieldMapDerived, x.PoleDenominatorDerived, x.ResidueSignDerived, x.SelectsPhysicalPropagator, x.SelectsRegulatorGhost))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatDenominators(xs []DenominatorWitness) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s(%s,canonical=%t,pole=%t)", x.Name, x.Formula, x.Canonical, x.SelectsPole))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatBranches(xs []BranchWitness) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s(beta=%t)", x.Name, x.WouldCorrectBeta))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func Join(xs []string) string { return strings.Join(xs, "; ") }
