// Package generation2flavororientationreadoutsourcemapandpmnsckmfirewallaudit implements
// Gate 788: Flavor Orientation Readout Source and PMNS-CKM Firewall Audit.
//
// Gate 787 isolated the non-native flavor-orientation component
// kappa_orient = sin^2(theta13)/4 - J_CKM. Gate 788 audits whether this
// component can be sourced from existing ASHA geometry or must remain an
// explicit flavor-orientation seal. This gate is forensic only: it does not
// derive Yukawa operators, PMNS, CKM, flavor hierarchy, scalar runtime lambda,
// Higgs pole mass, G_F, VEV, or a native HistoryLoopUnit theorem.
package generation2flavororientationreadoutsourcemapandpmnsckmfirewallaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE788-FLAVOR-ORIENTATION-READOUT-SOURCE-PMNS-CKM-FIREWALL-AUDIT"

	StatusGate787Inherited                       = "PASS_GATE787_FLAVOR_BOUNDARY_READOUT_STRESS_PULL_INHERITED"
	StatusKappaOrientMainNonNativeComponent      = "PASS_KAPPA_ORIENT_IDENTIFIED_AS_MAIN_NON_NATIVE_READOUT_COMPONENT"
	StatusKappaOrientDecomposed                  = "PASS_KAPPA_ORIENT_DECOMPOSED_INTO_PMNS_REACTOR_AND_CKM_ORIENTATION_TERMS"
	StatusPMNSReactorLeakageAudited              = "PASS_PMNS_REACTOR_LEAKAGE_TERM_AUDITED"
	StatusCKMJarlskogOrientationAudited          = "PASS_CKM_JARLSKOG_ORIENTATION_TERM_AUDITED"
	StatusBoundaryOnlyReplacementAudited         = "PASS_BOUNDARY_ONLY_REPLACEMENT_AUDITED"
	StatusExistingASHAGeometryCandidatesAudited  = "PASS_EXISTING_ASHA_GEOMETRY_SOURCE_CANDIDATES_AUDITED"
	StatusFlavorBoundaryReadoutSealRefined       = "PASS_FLAVOR_BOUNDARY_READOUT_SEAL_REFINED"
	StatusKappaOrientRuntimeTargetAbsenceAudited = "PASS_KAPPA_ORIENT_RUNTIME_TARGET_ABSENCE_AUDITED"
	StatusStatusPropagationRecorded              = "PASS_STATUS_PROPAGATION_RECORDED"
	StatusPhysicalFirewallsEnforced              = "PASS_PHYSICAL_FIREWALLS_ENFORCED"

	StatusKappaOrientFlavorOrientationReadoutShape       = "CONDITIONAL_SUPPORT_KAPPA_ORIENT_HAS_FLAVOR_ORIENTATION_READOUT_SHAPE"
	StatusOneQuarterK7PlusEventWeightResonance           = "CONDITIONAL_SUPPORT_ONE_QUARTER_HAS_EXISTING_K7_PLUS_EVENT_WEIGHT_RESONANCE"
	StatusJCKMOrientedFlavorAreaSourceType               = "CONDITIONAL_SUPPORT_J_CKM_HAS_ORIENTED_FLAVOR_AREA_SOURCE_TYPE"
	StatusNegativeCKMSignOrientationSubtractionCandidate = "CONDITIONAL_SUPPORT_NEGATIVE_CKM_SIGN_HAS_ORIENTATION_SUBTRACTION_CANDIDATE"
	StatusKappaBoundarySmallCorrectionToFlavorOrient     = "CONDITIONAL_SUPPORT_KAPPA_BOUNDARY_IS_SMALL_CORRECTION_TO_FLAVOR_ORIENTATION"
	StatusBoundaryGaugeCorrectionStronglyTyped           = "CONDITIONAL_SUPPORT_BOUNDARY_GAUGE_CORRECTION_IS_STRONGLY_SOURCE_TYPED"
	StatusFlavorOrientationTrueNonNativeObstruction      = "CONDITIONAL_SUPPORT_FLAVOR_ORIENTATION_READOUT_IS_TRUE_NON_NATIVE_OBSTRUCTION"
	StatusKappaOrientFormulaLevelRuntimeIndependent      = "CONDITIONAL_SUPPORT_KAPPA_ORIENT_IS_FORMULA_LEVEL_RUNTIME_INDEPENDENT"
	StatusKappaOrientCurrentFlavorReadoutBottleneck      = "CONDITIONAL_SUPPORT_KAPPA_ORIENT_IS_CURRENT_FLAVOR_READOUT_BOTTLENECK"

	StatusKappaOrientNotNativeFlavorTheorem             = "FAILED_ROUTE_KAPPA_ORIENT_NOT_NATIVE_FLAVOR_THEOREM"
	StatusNoNativeTheta13Theorem                        = "FAILED_ROUTE_NO_NATIVE_THETA13_THEOREM"
	StatusNoMapK7QuarterToPMNSReactor                   = "FAILED_ROUTE_NO_TYPED_MAP_FROM_K7_PLUS_RADIAL_EVENT_WEIGHT_TO_PMNS_REACTOR_LEAKAGE"
	StatusPMNSReactorTermRemainsFlavorSeal              = "FAILED_ROUTE_PMNS_REACTOR_TERM_REMAINS_FLAVOR_SEAL_INPUT"
	StatusNoNativeJCKMTheorem                           = "FAILED_ROUTE_NO_NATIVE_J_CKM_THEOREM"
	StatusNoNativeCKMOrientationSignTheorem             = "FAILED_ROUTE_NO_NATIVE_CKM_ORIENTATION_SIGN_THEOREM"
	StatusCKMTermRemainsFlavorSeal                      = "FAILED_ROUTE_CKM_TERM_REMAINS_FLAVOR_SEAL_INPUT"
	StatusBoundaryCorrectionDoesNotReplacePMNSCKM       = "FAILED_ROUTE_BOUNDARY_CORRECTION_DOES_NOT_REPLACE_PMNS_CKM_ORIENTATION_READOUT"
	StatusK7HodgePolarityDoesNotDerivePMNSCKM           = "FAILED_ROUTE_K7_HODGE_POLARITY_DOES_NOT_DERIVE_PMNS_CKM_ORIENTATION"
	StatusHiggsRadialEventDoesNotDeriveTheta13          = "FAILED_ROUTE_HIGGS_RADIAL_EVENT_DOES_NOT_DERIVE_THETA13"
	StatusBoundaryPairDoesNotDeriveFlavorMixing         = "FAILED_ROUTE_BOUNDARY_PAIR_DOES_NOT_DERIVE_FLAVOR_MIXING"
	StatusNEffDoesNotDeriveMixingAngles                 = "FAILED_ROUTE_N_EFF_DOES_NOT_DERIVE_MIXING_ANGLES"
	StatusNoNativeGenerationMixingOperatorFound         = "FAILED_ROUTE_NO_NATIVE_GENERATION_MIXING_OPERATOR_FOUND"
	StatusFlavorOrientationReadoutSealNotNative         = "FAILED_ROUTE_FLAVOR_ORIENTATION_READOUT_SEAL_NOT_NATIVE"
	StatusKappaOrientNotTheoremLevelIndependent         = "FAILED_ROUTE_KAPPA_ORIENT_NOT_THEOREM_LEVEL_INDEPENDENT"
	StatusCHistoryNotFullIndependentPredictionComponent = "FAILED_ROUTE_C_HISTORY_NOT_FULL_INDEPENDENT_PREDICTION_COMPONENT"
	StatusCHiggsNotLevelCPrediction                     = "FAILED_ROUTE_C_HIGGS_NOT_LEVEL_C_PREDICTION"
	StatusTreeProxyNotPoleMass                          = "FAILED_ROUTE_TREE_PROXY_NOT_POLE_MASS"
	StatusFirewallPreservedGate788                      = "FIREWALL_PRESERVED_GATE788_FLAVOR_ORIENTATION_READOUT_BOUNDARY"
)

const (
	pK7Snapshot           = 7.0 / 72.0
	sSplitSnapshot        = 0.0012924448188162962
	xiBoundarySnapshot    = 0.0503471644870914
	kappaERedSnapshot     = 0.005503554218475772
	kappaOrientSnapshot   = 0.00550633006471245
	kappaBoundarySnapshot = -2.775846236678231e-6
)

type Gate787Inheritance struct {
	Inherited          bool
	CompositeSeal      string
	KappaOrientIsFocus bool
	Verdict            string
}

type KappaOrientDecomposition struct {
	Recorded      bool
	Formula       string
	PMNSTerm      string
	QuarterFactor string
	CKMTerm       string
	NegativeSign  string
	ShapeTyped    bool
	Native        bool
	Verdict       string
}

type PMNSReactorLeakageAudit struct {
	Audited                 bool
	Term                    string
	Theta13SourceCandidates []string
	Theta13Native           bool
	QuarterResonance        bool
	QuarterResonanceSources []string
	TypedMapFromK7Quarter   bool
	RemainsFlavorSealInput  bool
	Verdict                 string
}

type CKMJarlskogOrientationAudit struct {
	Audited               bool
	Term                  string
	SourceType            string
	SourceCandidates      []string
	JCKMNative            bool
	NegativeSignCandidate bool
	NativeSignTheorem     bool
	RemainsFlavorSeal     bool
	Verdict               string
}

type BoundaryOnlyReplacementAudit struct {
	Audited                     bool
	KappaOrient                 float64
	KappaBoundary               float64
	AbsRatioBoundaryToOrient    float64
	BoundaryPartSmallCorrection bool
	BoundaryReplacesOrient      bool
	Verdict                     string
}

type ASHAGeometrySourceCandidateAudit struct {
	Audited                         bool
	Candidates                      map[string]string
	K7HodgeDerivesPMNSCKM           bool
	HiggsRadialDerivesTheta13       bool
	BoundaryPairDerivesFlavorMixing bool
	NEffDerivesMixingAngles         bool
	GenerationMixingOperatorFound   bool
	Verdict                         string
}

type FlavorBoundaryReadoutSealRefinement struct {
	Recorded                   bool
	OriginalSeal               string
	RefinedSeals               []string
	KappaOrient                float64
	KappaBoundary              float64
	BoundaryGaugeStronglyTyped bool
	OrientationTrueObstruction bool
	OrientationSealNative      bool
	Verdict                    string
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
	KappaOrient string
	KappaERed   string
	FWall3      string
	KappaLambda string
	CHistory    string
	CHiggs      string
	Verdict     string
}

type Firewalls struct {
	Enforced                        bool
	KappaOrientNativeFlavorTheorem  bool
	Theta13DerivedPMNSTheorem       bool
	JCKMDerivedCKMTheorem           bool
	QuarterResonanceProof           bool
	K7QuarterTheta13SourceTheorem   bool
	NEffMixingAngleTheorem          bool
	BoundaryPairFlavorMixingTheorem bool
	KappaBoundaryFullKappaETheorem  bool
	FWallNativeBoundaryTheorem      bool
	CHistoryFullPrediction          bool
	TreeProxyPoleMass               bool
	Verdict                         string
}

type Audit struct {
	Gate787        Gate787Inheritance
	Decomposition  KappaOrientDecomposition
	PMNS           PMNSReactorLeakageAudit
	CKM            CKMJarlskogOrientationAudit
	BoundaryOnly   BoundaryOnlyReplacementAudit
	Geometry       ASHAGeometrySourceCandidateAudit
	SealRefinement FlavorBoundaryReadoutSealRefinement
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
	kappaBoundary := (-5.0/3.0 + xiBoundarySnapshot*pK7Snapshot) * sSplitSnapshot * sSplitSnapshot
	kappaOrient := kappaERedSnapshot - kappaBoundary
	if !closeRel(kappaBoundary, kappaBoundarySnapshot, 2e-15) {
		return Audit{}, fmt.Errorf("kappa_boundary mismatch: got %.18g want %.18g", kappaBoundary, kappaBoundarySnapshot)
	}
	if !closeRel(kappaOrient, kappaOrientSnapshot, 2e-15) {
		return Audit{}, fmt.Errorf("kappa_orient mismatch: got %.18g want %.18g", kappaOrient, kappaOrientSnapshot)
	}
	if !closeRel(kappaOrient+kappaBoundary, kappaERedSnapshot, 1e-15) {
		return Audit{}, fmt.Errorf("kappa split does not reconstruct kappa_e_red")
	}

	forbidden := []string{"lambda_runtime", "lambda_runtime_eff", "m_H_tree", "m_H_pole", "C_Higgs", "G_F", "v"}
	formula := "sin^2(theta13)/4 - J_CKM"
	containsForbidden := containsAny(formula, forbidden)

	return Audit{
		Gate787: Gate787Inheritance{
			Inherited:          true,
			CompositeSeal:      "BoundaryExteriorResponsePackageSeal = (DegreeRuleSeal, FlavorBoundaryReadoutSeal, BoundaryStressPullOrientationSeal)",
			KappaOrientIsFocus: true,
			Verdict:            StatusGate787Inherited,
		},
		Decomposition: KappaOrientDecomposition{
			Recorded:      true,
			Formula:       formula,
			PMNSTerm:      "sin^2(theta13)/4 = PMNS reactor leakage candidate",
			QuarterFactor: "1/4 = projector/radial-event-style quarter normalization candidate; resonance only, no typed PMNS map",
			CKMTerm:       "J_CKM = CKM Jarlskog oriented flavor-area correction candidate",
			NegativeSign:  "minus sign = orientation-subtraction candidate between leptonic leakage and quark-sector area",
			ShapeTyped:    true,
			Native:        false,
			Verdict:       StatusKappaOrientFlavorOrientationReadoutShape,
		},
		PMNS: PMNSReactorLeakageAudit{
			Audited:                 true,
			Term:                    "sin^2(theta13)/4",
			Theta13SourceCandidates: []string{"generation carrier theorem", "Yukawa operator theorem", "PMNS mixing theorem", "Fock/projective selector theorem", "K7-to-generation orientation theorem", "flavor wall orientation theorem"},
			Theta13Native:           false,
			QuarterResonance:        true,
			QuarterResonanceSources: []string{"rho_plus=I_K7+/4", "Tr(rho_plus P_rad)=1/4", "dim_R K7+=4", "rank-one radial event weight"},
			TypedMapFromK7Quarter:   false,
			RemainsFlavorSealInput:  true,
			Verdict:                 StatusOneQuarterK7PlusEventWeightResonance,
		},
		CKM: CKMJarlskogOrientationAudit{
			Audited:               true,
			Term:                  "-J_CKM",
			SourceType:            "quark-sector CP/orientation area invariant; signed flavor mixing area candidate",
			SourceCandidates:      []string{"Yukawa operator theorem", "CKM mixing theorem", "generation orientation theorem", "quark-sector phase theorem", "triality/flavor carrier theorem", "boundary orientation theorem"},
			JCKMNative:            false,
			NegativeSignCandidate: true,
			NativeSignTheorem:     false,
			RemainsFlavorSeal:     true,
			Verdict:               StatusJCKMOrientedFlavorAreaSourceType,
		},
		BoundaryOnly: BoundaryOnlyReplacementAudit{
			Audited:                     true,
			KappaOrient:                 kappaOrient,
			KappaBoundary:               kappaBoundary,
			AbsRatioBoundaryToOrient:    math.Abs(kappaBoundary / kappaOrient),
			BoundaryPartSmallCorrection: true,
			BoundaryReplacesOrient:      false,
			Verdict:                     StatusKappaBoundarySmallCorrectionToFlavorOrient,
		},
		Geometry: ASHAGeometrySourceCandidateAudit{
			Audited: true,
			Candidates: map[string]string{
				"K7 Hodge polarity 4|3":          "supplies split-signature / Hodge polarity, not PMNS/CKM angles",
				"K7+ radial/Higgs event":         "supplies rank-one quarter weight, not theta13",
				"CP1 Higgs vacuum orbit":         "supplies Higgs socket orbit, not flavor mixing",
				"Boundary pair":                  "supplies wall coordinates, not flavor angles",
				"Fano/quaternionic structure":    "supplies twistor/socket geometry, not a generation mixing matrix",
				"spectral-action trace pair a,b": "supplies aggregate Yukawa participation N_eff, not mixing angles",
				"N_eff":                          "supplies aggregate trace spread, not PMNS/CKM orientation",
			},
			K7HodgeDerivesPMNSCKM:           false,
			HiggsRadialDerivesTheta13:       false,
			BoundaryPairDerivesFlavorMixing: false,
			NEffDerivesMixingAngles:         false,
			GenerationMixingOperatorFound:   false,
			Verdict:                         StatusNoNativeGenerationMixingOperatorFound,
		},
		SealRefinement: FlavorBoundaryReadoutSealRefinement{
			Recorded:                   true,
			OriginalSeal:               "FlavorBoundaryReadoutSeal",
			RefinedSeals:               []string{"FlavorOrientationReadoutSeal", "BoundaryGaugeCorrectionSeal"},
			KappaOrient:                kappaOrient,
			KappaBoundary:              kappaBoundary,
			BoundaryGaugeStronglyTyped: true,
			OrientationTrueObstruction: true,
			OrientationSealNative:      false,
			Verdict:                    StatusFlavorOrientationTrueNonNativeObstruction,
		},
		Runtime: RuntimeIndependenceAudit{
			Audited:                  true,
			ForbiddenDirectVariables: forbidden,
			ContainsForbidden:        containsForbidden,
			FormulaLevelIndependent:  !containsForbidden,
			TheoremLevelIndependent:  false,
			Verdict:                  StatusKappaOrientFormulaLevelRuntimeIndependent,
		},
		Propagation: StatusPropagation{
			Recorded:    true,
			KappaOrient: "formula-level runtime-independent FlavorOrientationReadoutSeal; not native",
			KappaERed:   "mixed flavor-boundary readout; boundary correction strongly typed, orientation part sealed",
			FWall3:      "Level B+ seal-factorized exterior response package; still not native",
			KappaLambda: "Level B formula-independent scalar complement; still not native",
			CHistory:    "Level B semi-independent correction; not full prediction",
			CHiggs:      "still not Level C",
			Verdict:     StatusKappaOrientCurrentFlavorReadoutBottleneck,
		},
		Firewalls: Firewalls{
			Enforced:                        true,
			KappaOrientNativeFlavorTheorem:  false,
			Theta13DerivedPMNSTheorem:       false,
			JCKMDerivedCKMTheorem:           false,
			QuarterResonanceProof:           false,
			K7QuarterTheta13SourceTheorem:   false,
			NEffMixingAngleTheorem:          false,
			BoundaryPairFlavorMixingTheorem: false,
			KappaBoundaryFullKappaETheorem:  false,
			FWallNativeBoundaryTheorem:      false,
			CHistoryFullPrediction:          false,
			TreeProxyPoleMass:               false,
			Verdict:                         StatusFirewallPreservedGate788,
		},
		Truth:          "Gate 788 audits the exposed flavor-orientation readout without promoting PMNS, CKM, Yukawa, or flavor-hierarchy data into native ASHA theorems.",
		FinalStatement: "Gate 788 does not source kappa_orient natively. It improves the ledger by separating the flavor-boundary readout into a strongly typed boundary-gauge correction and a true flavor-orientation seal: kappa_e_red = [sin^2(theta13)/4 - J_CKM] + [-5/3 + xi_boundary p]s^2. The next bottleneck is not the boundary correction. It is the missing native generation/mixing operator that could source theta13 and J_CKM, or else certify FlavorOrientationReadoutSeal as an external flavor input.",
	}, nil
}

func FormatKappaSplit(a Audit) string {
	return fmt.Sprintf("kappa_e_red = kappa_orient + kappa_boundary = %.18g + %.18g = %.18g", a.SealRefinement.KappaOrient, a.SealRefinement.KappaBoundary, a.SealRefinement.KappaOrient+a.SealRefinement.KappaBoundary)
}

func FormatBoundaryOnly(a BoundaryOnlyReplacementAudit) string {
	return fmt.Sprintf("|kappa_boundary/kappa_orient| = %.18g; boundary correction is small and cannot replace PMNS/CKM orientation readout", a.AbsRatioBoundaryToOrient)
}

func FormatSealRefinement(s FlavorBoundaryReadoutSealRefinement) string {
	return fmt.Sprintf("%s -> %s; kappa_orient=%.18g, kappa_boundary=%.18g", s.OriginalSeal, strings.Join(s.RefinedSeals, " + "), s.KappaOrient, s.KappaBoundary)
}

func Statuses() []string {
	return []string{
		StatusGate787Inherited,
		StatusKappaOrientMainNonNativeComponent,
		StatusKappaOrientDecomposed,
		StatusPMNSReactorLeakageAudited,
		StatusCKMJarlskogOrientationAudited,
		StatusBoundaryOnlyReplacementAudited,
		StatusExistingASHAGeometryCandidatesAudited,
		StatusFlavorBoundaryReadoutSealRefined,
		StatusKappaOrientRuntimeTargetAbsenceAudited,
		StatusStatusPropagationRecorded,
		StatusPhysicalFirewallsEnforced,
		StatusKappaOrientFlavorOrientationReadoutShape,
		StatusOneQuarterK7PlusEventWeightResonance,
		StatusJCKMOrientedFlavorAreaSourceType,
		StatusNegativeCKMSignOrientationSubtractionCandidate,
		StatusKappaBoundarySmallCorrectionToFlavorOrient,
		StatusBoundaryGaugeCorrectionStronglyTyped,
		StatusFlavorOrientationTrueNonNativeObstruction,
		StatusKappaOrientFormulaLevelRuntimeIndependent,
		StatusKappaOrientCurrentFlavorReadoutBottleneck,
		StatusKappaOrientNotNativeFlavorTheorem,
		StatusNoNativeTheta13Theorem,
		StatusNoMapK7QuarterToPMNSReactor,
		StatusPMNSReactorTermRemainsFlavorSeal,
		StatusNoNativeJCKMTheorem,
		StatusNoNativeCKMOrientationSignTheorem,
		StatusCKMTermRemainsFlavorSeal,
		StatusBoundaryCorrectionDoesNotReplacePMNSCKM,
		StatusK7HodgePolarityDoesNotDerivePMNSCKM,
		StatusHiggsRadialEventDoesNotDeriveTheta13,
		StatusBoundaryPairDoesNotDeriveFlavorMixing,
		StatusNEffDoesNotDeriveMixingAngles,
		StatusNoNativeGenerationMixingOperatorFound,
		StatusFlavorOrientationReadoutSealNotNative,
		StatusKappaOrientNotTheoremLevelIndependent,
		StatusCHistoryNotFullIndependentPredictionComponent,
		StatusCHiggsNotLevelCPrediction,
		StatusTreeProxyNotPoleMass,
		StatusFirewallPreservedGate788,
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
