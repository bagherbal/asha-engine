package modeclass

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/betamatching"
	"github.com/bagherbal/asha-engine/pkg/bridge/threshold"
)

type ModeClass string

const (
	BaselineContinuumSector     ModeClass = "baseline-continuum-sector"
	UnorientedScalarConstituent ModeClass = "unoriented-scalar-constituent"
	BridgeScalarInvariant       ModeClass = "bridge-scalar-invariant"
	ConstrainedFiniteVacuumMode ModeClass = "constrained-finite-vacuum-mode"
	OpenContactOverlapMode      ModeClass = "open-contact-overlap-mode"
	VacuumFrustrationInvariant  ModeClass = "vacuum-frustration-invariant"
)

type ClassStatus string

const (
	ClassDerived  ClassStatus = "derived"
	ClassOpen     ClassStatus = "open"
	ClassExcluded ClassStatus = "excluded"
)

type ClassRow struct {
	Name, ModeKind                                                                             string
	Class                                                                                      ModeClass
	Status                                                                                     ClassStatus
	PhysicalFieldDerived, RegulatorDerived, ConstrainedFiniteDerived, VacuumFrustrationDerived bool
	GaugeRepresentationRequired, GaugeRepresentationAvailable                                  bool
	MassUnitDerived, ActivationRuleDerived, DecouplingRuleDerived                              bool
	MayCorrectBeta, ExcludedFromThresholdBeta                                                  bool
	Reason                                                                                     string
}

type ClassifierAttempt struct {
	Name                                                                                             string
	Constructed, SelectsClass, RejectsPhysicalThreshold, RequiresExtraChoice, CanCorrectBeta, Sealed bool
	Detail                                                                                           string
}

type AmbiguityWitness struct {
	Name                     string
	CompatibleWithFiniteData bool
	SelectedClass            ModeClass
	WouldCorrectBeta         bool
	Reason                   string
}

type Analysis struct {
	BetaMatching                                                                                   betamatching.Analysis
	Rows                                                                                           []ClassRow
	ClassifierAttempts                                                                             []ClassifierAttempt
	AmbiguityWitnesses                                                                             []AmbiguityWitness
	TotalRows                                                                                      int
	BaselineSectorRows                                                                             int
	UnorientedScalarRows                                                                           int
	BridgeScalarRows                                                                               int
	BGapRows                                                                                       int
	ContactOpenRows                                                                                int
	VacuumFrustrationRows                                                                          int
	ConstrainedFiniteRows                                                                          int
	DerivedClassRows                                                                               int
	OpenClassRows                                                                                  int
	PhysicalHeavyThresholdRows                                                                     int
	RegulatorRowsDerived                                                                           int
	BetaCorrectionRowsAllowed                                                                      int
	BGapClassifiedAsConstrainedFinite                                                              bool
	BGapExcludedFromThresholdBeta                                                                  bool
	ContactOverlapClassDerived                                                                     bool
	ContactOverlapAmbiguityWitnessed                                                               bool
	ContactOverlapPhysicalFieldDerived                                                             bool
	ContactOverlapRegulatorDerived                                                                 bool
	ContactOverlapVacuumClassDerived                                                               bool
	ActivationPredicateDerived                                                                     bool
	DecouplingClassDerived                                                                         bool
	PhysicalMassUnitDerived                                                                        bool
	ThresholdMassSpectrumDerived                                                                   bool
	ThresholdCorrectedBetaDerived                                                                  bool
	FullFiniteBetaMatchingTensorDerived                                                            bool
	ResidualNullityBefore, ResidualNullityAfter                                                    int
	ResidualSymmetryBroken                                                                         bool
	PhysicalWeakAngleDerived, FineStructureDerived, PhysicalMassesDerived, HiddenObservedInputUsed bool
	TruthStatement                                                                                 string
	RejectedClaims, RemainingUnknowns                                                              []string
	RecommendedNextGate                                                                            string
}

var defaultOnce sync.Once
var defaultValue Analysis
var defaultErr error

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		bm, err := betamatching.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(bm)
	})
	return defaultValue, defaultErr
}

func Build(bm betamatching.Analysis) (Analysis, error) {
	if bm.ResidualNullityAfter != 3 || bm.ThresholdCorrectedBetaDerived || bm.FullFiniteBetaMatchingTensorDerived {
		return Analysis{}, fmt.Errorf("Gate 109 requires Gate 108 beta-matching obstruction with residual nullity 3")
	}
	if bm.HiddenObservedInputUsed || bm.PhysicalWeakAngleDerived || bm.FineStructureDerived || bm.PhysicalMassesDerived {
		return Analysis{}, fmt.Errorf("Gate 109 refuses hidden observed physical input")
	}
	if len(bm.Rows) == 0 || bm.IncompleteOpenRows == 0 {
		return Analysis{}, fmt.Errorf("Gate 109 requires representation-incomplete open rows from Gate 108")
	}
	rows := classifyRows(bm.Rows)
	attempts := buildAttempts()
	witnesses := buildAmbiguityWitnesses()
	counts := countRows(rows)
	truth := "Gate 109 classifies what can be classified before threshold beta matching. The B-sector first spectral gap is a constrained finite vacuum-action eigenmode and is excluded from heavy-threshold beta corrections. The seven contact partial-overlap modes remain open: current finite data does not decide whether they are physical fields, regulator modes, constrained overlap modes, or vacuum-frustration modes. Therefore no activation predicate, decoupling class, threshold mass spectrum, or threshold-corrected beta tensor is derived."
	return Analysis{BetaMatching: bm, Rows: rows, ClassifierAttempts: attempts, AmbiguityWitnesses: witnesses, TotalRows: counts.total, BaselineSectorRows: counts.baseline, UnorientedScalarRows: counts.unoriented, BridgeScalarRows: counts.bridge, BGapRows: counts.bgap, ContactOpenRows: counts.contactOpen, VacuumFrustrationRows: counts.vacuum, ConstrainedFiniteRows: counts.constrained, DerivedClassRows: counts.derived, OpenClassRows: counts.open, PhysicalHeavyThresholdRows: counts.physical, RegulatorRowsDerived: counts.regulator, BetaCorrectionRowsAllowed: counts.betaAllowed, BGapClassifiedAsConstrainedFinite: counts.bgap == 1 && counts.constrained >= 1, BGapExcludedFromThresholdBeta: bGapExcluded(rows), ContactOverlapClassDerived: false, ContactOverlapAmbiguityWitnessed: len(witnesses) >= 4, ContactOverlapPhysicalFieldDerived: false, ContactOverlapRegulatorDerived: false, ContactOverlapVacuumClassDerived: false, ActivationPredicateDerived: false, DecouplingClassDerived: false, PhysicalMassUnitDerived: false, ThresholdMassSpectrumDerived: false, ThresholdCorrectedBetaDerived: false, FullFiniteBetaMatchingTensorDerived: false, ResidualNullityBefore: bm.ResidualNullityAfter, ResidualNullityAfter: bm.ResidualNullityAfter, ResidualSymmetryBroken: false, PhysicalWeakAngleDerived: false, FineStructureDerived: false, PhysicalMassesDerived: false, HiddenObservedInputUsed: false, TruthStatement: truth, RejectedClaims: []string{"the B-sector finite spectral gap is a heavy continuum particle mass", "contact overlap multiplicity alone determines threshold beta rows", "open contact modes may be counted as scalar doublets, singlets, or regulators without a classifier theorem", "a dimensionless finite eigenvalue supplies a physical threshold scale", "mode-classification alone derives alpha, thetaW, W/Z masses, Higgs scale, or fermion masses"}, RemainingUnknowns: []string{"U-29A-CONTACT-CLASS: derive whether contact partial-overlap modes are physical fields, regulators, constrained finite modes, or frustration invariants", "U-29B-KINETIC-SIGN: derive a kinetic-sign/propagator test for contact overlap modes", "U-29C-LOCALITY: derive whether contact overlap modes admit local continuum field variables", "U-29D-REPRESENTATION: derive SU(3)c×SU(2)L×U(1)Y representation rows for any contact mode classified as physical", "U-29E-DECOUPLING: derive a mass unit, activation predicate, and matching rule before any Δb_i correction is allowed"}, RecommendedNextGate: "Gate 110 — contact-overlap kinetic-sign / locality / propagator classifier search"}, nil
}

type rowCounts struct{ total, baseline, unoriented, bridge, bgap, contactOpen, vacuum, constrained, derived, open, physical, regulator, betaAllowed int }

func classifyRows(rows []betamatching.MatchingRow) []ClassRow {
	out := make([]ClassRow, 0, len(rows))
	for _, r := range rows {
		c := ClassRow{Name: r.Name, ModeKind: r.ModeKind, GaugeRepresentationAvailable: r.RepresentationComplete, GaugeRepresentationRequired: true, Reason: r.Reason}
		if r.Name == "scalar/contact active sector aggregate" {
			c.Class = BaselineContinuumSector
			c.Status = ClassDerived
			c.PhysicalFieldDerived = true
			c.GaugeRepresentationRequired = false
			c.ExcludedFromThresholdBeta = true
			c.Reason = "derived as the baseline complex scalar doublet sector; not a heavy threshold row"
		} else {
			switch r.ModeKind {
			case string(threshold.ScalarActiveCandidate):
				c.Class = UnorientedScalarConstituent
				c.Status = ClassDerived
				c.ExcludedFromThresholdBeta = true
				c.Reason = "individual real scalar/contact constituents are not separately oriented as heavy threshold fields"
			case string(threshold.RadialCandidate):
				c.Class = BridgeScalarInvariant
				c.Status = ClassDerived
				c.ExcludedFromThresholdBeta = true
				c.Reason = "radial scalar response is a bridge-level singlet invariant, not an activated heavy threshold"
			case string(threshold.BGapCandidate):
				c.Class = ConstrainedFiniteVacuumMode
				c.Status = ClassDerived
				c.ConstrainedFiniteDerived = true
				c.ExcludedFromThresholdBeta = true
				c.Reason = "B-sector gap is an eigenvalue of the finite vacuum action; without locality, gauge representation, or decoupling law it is a constrained finite vacuum mode, not a continuum threshold"
			case string(threshold.ContactOverlapCandidate):
				c.Class = OpenContactOverlapMode
				c.Status = ClassOpen
				c.Reason = "contact partial-overlap mode has finite spectral data but no derived physical/regulator/frustration classifier"
			case string(threshold.LeakageCandidate):
				c.Class = VacuumFrustrationInvariant
				c.Status = ClassExcluded
				c.VacuumFrustrationDerived = true
				c.GaugeRepresentationRequired = false
				c.ExcludedFromThresholdBeta = true
				c.Reason = "bare leakage is already classified as a vacuum-frustration invariant"
			default:
				c.Class = OpenContactOverlapMode
				c.Status = ClassOpen
			}
		}
		out = append(out, c)
	}
	sort.SliceStable(out, func(i, j int) bool {
		if out[i].Class == out[j].Class {
			return out[i].Name < out[j].Name
		}
		return out[i].Class < out[j].Class
	})
	return out
}

func buildAttempts() []ClassifierAttempt {
	return []ClassifierAttempt{{Name: "B-sector vacuum-action classifier", Constructed: true, SelectsClass: true, RejectsPhysicalThreshold: true, Sealed: true, Detail: "finite positive B-spectrum is real, but it belongs to the vacuum action ledger and lacks continuum locality/representation/decoupling data"}, {Name: "contact-overlap physical-field classifier", Constructed: true, RequiresExtraChoice: true, Detail: "would require kinetic sign, locality, gauge representation, and mass unit; none are selected"}, {Name: "contact-overlap regulator classifier", Constructed: true, RequiresExtraChoice: true, Detail: "regulator interpretation is compatible, but no BRST/ghost/signature or cancellation theorem selects it"}, {Name: "contact-overlap vacuum-frustration classifier", Constructed: true, RequiresExtraChoice: true, Detail: "frustration interpretation is compatible, but the bare leakage theorem does not absorb the partial-overlap modes"}, {Name: "decoupling class classifier", Constructed: true, RequiresExtraChoice: true, Detail: "no scalar/Weyl/vector/constrained-mode decoupling rule is derived"}, {Name: "threshold beta permission", Constructed: true, CanCorrectBeta: false, Sealed: true, Detail: "no row is both physical-heavy, representation-complete, activation-derived, and decoupling-derived"}}
}
func buildAmbiguityWitnesses() []AmbiguityWitness {
	return []AmbiguityWitness{{Name: "contact modes as physical singlet scalars", CompatibleWithFiniteData: true, SelectedClass: OpenContactOverlapMode, WouldCorrectBeta: false, Reason: "singlet field reading gives zero gauge beta row, but locality and activation are not derived"}, {Name: "contact modes as scalar doublets", CompatibleWithFiniteData: true, SelectedClass: OpenContactOverlapMode, WouldCorrectBeta: true, Reason: "would alter beta coefficients, but the representation choice is not selected by finite data"}, {Name: "contact modes as regulator modes", CompatibleWithFiniteData: true, SelectedClass: OpenContactOverlapMode, WouldCorrectBeta: false, Reason: "possible but needs a regulator/ghost/signature theorem"}, {Name: "contact modes as constrained finite overlap modes", CompatibleWithFiniteData: true, SelectedClass: OpenContactOverlapMode, WouldCorrectBeta: false, Reason: "possible but no constraint-propagator theorem has been derived"}, {Name: "contact modes as vacuum-frustration modes", CompatibleWithFiniteData: true, SelectedClass: OpenContactOverlapMode, WouldCorrectBeta: false, Reason: "possible but the existing leakage invariant alone does not classify the partial modes"}}
}

func countRows(rows []ClassRow) rowCounts {
	var c rowCounts
	c.total = len(rows)
	for _, r := range rows {
		switch r.Class {
		case BaselineContinuumSector:
			c.baseline++
		case UnorientedScalarConstituent:
			c.unoriented++
		case BridgeScalarInvariant:
			c.bridge++
		case ConstrainedFiniteVacuumMode:
			c.constrained++
			if r.ModeKind == string(threshold.BGapCandidate) {
				c.bgap++
			}
		case OpenContactOverlapMode:
			if r.ModeKind == string(threshold.ContactOverlapCandidate) {
				c.contactOpen++
			}
		case VacuumFrustrationInvariant:
			c.vacuum++
		}
		if r.Status == ClassDerived || r.Status == ClassExcluded {
			c.derived++
		}
		if r.Status == ClassOpen {
			c.open++
		}
		if r.PhysicalFieldDerived && !r.ExcludedFromThresholdBeta {
			c.physical++
		}
		if r.RegulatorDerived {
			c.regulator++
		}
		if r.MayCorrectBeta {
			c.betaAllowed++
		}
	}
	return c
}
func bGapExcluded(rows []ClassRow) bool {
	for _, r := range rows {
		if r.ModeKind == string(threshold.BGapCandidate) {
			return r.Class == ConstrainedFiniteVacuumMode && r.ExcludedFromThresholdBeta
		}
	}
	return false
}
func FormatRows(rows []ClassRow, max int) string {
	if max <= 0 || max > len(rows) {
		max = len(rows)
	}
	parts := make([]string, 0, max+1)
	for i := 0; i < max; i++ {
		r := rows[i]
		parts = append(parts, fmt.Sprintf("%s:%s:%s", r.Name, r.Class, r.Status))
	}
	if max < len(rows) {
		parts = append(parts, fmt.Sprintf("... +%d more", len(rows)-max))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}
func FormatAttempts(attempts []ClassifierAttempt) string {
	parts := make([]string, 0, len(attempts))
	for _, a := range attempts {
		parts = append(parts, fmt.Sprintf("%s(select=%t,extra=%t,beta=%t)", a.Name, a.SelectsClass, a.RequiresExtraChoice, a.CanCorrectBeta))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}
func FormatWitnesses(ws []AmbiguityWitness) string {
	parts := make([]string, 0, len(ws))
	for _, w := range ws {
		parts = append(parts, fmt.Sprintf("%s(beta=%t)", w.Name, w.WouldCorrectBeta))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}
func Join(xs []string) string { return strings.Join(xs, "; ") }
