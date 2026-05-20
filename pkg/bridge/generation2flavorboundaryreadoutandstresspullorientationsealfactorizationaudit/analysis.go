// Package generation2flavorboundaryreadoutandstresspullorientationsealfactorizationaudit implements
// Gate 787: Flavor-Boundary Readout and Stress-Pull Orientation Seal Factorization Audit.
//
// Gate 786 showed that the boundary pair does not natively source the full
// exterior response package needed to represent F_wall_3_red. Gate 787 factors
// that composite package into sharper missing sub-objects: a degree rule, a
// flavor-boundary readout, and a stress-pull orientation/sign seal. This gate is
// forensic only: it does not derive scalar runtime lambda, Higgs pole mass,
// Yukawa operators, PMNS, CKM, flavor hierarchy, G_F, VEV, or a native
// HistoryLoopUnit theorem.
package generation2flavorboundaryreadoutandstresspullorientationsealfactorizationaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE787-FLAVOR-BOUNDARY-READOUT-STRESS-PULL-ORIENTATION-SEAL-FACTORIZATION-AUDIT"

	StatusGate786BoundaryPairReadoutNaturalityInherited = "PASS_GATE786_BOUNDARY_PAIR_READOUT_NATURALITY_INHERITED"
	StatusResponsePackageFactorizedIntoSubseals         = "PASS_RESPONSE_PACKAGE_FACTORIZED_INTO_SUBSEALS"
	StatusDegreeRuleAudited                             = "PASS_DEGREE_RULE_AUDITED"
	StatusDegreeOneFlavorBoundaryReadoutAudited         = "PASS_DEGREE_ONE_FLAVOR_BOUNDARY_READOUT_AUDITED"
	StatusBoundaryOnlyDegreeOneCandidatesAudited        = "PASS_BOUNDARY_ONLY_DEGREE_ONE_CANDIDATES_AUDITED"
	StatusBoundaryPartOfKappaERedAudited                = "PASS_BOUNDARY_PART_OF_KAPPA_E_RED_AUDITED"
	StatusDegreeTwoStressPullOrientationAudited         = "PASS_DEGREE_TWO_STRESS_PULL_ORIENTATION_AUDITED"
	StatusResponsePackageMinimalFactorClassification    = "PASS_RESPONSE_PACKAGE_MINIMAL_FACTOR_CLASSIFICATION_RECORDED"
	StatusResponsePackageRuntimeTargetAbsenceAudited    = "PASS_RESPONSE_PACKAGE_FORMULA_LEVEL_RUNTIME_TARGET_ABSENCE_AUDITED"
	StatusStatusPropagationRecorded                     = "PASS_STATUS_PROPAGATION_RECORDED"
	StatusPhysicalFirewallsEnforced                     = "PASS_PHYSICAL_FIREWALLS_ENFORCED"

	StatusBoundaryResponsePackageThreeMissingSubobjects = "CONDITIONAL_SUPPORT_BOUNDARY_RESPONSE_PACKAGE_HAS_THREE_NONTRIVIAL_MISSING_SUBOBJECTS"
	StatusDegreeRuleExplainsCubicStopIfSupplied         = "CONDITIONAL_SUPPORT_DEGREE_RULE_EXPLAINS_CUBIC_STOP_IF_SUPPLIED"
	StatusKappaERedMixedFlavorBoundaryReadout           = "CONDITIONAL_SUPPORT_KAPPA_E_RED_IS_MIXED_FLAVOR_BOUNDARY_READOUT"
	StatusBoundaryPartOfKappaETypedByGaugePXiSSquared   = "CONDITIONAL_SUPPORT_BOUNDARY_PART_OF_KAPPA_E_RED_IS_TYPED_BY_5_OVER_3_P_XI_AND_S_SQUARED"
	StatusKappaBoundaryStrongBoundaryGaugeSourceType    = "CONDITIONAL_SUPPORT_KAPPA_BOUNDARY_HAS_STRONG_BOUNDARY_GAUGE_SOURCE_TYPE"
	StatusMainNonNativeKappaEPartFlavorOrientationInput = "CONDITIONAL_SUPPORT_MAIN_NON_NATIVE_PART_OF_KAPPA_E_RED_IS_FLAVOR_ORIENTATION_INPUT"
	Status2PMagnitudeBoundaryPairTimesK7EventSource     = "CONDITIONAL_SUPPORT_2P_MAGNITUDE_HAS_BOUNDARY_PAIR_TIMES_K7_EVENT_SOURCE"
	StatusNegativeSignStressPullOrientationCandidate    = "CONDITIONAL_SUPPORT_NEGATIVE_SIGN_HAS_STRESS_PULL_ORIENTATION_CANDIDATE"
	StatusResponsePackageReducesToThreeSharpSeals       = "CONDITIONAL_SUPPORT_RESPONSE_PACKAGE_REDUCES_TO_THREE_SHARP_SEALS"
	StatusResponsePackageFormulaLevelRuntimeIndependent = "CONDITIONAL_SUPPORT_RESPONSE_PACKAGE_REMAINS_FORMULA_LEVEL_RUNTIME_INDEPENDENT"
	StatusFWall3RedSealFactorizedResponsePackageStatus  = "CONDITIONAL_SUPPORT_F_WALL_3_RED_HAS_SEAL_FACTORIZED_RESPONSE_PACKAGE_STATUS"

	StatusNoNativeDegreeRuleTheorem                         = "FAILED_ROUTE_NO_NATIVE_DEGREE_RULE_THEOREM"
	StatusThetaExtRemainsSealed                             = "FAILED_ROUTE_THETA_EXT_REMAINS_SEALED"
	StatusProjectorIdempotenceDoesNotForceCubicStop         = "FAILED_ROUTE_PROJECTOR_IDEMPOTENCE_DOES_NOT_FORCE_CUBIC_STOP"
	StatusKappaERedNotSourcedByBoundaryPairAlone            = "FAILED_ROUTE_KAPPA_E_RED_NOT_SOURCED_BY_BOUNDARY_PAIR_ALONE"
	StatusPMNSCKMOrientationNotNative                       = "FAILED_ROUTE_PMNS_CKM_ORIENTATION_NOT_NATIVE"
	StatusNoNativeFlavorBoundaryReadoutTheorem              = "FAILED_ROUTE_NO_NATIVE_FLAVOR_BOUNDARY_READOUT_THEOREM"
	StatusNoYukawaOperatorOrEigenvalueTheorem               = "FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusSplitAxisDoesNotSourceKappaE                      = "FAILED_ROUTE_SPLIT_AXIS_DOES_NOT_SOURCE_KAPPA_E_RED"
	StatusMidpointAxisDoesNotSourceKappaE                   = "FAILED_ROUTE_MIDPOINT_AXIS_DOES_NOT_SOURCE_KAPPA_E_RED"
	StatusBoundaryAxesDoNotReplaceFlavorReadout             = "FAILED_ROUTE_BOUNDARY_AXES_DO_NOT_REPLACE_FLAVOR_READOUT"
	StatusKappaBoundaryNotFullKappaESource                  = "FAILED_ROUTE_KAPPA_BOUNDARY_NOT_FULL_KAPPA_E_SOURCE"
	StatusNoNativeCouplingHyperchargeBoundarySquareToFlavor = "FAILED_ROUTE_NO_NATIVE_COUPLING_OF_HYPERCHARGE_BOUNDARY_SQUARE_TO_FLAVOR_READOUT"
	StatusNoNativeNegativeStressPullSignTheorem             = "FAILED_ROUTE_NO_NATIVE_NEGATIVE_STRESS_PULL_SIGN_THEOREM"
	StatusMatchingSignNotNativeOrientationTheorem           = "FAILED_ROUTE_MATCHING_SIGN_IS_NOT_NATIVE_ORIENTATION_THEOREM"
	StatusResponsePackageNotNativeAfterFactoring            = "FAILED_ROUTE_RESPONSE_PACKAGE_NOT_NATIVE_AFTER_FACTORING"
	StatusResponsePackageNotTheoremLevelIndependent         = "FAILED_ROUTE_RESPONSE_PACKAGE_NOT_THEOREM_LEVEL_INDEPENDENT"
	StatusFWall3RedNotNativeBoundaryResponseTheorem         = "FAILED_ROUTE_F_WALL_3_RED_NOT_NATIVE_BOUNDARY_RESPONSE_THEOREM"
	StatusCHistoryNotFullIndependentPredictionComponent     = "FAILED_ROUTE_C_HISTORY_NOT_FULL_INDEPENDENT_PREDICTION_COMPONENT"
	StatusTreeProxyNotPoleMass                              = "FAILED_ROUTE_TREE_PROXY_NOT_POLE_MASS"
	StatusFirewallPreservedGate787                          = "FIREWALL_PRESERVED_GATE787_FLAVOR_BOUNDARY_READOUT_STRESS_PULL_ORIENTATION_BOUNDARY"
)

const (
	pK7Snapshot        = 7.0 / 72.0
	sSplitSnapshot     = 0.0012924448188162962
	xiBoundarySnapshot = 0.0503471644870914
	kappaERedSnapshot  = 0.005503554218475772
	fWall3Snapshot     = 0.00012565521035653708
)

type Gate786Inheritance struct {
	Inherited      bool
	PackageSeal    string
	PriorNative    bool
	CurrentProblem string
	Verdict        string
}

type ResponsePackageFactorization struct {
	Recorded                  bool
	CompositeSeal             string
	Subseals                  []string
	DegreeZeroCanonical       bool
	ThreeNontrivialSubobjects bool
	Verdict                   string
}

type DegreeRuleAudit struct {
	Audited                 bool
	Rule                    string
	ExplainsCubicStop       bool
	Native                  bool
	ThetaExtSealed          bool
	ProjectorPowersContinue bool
	Verdict                 string
}

type FlavorBoundaryReadoutAudit struct {
	Audited             bool
	KappaERed           float64
	KappaOrient         float64
	KappaBoundary       float64
	Formula             string
	TermTypes           []string
	MixedReadout        bool
	NativeFromBoundary  bool
	NativeFlavorTheorem bool
	Verdict             string
}

type BoundaryAxisCandidatesAudit struct {
	Audited                   bool
	SplitAxisCandidate        bool
	SplitAxis                 string
	SplitAxisSourcesKappaE    bool
	MidpointAxisCandidate     bool
	MidpointAxis              string
	MidpointAxisSourcesKappaE bool
	BoundaryAxesReplaceFlavor bool
	Verdict                   string
}

type BoundaryKappaPartAudit struct {
	Audited            bool
	KappaBoundary      float64
	Coefficient        float64
	HyperchargeFactor  float64
	BoundaryStressTerm float64
	StrongSourceType   bool
	FullKappaESource   bool
	MainNonNativePart  string
	NativeCoupling     bool
	Verdict            string
}

type StressPullOrientationAudit struct {
	Audited              bool
	Magnitude            float64
	MagnitudeSource      string
	NegativeSign         bool
	CandidateSignSources []string
	NativeNegativeSign   bool
	MatchingSignNative   bool
	Verdict              string
}

type MinimalFactorClassification struct {
	Recorded bool
	Seals    []string
	Mapping  map[string]string
	Native   bool
	Verdict  string
}

type RuntimeIndependenceAudit struct {
	Audited                  bool
	ForbiddenDirectVariables []string
	ContainsForbidden        bool
	FormulaLevelIndependent  bool
	TheoremLevelIndependent  bool
	Verdict                  string
}

type StatusPropagation struct {
	Recorded    bool
	FWall3      string
	KappaLambda string
	CHistory    string
	CHiggs      string
	Verdict     string
}

type Firewalls struct {
	Enforced                        bool
	KappaENativeFlavorTheorem       bool
	KappaBoundaryFullFlavorTheorem  bool
	KappaOrientPMNSCKMTheorem       bool
	SplitMidpointAxesKappaETheorem  bool
	TwoPMagnitudeFullSignTheorem    bool
	NegativeCubicSignNative         bool
	ResponsePackageNativeGenerating bool
	FWallNative                     bool
	KappaLambdaNative               bool
	CHistoryIndependent             bool
	TreeProxyPoleMass               bool
	Verdict                         string
}

type Audit struct {
	Gate786        Gate786Inheritance
	Factorization  ResponsePackageFactorization
	DegreeRule     DegreeRuleAudit
	FlavorReadout  FlavorBoundaryReadoutAudit
	Axis           BoundaryAxisCandidatesAudit
	KappaBoundary  BoundaryKappaPartAudit
	StressPull     StressPullOrientationAudit
	Minimal        MinimalFactorClassification
	Runtime        RuntimeIndependenceAudit
	Propagation    StatusPropagation
	Firewalls      Firewalls
	Truth          string
	FinalStatement string
}

var (
	defaultOnce  sync.Once
	defaultAudit Audit
	defaultErr   error
)

func BuildDefault() (Audit, error) {
	defaultOnce.Do(func() {
		defaultAudit, defaultErr = build()
	})
	return defaultAudit, defaultErr
}

func build() (Audit, error) {
	if !finite(pK7Snapshot, sSplitSnapshot, xiBoundarySnapshot, kappaERedSnapshot) {
		return Audit{}, fmt.Errorf("non-finite Gate787 ledger")
	}
	kappaBoundary := (-5.0/3.0 + xiBoundarySnapshot*pK7Snapshot) * sSplitSnapshot * sSplitSnapshot
	kappaOrient := kappaERedSnapshot - kappaBoundary
	magnitude := 2 * pK7Snapshot

	a := Audit{}
	a.Gate786 = Gate786Inheritance{
		Inherited:      true,
		PackageSeal:    "BoundaryExteriorResponsePackageSeal = Theta_ext + chi_ext + degree-one readout + ordered boundary orientation + negative stress-pull sign",
		PriorNative:    false,
		CurrentProblem: "factor the composite response-package seal into sharper source objects",
		Verdict:        StatusGate786BoundaryPairReadoutNaturalityInherited,
	}
	a.Factorization = ResponsePackageFactorization{
		Recorded:                  true,
		CompositeSeal:             "BoundaryExteriorResponsePackageSeal",
		Subseals:                  []string{"DegreeRuleSeal", "FlavorBoundaryReadoutSeal", "BoundaryStressPullOrientationSeal"},
		DegreeZeroCanonical:       true,
		ThreeNontrivialSubobjects: true,
		Verdict:                   StatusResponsePackageReducesToThreeSharpSeals,
	}
	a.DegreeRule = DegreeRuleAudit{
		Audited:                 true,
		Rule:                    "Theta_ext(M_n) in Lambda^(n-1)B_boundary; Theta_ext(M_n>=4)=0 if Lambda^3 B_boundary=0",
		ExplainsCubicStop:       true,
		Native:                  false,
		ThetaExtSealed:          true,
		ProjectorPowersContinue: true,
		Verdict:                 StatusNoNativeDegreeRuleTheorem,
	}
	a.FlavorReadout = FlavorBoundaryReadoutAudit{
		Audited:       true,
		KappaERed:     kappaERedSnapshot,
		KappaOrient:   kappaOrient,
		KappaBoundary: kappaBoundary,
		Formula:       "kappa_e_red = sin^2(theta13)/4 - J_CKM - (5/3)s^2 + xi_boundary p s^2",
		TermTypes: []string{
			"sin^2(theta13)/4: PMNS reactor leakage / flavor orientation candidate",
			"-J_CKM: CKM orientation correction candidate",
			"-(5/3)s^2: hypercharge-normalized boundary-square correction",
			"+xi_boundary p s^2: boundary-stress-weighted K7 second raw moment correction",
		},
		MixedReadout:        true,
		NativeFromBoundary:  false,
		NativeFlavorTheorem: false,
		Verdict:             StatusKappaERedMixedFlavorBoundaryReadout,
	}
	a.Axis = BoundaryAxisCandidatesAudit{
		Audited:                   true,
		SplitAxisCandidate:        true,
		SplitAxis:                 "beta_s ~ b_R - b_lambda, matching s=(R3-1)-|lambda|",
		SplitAxisSourcesKappaE:    false,
		MidpointAxisCandidate:     true,
		MidpointAxis:              "beta_xi ~ b_lambda + b_R, matching xi_boundary=0.5(|lambda|+R3-1)",
		MidpointAxisSourcesKappaE: false,
		BoundaryAxesReplaceFlavor: false,
		Verdict:                   StatusBoundaryAxesDoNotReplaceFlavorReadout,
	}
	a.KappaBoundary = BoundaryKappaPartAudit{
		Audited:            true,
		KappaBoundary:      kappaBoundary,
		Coefficient:        -5.0/3.0 + xiBoundarySnapshot*pK7Snapshot,
		HyperchargeFactor:  5.0 / 3.0,
		BoundaryStressTerm: xiBoundarySnapshot * pK7Snapshot,
		StrongSourceType:   true,
		FullKappaESource:   false,
		MainNonNativePart:  "kappa_orient = sin^2(theta13)/4 - J_CKM",
		NativeCoupling:     false,
		Verdict:            StatusKappaBoundaryStrongBoundaryGaugeSourceType,
	}
	a.StressPull = StressPullOrientationAudit{
		Audited:              true,
		Magnitude:            magnitude,
		MagnitudeSource:      "2p = dim(B_boundary) * p_K7 = 7/36",
		NegativeSign:         true,
		CandidateSignSources: []string{"ordered boundary orientation", "opposite scalar-wall/gauge-stress orientation", "restorative stress-pull convention", "matching sign from residual compression"},
		NativeNegativeSign:   false,
		MatchingSignNative:   false,
		Verdict:              StatusNegativeSignStressPullOrientationCandidate,
	}
	a.Minimal = MinimalFactorClassification{
		Recorded: true,
		Seals:    []string{"DegreeRuleSeal", "FlavorBoundaryReadoutSeal", "BoundaryStressPullOrientationSeal"},
		Mapping: map[string]string{
			"DegreeRuleSeal":                    "Theta_ext(M_n) in Lambda^(n-1)B_boundary; explains cubic stop if supplied",
			"FlavorBoundaryReadoutSeal":         "chi_ext(beta_B)=kappa_e_red",
			"BoundaryStressPullOrientationSeal": "ordered omega_B and chi_ext(omega_B)=-2p, especially negative sign",
		},
		Native:  false,
		Verdict: StatusResponsePackageReducesToThreeSharpSeals,
	}
	a.Runtime = RuntimeIndependenceAudit{
		Audited:                  true,
		ForbiddenDirectVariables: []string{"lambda_runtime", "lambda_runtime_eff", "m_H_tree", "m_H_pole", "C_Higgs", "G_F", "v"},
		ContainsForbidden:        false,
		FormulaLevelIndependent:  true,
		TheoremLevelIndependent:  false,
		Verdict:                  StatusResponsePackageFormulaLevelRuntimeIndependent,
	}
	a.Propagation = StatusPropagation{
		Recorded:    true,
		FWall3:      "Level B+ seal-factorized exterior response package; formula-level runtime independent; not native",
		KappaLambda: "Level B formula-independent boundary-flavor complement; not native scalar matching theorem",
		CHistory:    "Level B semi-independent History correction; not full independent prediction component",
		CHiggs:      "not Level C until C_History, N_eff, and scale dependencies are theorem-sourced",
		Verdict:     StatusFWall3RedSealFactorizedResponsePackageStatus,
	}
	a.Firewalls = Firewalls{
		Enforced:                        true,
		KappaENativeFlavorTheorem:       false,
		KappaBoundaryFullFlavorTheorem:  false,
		KappaOrientPMNSCKMTheorem:       false,
		SplitMidpointAxesKappaETheorem:  false,
		TwoPMagnitudeFullSignTheorem:    false,
		NegativeCubicSignNative:         false,
		ResponsePackageNativeGenerating: false,
		FWallNative:                     false,
		KappaLambdaNative:               false,
		CHistoryIndependent:             false,
		TreeProxyPoleMass:               false,
		Verdict:                         StatusFirewallPreservedGate787,
	}
	a.Truth = "Gate 787 factors the exterior response package into degree-rule, flavor-boundary readout, and stress-pull orientation seals without promoting any bridge coefficient to a native law."
	a.FinalStatement = "Gate 787 does not make the exterior response package native. It reduces the package into three sharper missing subobjects: DegreeRuleSeal, FlavorBoundaryReadoutSeal, and BoundaryStressPullOrientationSeal. It further shows that the boundary-only part of kappa_e_red is strongly typed, while the main non-native part is the flavor-orientation input sin^2(theta13)/4 - J_CKM. The next bottleneck is the FlavorBoundaryReadoutSeal, especially whether the orientation term sin^2(theta13)/4 - J_CKM can be sourced without empirical PMNS/CKM input."
	return a, nil
}

func Statuses() []string {
	return []string{
		StatusGate786BoundaryPairReadoutNaturalityInherited,
		StatusResponsePackageFactorizedIntoSubseals,
		StatusDegreeRuleAudited,
		StatusDegreeOneFlavorBoundaryReadoutAudited,
		StatusBoundaryOnlyDegreeOneCandidatesAudited,
		StatusBoundaryPartOfKappaERedAudited,
		StatusDegreeTwoStressPullOrientationAudited,
		StatusResponsePackageMinimalFactorClassification,
		StatusResponsePackageRuntimeTargetAbsenceAudited,
		StatusStatusPropagationRecorded,
		StatusPhysicalFirewallsEnforced,
		StatusBoundaryResponsePackageThreeMissingSubobjects,
		StatusDegreeRuleExplainsCubicStopIfSupplied,
		StatusKappaERedMixedFlavorBoundaryReadout,
		StatusBoundaryPartOfKappaETypedByGaugePXiSSquared,
		StatusKappaBoundaryStrongBoundaryGaugeSourceType,
		StatusMainNonNativeKappaEPartFlavorOrientationInput,
		Status2PMagnitudeBoundaryPairTimesK7EventSource,
		StatusNegativeSignStressPullOrientationCandidate,
		StatusResponsePackageReducesToThreeSharpSeals,
		StatusResponsePackageFormulaLevelRuntimeIndependent,
		StatusFWall3RedSealFactorizedResponsePackageStatus,
		StatusNoNativeDegreeRuleTheorem,
		StatusThetaExtRemainsSealed,
		StatusProjectorIdempotenceDoesNotForceCubicStop,
		StatusKappaERedNotSourcedByBoundaryPairAlone,
		StatusPMNSCKMOrientationNotNative,
		StatusNoNativeFlavorBoundaryReadoutTheorem,
		StatusNoYukawaOperatorOrEigenvalueTheorem,
		StatusSplitAxisDoesNotSourceKappaE,
		StatusMidpointAxisDoesNotSourceKappaE,
		StatusBoundaryAxesDoNotReplaceFlavorReadout,
		StatusKappaBoundaryNotFullKappaESource,
		StatusNoNativeCouplingHyperchargeBoundarySquareToFlavor,
		StatusNoNativeNegativeStressPullSignTheorem,
		StatusMatchingSignNotNativeOrientationTheorem,
		StatusResponsePackageNotNativeAfterFactoring,
		StatusResponsePackageNotTheoremLevelIndependent,
		StatusFWall3RedNotNativeBoundaryResponseTheorem,
		StatusCHistoryNotFullIndependentPredictionComponent,
		StatusTreeProxyNotPoleMass,
		StatusFirewallPreservedGate787,
	}
}

func FormatFactorization(f ResponsePackageFactorization) string {
	return fmt.Sprintf("%s -> %s", f.CompositeSeal, strings.Join(f.Subseals, ", "))
}

func FormatFlavorReadout(f FlavorBoundaryReadoutAudit) string {
	return fmt.Sprintf("kappa_e_red=%.16g = kappa_orient %.16g + kappa_boundary %.16g", f.KappaERed, f.KappaOrient, f.KappaBoundary)
}

func FormatKappaBoundary(k BoundaryKappaPartAudit) string {
	return fmt.Sprintf("kappa_boundary=[-5/3 + xi*p]s^2 = %.16g; coefficient=%.16g", k.KappaBoundary, k.Coefficient)
}

func FormatStressPull(s StressPullOrientationAudit) string {
	return fmt.Sprintf("magnitude %.16g from %s; sign native=%v", s.Magnitude, s.MagnitudeSource, s.NativeNegativeSign)
}

func FormatPropagation(p StatusPropagation) string {
	return strings.Join([]string{p.FWall3, p.KappaLambda, p.CHistory, p.CHiggs}, " | ")
}

func containsAll(haystack, needles []string) bool {
	joined := strings.Join(haystack, "\n")
	for _, needle := range needles {
		if !strings.Contains(joined, needle) {
			return false
		}
	}
	return true
}

func closeRel(got, want, tol float64) bool {
	if math.IsNaN(got) || math.IsNaN(want) || math.IsInf(got, 0) || math.IsInf(want, 0) {
		return false
	}
	d := math.Abs(got - want)
	if math.Abs(want) < 1 {
		return d <= tol
	}
	return d/math.Abs(want) <= tol
}

func finite(vals ...float64) bool {
	for _, v := range vals {
		if math.IsNaN(v) || math.IsInf(v, 0) {
			return false
		}
	}
	return true
}
