// Package generation2scalarquaternionicmomentewcurvatureprojectionaudit implements Gate 563:
// Scalar/Quaternionic Moment to Electroweak Curvature Projection Audit.
//
// Gate 562 located the lawful structural target for the sealed scalar Pauli/Hopf
// moment map: the quaternionic weak socket Im(H) inside A_F=C⊕H⊕M3(C). Gate 563
// asks whether that scalar/quaternionic moment already enters the finite
// one-form, curvature, scalar kinetic projection, or electroweak mass dynamics
// lanes. The result is intentionally split: the finite one-form lane contains
// the scalar SU(2)/H doublet and the product spectral-action lane contains a
// symbolic D_phi squared kinetic channel, but no current theorem inserts the
// moment map into a native curvature projection or derives U(1) mixing,
// kinetic normalization, physical W/Z/photon states, Higgs potential, masses,
// or flavor data.
package generation2scalarquaternionicmomentewcurvatureprojectionaudit

import (
	"fmt"
	"strings"
	"sync"

	gate493 "github.com/bagherbal/asha-engine/pkg/bridge/generation2electroweakcurvatureaction"
	gate499 "github.com/bagherbal/asha-engine/pkg/bridge/generation2innerfluctuationdphiprovenance"
	gate562 "github.com/bagherbal/asha-engine/pkg/bridge/generation2paulihopfquaternionicweaksocketaudit"
	gate500 "github.com/bagherbal/asha-engine/pkg/bridge/generation2productspectralactionkineticprojection"
)

const (
	AuditID = "GATE563-SCALAR-QUATERNIONIC-MOMENT-ELECTROWEAK-CURVATURE-PROJECTION-AUDIT"

	StatusGate562Inherited                 = "CONDITIONAL_SUPPORT_GATE562_SCALAR_QUATERNIONIC_MOMENT_BRIDGE_INHERITED"
	StatusFiniteOneFormScalarLaneRecovered = "PASS_FINITE_ONE_FORM_SCALAR_SU2_H_DOUBLE_MODULE_LANE_RECOVERED"
	StatusQuaternionicActionStructural     = "PASS_IM_H_ACTION_ON_HPHI_STRUCTURAL_PAIRING_AVAILABLE"
	StatusMomentPairingBookkeeping         = "CONDITIONAL_SUPPORT_MOMENT_PAIRING_AVAILABLE_IN_SCALAR_QUATERNIONIC_REPRESENTATION_BOOKKEEPING"
	StatusProductDphiSquaredSymbolic       = "CONDITIONAL_SUPPORT_PRODUCT_SPECTRAL_ACTION_SYMBOLIC_DPHI_SQUARED_CHANNEL_PRESENT"
	StatusCurvatureSocketAvailable         = "CONDITIONAL_SUPPORT_ELECTROWEAK_CURVATURE_SOCKET_AVAILABLE_BRIDGE_LEVEL"
	StatusMomentNotInCurvatureProjection   = "FAILED_ROUTE_MOMENT_MAP_NOT_FOUND_IN_NATIVE_CURVATURE_OR_KINETIC_PROJECTION"
	StatusOrbitSplitRepresentationOnly     = "CONDITIONAL_SUPPORT_NONZERO_MU_STABILIZER_ORBIT_SPLIT_RECOGNIZED_REPRESENTATION_LEVEL"
	StatusNoCurvatureStabilizerProjection  = "FAILED_ROUTE_NO_NATIVE_CURVATURE_PROJECTION_DISTINGUISHES_STABILIZER_AND_ORBIT_COMPONENTS"
	StatusNoNativeU1PhotonDirection        = "FAILED_ROUTE_NO_NATIVE_ELECTROWEAK_U1_MIXING_OR_PHOTON_DIRECTION"
	StatusNoKineticNormalization           = "FAILED_ROUTE_NO_NATIVE_KINETIC_NORMALIZATION_FOR_WZ_MASS_DYNAMICS"
	StatusNoFlavorData                     = "FAILED_ROUTE_SCALAR_QUATERNIONIC_MOMENT_DOES_NOT_DERIVE_FLAVOR_DATA"
	StatusPreviousFirewallsPreserved       = "FIREWALL_PRESERVED_Q4_TAU_ETA_WSPATIAL_BOUNDARIES"
	StatusFirewallPreserved                = "FIREWALL_PRESERVED_GATE563_SCALAR_QUATERNIONIC_ELECTROWEAK_PROJECTION_BOUNDARY"
)

type InheritedAudit struct {
	Gate562QuaternionicBridge         bool
	Gate562MomentMap                  bool
	Gate562StabilizerOrbitSplit       bool
	Gate562PhysicalDynamicsFirewalled bool
	Gate562WSpatialTransferBlocked    bool
	Verdict                           string
}

type FiniteOneFormScalarLaneAudit struct {
	OneFormsFormula               string
	FluctuatedDiracFormula        string
	ScalarCarrier                 string
	Algebra                       string
	HiggsDoubletRecovered         bool
	ComplexDoublets               int
	RealScalarDimension           int
	WeakRepresentation            string
	ColorRepresentation           string
	StructuralDphiSocketFound     bool
	ScalarSU2RepresentationClosed bool
	NumericalYukawaFree           bool
	HiggsPotentialDerived         bool
	HeatKernelProjectionAvailable bool
	Verdict                       string
}

type QuaternionicActionAudit struct {
	ImHSocketAvailable           bool
	HphiDoubletModuleAvailable   bool
	StructuralActionAvailable    bool
	PairingFormula               string
	MomentPairingAvailable       bool
	AvailableLayer               string
	CouplingNormalizationDerived bool
	Verdict                      string
}

type CurvatureKineticProjectionAudit struct {
	StructuralDphiSocketFound             bool
	ProductActionContainsDphiSquared      bool
	SymbolicKineticProjectionReadOff      bool
	KineticCoefficientSymbol              string
	NativeScalarKineticCoefficientDerived bool
	NativeCanonicalScalarMetricDerived    bool
	ElectroweakCurvatureCarrierTyped      bool
	ElectroweakQuadraticFamilyTyped       bool
	NativeElectroweakCurvatureAction      bool
	FullSecondVariationComputed           bool
	GaugeHessianActionSelected            bool
	PhysicalGaugeCouplingsDerived         bool
	MomentMapTermFoundInFiniteCurvature   bool
	MomentMapTermFoundInKineticProjection bool
	MomentMapLayer                        string
	Verdict                               string
}

type MomentMapAppearanceAudit struct {
	PhiPhiDaggerIdentityAvailable    bool
	MuSigmaExpressionAvailable       bool
	PairingMuXAvailable              bool
	AppearsInFiniteOneForm           bool
	AppearsInCurvature               bool
	AppearsInScalarKineticProjection bool
	ExactTheoremLayer                string
	Verdict                          string
}

type StabilizerOrbitProjectionAudit struct {
	NonzeroMuSplitAvailable               bool
	Split                                 string
	RecognizedAtRepresentationLevel       bool
	CurvatureProjectionDistinguishesParts bool
	StabilizerUsedForPhotonDirection      bool
	OrbitUsedForWZDirections              bool
	Verdict                               string
}

type U1MixingFirewallAudit struct {
	AbelianSocketPresent                  bool
	AbelianNullDirectionDiagnostic        bool
	U1CompletionCoefficientSelected       bool
	HyperchargeAbsoluteNormalizationFixed bool
	WeakMixingAngleDerived                bool
	PhotonDirectionDerived                bool
	PhysicalElectroweakMixingNative       bool
	Verdict                               string
}

type KineticMassFirewallAudit struct {
	SymbolicDphiSquaredChannel     bool
	NativeScalarKineticCoefficient bool
	NativeGaugeKineticHessian      bool
	ScalarVacuumOrientationDerived bool
	HiggsVEVDerived                bool
	PhysicalWZMassMatrixDerived    bool
	GaugeCouplingsDerived          bool
	HiggsPotentialDerived          bool
	Verdict                        string
}

type FlavorFirewallAudit struct {
	YukawaEigenvaluesDerived       bool
	GenerationHierarchyDerived     bool
	CKMPMNSDerived                 bool
	ObservedFlavorImported         bool
	Q4PromotedToHiggsFlavor        bool
	TauEtaPromotedToSpectrum       bool
	WSpatialWeakPlaneRouteReopened bool
	PauliRouteSeparateFromWSpatial bool
	Verdict                        string
}

type FinalVerdict struct {
	FiniteOneFormContainsScalarDoublet     bool
	ImHActsStructurallyOnHphi              bool
	MomentMapInsideCurvatureOrKineticData  bool
	MomentMapRepresentationBookkeepingOnly bool
	NonzeroMuCurvatureLevelSplit           bool
	ElectroweakU1MixingDerived             bool
	KineticNormalizationAndMassDynamics    bool
	FlavorDataDerived                      bool
	MissingNextTheorem                     string
	Verdict                                string
}

type Analysis struct {
	Inherited    InheritedAudit
	OneForm      FiniteOneFormScalarLaneAudit
	Quaternionic QuaternionicActionAudit
	Kinetic      CurvatureKineticProjectionAudit
	Moment       MomentMapAppearanceAudit
	Orbit        StabilizerOrbitProjectionAudit
	U1           U1MixingFirewallAudit
	Mass         KineticMassFirewallAudit
	Flavor       FlavorFirewallAudit
	Final        FinalVerdict
	Truth        string
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
	g562, err := gate562.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate 562 scalar/quaternionic moment audit: %w", err)
	}
	g499, err := gate499.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate 499 Dphi provenance audit: %w", err)
	}
	g500, err := gate500.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate 500 product spectral-action kinetic projection audit: %w", err)
	}
	g493, err := gate493.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate 493 electroweak curvature action audit: %w", err)
	}

	a := Analysis{}
	a.Inherited = auditInherited(g562)
	a.OneForm = auditFiniteOneForm(g499)
	a.Quaternionic = auditQuaternionicAction(g562, a.OneForm)
	a.Kinetic = auditCurvatureKinetic(g499, g500, g493)
	a.Moment = auditMomentAppearance(g562, a.Quaternionic, a.Kinetic)
	a.Orbit = auditOrbitProjection(g562, a.Kinetic)
	a.U1 = auditU1Mixing(g493, g499)
	a.Mass = auditKineticMass(g500, g493)
	a.Flavor = auditFlavorFirewall(g500)
	a.Final = auditFinal(a)
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func auditInherited(g gate562.Analysis) InheritedAudit {
	return InheritedAudit{
		Gate562QuaternionicBridge:         g.Final.PauliTripletEquivalentToImH,
		Gate562MomentMap:                  g.Final.HopfMomentQuaternionicMomentMap,
		Gate562StabilizerOrbitSplit:       g.Final.NonzeroMuQuaternionicThreeSplit,
		Gate562PhysicalDynamicsFirewalled: !g.Final.PhysicalElectroweakDynamicsDerived,
		Gate562WSpatialTransferBlocked:    !g.Final.LawfulTransferToWSpatialOrGeneration,
		Verdict:                           join(StatusGate562Inherited, "Gate 563 inherits the sealed scalar/quaternionic moment map and its physical dynamics firewall"),
	}
}

func auditFiniteOneForm(g gate499.Analysis) FiniteOneFormScalarLaneAudit {
	return FiniteOneFormScalarLaneAudit{
		OneFormsFormula:               g.InnerFluctuation.FluctuatedDiracFormula,
		FluctuatedDiracFormula:        "D_A = D_F + A + JAJ^{-1}",
		ScalarCarrier:                 "H_phi ~= R^4 ~= C^2, one complex scalar SU(2)/H doublet",
		Algebra:                       g.InnerFluctuation.Algebra,
		HiggsDoubletRecovered:         g.InnerFluctuation.HiggsDoubletRecovered,
		ComplexDoublets:               g.InnerFluctuation.ComplexDoublets,
		RealScalarDimension:           g.InnerFluctuation.RealScalarDimension,
		WeakRepresentation:            g.InnerFluctuation.WeakRepresentation,
		ColorRepresentation:           g.InnerFluctuation.ColorRepresentation,
		StructuralDphiSocketFound:     g.Dphi.StructuralDphiSocketFound,
		ScalarSU2RepresentationClosed: g.Dphi.ScalarSU2RepresentationProvenanceClosed,
		NumericalYukawaFree:           g.InnerFluctuation.NumericalYukawaFree,
		HiggsPotentialDerived:         !g.InnerFluctuation.HiggsPotentialNotDerived,
		HeatKernelProjectionAvailable: !g.InnerFluctuation.HeatKernelProjectionMissing,
		Verdict:                       join(StatusFiniteOneFormScalarLaneRecovered, "The finite one-form lane recovers one complex scalar SU(2)/H doublet structurally, without Higgs-potential or heat-kernel dynamics"),
	}
}

func auditQuaternionicAction(g gate562.Analysis, one FiniteOneFormScalarLaneAudit) QuaternionicActionAudit {
	available := g.Final.ImHNativeOrientedMetricThreeSpace && one.HiggsDoubletRecovered && one.StructuralDphiSocketFound && g.Final.HPhiWeakDoubletModule
	return QuaternionicActionAudit{
		ImHSocketAvailable:           g.Final.ImHNativeOrientedMetricThreeSpace,
		HphiDoubletModuleAvailable:   g.Final.HPhiWeakDoubletModule && one.HiggsDoubletRecovered,
		StructuralActionAvailable:    available,
		PairingFormula:               "mu_X(phi)=phi^dagger X_H phi, with X_H a Hermitian Pauli representative of X in Im(H)",
		MomentPairingAvailable:       available,
		AvailableLayer:               "sealed scalar/quaternionic representation bookkeeping and finite one-form scalar-doublet provenance",
		CouplingNormalizationDerived: false,
		Verdict:                      join(StatusQuaternionicActionStructural, StatusMomentPairingBookkeeping, "Im(H) acts structurally on H_phi through the scalar doublet representation; the moment pairing is available before physical coupling normalization"),
	}
}

func auditCurvatureKinetic(g499 gate499.Analysis, g500 gate500.Analysis, g493 gate493.Analysis) CurvatureKineticProjectionAudit {
	return CurvatureKineticProjectionAudit{
		StructuralDphiSocketFound:             g499.Dphi.StructuralDphiSocketFound,
		ProductActionContainsDphiSquared:      g500.KineticProjection.ProductActionContainsDphiSquared,
		SymbolicKineticProjectionReadOff:      g500.KineticProjection.SymbolicKineticProjectionReadOff,
		KineticCoefficientSymbol:              g500.KineticProjection.CoefficientSymbol,
		NativeScalarKineticCoefficientDerived: g500.Boundary.NativeScalarKineticCoefficientDerived,
		NativeCanonicalScalarMetricDerived:    g500.Boundary.NativeCanonicalScalarMetricDerived,
		ElectroweakCurvatureCarrierTyped:      g493.Curvature.CurvatureCarrierTyped && g493.Curvature.FullFieldStrengthTyped,
		ElectroweakQuadraticFamilyTyped:       g493.Quadratic.FullQuadraticFamilyTyped,
		NativeElectroweakCurvatureAction:      g493.Boundary.NativeEWCurvatureActionDerived,
		FullSecondVariationComputed:           g493.Boundary.SecondVariationComputed,
		GaugeHessianActionSelected:            g493.Boundary.GaugeHessianActionSelected,
		PhysicalGaugeCouplingsDerived:         g493.Boundary.PhysicalGaugeCouplingsDerived,
		MomentMapTermFoundInFiniteCurvature:   false,
		MomentMapTermFoundInKineticProjection: false,
		MomentMapLayer:                        "the moment map is certified in the scalar/quaternionic representation layer; existing curvature/kinetic ledgers contain structural D_phi and symbolic |D_phi phi|^2, not an explicit mu_a sigma_a projection term",
		Verdict:                               join(StatusProductDphiSquaredSymbolic, StatusCurvatureSocketAvailable, StatusMomentNotInCurvatureProjection),
	}
}

func auditMomentAppearance(g gate562.Analysis, q QuaternionicActionAudit, k CurvatureKineticProjectionAudit) MomentMapAppearanceAudit {
	return MomentMapAppearanceAudit{
		PhiPhiDaggerIdentityAvailable:    g.Final.HopfMomentQuaternionicMomentMap,
		MuSigmaExpressionAvailable:       g.Final.HopfMomentQuaternionicMomentMap,
		PairingMuXAvailable:              q.MomentPairingAvailable,
		AppearsInFiniteOneForm:           false,
		AppearsInCurvature:               k.MomentMapTermFoundInFiniteCurvature,
		AppearsInScalarKineticProjection: k.MomentMapTermFoundInKineticProjection,
		ExactTheoremLayer:                "Gate 560/562 scalar-quaternionic representation bookkeeping; not finite curvature, not scalar kinetic projection, not spectral-action coefficient closure",
		Verdict:                          join(StatusMomentPairingBookkeeping, StatusMomentNotInCurvatureProjection),
	}
}

func auditOrbitProjection(g gate562.Analysis, k CurvatureKineticProjectionAudit) StabilizerOrbitProjectionAudit {
	return StabilizerOrbitProjectionAudit{
		NonzeroMuSplitAvailable:               g.Final.NonzeroMuQuaternionicThreeSplit,
		Split:                                 "Im(H)=R mu ⊕ mu^perp for mu != 0",
		RecognizedAtRepresentationLevel:       g.Final.NonzeroMuQuaternionicThreeSplit,
		CurvatureProjectionDistinguishesParts: false,
		StabilizerUsedForPhotonDirection:      false,
		OrbitUsedForWZDirections:              false,
		Verdict:                               join(StatusOrbitSplitRepresentationOnly, StatusNoCurvatureStabilizerProjection, "The split is scalar/quaternionic orbit geometry; no finite curvature projection distinguishes stabilizer/orbit components"),
	}
}

func auditU1Mixing(g493 gate493.Analysis, g499 gate499.Analysis) U1MixingFirewallAudit {
	return U1MixingFirewallAudit{
		AbelianSocketPresent:                  true,
		AbelianNullDirectionDiagnostic:        g493.Curvature.AbelianNullDirectionFound,
		U1CompletionCoefficientSelected:       g493.Boundary.AbelianCoefficientSelected,
		HyperchargeAbsoluteNormalizationFixed: !g499.InnerFluctuation.HeatKernelProjectionMissing && !g499.Firewall.NativeKappaWritten,
		WeakMixingAngleDerived:                g493.Boundary.WeakMixingAngleDerived,
		PhotonDirectionDerived:                false,
		PhysicalElectroweakMixingNative:       false,
		Verdict:                               join(StatusNoNativeU1PhotonDirection, "The abelian/hypercharge socket and null-direction diagnostics exist, but no native U(1) mixing or physical photon direction is selected"),
	}
}

func auditKineticMass(g500 gate500.Analysis, g493 gate493.Analysis) KineticMassFirewallAudit {
	return KineticMassFirewallAudit{
		SymbolicDphiSquaredChannel:     g500.KineticProjection.SymbolicKineticProjectionReadOff,
		NativeScalarKineticCoefficient: g500.Boundary.NativeScalarKineticCoefficientDerived,
		NativeGaugeKineticHessian:      g493.Boundary.GaugeHessianActionSelected,
		ScalarVacuumOrientationDerived: g500.Boundary.NativeVacuumOrientationDerived,
		HiggsVEVDerived:                false,
		PhysicalWZMassMatrixDerived:    g493.Boundary.PhysicalWZMassMatrixDerived || g500.Boundary.NativeWZMassMatrixDerived,
		GaugeCouplingsDerived:          g493.Boundary.PhysicalGaugeCouplingsDerived,
		HiggsPotentialDerived:          false,
		Verdict:                        join(StatusNoKineticNormalization, "The product action gives only a symbolic D_phi squared channel; kinetic coefficient, scalar metric, vacuum, gauge Hessian, couplings, and W/Z mass dynamics remain unclosed"),
	}
}

func auditFlavorFirewall(g500 gate500.Analysis) FlavorFirewallAudit {
	return FlavorFirewallAudit{
		YukawaEigenvaluesDerived:       false,
		GenerationHierarchyDerived:     false,
		CKMPMNSDerived:                 false,
		ObservedFlavorImported:         false,
		Q4PromotedToHiggsFlavor:        false,
		TauEtaPromotedToSpectrum:       false,
		WSpatialWeakPlaneRouteReopened: false,
		PauliRouteSeparateFromWSpatial: true,
		Verdict:                        join(StatusNoFlavorData, StatusPreviousFirewallsPreserved, "q4 remains contact-only, tau_eta remains Sigma_3-axis trace shadow, and the W_spatial weak-plane route remains blocked"),
	}
}

func auditFinal(a Analysis) FinalVerdict {
	return FinalVerdict{
		FiniteOneFormContainsScalarDoublet:     a.OneForm.HiggsDoubletRecovered && a.OneForm.ComplexDoublets == 1 && a.OneForm.RealScalarDimension == 4,
		ImHActsStructurallyOnHphi:              a.Quaternionic.StructuralActionAvailable && a.Quaternionic.MomentPairingAvailable,
		MomentMapInsideCurvatureOrKineticData:  a.Kinetic.MomentMapTermFoundInFiniteCurvature || a.Kinetic.MomentMapTermFoundInKineticProjection,
		MomentMapRepresentationBookkeepingOnly: a.Moment.PhiPhiDaggerIdentityAvailable && !a.Moment.AppearsInCurvature && !a.Moment.AppearsInScalarKineticProjection,
		NonzeroMuCurvatureLevelSplit:           a.Orbit.CurvatureProjectionDistinguishesParts,
		ElectroweakU1MixingDerived:             a.U1.PhysicalElectroweakMixingNative || a.U1.PhotonDirectionDerived,
		KineticNormalizationAndMassDynamics:    a.Mass.NativeScalarKineticCoefficient && a.Mass.NativeGaugeKineticHessian && a.Mass.PhysicalWZMassMatrixDerived,
		FlavorDataDerived:                      a.Flavor.YukawaEigenvaluesDerived || a.Flavor.CKMPMNSDerived || a.Flavor.ObservedFlavorImported,
		MissingNextTheorem:                     "Gate 564 should construct or obstruct an explicit scalar/quaternionic moment insertion into a finite curvature or product kinetic projection, including native U(1) mixing, scalar/gauge kinetic normalization, and zero observed-flavor input.",
		Verdict:                                join(StatusFiniteOneFormScalarLaneRecovered, StatusQuaternionicActionStructural, StatusMomentNotInCurvatureProjection, StatusNoNativeU1PhotonDirection, StatusNoKineticNormalization, StatusFirewallPreserved),
	}
}

func Statuses() []string {
	return []string{
		StatusGate562Inherited,
		StatusFiniteOneFormScalarLaneRecovered,
		StatusQuaternionicActionStructural,
		StatusMomentPairingBookkeeping,
		StatusProductDphiSquaredSymbolic,
		StatusCurvatureSocketAvailable,
		StatusMomentNotInCurvatureProjection,
		StatusOrbitSplitRepresentationOnly,
		StatusNoCurvatureStabilizerProjection,
		StatusNoNativeU1PhotonDirection,
		StatusNoKineticNormalization,
		StatusNoFlavorData,
		StatusPreviousFirewallsPreserved,
		StatusFirewallPreserved,
	}
}

func truth(a Analysis) string {
	return "Gate 563 keeps the scalar/quaternionic bridge honest. The finite one-form lane contains one complex SU(2)/H scalar doublet, and Im(H) acts structurally on H_phi with the Gate 560/562 moment pairing. The product spectral action also contains a symbolic D_phi squared channel. But the current project does not yet contain an explicit native mu_a sigma_a curvature projection, a curvature-level stabilizer/orbit split, native U(1) mixing or photon direction, kinetic normalization, physical W/Z dynamics, Higgs potential coefficients, or flavor data. The Pauli/quaternionic moment remains a structural scalar/H socket, not a physical electroweak mass theorem."
}

func validate(a Analysis) error {
	if !a.Inherited.Gate562QuaternionicBridge || !a.Inherited.Gate562MomentMap || !a.Inherited.Gate562PhysicalDynamicsFirewalled || !a.Inherited.Gate562WSpatialTransferBlocked {
		return fmt.Errorf("Gate562 inheritance failed: %s", FormatInherited(a.Inherited))
	}
	if !a.OneForm.HiggsDoubletRecovered || a.OneForm.ComplexDoublets != 1 || a.OneForm.RealScalarDimension != 4 || !a.OneForm.StructuralDphiSocketFound || !a.OneForm.ScalarSU2RepresentationClosed || a.OneForm.HiggsPotentialDerived || a.OneForm.HeatKernelProjectionAvailable {
		return fmt.Errorf("finite one-form lane failed: %s", FormatOneForm(a.OneForm))
	}
	if !a.Quaternionic.ImHSocketAvailable || !a.Quaternionic.HphiDoubletModuleAvailable || !a.Quaternionic.StructuralActionAvailable || !a.Quaternionic.MomentPairingAvailable || a.Quaternionic.CouplingNormalizationDerived {
		return fmt.Errorf("quaternionic action failed: %s", FormatQuaternionic(a.Quaternionic))
	}
	if !a.Kinetic.StructuralDphiSocketFound || !a.Kinetic.ProductActionContainsDphiSquared || !a.Kinetic.SymbolicKineticProjectionReadOff || a.Kinetic.NativeScalarKineticCoefficientDerived || a.Kinetic.NativeCanonicalScalarMetricDerived || !a.Kinetic.ElectroweakCurvatureCarrierTyped || !a.Kinetic.ElectroweakQuadraticFamilyTyped || a.Kinetic.NativeElectroweakCurvatureAction || a.Kinetic.FullSecondVariationComputed || a.Kinetic.GaugeHessianActionSelected || a.Kinetic.PhysicalGaugeCouplingsDerived || a.Kinetic.MomentMapTermFoundInFiniteCurvature || a.Kinetic.MomentMapTermFoundInKineticProjection {
		return fmt.Errorf("curvature/kinetic audit failed: %s", FormatKinetic(a.Kinetic))
	}
	if !a.Moment.PhiPhiDaggerIdentityAvailable || !a.Moment.MuSigmaExpressionAvailable || !a.Moment.PairingMuXAvailable || a.Moment.AppearsInCurvature || a.Moment.AppearsInScalarKineticProjection {
		return fmt.Errorf("moment appearance failed: %s", FormatMoment(a.Moment))
	}
	if !a.Orbit.NonzeroMuSplitAvailable || !a.Orbit.RecognizedAtRepresentationLevel || a.Orbit.CurvatureProjectionDistinguishesParts || a.Orbit.StabilizerUsedForPhotonDirection || a.Orbit.OrbitUsedForWZDirections {
		return fmt.Errorf("orbit projection failed: %s", FormatOrbit(a.Orbit))
	}
	if !a.U1.AbelianSocketPresent || !a.U1.AbelianNullDirectionDiagnostic || a.U1.U1CompletionCoefficientSelected || a.U1.WeakMixingAngleDerived || a.U1.PhotonDirectionDerived || a.U1.PhysicalElectroweakMixingNative {
		return fmt.Errorf("U1 mixing firewall failed: %s", FormatU1(a.U1))
	}
	if !a.Mass.SymbolicDphiSquaredChannel || a.Mass.NativeScalarKineticCoefficient || a.Mass.NativeGaugeKineticHessian || a.Mass.ScalarVacuumOrientationDerived || a.Mass.HiggsVEVDerived || a.Mass.PhysicalWZMassMatrixDerived || a.Mass.GaugeCouplingsDerived || a.Mass.HiggsPotentialDerived {
		return fmt.Errorf("kinetic/mass firewall failed: %s", FormatMass(a.Mass))
	}
	if a.Flavor.YukawaEigenvaluesDerived || a.Flavor.GenerationHierarchyDerived || a.Flavor.CKMPMNSDerived || a.Flavor.ObservedFlavorImported || a.Flavor.Q4PromotedToHiggsFlavor || a.Flavor.TauEtaPromotedToSpectrum || a.Flavor.WSpatialWeakPlaneRouteReopened || !a.Flavor.PauliRouteSeparateFromWSpatial {
		return fmt.Errorf("flavor firewall failed: %s", FormatFlavor(a.Flavor))
	}
	return nil
}

func join(parts ...string) string { return strings.Join(parts, "; ") }
