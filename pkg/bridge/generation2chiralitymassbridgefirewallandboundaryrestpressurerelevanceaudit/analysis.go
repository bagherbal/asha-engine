// Package generation2chiralitymassbridgefirewallandboundaryrestpressurerelevanceaudit implements
// Gate 812: Chirality Mass-Bridge Firewall and Boundary RestPressure Relevance Audit.
//
// Gate 812 follows Gate 811 by checking whether a Cl(1,7) chirality/mass-edge
// idea can source Yukawa trace magnitudes, top dominance, rest pressure, or the
// boundary second-moment correction. It preserves the real-form firewall:
// omega^2=-1 blocks naive real projectors (1±omega)/2, and any chirality split
// requires a complex or equivalent airlock. The mass bridge is useful as an edge
// template audit, not as a Yukawa value or rest-pressure theorem.
package generation2chiralitymassbridgefirewallandboundaryrestpressurerelevanceaudit

import (
	"fmt"
	"strings"
)

const (
	AuditID = "GATE812-CHIRALITY-MASS-BRIDGE-FIREWALL-BOUNDARY-RESTPRESSURE-RELEVANCE-AUDIT"

	NEff      = 3.0023273474722147
	DeltaN    = NEff - 3.0
	SBoundary = 0.0012924448188162962
	PBoundary = 7.0 / 72.0
	CYukawa   = 0.9992248188812008
	CHistory  = 1.038025177923625
	CHiggs    = 1.0372205204048603

	OmegaSquared = -1.0

	StatusGate811Inherited              = "PASS_GATE811_HYPERCHARGE_COLOR_POSITIVE_REST_CORRECTION_INHERITED"
	StatusBoundarySecondMomentSelected  = "PASS_BOUNDARY_SECOND_MOMENT_CORRECTION_SELECTED_AS_CURRENT_REST_PRESSURE_TARGET"
	StatusCL17PseudoscalarInherited     = "PASS_CL17_PSEUDOSCALAR_STATUS_INHERITED"
	StatusRealChiralityProjectorAudited = "PASS_REAL_CHIRALITY_PROJECTOR_VALIDITY_AUDITED"
	StatusWeakRestrictionRequirements   = "PASS_WEAK_CHIRALITY_RESTRICTION_REQUIREMENTS_DEFINED"
	StatusHiggsMassBridgeAudited        = "PASS_HIGGS_SCALAR_MASS_BRIDGE_AUDITED"
	StatusFiniteTripleEdgeCompat        = "PASS_FINITE_TRIPLE_EDGE_TEMPLATE_COMPATIBILITY_RECORDED"
	StatusGradeZeroAudited              = "PASS_GRADE_ZERO_COMMUTING_CLAIM_AUDITED"
	StatusTraceMagnitudeRelation        = "PASS_RELATION_TO_TRACE_MAGNITUDE_OPERATOR_SEAL_AUDITED"
	StatusBoundaryFNRelation            = "PASS_RELATION_TO_BOUNDARY_FN_PACKAGE_AUDITED"
	StatusLagrangianHierarchySeparated  = "PASS_LAGRANGIAN_HIERARCHY_CLAIM_SEPARATED_FROM_YUKAWA_HIERARCHY"
	StatusChiralityIdeaClassified       = "PASS_CHIRALITY_IDEA_STATUS_CLASSIFIED"
	StatusCHiggsFirewall                = "PASS_C_HIGGS_FIREWALL_PRESERVED"
	StatusBranchDecision                = "PASS_BRANCH_DECISION_RECORDED"
	StatusPhysicalFirewalls             = "PASS_PHYSICAL_FIREWALLS_ENFORCED"

	StatusChiralityAirlockRequired = "CONDITIONAL_SUPPORT_CHIRALITY_PROJECTORS_REQUIRE_COMPLEX_OR_EQUIVALENT_AIRLOCK_IN_ACTIVE_CL17_BOARD"
	StatusChiralRestrictionSearch  = "CONDITIONAL_SUPPORT_CHIRAL_RESTRICTION_MAY_BE_A_GAUGE_REPRESENTATION_SEARCH_CANDIDATE"
	StatusMassBridgeTyped          = "CONDITIONAL_SUPPORT_HIGGS_SCALAR_OR_ONE_FORM_EDGE_CAN_BE_TYPED_AS_LEFT_RIGHT_MASS_BRIDGE"
	StatusMassBridgeEdgeOnly       = "CONDITIONAL_SUPPORT_CHIRAL_MASS_BRIDGE_REINFORCES_EDGE_TEMPLATE_NOT_YUKAWA_VALUE_SOURCE"
	StatusScalarBilinearCompat     = "CONDITIONAL_SUPPORT_SCALAR_FIELD_IS_COMPATIBLE_WITH_LEFT_RIGHT_YUKAWA_BILINEAR"
	StatusBridgeEdgeCompatOnly     = "CONDITIONAL_SUPPORT_CHIRALITY_MASS_BRIDGE_SUPPLIES_EDGE_COMPATIBILITY_ONLY"
	StatusOrthogonalToBoundaryFN   = "CONDITIONAL_SUPPORT_CHIRALITY_MASS_BRIDGE_IS_ORTHOGONAL_TO_BOUNDARY_FN_COEFFICIENT_SOURCE"
	StatusGaugeGravitySeparate     = "CONDITIONAL_SUPPORT_GAUGE_GRAVITY_HIERARCHY_MAY_BE_SEPARATE_STRUCTURAL_LANE"
	StatusUsefulFirewallAudit      = "CONDITIONAL_SUPPORT_CHIRALITY_IDEA_IS_USEFUL_AS_MASS_EDGE_FIREWALL_AUDIT"

	StatusNaiveRealProjectorsInvalid = "FAILED_ROUTE_REAL_PROJECTORS_ONE_PLUS_MINUS_OMEGA_NOT_VALID_WHEN_OMEGA_SQUARED_MINUS_ONE"
	StatusNoNativeRealChirality      = "FAILED_ROUTE_NATIVE_REAL_CL17_CHIRALITY_SPLIT_NOT_CERTIFIED_BY_NAIVE_PSEUDOSCALAR_PROJECTORS"
	StatusComplexAirlockNotNative    = "FAILED_ROUTE_COMPLEX_CHIRALITY_AIRLOCK_NOT_NATIVE_REAL_PROJECTOR_THEOREM"
	StatusNoNativeSU2L               = "FAILED_ROUTE_NO_NATIVE_SU2L_THEOREM_FROM_G_I_P_L_IN_THIS_GATE"
	StatusAirlockNoGaugeAssignments  = "FAILED_ROUTE_CHIRAL_PROJECTOR_AIRLOCK_DOES_NOT_DERIVE_STANDARD_MODEL_GAUGE_ASSIGNMENTS"
	StatusWeakNoTraceMagnitudes      = "FAILED_ROUTE_WEAK_CHIRALITY_RESTRICTION_DOES_NOT_SOURCE_YUKAWA_TRACE_MAGNITUDES"
	StatusHiggsScalarNoYf            = "FAILED_ROUTE_HIGGS_GRADE_ZERO_OR_SCALAR_STATUS_DOES_NOT_DERIVE_Y_F_OPERATORS"
	StatusMassBridgeNoEigenvalues    = "FAILED_ROUTE_MASS_BRIDGE_TEMPLATE_DOES_NOT_DERIVE_YUKAWA_EIGENVALUES"
	StatusEdgeNoDeltaN               = "FAILED_ROUTE_CHIRAL_EDGE_EXISTENCE_DOES_NOT_SOURCE_N_EFF_MINUS_THREE"
	StatusMassBridgeNoTopRest        = "FAILED_ROUTE_HIGGS_MASS_BRIDGE_DOES_NOT_SUPPLY_TOP_DOMINANCE_OR_REST_PRESSURE_OPERATOR"
	StatusCommutingNoEdge            = "FAILED_ROUTE_COMMUTING_WITH_CHIRALITY_DOES_NOT_BY_ITSELF_SOURCE_CHIRAL_MASS_EDGE"
	StatusGradeZeroNotUniqueYukawa   = "FAILED_ROUTE_GRADE_ZERO_SCALAR_NOT_UNIQUE_YUKAWA_VALUE_SOURCE"
	StatusScalarNotHierarchyOperator = "FAILED_ROUTE_SCALAR_HIGGS_TEMPLATE_NOT_HIERARCHY_BREAKING_OPERATOR"
	StatusNoHermitianTraceOps        = "FAILED_ROUTE_CHIRALITY_MASS_BRIDGE_DOES_NOT_SUPPLY_HERMITIAN_TRACE_MAGNITUDE_OPERATORS"
	StatusNoPositiveTraceAtoms       = "FAILED_ROUTE_CHIRALITY_MASS_BRIDGE_DOES_NOT_SUPPLY_POSITIVE_TRACE_ATOMS"
	StatusNoTopColorBlock            = "FAILED_ROUTE_CHIRALITY_MASS_BRIDGE_DOES_NOT_DERIVE_TOP_COLOR_BLOCK"
	StatusNoRestPressureOperator     = "FAILED_ROUTE_CHIRALITY_MASS_BRIDGE_DOES_NOT_DERIVE_REST_PRESSURE_OPERATOR"
	StatusChiralityNoNineFive        = "FAILED_ROUTE_CHIRALITY_PROJECTOR_DOES_NOT_SOURCE_NINE_OVER_FIVE"
	StatusHiggsNoSixPS2              = "FAILED_ROUTE_HIGGS_SCALAR_MASS_BRIDGE_DOES_NOT_SOURCE_SIX_P_S_SQUARED"
	StatusNoBoundaryTraceMap         = "FAILED_ROUTE_CHIRALITY_STRUCTURE_DOES_NOT_REPLACE_BOUNDARY_TO_TRACE_MAGNITUDE_MAP"
	StatusNoPositiveRestSpectrum     = "FAILED_ROUTE_CHIRALITY_STRUCTURE_DOES_NOT_CONSTRUCT_POSITIVE_REST_SPECTRUM"
	StatusForceNotYukawa             = "FAILED_ROUTE_FORCE_HIERARCHY_NOT_YUKAWA_HIERARCHY_THEOREM"
	StatusAlphaOverAlphaGNoDeltaN    = "FAILED_ROUTE_ALPHA_OVER_ALPHA_G_DOES_NOT_SOURCE_N_EFF_MINUS_THREE"
	StatusGravityNotRestPressure     = "FAILED_ROUTE_GRAVITATIONAL_STIFFNESS_NOT_REST_PRESSURE_OPERATOR"
	StatusDoesNotSolveGate811        = "FAILED_ROUTE_CHIRALITY_IDEA_DOES_NOT_SOLVE_GATE811_REST_PRESSURE_MISSING_OBJECT"
	StatusNoCYukawaUpdate            = "FAILED_ROUTE_GATE812_DOES_NOT_UPDATE_C_YUKAWA"
	StatusCHiggsLevelB               = "FAILED_ROUTE_C_HIGGS_REMAINS_LEVEL_B"
	StatusFirewallGate812            = "FIREWALL_PRESERVED_GATE812_CHIRALITY_MASS_BRIDGE_REST_PRESSURE_RELEVANCE_BOUNDARY"
)

type Inheritance struct {
	Gate811Inherited, BoundarySecondMomentSelected bool
	NEff, DeltaN, S, P, CYukawa, CHistory, CHiggs  float64
	M2, ResidualBoundary, C2Obs                    float64
	Verdicts                                       []string
}

type PseudoscalarAudit struct {
	Inherited, Audited                                                 bool
	OmegaSquared                                                       float64
	NaiveProjectorsIdempotent                                          bool
	PPlusScalar, PPlusOmega, PPlusSquaredScalar, PPlusSquaredOmega     float64
	PMinusScalar, PMinusOmega, PMinusSquaredScalar, PMinusSquaredOmega float64
	ComplexGammaSquared                                                float64
	Supports, Failures, Verdicts                                       []string
}

type RequirementAudit struct {
	Audited            bool
	Name               string
	Requirements       []string
	Supports, Failures []string
	Verdicts           []string
	Description        string
}

type MassBridgeAudit struct {
	Audited            bool
	Facts              []string
	Supports, Failures []string
	Verdicts           []string
}

type BoundaryFNRelevance struct {
	Audited            bool
	ExistingSources    []string
	NotSourced         []string
	Supports, Failures []string
	Verdicts           []string
}

type LagrangianHierarchyAudit struct {
	Audited            bool
	Separated          bool
	Supports, Failures []string
	Verdicts           []string
}

type ChiralityIdeaStatus struct {
	Classified bool
	UsefulFor  []string
	NotUseful  []string
	Supports   []string
	Failures   []string
	Verdict    string
}

type Impact struct {
	Preserved                       bool
	NEff, CYukawa, CHistory, CHiggs float64
	Failures, Verdicts              []string
}

type Firewalls struct {
	Enforced                                                                                      bool
	NoNaiveRealProjectors, NoSU2Shortcut, NoYukawaValues, NoRestPressure, NoBoundaryFNReplacement bool
	NoForceHierarchyShortcut, NoLedgerUpdate, NoPoleMass                                          bool
	Verdict                                                                                       string
}

type BranchDecision struct {
	Recorded                   bool
	Next, Alternative, Verdict string
}

type Analysis struct {
	Inheritance         Inheritance
	Pseudoscalar        PseudoscalarAudit
	WeakRestriction     RequirementAudit
	HiggsMassBridge     MassBridgeAudit
	GradeZero           RequirementAudit
	TraceMagnitude      RequirementAudit
	BoundaryFN          BoundaryFNRelevance
	LagrangianHierarchy LagrangianHierarchyAudit
	IdeaStatus          ChiralityIdeaStatus
	Impact              Impact
	Firewalls           Firewalls
	Branch              BranchDecision
	Truth               string
	Final               string
}

func M2(s float64) float64                       { return PBoundary * s * s }
func BoundaryResidual(deltaN, s float64) float64 { return deltaN - (9.0/5.0)*s }
func C2Observed(deltaN, s float64) float64       { return BoundaryResidual(deltaN, s) / M2(s) }

func BuildDefault() (Analysis, error) {
	m2 := M2(SBoundary)
	residual := BoundaryResidual(DeltaN, SBoundary)
	c2 := C2Observed(DeltaN, SBoundary)
	return Analysis{
		Inheritance: Inheritance{Gate811Inherited: true, BoundarySecondMomentSelected: true, NEff: NEff, DeltaN: DeltaN, S: SBoundary, P: PBoundary, CYukawa: CYukawa, CHistory: CHistory, CHiggs: CHiggs, M2: m2, ResidualBoundary: residual, C2Obs: c2, Verdicts: []string{StatusGate811Inherited, StatusBoundarySecondMomentSelected}},
		Pseudoscalar: PseudoscalarAudit{Inherited: true, Audited: true, OmegaSquared: OmegaSquared, NaiveProjectorsIdempotent: false,
			PPlusScalar: 0.5, PPlusOmega: 0.5, PPlusSquaredScalar: 0, PPlusSquaredOmega: 0.5,
			PMinusScalar: 0.5, PMinusOmega: -0.5, PMinusSquaredScalar: 0, PMinusSquaredOmega: -0.5,
			ComplexGammaSquared: 1,
			Verdicts:            []string{StatusCL17PseudoscalarInherited, StatusRealChiralityProjectorAudited},
			Supports:            []string{StatusChiralityAirlockRequired},
			Failures:            []string{StatusNaiveRealProjectorsInvalid, StatusNoNativeRealChirality, StatusComplexAirlockNotNative}},
		WeakRestriction:     RequirementAudit{Audited: true, Name: "weak chirality restriction T_i = G_i P_L", Requirements: []string{"valid chirality projector P_L", "typed generator set G_i", "closure [G_i P_L,G_j P_L] as su(2)", "action on finite Standard Model left-handed doublets", "hypercharge/color/right-handed singlet preservation", "finite spectral triple compatibility"}, Verdicts: []string{StatusWeakRestrictionRequirements}, Supports: []string{StatusChiralRestrictionSearch}, Failures: []string{StatusNoNativeSU2L, StatusAirlockNoGaugeAssignments, StatusWeakNoTraceMagnitudes}, Description: "structural resonance only; finite spectral triple already supplies certified gauge-compatible chiral edge templates"},
		HiggsMassBridge:     MassBridgeAudit{Audited: true, Facts: []string{"a scalar field can appear in a left-right Yukawa bilinear", "finite spectral triple certifies Q_L->u_R, Q_L->d_R, L_L->e_R, L_L->nu_R edge locations", "allowed mass bridge tells where a coupling may live, not what the coupling values are"}, Verdicts: []string{StatusHiggsMassBridgeAudited, StatusFiniteTripleEdgeCompat}, Supports: []string{StatusMassBridgeTyped, StatusMassBridgeEdgeOnly}, Failures: []string{StatusHiggsScalarNoYf, StatusMassBridgeNoEigenvalues, StatusEdgeNoDeltaN, StatusMassBridgeNoTopRest}},
		GradeZero:           RequirementAudit{Audited: true, Name: "grade-zero commuting claim", Requirements: []string{"grade-0 scalars commute with Clifford elements", "commuting with chirality is not the same as chirality flipping", "mass bridge is typed through bilinear/finite Dirac edge structure"}, Verdicts: []string{StatusGradeZeroAudited}, Supports: []string{StatusScalarBilinearCompat}, Failures: []string{StatusCommutingNoEdge, StatusGradeZeroNotUniqueYukawa, StatusScalarNotHierarchyOperator}, Description: "scalar Higgs is compatible with a left-right bilinear but does not source Yukawa values or hierarchy"},
		TraceMagnitude:      RequirementAudit{Audited: true, Name: "TraceMagnitudeOperatorSeal relation", Requirements: []string{"H_f=Y_f†Y_f", "positive spectra Spec(H_f)", "T=h_t", "a_rest,b_rest,alpha,beta,q_rest", "sector trace atoms"}, Verdicts: []string{StatusTraceMagnitudeRelation}, Supports: []string{StatusBridgeEdgeCompatOnly}, Failures: []string{StatusNoHermitianTraceOps, StatusNoPositiveTraceAtoms, StatusNoTopColorBlock, StatusNoRestPressureOperator}, Description: "chirality/mass bridge supplies edge compatibility only, not Hermitian trace magnitudes"},
		BoundaryFN:          BoundaryFNRelevance{Audited: true, ExistingSources: []string{"color factor 3 from finite trace multiplicity", "hypercharge coefficient 5/3 from reduced flavor/boundary ledger", "boundary-pair factor 2 from boundary exterior response package"}, NotSourced: []string{"9/5 by chirality projector", "6 p s^2 by Higgs scalar mass bridge", "boundary-to-trace-magnitude map", "positive rest spectrum"}, Verdicts: []string{StatusBoundaryFNRelation}, Supports: []string{StatusOrthogonalToBoundaryFN}, Failures: []string{StatusChiralityNoNineFive, StatusHiggsNoSixPS2, StatusNoBoundaryTraceMap, StatusNoPositiveRestSpectrum}},
		LagrangianHierarchy: LagrangianHierarchyAudit{Audited: true, Separated: true, Verdicts: []string{StatusLagrangianHierarchySeparated}, Supports: []string{StatusGaugeGravitySeparate}, Failures: []string{StatusForceNotYukawa, StatusAlphaOverAlphaGNoDeltaN, StatusGravityNotRestPressure}},
		IdeaStatus:          ChiralityIdeaStatus{Classified: true, UsefulFor: []string{"real-form chirality audit", "left-right mass-edge compatibility", "finite triple edge interpretation", "scalar/Higgs bridge firewalls"}, NotUseful: []string{"Yukawa magnitudes", "top dominance", "rest pressure", "Gate811 boundary second-moment correction", "N_eff", "C_Yukawa", "C_Higgs"}, Supports: []string{StatusUsefulFirewallAudit}, Failures: []string{StatusDoesNotSolveGate811}, Verdict: StatusChiralityIdeaClassified},
		Impact:              Impact{Preserved: true, NEff: NEff, CYukawa: CYukawa, CHistory: CHistory, CHiggs: CHiggs, Verdicts: []string{StatusCHiggsFirewall}, Failures: []string{StatusNoCYukawaUpdate, StatusCHiggsLevelB}},
		Firewalls:           Firewalls{Enforced: true, NoNaiveRealProjectors: true, NoSU2Shortcut: true, NoYukawaValues: true, NoRestPressure: true, NoBoundaryFNReplacement: true, NoForceHierarchyShortcut: true, NoLedgerUpdate: true, NoPoleMass: true, Verdict: StatusFirewallGate812},
		Branch:              BranchDecision{Recorded: true, Next: "Gate 813 — Boundary Second-Moment RestPressure Correction and Positive Spectrum Construction Audit", Alternative: "Gate 813 — ChiralityMassBridge TraceMagnitudeOperator Source Audit only if a valid chirality-to-trace-magnitude map is unexpectedly found", Verdict: StatusBranchDecision},
		Truth:               "Gate 812 protects ASHA from using a chirality/mass-edge resonance as a Yukawa trace-magnitude source.",
		Final:               "The chirality idea is useful as a mass-edge firewall audit, but the missing object remains the Gate811 boundary-to-rest-pressure map and positive spectrum construction.",
	}, nil
}

func Statuses() []string {
	return []string{StatusGate811Inherited, StatusBoundarySecondMomentSelected, StatusCL17PseudoscalarInherited, StatusRealChiralityProjectorAudited, StatusWeakRestrictionRequirements, StatusHiggsMassBridgeAudited, StatusFiniteTripleEdgeCompat, StatusGradeZeroAudited, StatusTraceMagnitudeRelation, StatusBoundaryFNRelation, StatusLagrangianHierarchySeparated, StatusChiralityIdeaClassified, StatusCHiggsFirewall, StatusBranchDecision, StatusPhysicalFirewalls, StatusChiralityAirlockRequired, StatusChiralRestrictionSearch, StatusMassBridgeTyped, StatusMassBridgeEdgeOnly, StatusScalarBilinearCompat, StatusBridgeEdgeCompatOnly, StatusOrthogonalToBoundaryFN, StatusGaugeGravitySeparate, StatusUsefulFirewallAudit, StatusNaiveRealProjectorsInvalid, StatusNoNativeRealChirality, StatusComplexAirlockNotNative, StatusNoNativeSU2L, StatusAirlockNoGaugeAssignments, StatusWeakNoTraceMagnitudes, StatusHiggsScalarNoYf, StatusMassBridgeNoEigenvalues, StatusEdgeNoDeltaN, StatusMassBridgeNoTopRest, StatusCommutingNoEdge, StatusGradeZeroNotUniqueYukawa, StatusScalarNotHierarchyOperator, StatusNoHermitianTraceOps, StatusNoPositiveTraceAtoms, StatusNoTopColorBlock, StatusNoRestPressureOperator, StatusChiralityNoNineFive, StatusHiggsNoSixPS2, StatusNoBoundaryTraceMap, StatusNoPositiveRestSpectrum, StatusForceNotYukawa, StatusAlphaOverAlphaGNoDeltaN, StatusGravityNotRestPressure, StatusDoesNotSolveGate811, StatusNoCYukawaUpdate, StatusCHiggsLevelB, StatusFirewallGate812}
}

func FormatPseudoscalar(a PseudoscalarAudit) string {
	return fmt.Sprintf("omega^2=%.0f Pplus=(%.1f%+.1fω) Pplus^2=(%.1f%+.1fω) Pminus=(%.1f%+.1fω) Pminus^2=(%.1f%+.1fω) gamma_chi^2=%.0f idempotent=%v supports=[%s] failures=[%s]", a.OmegaSquared, a.PPlusScalar, a.PPlusOmega, a.PPlusSquaredScalar, a.PPlusSquaredOmega, a.PMinusScalar, a.PMinusOmega, a.PMinusSquaredScalar, a.PMinusSquaredOmega, a.ComplexGammaSquared, a.NaiveProjectorsIdempotent, strings.Join(a.Supports, "; "), strings.Join(a.Failures, "; "))
}

func FormatRequirement(a RequirementAudit) string {
	return fmt.Sprintf("%s requirements=[%s] supports=[%s] failures=[%s]", a.Name, strings.Join(a.Requirements, "; "), strings.Join(a.Supports, "; "), strings.Join(a.Failures, "; "))
}

func FormatMassBridge(a MassBridgeAudit) string {
	return fmt.Sprintf("facts=[%s] supports=[%s] failures=[%s]", strings.Join(a.Facts, "; "), strings.Join(a.Supports, "; "), strings.Join(a.Failures, "; "))
}

func FormatBoundaryFN(a BoundaryFNRelevance) string {
	return fmt.Sprintf("existing=[%s] not_sourced=[%s] supports=[%s] failures=[%s]", strings.Join(a.ExistingSources, "; "), strings.Join(a.NotSourced, "; "), strings.Join(a.Supports, "; "), strings.Join(a.Failures, "; "))
}

func FormatStatus(a ChiralityIdeaStatus) string {
	return fmt.Sprintf("useful_for=[%s] not_useful_for=[%s] supports=[%s] failures=[%s]", strings.Join(a.UsefulFor, "; "), strings.Join(a.NotUseful, "; "), strings.Join(a.Supports, "; "), strings.Join(a.Failures, "; "))
}

func containsAll(hay []string, needles []string) bool {
	m := map[string]bool{}
	for _, h := range hay {
		m[h] = true
	}
	for _, n := range needles {
		if !m[n] {
			return false
		}
	}
	return true
}
