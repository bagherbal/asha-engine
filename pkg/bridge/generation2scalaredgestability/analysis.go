// Package generation2scalaredgestability implements Gate 491:
// Scalar-Edge Stability and Higgs One-Form Positivity Audit.
//
// Gate 490 redirected ASHA away from closed flavor/Yukawa prediction and into
// mass-independent topological and variational invariants. Gate 491 audits the
// scalar/Higgs edge lane: whether the already-derived Higgs one-form edge
// support and scalar kinetic trace positivity are enough to prove a native
// stability theorem independent of flavor moduli.
//
// The answer is deliberately bounded. The edge support is native as a finite
// one-form support theorem, and the kinetic carrier is positive-semidefinite as
// a Hilbert-Schmidt sum of edge squares. A count-level Goldstone resonance is
// present. But the full scalar Hessian/vacuum-stability theorem, numerical Z_H,
// Higgs quartic, pole mass, W/Z mass matrix, and gauge-eating map remain blocked
// by the amplitude, cutoff, sign, subtraction, and covariant-derivative seals.
package generation2scalaredgestability

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/goldstone"
	"github.com/bagherbal/asha-engine/pkg/bridge/innerfluctuationedgemeasure"
	"github.com/bagherbal/asha-engine/pkg/bridge/scalarkinetictracepositivity"
)

const (
	AuditID = "GATE491-SCALAR-EDGE-STABILITY-HIGGS-ONEFORM-POSITIVITY-AUDIT"

	StatusGate490Inherited                 = "CONDITIONAL_SUPPORT_GATE490_TOPOLOGICAL_REDIRECT_INHERITED"
	StatusHiggsOneFormEdgeSupportInherited = "CONDITIONAL_SUPPORT_HIGGS_ONEFORM_EDGE_SUPPORT_INHERITED"
	StatusScalarKineticTracePositive       = "CONDITIONAL_SUPPORT_SCALAR_KINETIC_TRACE_POSITIVE_SEMIDEFINITE"
	StatusGhostSignObstructionRemoved      = "CONDITIONAL_SUPPORT_NEGATIVE_GHOST_KINETIC_ROUTE_BLOCKED"
	StatusStrictPositivityConditional      = "CONDITIONAL_SUPPORT_STRICT_ZH_POSITIVITY_CONDITION_IDENTIFIED"
	StatusGoldstoneCountResonance          = "CONDITIONAL_SUPPORT_GOLDSTONE_COUNT_RESONANCE_CONFIRMED"
	StatusFlavorIndependencePreserved      = "CONDITIONAL_SUPPORT_SCALAR_EDGE_AUDIT_FLAVOR_INDEPENDENT"
	StatusFirewallPreserved                = "FIREWALL_PRESERVED_NO_MASS_OR_FLAVOR_DATA_IMPORTED"

	StatusFailedNumericalZHSealed           = "FAILED_ROUTE_NUMERICAL_ZH_VALUE_STILL_SEALED"
	StatusFailedFullScalarHessianNotDerived = "FAILED_ROUTE_FULL_SCALAR_EDGE_HESSIAN_NOT_DERIVED"
	StatusFailedVacuumStabilityNotDerived   = "FAILED_ROUTE_ABSOLUTE_VACUUM_STABILITY_NOT_DERIVED"
	StatusFailedQuarticMassNotDerived       = "FAILED_ROUTE_HIGGS_QUARTIC_AND_MASS_NOT_DERIVED"
	StatusFailedGaugeEatingMapNotDerived    = "FAILED_ROUTE_CANONICAL_GOLDSTONE_GAUGE_EATING_MAP_NOT_DERIVED"
	StatusFailedContinuumPermissionNotFull  = "FAILED_ROUTE_CONTINUUM_SCALAR_MATCHING_PERMISSION_NOT_COMPLETE"

	StatusGate492RedirectDefined = "CONDITIONAL_SUPPORT_GATE492_SCALAR_COVARIANT_DERIVATIVE_REDIRECT_DEFINED"
)

const (
	NativeFlavorDim = 13
	KXYCoeffDim     = 9
)

type Inheritance struct {
	Executed                     bool
	Gate489FlavorAirlockClosed   bool
	Gate490AnomalyLedgerStable   bool
	NonFlavorFrontierSelected    bool
	YukawaEntriesEnvironmental   bool
	CKMOrientationEnvironmental  bool
	NoObservedFlavorDataImported bool
	Verdict                      string
	Reason                       string
}

type OneFormSupport struct {
	Executed                bool
	HiggsIsFiniteOneForm    bool
	EdgeMeasureSelected     bool
	NodeMeasureAdmissible   bool
	JDoubledEdgeCount       int
	PhysicalPoleMassDerived bool
	FullNumericalTOEClosed  bool
	SupportFormula          string
	Verdict                 string
	Reason                  string
}

type KineticPositivity struct {
	Executed                        bool
	TraceFunctionalFormalized       bool
	DoubledCarrierEvaluated         bool
	EdgeTerms                       int
	QuarkEdges                      int
	LeptonEdges                     int
	UsesHilbertSchmidtSquares       bool
	PositiveSemidefinite            bool
	NegativeTermsPermitted          bool
	ImaginaryKineticPermitted       bool
	GhostRiskEliminated             bool
	StrictPositiveConditional       bool
	StrictPositiveProvedNumerically bool
	NumericalZHComputed             bool
	YukawaAmplitudesSealed          bool
	CutoffMomentSealed              bool
	SignConventionSealed            bool
	Verdict                         string
	Reason                          string
}

type GoldstoneAudit struct {
	Executed                             bool
	ActiveRealDirections                 int
	RadialDirections                     int
	ScalarAngularDirections              int
	ProtectedContactDirections           int
	BrokenElectroweakDirections          int
	CountResonance                       bool
	CanonicalProtectedToBrokenMapDerived bool
	SU2ActionOnContactScalarDerived      bool
	CovariantDerivativeDerived           bool
	GaugeBosonMassMatrixDerived          bool
	GaugeEatingTheoremDerived            bool
	Verdict                              string
	Reason                               string
}

type StabilityBoundary struct {
	Executed                        bool
	EdgeSupportNative               bool
	KineticSemidefiniteNative       bool
	GhostInstabilityBlocked         bool
	FullHessianDerived              bool
	VacuumStabilityDerived          bool
	QuarticDerived                  bool
	HiggsMassDerived                bool
	ContinuumScalarMatchingComplete bool
	Verdict                         string
	Reason                          string
}

type Firewall struct {
	Executed                  bool
	ObservedMassesImported    bool
	ObservedYukawaImported    bool
	ObservedCKMImported       bool
	ObservedPMNSImported      bool
	ObservedHiggsMassImported bool
	NativeYukawaMatrixWritten bool
	NativeCKMMatrixWritten    bool
	NativeQuarticMassWritten  bool
	NativeFlavorModuliChanged bool
	NativeFlavorDimAfter      int
	KXYCoeffDimAfter          int
	Verdict                   string
	Reason                    string
}

type RegistryUpdate struct {
	NativeEntries        []string
	BridgeEntries        []string
	EnvironmentalEntries []string
	FailedRoutes         []string
	OpenTheorems         []string
}

type NextStep struct {
	Gate                       int
	Title, Reason, PrimaryTask string
}

type Analysis struct {
	Inheritance Inheritance
	Support     OneFormSupport
	Kinetic     KineticPositivity
	Goldstone   GoldstoneAudit
	Boundary    StabilityBoundary
	Firewall    Firewall
	Registry    RegistryUpdate
	Next        NextStep
	Truth       string
}

var cache struct {
	sync.Once
	a   Analysis
	err error
}

func BuildDefault() (Analysis, error) {
	cache.Once.Do(func() { cache.a, cache.err = Build() })
	return cache.a, cache.err
}

func Build() (Analysis, error) {
	supportSrc, err := innerfluctuationedgemeasure.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate385 edge measure: %w", err)
	}
	kineticSrc, err := scalarkinetictracepositivity.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate301 scalar kinetic positivity: %w", err)
	}
	goldstoneSrc, err := goldstone.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Goldstone audit: %w", err)
	}

	a := Analysis{Inheritance: buildInheritance()}
	a.Support = buildOneFormSupport(supportSrc)
	a.Kinetic = buildKineticPositivity(kineticSrc)
	a.Goldstone = buildGoldstoneAudit(goldstoneSrc)
	a.Boundary = buildStabilityBoundary(a.Support, a.Kinetic, a.Goldstone)
	a.Firewall = buildFirewall()
	a.Registry = buildRegistryUpdate(a)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance() Inheritance {
	return Inheritance{
		Executed:                     true,
		Gate489FlavorAirlockClosed:   true,
		Gate490AnomalyLedgerStable:   true,
		NonFlavorFrontierSelected:    true,
		YukawaEntriesEnvironmental:   true,
		CKMOrientationEnvironmental:  true,
		NoObservedFlavorDataImported: true,
		Verdict:                      StatusGate490Inherited,
		Reason:                       "Gate490 proved a mass-independent anomaly ledger after Gate489 closed native flavor prediction; Gate491 therefore audits scalar-edge stability without importing masses or flavor moduli.",
	}
}

func buildOneFormSupport(src innerfluctuationedgemeasure.Analysis) OneFormSupport {
	c := src.Calculation
	return OneFormSupport{
		Executed:                c.Executed,
		HiggsIsFiniteOneForm:    c.InnerFluctuation.IsFiniteOneForm,
		EdgeMeasureSelected:     c.EdgeMeasureSelected && c.Support.EdgeMeasureMandated,
		NodeMeasureAdmissible:   c.Support.NodeMeasureAdmissible,
		JDoubledEdgeCount:       int(c.Support.EdgeCount),
		PhysicalPoleMassDerived: c.PhysicalPoleMassDerived,
		FullNumericalTOEClosed:  c.FullNumericalTOEClosed,
		SupportFormula:          c.Support.SupportProjectionFormula,
		Verdict:                 StatusHiggsOneFormEdgeSupportInherited,
		Reason:                  "the Higgs carrier is inherited as a finite one-form supported on the J-doubled finite Dirac edge module, not as a scalar placed on contact nodes",
	}
}

func buildKineticPositivity(src scalarkinetictracepositivity.Analysis) KineticPositivity {
	return KineticPositivity{
		Executed:                        true,
		TraceFunctionalFormalized:       src.Summary.TraceFunctionalFormalized,
		DoubledCarrierEvaluated:         src.Summary.DoubledCarrierEvaluated,
		EdgeTerms:                       src.Doubled.TotalEdgesMapped,
		QuarkEdges:                      src.Doubled.QuarkEdgesMapped,
		LeptonEdges:                     src.Doubled.LeptonEdgesMapped,
		UsesHilbertSchmidtSquares:       src.Trace.UsesHilbertSchmidtNorm,
		PositiveSemidefinite:            src.Summary.PositiveSemidefiniteProved,
		NegativeTermsPermitted:          src.Positivity.NegativeTermsPermitted,
		ImaginaryKineticPermitted:       src.Positivity.ImaginaryKineticPermitted,
		GhostRiskEliminated:             src.Positivity.GhostRiskEliminatedStructurally,
		StrictPositiveConditional:       src.Summary.StrictPositiveConditionIdentified,
		StrictPositiveProvedNumerically: src.Positivity.StrictPositiveProved,
		NumericalZHComputed:             src.Summary.NumericalZHComputed,
		YukawaAmplitudesSealed:          src.Seals.AllNumericalValuesSealed,
		CutoffMomentSealed:              src.ZH.RequiresPositiveF0,
		SignConventionSealed:            src.ZH.RequiresEuclideanSignLedger,
		Verdict:                         strings.Join([]string{StatusScalarKineticTracePositive, StatusGhostSignObstructionRemoved, StatusStrictPositivityConditional, StatusFailedNumericalZHSealed}, "; "),
		Reason:                          "the scalar kinetic carrier is a positive-semidefinite Hilbert-Schmidt edge-square trace; strict numerical Z_H still requires a nonzero amplitude theorem/seal plus f0, sign, and trace conventions",
	}
}

func buildGoldstoneAudit(src goldstone.Analysis) GoldstoneAudit {
	return GoldstoneAudit{
		Executed:                             true,
		ActiveRealDirections:                 src.ActiveRealDirections,
		RadialDirections:                     src.RadialDirections,
		ScalarAngularDirections:              src.ScalarAngularDirections,
		ProtectedContactDirections:           src.ProtectedContactDirections,
		BrokenElectroweakDirections:          src.BrokenGaugeDirections,
		CountResonance:                       src.GoldstoneCountResonance,
		CanonicalProtectedToBrokenMapDerived: src.CanonicalProtectedToBrokenMapDerived,
		SU2ActionOnContactScalarDerived:      src.SU2LActionOnContactScalarDerived,
		CovariantDerivativeDerived:           src.CovariantDerivativeDerived,
		GaugeBosonMassMatrixDerived:          src.GaugeBosonMassMatrixDerived,
		GaugeEatingTheoremDerived:            src.GaugeEatingTheoremDerived,
		Verdict:                              strings.Join([]string{StatusGoldstoneCountResonance, StatusFailedGaugeEatingMapNotDerived}, "; "),
		Reason:                               "four real scalar directions split as one radial plus three angular directions, matching three protected contact directions and three broken electroweak directions, but no canonical gauge-eating map or covariant derivative is derived",
	}
}

func buildStabilityBoundary(s OneFormSupport, k KineticPositivity, g GoldstoneAudit) StabilityBoundary {
	return StabilityBoundary{
		Executed:                        true,
		EdgeSupportNative:               s.HiggsIsFiniteOneForm && s.EdgeMeasureSelected && !s.NodeMeasureAdmissible,
		KineticSemidefiniteNative:       k.PositiveSemidefinite && !k.NegativeTermsPermitted && !k.ImaginaryKineticPermitted,
		GhostInstabilityBlocked:         k.GhostRiskEliminated,
		FullHessianDerived:              false,
		VacuumStabilityDerived:          false,
		QuarticDerived:                  false,
		HiggsMassDerived:                false,
		ContinuumScalarMatchingComplete: false,
		Verdict:                         strings.Join([]string{StatusScalarKineticTracePositive, StatusFailedFullScalarHessianNotDerived, StatusFailedVacuumStabilityNotDerived, StatusFailedQuarticMassNotDerived, StatusFailedContinuumPermissionNotFull}, "; "),
		Reason:                          fmt.Sprintf("edge support and kinetic positivity pass, and Goldstone count resonance=%t; however a full scalar-edge Hessian needs quartic/potential coefficients, subtraction/sign conventions, covariant derivative, and continuum matching", g.CountResonance),
	}
}

func buildFirewall() Firewall {
	return Firewall{
		Executed:                  true,
		ObservedMassesImported:    false,
		ObservedYukawaImported:    false,
		ObservedCKMImported:       false,
		ObservedPMNSImported:      false,
		ObservedHiggsMassImported: false,
		NativeYukawaMatrixWritten: false,
		NativeCKMMatrixWritten:    false,
		NativeQuarticMassWritten:  false,
		NativeFlavorModuliChanged: false,
		NativeFlavorDimAfter:      NativeFlavorDim,
		KXYCoeffDimAfter:          KXYCoeffDim,
		Verdict:                   StatusFirewallPreserved,
		Reason:                    "Gate491 imports no masses, Yukawa matrices, CKM/PMNS data, observed Higgs mass, or pole-matching constants and writes no native flavor or mass prediction.",
	}
}

func buildRegistryUpdate(_ Analysis) RegistryUpdate {
	return RegistryUpdate{
		NativeEntries: []string{
			"Higgs/scalar support is inherited as a finite one-form edge module rather than a contact-node scalar measure",
			"the scalar kinetic trace carrier is positive-semidefinite because allowed finite Dirac scalar edges enter as Hilbert-Schmidt squares",
			"negative or imaginary scalar kinetic ghost terms are structurally blocked at the finite edge-trace level",
		},
		BridgeEntries: []string{
			"strict Z_H>0 is conditional on a positive convention ledger and at least one nonzero scalar edge amplitude theorem/seal",
			"the 1+3 scalar split has a count-level Goldstone resonance with three broken electroweak directions",
		},
		EnvironmentalEntries: []string{
			"Yukawa amplitudes, cutoff moment f0, observed Higgs mass, quartic value, pole scheme, CKM, PMNS, and continuum threshold data remain sealed bridge/environmental data",
		},
		FailedRoutes: []string{
			StatusFailedNumericalZHSealed,
			StatusFailedFullScalarHessianNotDerived,
			StatusFailedVacuumStabilityNotDerived,
			StatusFailedQuarticMassNotDerived,
			StatusFailedGaugeEatingMapNotDerived,
			StatusFailedContinuumPermissionNotFull,
		},
		OpenTheorems: []string{
			StatusGate492RedirectDefined,
			"derive a native scalar covariant derivative and protected-to-broken intertwiner before any W/Z or gauge-eating mass theorem",
		},
	}
}

func buildNext() NextStep {
	return NextStep{
		Gate:        492,
		Title:       "Scalar Covariant Derivative and Goldstone Intertwiner Audit",
		Reason:      "Gate491 proves edge-support and kinetic positivity but blocks the full Hessian because the scalar contact frame is not yet canonically mapped into the broken electroweak generator frame.",
		PrimaryTask: "construct or reject a native DΦ and protected-to-broken intertwiner without importing W/Z masses, Higgs pole data, or gauge-coupling numerics",
	}
}

func validate(a Analysis) error {
	if !a.Inheritance.Executed || !a.Inheritance.Gate489FlavorAirlockClosed || !a.Inheritance.Gate490AnomalyLedgerStable || !a.Inheritance.NonFlavorFrontierSelected || !a.Inheritance.YukawaEntriesEnvironmental || !a.Inheritance.CKMOrientationEnvironmental || !a.Inheritance.NoObservedFlavorDataImported {
		return fmt.Errorf("Gate491 inheritance invalid: %+v", a.Inheritance)
	}
	if !a.Support.Executed || !a.Support.HiggsIsFiniteOneForm || !a.Support.EdgeMeasureSelected || a.Support.NodeMeasureAdmissible || a.Support.JDoubledEdgeCount != 10 || a.Support.PhysicalPoleMassDerived || a.Support.FullNumericalTOEClosed {
		return fmt.Errorf("Gate491 one-form support invalid: %+v", a.Support)
	}
	if !a.Kinetic.Executed || !a.Kinetic.TraceFunctionalFormalized || !a.Kinetic.DoubledCarrierEvaluated || a.Kinetic.EdgeTerms != 4 || a.Kinetic.QuarkEdges != 2 || a.Kinetic.LeptonEdges != 2 || !a.Kinetic.UsesHilbertSchmidtSquares || !a.Kinetic.PositiveSemidefinite || a.Kinetic.NegativeTermsPermitted || a.Kinetic.ImaginaryKineticPermitted || !a.Kinetic.GhostRiskEliminated || !a.Kinetic.StrictPositiveConditional || a.Kinetic.StrictPositiveProvedNumerically || a.Kinetic.NumericalZHComputed || !a.Kinetic.YukawaAmplitudesSealed || !a.Kinetic.CutoffMomentSealed || !a.Kinetic.SignConventionSealed {
		return fmt.Errorf("Gate491 kinetic positivity invalid: %+v", a.Kinetic)
	}
	if !a.Goldstone.Executed || a.Goldstone.ActiveRealDirections != 4 || a.Goldstone.RadialDirections != 1 || a.Goldstone.ScalarAngularDirections != 3 || a.Goldstone.ProtectedContactDirections != 3 || a.Goldstone.BrokenElectroweakDirections != 3 || !a.Goldstone.CountResonance || a.Goldstone.CanonicalProtectedToBrokenMapDerived || a.Goldstone.SU2ActionOnContactScalarDerived || a.Goldstone.CovariantDerivativeDerived || a.Goldstone.GaugeBosonMassMatrixDerived || a.Goldstone.GaugeEatingTheoremDerived {
		return fmt.Errorf("Gate491 Goldstone audit invalid: %+v", a.Goldstone)
	}
	if !a.Boundary.Executed || !a.Boundary.EdgeSupportNative || !a.Boundary.KineticSemidefiniteNative || !a.Boundary.GhostInstabilityBlocked || a.Boundary.FullHessianDerived || a.Boundary.VacuumStabilityDerived || a.Boundary.QuarticDerived || a.Boundary.HiggsMassDerived || a.Boundary.ContinuumScalarMatchingComplete {
		return fmt.Errorf("Gate491 stability boundary invalid: %+v", a.Boundary)
	}
	if !a.Firewall.Executed || a.Firewall.ObservedMassesImported || a.Firewall.ObservedYukawaImported || a.Firewall.ObservedCKMImported || a.Firewall.ObservedPMNSImported || a.Firewall.ObservedHiggsMassImported || a.Firewall.NativeYukawaMatrixWritten || a.Firewall.NativeCKMMatrixWritten || a.Firewall.NativeQuarticMassWritten || a.Firewall.NativeFlavorModuliChanged || a.Firewall.NativeFlavorDimAfter != NativeFlavorDim || a.Firewall.KXYCoeffDimAfter != KXYCoeffDim {
		return fmt.Errorf("Gate491 firewall invalid: %+v", a.Firewall)
	}
	return nil
}

func truth(a Analysis) string {
	return fmt.Sprintf("Gate491 proves a bounded scalar-edge stability theorem: the Higgs carrier is a finite one-form on %d J-doubled edge slots, and its kinetic trace over %d scalar Dirac edge classes is positive-semidefinite, eliminating native negative/imaginary ghost kinetic terms. This is not yet a full Higgs potential or vacuum-stability theorem: numerical Z_H, the quartic, pole mass, W/Z mass matrix, covariant derivative, and Goldstone gauge-eating map remain sealed.", a.Support.JDoubledEdgeCount, a.Kinetic.EdgeTerms)
}

func FormatInheritance(i Inheritance) string {
	return fmt.Sprintf("%s: Gate489_airlock=%t Gate490_anomaly=%t nonflavor=%t Yukawa_env=%t CKM_env=%t observed_flavor=%t; %s", i.Verdict, i.Gate489FlavorAirlockClosed, i.Gate490AnomalyLedgerStable, i.NonFlavorFrontierSelected, i.YukawaEntriesEnvironmental, i.CKMOrientationEnvironmental, !i.NoObservedFlavorDataImported, i.Reason)
}

func FormatSupport(s OneFormSupport) string {
	return fmt.Sprintf("%s: oneform=%t edge_selected=%t node_admissible=%t J_edges=%d pole_mass=%t toe_closed=%t formula=%q; %s", s.Verdict, s.HiggsIsFiniteOneForm, s.EdgeMeasureSelected, s.NodeMeasureAdmissible, s.JDoubledEdgeCount, s.PhysicalPoleMassDerived, s.FullNumericalTOEClosed, s.SupportFormula, s.Reason)
}

func FormatKinetic(k KineticPositivity) string {
	return fmt.Sprintf("%s: trace=%t doubled=%t edges=%d quark=%d lepton=%d HS=%t semidef=%t negative=%t imaginary=%t ghost_blocked=%t strict_conditional=%t strict_numeric=%t numerical_ZH=%t Yukawa_sealed=%t f0_sealed=%t sign_sealed=%t; %s", k.Verdict, k.TraceFunctionalFormalized, k.DoubledCarrierEvaluated, k.EdgeTerms, k.QuarkEdges, k.LeptonEdges, k.UsesHilbertSchmidtSquares, k.PositiveSemidefinite, k.NegativeTermsPermitted, k.ImaginaryKineticPermitted, k.GhostRiskEliminated, k.StrictPositiveConditional, k.StrictPositiveProvedNumerically, k.NumericalZHComputed, k.YukawaAmplitudesSealed, k.CutoffMomentSealed, k.SignConventionSealed, k.Reason)
}

func FormatGoldstone(g GoldstoneAudit) string {
	return fmt.Sprintf("%s: active=%d radial=%d angular=%d protected=%d brokenEW=%d count=%t map=%t SU2_on_scalar=%t Dphi=%t mass_matrix=%t eating=%t; %s", g.Verdict, g.ActiveRealDirections, g.RadialDirections, g.ScalarAngularDirections, g.ProtectedContactDirections, g.BrokenElectroweakDirections, g.CountResonance, g.CanonicalProtectedToBrokenMapDerived, g.SU2ActionOnContactScalarDerived, g.CovariantDerivativeDerived, g.GaugeBosonMassMatrixDerived, g.GaugeEatingTheoremDerived, g.Reason)
}

func FormatBoundary(b StabilityBoundary) string {
	return fmt.Sprintf("%s: edge_native=%t kinetic_semidef=%t ghost_blocked=%t full_hessian=%t vacuum_stability=%t quartic=%t higgs_mass=%t continuum_matching=%t; %s", b.Verdict, b.EdgeSupportNative, b.KineticSemidefiniteNative, b.GhostInstabilityBlocked, b.FullHessianDerived, b.VacuumStabilityDerived, b.QuarticDerived, b.HiggsMassDerived, b.ContinuumScalarMatchingComplete, b.Reason)
}

func FormatFirewall(f Firewall) string {
	return fmt.Sprintf("%s: masses=%t Yukawa=%t CKM=%t PMNS=%t Higgs_obs=%t native_Yukawa=%t native_CKM=%t native_quartic_mass=%t flavor_changed=%t dim=%d KXY=%d; %s", f.Verdict, f.ObservedMassesImported, f.ObservedYukawaImported, f.ObservedCKMImported, f.ObservedPMNSImported, f.ObservedHiggsMassImported, f.NativeYukawaMatrixWritten, f.NativeCKMMatrixWritten, f.NativeQuarticMassWritten, f.NativeFlavorModuliChanged, f.NativeFlavorDimAfter, f.KXYCoeffDimAfter, f.Reason)
}

func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 491 Registry Audit — Scalar-Edge Stability and Higgs One-Form Positivity Audit\n\n")
	b.WriteString("## Verdict\n\n")
	for _, v := range []string{
		StatusGate490Inherited,
		StatusHiggsOneFormEdgeSupportInherited,
		StatusScalarKineticTracePositive,
		StatusGhostSignObstructionRemoved,
		StatusStrictPositivityConditional,
		StatusGoldstoneCountResonance,
		StatusFailedNumericalZHSealed,
		StatusFailedFullScalarHessianNotDerived,
		StatusFailedVacuumStabilityNotDerived,
		StatusFailedQuarticMassNotDerived,
		StatusFailedGaugeEatingMapNotDerived,
		StatusFirewallPreserved,
	} {
		b.WriteString("- `" + v + "`\n")
	}
	b.WriteString("\n## Inherited boundary\n\n")
	b.WriteString(FormatInheritance(a.Inheritance) + "\n\n")
	b.WriteString("Gate489 closed native flavor prediction, and Gate490 proved a non-flavor anomaly ledger. Gate491 therefore audits scalar-edge stability only through finite one-form support, kinetic positivity, and Goldstone count data. It does not import masses, Yukawa entries, CKM/PMNS, observed Higgs mass, or continuum pole matching.\n\n")

	b.WriteString("## Higgs one-form edge support\n\n")
	b.WriteString(FormatSupport(a.Support) + "\n\n")
	b.WriteString("The admissible support is the represented one-form module `Ω¹_D(A_F)`, projected onto the finite Dirac edge graph. The node measure remains rejected for this kinetic-support question.\n\n")

	b.WriteString("## Scalar kinetic positivity audit\n\n")
	b.WriteString(FormatKinetic(a.Kinetic) + "\n\n")
	b.WriteString("The kinetic carrier has the form `C_H · (3||Y_u||² + 3||Y_d||² + ||Y_e||² + ||Y_ν||²)` before sealed convention factors. This proves non-negativity and blocks negative/imaginary finite scalar kinetic ghosts. It does not compute numerical `Z_H`.\n\n")

	b.WriteString("## Goldstone and Hessian boundary\n\n")
	b.WriteString(FormatGoldstone(a.Goldstone) + "\n\n")
	b.WriteString(FormatBoundary(a.Boundary) + "\n\n")
	b.WriteString("The count-level resonance `4 = 1 + 3` matches three protected contact directions and three broken electroweak directions. But a full gauge-eating theorem still requires a native protected-to-broken intertwiner, scalar `SU(2)_L` action, covariant derivative, and gauge-boson mass matrix.\n\n")

	b.WriteString("## Firewall result\n\n")
	b.WriteString(FormatFirewall(a.Firewall) + "\n\n")
	b.WriteString("No native Higgs mass, quartic, Yukawa matrix, CKM/PMNS matrix, or continuum scalar matching was written.\n\n")

	b.WriteString("## Registry update\n\n")
	b.WriteString("### Native\n\n")
	for _, x := range a.Registry.NativeEntries {
		b.WriteString("- " + x + "\n")
	}
	b.WriteString("\n### Bridge\n\n")
	for _, x := range a.Registry.BridgeEntries {
		b.WriteString("- " + x + "\n")
	}
	b.WriteString("\n### Environmental\n\n")
	for _, x := range a.Registry.EnvironmentalEntries {
		b.WriteString("- " + x + "\n")
	}
	b.WriteString("\n### Failed routes\n\n")
	for _, x := range a.Registry.FailedRoutes {
		b.WriteString("- `" + x + "`\n")
	}
	b.WriteString("\n### Open theorems\n\n")
	for _, x := range a.Registry.OpenTheorems {
		b.WriteString("- " + x + "\n")
	}

	b.WriteString("\n## Next step\n\n")
	b.WriteString(fmt.Sprintf("**Gate %d — %s.** %s Primary task: %s\n\n", a.Next.Gate, a.Next.Title, a.Next.Reason, a.Next.PrimaryTask))
	b.WriteString("## Truth statement\n\n")
	b.WriteString(a.Truth + "\n")
	return b.String()
}
