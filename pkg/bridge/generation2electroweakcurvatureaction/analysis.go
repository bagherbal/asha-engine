// Package generation2electroweakcurvatureaction implements Gate 493:
// Full Electroweak Curvature Action and Gauge Hessian Selection Audit.
//
// Gate 492 found the bridge-level DΦ socket: a rank-three broken-image
// Goldstone diagnostic and a photon-null direction.  Gate 493 asks whether the
// full finite electroweak curvature action is now strong enough to select the
// gauge kinetic Hessian and promote the W/Z mass diagnostic to a native theorem.
//
// The result is intentionally conservative.  Earlier electroweak curvature
// gates already prove that the full connection {T1,T2,Z,Q} closes and that a
// field-strength carrier can be typed.  They also expose a positive one-parameter
// abelian completion family whose broken-coordinate slice can reach the
// diag(1,1,4) whitening candidate.  But the finite action still does not select
// the abelian coefficient, compute the second variation, or derive physical
// couplings/masses.  Gate 493 therefore preserves the electroweak action socket
// while blocking the native W/Z and weak-angle registry writes.
package generation2electroweakcurvatureaction

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/ewcurvature"
	"github.com/bagherbal/asha-engine/pkg/bridge/ewquadratic"
	"github.com/bagherbal/asha-engine/pkg/bridge/gaugekineticdiag"
	"github.com/bagherbal/asha-engine/pkg/bridge/generation2scalarcovariantintertwiner"
)

const (
	AuditID = "GATE493-FULL-ELECTROWEAK-CURVATURE-ACTION-GAUGE-HESSIAN-SELECTION-AUDIT"

	StatusGate492Inherited                  = "CONDITIONAL_SUPPORT_GATE492_DPHI_GOLDSTONE_SOCKET_INHERITED"
	StatusFullEWConnectionClosed            = "CONDITIONAL_SUPPORT_FULL_ELECTROWEAK_CONNECTION_CLOSED"
	StatusFieldStrengthCarrierTyped         = "CONDITIONAL_SUPPORT_FIELD_STRENGTH_CARRIER_TYPED"
	StatusSemisimpleCurvatureRankThree      = "CONDITIONAL_SUPPORT_SEMISIMPLE_CURVATURE_RANK_THREE"
	StatusAbelianNullDirectionIdentified    = "CONDITIONAL_SUPPORT_ABELIAN_NULL_DIRECTION_IDENTIFIED"
	StatusEWQuadraticFamilyTyped            = "CONDITIONAL_SUPPORT_ELECTROWEAK_QUADRATIC_ACTION_FAMILY_TYPED"
	StatusPositiveAbelianCompletionFamily   = "CONDITIONAL_SUPPORT_POSITIVE_ABELIAN_COMPLETION_FAMILY_EXISTS"
	StatusDiag114ReachableAsBridgeCandidate = "CONDITIONAL_SUPPORT_DIAG114_REACHABLE_AS_WHITENING_CANDIDATE"
	StatusCoupledScalarGaugeSocketTyped     = "CONDITIONAL_SUPPORT_COUPLED_SCALAR_GAUGE_ACTION_SOCKET_TYPED"
	StatusFirewallPreserved                 = "FIREWALL_PRESERVED_NO_ELECTROWEAK_OR_FLAVOR_DATA_IMPORTED"
	StatusEWRegistryWriteBlocked            = "FIREWALL_BLOCKED_PHYSICAL_ELECTROWEAK_REGISTRY_WRITE"

	StatusFailedNativeEWActionNotDerived      = "FAILED_ROUTE_NATIVE_ELECTROWEAK_CURVATURE_ACTION_NOT_DERIVED"
	StatusFailedSecondVariationNotComputed    = "FAILED_ROUTE_FULL_ELECTROWEAK_ACTION_SECOND_VARIATION_NOT_COMPUTED"
	StatusFailedAbelianCoefficientNotSelected = "FAILED_ROUTE_U1_COMPLETION_COEFFICIENT_NOT_SELECTED"
	StatusFailedGaugeHessianNotSelected       = "FAILED_ROUTE_GAUGE_HESSIAN_NOT_ACTION_SELECTED"
	StatusFailedDiag114NotActionSelected      = "FAILED_ROUTE_DIAG114_NOT_ACTION_SELECTED"
	StatusFailedScalarGaugeActionNotNative    = "FAILED_ROUTE_COUPLED_SCALAR_GAUGE_ACTION_NOT_NATIVE"
	StatusFailedPhysicalCouplingsNotDerived   = "FAILED_ROUTE_PHYSICAL_GAUGE_COUPLINGS_NOT_DERIVED"
	StatusFailedWeakAngleNotDerived           = "FAILED_ROUTE_WEAK_MIXING_ANGLE_NOT_DERIVED"
	StatusFailedWZMassMatrixNotDerived        = "FAILED_ROUTE_PHYSICAL_WZ_MASS_MATRIX_NOT_DERIVED"
	StatusFailedHiggsVacuumScaleNotDerived    = "FAILED_ROUTE_HIGGS_VEV_AND_SCALAR_NORMALIZATION_NOT_DERIVED"
	StatusGate494RedirectDefined              = "CONDITIONAL_SUPPORT_GATE494_ABELIAN_COEFFICIENT_SELECTION_REDIRECT_DEFINED"
)

const eps = 1e-8

type Inheritance struct {
	Executed                         bool
	Gate492DphiTemplateFound         bool
	Gate492GoldstoneIntertwinerFound bool
	Gate492PhotonExemptionConfirmed  bool
	Gate492WZSignatureDiagnostic     bool
	Gate492NativeDphiStillBlocked    bool
	Gate492WZMassWriteBlocked        bool
	FullCurvatureNextGateRequested   bool
	NoMassFlavorDataImported         bool
	Verdict                          string
	Reason                           string
}

type CurvatureCarrierAudit struct {
	Executed                    bool
	Variables                   []string
	Dimension                   int
	Closed                      bool
	ClosureResidual             float64
	FullFieldStrengthTyped      bool
	CurvatureCarrierTyped       bool
	CurvatureQuadraticCandidate bool
	AdjointRank                 int
	AdjointMetricPositive       bool
	AbelianNullVector           []float64
	AbelianNullDirectionFound   bool
	U1KineticSelected           bool
	SecondVariationComputed     bool
	PhysicalCouplingsDerived    bool
	NativeCurvatureAction       bool
	Verdict                     string
	Reason                      string
}

type QuadraticActionAudit struct {
	Executed                       bool
	FullQuadraticFamilyTyped       bool
	PositiveQuadraticFamilyExists  bool
	SemisimpleRank                 int
	SemisimplePositiveSemidefinite bool
	AbelianCompletionTyped         bool
	AbelianCoefficientSelected     bool
	AbelianPositiveFor             string
	Diag114ReachableInFamily       bool
	Diag114Kappa                   float64
	Diag114SelectedByAction        bool
	GaugeKineticHessianSelected    bool
	PhysicalCouplingsOrMasses      bool
	Verdict                        string
	Reason                         string
}

type GaugeHessianAudit struct {
	Executed                    bool
	BrokenDiag114CandidateFound bool
	CandidatePositive           bool
	WhitenedExact               bool
	NeutralFactor               float64
	SelectedByFiniteAction      bool
	ScalarKineticActionSelected bool
	GaugeHessianSelected        bool
	PhysicalCouplingsDerived    bool
	PhysicalMassesDerived       bool
	Verdict                     string
	Reason                      string
}

type CoupledScalarGaugeAudit struct {
	Executed                            bool
	DphiTemplateAvailable               bool
	GoldstoneImageRank                  int
	PhotonNullDirection                 bool
	EWQuadraticFamilyAvailable          bool
	CoupledActionSocketTyped            bool
	NativeScalarGaugeActionDerived      bool
	ScalarKineticMetricNative           bool
	ScalarVacuumOrientationNative       bool
	GaugeHessianCouplingsActionSelected bool
	HiggsVEVDerived                     bool
	PhysicalWZMassMatrixDerived         bool
	WeakMixingAngleDerived              bool
	Verdict                             string
	Reason                              string
}

type Boundary struct {
	Executed                        bool
	EWCurvatureCarrierPromotable    bool
	EWQuadraticFamilyPromotable     bool
	Diag114BridgeCandidatePreserved bool
	NativeEWCurvatureActionDerived  bool
	SecondVariationComputed         bool
	AbelianCoefficientSelected      bool
	GaugeHessianActionSelected      bool
	PhysicalGaugeCouplingsDerived   bool
	WeakMixingAngleDerived          bool
	PhysicalWZMassMatrixDerived     bool
	NativeElectroweakMassTheorem    bool
	Verdict                         string
	Reason                          string
}

type Firewall struct {
	Executed                   bool
	ObservedWMassImported      bool
	ObservedZMassImported      bool
	ObservedHiggsMassImported  bool
	FermiConstantImported      bool
	WeakAngleImported          bool
	FineStructureImported      bool
	GaugeCouplingImported      bool
	YukawaImported             bool
	CKMPMNSImported            bool
	NativeWZMassWritten        bool
	NativeWeakAngleWritten     bool
	NativeGaugeCouplingWritten bool
	NativeHiggsVEVWritten      bool
	Verdict                    string
	Reason                     string
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
	Curvature   CurvatureCarrierAudit
	Quadratic   QuadraticActionAudit
	Gauge       GaugeHessianAudit
	Coupled     CoupledScalarGaugeAudit
	Boundary    Boundary
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
	g492, err := generation2scalarcovariantintertwiner.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate492 scalar covariant audit: %w", err)
	}
	ewc, err := ewcurvature.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit electroweak curvature audit: %w", err)
	}
	ewq, err := ewquadratic.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit electroweak quadratic audit: %w", err)
	}
	gkd, err := gaugekineticdiag.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit gauge kinetic diag audit: %w", err)
	}

	a := Analysis{}
	a.Inheritance = buildInheritance(g492)
	a.Curvature = buildCurvatureAudit(ewc)
	a.Quadratic = buildQuadraticAudit(ewq)
	a.Gauge = buildGaugeHessianAudit(gkd)
	a.Coupled = buildCoupledAudit(g492, ewq)
	a.Boundary = buildBoundary(a.Curvature, a.Quadratic, a.Gauge, a.Coupled)
	a.Firewall = buildFirewall()
	a.Registry = buildRegistryUpdate(a)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance(g492 generation2scalarcovariantintertwiner.Analysis) Inheritance {
	return Inheritance{
		Executed:                         true,
		Gate492DphiTemplateFound:         g492.Dphi.AbstractTemplateAvailable,
		Gate492GoldstoneIntertwinerFound: g492.Intertwiner.GoldstoneImageDiagnostic && g492.Intertwiner.BrokenImageRank == 3,
		Gate492PhotonExemptionConfirmed:  g492.Photon.QEMAnnihilatesVacuum,
		Gate492WZSignatureDiagnostic:     g492.Dphi.DimensionlessWZPhotonSignature && g492.Dphi.MassMatrixRank == 3,
		Gate492NativeDphiStillBlocked:    !g492.Dphi.NativeDphiDerived && !g492.Boundary.NativeDphiDerived,
		Gate492WZMassWriteBlocked:        g492.Firewall.NativeWZMassWritten == false && g492.Boundary.PhysicalMassMatrixDerived == false,
		FullCurvatureNextGateRequested:   g492.Next.Gate == 493,
		NoMassFlavorDataImported:         !g492.Firewall.ObservedWMassImported && !g492.Firewall.ObservedZMassImported && !g492.Firewall.ObservedHiggsMassImported && !g492.Firewall.WeakAngleImported && !g492.Firewall.YukawaImported && !g492.Firewall.CKMPMNSImported,
		Verdict:                          StatusGate492Inherited,
		Reason:                           "Gate492 supplies only a bridge DΦ/Goldstone/photon diagnostic socket and explicitly asks for a full electroweak curvature/action audit before any W/Z promotion.",
	}
}

func buildCurvatureAudit(c ewcurvature.Analysis) CurvatureCarrierAudit {
	vars := make([]string, 0, len(c.Variables))
	for _, v := range c.Variables {
		vars = append(vars, v.Basis)
	}
	nullFound := len(c.AbelianNullVector) == 4 && vectorNorm(c.AbelianNullVector) > eps
	return CurvatureCarrierAudit{
		Executed:                    true,
		Variables:                   vars,
		Dimension:                   c.Dimension,
		Closed:                      c.Closed,
		ClosureResidual:             c.ClosureResidual,
		FullFieldStrengthTyped:      c.FullFieldStrengthTyped,
		CurvatureCarrierTyped:       c.CurvatureCarrierTyped,
		CurvatureQuadraticCandidate: c.CurvatureQuadraticCandidate,
		AdjointRank:                 c.AdjointRank,
		AdjointMetricPositive:       c.AdjointMetricPositive,
		AbelianNullVector:           append([]float64(nil), c.AbelianNullVector...),
		AbelianNullDirectionFound:   nullFound,
		U1KineticSelected:           c.U1KineticSelected,
		SecondVariationComputed:     c.SecondVariationComputed,
		PhysicalCouplingsDerived:    c.PhysicalCouplings,
		NativeCurvatureAction:       c.SecondVariationComputed && c.U1KineticSelected && c.AdjointMetricPositive,
		Verdict:                     StatusFieldStrengthCarrierTyped,
		Reason:                      "The full electroweak carrier closes and types a field-strength object, but the adjoint diagnostic is rank-three and leaves the abelian direction null; it is not yet a native positive physical action.",
	}
}

func buildQuadraticAudit(q ewquadratic.Analysis) QuadraticActionAudit {
	return QuadraticActionAudit{
		Executed:                       true,
		FullQuadraticFamilyTyped:       q.FullQuadraticActionFamilyTyped,
		PositiveQuadraticFamilyExists:  q.PositiveQuadraticFamilyExists,
		SemisimpleRank:                 q.SemisimpleRank,
		SemisimplePositiveSemidefinite: q.SemisimplePositiveSemidefinite,
		AbelianCompletionTyped:         q.AbelianCompletionTyped,
		AbelianCoefficientSelected:     q.AbelianCoefficientSelected,
		AbelianPositiveFor:             q.AbelianCompletion.PositiveFor,
		Diag114ReachableInFamily:       q.Diag114ReachableInFamily,
		Diag114Kappa:                   q.Diag114Kappa,
		Diag114SelectedByAction:        q.Diag114SelectedByAction,
		GaugeKineticHessianSelected:    q.GaugeKineticHessianSelected,
		PhysicalCouplingsOrMasses:      q.PhysicalCouplingsOrMasses,
		Verdict:                        StatusEWQuadraticFamilyTyped,
		Reason:                         "A positive full electroweak quadratic family exists after adding the abelian completion, but kappa_U1 remains a free bridge coefficient, not an action-selected theorem.",
	}
}

func buildGaugeHessianAudit(g gaugekineticdiag.Analysis) GaugeHessianAudit {
	return GaugeHessianAudit{
		Executed:                    true,
		BrokenDiag114CandidateFound: len(g.CandidateDiagonal) == 3 && math.Abs(g.CandidateDiagonal[0]-1) < eps && math.Abs(g.CandidateDiagonal[1]-1) < eps && math.Abs(g.CandidateDiagonal[2]-4) < eps,
		CandidatePositive:           g.CandidatePositive,
		WhitenedExact:               g.WhitenedExact,
		NeutralFactor:               g.NeutralFactor,
		SelectedByFiniteAction:      g.SelectedByFiniteAction,
		ScalarKineticActionSelected: g.ScalarKineticActionSelected,
		GaugeHessianSelected:        g.GaugeHessianSelected,
		PhysicalCouplingsDerived:    g.PhysicalCouplingsDerived,
		PhysicalMassesDerived:       g.PhysicalMassesDerived,
		Verdict:                     StatusDiag114ReachableAsBridgeCandidate,
		Reason:                      "The diag(1,1,4) broken-coordinate Hessian remains the exact whitening candidate, but it has not been selected by a finite action second variation.",
	}
}

func buildCoupledAudit(g492 generation2scalarcovariantintertwiner.Analysis, q ewquadratic.Analysis) CoupledScalarGaugeAudit {
	socket := g492.Dphi.AbstractTemplateAvailable && g492.Intertwiner.BrokenImageRank == 3 && g492.Photon.QEMAnnihilatesVacuum && q.FullQuadraticActionFamilyTyped
	return CoupledScalarGaugeAudit{
		Executed:                            true,
		DphiTemplateAvailable:               g492.Dphi.AbstractTemplateAvailable,
		GoldstoneImageRank:                  g492.Intertwiner.BrokenImageRank,
		PhotonNullDirection:                 g492.Photon.QEMAnnihilatesVacuum,
		EWQuadraticFamilyAvailable:          q.FullQuadraticActionFamilyTyped && q.PositiveQuadraticFamilyExists,
		CoupledActionSocketTyped:            socket,
		NativeScalarGaugeActionDerived:      false,
		ScalarKineticMetricNative:           g492.Boundary.NativeKineticMetricDerived,
		ScalarVacuumOrientationNative:       g492.Boundary.NativeVacuumOrientationDerived,
		GaugeHessianCouplingsActionSelected: g492.Boundary.NativeGaugeHessianDerived && q.GaugeKineticHessianSelected,
		HiggsVEVDerived:                     false,
		PhysicalWZMassMatrixDerived:         g492.Boundary.PhysicalMassMatrixDerived,
		WeakMixingAngleDerived:              false,
		Verdict:                             StatusCoupledScalarGaugeSocketTyped,
		Reason:                              "The scalar DΦ diagnostic and the electroweak quadratic family can be placed in the same bridge template, but the scalar metric, vacuum orientation, gauge Hessian, and couplings are still unselected.",
	}
}

func buildBoundary(c CurvatureCarrierAudit, q QuadraticActionAudit, g GaugeHessianAudit, s CoupledScalarGaugeAudit) Boundary {
	promotableCarrier := c.Closed && c.FullFieldStrengthTyped && c.AbelianNullDirectionFound
	promotableFamily := q.FullQuadraticFamilyTyped && q.PositiveQuadraticFamilyExists && q.AbelianCompletionTyped
	diagPreserved := g.BrokenDiag114CandidateFound && g.CandidatePositive && g.WhitenedExact && q.Diag114ReachableInFamily
	return Boundary{
		Executed:                        true,
		EWCurvatureCarrierPromotable:    promotableCarrier,
		EWQuadraticFamilyPromotable:     promotableFamily,
		Diag114BridgeCandidatePreserved: diagPreserved,
		NativeEWCurvatureActionDerived:  c.NativeCurvatureAction,
		SecondVariationComputed:         c.SecondVariationComputed,
		AbelianCoefficientSelected:      q.AbelianCoefficientSelected,
		GaugeHessianActionSelected:      q.GaugeKineticHessianSelected || g.GaugeHessianSelected,
		PhysicalGaugeCouplingsDerived:   c.PhysicalCouplingsDerived || q.PhysicalCouplingsOrMasses || g.PhysicalCouplingsDerived,
		WeakMixingAngleDerived:          s.WeakMixingAngleDerived,
		PhysicalWZMassMatrixDerived:     s.PhysicalWZMassMatrixDerived || g.PhysicalMassesDerived,
		NativeElectroweakMassTheorem:    false,
		Verdict:                         StatusEWRegistryWriteBlocked,
		Reason:                          "Gate493 preserves the full electroweak curvature/action socket but blocks promotion because no finite second variation selects kappa_U1, the gauge Hessian, the scalar normalization, or the physical W/Z matrix.",
	}
}

func buildFirewall() Firewall {
	return Firewall{
		Executed:                   true,
		ObservedWMassImported:      false,
		ObservedZMassImported:      false,
		ObservedHiggsMassImported:  false,
		FermiConstantImported:      false,
		WeakAngleImported:          false,
		FineStructureImported:      false,
		GaugeCouplingImported:      false,
		YukawaImported:             false,
		CKMPMNSImported:            false,
		NativeWZMassWritten:        false,
		NativeWeakAngleWritten:     false,
		NativeGaugeCouplingWritten: false,
		NativeHiggsVEVWritten:      false,
		Verdict:                    StatusFirewallPreserved,
		Reason:                     "No observed electroweak scale, weak angle, gauge coupling, Higgs/VEV datum, Yukawa texture, CKM, or PMNS data are used; all numerical electroweak data remain environmental/bridge.",
	}
}

func buildRegistryUpdate(a Analysis) RegistryUpdate {
	return RegistryUpdate{
		NativeEntries: []string{
			"Full electroweak Lie carrier {T1,T2,Z,Q} closes as a structural field-strength carrier.",
			"The semisimple curvature diagnostic has rank three and exposes an abelian null direction requiring completion.",
		},
		BridgeEntries: []string{
			"Positive quadratic family K(kappa)=K_SU2+kappa(Q-Z)(Q-Z)^T is typed for kappa_U1>0.",
			"diag(1,1,4) remains reachable as the broken-coordinate whitening candidate at kappa_U1=6 in the chosen convention.",
			"The scalar DΦ diagnostic, Goldstone image rank three, photon-null direction, and EW quadratic family form a consistent coupled bridge socket.",
		},
		EnvironmentalEntries: []string{
			"Physical W and Z masses, Higgs VEV/Fermi constant, weak mixing angle, fine-structure constant, and gauge coupling normalizations.",
			"Yukawa/CKM/PMNS data remain irrelevant to this non-flavor audit and stay quarantined.",
		},
		FailedRoutes: []string{
			StatusFailedNativeEWActionNotDerived,
			StatusFailedSecondVariationNotComputed,
			StatusFailedAbelianCoefficientNotSelected,
			StatusFailedGaugeHessianNotSelected,
			StatusFailedDiag114NotActionSelected,
			StatusFailedScalarGaugeActionNotNative,
			StatusFailedPhysicalCouplingsNotDerived,
			StatusFailedWeakAngleNotDerived,
			StatusFailedWZMassMatrixNotDerived,
			StatusFailedHiggsVacuumScaleNotDerived,
		},
		OpenTheorems: []string{
			"Derive kappa_U1 from finite spectral/action data instead of the whitening convention.",
			"Compute an actual second variation δ²S/δA_iδA_j of the full finite electroweak action.",
			"Derive scalar/contact kinetic normalization and vacuum orientation before interpreting W/Z mass eigenvalues.",
			"Only after action-selected kinetic terms may continuum coupling, weak-angle, RG, and physical mass bridges be opened.",
		},
	}
}

func buildNext() NextStep {
	return NextStep{
		Gate:        494,
		Title:       "Abelian U(1) Completion Coefficient Selection Audit",
		Reason:      "Gate493 shows the full electroweak action family exists but leaves kappa_U1 free; the only mathematically sharp next move is to search finite trace/spectral/unimodularity data for a native abelian completion coefficient.",
		PrimaryTask: "test whether kappa_U1 is selected by a native finite spectral trace, representation metric, unimodularity constraint, or topological normalization without importing theta_W, alpha, W/Z masses, or continuum RG data",
	}
}

func validate(a Analysis) error {
	switch {
	case !a.Inheritance.Executed || !a.Inheritance.Gate492DphiTemplateFound || !a.Inheritance.FullCurvatureNextGateRequested:
		return fmt.Errorf("invalid Gate493 inheritance: %+v", a.Inheritance)
	case !a.Curvature.Executed || !a.Curvature.Closed || !a.Curvature.FullFieldStrengthTyped || a.Curvature.Dimension != 4 || a.Curvature.AdjointRank != 3 || !a.Curvature.AbelianNullDirectionFound:
		return fmt.Errorf("invalid curvature audit: %+v", a.Curvature)
	case a.Curvature.SecondVariationComputed || a.Curvature.U1KineticSelected || a.Curvature.NativeCurvatureAction || a.Curvature.PhysicalCouplingsDerived:
		return fmt.Errorf("curvature audit over-promoted native action: %+v", a.Curvature)
	case !a.Quadratic.Executed || !a.Quadratic.FullQuadraticFamilyTyped || !a.Quadratic.PositiveQuadraticFamilyExists || !a.Quadratic.AbelianCompletionTyped || !a.Quadratic.Diag114ReachableInFamily || math.Abs(a.Quadratic.Diag114Kappa-6) > eps:
		return fmt.Errorf("invalid quadratic audit: %+v", a.Quadratic)
	case a.Quadratic.AbelianCoefficientSelected || a.Quadratic.Diag114SelectedByAction || a.Quadratic.GaugeKineticHessianSelected || a.Quadratic.PhysicalCouplingsOrMasses:
		return fmt.Errorf("quadratic audit over-promoted selected action: %+v", a.Quadratic)
	case !a.Gauge.Executed || !a.Gauge.BrokenDiag114CandidateFound || !a.Gauge.CandidatePositive || !a.Gauge.WhitenedExact:
		return fmt.Errorf("invalid gauge Hessian audit: %+v", a.Gauge)
	case a.Gauge.SelectedByFiniteAction || a.Gauge.GaugeHessianSelected || a.Gauge.PhysicalCouplingsDerived || a.Gauge.PhysicalMassesDerived:
		return fmt.Errorf("gauge Hessian audit over-promoted selected data: %+v", a.Gauge)
	case !a.Coupled.Executed || !a.Coupled.CoupledActionSocketTyped || a.Coupled.NativeScalarGaugeActionDerived || a.Coupled.GaugeHessianCouplingsActionSelected || a.Coupled.PhysicalWZMassMatrixDerived || a.Coupled.WeakMixingAngleDerived:
		return fmt.Errorf("invalid coupled scalar/gauge audit: %+v", a.Coupled)
	case !a.Boundary.Executed || !a.Boundary.EWCurvatureCarrierPromotable || !a.Boundary.EWQuadraticFamilyPromotable || !a.Boundary.Diag114BridgeCandidatePreserved:
		return fmt.Errorf("invalid boundary: %+v", a.Boundary)
	case a.Boundary.NativeEWCurvatureActionDerived || a.Boundary.SecondVariationComputed || a.Boundary.AbelianCoefficientSelected || a.Boundary.GaugeHessianActionSelected || a.Boundary.PhysicalGaugeCouplingsDerived || a.Boundary.WeakMixingAngleDerived || a.Boundary.PhysicalWZMassMatrixDerived || a.Boundary.NativeElectroweakMassTheorem:
		return fmt.Errorf("boundary over-promoted electroweak theorem: %+v", a.Boundary)
	case !a.Firewall.Executed || a.Firewall.ObservedWMassImported || a.Firewall.ObservedZMassImported || a.Firewall.ObservedHiggsMassImported || a.Firewall.FermiConstantImported || a.Firewall.WeakAngleImported || a.Firewall.FineStructureImported || a.Firewall.GaugeCouplingImported || a.Firewall.YukawaImported || a.Firewall.CKMPMNSImported || a.Firewall.NativeWZMassWritten || a.Firewall.NativeWeakAngleWritten || a.Firewall.NativeGaugeCouplingWritten || a.Firewall.NativeHiggsVEVWritten:
		return fmt.Errorf("firewall leak: %+v", a.Firewall)
	}
	return nil
}

func truth(a Analysis) string {
	return "Gate493 proves the correct electroweak action boundary: ASHA has a closed full {T1,T2,Z,Q} curvature carrier and a positive abelian-completed quadratic family, with diag(1,1,4) reachable as a bridge whitening candidate. But a family of actions is not an action-selected Hessian. Because kappa_U1, scalar normalization, vacuum orientation, gauge couplings, weak angle, and W/Z mass scale are not derived by finite second variation, the native electroweak mass theorem remains blocked and the firewall stays intact."
}

func vectorNorm(xs []float64) float64 {
	s := 0.0
	for _, x := range xs {
		s += x * x
	}
	return math.Sqrt(s)
}

func boolWord(v bool) string {
	if v {
		return "yes"
	}
	return "no"
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("executed=%v dphi_template=%v goldstone_rank3=%v photon_exempt=%v wz_signature=%v native_dphi_blocked=%v wz_write_blocked=%v gate493_requested=%v no_mass_flavor_data=%v verdict=%s reason=%s", x.Executed, x.Gate492DphiTemplateFound, x.Gate492GoldstoneIntertwinerFound, x.Gate492PhotonExemptionConfirmed, x.Gate492WZSignatureDiagnostic, x.Gate492NativeDphiStillBlocked, x.Gate492WZMassWriteBlocked, x.FullCurvatureNextGateRequested, x.NoMassFlavorDataImported, x.Verdict, x.Reason)
}

func FormatCurvature(x CurvatureCarrierAudit) string {
	return fmt.Sprintf("variables=%s dim=%d closed=%v residual=%.3g field_strength_typed=%v curvature_carrier=%v quadratic_candidate=%v adjoint_rank=%d adjoint_positive=%v abelian_null=%s u1_selected=%v second_variation=%v physical_couplings=%v native_action=%v verdict=%s reason=%s", strings.Join(x.Variables, ","), x.Dimension, x.Closed, x.ClosureResidual, x.FullFieldStrengthTyped, x.CurvatureCarrierTyped, x.CurvatureQuadraticCandidate, x.AdjointRank, x.AdjointMetricPositive, formatFloatList(x.AbelianNullVector), x.U1KineticSelected, x.SecondVariationComputed, x.PhysicalCouplingsDerived, x.NativeCurvatureAction, x.Verdict, x.Reason)
}

func FormatQuadratic(x QuadraticActionAudit) string {
	return fmt.Sprintf("family_typed=%v positive_family=%v semisimple_rank=%d semisimple_psd=%v abelian_completion=%v abelian_selected=%v positive_for=%q diag114_reachable=%v kappa=%.10f diag114_action_selected=%v hessian_selected=%v physical_couplings_masses=%v verdict=%s reason=%s", x.FullQuadraticFamilyTyped, x.PositiveQuadraticFamilyExists, x.SemisimpleRank, x.SemisimplePositiveSemidefinite, x.AbelianCompletionTyped, x.AbelianCoefficientSelected, x.AbelianPositiveFor, x.Diag114ReachableInFamily, x.Diag114Kappa, x.Diag114SelectedByAction, x.GaugeKineticHessianSelected, x.PhysicalCouplingsOrMasses, x.Verdict, x.Reason)
}

func FormatGauge(x GaugeHessianAudit) string {
	return fmt.Sprintf("diag114_candidate=%v positive=%v whitened_exact=%v neutral_factor=%.10f selected_by_action=%v scalar_kinetic_selected=%v gauge_hessian_selected=%v physical_couplings=%v physical_masses=%v verdict=%s reason=%s", x.BrokenDiag114CandidateFound, x.CandidatePositive, x.WhitenedExact, x.NeutralFactor, x.SelectedByFiniteAction, x.ScalarKineticActionSelected, x.GaugeHessianSelected, x.PhysicalCouplingsDerived, x.PhysicalMassesDerived, x.Verdict, x.Reason)
}

func FormatCoupled(x CoupledScalarGaugeAudit) string {
	return fmt.Sprintf("dphi=%v goldstone_rank=%d photon_null=%v ew_quadratic=%v coupled_socket=%v native_scalar_gauge_action=%v scalar_metric_native=%v vacuum_native=%v gauge_hessian_couplings_selected=%v higgs_vev=%v wz_mass=%v weak_angle=%v verdict=%s reason=%s", x.DphiTemplateAvailable, x.GoldstoneImageRank, x.PhotonNullDirection, x.EWQuadraticFamilyAvailable, x.CoupledActionSocketTyped, x.NativeScalarGaugeActionDerived, x.ScalarKineticMetricNative, x.ScalarVacuumOrientationNative, x.GaugeHessianCouplingsActionSelected, x.HiggsVEVDerived, x.PhysicalWZMassMatrixDerived, x.WeakMixingAngleDerived, x.Verdict, x.Reason)
}

func FormatBoundary(x Boundary) string {
	return fmt.Sprintf("carrier_promotable=%v quadratic_family_promotable=%v diag114_preserved=%v native_ew_action=%v second_variation=%v abelian_selected=%v gauge_hessian_selected=%v physical_couplings=%v weak_angle=%v wz_mass=%v native_ew_mass_theorem=%v verdict=%s reason=%s", x.EWCurvatureCarrierPromotable, x.EWQuadraticFamilyPromotable, x.Diag114BridgeCandidatePreserved, x.NativeEWCurvatureActionDerived, x.SecondVariationComputed, x.AbelianCoefficientSelected, x.GaugeHessianActionSelected, x.PhysicalGaugeCouplingsDerived, x.WeakMixingAngleDerived, x.PhysicalWZMassMatrixDerived, x.NativeElectroweakMassTheorem, x.Verdict, x.Reason)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("observed_W=%v observed_Z=%v observed_Higgs=%v Fermi=%v weak_angle=%v alpha=%v gauge_coupling=%v Yukawa=%v CKM_PMNS=%v native_WZ=%v native_theta=%v native_gauge_coupling=%v native_vev=%v verdict=%s reason=%s", x.ObservedWMassImported, x.ObservedZMassImported, x.ObservedHiggsMassImported, x.FermiConstantImported, x.WeakAngleImported, x.FineStructureImported, x.GaugeCouplingImported, x.YukawaImported, x.CKMPMNSImported, x.NativeWZMassWritten, x.NativeWeakAngleWritten, x.NativeGaugeCouplingWritten, x.NativeHiggsVEVWritten, x.Verdict, x.Reason)
}

func formatFloatList(xs []float64) string {
	parts := make([]string, len(xs))
	for i, x := range xs {
		if math.Abs(x) < 1e-12 {
			x = 0
		}
		parts[i] = fmt.Sprintf("%.10g", x)
	}
	return "[" + strings.Join(parts, ",") + "]"
}

func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 493 Registry Audit — Full Electroweak Curvature Action and Gauge Hessian Selection Audit\n\n")
	b.WriteString("## Verdict\n\n")
	for _, v := range []string{
		StatusGate492Inherited,
		StatusFullEWConnectionClosed,
		StatusFieldStrengthCarrierTyped,
		StatusSemisimpleCurvatureRankThree,
		StatusAbelianNullDirectionIdentified,
		StatusEWQuadraticFamilyTyped,
		StatusPositiveAbelianCompletionFamily,
		StatusDiag114ReachableAsBridgeCandidate,
		StatusCoupledScalarGaugeSocketTyped,
		StatusFailedNativeEWActionNotDerived,
		StatusFailedSecondVariationNotComputed,
		StatusFailedAbelianCoefficientNotSelected,
		StatusFailedGaugeHessianNotSelected,
		StatusFailedDiag114NotActionSelected,
		StatusFailedScalarGaugeActionNotNative,
		StatusFailedPhysicalCouplingsNotDerived,
		StatusFailedWeakAngleNotDerived,
		StatusFailedWZMassMatrixNotDerived,
		StatusFailedHiggsVacuumScaleNotDerived,
		StatusFirewallPreserved,
		StatusEWRegistryWriteBlocked,
	} {
		b.WriteString("- `" + v + "`\n")
	}

	b.WriteString("\n## Inherited boundary\n\n")
	b.WriteString(FormatInheritance(a.Inheritance) + "\n\n")
	b.WriteString("Gate492 provides the DΦ bridge template, rank-three Goldstone image diagnostic, and photon exemption. Gate493 may test the full electroweak curvature/action socket, but it may not import W/Z masses, Higgs VEV, Fermi constant, weak mixing angle, fine-structure constant, gauge couplings, Yukawas, CKM, or PMNS.\n\n")

	b.WriteString("## Full electroweak curvature carrier\n\n")
	b.WriteString(FormatCurvature(a.Curvature) + "\n\n")
	b.WriteString("The full connection closes only after the photon direction is included. The semisimple adjoint diagnostic has rank three and leaves the pure abelian direction null, so closure alone cannot be a positive physical gauge Hessian.\n\n")

	b.WriteString("## Quadratic action family and abelian completion\n\n")
	b.WriteString(FormatQuadratic(a.Quadratic) + "\n\n")
	b.WriteString("The one-parameter abelian completion is the correct mathematical socket. In the current convention, the old `diag(1,1,4)` broken-coordinate candidate is reachable at `kappa_U1 = 6`, but reachability is not selection.\n\n")

	b.WriteString("## Gauge Hessian audit\n\n")
	b.WriteString(FormatGauge(a.Gauge) + "\n\n")
	b.WriteString("The whitening candidate remains coherent and useful, but it is still not the second variation of a finite action. Therefore it cannot fix `g_2`, `g_Y`, `theta_W`, `alpha`, or W/Z masses.\n\n")

	b.WriteString("## Coupled scalar-gauge action socket\n\n")
	b.WriteString(FormatCoupled(a.Coupled) + "\n\n")
	b.WriteString("The scalar and gauge diagnostics fit together structurally, but the scalar metric, scalar vacuum orientation, abelian coefficient, and gauge Hessian are all still unselected. This blocks native electroweak mass promotion.\n\n")

	b.WriteString("## Firewall result\n\n")
	b.WriteString(FormatBoundary(a.Boundary) + "\n\n")
	b.WriteString(FormatFirewall(a.Firewall) + "\n\n")
	b.WriteString("No physical electroweak number entered the theorem lane, and no native W/Z, weak-angle, gauge-coupling, Higgs-VEV, Yukawa, CKM, or PMNS registry write occurred.\n\n")

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
	b.WriteString(fmt.Sprintf("**Gate %d — %s.** %s Primary task: %s.\n\n", a.Next.Gate, a.Next.Title, a.Next.Reason, a.Next.PrimaryTask))
	b.WriteString("## Truth statement\n\n")
	b.WriteString(a.Truth + "\n")
	return b.String()
}
