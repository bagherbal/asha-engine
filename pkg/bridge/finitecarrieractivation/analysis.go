// Package finitecarrieractivation implements Gate 205: finite carrier
// activation / contact-to-row semantics obstruction audit.
//
// Gate 204 proved a narrow but important positive result: the two
// non-universal Gate-201 shapes are exact rational representation-row lattice
// generators.  It also refused to promote the seven contact partial-overlap
// modes into those rows, because the contact carrier still lacks charge,
// spin-statistics, mass-activation, and decoupling semantics.
//
// Gate 205 turns that missing bridge into an explicit three-pillar obstruction.
// A finite contact mode may act as a heavy threshold beta row only after it has
// all of:
//
//  1. gauge charge semantics: SU(3)c, SU(2)L, U(1)Y labels and Dynkin indices;
//  2. spin-statistics semantics: Weyl/Dirac/scalar coefficient and local kinetic class;
//  3. mass-activation semantics: a VEV-independent scale/activation/decoupling rule.
//
// The audit intentionally does not assign contact modes to the Gate-201 shapes
// by multiplicity, numeric proximity, or desired RG outcome.  If any pillar is
// missing, the contact-to-row map remains a FAILED_ROUTE.
package finitecarrieractivation

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/representationrowlattice"
	"github.com/bagherbal/asha-engine/pkg/bridge/threshold"
)

type TargetShape struct {
	Name                  string
	TargetDeltaB          representationrowlattice.RationalTriple
	OnLattice             bool
	DirectGenerator       bool
	MatchedRepresentation string
	ConditionalSupport    bool
	FiniteDerived         bool
}

type Gate204Snapshot struct {
	Gate204Inherited                     bool
	Gate204ConditionalSupportPreserved   bool
	RepresentationLatticeConstructed     bool
	Gate201ShapesOnLattice               bool
	ContactMapFailed                     bool
	UniversalBetaSourceStillExternal     bool
	UniversalFitAvoided                  bool
	NoPhysicalPredictionClaim            bool
	ContactPartialOverlapModes           int
	ContactModesHaveChargeLabels         bool
	ContactModesHaveGaugeRepSemantics    bool
	ContactModesHaveDynkinIndices        bool
	ContactModesHaveSpinStatistics       bool
	ContactModesHaveMassActivation       bool
	ContactModesHaveDecouplingLaw        bool
	ContactModesPromotedToBetaRows       bool
	PhysicalUnificationClaimed           bool
	ThresholdCorrectedPhysicalFitClaimed bool
	AbsoluteMassPredicted                bool
	FiniteMatchingCorrectionsDerived     bool
	StrictNullityAfter                   int
	PhysicalPredictionNullityAfter       int
	TargetShapes                         []TargetShape
	TruthStatement                       string
}

func DefaultGate204Snapshot() (Gate204Snapshot, error) {
	prev, err := representationrowlattice.BuildDefault()
	if err != nil {
		return Gate204Snapshot{}, err
	}
	targets := make([]TargetShape, 0, len(prev.Memberships))
	for _, m := range prev.Memberships {
		targets = append(targets, TargetShape{
			Name:                  m.ShapeName,
			TargetDeltaB:          m.TargetDeltaB,
			OnLattice:             m.Found,
			DirectGenerator:       m.DirectGenerator,
			MatchedRepresentation: m.MatchedRepresentation,
			ConditionalSupport:    m.ConditionalSupport,
			FiniteDerived:         m.FiniteDerived,
		})
	}
	return Gate204Snapshot{
		Gate204Inherited:                     true,
		Gate204ConditionalSupportPreserved:   prev.Summary.ConditionalSupportLogged && prev.Summary.Gate201ShapesOnLattice,
		RepresentationLatticeConstructed:     prev.Summary.RationalGrammarConstructed && prev.Summary.LatticeConstructed,
		Gate201ShapesOnLattice:               prev.MembershipAudit.AllGate201ShapesSupported,
		ContactMapFailed:                     prev.Summary.ContactMapFailed && !prev.ContactInventory.CanonicalMapToRowBasisFound && !prev.ContactInventory.FiniteHeavySectorBasisDerived,
		UniversalBetaSourceStillExternal:     prev.Firewall.UniversalBetaSourceStillExternal,
		UniversalFitAvoided:                  prev.Summary.UniversalFitAvoided && !prev.Firewall.UniversalBetaFitAttempted,
		NoPhysicalPredictionClaim:            prev.Summary.NoPhysicalPredictionClaim,
		ContactPartialOverlapModes:           prev.ContactInventory.ContactPartialOverlapModes,
		ContactModesHaveChargeLabels:         prev.ContactInventory.ContactModesHaveChargeLabels,
		ContactModesHaveGaugeRepSemantics:    prev.ContactInventory.ContactModesHaveGaugeRepSemantics,
		ContactModesHaveDynkinIndices:        prev.ContactInventory.ContactModesHaveDynkinIndices,
		ContactModesHaveSpinStatistics:       prev.ContactInventory.ContactModesHaveSpinStatistics,
		ContactModesHaveMassActivation:       prev.ContactInventory.ContactModesHaveMassActivation,
		ContactModesHaveDecouplingLaw:        prev.ContactInventory.ContactModesHaveDecouplingLaw,
		ContactModesPromotedToBetaRows:       prev.Firewall.ContactModesPromotedToBetaRows,
		PhysicalUnificationClaimed:           prev.Firewall.PhysicalUnificationClaimed,
		ThresholdCorrectedPhysicalFitClaimed: prev.Firewall.ThresholdCorrectedPhysicalFitClaimed,
		AbsoluteMassPredicted:                prev.Firewall.AbsoluteMassPredicted,
		FiniteMatchingCorrectionsDerived:     prev.Firewall.FiniteMatchingCorrectionsDerived,
		StrictNullityAfter:                   prev.Firewall.StrictNullityAfter,
		PhysicalPredictionNullityAfter:       prev.Firewall.PhysicalPredictionNullityAfter,
		TargetShapes:                         targets,
		TruthStatement:                       prev.TruthStatement,
	}, nil
}

type ContactMode struct {
	Name                   string
	Value                  float64
	FiniteSpectralAnchor   bool
	PositivePartialOverlap bool
	ChargeSemantics        bool
	SpinStatistics         bool
	MassActivation         bool
	DecouplingLaw          bool
	AssignedTargetShape    string
	BetaRowAllowed         bool
	Verdict                string
}

type GaugeChargeAudit struct {
	ContactModesAudited           int
	TargetShapesAudited           int
	FiniteOverlapCarrierAvailable bool
	NativeSU3DynkinIndicesDerived bool
	NativeSU2DynkinIndicesDerived bool
	NativeHyperchargeDerived      bool
	CanonicalGaugeRepInheritance  bool
	CanFormDiracVectorlikeDoublet bool
	CanFormWeylSU2Adjoint         bool
	CandidateRowsAssigned         int
	GaugeChargeSemanticsComplete  bool
	Verdict                       string
}

type SpinStatisticsAudit struct {
	ContactModesAudited             int
	LocalContinuumFieldClassDerived bool
	LorentzKineticOperatorDerived   bool
	WeylCoefficientDerived          bool
	DiracCoefficientDerived         bool
	ScalarCoefficientDerived        bool
	SpinStatisticsAssigned          bool
	StandardBetaCoefficientSelected bool
	Verdict                         string
}

type MassActivationAudit struct {
	ContactModesAudited                  int
	DimensionlessSpectralValuesAvailable bool
	CanonicalPhysicalMassUnitDerived     bool
	VEVIndependentActivationDerived      bool
	DecouplingScaleDerived               bool
	ActivationPredicateDerived           bool
	MatchingSchemeDerived                bool
	ThresholdCorrectedBetaRowsAllowed    bool
	Verdict                              string
}

type CarrierActivationClassification struct {
	RequiredPillars               int
	CompletePillars               int
	MissingPillars                []string
	CarrierActivationDerived      bool
	ContactModesCanBeHeavyRows    bool
	ContactModesCanBeTargetShapes bool
	Classification                string
	Verdict                       string
}

type FirewallAudit struct {
	Gate204Inherited                     bool
	Gate204ConditionalSupportPreserved   bool
	RepresentationLatticeConstructed     bool
	Gate201ShapesRemainConditional       bool
	ContactModesPromotedToBetaRows       bool
	ContactModesAssignedToGate201Shapes  bool
	ArbitraryChargeAssignmentInserted    bool
	ArbitrarySpinStatisticInserted       bool
	ArbitraryMassScaleInserted           bool
	PhenomenologicalVEVUsedForActivation bool
	UniversalBetaFitAttempted            bool
	ContinuousScalesSolved               bool
	PhysicalUnificationClaimed           bool
	ThresholdCorrectedPhysicalFitClaimed bool
	AbsoluteMassPredicted                bool
	FiniteMatchingCorrectionsDerived     bool
	StrictNullityBefore                  int
	StrictNullityAfter                   int
	PhysicalPredictionNullityBefore      int
	PhysicalPredictionNullityAfter       int
	RecommendedNextGate                  string
	OpenRequirements                     []string
	Verdict                              string
}

type Summary struct {
	TestsAudited                   int
	Gate204Inherited               bool
	ContactModesAudited            bool
	GaugeChargeObstructed          bool
	SpinStatisticsObstructed       bool
	MassActivationObstructed       bool
	CarrierActivationObstructed    bool
	Gate201ShapesRemainConditional bool
	FailedRouteLogged              bool
	NoPhysicalPredictionClaim      bool
	Status                         string
	Comment                        string
}

type Analysis struct {
	PreviousGate204 Gate204Snapshot
	ContactModes    []ContactMode
	GaugeCharge     GaugeChargeAudit
	SpinStatistics  SpinStatisticsAudit
	MassActivation  MassActivationAudit
	Classification  CarrierActivationClassification
	Firewall        FirewallAudit
	Summary         Summary
	TruthStatement  string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		prev, err := DefaultGate204Snapshot()
		if err != nil {
			defaultErr = err
			return
		}
		th, err := threshold.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultA, defaultErr = Build(prev, th.ContactPartialOverlap)
	})
	return defaultA, defaultErr
}

func Build(prev Gate204Snapshot, contactValues []float64) (Analysis, error) {
	if !prev.Gate204Inherited || !prev.Gate204ConditionalSupportPreserved || !prev.RepresentationLatticeConstructed || !prev.Gate201ShapesOnLattice {
		return Analysis{}, fmt.Errorf("Gate 205 requires Gate 204 conditional-support lattice membership")
	}
	if !prev.ContactMapFailed || prev.ContactModesPromotedToBetaRows || prev.PhysicalUnificationClaimed || prev.ThresholdCorrectedPhysicalFitClaimed || prev.AbsoluteMassPredicted || prev.FiniteMatchingCorrectionsDerived {
		return Analysis{}, fmt.Errorf("Gate 205 refuses inherited contact-map or physical-prediction leakage")
	}
	if len(contactValues) != 7 || prev.ContactPartialOverlapModes != 7 {
		return Analysis{}, fmt.Errorf("Gate 205 requires exactly seven contact partial-overlap modes, got snapshot=%d values=%d", prev.ContactPartialOverlapModes, len(contactValues))
	}

	modes := buildContactModes(contactValues)
	gauge := auditGaugeCharge(prev, modes)
	spin := auditSpinStatistics(modes)
	mass := auditMassActivation(modes)
	class := classifyCarrierActivation(gauge, spin, mass)
	fw := auditFirewall(prev, gauge, spin, mass, class)
	summary := Summary{
		TestsAudited:                   6,
		Gate204Inherited:               fw.Gate204Inherited && fw.Gate204ConditionalSupportPreserved,
		ContactModesAudited:            len(modes) == 7,
		GaugeChargeObstructed:          !gauge.GaugeChargeSemanticsComplete && gauge.CandidateRowsAssigned == 0,
		SpinStatisticsObstructed:       !spin.SpinStatisticsAssigned && !spin.StandardBetaCoefficientSelected,
		MassActivationObstructed:       !mass.VEVIndependentActivationDerived && !mass.ThresholdCorrectedBetaRowsAllowed,
		CarrierActivationObstructed:    !class.CarrierActivationDerived && !class.ContactModesCanBeHeavyRows,
		Gate201ShapesRemainConditional: fw.Gate201ShapesRemainConditional,
		FailedRouteLogged:              true,
		NoPhysicalPredictionClaim:      !fw.PhysicalUnificationClaimed && !fw.ThresholdCorrectedPhysicalFitClaimed && !fw.AbsoluteMassPredicted && fw.PhysicalPredictionNullityBefore == fw.PhysicalPredictionNullityAfter,
		Status:                         "FAILED_ROUTE",
		Comment:                        "Gate 205 audits the exact semantic requirements needed to promote seven contact partial-overlap modes into heavy beta rows. The finite carrier is real and positive, and the Gate-201 target shapes are legal row-lattice generators, but the contact modes still lack gauge charge/Dynkin labels, spin-statistics coefficients, and a VEV-independent activation/decoupling law. Carrier activation is therefore an obstruction, not a prediction.",
	}
	truth := "Gate 205 consolidates the contact-to-row semantic gap into a three-pillar carrier-activation obstruction. The seven contact partial-overlap modes are finite positive spectral anchors, and Gate 204 proved the target rows (3,2,1/6) Dirac and (1,3,0) Weyl are legal rational lattice generators. However, no theorem assigns the contact modes SU(3)c/SU(2)L/U(1)Y Dynkin data, Weyl/Dirac/scalar statistics, or a VEV-independent mass activation and decoupling law. Therefore zero contact modes are promoted to heavy beta rows and the Gate-201 shapes remain conditional representation support only."

	return Analysis{PreviousGate204: prev, ContactModes: modes, GaugeCharge: gauge, SpinStatistics: spin, MassActivation: mass, Classification: class, Firewall: fw, Summary: summary, TruthStatement: truth}, nil
}

func buildContactModes(values []float64) []ContactMode {
	vals := append([]float64(nil), values...)
	sort.Sort(sort.Reverse(sort.Float64Slice(vals)))
	out := make([]ContactMode, 0, len(vals))
	for i, v := range vals {
		out = append(out, ContactMode{
			Name:                   fmt.Sprintf("contact partial-overlap mode %d", i+1),
			Value:                  v,
			FiniteSpectralAnchor:   true,
			PositivePartialOverlap: v > 0,
			ChargeSemantics:        false,
			SpinStatistics:         false,
			MassActivation:         false,
			DecouplingLaw:          false,
			AssignedTargetShape:    "",
			BetaRowAllowed:         false,
			Verdict:                "finite positive overlap mode; carrier activation semantics absent",
		})
	}
	return out
}

func auditGaugeCharge(prev Gate204Snapshot, modes []ContactMode) GaugeChargeAudit {
	assigned := 0
	for _, m := range modes {
		if m.ChargeSemantics || m.AssignedTargetShape != "" || m.BetaRowAllowed {
			assigned++
		}
	}
	complete := false
	return GaugeChargeAudit{
		ContactModesAudited:           len(modes),
		TargetShapesAudited:           len(prev.TargetShapes),
		FiniteOverlapCarrierAvailable: allFinitePositive(modes),
		NativeSU3DynkinIndicesDerived: false,
		NativeSU2DynkinIndicesDerived: false,
		NativeHyperchargeDerived:      false,
		CanonicalGaugeRepInheritance:  false,
		CanFormDiracVectorlikeDoublet: false,
		CanFormWeylSU2Adjoint:         false,
		CandidateRowsAssigned:         assigned,
		GaugeChargeSemanticsComplete:  complete,
		Verdict:                       "FAILED_ROUTE: contact modes have finite overlap values but no canonical SU(3)c×SU(2)L×U(1)Y charge/Dynkin labels",
	}
}

func auditSpinStatistics(modes []ContactMode) SpinStatisticsAudit {
	return SpinStatisticsAudit{
		ContactModesAudited:             len(modes),
		LocalContinuumFieldClassDerived: false,
		LorentzKineticOperatorDerived:   false,
		WeylCoefficientDerived:          false,
		DiracCoefficientDerived:         false,
		ScalarCoefficientDerived:        false,
		SpinStatisticsAssigned:          false,
		StandardBetaCoefficientSelected: false,
		Verdict:                         "FAILED_ROUTE: no local Lorentz kinetic class selects Weyl, Dirac, complex-scalar, or real-scalar beta coefficient",
	}
}

func auditMassActivation(modes []ContactMode) MassActivationAudit {
	return MassActivationAudit{
		ContactModesAudited:                  len(modes),
		DimensionlessSpectralValuesAvailable: allFinitePositive(modes),
		CanonicalPhysicalMassUnitDerived:     false,
		VEVIndependentActivationDerived:      false,
		DecouplingScaleDerived:               false,
		ActivationPredicateDerived:           false,
		MatchingSchemeDerived:                false,
		ThresholdCorrectedBetaRowsAllowed:    false,
		Verdict:                              "FAILED_ROUTE: contact values are dimensionless anchors only; no VEV-independent mass unit, activation predicate, or decoupling law is derived",
	}
}

func classifyCarrierActivation(g GaugeChargeAudit, s SpinStatisticsAudit, m MassActivationAudit) CarrierActivationClassification {
	missing := make([]string, 0, 3)
	complete := 0
	if g.GaugeChargeSemanticsComplete {
		complete++
	} else {
		missing = append(missing, "gauge charge / Dynkin / hypercharge semantics")
	}
	if s.SpinStatisticsAssigned && s.StandardBetaCoefficientSelected {
		complete++
	} else {
		missing = append(missing, "spin-statistics / local kinetic coefficient semantics")
	}
	if m.VEVIndependentActivationDerived && m.DecouplingScaleDerived && m.ActivationPredicateDerived && m.MatchingSchemeDerived {
		complete++
	} else {
		missing = append(missing, "mass activation / decoupling / threshold matching semantics")
	}
	derived := complete == 3
	return CarrierActivationClassification{
		RequiredPillars:               3,
		CompletePillars:               complete,
		MissingPillars:                missing,
		CarrierActivationDerived:      derived,
		ContactModesCanBeHeavyRows:    derived,
		ContactModesCanBeTargetShapes: derived,
		Classification:                "Carrier Activation is a bridge obstruction under current axioms; it may require a future spontaneous semantic seal or a new finite local-field theorem.",
		Verdict:                       "FAILED_ROUTE",
	}
}

func auditFirewall(prev Gate204Snapshot, g GaugeChargeAudit, s SpinStatisticsAudit, m MassActivationAudit, c CarrierActivationClassification) FirewallAudit {
	return FirewallAudit{
		Gate204Inherited:                     prev.Gate204Inherited,
		Gate204ConditionalSupportPreserved:   prev.Gate204ConditionalSupportPreserved,
		RepresentationLatticeConstructed:     prev.RepresentationLatticeConstructed,
		Gate201ShapesRemainConditional:       allTargetsConditional(prev.TargetShapes),
		ContactModesPromotedToBetaRows:       false,
		ContactModesAssignedToGate201Shapes:  g.CandidateRowsAssigned > 0,
		ArbitraryChargeAssignmentInserted:    false,
		ArbitrarySpinStatisticInserted:       false,
		ArbitraryMassScaleInserted:           false,
		PhenomenologicalVEVUsedForActivation: false,
		UniversalBetaFitAttempted:            false,
		ContinuousScalesSolved:               false,
		PhysicalUnificationClaimed:           false,
		ThresholdCorrectedPhysicalFitClaimed: false,
		AbsoluteMassPredicted:                false,
		FiniteMatchingCorrectionsDerived:     false,
		StrictNullityBefore:                  prev.StrictNullityAfter,
		StrictNullityAfter:                   prev.StrictNullityAfter,
		PhysicalPredictionNullityBefore:      prev.PhysicalPredictionNullityAfter,
		PhysicalPredictionNullityAfter:       prev.PhysicalPredictionNullityAfter,
		RecommendedNextGate:                  "Gate 206 — carrier-activation seal / local-field semantic bifurcation audit",
		OpenRequirements: []string{
			"derive a canonical gauge-representation functor from contact modes to SU(3)c×SU(2)L×U(1)Y rows",
			"derive a local Lorentz kinetic class selecting Weyl, Dirac, complex-scalar, or real-scalar coefficients",
			"derive a VEV-independent activation scale and decoupling predicate, or seal carrier activation as an empirical/spontaneous boundary datum",
		},
		Verdict: formatFirewallVerdict(g, s, m, c),
	}
}

func allFinitePositive(modes []ContactMode) bool {
	if len(modes) == 0 {
		return false
	}
	for _, m := range modes {
		if !m.FiniteSpectralAnchor || !m.PositivePartialOverlap || !(m.Value > 0) {
			return false
		}
	}
	return true
}

func allTargetsConditional(targets []TargetShape) bool {
	if len(targets) == 0 {
		return false
	}
	for _, t := range targets {
		if !t.OnLattice || !t.DirectGenerator || !t.ConditionalSupport || t.FiniteDerived {
			return false
		}
	}
	return true
}

func formatFirewallVerdict(g GaugeChargeAudit, s SpinStatisticsAudit, m MassActivationAudit, c CarrierActivationClassification) string {
	return fmt.Sprintf("%s; chargeComplete=%t spinComplete=%t activationComplete=%t missing=%s", c.Verdict, g.GaugeChargeSemanticsComplete, s.SpinStatisticsAssigned, m.VEVIndependentActivationDerived, strings.Join(c.MissingPillars, "; "))
}

func FormatGate204(g Gate204Snapshot) string {
	return fmt.Sprintf("inherited=%t conditionalSupport=%t lattice=%t shapesOnLattice=%t contactMapFailed=%t universalExternal=%t noPrediction=%t contactModes=%d targets=%s", g.Gate204Inherited, g.Gate204ConditionalSupportPreserved, g.RepresentationLatticeConstructed, g.Gate201ShapesOnLattice, g.ContactMapFailed, g.UniversalBetaSourceStillExternal, g.NoPhysicalPredictionClaim, g.ContactPartialOverlapModes, FormatTargets(g.TargetShapes))
}

func FormatTargets(targets []TargetShape) string {
	parts := make([]string, 0, len(targets))
	for _, t := range targets {
		parts = append(parts, fmt.Sprintf("%s:%s via %s conditional=%t finiteDerived=%t", t.Name, t.TargetDeltaB, t.MatchedRepresentation, t.ConditionalSupport, t.FiniteDerived))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatContactModes(modes []ContactMode, max int) string {
	if max <= 0 || max > len(modes) {
		max = len(modes)
	}
	parts := make([]string, 0, max)
	for i := 0; i < max; i++ {
		m := modes[i]
		parts = append(parts, fmt.Sprintf("%s=%.10f betaAllowed=%t verdict=%s", m.Name, m.Value, m.BetaRowAllowed, m.Verdict))
	}
	if max < len(modes) {
		parts = append(parts, fmt.Sprintf("... +%d more", len(modes)-max))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatGaugeCharge(g GaugeChargeAudit) string {
	return fmt.Sprintf("modes=%d targets=%d carrier=%t su3Dynkin=%t su2Dynkin=%t hypercharge=%t inheritance=%t diracQ=%t weylAdj=%t rowsAssigned=%d complete=%t verdict=%s", g.ContactModesAudited, g.TargetShapesAudited, g.FiniteOverlapCarrierAvailable, g.NativeSU3DynkinIndicesDerived, g.NativeSU2DynkinIndicesDerived, g.NativeHyperchargeDerived, g.CanonicalGaugeRepInheritance, g.CanFormDiracVectorlikeDoublet, g.CanFormWeylSU2Adjoint, g.CandidateRowsAssigned, g.GaugeChargeSemanticsComplete, g.Verdict)
}

func FormatSpinStatistics(s SpinStatisticsAudit) string {
	return fmt.Sprintf("modes=%d localField=%t lorentzKinetic=%t weyl=%t dirac=%t scalar=%t assigned=%t betaCoeff=%t verdict=%s", s.ContactModesAudited, s.LocalContinuumFieldClassDerived, s.LorentzKineticOperatorDerived, s.WeylCoefficientDerived, s.DiracCoefficientDerived, s.ScalarCoefficientDerived, s.SpinStatisticsAssigned, s.StandardBetaCoefficientSelected, s.Verdict)
}

func FormatMassActivation(m MassActivationAudit) string {
	return fmt.Sprintf("modes=%d dimensionlessValues=%t massUnit=%t vevIndependent=%t decouplingScale=%t activation=%t matching=%t betaRowsAllowed=%t verdict=%s", m.ContactModesAudited, m.DimensionlessSpectralValuesAvailable, m.CanonicalPhysicalMassUnitDerived, m.VEVIndependentActivationDerived, m.DecouplingScaleDerived, m.ActivationPredicateDerived, m.MatchingSchemeDerived, m.ThresholdCorrectedBetaRowsAllowed, m.Verdict)
}

func FormatClassification(c CarrierActivationClassification) string {
	return fmt.Sprintf("pillars=%d complete=%d missing=%s derived=%t heavyRows=%t targetShapes=%t class=%s verdict=%s", c.RequiredPillars, c.CompletePillars, strings.Join(c.MissingPillars, "; "), c.CarrierActivationDerived, c.ContactModesCanBeHeavyRows, c.ContactModesCanBeTargetShapes, c.Classification, c.Verdict)
}

func FormatFirewall(f FirewallAudit) string {
	return fmt.Sprintf("gate204=%t conditional=%t lattice=%t shapesConditional=%t contactBeta=%t contactAssigned=%t arbitraryCharge=%t arbitrarySpin=%t arbitraryMass=%t vevUsed=%t universalFit=%t scalesSolved=%t unification=%t thresholdFit=%t mass=%t matching=%t strictNullity=%d->%d physicalNullity=%d->%d next=%s", f.Gate204Inherited, f.Gate204ConditionalSupportPreserved, f.RepresentationLatticeConstructed, f.Gate201ShapesRemainConditional, f.ContactModesPromotedToBetaRows, f.ContactModesAssignedToGate201Shapes, f.ArbitraryChargeAssignmentInserted, f.ArbitrarySpinStatisticInserted, f.ArbitraryMassScaleInserted, f.PhenomenologicalVEVUsedForActivation, f.UniversalBetaFitAttempted, f.ContinuousScalesSolved, f.PhysicalUnificationClaimed, f.ThresholdCorrectedPhysicalFitClaimed, f.AbsoluteMassPredicted, f.FiniteMatchingCorrectionsDerived, f.StrictNullityBefore, f.StrictNullityAfter, f.PhysicalPredictionNullityBefore, f.PhysicalPredictionNullityAfter, f.RecommendedNextGate)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("tests=%d gate204=%t contacts=%t chargeBlocked=%t spinBlocked=%t activationBlocked=%t carrierBlocked=%t shapesConditional=%t failed=%t noPrediction=%t status=%s comment=%s", s.TestsAudited, s.Gate204Inherited, s.ContactModesAudited, s.GaugeChargeObstructed, s.SpinStatisticsObstructed, s.MassActivationObstructed, s.CarrierActivationObstructed, s.Gate201ShapesRemainConditional, s.FailedRouteLogged, s.NoPhysicalPredictionClaim, s.Status, s.Comment)
}
