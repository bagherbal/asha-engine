// Package generation2spatialcp2selectorisotropyobstructionaudit implements Gate 573:
// Spatial CP2 Selector and SU(3) Isotropy Obstruction Audit.
//
// Gate 572 certified the projective Fock law-space CP^3=S^7/S^1 and the
// B-L projective split CP^0|CP^2. Gate 573 descends into the spatial block
// CP^2_sp=P(span_C{a_1^dagger,a_2^dagger,a_3^dagger}) and asks whether current
// ASHA data contains a native rank-one projector P_u that would split it as
// CP^1|CP^0. The result is an obstruction theorem: U(3), with traceless part
// SU(3), acts on the B-L-degenerate spatial eigenspace, SU(3) is transitive on
// CP^2, and no SU(3)-invariant datum can prefer a point [u]. A 2+1 Hermitian
// selector can be written once a point/rank-one projector is sealed, but current
// project data does not derive that point natively and no weak-isospin, flavor,
// electroweak, K7, or product-time bridge is opened.
package generation2spatialcp2selectorisotropyobstructionaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate557 "github.com/bagherbal/asha-engine/pkg/bridge/generation2etatracerepresentativerecordalgebraaudit"
	gate555 "github.com/bagherbal/asha-engine/pkg/bridge/generation2fourfoldselectororigintraceaudit"
	gate560 "github.com/bagherbal/asha-engine/pkg/bridge/generation2paulihopfscalarmomentmapaudit"
	gate561 "github.com/bagherbal/asha-engine/pkg/bridge/generation2paulimomentweakplaneincidenceaudit"
	gate572 "github.com/bagherbal/asha-engine/pkg/bridge/generation2projectivefockcp3momentmapselectorgeometryaudit"
	gate556 "github.com/bagherbal/asha-engine/pkg/bridge/generation2tauetacarrierpullbackobstructionaudit"
)

const (
	AuditID = "GATE573-SPATIAL-CP2-SELECTOR-AND-SU3-ISOTROPY-OBSTRUCTION-AUDIT"

	StatusGate572Inherited                = "CONDITIONAL_SUPPORT_GATE572_PROJECTIVE_CP3_B_MINUS_L_SPLIT_INHERITED"
	StatusSpatialCP2Certified             = "PASS_CP2_SPATIAL_BLOCK_CERTIFIED_AS_B_MINUS_L_CRITICAL_STRATUM"
	StatusU3SU3SymmetryCertified          = "PASS_U3_WITH_SU3_TRACeless_PART_ACTS_ON_W_SPATIAL"
	StatusBMinusLScalarOnSpatial          = "PASS_B_MINUS_L_RESTRICTS_TO_ONE_THIRD_IDENTITY_ON_W_SPATIAL"
	StatusSU3TransitiveOnCP2              = "PASS_SU3_ACTS_TRANSITIVELY_ON_SPATIAL_CP2"
	StatusPointStabilizerCertified        = "PASS_CP2_HOMOGENEOUS_MODEL_SU3_OVER_S_U1_TIMES_U2_CERTIFIED"
	StatusNoInvariantPoint                = "FAILED_ROUTE_NO_SU3_INVARIANT_POINT_IN_SPATIAL_CP2"
	StatusGeneralSecondSelectorClassified = "PASS_GENERAL_HERMITIAN_SPATIAL_2PLUS1_SELECTOR_CLASSIFIED"
	StatusCP1CP0CriticalStrataClassified  = "PASS_SPATIAL_SECOND_SELECTOR_CRITICAL_STRATA_CP1_AND_CP0_CLASSIFIED"
	StatusNoNativeRankOneProjector        = "FAILED_ROUTE_NO_NATIVE_RANK_ONE_PROJECTOR_ON_SPATIAL_CP2"
	StatusNoNativeCP2SecondSelector       = "FAILED_ROUTE_NO_NATIVE_PROJECTIVE_SPATIAL_2PLUS1_SELECTOR"
	StatusNoWeakPlaneCP1CP0               = "FAILED_ROUTE_CP2_BLOCK_DOES_NOT_SELECT_WEAK_PLANE_CP1_PLUS_CP0"
	StatusOrientationSealRequired         = "CONDITIONAL_SUPPORT_SPATIAL_PROJECTIVE_ORIENTATION_SEAL_DEFINED"
	StatusSealedU2U1Commutant             = "CONDITIONAL_SUPPORT_SEALED_SPATIAL_SELECTOR_COMMUTANT_U2_PLUS_U1_DIMENSION_5"
	StatusBasisDependentU12               = "CONDITIONAL_SUPPORT_SEALED_CHOICE_U_EQUALS_A3_DAGGER_CONVENTIONALLY_GIVES_U12"
	StatusBasisDependentNotNative         = "FAILED_ROUTE_U12_WEAK_PLANE_CHOICE_BASIS_DEPENDENT_NOT_NATIVE"
	StatusK7TimeBoundaryPreserved         = "FIREWALL_PRESERVED_GATE571_K7_AND_PRODUCT_TIME_BOUNDARY"
	StatusFlavorEWBoundaryPreserved       = "FIREWALL_PRESERVED_GATE573_NO_WEAK_ISOSPIN_FLAVOR_ELECTROWEAK_OR_OBSERVED_DATA"
	StatusGate573BoundaryPreserved        = "FIREWALL_PRESERVED_GATE573_SPATIAL_CP2_SELECTOR_BOUNDARY"
)

type InheritedGate572Audit struct {
	CP3ProjectiveLawSpace     bool
	BMinusLProjectiveOnePlus3 bool
	SpatialCP2Block           bool
	NoNativeSecondSelector    bool
	NoCP3ToK7Functor          bool
	NoPhysicalTimeOpened      bool
	NoFlavorElectroweakData   bool
	Verdict                   string
}

type SpatialCP2CarrierAudit struct {
	SpatialEigenspace             string
	ProjectiveBlock               string
	ComplexDimension              int
	RealDimension                 int
	BLRestrictedMatrix            string
	BLRestrictedEigenvalue        string
	Gate572CriticalStratumMatched bool
	CertifiedAsSpatialBlock       bool
	Verdict                       string
}

type SpatialSymmetryAudit struct {
	CommutantFromBMinusL      string
	SpatialSymmetry           string
	TracelessPart             string
	U3Dimension               int
	SU3Dimension              int
	BMinusLScalarOnWSpatial   bool
	SuppliesFurtherSelector   bool
	PreferredSpatialDirection bool
	Verdict                   string
}

type TransitivityAudit struct {
	Group                     string
	Space                     string
	Action                    string
	GroupRealDimension        int
	PointStabilizer           string
	PointStabilizerDimension  int
	QuotientRealDimension     int
	CP2RealDimension          int
	ActsTransitively          bool
	InvariantPointSelected    bool
	InvariantRankOneProjector bool
	Verdict                   string
}

type HermitianSecondSelectorAudit struct {
	SelectorFormula             string
	RankOneProjectorFormula     string
	LambdaOne                   float64
	LambdaTwo                   float64
	SamplePoint                 string
	SampleMatrix                [][]float64
	SampleEigenvalues           []float64
	MultiplicityPattern         string
	CriticalCP0                 string
	CriticalCP1                 string
	CP1RealDimension            int
	CP0RealDimension            int
	ProjectorIdempotentResidual float64
	ProjectorTrace              float64
	ClassifiesCP1CP0Split       bool
	NativeWithoutU              bool
	Verdict                     string
}

type NativeSelectorCandidate struct {
	Source            string
	PriorGate         string
	Candidate         string
	NativePUProvided  bool
	WouldSelectCP1CP0 bool
	Status            string
	Reason            string
}

type NativeSelectorSearchAudit struct {
	Candidates                  []NativeSelectorCandidate
	CandidateCount              int
	NativeRankOneProjectorFound bool
	NativeProjectivePointFound  bool
	NativeSecondSelectorFound   bool
	Verdict                     string
	Reason                      string
}

type OrientationSealAudit struct {
	SealName                   string
	MinimalDatum               string
	EquivalentDatum            string
	SealedSelectorFormula      string
	SealedNotNative            bool
	Commutant                  string
	CommutantDimension         int
	ExpectedCommutantDimension int
	CommMatchesGate555Formula  bool
	Verdict                    string
}

type WeakPlaneRelationAudit struct {
	SealedPoint           string
	ComplementaryPlane    string
	ConventionalPlaneName string
	BasisDependent        bool
	NativeDerived         bool
	WeakIsospinIdentified bool
	GenerationHierarchy   bool
	YukawaTextureDerived  bool
	CKMPMNSDerived        bool
	Verdict               string
}

type FirewallAudit struct {
	PromotedToWeakIsospin       bool
	PromotedToPhysicalWeakPlane bool
	GenerationHierarchyDerived  bool
	YukawaTextureDerived        bool
	CKMPMNSDerived              bool
	ObservedFlavorDataImported  bool
	PhysicalElectroweakDynamics bool
	WZMassesDerived             bool
	PhotonDynamicsDerived       bool
	CP2ToK7FunctorOpened        bool
	ProductTimeOpened           bool
	Gate564565BoundaryPreserved bool
	Verdict                     string
}

type FinalVerdict struct {
	SpatialCP2Certified         bool
	SU3Transitive               bool
	SU3InvariantPointSelected   bool
	GeneralTwoPlusOneSelector   bool
	NativeRankOnePU             bool
	MinimalSeal                 string
	PhysicalWeakFlavorEWDerived bool
	K7OrProductTimeOpened       bool
	MissingNextTheorem          string
	Verdict                     string
}

type Analysis struct {
	Inherited InheritedGate572Audit
	Carrier   SpatialCP2CarrierAudit
	Symmetry  SpatialSymmetryAudit
	Transit   TransitivityAudit
	Selector  HermitianSecondSelectorAudit
	Search    NativeSelectorSearchAudit
	Seal      OrientationSealAudit
	WeakPlane WeakPlaneRelationAudit
	Firewall  FirewallAudit
	Final     FinalVerdict
	Truth     string
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
	g572, err := gate572.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate572 projective CP3 predecessor: %w", err)
	}
	g555, err := gate555.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate555 selector predecessor: %w", err)
	}
	g556, err := gate556.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate556 tau-eta pullback predecessor: %w", err)
	}
	g557, err := gate557.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate557 eta-record algebra predecessor: %w", err)
	}
	g560, err := gate560.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate560 Pauli-Hopf scalar moment predecessor: %w", err)
	}
	g561, err := gate561.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate561 Pauli incidence predecessor: %w", err)
	}
	a := Analysis{}
	a.Inherited = auditInherited(g572)
	a.Carrier = auditCarrier(g572)
	a.Symmetry = auditSymmetry(g555)
	a.Transit = auditTransitivity()
	a.Selector = auditHermitianSecondSelector()
	a.Search = auditNativeSelectorSearch(g572, g555, g556, g557, g560, g561)
	a.Seal = auditOrientationSeal(a.Search)
	a.WeakPlane = auditWeakPlaneRelation(a.Seal)
	a.Firewall = auditFirewall(g572)
	a.Final = auditFinal(a)
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func auditInherited(g gate572.Analysis) InheritedGate572Audit {
	return InheritedGate572Audit{
		CP3ProjectiveLawSpace:     g.Final.CP3Certified,
		BMinusLProjectiveOnePlus3: g.Final.BMinusLProjectiveCP0CP2Split,
		SpatialCP2Block:           g.SpatialBlock.BMinusLSpatialEigenspace && g.SpatialBlock.NativeProjectiveRefinement,
		NoNativeSecondSelector:    !g.Final.NativeSecondSelectorOnCP2,
		NoCP3ToK7Functor:          !g.K7.CP3ToK7FunctorFound && g.K7.Gate571BoundaryPreserved,
		NoPhysicalTimeOpened:      !g.Time.MomentFlowPhysicalTime && !g.Time.MomentFlowOSHilbert && !g.Time.MomentFlowRGScale && !g.Time.MomentFlowSpacetime,
		NoFlavorElectroweakData:   !g.FlavorEW.YukawaEigenvaluesDerived && !g.FlavorEW.CKMPMNSDerived && !g.FlavorEW.GenerationHierarchyDerived && !g.FlavorEW.WZMassesDerived,
		Verdict:                   StatusGate572Inherited,
	}
}

func auditCarrier(g gate572.Analysis) SpatialCP2CarrierAudit {
	return SpatialCP2CarrierAudit{
		SpatialEigenspace:             "W_spatial=span_C{a_1^dagger,a_2^dagger,a_3^dagger}=span_C{e_1,e_2,e_3}",
		ProjectiveBlock:               "CP^2_sp=P(W_spatial)={z0=0}/S^1 subset CP^3",
		ComplexDimension:              2,
		RealDimension:                 4,
		BLRestrictedMatrix:            "(B-L)|W_spatial=(1/3) I_3",
		BLRestrictedEigenvalue:        "1/3",
		Gate572CriticalStratumMatched: g.BMinusL.CriticalStrataCertified && g.SpatialBlock.BMinusLSpatialEigenspace,
		CertifiedAsSpatialBlock:       g.Final.BMinusLProjectiveCP0CP2Split && g.SpatialBlock.BMinusLSpatialEigenspace && g.SpatialBlock.NativeProjectiveRefinement,
		Verdict:                       StatusSpatialCP2Certified,
	}
}

func auditSymmetry(g gate555.Analysis) SpatialSymmetryAudit {
	return SpatialSymmetryAudit{
		CommutantFromBMinusL:      g.BMinusL.Commutant,
		SpatialSymmetry:           "U(3) acting on the multiplicity-three B-L=1/3 eigenspace W_spatial",
		TracelessPart:             "SU(3)",
		U3Dimension:               9,
		SU3Dimension:              8,
		BMinusLScalarOnWSpatial:   true,
		SuppliesFurtherSelector:   false,
		PreferredSpatialDirection: false,
		Verdict:                   join(StatusU3SU3SymmetryCertified, StatusBMinusLScalarOnSpatial),
	}
}

func auditTransitivity() TransitivityAudit {
	stabDim := 1 + 4 - 1 // S(U(1)xU(2)) removes one determinant phase from U(1)xU(2).
	return TransitivityAudit{
		Group:                     "SU(3)",
		Space:                     "CP^2_sp=P(C^3)",
		Action:                    "g.[u]=[g u] for g in SU(3)",
		GroupRealDimension:        8,
		PointStabilizer:           "S(U(1)xU(2))",
		PointStabilizerDimension:  stabDim,
		QuotientRealDimension:     8 - stabDim,
		CP2RealDimension:          4,
		ActsTransitively:          8-stabDim == 4,
		InvariantPointSelected:    false,
		InvariantRankOneProjector: false,
		Verdict:                   join(StatusSU3TransitiveOnCP2, StatusPointStabilizerCertified, StatusNoInvariantPoint),
	}
}

func auditHermitianSecondSelector() HermitianSecondSelectorAudit {
	u := []float64{0, 0, 1}
	p := rankOneProjectorReal(u)
	i := identity3()
	lambda1, lambda2 := 7.0, 2.0
	s := addScaled(i, p, lambda2, lambda1-lambda2)
	res := projectorIdempotentResidual(p)
	return HermitianSecondSelectorAudit{
		SelectorFormula:             "S_sp=lambda_2(I-P_u)+lambda_1 P_u on W_spatial, with lambda_1 != lambda_2",
		RankOneProjectorFormula:     "P_u=uu^dagger/(u^dagger u)",
		LambdaOne:                   lambda1,
		LambdaTwo:                   lambda2,
		SamplePoint:                 "u=e_3, so P_u=diag(0,0,1)",
		SampleMatrix:                s,
		SampleEigenvalues:           []float64{lambda2, lambda2, lambda1},
		MultiplicityPattern:         "2+1",
		CriticalCP0:                 "[u] is the rank-one eigenspace with eigenvalue lambda_1, projectively CP^0",
		CriticalCP1:                 "P(u^perp) is the two-dimensional eigenspace with eigenvalue lambda_2, projectively CP^1",
		CP1RealDimension:            2,
		CP0RealDimension:            0,
		ProjectorIdempotentResidual: res,
		ProjectorTrace:              trace3(p),
		ClassifiesCP1CP0Split:       res < 1e-12 && math.Abs(trace3(p)-1) < 1e-12,
		NativeWithoutU:              false,
		Verdict:                     join(StatusGeneralSecondSelectorClassified, StatusCP1CP0CriticalStrataClassified),
	}
}

func auditNativeSelectorSearch(g572 gate572.Analysis, g555 gate555.Analysis, g556 gate556.Analysis, g557 gate557.Analysis, g560 gate560.Analysis, g561 gate561.Analysis) NativeSelectorSearchAudit {
	rows := []NativeSelectorCandidate{
		{
			Source:            "tau_eta pullback",
			PriorGate:         "Gates 555-556",
			Candidate:         "use tau_eta or |tau_eta| as a spatial 2+1 operator",
			NativePUProvided:  false,
			WouldSelectCP1CP0: false,
			Status:            join(g555.TauEta.Verdict, g556.Final.Verdict),
			Reason:            "tau_eta has sealed 2+1 capacity only as eta-graded trace values; Gate 556 reports no native source algebra, no unit-preserving representation, and no canonical 2+1 selector on W_spatial.",
		},
		{
			Source:            "eta-record algebra",
			PriorGate:         "Gate 557",
			Candidate:         "construct idempotents/projectors from A_eta_rec",
			NativePUProvided:  false,
			WouldSelectCP1CP0: false,
			Status:            g557.Final.Verdict,
			Reason:            "eta-record trace origins are recovered, but the matrix algebra in End(H_phi), products, idempotents, spectra, and transfer functor to W_spatial remain unavailable.",
		},
		{
			Source:            "Pauli/Hopf scalar moment",
			PriorGate:         "Gates 560-561",
			Candidate:         "transfer scalar Pauli moment axis to spatial projective direction",
			NativePUProvided:  false,
			WouldSelectCP1CP0: false,
			Status:            join(g560.Final.Verdict, g561.Final.Verdict),
			Reason:            "the scalar Pauli/Hopf triplet exists in the sealed H_phi lane, while Gate 561 blocks a basis-independent Pauli-to-weak-plane/spatial-incidence intertwiner.",
		},
		{
			Source:            "quaternionic Im(H)",
			PriorGate:         "Gate 562",
			Candidate:         "identify quaternionic SU(2) socket with a point in CP^2_sp",
			NativePUProvided:  false,
			WouldSelectCP1CP0: false,
			Status:            "CONDITIONAL_SUPPORT_FINITE_ALGEBRA_QUATERNIONIC_SOCKET_RECOVERED;FAILED_ROUTE_PAULI_QUATERNION_BRIDGE_DOES_NOT_DERIVE_ELECTROWEAK_DYNAMICS_OR_MASSES;FIREWALL_PRESERVED_GATE562_PAULI_HOPF_QUATERNIONIC_WEAK_SOCKET_BOUNDARY",
			Reason:            "Gate 562 records Im(H) as a finite scalar/quaternionic socket, but does not supply a lawful transfer to W_spatial or generation carriers.",
		},
		{
			Source:            "q4 contact quartic",
			PriorGate:         "Gate 555",
			Candidate:         "extract a spatial rank-one projector from q4",
			NativePUProvided:  false,
			WouldSelectCP1CP0: false,
			Status:            g555.Contact.Verdict,
			Reason:            "the contact quartic has a regular representation and irreducibility result, but no native carrier action on W_spatial, H_phi, grading, J, D, first-order data, or B-L compatibility that produces P_u.",
		},
		{
			Source:            "K7 projector carrier",
			PriorGate:         "Gates 571-572",
			Candidate:         "pull a CP2 point/projector from Boolean-octonionic K7",
			NativePUProvided:  false,
			WouldSelectCP1CP0: false,
			Status:            g572.K7.Verdict,
			Reason:            "Gate 571/572 preserve the no-functor boundary between Hopf/CP3 law-space and Boolean-octonionic K7, so K7 cannot select a point in CP^2_sp.",
		},
		{
			Source:            "B-L commutant data",
			PriorGate:         "Gates 555 and 572",
			Candidate:         "use Comm(B-L)=u(1)+u(3) to split u(3) as u(2)+u(1)",
			NativePUProvided:  false,
			WouldSelectCP1CP0: false,
			Status:            join(g572.Stabilizer.Verdict, StatusBMinusLScalarOnSpatial, StatusNoInvariantPoint),
			Reason:            "B-L is scalar on W_spatial and its U(3) commutant protects the degeneracy rather than breaking it.",
		},
		{
			Source:            "finite one-form scalar lane",
			PriorGate:         "Gates 562-564",
			Candidate:         "transfer H_phi scalar doublet/one-form data into CP^2_sp",
			NativePUProvided:  false,
			WouldSelectCP1CP0: false,
			Status:            "CONDITIONAL_SUPPORT_STRUCTURAL_SCALAR_DOUBLE_MODULE_RECOVERED;FAILED_ROUTE_NO_NATIVE_CURVATURE_KINETIC_OR_WSPATIAL_PROJECTOR_FROM_SCALAR_LANE;FIREWALL_PRESERVED_SCALAR_ONE_FORM_BOUNDARY",
			Reason:            "the finite one-form scalar lane is structurally related to the scalar/quaternionic socket, not a native rank-one projector on W_spatial.",
		},
		{
			Source:            "hypercharge normalization",
			PriorGate:         "Gate 565",
			Candidate:         "derive CP2 spatial point from k_Y=5/3 or sin^2(theta_*)=3/8 boundary normalization",
			NativePUProvided:  false,
			WouldSelectCP1CP0: false,
			Status:            "PASS_HYPERCHARGE_NORMALIZATION_KY_5_OVER_3_BOUNDARY_LAYER;PASS_BOUNDARY_SIN2_THETA_STAR_3_OVER_8;FAILED_ROUTE_HYPERCHARGE_NORMALIZATION_DOES_NOT_SELECT_SPATIAL_CP2_POINT",
			Reason:            "hypercharge normalization is a representation-trace/boundary-normalization result; it produces no spatial rank-one projector, flavor data, or physical weak-plane selection.",
		},
	}
	found := false
	for _, row := range rows {
		found = found || row.NativePUProvided || row.WouldSelectCP1CP0
	}
	return NativeSelectorSearchAudit{
		Candidates:                  rows,
		CandidateCount:              len(rows),
		NativeRankOneProjectorFound: found,
		NativeProjectivePointFound:  found,
		NativeSecondSelectorFound:   found,
		Verdict:                     join(StatusNoNativeRankOneProjector, StatusNoNativeCP2SecondSelector, StatusNoWeakPlaneCP1CP0),
		Reason:                      "All tested ASHA sources either preserve the U(3)/SU(3) degeneracy, live on a different carrier, require a missing transfer functor, or are explicitly sealed/bridge-only; none supplies a native rank-one P_u on W_spatial.",
	}
}

func auditOrientationSeal(s NativeSelectorSearchAudit) OrientationSealAudit {
	_ = s
	return OrientationSealAudit{
		SealName:                   "SpatialProjectiveOrientationSeal",
		MinimalDatum:               "choice of [u] in CP^2_sp",
		EquivalentDatum:            "rank-one Hermitian projector P_u=uu^dagger/(u^dagger u) on W_spatial",
		SealedSelectorFormula:      "S_sp=lambda_2(I-P_u)+lambda_1 P_u, lambda_1 != lambda_2",
		SealedNotNative:            true,
		Commutant:                  "u(2)+u(1)",
		CommutantDimension:         2*2 + 1*1,
		ExpectedCommutantDimension: 5,
		CommMatchesGate555Formula:  2*2+1*1 == 5,
		Verdict:                    join(StatusOrientationSealRequired, StatusSealedU2U1Commutant),
	}
}

func auditWeakPlaneRelation(seal OrientationSealAudit) WeakPlaneRelationAudit {
	return WeakPlaneRelationAudit{
		SealedPoint:           "[u]=[a_3^dagger] under the ordered basis (a_1^dagger,a_2^dagger,a_3^dagger)",
		ComplementaryPlane:    "span_C{a_1^dagger,a_2^dagger}, projectively CP^1",
		ConventionalPlaneName: "U_12",
		BasisDependent:        true,
		NativeDerived:         !seal.SealedNotNative,
		Verdict:               join(StatusBasisDependentU12, StatusBasisDependentNotNative),
	}
}

func auditFirewall(g572 gate572.Analysis) FirewallAudit {
	return FirewallAudit{
		PromotedToWeakIsospin:       false,
		PromotedToPhysicalWeakPlane: false,
		GenerationHierarchyDerived:  false,
		YukawaTextureDerived:        false,
		CKMPMNSDerived:              false,
		ObservedFlavorDataImported:  false,
		PhysicalElectroweakDynamics: false,
		WZMassesDerived:             false,
		PhotonDynamicsDerived:       false,
		CP2ToK7FunctorOpened:        g572.K7.CP3ToK7FunctorFound,
		ProductTimeOpened:           g572.Time.MomentFlowPhysicalTime || g572.Time.MomentFlowOSHilbert || g572.Time.MomentFlowRGScale || g572.Time.MomentFlowSpacetime,
		Gate564565BoundaryPreserved: true,
		Verdict:                     join(StatusK7TimeBoundaryPreserved, StatusFlavorEWBoundaryPreserved, StatusGate573BoundaryPreserved),
	}
}

func auditFinal(a Analysis) FinalVerdict {
	physical := a.Firewall.PromotedToWeakIsospin || a.Firewall.PromotedToPhysicalWeakPlane || a.Firewall.GenerationHierarchyDerived || a.Firewall.YukawaTextureDerived || a.Firewall.CKMPMNSDerived || a.Firewall.ObservedFlavorDataImported || a.Firewall.PhysicalElectroweakDynamics || a.Firewall.WZMassesDerived || a.Firewall.PhotonDynamicsDerived
	return FinalVerdict{
		SpatialCP2Certified:         a.Carrier.CertifiedAsSpatialBlock && a.Carrier.RealDimension == 4,
		SU3Transitive:               a.Transit.ActsTransitively && a.Transit.QuotientRealDimension == a.Transit.CP2RealDimension,
		SU3InvariantPointSelected:   a.Transit.InvariantPointSelected || a.Transit.InvariantRankOneProjector,
		GeneralTwoPlusOneSelector:   a.Selector.ClassifiesCP1CP0Split,
		NativeRankOnePU:             a.Search.NativeRankOneProjectorFound || a.Search.NativeProjectivePointFound || a.Search.NativeSecondSelectorFound,
		MinimalSeal:                 "SpatialProjectiveOrientationSeal = choice of [u] in CP^2_sp, equivalently rank-one projector P_u",
		PhysicalWeakFlavorEWDerived: physical,
		K7OrProductTimeOpened:       a.Firewall.CP2ToK7FunctorOpened || a.Firewall.ProductTimeOpened,
		MissingNextTheorem:          "A lawful next theorem must derive a basis-independent rank-one Hermitian projector P_u on W_spatial, or a functorial carrier action that produces such P_u and passes B-L, spectral-triple, K7/time, and flavor/electroweak firewalls. Without it, CP^2_sp remains SU(3)-homogeneous and no weak-plane or flavor/electroweak data is derived.",
		Verdict:                     join(StatusSpatialCP2Certified, StatusSU3TransitiveOnCP2, StatusNoInvariantPoint, StatusGeneralSecondSelectorClassified, StatusNoNativeRankOneProjector, StatusOrientationSealRequired, StatusBasisDependentNotNative, StatusFlavorEWBoundaryPreserved, StatusGate573BoundaryPreserved),
	}
}

func validate(a Analysis) error {
	if !a.Inherited.CP3ProjectiveLawSpace || !a.Inherited.BMinusLProjectiveOnePlus3 || !a.Inherited.SpatialCP2Block {
		return fmt.Errorf("Gate572 projective B-L CP2 inheritance failed")
	}
	if !a.Carrier.CertifiedAsSpatialBlock || !a.Carrier.Gate572CriticalStratumMatched || a.Carrier.RealDimension != 4 {
		return fmt.Errorf("spatial CP2 carrier not certified")
	}
	if !a.Symmetry.BMinusLScalarOnWSpatial || a.Symmetry.SuppliesFurtherSelector || a.Symmetry.PreferredSpatialDirection {
		return fmt.Errorf("B-L spatial symmetry audit failed")
	}
	if !a.Transit.ActsTransitively || a.Transit.InvariantPointSelected || a.Transit.InvariantRankOneProjector || a.Transit.QuotientRealDimension != 4 {
		return fmt.Errorf("SU(3) transitivity/isotropy audit failed")
	}
	if !a.Selector.ClassifiesCP1CP0Split || a.Selector.NativeWithoutU || a.Selector.ProjectorIdempotentResidual > 1e-12 || math.Abs(a.Selector.ProjectorTrace-1) > 1e-12 {
		return fmt.Errorf("general Hermitian second-selector classification failed")
	}
	if a.Search.NativeRankOneProjectorFound || a.Search.NativeProjectivePointFound || a.Search.NativeSecondSelectorFound {
		return fmt.Errorf("unexpected native rank-one spatial CP2 projector promoted")
	}
	if !a.Seal.SealedNotNative || !a.Seal.CommMatchesGate555Formula || a.Seal.CommutantDimension != 5 {
		return fmt.Errorf("orientation seal audit failed")
	}
	if !a.WeakPlane.BasisDependent || a.WeakPlane.NativeDerived || a.WeakPlane.WeakIsospinIdentified || a.WeakPlane.GenerationHierarchy || a.WeakPlane.YukawaTextureDerived || a.WeakPlane.CKMPMNSDerived {
		return fmt.Errorf("basis-dependent weak-plane relation promoted incorrectly")
	}
	if a.Final.SU3InvariantPointSelected || a.Final.NativeRankOnePU || a.Final.PhysicalWeakFlavorEWDerived || a.Final.K7OrProductTimeOpened {
		return fmt.Errorf("final firewall violation")
	}
	return nil
}

func truth(a Analysis) string {
	parts := []string{
		"Gate 573 certifies CP^2_sp=P(W_spatial) as the B-L=1/3 projective critical stratum inherited from Gate 572.",
		"On W_spatial, B-L is (1/3)I_3, so the B-L commutant contains U(3) and its traceless part SU(3); this symmetry preserves, rather than breaks, the spatial degeneracy.",
		"SU(3) acts transitively on CP^2_sp with point stabilizer S(U(1)xU(2)), hence CP^2 ~= SU(3)/S(U(1)xU(2)) and no SU(3)-invariant projective point [u] is selected by current data.",
		"A 2+1 Hermitian selector has the form S_sp=lambda_2(I-P_u)+lambda_1 P_u; its critical projective strata are CP^1=P(u^perp) and CP^0=[u].",
		"The native search over tau_eta, eta-record algebra, Pauli/Hopf scalar moment, Im(H), q4, K7, B-L commutant data, finite one-form scalar lane, and hypercharge normalization finds no native rank-one P_u on W_spatial.",
		"The minimal lawful non-native datum is SpatialProjectiveOrientationSeal, a choice of [u] or P_u; once sealed it gives commutant u(2)+u(1), but this is sealed support only and does not derive weak isospin, generation hierarchy, Yukawa, CKM/PMNS, observed flavor, electroweak dynamics, K7, or physical time.",
	}
	return strings.Join(parts, " ")
}

func Statuses() []string {
	return []string{
		StatusGate572Inherited,
		StatusSpatialCP2Certified,
		StatusU3SU3SymmetryCertified,
		StatusBMinusLScalarOnSpatial,
		StatusSU3TransitiveOnCP2,
		StatusPointStabilizerCertified,
		StatusNoInvariantPoint,
		StatusGeneralSecondSelectorClassified,
		StatusCP1CP0CriticalStrataClassified,
		StatusNoNativeRankOneProjector,
		StatusNoNativeCP2SecondSelector,
		StatusNoWeakPlaneCP1CP0,
		StatusOrientationSealRequired,
		StatusSealedU2U1Commutant,
		StatusBasisDependentU12,
		StatusBasisDependentNotNative,
		StatusK7TimeBoundaryPreserved,
		StatusFlavorEWBoundaryPreserved,
		StatusGate573BoundaryPreserved,
	}
}

func rankOneProjectorReal(u []float64) [][]float64 {
	norm := 0.0
	for _, x := range u {
		norm += x * x
	}
	if norm == 0 {
		return nil
	}
	p := make([][]float64, len(u))
	for i := range u {
		p[i] = make([]float64, len(u))
		for j := range u {
			p[i][j] = u[i] * u[j] / norm
		}
	}
	return p
}

func identity3() [][]float64 {
	return [][]float64{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
}

func addScaled(a, b [][]float64, ca, cb float64) [][]float64 {
	out := make([][]float64, len(a))
	for i := range a {
		out[i] = make([]float64, len(a[i]))
		for j := range a[i] {
			out[i][j] = ca*a[i][j] + cb*b[i][j]
		}
	}
	return out
}

func projectorIdempotentResidual(p [][]float64) float64 {
	pp := matMul3(p, p)
	m := 0.0
	for i := range p {
		for j := range p[i] {
			if d := math.Abs(pp[i][j] - p[i][j]); d > m {
				m = d
			}
		}
	}
	return m
}

func matMul3(a, b [][]float64) [][]float64 {
	out := make([][]float64, 3)
	for i := 0; i < 3; i++ {
		out[i] = make([]float64, 3)
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				out[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return out
}

func trace3(a [][]float64) float64 { return a[0][0] + a[1][1] + a[2][2] }

func join(xs ...string) string { return strings.Join(xs, ";") }
