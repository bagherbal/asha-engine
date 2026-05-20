// Package generation2generationmixingoperatorsourceandflavororientationreadoutsealaudit implements
// Gate 789: Generation Mixing Operator Source and FlavorOrientationReadoutSeal Audit.
//
// Gate 788 isolated kappa_orient = sin^2(theta13)/4 - J_CKM as the true
// non-native flavor-orientation obstruction inside the scalar-Higgs bridge.
// Gate 789 audits the exact typed object required before theta13 or J_CKM can
// be sourced natively. This gate is forensic only: it does not derive Yukawa
// eigenvalues, PMNS, CKM, flavor hierarchy, scalar runtime lambda, Higgs pole
// mass, G_F, VEV, or a native HistoryLoopUnit theorem.
package generation2generationmixingoperatorsourceandflavororientationreadoutsealaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE789-GENERATION-MIXING-OPERATOR-SOURCE-FLAVOR-ORIENTATION-READOUT-SEAL-AUDIT"

	StatusGate788Inherited                    = "PASS_GATE788_FLAVOR_ORIENTATION_READOUT_INHERITED"
	StatusKappaOrientFlavorBottleneck         = "PASS_KAPPA_ORIENT_SELECTED_AS_CURRENT_FLAVOR_BOTTLENECK"
	StatusRequiredMixingObjectsDefined        = "PASS_REQUIRED_GENERATION_MIXING_OBJECTS_DEFINED"
	StatusYukawaTracePairAudited              = "PASS_YUKAWA_TRACE_PAIR_SOURCE_AUDITED"
	StatusYukawaSingularLedgerAudited         = "PASS_YUKAWA_SINGULAR_VALUE_LEDGER_AUDITED"
	StatusFiniteSpectralTripleAudited         = "PASS_FINITE_SPECTRAL_TRIPLE_FLAVOR_SOURCE_AUDITED"
	StatusK7HodgePolarityAudited              = "PASS_K7_HODGE_POLARITY_SOURCE_AUDITED"
	StatusFockProjectiveSelectorAudited       = "PASS_FOCK_PROJECTIVE_SELECTOR_SOURCE_AUDITED"
	StatusTrialityGenerationCarrierAudited    = "PASS_TRIALITY_GENERATION_CARRIER_CANDIDATE_AUDITED"
	StatusBoundaryDataAudited                 = "PASS_BOUNDARY_DATA_SOURCE_AUDITED"
	StatusGenerationMixingOperatorSealDefined = "PASS_GENERATION_MIXING_OPERATOR_SEAL_DEFINED"
	StatusGenerationMixingSealMinimality      = "PASS_GENERATION_MIXING_OPERATOR_SEAL_MINIMALITY_AUDITED"
	StatusGenerationMixingRuntimeAbsence      = "PASS_GENERATION_MIXING_SEAL_RUNTIME_TARGET_ABSENCE_AUDITED"
	StatusStatusPropagationRecorded           = "PASS_STATUS_PROPAGATION_RECORDED"
	StatusBranchDecisionRecorded              = "PASS_BRANCH_DECISION_RECORDED"
	StatusPhysicalFirewallsEnforced           = "PASS_PHYSICAL_FIREWALLS_ENFORCED"

	StatusTheta13AndJCKMRequireSectorMisalignment = "CONDITIONAL_SUPPORT_THETA13_AND_J_CKM_REQUIRE_SECTOR_MISALIGNMENT_OPERATORS"
	StatusABNEffAggregateYukawaParticipation      = "CONDITIONAL_SUPPORT_A_B_N_EFF_SOURCE_AGGREGATE_YUKAWA_PARTICIPATION"
	StatusProjectiveSelectorFutureCandidate       = "CONDITIONAL_SUPPORT_PROJECTIVE_SELECTOR_GEOMETRY_IS_A_FUTURE_GENERATION_CARRIER_CANDIDATE"
	StatusTrialityThreefoldRelevantCandidate      = "CONDITIONAL_SUPPORT_TRIALITY_OR_THREEFOLD_STRUCTURE_IS_RELEVANT_GENERATION_CARRIER_CANDIDATE"
	StatusBoundaryDataSmallFlavorCorrection       = "CONDITIONAL_SUPPORT_BOUNDARY_DATA_SUPPLIES_SMALL_CORRECTION_TO_FLAVOR_READOUT"
	StatusFlavorOrientationRequiresMixingSeal     = "CONDITIONAL_SUPPORT_FLAVOR_ORIENTATION_READOUT_REQUIRES_GENERATION_MIXING_OPERATOR_SEAL"
	StatusFlavorOrientationRuntimeIndependent     = "CONDITIONAL_SUPPORT_FLAVOR_ORIENTATION_READOUT_IS_RUNTIME_TARGET_INDEPENDENT"
	StatusKappaOrientReadoutOfMixingSeal          = "CONDITIONAL_SUPPORT_KAPPA_ORIENT_IS_READOUT_OF_GENERATION_MIXING_OPERATOR_SEAL"

	StatusNoNativeUPMNSOrVCKM                       = "FAILED_ROUTE_CURRENT_LEDGER_DOES_NOT_YET_SUPPLY_NATIVE_U_PMNS_OR_V_CKM"
	StatusNEffDoesNotDeterminePMNSOrCKM             = "FAILED_ROUTE_N_EFF_DOES_NOT_DETERMINE_PMNS_OR_CKM_MIXING"
	StatusTraceInvariantsNoEigenvectorMisalignment  = "FAILED_ROUTE_TRACE_INVARIANTS_DO_NOT_SUPPLY_EIGENVECTOR_MISALIGNMENT"
	StatusSingularValuesDoNotDeterminePMNSCKM       = "FAILED_ROUTE_SINGULAR_VALUES_ALONE_DO_NOT_DETERMINE_PMNS_CKM"
	StatusNoNativeYukawaEigenvectorOrientation      = "FAILED_ROUTE_NO_NATIVE_YUKAWA_EIGENVECTOR_ORIENTATION_THEOREM"
	StatusFiniteTripleNoGenerationMixingOperator    = "FAILED_ROUTE_FINITE_TRIPLE_DOES_NOT_YET_SOURCE_GENERATION_MIXING_OPERATOR"
	StatusK7PolarityNoGenerationMixingOperator      = "FAILED_ROUTE_K7_HODGE_POLARITY_DOES_NOT_DEFINE_GENERATION_MIXING_OPERATOR"
	StatusK7QuarterNoTheta13                        = "FAILED_ROUTE_K7_PLUS_QUARTER_WEIGHT_DOES_NOT_DERIVE_THETA13"
	StatusNoSelectorToPMNSCKMMap                    = "FAILED_ROUTE_NO_TYPED_SELECTOR_TO_PMNS_CKM_MAP"
	StatusThreefoldAloneNoPMNSCKM                   = "FAILED_ROUTE_THREEFOLD_STRUCTURE_ALONE_DOES_NOT_DERIVE_PMNS_CKM"
	StatusNoSectorMisalignmentOperator              = "FAILED_ROUTE_NO_SECTOR_MISALIGNMENT_OPERATOR_CERTIFIED"
	StatusBoundaryPairNoFlavorMixing                = "FAILED_ROUTE_BOUNDARY_PAIR_DOES_NOT_DERIVE_FLAVOR_MIXING"
	StatusFlavorOrientationNotNativeWithoutMixingOp = "FAILED_ROUTE_FLAVOR_ORIENTATION_READOUT_NOT_NATIVE_WITHOUT_GENERATION_MIXING_OPERATOR"
	StatusFlavorOrientationNotTheoremIndependent    = "FAILED_ROUTE_FLAVOR_ORIENTATION_READOUT_NOT_THEOREM_LEVEL_INDEPENDENT"
	StatusCHistoryNotFullPrediction                 = "FAILED_ROUTE_C_HISTORY_NOT_FULL_INDEPENDENT_PREDICTION_COMPONENT"
	StatusCHiggsNotLevelC                           = "FAILED_ROUTE_C_HIGGS_NOT_LEVEL_C_PREDICTION"
	StatusTreeProxyNotPoleMass                      = "FAILED_ROUTE_TREE_PROXY_NOT_POLE_MASS"
	StatusFirewallPreservedGate789                  = "FIREWALL_PRESERVED_GATE789_GENERATION_MIXING_OPERATOR_SOURCE_BOUNDARY"
)

const (
	nEffSnapshot          = 3.0023273474722147
	kappaOrientSnapshot   = 0.00550633006471245
	kappaBoundarySnapshot = -2.775846236678231e-6
	kappaERedSnapshot     = 0.005503554218475772
)

type Gate788Inheritance struct {
	Inherited         bool
	KappaOrientFocus  bool
	Formula           string
	PMNSTyping        string
	CKMTyping         string
	NegativeSign      string
	Verdict           string
	BottleneckVerdict string
}

type RequiredGenerationMixingObjects struct {
	Defined                 bool
	RequiredObjects         []string
	LeptonMisalignment      string
	LeptonReadout           string
	QuarkMisalignment       string
	QuarkReadout            string
	SectorMisalignmentNeed  bool
	NativeUPMNSOrVCKMExists bool
	Verdict                 string
}

type YukawaTracePairAudit struct {
	Audited                         bool
	A                               string
	B                               string
	NEff                            float64
	AggregateParticipation          bool
	DeterminesPMNSOrCKM             bool
	SuppliesEigenvectorMisalignment bool
	Verdict                         string
}

type YukawaSingularValueLedgerAudit struct {
	Audited                       bool
	SingularValuesCanSourceTraces bool
	SingularValuesDetermineMixing bool
	NativeEigenvectorOrientation  bool
	Verdict                       string
}

type FiniteSpectralTripleFlavorAudit struct {
	Audited                         bool
	AllowedYukawaEdgeShapes         bool
	GenerationMixingOperatorSourced bool
	Verdict                         string
}

type K7HodgePolarityAudit struct {
	Audited                         bool
	Split                           string
	SelectorResonance               bool
	DefinesGenerationMixingOperator bool
	QuarterWeightDerivesTheta13     bool
	Verdict                         string
}

type FockProjectiveSelectorAudit struct {
	Audited                   bool
	Patterns                  []string
	FutureGenerationCandidate bool
	TypedSelectorToPMNSCKMMap bool
	Verdict                   string
}

type TrialityGenerationCarrierAudit struct {
	Audited                         bool
	ThreefoldRelevantCandidate      bool
	SuppliesSectorOperators         bool
	SuppliesRelativeOrientations    bool
	SuppliesPhaseData               bool
	SuppliesMixingReadoutMaps       bool
	SectorMisalignmentOperatorFound bool
	Verdict                         string
}

type BoundaryDataSourceAudit struct {
	Audited                  bool
	Coordinates              []string
	SmallCorrectionToReadout bool
	DerivesFlavorMixing      bool
	Verdict                  string
}

type GenerationMixingOperatorSeal struct {
	Defined    bool
	Name       string
	Components []string
	Readout    string
	Native     bool
	Verdict    string
}

type MinimalityAudit struct {
	Audited       bool
	RemoveEffects map[string]string
	Minimal       bool
	Verdict       string
}

type RuntimeIndependenceAudit struct {
	Audited                  bool
	ForbiddenDirectVariables []string
	ContainsForbidden        bool
	RuntimeTargetIndependent bool
	TheoremLevelIndependent  bool
	Verdict                  string
}

type StatusPropagation struct {
	Recorded    bool
	KappaOrient string
	KappaERed   string
	FWall3      string
	KappaLambda string
	CHistory    string
	CHiggs      string
	Verdict     string
}

type BranchDecision struct {
	Recorded      bool
	SuccessBranch string
	FailureBranch string
	Selected      string
	Reason        string
	Verdict       string
}

type Firewalls struct {
	Enforced                          bool
	NEffPMNSCKMTheorem                bool
	YukawaSingularValuesMixingTheorem bool
	K7PolarityMixingTheorem           bool
	RadialQuarterTheta13Theorem       bool
	ProjectiveSelectorPMNSCKMTheorem  bool
	TrialityPMNSCKMTheorem            bool
	BoundaryPairFlavorMixingTheorem   bool
	KappaOrientNativeFlavorTheorem    bool
	FlavorOrientationSealNative       bool
	TreeProxyPoleMass                 bool
	Verdict                           string
}

type Audit struct {
	Gate788        Gate788Inheritance
	Required       RequiredGenerationMixingObjects
	YukawaTrace    YukawaTracePairAudit
	SingularLedger YukawaSingularValueLedgerAudit
	FiniteTriple   FiniteSpectralTripleFlavorAudit
	K7Polarity     K7HodgePolarityAudit
	FockSelector   FockProjectiveSelectorAudit
	Triality       TrialityGenerationCarrierAudit
	BoundaryData   BoundaryDataSourceAudit
	Seal           GenerationMixingOperatorSeal
	Minimality     MinimalityAudit
	Runtime        RuntimeIndependenceAudit
	Propagation    StatusPropagation
	Branch         BranchDecision
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
	if !closeRel(kappaOrientSnapshot+kappaBoundarySnapshot, kappaERedSnapshot, 1e-15) {
		return Audit{}, fmt.Errorf("kappa_orient + kappa_boundary does not reconstruct kappa_e_red")
	}
	forbidden := []string{"lambda_runtime", "lambda_runtime_eff", "m_H_tree", "m_H_pole", "C_Higgs", "G_F", "v"}
	sealFormula := "sin^2(theta13)/4 - J_CKM"
	containsForbidden := containsAny(sealFormula, forbidden)

	return Audit{
		Gate788: Gate788Inheritance{
			Inherited:         true,
			KappaOrientFocus:  true,
			Formula:           sealFormula,
			PMNSTyping:        "sin^2(theta13)/4 = PMNS reactor leakage candidate",
			CKMTyping:         "J_CKM = CKM Jarlskog oriented flavor-area correction candidate",
			NegativeSign:      "minus sign = orientation-subtraction candidate between PMNS leakage and CKM area",
			Verdict:           StatusGate788Inherited,
			BottleneckVerdict: StatusKappaOrientFlavorBottleneck,
		},
		Required: RequiredGenerationMixingObjects{
			Defined: true,
			RequiredObjects: []string{
				"generation carrier G_gen",
				"sector Yukawa or mass operators on G_gen",
				"typed diagonalization maps",
				"misalignment unitaries between sectors",
				"readout maps theta13 and J_CKM",
				"orientation/sign convention explaining sin^2(theta13)/4 - J_CKM",
			},
			LeptonMisalignment:      "U_PMNS = U_e^† U_nu",
			LeptonReadout:           "sin^2(theta13) = |(U_PMNS)13|^2",
			QuarkMisalignment:       "V_CKM = U_u^† U_d",
			QuarkReadout:            "J_CKM = Im(V_us V_cb V_ub^* V_cs^*) or equivalent invariant",
			SectorMisalignmentNeed:  true,
			NativeUPMNSOrVCKMExists: false,
			Verdict:                 StatusTheta13AndJCKMRequireSectorMisalignment,
		},
		YukawaTrace: YukawaTracePairAudit{
			Audited:                         true,
			A:                               "a = quadratic Yukawa trace",
			B:                               "b = quartic Yukawa trace",
			NEff:                            nEffSnapshot,
			AggregateParticipation:          true,
			DeterminesPMNSOrCKM:             false,
			SuppliesEigenvectorMisalignment: false,
			Verdict:                         StatusABNEffAggregateYukawaParticipation,
		},
		SingularLedger: YukawaSingularValueLedgerAudit{
			Audited:                       true,
			SingularValuesCanSourceTraces: true,
			SingularValuesDetermineMixing: false,
			NativeEigenvectorOrientation:  false,
			Verdict:                       StatusSingularValuesDoNotDeterminePMNSCKM,
		},
		FiniteTriple: FiniteSpectralTripleFlavorAudit{
			Audited:                         true,
			AllowedYukawaEdgeShapes:         true,
			GenerationMixingOperatorSourced: false,
			Verdict:                         StatusFiniteTripleNoGenerationMixingOperator,
		},
		K7Polarity: K7HodgePolarityAudit{
			Audited:                         true,
			Split:                           "K7 = K7+ ⊕ K7-, dim K7+=4, dim K7-=3",
			SelectorResonance:               true,
			DefinesGenerationMixingOperator: false,
			QuarterWeightDerivesTheta13:     false,
			Verdict:                         StatusK7PolarityNoGenerationMixingOperator,
		},
		FockSelector: FockProjectiveSelectorAudit{
			Audited:                   true,
			Patterns:                  []string{"4 = 1 + 3", "CP3", "projective selector patterns"},
			FutureGenerationCandidate: true,
			TypedSelectorToPMNSCKMMap: false,
			Verdict:                   StatusProjectiveSelectorFutureCandidate,
		},
		Triality: TrialityGenerationCarrierAudit{
			Audited:                         true,
			ThreefoldRelevantCandidate:      true,
			SuppliesSectorOperators:         false,
			SuppliesRelativeOrientations:    false,
			SuppliesPhaseData:               false,
			SuppliesMixingReadoutMaps:       false,
			SectorMisalignmentOperatorFound: false,
			Verdict:                         StatusTrialityThreefoldRelevantCandidate,
		},
		BoundaryData: BoundaryDataSourceAudit{
			Audited:                  true,
			Coordinates:              []string{"s", "xi_boundary", "lambda(Lambda12)", "R3-1"},
			SmallCorrectionToReadout: true,
			DerivesFlavorMixing:      false,
			Verdict:                  StatusBoundaryDataSmallFlavorCorrection,
		},
		Seal: GenerationMixingOperatorSeal{
			Defined: true,
			Name:    "GenerationMixingOperatorSeal",
			Components: []string{
				"G_gen",
				"Y_u, Y_d, Y_e, Y_nu or equivalent sector operators",
				"sector diagonalization frames",
				"U_PMNS",
				"V_CKM",
				"readout maps theta13 and J_CKM",
				"orientation/sign convention",
			},
			Readout: "FlavorOrientationReadoutSeal = Readout[GenerationMixingOperatorSeal] = sin^2(theta13)/4 - J_CKM",
			Native:  false,
			Verdict: StatusFlavorOrientationRequiresMixingSeal,
		},
		Minimality: MinimalityAudit{
			Audited: true,
			RemoveEffects: map[string]string{
				"G_gen":                        "no generation carrier",
				"sector operators":             "no flavor sector structure",
				"diagonalization frames":       "no eigenbasis comparison",
				"U_PMNS":                       "no theta13 readout",
				"V_CKM":                        "no J_CKM readout",
				"phase/orientation convention": "no signed subtraction and no CP-area orientation",
			},
			Minimal: true,
			Verdict: StatusGenerationMixingSealMinimality,
		},
		Runtime: RuntimeIndependenceAudit{
			Audited:                  true,
			ForbiddenDirectVariables: forbidden,
			ContainsForbidden:        containsForbidden,
			RuntimeTargetIndependent: !containsForbidden,
			TheoremLevelIndependent:  false,
			Verdict:                  StatusFlavorOrientationRuntimeIndependent,
		},
		Propagation: StatusPropagation{
			Recorded:    true,
			KappaOrient: "readout of GenerationMixingOperatorSeal; runtime-target independent but not native",
			KappaERed:   "mixed flavor-boundary readout: GenerationMixingOperatorSeal + strongly typed boundary correction",
			FWall3:      "Level B+ seal-factorized exterior response package; not native",
			KappaLambda: "Level B formula-independent scalar complement; not native",
			CHistory:    "Level B semi-independent History correction; not full prediction",
			CHiggs:      "still not Level C",
			Verdict:     StatusKappaOrientReadoutOfMixingSeal,
		},
		Branch: BranchDecision{
			Recorded:      true,
			SuccessBranch: "Gate 790 — Native Generation Mixing Readout Derivation Audit",
			FailureBranch: "Gate 790 — C_Higgs Dependency Freeze and Level-B Prediction Interface Audit",
			Selected:      "failure branch",
			Reason:        "current ASHA ledger does not certify native PMNS/CKM generation mixing or sector-misalignment operators",
			Verdict:       StatusBranchDecisionRecorded,
		},
		Firewalls: Firewalls{
			Enforced:                          true,
			NEffPMNSCKMTheorem:                false,
			YukawaSingularValuesMixingTheorem: false,
			K7PolarityMixingTheorem:           false,
			RadialQuarterTheta13Theorem:       false,
			ProjectiveSelectorPMNSCKMTheorem:  false,
			TrialityPMNSCKMTheorem:            false,
			BoundaryPairFlavorMixingTheorem:   false,
			KappaOrientNativeFlavorTheorem:    false,
			FlavorOrientationSealNative:       false,
			TreeProxyPoleMass:                 false,
			Verdict:                           StatusFirewallPreservedGate789,
		},
		Truth:          "Gate 789 audits the missing flavor carrier/mixing operator without promoting aggregate traces, singular values, projective resonances, triality, or empirical PMNS/CKM data into native theorems.",
		FinalStatement: "Gate 789 does not source theta13 or J_CKM natively. It identifies the exact missing object as GenerationMixingOperatorSeal: a generation carrier with sector operators, diagonalization frames, PMNS/CKM misalignment matrices, readout maps, and orientation/sign convention. The honest failure branch is to freeze kappa_orient as a flavor-orientation seal for the current scalar-Higgs bridge and audit the full Level-B C_Higgs prediction interface, unless a new native generation-mixing construction is introduced.",
	}, nil
}

func FormatSeal(s GenerationMixingOperatorSeal) string {
	return fmt.Sprintf("%s = (%s); %s", s.Name, strings.Join(s.Components, ", "), s.Readout)
}

func FormatBranch(b BranchDecision) string {
	return fmt.Sprintf("selected %s: %s", b.Selected, b.Reason)
}

func Statuses() []string {
	return []string{
		StatusGate788Inherited,
		StatusKappaOrientFlavorBottleneck,
		StatusRequiredMixingObjectsDefined,
		StatusYukawaTracePairAudited,
		StatusYukawaSingularLedgerAudited,
		StatusFiniteSpectralTripleAudited,
		StatusK7HodgePolarityAudited,
		StatusFockProjectiveSelectorAudited,
		StatusTrialityGenerationCarrierAudited,
		StatusBoundaryDataAudited,
		StatusGenerationMixingOperatorSealDefined,
		StatusGenerationMixingSealMinimality,
		StatusGenerationMixingRuntimeAbsence,
		StatusStatusPropagationRecorded,
		StatusBranchDecisionRecorded,
		StatusPhysicalFirewallsEnforced,
		StatusTheta13AndJCKMRequireSectorMisalignment,
		StatusABNEffAggregateYukawaParticipation,
		StatusProjectiveSelectorFutureCandidate,
		StatusTrialityThreefoldRelevantCandidate,
		StatusBoundaryDataSmallFlavorCorrection,
		StatusFlavorOrientationRequiresMixingSeal,
		StatusFlavorOrientationRuntimeIndependent,
		StatusKappaOrientReadoutOfMixingSeal,
		StatusNoNativeUPMNSOrVCKM,
		StatusNEffDoesNotDeterminePMNSOrCKM,
		StatusTraceInvariantsNoEigenvectorMisalignment,
		StatusSingularValuesDoNotDeterminePMNSCKM,
		StatusNoNativeYukawaEigenvectorOrientation,
		StatusFiniteTripleNoGenerationMixingOperator,
		StatusK7PolarityNoGenerationMixingOperator,
		StatusK7QuarterNoTheta13,
		StatusNoSelectorToPMNSCKMMap,
		StatusThreefoldAloneNoPMNSCKM,
		StatusNoSectorMisalignmentOperator,
		StatusBoundaryPairNoFlavorMixing,
		StatusFlavorOrientationNotNativeWithoutMixingOp,
		StatusFlavorOrientationNotTheoremIndependent,
		StatusCHistoryNotFullPrediction,
		StatusCHiggsNotLevelC,
		StatusTreeProxyNotPoleMass,
		StatusFirewallPreservedGate789,
	}
}

func containsAny(text string, needles []string) bool {
	for _, n := range needles {
		if strings.Contains(text, n) {
			return true
		}
	}
	return false
}

func containsAll(haystack []string, needles []string) bool {
	seen := map[string]bool{}
	for _, h := range haystack {
		seen[h] = true
	}
	for _, n := range needles {
		if !seen[n] {
			return false
		}
	}
	return true
}

func closeRel(a, b, tol float64) bool {
	if a == b {
		return true
	}
	if math.IsNaN(a) || math.IsNaN(b) || math.IsInf(a, 0) || math.IsInf(b, 0) {
		return false
	}
	scale := math.Max(1, math.Max(math.Abs(a), math.Abs(b)))
	return math.Abs(a-b) <= tol*scale
}
