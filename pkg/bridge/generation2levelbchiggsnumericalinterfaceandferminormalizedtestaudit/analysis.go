// Package generation2levelbchiggsnumericalinterfaceandferminormalizedtestaudit implements
// Gate 791: Level-B C_Higgs Numerical Interface and Fermi-Normalized Test Audit.
//
// Gate 790 froze C_Higgs as a clean Level-B dimensionless scalar-Higgs bridge
// interface. Gate 791 turns that freeze into the explicit numerical and
// Fermi-normalized test interface while preserving the pole-mass, VEV/G_F,
// Yukawa, PMNS/CKM, boundary-response, and HistoryLoop firewalls.
package generation2levelbchiggsnumericalinterfaceandferminormalizedtestaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE791-LEVEL-B-C-HIGGS-NUMERICAL-INTERFACE-FERMI-NORMALIZED-TEST-AUDIT"

	StatusGate790Inherited                 = "PASS_GATE790_C_HIGGS_LEVEL_B_INTERFACE_INHERITED"
	StatusInternalCHiggsObjectDefined      = "PASS_INTERNAL_C_HIGGS_OBJECT_DEFINED"
	StatusBridgeQuarticRecorded            = "PASS_BRIDGE_QUARTIC_FROM_C_HIGGS_RECORDED"
	StatusFermiTreeInterfaceDefined        = "PASS_FERMI_NORMALIZED_TREE_INTERFACE_DEFINED"
	StatusTreeProxyLedgerRecomputed        = "PASS_TREE_PROXY_NUMERICAL_LEDGER_RECOMPUTED"
	StatusObservableDiagnosticDefined      = "PASS_OBSERVABLE_SIDE_DIAGNOSTIC_RATIO_DEFINED"
	StatusNonCircularProtocolDefined       = "PASS_NONCIRCULAR_LEVEL_B_PROTOCOL_DEFINED"
	StatusCorrectionDecompositionRecorded  = "PASS_CORRECTION_FACTOR_DECOMPOSITION_RECORDED"
	StatusSensitivityFormulasRecorded      = "PASS_LEVEL_B_SENSITIVITY_FORMULAS_RECORDED"
	StatusTestStatusClassificationRecorded = "PASS_TEST_STATUS_CLASSIFICATION_RECORDED"
	StatusSourcePressureMapRecorded        = "PASS_SOURCE_PRESSURE_MAP_RECORDED"
	StatusNextBranchRecommendationRecorded = "PASS_NEXT_BRANCH_RECOMMENDATION_RECORDED"
	StatusPhysicalFirewallsEnforced        = "PASS_PHYSICAL_FIREWALLS_ENFORCED"

	StatusCHiggsDimensionlessLevelBOutput     = "CONDITIONAL_SUPPORT_C_HIGGS_IS_DIMENSIONLESS_LEVEL_B_ASHA_OUTPUT"
	StatusFermiTreeIdentityAtProxyLevel       = "CONDITIONAL_SUPPORT_4_SQRT2_GF_M_TREE_SQUARED_EQUALS_C_HIGGS_AT_TREE_PROXY_LEVEL"
	StatusDeltaRPoleDiagnosticGap             = "CONDITIONAL_SUPPORT_DELTA_R_POLE_IS_VALID_ONLY_AS_LEVEL_1C_DIAGNOSTIC_GAP"
	StatusHistoryUpliftDominates              = "CONDITIONAL_SUPPORT_HISTORY_UPLIFT_DOMINATES_LEVEL_B_CORRECTION"
	StatusYukawaDilutesHistory                = "CONDITIONAL_SUPPORT_YUKAWA_PARTICIPATION_DILUTES_HISTORY_UPLIFT"
	StatusKeySensitivityChannels              = "CONDITIONAL_SUPPORT_KAPPA_ORIENT_AND_N_EFF_ARE_KEY_NON_NATIVE_SENSITIVITY_CHANNELS"
	StatusCleanNonCircularDimensionlessObject = "CONDITIONAL_SUPPORT_C_HIGGS_IS_CURRENT_CLEANEST_NONCIRCULAR_DIMENSIONLESS_TEST_OBJECT"

	StatusCHiggsNotNativeHiggsTheorem      = "FAILED_ROUTE_C_HIGGS_NOT_NATIVE_HIGGS_THEOREM"
	StatusFermiTreeIdentityNotPoleTheorem  = "FAILED_ROUTE_FERMI_NORMALIZED_TREE_IDENTITY_NOT_POLE_MASS_THEOREM"
	StatusDeltaRPoleNotCorrectionTheorem   = "FAILED_ROUTE_DELTA_R_POLE_NOT_NATIVE_TREE_TO_POLE_CORRECTION_THEOREM"
	StatusExternalPoleNotASHADerivation    = "FAILED_ROUTE_EXTERNAL_POLE_OBSERVABLE_NOT_ASHA_DERIVATION"
	StatusObservedHiggsMassForbidden       = "FAILED_ROUTE_OBSERVED_HIGGS_MASS_MUST_NOT_SOURCE_C_HIGGS_COMPONENTS"
	StatusLevelBNotLevelCNativePrediction  = "FAILED_ROUTE_LEVEL_B_TEST_INTERFACE_NOT_LEVEL_C_NATIVE_PREDICTION"
	StatusTreeProxyNotPoleMass             = "FAILED_ROUTE_TREE_PROXY_NOT_POLE_MASS"
	StatusPoleComparisonRequiresCorrection = "FAILED_ROUTE_POLE_COMPARISON_REQUIRES_CORRECTION_PACKAGE"
	StatusNEffNotNativeYukawa              = "FAILED_ROUTE_N_EFF_NOT_NATIVE_YUKAWA_THEOREM"
	StatusKappaOrientNotNativePMNSCKM      = "FAILED_ROUTE_KAPPA_ORIENT_NOT_NATIVE_PMNS_CKM_THEOREM"
	StatusLHopfNotNativeHistoryLoop        = "FAILED_ROUTE_L_HOPF_NOT_NATIVE_HISTORYLOOP_THEOREM"
	StatusFWallNotNativeBoundary           = "FAILED_ROUTE_F_WALL_3_RED_NOT_NATIVE_BOUNDARY_RESPONSE_THEOREM"
	StatusVOrGFNotNativeElectroweakScale   = "FAILED_ROUTE_V_OR_GF_NOT_NATIVE_ELECTROWEAK_SCALE_THEOREM"
	StatusFirewallPreservedGate791         = "FIREWALL_PRESERVED_GATE791_LEVEL_B_C_HIGGS_TEST_INTERFACE_BOUNDARY"
)

const (
	pSnapshot              = 7.0 / 72.0
	sSnapshot              = 0.0012924448188162962
	xiBoundarySnapshot     = 0.0503471644870914
	kappaOrientSnapshot    = 0.00550633006471245
	kappaBoundarySnapshot  = -2.775846236678231e-6
	kappaERedSnapshot      = 0.005503554218475772
	fWall3Snapshot         = 0.00012565521035653708
	absLambdaSnapshot      = 0.049700942077680596
	kappaLambdaRedSnapshot = 0.04432304306956136
	lHopfSnapshot          = 0.039788735772973836
	nEffSnapshot           = 3.0023273474722147
	cYukawaSnapshot        = 0.9992248188812008
	cHistorySnapshot       = 1.038025177923625
	cHiggsSnapshot         = 1.0372205204048603
	lambdaHBridgeSnapshot  = 0.12965256505060754
	vSealSnapshot          = 246.2196508
	gFEquivalentSnapshot   = 1.1663786999444556e-05
)

type Gate790Inheritance struct {
	Inherited         bool
	FrozenLevelB      bool
	CleanInterface    bool
	DirectRuntimeFree bool
	Verdict           string
}

type InternalObject struct {
	Defined              bool
	CHiggsFormula        string
	CHiggsValue          float64
	LambdaHBridgeFormula string
	LambdaHBridgeValue   float64
	Dimensionless        bool
	NativeHiggsTheorem   bool
	Verdict              string
}

type FermiTreeInterface struct {
	Defined          bool
	VEVRelation      string
	TreeProxyFormula string
	DimensionlessID  string
	AtTreeProxyLevel bool
	PoleMassTheorem  bool
	Verdict          string
}

type TreeProxyLedger struct {
	Recomputed     bool
	CHiggs         float64
	SqrtCHiggs     float64
	V              float64
	VHalf          float64
	GF             float64
	MTree          float64
	MTreeSquared   float64
	FermiRatio     float64
	ReducedQuarter float64
	LambdaHBridge  float64
	Verdict        string
}

type ObservableDiagnostic struct {
	Defined              bool
	ExternalRatio        string
	DiagnosticGap        string
	GapInterpretation    []string
	Level1COnly          bool
	NativePoleCorrection bool
	ExternalPoleDerived  bool
	Verdict              string
}

type NonCircularProtocol struct {
	Defined                 bool
	AllowedInputs           []string
	Compute                 []string
	ForbiddenCircularInputs []string
	ObservedHiggsForbidden  bool
	Verdict                 string
}

type CorrectionDecomposition struct {
	Recorded        bool
	CYukawa         float64
	CHistory        float64
	CHiggs          float64
	EpsilonYukawa   float64
	DeltaHistory    float64
	DeltaHiggs      float64
	Multiplicative  string
	HistoryDominant bool
	YukawaDilution  bool
	Verdict         string
}

type SensitivityFormulas struct {
	Recorded                   bool
	CHiggsRelative             string
	CYukawaRelative            string
	CHistoryDifferential       string
	KappaLambdaRedDifferential string
	KappaERedDifferential      string
	KeyChannels                []string
	Verdict                    string
}

type TestStatusClassification struct {
	Recorded       bool
	CHiggs         string
	LambdaHBridge  string
	TreeProxy      string
	PoleDiagnostic string
	PoleMass       string
	CleanObject    bool
	LevelCNative   bool
	Verdict        string
}

type SourcePressureMap struct {
	Recorded  bool
	Pressures []string
	Verdict   string
}

type NextBranchRecommendation struct {
	Recorded    bool
	Recommended string
	Alternative string
	Reason      string
	Verdict     string
}

type Firewalls struct {
	Enforced                    bool
	CHiggsNativeHiggsTheorem    bool
	CHiggsPoleMassPrediction    bool
	FermiTreeIdentityPoleMass   bool
	LambdaHIndependentRuntime   bool
	TreeProxyPoleMass           bool
	ExternalPoleASHADerived     bool
	DeltaRPoleCorrectionTheorem bool
	NEffNativeYukawa            bool
	KappaOrientNativePMNSCKM    bool
	LHopfNativeHistoryLoop      bool
	FWallNativeBoundary         bool
	VOrGFNativeScale            bool
	Verdict                     string
}

type Analysis struct {
	Gate790        Gate790Inheritance
	Internal       InternalObject
	FermiTree      FermiTreeInterface
	TreeLedger     TreeProxyLedger
	Diagnostic     ObservableDiagnostic
	Protocol       NonCircularProtocol
	Decomposition  CorrectionDecomposition
	Sensitivity    SensitivityFormulas
	Classification TestStatusClassification
	Pressure       SourcePressureMap
	Next           NextBranchRecommendation
	Firewalls      Firewalls
	Truth          string
	FinalStatement string
}

func BuildDefault() (Analysis, error) {
	lHopf := 1.0 / (8.0 * math.Pi)
	cYukawa := 3.0 / nEffSnapshot
	cHistory := 1.0 + lHopf*(1.0-kappaLambdaRedSnapshot)
	cHiggs := cYukawa * cHistory
	lambdaHBridge := cHiggs / 8.0
	gF := 1.0 / (math.Sqrt2 * vSealSnapshot * vSealSnapshot)
	sqrtC := math.Sqrt(cHiggs)
	mTree := (vSealSnapshot / 2.0) * sqrtC
	mTreeSquared := mTree * mTree
	fermiRatio := 4.0 * math.Sqrt2 * gF * mTreeSquared
	reducedQuarter := math.Sqrt2 * gF * mTreeSquared
	epsilonYukawa := 1.0 - cYukawa
	deltaHistory := cHistory - 1.0
	deltaHiggs := cHiggs - 1.0
	checks := []struct {
		name           string
		got, want, tol float64
	}{
		{"L_Hopf", lHopf, lHopfSnapshot, 1e-16},
		{"C_Yukawa", cYukawa, cYukawaSnapshot, 1e-15},
		{"C_History", cHistory, cHistorySnapshot, 1e-15},
		{"C_Higgs", cHiggs, cHiggsSnapshot, 1e-15},
		{"lambda_H_bridge", lambdaHBridge, lambdaHBridgeSnapshot, 1e-15},
		{"G_F equivalent", gF, gFEquivalentSnapshot, 1e-18},
		{"Fermi ratio", fermiRatio, cHiggsSnapshot, 1e-15},
		{"reduced quarter", reducedQuarter, cHiggsSnapshot / 4.0, 1e-15},
	}
	for _, c := range checks {
		if !closeAbs(c.got, c.want, c.tol) {
			return Analysis{}, fmt.Errorf("%s mismatch: got %.17g want %.17g", c.name, c.got, c.want)
		}
	}
	if !closeAbs(deltaHiggs, deltaHistory-epsilonYukawa*(1.0+deltaHistory), 1e-15) {
		return Analysis{}, fmt.Errorf("multiplicative Delta_Higgs identity failed")
	}

	a := Analysis{
		Gate790: Gate790Inheritance{
			Inherited:         true,
			FrozenLevelB:      true,
			CleanInterface:    true,
			DirectRuntimeFree: true,
			Verdict:           StatusGate790Inherited,
		},
		Internal: InternalObject{
			Defined:              true,
			CHiggsFormula:        "C_Higgs_ASHA_LevelB = (3/N_eff)[1 + L_Hopf(1-kappa_lambda_red)]",
			CHiggsValue:          cHiggs,
			LambdaHBridgeFormula: "lambda_H_bridge = C_Higgs/8 after Gate770 quartic airlock",
			LambdaHBridgeValue:   lambdaHBridge,
			Dimensionless:        true,
			NativeHiggsTheorem:   false,
			Verdict:              StatusInternalCHiggsObjectDefined,
		},
		FermiTree: FermiTreeInterface{
			Defined:          true,
			VEVRelation:      "v = (sqrt(2)G_F)^(-1/2)",
			TreeProxyFormula: "m_H_tree_proxy = (v/2)sqrt(C_Higgs)",
			DimensionlessID:  "4 sqrt(2) G_F m_H_tree_proxy^2 = C_Higgs",
			AtTreeProxyLevel: true,
			PoleMassTheorem:  false,
			Verdict:          StatusFermiTreeInterfaceDefined,
		},
		TreeLedger: TreeProxyLedger{
			Recomputed:     true,
			CHiggs:         cHiggs,
			SqrtCHiggs:     sqrtC,
			V:              vSealSnapshot,
			VHalf:          vSealSnapshot / 2.0,
			GF:             gF,
			MTree:          mTree,
			MTreeSquared:   mTreeSquared,
			FermiRatio:     fermiRatio,
			ReducedQuarter: reducedQuarter,
			LambdaHBridge:  lambdaHBridge,
			Verdict:        StatusTreeProxyLedgerRecomputed,
		},
		Diagnostic: ObservableDiagnostic{
			Defined:       true,
			ExternalRatio: "R_pole_external = 4 sqrt(2) G_F_external m_H_pole_external^2",
			DiagnosticGap: "Delta_R_pole = R_pole_external - C_Higgs",
			GapInterpretation: []string{
				"tree-to-pole corrections",
				"RG/threshold corrections",
				"scheme dependence",
				"top/gauge loop effects",
				"measurement uncertainty",
				"external pole convention",
			},
			Level1COnly:          true,
			NativePoleCorrection: false,
			ExternalPoleDerived:  false,
			Verdict:              StatusObservableDiagnosticDefined,
		},
		Protocol: NonCircularProtocol{
			Defined: true,
			AllowedInputs: []string{
				"N_eff from non-Higgs Yukawa trace ledger",
				"kappa_orient from flavor mixing data or future GenerationMixingOperatorSeal",
				"boundary coordinates |lambda|, s, xi_boundary",
				"p from K7 event weight",
				"L_Hopf from RadialHessianHopfTransportSeal",
			},
			Compute: []string{
				"C_Higgs",
				"lambda_H_bridge = C_Higgs/8",
				"m_H_tree_proxy = (v/2)sqrt(C_Higgs) only after VEV/Fermi seal",
				"4 sqrt(2)G_F m_H_tree_proxy^2 = C_Higgs",
			},
			ForbiddenCircularInputs: []string{
				"observed Higgs mass must not choose N_eff",
				"observed Higgs mass must not choose kappa_orient",
				"observed Higgs mass must not choose F_wall_3_red",
				"observed Higgs mass must not choose L_Hopf",
			},
			ObservedHiggsForbidden: true,
			Verdict:                StatusNonCircularProtocolDefined,
		},
		Decomposition: CorrectionDecomposition{
			Recorded:        true,
			CYukawa:         cYukawa,
			CHistory:        cHistory,
			CHiggs:          cHiggs,
			EpsilonYukawa:   epsilonYukawa,
			DeltaHistory:    deltaHistory,
			DeltaHiggs:      deltaHiggs,
			Multiplicative:  "Delta_Higgs = delta_History - epsilon_Yukawa(1+delta_History)",
			HistoryDominant: deltaHistory > 40.0*epsilonYukawa,
			YukawaDilution:  true,
			Verdict:         StatusCorrectionDecompositionRecorded,
		},
		Sensitivity: SensitivityFormulas{
			Recorded:                   true,
			CHiggsRelative:             "delta C_Higgs / C_Higgs = delta C_Yukawa / C_Yukawa + delta C_History / C_History",
			CYukawaRelative:            "delta C_Yukawa / C_Yukawa = - delta N_eff / N_eff",
			CHistoryDifferential:       "delta C_History = (1-kappa_lambda_red) delta L_Hopf - L_Hopf delta kappa_lambda_red",
			KappaLambdaRedDifferential: "delta kappa_lambda_red = delta |lambda| + delta F_wall_3_red - delta kappa_e_red",
			KappaERedDifferential:      "delta kappa_e_red = delta kappa_orient + delta kappa_boundary",
			KeyChannels: []string{
				"kappa_orient",
				"N_eff",
			},
			Verdict: StatusSensitivityFormulasRecorded,
		},
		Classification: TestStatusClassification{
			Recorded:       true,
			CHiggs:         "Level-B dimensionless test target",
			LambdaHBridge:  "bridge quartic after airlock: lambda_H_bridge=C_Higgs/8",
			TreeProxy:      "Level-1B tree Hessian proxy after VEV/Fermi seal",
			PoleDiagnostic: "Level-1C diagnostic only if external pole observable and correction package are supplied",
			PoleMass:       "m_H_pole not predicted",
			CleanObject:    true,
			LevelCNative:   false,
			Verdict:        StatusTestStatusClassificationRecorded,
		},
		Pressure: SourcePressureMap{
			Recorded: true,
			Pressures: []string{
				"1. GenerationMixingOperatorSeal: needed for kappa_orient",
				"2. Yukawa operator/eigenvector theorem: needed for N_eff and eventually mixing",
				"3. RadialHessianHopfTransportSeal: needed for native L_Hopf transport law",
				"4. BoundaryExteriorResponsePackageSeal: needed for native F_wall_3_red",
				"5. Electroweak scale theorem: needed for native v or G_F",
				"6. Tree-to-pole correction package: needed for physical pole comparison",
			},
			Verdict: StatusSourcePressureMapRecorded,
		},
		Next: NextBranchRecommendation{
			Recorded:    true,
			Recommended: "Gate 792 — Level-B Error Budget and Independent-Input Sensitivity Audit",
			Alternative: "Gate 792 — N_eff and FlavorOrientation Joint Yukawa-Mixing Source Audit",
			Reason:      "before chasing another deep theorem branch, quantify which sealed inputs dominate the Level-B C_Higgs uncertainty and scientific testability",
			Verdict:     StatusNextBranchRecommendationRecorded,
		},
		Firewalls: Firewalls{
			Enforced:                    true,
			CHiggsNativeHiggsTheorem:    false,
			CHiggsPoleMassPrediction:    false,
			FermiTreeIdentityPoleMass:   false,
			LambdaHIndependentRuntime:   false,
			TreeProxyPoleMass:           false,
			ExternalPoleASHADerived:     false,
			DeltaRPoleCorrectionTheorem: false,
			NEffNativeYukawa:            false,
			KappaOrientNativePMNSCKM:    false,
			LHopfNativeHistoryLoop:      false,
			FWallNativeBoundary:         false,
			VOrGFNativeScale:            false,
			Verdict:                     StatusFirewallPreservedGate791,
		},
		Truth:          "C_Higgs is the current clean Level-B dimensionless ASHA output: noncircular at formula level, explicit about its seals, and not promoted to pole mass or native Higgs theorem.",
		FinalStatement: "Gate 791 does not turn C_Higgs into a native prediction. It turns the frozen scalar-Higgs bridge into a clean Level-B dimensionless test interface: C_Higgs is computed internally from declared bridge/sealed inputs, and the lawful comparison channel is the Fermi-normalized tree identity 4 sqrt(2) G_F m_H_tree_proxy^2 = C_Higgs. Gate 792 should audit the Level-B error budget and independent-input sensitivities, so ASHA knows which remaining seal most limits scientific testability.",
	}
	return a, nil
}

func Statuses() []string {
	return []string{
		StatusGate790Inherited,
		StatusInternalCHiggsObjectDefined,
		StatusBridgeQuarticRecorded,
		StatusFermiTreeInterfaceDefined,
		StatusTreeProxyLedgerRecomputed,
		StatusObservableDiagnosticDefined,
		StatusNonCircularProtocolDefined,
		StatusCorrectionDecompositionRecorded,
		StatusSensitivityFormulasRecorded,
		StatusTestStatusClassificationRecorded,
		StatusSourcePressureMapRecorded,
		StatusNextBranchRecommendationRecorded,
		StatusPhysicalFirewallsEnforced,
		StatusCHiggsDimensionlessLevelBOutput,
		StatusFermiTreeIdentityAtProxyLevel,
		StatusDeltaRPoleDiagnosticGap,
		StatusHistoryUpliftDominates,
		StatusYukawaDilutesHistory,
		StatusKeySensitivityChannels,
		StatusCleanNonCircularDimensionlessObject,
		StatusCHiggsNotNativeHiggsTheorem,
		StatusFermiTreeIdentityNotPoleTheorem,
		StatusDeltaRPoleNotCorrectionTheorem,
		StatusExternalPoleNotASHADerivation,
		StatusObservedHiggsMassForbidden,
		StatusLevelBNotLevelCNativePrediction,
		StatusTreeProxyNotPoleMass,
		StatusPoleComparisonRequiresCorrection,
		StatusNEffNotNativeYukawa,
		StatusKappaOrientNotNativePMNSCKM,
		StatusLHopfNotNativeHistoryLoop,
		StatusFWallNotNativeBoundary,
		StatusVOrGFNotNativeElectroweakScale,
		StatusFirewallPreservedGate791,
	}
}

func FormatTreeLedger(l TreeProxyLedger) string {
	return fmt.Sprintf("C_Higgs=%.17g; sqrt(C_Higgs)=%.17g; v=%.17g GeV; G_F=%.17g GeV^-2; v/2=%.17g GeV; m_H_tree_proxy=%.17g GeV; m_H_tree_proxy^2=%.17g GeV^2; 4sqrt(2)G_Fm_tree^2=%.17g; lambda_H_bridge=%.17g", l.CHiggs, l.SqrtCHiggs, l.V, l.GF, l.VHalf, l.MTree, l.MTreeSquared, l.FermiRatio, l.LambdaHBridge)
}

func FormatDecomposition(d CorrectionDecomposition) string {
	return fmt.Sprintf("C_Yukawa=%.17g; C_History=%.17g; C_Higgs=%.17g; epsilon_Yukawa=%.17g; delta_History=%.17g; Delta_Higgs=%.17g; %s", d.CYukawa, d.CHistory, d.CHiggs, d.EpsilonYukawa, d.DeltaHistory, d.DeltaHiggs, d.Multiplicative)
}

func FormatPressure(p SourcePressureMap) string {
	return strings.Join(p.Pressures, "; ")
}

func containsAll(haystack []string, needles []string) bool {
	for _, n := range needles {
		found := false
		for _, h := range haystack {
			if strings.Contains(h, n) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func closeAbs(a, b, tol float64) bool { return math.Abs(a-b) <= tol }

var buildOnce struct {
	sync.Once
	a   Analysis
	err error
}

func Cached() (Analysis, error) {
	buildOnce.Do(func() { buildOnce.a, buildOnce.err = BuildDefault() })
	return buildOnce.a, buildOnce.err
}
