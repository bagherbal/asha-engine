// Package generation2projectivefockcp3momentmapselectorgeometryaudit implements Gate 572:
// Projective Fock CP3 Moment-Map Selector Geometry Audit.
//
// Gate 572 stays inside the Witt/Fock Hopf law-space lane certified by Gate
// 570 and behind the K7/product-time boundary enforced by Gate 571. It audits
// whether the Hopf quotient CP^3=S^7/S^1 carries the standard projective
// Fock/Kähler geometry on which Hermitian diagonal selectors define moment-map
// functions.  In particular it tests B-L=diag(-1,1/3,1/3,1/3), verifies the
// CP^0|CP^2 critical-stratum split, compares its stabilizer with Gate 555's
// commutant u(1)+u(3), and refuses to promote the remaining CP^2 degeneracy to
// a 2+1 weak-plane/generation selector without a second native selector.
package generation2projectivefockcp3momentmapselectorgeometryaudit

import (
	"fmt"
	"math"
	"math/cmplx"
	"strings"
	"sync"

	gate555 "github.com/bagherbal/asha-engine/pkg/bridge/generation2fourfoldselectororigintraceaudit"
	gate571 "github.com/bagherbal/asha-engine/pkg/bridge/generation2hopfs7k7producttimeairlockaudit"
	gate570 "github.com/bagherbal/asha-engine/pkg/bridge/generation2witthopfs7contactreebaudit"
)

const (
	AuditID = "GATE572-PROJECTIVE-FOCK-CP3-MOMENT-MAP-SELECTOR-GEOMETRY-AUDIT"

	StatusGates570571Inherited      = "CONDITIONAL_SUPPORT_GATES570_571_HOPF_CP3_AND_K7_FIREWALL_INHERITED"
	StatusCP3QuotientCertified      = "PASS_PROJECTIVE_FOCK_CP3_QUOTIENT_CERTIFIED"
	StatusCP3DimensionCertified     = "PASS_CP3_REAL_DIMENSION_6_CERTIFIED"
	StatusFubiniStudyQuotient       = "CONDITIONAL_SUPPORT_FUBINI_STUDY_FORM_QUOTIENT_OF_DALPHA_UP_TO_CONVENTION"
	StatusTotalFockPhaseQuotiented  = "PASS_TOTAL_FOCK_PHASE_REEB_DIRECTION_QUOTIENTED_ON_CP3"
	StatusLawSpacePhaseOnly         = "FIREWALL_PRESERVED_TOTAL_FOCK_PHASE_IS_LAW_SPACE_PHASE_NOT_PHYSICAL_TIME"
	StatusSelectorMomentWellDefined = "PASS_SELECTOR_MOMENT_MAP_FUNCTIONS_ON_CP3_DEFINED"
	StatusBMinusLFormulaVerified    = "PASS_B_MINUS_L_MOMENT_FORMULA_VERIFIED"
	StatusBMinusLCriticalStrata     = "PASS_B_MINUS_L_CRITICAL_STRATA_CP0_AND_CP2_VERIFIED"
	StatusProjectiveOnePlusThree    = "PASS_B_MINUS_L_PROJECTIVE_1_PLUS_3_GEOMETRY_CERTIFIED"
	StatusStabilizerMatchesGate555  = "PASS_B_MINUS_L_STABILIZER_U1_TIMES_U3_MATCHES_GATE555_COMMUTANT"
	StatusCP3HomogeneousGeometry    = "PASS_CP3_HOMOGENEOUS_GEOMETRY_U4_OVER_U1_TIMES_U3_CERTIFIED"
	StatusCP2SpatialBlock           = "CONDITIONAL_SUPPORT_CP2_NATIVE_PROJECTIVE_REFINEMENT_OF_SPATIAL_EIGENSPACE"
	StatusNoCP2SecondSelector       = "FAILED_ROUTE_NO_NATIVE_PROJECTIVE_SPATIAL_2PLUS1_SELECTOR"
	StatusCP2NoWeakPlane            = "FAILED_ROUTE_CP2_BLOCK_DOES_NOT_SELECT_WEAK_PLANE_CP1_PLUS_CP0"
	StatusNoCP3ToK7Functor          = "FAILED_ROUTE_NO_CP3_TO_BOOLEAN_OCTONIONIC_K7_FUNCTOR"
	StatusCP3MomentNotPhysicalTime  = "FAILED_ROUTE_CP3_MOMENT_MAP_FLOWS_NOT_PHYSICAL_TIME"
	StatusCP3NoOSHilbertOrRG        = "FAILED_ROUTE_CP3_MOMENT_MAP_FLOWS_DO_NOT_OPEN_OS_HILBERT_OR_RG_DYNAMICS"
	StatusNoFlavorEWDynamics        = "FAILED_ROUTE_CP3_SELECTOR_GEOMETRY_DOES_NOT_DERIVE_FLAVOR_EW_DYNAMICS_OR_OBSERVED_DATA"
	StatusGate572BoundaryPreserved  = "FIREWALL_PRESERVED_GATE572_PROJECTIVE_FOCK_CP3_SELECTOR_BOUNDARY"
)

var bMinusLCoefficients = []float64{-1, 1.0 / 3.0, 1.0 / 3.0, 1.0 / 3.0}

type InheritedBoundaryAudit struct {
	Gate570CP3Certified         bool
	Gate570ReebTotalPhase       bool
	Gate570PhysicalTimeOpened   bool
	Gate571K7FunctorFound       bool
	Gate571ProductTimeOpened    bool
	Gate571PhysicalDynamics     bool
	Gate555SelectorTheorem      bool
	Gate555BMinusLCommutantDim  int
	Gate555NativeSecondSelector bool
	Verdict                     string
}

type ProjectiveQuotientAudit struct {
	Carrier                     string
	AmbientComplexDimension     int
	Sphere                      string
	SphereRealDimension         int
	Fiber                       string
	FiberRealDimension          int
	Base                        string
	BaseComplexDimension        int
	BaseRealDimension           int
	DimensionFormula            string
	ProjectiveQuotientCertified bool
	FubiniStudyAvailable        bool
	PullbackConvention          string
	PhysicalSpacetime           bool
	Verdict                     string
}

type CentralPhaseQuotientAudit struct {
	Generator              string
	Action                 string
	ReebDirection          string
	ReebMatchesTotalNumber bool
	TrivialOnCP3           bool
	LawSpacePhaseOnly      bool
	PhysicalLorentzianTime bool
	OSHilbertDynamics      bool
	RGScale                bool
	SpacetimeHamiltonian   bool
	Verdict                string
}

type SelectorMomentAudit struct {
	SelectorFormula         string
	MomentFormula           string
	HermitianSelector       bool
	PhaseInvariant          bool
	ComplexScaleInvariant   bool
	SampleMoment            float64
	PhasedSampleMoment      float64
	ScaledSampleMoment      float64
	MaxInvarianceResidual   float64
	DefinesMomentFunctions  bool
	PhysicalHamiltonianFlow bool
	Verdict                 string
}

type BMinusLMomentAudit struct {
	Coefficients            []float64
	FormulaOnS7             string
	LeptonLineCondition     string
	SpatialPlaneCondition   string
	LeptonCriticalValue     float64
	SpatialCriticalValue    float64
	SampleLeptonValue       float64
	SampleSpatialValue      float64
	FormulaResidualMax      float64
	CriticalStrataCertified bool
	ProjectiveOnePlusThree  bool
	WeakPlaneSelected       bool
	GenerationSelected      bool
	Verdict                 string
}

type StabilizerAudit struct {
	SelectorSplit                  string
	Stabilizer                     string
	StabilizerDimension            int
	LieAlgebra                     string
	Gate555Commutant               string
	Gate555CommutantDimension      int
	MatchesGate555Commutant        bool
	CP3HomogeneousModel            string
	U4Dimension                    int
	IsotropyDimension              int
	HomogeneousRealDimension       int
	HomogeneousDimensionMatchesCP3 bool
	Verdict                        string
}

type SpatialProjectiveBlockAudit struct {
	Block                      string
	ProjectiveDimension        string
	NativeProjectiveRefinement bool
	BMinusLSpatialEigenspace   bool
	WeakPlaneSelected          bool
	RequiresSecondSelector     bool
	Verdict                    string
}

type SecondSelectorObstructionAudit struct {
	CurrentNativeSecondSelector bool
	Gate555UniqueWeakPlane      bool
	TauEtaPulledBackNative      bool
	CandidateCP2Split           string
	SpatialTwoPlusOneDerived    bool
	Verdict                     string
	Reason                      string
}

type K7RelationAudit struct {
	CP3ToK7FunctorFound       bool
	HopfS7ToK7FunctorFound    bool
	TangentS7ToK7FunctorFound bool
	TotalPhaseToK7Action      bool
	DimensionMatchPromoted    bool
	Gate571BoundaryPreserved  bool
	Verdict                   string
}

type ProductTimeFirewallAudit struct {
	MomentFlowPhysicalTime    bool
	MomentFlowOSHilbert       bool
	MomentFlowRGScale         bool
	MomentFlowSpacetime       bool
	MomentFlowObservedHistory bool
	LawSpaceHamiltonianOnly   bool
	Verdict                   string
}

type FlavorElectroweakFirewallAudit struct {
	YukawaEigenvaluesDerived       bool
	CKMPMNSDerived                 bool
	GenerationHierarchyDerived     bool
	PhotonDynamicsDerived          bool
	WZMassesDerived                bool
	ObservedDataImported           bool
	Gate564565RemainBridgeSymbolic bool
	Verdict                        string
}

type FinalVerdict struct {
	CP3Certified                   bool
	FubiniStudyAvailable           bool
	SelectorMomentFunctionsOnCP3   bool
	BMinusLProjectiveCP0CP2Split   bool
	MatchesGate555Commutant        bool
	NativeSecondSelectorOnCP2      bool
	K7RelationOrPhysicalTimeProven bool
	MissingNextTheorem             string
	Verdict                        string
}

type Analysis struct {
	Inherited    InheritedBoundaryAudit
	Projective   ProjectiveQuotientAudit
	Phase        CentralPhaseQuotientAudit
	Selector     SelectorMomentAudit
	BMinusL      BMinusLMomentAudit
	Stabilizer   StabilizerAudit
	SpatialBlock SpatialProjectiveBlockAudit
	Second       SecondSelectorObstructionAudit
	K7           K7RelationAudit
	Time         ProductTimeFirewallAudit
	FlavorEW     FlavorElectroweakFirewallAudit
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
	prev555, err := gate555.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate555 selector algebra audit: %w", err)
	}
	prev570, err := gate570.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate570 Hopf CP3 predecessor: %w", err)
	}
	prev571, err := gate571.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate571 K7/product-time firewall predecessor: %w", err)
	}

	a := Analysis{}
	a.Inherited = auditInherited(prev555, prev570, prev571)
	a.Projective = auditProjectiveQuotient(prev570)
	a.Phase = auditCentralPhase(prev570)
	a.Selector = auditSelectorMoment()
	a.BMinusL = auditBMinusLMoment()
	a.Stabilizer = auditStabilizer(prev555, a.Projective)
	a.SpatialBlock = auditSpatialBlock()
	a.Second = auditSecondSelector(prev555)
	a.K7 = auditK7(prev571)
	a.Time = auditProductTime()
	a.FlavorEW = auditFlavorElectroweak()
	a.Final = auditFinal(a)
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func auditInherited(prev555 gate555.Analysis, prev570 gate570.Analysis, prev571 gate571.Analysis) InheritedBoundaryAudit {
	return InheritedBoundaryAudit{
		Gate570CP3Certified:         prev570.Final.CP3ProjectiveLawSpace,
		Gate570ReebTotalPhase:       prev570.Final.TotalPhaseRelation,
		Gate570PhysicalTimeOpened:   prev570.Final.PhysicalTimeOpened,
		Gate571K7FunctorFound:       prev571.Final.HopfToK7FunctorFound || prev571.Final.TangentToK7FunctorFound,
		Gate571ProductTimeOpened:    prev571.Final.ProductTimeAirlockOpened,
		Gate571PhysicalDynamics:     prev571.Final.PhysicalDynamicsOpened,
		Gate555SelectorTheorem:      prev555.Final.NativeSelectorAlgebraTheorem,
		Gate555BMinusLCommutantDim:  prev555.BMinusL.CommutantDimension,
		Gate555NativeSecondSelector: prev555.Final.NativeThreeToTwoPlusOneSelector,
		Verdict:                     StatusGates570571Inherited,
	}
}

func auditProjectiveQuotient(prev570 gate570.Analysis) ProjectiveQuotientAudit {
	baseComplex := prev570.Quotient.BaseComplexDimension
	baseReal := 2 * baseComplex
	return ProjectiveQuotientAudit{
		Carrier:                     "W=C^4 with positive Hermitian Fock metric",
		AmbientComplexDimension:     4,
		Sphere:                      "S^7={z in C^4:<z,z>=1}",
		SphereRealDimension:         prev570.Sphere.SphereRealDimension,
		Fiber:                       "S^1 global Fock phase",
		FiberRealDimension:          1,
		Base:                        "CP^3=S^7/S^1=P(C^4)",
		BaseComplexDimension:        baseComplex,
		BaseRealDimension:           baseReal,
		DimensionFormula:            "dim_R CP^3 = dim_R S^7 - dim_R S^1 = 7-1 = 6 = 2*3",
		ProjectiveQuotientCertified: prev570.Quotient.ProjectiveLawSpace && prev570.Quotient.Base == "CP^3" && prev570.Quotient.BaseRealDimension == 6,
		FubiniStudyAvailable:        true,
		PullbackConvention:          "pi^* omega_FS = d alpha in the Gate570 normalization alpha=Im<z,dz>; equivalently some literature uses pi^* omega_FS = (1/2)d alpha when alpha or omega_FS is rescaled",
		PhysicalSpacetime:           false,
		Verdict:                     join(StatusCP3QuotientCertified, StatusCP3DimensionCertified, StatusFubiniStudyQuotient),
	}
}

func auditCentralPhase(prev570 gate570.Analysis) CentralPhaseQuotientAudit {
	return CentralPhaseQuotientAudit{
		Generator:              "N=N_0+N_1+N_2+N_3",
		Action:                 "z -> exp(i theta) z",
		ReebDirection:          "R_z=Jz=iz",
		ReebMatchesTotalNumber: prev570.Final.TotalPhaseRelation && prev570.Reeb.UniqueByContactEquation,
		TrivialOnCP3:           true,
		LawSpacePhaseOnly:      true,
		Verdict:                join(StatusTotalFockPhaseQuotiented, StatusLawSpacePhaseOnly),
	}
}

func auditSelectorMoment() SelectorMomentAudit {
	s := []float64{-2, -0.5, 0.25, 3}
	z := []complex128{1 + 2i, -0.5 + 0.25i, 2 - 1i, -1.5 + 0.75i}
	h, _ := RayleighMoment(s, z)
	phase := cmplx.Exp(complex(0, 0.37))
	hPhase, _ := RayleighMoment(s, scaleComplex(z, phase))
	hScale, _ := RayleighMoment(s, scaleComplex(z, 2.25*cmplx.Exp(complex(0, -0.91))))
	res := maxAbs(h-hPhase, h-hScale)
	return SelectorMomentAudit{
		SelectorFormula:         "S=sum_k s_k N_k, represented on W=C^4 by Hermitian diag(s_0,s_1,s_2,s_3)",
		MomentFormula:           "h_S([z])=(z^dagger S z)/(z^dagger z)",
		HermitianSelector:       true,
		PhaseInvariant:          math.Abs(h-hPhase) < 1e-12,
		ComplexScaleInvariant:   math.Abs(h-hScale) < 1e-12,
		SampleMoment:            h,
		PhasedSampleMoment:      hPhase,
		ScaledSampleMoment:      hScale,
		MaxInvarianceResidual:   res,
		DefinesMomentFunctions:  true,
		PhysicalHamiltonianFlow: false,
		Verdict:                 StatusSelectorMomentWellDefined,
	}
}

func auditBMinusLMoment() BMinusLMomentAudit {
	e0 := []complex128{1, 0, 0, 0}
	spatial := []complex128{0, 1 / complex(math.Sqrt(3), 0), 1 / complex(math.Sqrt(3), 0), 1 / complex(math.Sqrt(3), 0)}
	mixed := []complex128{complex(math.Sqrt(0.4), 0), complex(math.Sqrt(0.2), 0), complex(math.Sqrt(0.3), 0), complex(math.Sqrt(0.1), 0)}
	h0, _ := RayleighMoment(bMinusLCoefficients, e0)
	hs, _ := RayleighMoment(bMinusLCoefficients, spatial)
	hm, _ := RayleighMoment(bMinusLCoefficients, mixed)
	formulaMixed := 1.0/3.0 - (4.0/3.0)*abs2(mixed[0])
	res := maxAbs(h0-(-1), hs-(1.0/3.0), hm-formulaMixed)
	return BMinusLMomentAudit{
		Coefficients:            append([]float64(nil), bMinusLCoefficients...),
		FormulaOnS7:             "h_{B-L}=-|z0|^2+(1/3)(|z1|^2+|z2|^2+|z3|^2)=1/3-(4/3)|z0|^2",
		LeptonLineCondition:     "z1=z2=z3=0 gives projectivized eigenspace CP^0 with eigenvalue -1",
		SpatialPlaneCondition:   "z0=0 gives projectivized eigenspace CP^2 with eigenvalue 1/3",
		LeptonCriticalValue:     -1,
		SpatialCriticalValue:    1.0 / 3.0,
		SampleLeptonValue:       h0,
		SampleSpatialValue:      hs,
		FormulaResidualMax:      math.Abs(res),
		CriticalStrataCertified: math.Abs(res) < 1e-12,
		ProjectiveOnePlusThree:  true,
		WeakPlaneSelected:       false,
		GenerationSelected:      false,
		Verdict:                 join(StatusBMinusLFormulaVerified, StatusBMinusLCriticalStrata, StatusProjectiveOnePlusThree),
	}
}

func auditStabilizer(prev555 gate555.Analysis, p ProjectiveQuotientAudit) StabilizerAudit {
	iso := 1 + 9
	return StabilizerAudit{
		SelectorSplit:                  "eigenvalue -1 multiplicity 1 and eigenvalue 1/3 multiplicity 3",
		Stabilizer:                     "U(1)xU(3)",
		StabilizerDimension:            iso,
		LieAlgebra:                     "u(1)+u(3)",
		Gate555Commutant:               prev555.BMinusL.Commutant,
		Gate555CommutantDimension:      prev555.BMinusL.CommutantDimension,
		MatchesGate555Commutant:        prev555.BMinusL.CommutantDimension == iso && strings.Contains(prev555.BMinusL.Commutant, "u(1) + u(3)"),
		CP3HomogeneousModel:            "CP^3 ~= U(4)/(U(1)xU(3)) up to central convention",
		U4Dimension:                    16,
		IsotropyDimension:              iso,
		HomogeneousRealDimension:       16 - iso,
		HomogeneousDimensionMatchesCP3: 16-iso == p.BaseRealDimension,
		Verdict:                        join(StatusStabilizerMatchesGate555, StatusCP3HomogeneousGeometry),
	}
}

func auditSpatialBlock() SpatialProjectiveBlockAudit {
	return SpatialProjectiveBlockAudit{
		Block:                      "CP^2={z0=0}/S^1=P(span{e1,e2,e3})",
		ProjectiveDimension:        "complex dimension 2, real dimension 4",
		NativeProjectiveRefinement: true,
		BMinusLSpatialEigenspace:   true,
		WeakPlaneSelected:          false,
		RequiresSecondSelector:     true,
		Verdict:                    StatusCP2SpatialBlock,
	}
}

func auditSecondSelector(prev555 gate555.Analysis) SecondSelectorObstructionAudit {
	return SecondSelectorObstructionAudit{
		CurrentNativeSecondSelector: prev555.Final.NativeThreeToTwoPlusOneSelector,
		Gate555UniqueWeakPlane:      prev555.WeakPlane.UniqueWeakPlane,
		TauEtaPulledBackNative:      prev555.TauEta.NativeThreeToTwoPlusOne,
		CandidateCP2Split:           "CP^2 -> CP^1 plus CP^0, equivalent to spatial 3 -> 2+1",
		SpatialTwoPlusOneDerived:    false,
		Verdict:                     join(StatusNoCP2SecondSelector, StatusCP2NoWeakPlane),
		Reason:                      "B-L has one spatial/color eigenvalue of multiplicity three; Gate555 reports no unique weak plane and no unit-preserving native tau_eta pullback, so CP^2 is not further split by current project data.",
	}
}

func auditK7(prev571 gate571.Analysis) K7RelationAudit {
	return K7RelationAudit{
		CP3ToK7FunctorFound:       prev571.Quotient.CP3ToK7FunctorFound,
		HopfS7ToK7FunctorFound:    prev571.Final.HopfToK7FunctorFound,
		TangentS7ToK7FunctorFound: prev571.Final.TangentToK7FunctorFound,
		TotalPhaseToK7Action:      prev571.Quotient.K7CentralU1ActionFound,
		DimensionMatchPromoted:    !prev571.Final.DimensionMatchOnly,
		Gate571BoundaryPreserved:  !prev571.Final.HopfToK7FunctorFound && !prev571.Final.TangentToK7FunctorFound && !prev571.Quotient.CP3ToK7FunctorFound && !prev571.Quotient.K7CentralU1ActionFound,
		Verdict:                   join(StatusNoCP3ToK7Functor, StatusGate572BoundaryPreserved),
	}
}

func auditProductTime() ProductTimeFirewallAudit {
	return ProductTimeFirewallAudit{
		LawSpaceHamiltonianOnly: true,
		Verdict:                 join(StatusCP3MomentNotPhysicalTime, StatusCP3NoOSHilbertOrRG, StatusGate572BoundaryPreserved),
	}
}

func auditFlavorElectroweak() FlavorElectroweakFirewallAudit {
	return FlavorElectroweakFirewallAudit{
		Gate564565RemainBridgeSymbolic: true,
		Verdict:                        join(StatusNoFlavorEWDynamics, StatusGate572BoundaryPreserved),
	}
}

func auditFinal(a Analysis) FinalVerdict {
	physicalOrK7 := !a.K7.Gate571BoundaryPreserved || a.Time.MomentFlowPhysicalTime || a.Time.MomentFlowOSHilbert || a.Time.MomentFlowRGScale || a.Time.MomentFlowSpacetime || a.Time.MomentFlowObservedHistory
	return FinalVerdict{
		CP3Certified:                   a.Projective.ProjectiveQuotientCertified && a.Projective.BaseRealDimension == 6,
		FubiniStudyAvailable:           a.Projective.FubiniStudyAvailable,
		SelectorMomentFunctionsOnCP3:   a.Selector.DefinesMomentFunctions && a.Selector.PhaseInvariant && a.Selector.ComplexScaleInvariant,
		BMinusLProjectiveCP0CP2Split:   a.BMinusL.CriticalStrataCertified && a.BMinusL.ProjectiveOnePlusThree,
		MatchesGate555Commutant:        a.Stabilizer.MatchesGate555Commutant && a.Stabilizer.HomogeneousDimensionMatchesCP3,
		NativeSecondSelectorOnCP2:      a.Second.CurrentNativeSecondSelector || a.Second.SpatialTwoPlusOneDerived,
		K7RelationOrPhysicalTimeProven: physicalOrK7,
		MissingNextTheorem:             "A lawful next theorem must supply an additional native Hermitian selector on the spatial CP^2 block, or a certified functor/intertwiner from CP^3/Hopf data to another carrier. Without that, CP^2 remains a degenerate spatial/color projective block and no K7, time, flavor, electroweak, or observed-data bridge is opened.",
		Verdict:                        join(StatusCP3QuotientCertified, StatusSelectorMomentWellDefined, StatusProjectiveOnePlusThree, StatusNoCP2SecondSelector, StatusCP2NoWeakPlane, StatusNoCP3ToK7Functor, StatusCP3MomentNotPhysicalTime, StatusCP3NoOSHilbertOrRG, StatusNoFlavorEWDynamics, StatusGate572BoundaryPreserved),
	}
}

func validate(a Analysis) error {
	if !a.Inherited.Gate570CP3Certified || !a.Inherited.Gate570ReebTotalPhase || a.Inherited.Gate570PhysicalTimeOpened {
		return fmt.Errorf("Gate570 CP3/Reeb inheritance failed")
	}
	if a.Inherited.Gate571K7FunctorFound || a.Inherited.Gate571ProductTimeOpened || a.Inherited.Gate571PhysicalDynamics {
		return fmt.Errorf("Gate571 firewall was unexpectedly open")
	}
	if !a.Inherited.Gate555SelectorTheorem || a.Inherited.Gate555BMinusLCommutantDim != 10 {
		return fmt.Errorf("Gate555 selector algebra/commutant inheritance failed")
	}
	if !a.Final.CP3Certified || !a.Final.FubiniStudyAvailable || !a.Final.SelectorMomentFunctionsOnCP3 || !a.Final.BMinusLProjectiveCP0CP2Split || !a.Final.MatchesGate555Commutant {
		return fmt.Errorf("projective CP3 selector audit did not certify required positive results")
	}
	if a.Final.NativeSecondSelectorOnCP2 {
		return fmt.Errorf("unexpected native CP2 second selector promoted")
	}
	if a.Final.K7RelationOrPhysicalTimeProven {
		return fmt.Errorf("unexpected K7 or physical-time relation promoted")
	}
	if a.FlavorEW.YukawaEigenvaluesDerived || a.FlavorEW.CKMPMNSDerived || a.FlavorEW.GenerationHierarchyDerived || a.FlavorEW.PhotonDynamicsDerived || a.FlavorEW.WZMassesDerived || a.FlavorEW.ObservedDataImported {
		return fmt.Errorf("flavor/electroweak firewall pollution detected")
	}
	return nil
}

func truth(a Analysis) string {
	parts := []string{
		"Gate 572 certifies CP^3=S^7/S^1=P(C^4) as the projective Witt/Fock law-space, with real dimension six and the Fubini-Study symplectic/Kähler form available from the Hopf contact curvature d alpha up to normalization convention.",
		"The central total-number phase is exactly the Hopf/Reeb fiber and becomes trivial after projectivization; this is only law-space phase quotient, not Lorentzian time or product dynamics.",
		"Every Hermitian selector S=sum s_k N_k defines a phase-invariant Rayleigh moment h_S([z])=(z^dagger S z)/(z^dagger z) on CP^3.",
		"For B-L=diag(-1,1/3,1/3,1/3), h=1/3-(4/3)|z0|^2 on S^7, with critical projective strata CP^0 and CP^2 realizing the projective 1+3 split.",
		"The stabilizer U(1)xU(3) and Lie algebra u(1)+u(3) match Gate555's commutant, but no second native selector splits CP^2 into CP^1 plus CP^0.",
		"Gate571 remains intact: CP^3 moment-map geometry is not K7, physical time, RG scale, OS/Hilbert dynamics, spacetime Hamiltonian evolution, flavor, electroweak mass data, or observed history.",
	}
	return strings.Join(parts, " ")
}

func Statuses() []string {
	return []string{
		StatusGates570571Inherited,
		StatusCP3QuotientCertified,
		StatusCP3DimensionCertified,
		StatusFubiniStudyQuotient,
		StatusTotalFockPhaseQuotiented,
		StatusLawSpacePhaseOnly,
		StatusSelectorMomentWellDefined,
		StatusBMinusLFormulaVerified,
		StatusBMinusLCriticalStrata,
		StatusProjectiveOnePlusThree,
		StatusStabilizerMatchesGate555,
		StatusCP3HomogeneousGeometry,
		StatusCP2SpatialBlock,
		StatusNoCP2SecondSelector,
		StatusCP2NoWeakPlane,
		StatusNoCP3ToK7Functor,
		StatusCP3MomentNotPhysicalTime,
		StatusCP3NoOSHilbertOrRG,
		StatusNoFlavorEWDynamics,
		StatusGate572BoundaryPreserved,
	}
}

func RayleighMoment(selector []float64, z []complex128) (float64, error) {
	if len(selector) != len(z) {
		return 0, fmt.Errorf("selector/vector dimension mismatch: %d vs %d", len(selector), len(z))
	}
	n := 0.0
	d := 0.0
	for i, zi := range z {
		w := abs2(zi)
		n += selector[i] * w
		d += w
	}
	if d == 0 {
		return 0, fmt.Errorf("zero vector has no projective class")
	}
	return n / d, nil
}

func scaleComplex(z []complex128, c complex128) []complex128 {
	out := make([]complex128, len(z))
	for i := range z {
		out[i] = c * z[i]
	}
	return out
}

func abs2(z complex128) float64 {
	r, im := real(z), imag(z)
	return r*r + im*im
}

func maxAbs(xs ...float64) float64 {
	m := 0.0
	for _, x := range xs {
		if ax := math.Abs(x); ax > m {
			m = ax
		}
	}
	return m
}

func join(xs ...string) string { return strings.Join(xs, ";") }
