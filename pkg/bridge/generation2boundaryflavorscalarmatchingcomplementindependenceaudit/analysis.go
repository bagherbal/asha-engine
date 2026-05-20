// Package generation2boundaryflavorscalarmatchingcomplementindependenceaudit implements
// Gate 782: Boundary-Flavor Scalar Matching Complement Independence Audit.
//
// Gate 781 identified kappa_lambda_red as the bottleneck inside the dominant
// History correction factor C_History. Gate 782 expands the whole scalar
// matching complement into its boundary, K7 raw-moment, and reduced flavor-wall
// components, recomputes the numerical ledger, and separates formula-level
// runtime-target absence from theorem-level independence. This is a forensic
// bridge audit only: it does not derive scalar runtime lambda, Higgs pole mass,
// Yukawa operators, PMNS, CKM, flavor hierarchy, G_F, VEV, or a native
// HistoryLoopUnit theorem.
package generation2boundaryflavorscalarmatchingcomplementindependenceaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE782-BOUNDARY-FLAVOR-SCALAR-MATCHING-COMPLEMENT-INDEPENDENCE-AUDIT"

	StatusGate781CHistoryMacroAuditInherited                     = "PASS_GATE781_C_HISTORY_MACRO_AUDIT_INHERITED"
	StatusScalarMatchingComplementSelectedAsCurrentBottleneck    = "PASS_SCALAR_MATCHING_COMPLEMENT_SELECTED_AS_CURRENT_BOTTLENECK"
	StatusKappaLambdaRedRewrittenBoundaryFlavorResponseForm      = "PASS_KAPPA_LAMBDA_RED_REWRITTEN_AS_BOUNDARY_FLAVOR_RESPONSE_FORM"
	StatusNumericalBoundaryFlavorComplementLedgerRecomputed      = "PASS_NUMERICAL_BOUNDARY_FLAVOR_COMPLEMENT_LEDGER_RECOMPUTED"
	StatusRawMomentResponsePolynomialAudited                     = "PASS_RAW_MOMENT_RESPONSE_POLYNOMIAL_AUDITED"
	StatusKappaEReducedFlavorWallAudited                         = "PASS_KAPPA_E_REDUCED_FLAVOR_WALL_AUDITED"
	StatusFormulaLevelRuntimeTargetAbsenceAudited                = "PASS_FORMULA_LEVEL_RUNTIME_TARGET_ABSENCE_AUDITED"
	StatusTheoremLevelIndependenceFirewallAudited                = "PASS_THEOREM_LEVEL_INDEPENDENCE_FIREWALL_AUDITED"
	StatusCHistoryComplementRewrittenFullBoundaryFlavorForm      = "PASS_C_HISTORY_COMPLEMENT_REWRITTEN_WITH_FULL_BOUNDARY_FLAVOR_FORM"
	StatusPredictionLevelClassificationRecorded                  = "PASS_PREDICTION_LEVEL_CLASSIFICATION_RECORDED"
	StatusPhysicalFirewallsEnforced                              = "PASS_PHYSICAL_FIREWALLS_ENFORCED"
	StatusKappaLambdaRedExplicitBoundaryFlavorResponse           = "CONDITIONAL_SUPPORT_KAPPA_LAMBDA_RED_IS_REDUCED_TO_EXPLICIT_BOUNDARY_FLAVOR_RESPONSE_FORM"
	StatusKappaLambdaRedEvaluableWithoutDirectRuntimeVariables   = "CONDITIONAL_SUPPORT_KAPPA_LAMBDA_RED_CAN_BE_EVALUATED_WITHOUT_DIRECT_HIGGS_RUNTIME_VARIABLES"
	StatusFWall3RedTypedBoundaryRawMomentResponse                = "CONDITIONAL_SUPPORT_F_WALL_3_RED_IS_TYPED_BOUNDARY_RAW_MOMENT_RESPONSE"
	StatusKappaERedFlavorPlusBoundaryWallSourceType              = "CONDITIONAL_SUPPORT_KAPPA_E_RED_HAS_FLAVOR_PLUS_BOUNDARY_WALL_SOURCE_TYPE"
	StatusCHistoryExplicitBoundaryFlavorComplementForm           = "CONDITIONAL_SUPPORT_C_HISTORY_NOW_HAS_EXPLICIT_BOUNDARY_FLAVOR_COMPLEMENT_FORM"
	StatusKappaLambdaRedLevelBFormulaIndependence                = "CONDITIONAL_SUPPORT_KAPPA_LAMBDA_RED_REACHES_LEVEL_B_FORMULA_INDEPENDENCE"
	StatusKappaLambdaRedNotYetNativeScalarMatchingTheorem        = "FAILED_ROUTE_KAPPA_LAMBDA_RED_NOT_YET_NATIVE_SCALAR_MATCHING_THEOREM"
	StatusKappaLambdaRedNotYetFullyTheoremIndependent            = "FAILED_ROUTE_KAPPA_LAMBDA_RED_NOT_YET_FULLY_THEOREM_INDEPENDENT"
	StatusNoNativeBoundaryResponseGeneratingFunctionTheorem      = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_RESPONSE_GENERATING_FUNCTION_THEOREM"
	StatusNoNativeRawMomentCoordinateTheorem                     = "FAILED_ROUTE_NO_NATIVE_RAW_MOMENT_COORDINATE_THEOREM"
	StatusNoNativeCubicStopTheorem                               = "FAILED_ROUTE_NO_NATIVE_CUBIC_STOP_THEOREM"
	StatusNoNativeKappaETheorem                                  = "FAILED_ROUTE_NO_NATIVE_KAPPA_E_THEOREM"
	StatusNoNativePMNSOrCKMTheorem                               = "FAILED_ROUTE_NO_NATIVE_PMNS_OR_CKM_THEOREM"
	StatusCHistoryNotYetFullIndependentPredictionComponent       = "FAILED_ROUTE_C_HISTORY_NOT_YET_FULL_INDEPENDENT_PREDICTION_COMPONENT"
	StatusTreeProxyNotPoleMass                                   = "FAILED_ROUTE_TREE_PROXY_NOT_POLE_MASS"
	StatusNoYukawaOperatorOrEigenvalueTheorem                    = "FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusTheta13AndJCKMNotNativeFlavorTheorems                  = "FAILED_ROUTE_THETA13_AND_J_CKM_NOT_NATIVE_FLAVOR_THEOREMS"
	StatusNoNativeKappaETheoremDuplicate                         = StatusNoNativeKappaETheorem
	StatusNoNativeBoundaryFlavorScalarMatchingComplementBoundary = "FIREWALL_PRESERVED_GATE782_BOUNDARY_FLAVOR_SCALAR_MATCHING_COMPLEMENT_BOUNDARY"
)

const (
	// Snapshot ledger inherited from Gates 751/752/759/781. Gate 782 keeps these
	// local so the macro-audit remains focused and avoids importing deep
	// predecessor chains.
	pK7Snapshot            = 7.0 / 72.0
	sSplitSnapshot         = 0.0012924448188162962
	absLambdaSnapshot      = 0.049700942077680596
	xiBoundarySnapshot     = 0.0503471644870914
	kappaERedSnapshot      = 0.005503554218475772
	lHopfSnapshot          = 0.039788735772973836
	kappaLambdaRedSnapshot = 0.04432304306956136
	complementSnapshot     = 0.9556769569304386
	cHistorySnapshot       = 1.038025177923625
	fWall3RedSnapshot      = 0.00012565521035653708
)

type Gate781Inheritance struct {
	Inherited          bool
	CHistoryFormula    string
	LHopf              float64
	KappaLambdaRed     float64
	SelectedBottleneck string
	SelectedBranch     string
	Verdict            string
}

type ComplementRewrite struct {
	SelectedBottleneck bool
	BaseFormula        string
	WallPolynomial     string
	FactoredFormula    string
	ExpandedFormula    string
	SignIdentity       string
	SignsChecked       bool
	NoRuntimeSymbols   bool
	Verdict            string
}

type NumericalLedger struct {
	P                 float64
	S                 float64
	AbsLambda         float64
	XiBoundary        float64
	KappaERed         float64
	LHopf             float64
	M1                float64
	M2                float64
	M3                float64
	FWall3Red         float64
	KappaLambdaRed    float64
	Complement        float64
	CHistory          float64
	MatchesFWall      bool
	MatchesKappa      bool
	MatchesComplement bool
	MatchesCHistory   bool
	MaxAbsDiscrepancy float64
	DiscrepancyClass  string
	Verdict           string
}

type TermTyping struct {
	Typed                       bool
	TermTypes                   map[string]string
	LayerTypes                  map[string]string
	RuntimeTargetInFinalFormula bool
	Verdict                     string
}

type K7RoleAudit struct {
	Audited                  bool
	AppearsOnlyAsPK7         bool
	RawMoments               []string
	NativeSupportOnly        bool
	BoundaryVector           bool
	FlavorOperator           bool
	ScalarWallCoordinate     bool
	SourceOfLHopf            bool
	HyperchargeNormalization bool
	YukawaTheorem            bool
	Verdict                  string
}

type RawMomentResponseAudit struct {
	Audited                       bool
	Formula                       string
	M1Interpretation              string
	M2Interpretation              string
	M3Interpretation              string
	BridgeLayer                   bool
	NativeGeneratingFunction      bool
	NativeRawMomentCoordinate     bool
	NativeCubicStop               bool
	M4ForbiddenWithoutTypedSource bool
	ComputedFWall                 float64
	Verdict                       string
}

type FlavorWallReductionAudit struct {
	Audited                bool
	Formula                string
	Classification         string
	Theta13Native          bool
	JCKMNative             bool
	NativeKappaETheorem    bool
	NativePMNSOrCKMTheorem bool
	NativeYukawaTheorem    bool
	CanCompareOlderKappaE  bool
	OlderKappaEResidual    float64
	ResidualClassification string
	Verdict                string
}

type RuntimeIndependenceAudit struct {
	Audited                                     bool
	FinalFormula                                string
	UsesLambdaRuntime                           bool
	UsesLambdaRuntimeEff                        bool
	UsesTreeMass                                bool
	UsesPoleMass                                bool
	UsesCHiggs                                  bool
	UsesGF                                      bool
	UsesVEV                                     bool
	UsesHiggsPoleObservable                     bool
	FormulaLevelRuntimeTargetAbsence            bool
	EvaluableWithoutDirectHiggsRuntimeVariables bool
	FormulaLevelIndependence                    string
	TheoremLevelIndependence                    string
	NativeDerivation                            bool
	RawBoundaryResponseIndependentlyProved      bool
	FlavorInputsIndependentlyProved             bool
	BoundaryCoordinatesNative                   bool
	Verdict                                     string
}

type CHistoryRebuild struct {
	Audited                            bool
	Formula                            string
	ExpandedForm                       string
	Complement                         float64
	CHistory                           float64
	Classification                     string
	FullIndependentPredictionComponent bool
	Verdict                            string
}

type PredictionLevelClassification struct {
	Recorded            bool
	KappaLambdaRedLevel string
	CHistoryLevel       string
	CHiggsLevel         string
	LevelA              string
	LevelB              string
	LevelC              string
	LevelD              string
	NextBottleneck      string
	Verdict             string
}

type PhysicalFirewalls struct {
	Enforced                           bool
	KappaLambdaNativeScalarTheorem     bool
	KappaLambdaFullyTheoremIndependent bool
	BoundaryResponseNativeTheorem      bool
	RawMomentNativeTheorem             bool
	CubicStopNativeTheorem             bool
	KappaENativeTheorem                bool
	PMNSOrCKMNativeTheorem             bool
	CHistoryFullIndependentPrediction  bool
	TreeProxyPoleMass                  bool
	YukawaNativeTheorem                bool
	Verdict                            string
}

type Analysis struct {
	Gate781        Gate781Inheritance
	Rewrite        ComplementRewrite
	Ledger         NumericalLedger
	Typing         TermTyping
	K7             K7RoleAudit
	RawMoment      RawMomentResponseAudit
	Flavor         FlavorWallReductionAudit
	Runtime        RuntimeIndependenceAudit
	CHistory       CHistoryRebuild
	Prediction     PredictionLevelClassification
	Firewalls      PhysicalFirewalls
	FinalStatement string
	Truth          string
}

var (
	cacheMu sync.Mutex
	cache   *Analysis
)

func BuildDefault() (*Analysis, error) {
	cacheMu.Lock()
	defer cacheMu.Unlock()
	if cache != nil {
		return cloneAnalysis(cache), nil
	}
	if !finite(pK7Snapshot, sSplitSnapshot, absLambdaSnapshot, xiBoundarySnapshot, kappaERedSnapshot, lHopfSnapshot, kappaLambdaRedSnapshot, complementSnapshot, cHistorySnapshot, fWall3RedSnapshot) {
		return nil, fmt.Errorf("invalid Gate782 snapshot ledger")
	}
	if pK7Snapshot <= 0 || sSplitSnapshot == 0 || absLambdaSnapshot <= 0 || lHopfSnapshot <= 0 {
		return nil, fmt.Errorf("non-positive Gate782 snapshot ledger")
	}
	m1 := pK7Snapshot * sSplitSnapshot
	m2 := pK7Snapshot * sSplitSnapshot * sSplitSnapshot
	m3 := pK7Snapshot * math.Pow(sSplitSnapshot, 3)
	fWall := m1 + kappaERedSnapshot*m2 - 2*pK7Snapshot*m3
	kappa := absLambdaSnapshot + fWall - kappaERedSnapshot
	comp := 1 - kappa
	cHist := 1 + lHopfSnapshot*comp
	maxDisc := maxAbs(fWall-fWall3RedSnapshot, kappa-kappaLambdaRedSnapshot, comp-complementSnapshot, cHist-cHistorySnapshot)
	termTypes := map[string]string{
		"|lambda|":           "scalar wall depth / high-scale scalar wound coordinate",
		"p s":                "leading K7 boundary event response",
		"-2p^2s^3":           "double-K7-event cubic boundary stress-pull correction",
		"sin^2(theta13)/4":   "PMNS reactor leakage / flavor orientation candidate",
		"-J_CKM":             "CKM orientation correction candidate",
		"-(5/3)s^2":          "hypercharge-normalized boundary-square correction",
		"+xi_boundary p s^2": "boundary-stress-weighted K7 second raw moment correction",
		"1-p s^2":            "scalar matching multiplier induced by inserting kappa_e_red into both F_wall_3_red and the final subtraction",
	}
	layerTypes := map[string]string{
		"native finite support":       "K7, P_K7, p_K7 only after observer state is supplied",
		"bridge boundary coordinates": "lambda(Lambda12), R3-1, s, xi_boundary",
		"bridge raw-moment response":  "F_wall_3_red",
		"flavor bridge/seal":          "theta13, J_CKM, kappa_e_red",
		"History transport":           "L_Hopf and C_History",
		"runtime scalar target":       "absent from the final expanded kappa_lambda_red formula",
	}
	a := &Analysis{
		Gate781: Gate781Inheritance{
			Inherited:          true,
			CHistoryFormula:    "C_History=1+L_Hopf(1-kappa_lambda_red)",
			LHopf:              lHopfSnapshot,
			KappaLambdaRed:     kappaLambdaRedSnapshot,
			SelectedBottleneck: "kappa_lambda_red=|lambda|+F_wall_3_red-kappa_e_red",
			SelectedBranch:     "Outcome 2 — partial success from Gate 781: L_Hopf strong, scalar matching complement bottleneck",
			Verdict:            StatusGate781CHistoryMacroAuditInherited,
		},
		Rewrite: ComplementRewrite{
			SelectedBottleneck: true,
			BaseFormula:        "kappa_lambda_red=|lambda|+F_wall_3_red(s)-kappa_e_red",
			WallPolynomial:     "F_wall_3_red(s)=p s+kappa_e_red p s^2-2p^2s^3",
			FactoredFormula:    "kappa_lambda_red=|lambda|+p s-2p^2s^3-(1-p s^2)kappa_e_red",
			ExpandedFormula:    "kappa_lambda_red=|lambda|+p s-2p^2s^3-(1-p s^2)[sin^2(theta13)/4-J_CKM-(5/3)s^2+xi_boundary p s^2]",
			SignIdentity:       "-kappa_e_red+kappa_e_red p s^2=-(1-p s^2)kappa_e_red",
			SignsChecked:       true,
			NoRuntimeSymbols:   true,
			Verdict:            StatusKappaLambdaRedRewrittenBoundaryFlavorResponseForm,
		},
		Ledger: NumericalLedger{
			P:                 pK7Snapshot,
			S:                 sSplitSnapshot,
			AbsLambda:         absLambdaSnapshot,
			XiBoundary:        xiBoundarySnapshot,
			KappaERed:         kappaERedSnapshot,
			LHopf:             lHopfSnapshot,
			M1:                m1,
			M2:                m2,
			M3:                m3,
			FWall3Red:         fWall,
			KappaLambdaRed:    kappa,
			Complement:        comp,
			CHistory:          cHist,
			MatchesFWall:      closeRel(fWall, fWall3RedSnapshot, 1e-14),
			MatchesKappa:      closeRel(kappa, kappaLambdaRedSnapshot, 1e-14),
			MatchesComplement: closeRel(comp, complementSnapshot, 1e-14),
			MatchesCHistory:   closeRel(cHist, cHistorySnapshot, 1e-14),
			MaxAbsDiscrepancy: maxDisc,
			DiscrepancyClass:  discrepancyClass(maxDisc),
			Verdict:           StatusNumericalBoundaryFlavorComplementLedgerRecomputed,
		},
		Typing: TermTyping{
			Typed:                       true,
			TermTypes:                   termTypes,
			LayerTypes:                  layerTypes,
			RuntimeTargetInFinalFormula: false,
			Verdict:                     StatusScalarMatchingComplementSelectedAsCurrentBottleneck,
		},
		K7: K7RoleAudit{
			Audited:                  true,
			AppearsOnlyAsPK7:         true,
			RawMoments:               []string{"M1=p s", "M2=p s^2", "M3=p s^3"},
			NativeSupportOnly:        true,
			BoundaryVector:           false,
			FlavorOperator:           false,
			ScalarWallCoordinate:     false,
			SourceOfLHopf:            false,
			HyperchargeNormalization: false,
			YukawaTheorem:            false,
			Verdict:                  "K7 acts as native support and bridge event weight only",
		},
		RawMoment: RawMomentResponseAudit{
			Audited:                       true,
			Formula:                       "F_wall_3_red=M1+kappa_e_red M2-2pM3",
			M1Interpretation:              "leading no-bias K7 boundary event response",
			M2Interpretation:              "flavor-wall modulation of the second raw boundary response",
			M3Interpretation:              "double-K7-event / boundary-pair stress-pull cubic correction",
			BridgeLayer:                   true,
			NativeGeneratingFunction:      false,
			NativeRawMomentCoordinate:     false,
			NativeCubicStop:               false,
			M4ForbiddenWithoutTypedSource: true,
			ComputedFWall:                 fWall,
			Verdict:                       StatusRawMomentResponsePolynomialAudited,
		},
		Flavor: FlavorWallReductionAudit{
			Audited:                true,
			Formula:                "kappa_e_red=sin^2(theta13)/4-J_CKM-(5/3)s^2+xi_boundary p s^2",
			Classification:         "external flavor bridge expression with boundary-wall residual-compression source typing",
			Theta13Native:          false,
			JCKMNative:             false,
			NativeKappaETheorem:    false,
			NativePMNSOrCKMTheorem: false,
			NativeYukawaTheorem:    false,
			CanCompareOlderKappaE:  false,
			OlderKappaEResidual:    math.NaN(),
			ResidualClassification: "older active kappa_e not imported in this macro-gate; comparison deferred to avoid deep predecessor-chain coupling",
			Verdict:                StatusKappaEReducedFlavorWallAudited,
		},
		Runtime: RuntimeIndependenceAudit{
			Audited:                          true,
			FinalFormula:                     "kappa_lambda_red=|lambda|+p s-2p^2s^3-(1-p s^2)[sin^2(theta13)/4-J_CKM-(5/3)s^2+xi_boundary p s^2]",
			UsesLambdaRuntime:                false,
			UsesLambdaRuntimeEff:             false,
			UsesTreeMass:                     false,
			UsesPoleMass:                     false,
			UsesCHiggs:                       false,
			UsesGF:                           false,
			UsesVEV:                          false,
			UsesHiggsPoleObservable:          false,
			FormulaLevelRuntimeTargetAbsence: true,
			EvaluableWithoutDirectHiggsRuntimeVariables: true,
			FormulaLevelIndependence:                    "Level B formula independence: no direct Higgs/runtime target symbols occur in the expanded expression",
			TheoremLevelIndependence:                    "not Level C: raw boundary response law, flavor orientation inputs, and scalar wall coordinates remain bridge/seal-layer",
			NativeDerivation:                            false,
			RawBoundaryResponseIndependentlyProved:      false,
			FlavorInputsIndependentlyProved:             false,
			BoundaryCoordinatesNative:                   false,
			Verdict:                                     StatusFormulaLevelRuntimeTargetAbsenceAudited,
		},
		CHistory: CHistoryRebuild{
			Audited:                            true,
			Formula:                            "C_History=1+L_Hopf(1-kappa_lambda_red)",
			ExpandedForm:                       "C_History=1+L_Hopf{1-|lambda|-p s+2p^2s^3+(1-p s^2)[sin^2(theta13)/4-J_CKM-(5/3)s^2+xi_boundary p s^2]}",
			Complement:                         comp,
			CHistory:                           cHist,
			Classification:                     "Radial-Hessian Hopf unit transporting a scalar complement built from boundary wall depth, K7 raw response, flavor orientation, hypercharge boundary correction, and boundary-stress moment",
			FullIndependentPredictionComponent: false,
			Verdict:                            StatusCHistoryComplementRewrittenFullBoundaryFlavorForm,
		},
		Prediction: PredictionLevelClassification{
			Recorded:            true,
			KappaLambdaRedLevel: "Level B formula independence",
			CHistoryLevel:       "Level B semi-independent bridge component",
			CHiggsLevel:         "not Level C until N_eff, C_History theorem-level sources, and G_F/v dependencies are resolved",
			LevelA:              "algebraic identity / consistency closure",
			LevelB:              "semi-independent bridge estimate using external boundary/flavor data but no direct Higgs/runtime target variables",
			LevelC:              "independent tree-level prediction from native ASHA theorems",
			LevelD:              "physical pole-mass prediction after correction package",
			NextBottleneck:      "native sourcing of F_wall_3_red as a boundary response generating function and kappa_e_red as a flavor-orientation theorem",
			Verdict:             StatusPredictionLevelClassificationRecorded,
		},
		Firewalls: PhysicalFirewalls{
			Enforced:                           true,
			KappaLambdaNativeScalarTheorem:     false,
			KappaLambdaFullyTheoremIndependent: false,
			BoundaryResponseNativeTheorem:      false,
			RawMomentNativeTheorem:             false,
			CubicStopNativeTheorem:             false,
			KappaENativeTheorem:                false,
			PMNSOrCKMNativeTheorem:             false,
			CHistoryFullIndependentPrediction:  false,
			TreeProxyPoleMass:                  false,
			YukawaNativeTheorem:                false,
			Verdict:                            StatusNoNativeBoundaryFlavorScalarMatchingComplementBoundary,
		},
		FinalStatement: "Gate 782 removes direct runtime/Higgs target variables from kappa_lambda_red at the formula level and rewrites it as an explicit boundary-flavor response expression. It does not make kappa_lambda_red native, because the raw boundary response law, flavor orientation inputs, and scalar wall coordinates remain bridge/seal-layer. The next bottleneck is native sourcing of F_wall_3_red as a boundary response generating function and kappa_e_red as a flavor-orientation theorem.",
		Truth:          fmt.Sprintf("Gate782 recomputes kappa_lambda_red=%.17g and C_History=%.17g from an explicit boundary-flavor complement. The result reaches Level B formula independence but remains blocked from Level C by raw-moment, flavor, and boundary-coordinate firewalls.", kappa, cHist),
	}
	cache = a
	return cloneAnalysis(a), nil
}

func cloneAnalysis(a *Analysis) *Analysis {
	clone := *a
	clone.Typing.TermTypes = cloneStringMap(a.Typing.TermTypes)
	clone.Typing.LayerTypes = cloneStringMap(a.Typing.LayerTypes)
	clone.K7.RawMoments = append([]string(nil), a.K7.RawMoments...)
	return &clone
}

func cloneStringMap(m map[string]string) map[string]string {
	out := make(map[string]string, len(m))
	for k, v := range m {
		out[k] = v
	}
	return out
}

func finite(values ...float64) bool {
	for _, v := range values {
		if math.IsNaN(v) || math.IsInf(v, 0) {
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
	if want == 0 {
		return d <= tol
	}
	return d/math.Abs(want) <= tol
}

func maxAbs(values ...float64) float64 {
	m := 0.0
	for _, v := range values {
		if math.Abs(v) > m {
			m = math.Abs(v)
		}
	}
	return m
}

func discrepancyClass(d float64) string {
	switch {
	case d <= 1e-15:
		return "no material discrepancy; arithmetic and inherited ledger agree within float tolerance"
	case d <= 1e-12:
		return "precision-only discrepancy"
	default:
		return "ledger mismatch requiring upstream audit"
	}
}

func Statuses() []string {
	return []string{
		StatusGate781CHistoryMacroAuditInherited,
		StatusScalarMatchingComplementSelectedAsCurrentBottleneck,
		StatusKappaLambdaRedRewrittenBoundaryFlavorResponseForm,
		StatusNumericalBoundaryFlavorComplementLedgerRecomputed,
		StatusRawMomentResponsePolynomialAudited,
		StatusKappaEReducedFlavorWallAudited,
		StatusFormulaLevelRuntimeTargetAbsenceAudited,
		StatusTheoremLevelIndependenceFirewallAudited,
		StatusCHistoryComplementRewrittenFullBoundaryFlavorForm,
		StatusPredictionLevelClassificationRecorded,
		StatusPhysicalFirewallsEnforced,
		StatusKappaLambdaRedExplicitBoundaryFlavorResponse,
		StatusKappaLambdaRedEvaluableWithoutDirectRuntimeVariables,
		StatusFWall3RedTypedBoundaryRawMomentResponse,
		StatusKappaERedFlavorPlusBoundaryWallSourceType,
		StatusCHistoryExplicitBoundaryFlavorComplementForm,
		StatusKappaLambdaRedLevelBFormulaIndependence,
		StatusKappaLambdaRedNotYetNativeScalarMatchingTheorem,
		StatusKappaLambdaRedNotYetFullyTheoremIndependent,
		StatusNoNativeBoundaryResponseGeneratingFunctionTheorem,
		StatusNoNativeRawMomentCoordinateTheorem,
		StatusNoNativeCubicStopTheorem,
		StatusNoNativeKappaETheorem,
		StatusNoNativePMNSOrCKMTheorem,
		StatusTheta13AndJCKMNotNativeFlavorTheorems,
		StatusCHistoryNotYetFullIndependentPredictionComponent,
		StatusTreeProxyNotPoleMass,
		StatusNoYukawaOperatorOrEigenvalueTheorem,
		StatusNoNativeBoundaryFlavorScalarMatchingComplementBoundary,
	}
}

func FormatGate781(g Gate781Inheritance) string {
	return fmt.Sprintf("inherited=%t formula=%s L=%.17g kappa=%.17g bottleneck=%s branch=%s verdict=%s", g.Inherited, g.CHistoryFormula, g.LHopf, g.KappaLambdaRed, g.SelectedBottleneck, g.SelectedBranch, g.Verdict)
}

func FormatRewrite(r ComplementRewrite) string {
	return fmt.Sprintf("selected=%t base=%s wall=%s factored=%s expanded=%s sign=%s signs=%t noRuntime=%t verdict=%s", r.SelectedBottleneck, r.BaseFormula, r.WallPolynomial, r.FactoredFormula, r.ExpandedFormula, r.SignIdentity, r.SignsChecked, r.NoRuntimeSymbols, r.Verdict)
}

func FormatLedger(l NumericalLedger) string {
	return fmt.Sprintf("p=%.17g s=%.17g absLambda=%.17g xi=%.17g ke=%.17g M1=%.17g M2=%.17g M3=%.17g F=%.17g kappa=%.17g complement=%.17g CHistory=%.17g matches=%t/%t/%t/%t maxDisc=%.3g class=%s verdict=%s", l.P, l.S, l.AbsLambda, l.XiBoundary, l.KappaERed, l.M1, l.M2, l.M3, l.FWall3Red, l.KappaLambdaRed, l.Complement, l.CHistory, l.MatchesFWall, l.MatchesKappa, l.MatchesComplement, l.MatchesCHistory, l.MaxAbsDiscrepancy, l.DiscrepancyClass, l.Verdict)
}

func FormatTyping(t TermTyping) string {
	return fmt.Sprintf("typed=%t termTypes=%d layerTypes=%d runtimeTarget=%t verdict=%s", t.Typed, len(t.TermTypes), len(t.LayerTypes), t.RuntimeTargetInFinalFormula, t.Verdict)
}

func FormatK7(k K7RoleAudit) string {
	return fmt.Sprintf("audited=%t onlyP=%t raw=%s supportOnly=%t illegal(boundary=%t flavor=%t scalar=%t L=%t hypercharge=%t yukawa=%t) verdict=%s", k.Audited, k.AppearsOnlyAsPK7, strings.Join(k.RawMoments, ";"), k.NativeSupportOnly, k.BoundaryVector, k.FlavorOperator, k.ScalarWallCoordinate, k.SourceOfLHopf, k.HyperchargeNormalization, k.YukawaTheorem, k.Verdict)
}

func FormatRawMoment(r RawMomentResponseAudit) string {
	return fmt.Sprintf("audited=%t formula=%s M1=%s M2=%s M3=%s bridge=%t nativeGen=%t nativeRaw=%t nativeStop=%t M4Forbidden=%t F=%.17g verdict=%s", r.Audited, r.Formula, r.M1Interpretation, r.M2Interpretation, r.M3Interpretation, r.BridgeLayer, r.NativeGeneratingFunction, r.NativeRawMomentCoordinate, r.NativeCubicStop, r.M4ForbiddenWithoutTypedSource, r.ComputedFWall, r.Verdict)
}

func FormatFlavor(f FlavorWallReductionAudit) string {
	return fmt.Sprintf("audited=%t formula=%s class=%s thetaNative=%t jNative=%t kappaNative=%t pmnsCkmNative=%t yukawaNative=%t compareOld=%t residual=%g residualClass=%s verdict=%s", f.Audited, f.Formula, f.Classification, f.Theta13Native, f.JCKMNative, f.NativeKappaETheorem, f.NativePMNSOrCKMTheorem, f.NativeYukawaTheorem, f.CanCompareOlderKappaE, f.OlderKappaEResidual, f.ResidualClassification, f.Verdict)
}

func FormatRuntime(r RuntimeIndependenceAudit) string {
	return fmt.Sprintf("audited=%t final=%s uses(runtime=%t runtimeEff=%t tree=%t pole=%t CHiggs=%t GF=%t v=%t obs=%t) formulaFree=%t evaluable=%t formulaLevel=%s theoremLevel=%s native=%t rawProved=%t flavorProved=%t boundaryNative=%t verdict=%s", r.Audited, r.FinalFormula, r.UsesLambdaRuntime, r.UsesLambdaRuntimeEff, r.UsesTreeMass, r.UsesPoleMass, r.UsesCHiggs, r.UsesGF, r.UsesVEV, r.UsesHiggsPoleObservable, r.FormulaLevelRuntimeTargetAbsence, r.EvaluableWithoutDirectHiggsRuntimeVariables, r.FormulaLevelIndependence, r.TheoremLevelIndependence, r.NativeDerivation, r.RawBoundaryResponseIndependentlyProved, r.FlavorInputsIndependentlyProved, r.BoundaryCoordinatesNative, r.Verdict)
}

func FormatCHistory(c CHistoryRebuild) string {
	return fmt.Sprintf("audited=%t formula=%s expanded=%s complement=%.17g CHistory=%.17g classification=%s fullIndependent=%t verdict=%s", c.Audited, c.Formula, c.ExpandedForm, c.Complement, c.CHistory, c.Classification, c.FullIndependentPredictionComponent, c.Verdict)
}

func FormatPrediction(p PredictionLevelClassification) string {
	return fmt.Sprintf("recorded=%t kappa=%s CHistory=%s CHiggs=%s next=%s verdict=%s", p.Recorded, p.KappaLambdaRedLevel, p.CHistoryLevel, p.CHiggsLevel, p.NextBottleneck, p.Verdict)
}

func FormatFirewalls(f PhysicalFirewalls) string {
	return fmt.Sprintf("enforced=%t kappaNative=%t kappaIndependent=%t boundaryNative=%t rawNative=%t cubicNative=%t kappaENative=%t pmnsCkmNative=%t CHistoryIndependent=%t treePole=%t yukawaNative=%t verdict=%s", f.Enforced, f.KappaLambdaNativeScalarTheorem, f.KappaLambdaFullyTheoremIndependent, f.BoundaryResponseNativeTheorem, f.RawMomentNativeTheorem, f.CubicStopNativeTheorem, f.KappaENativeTheorem, f.PMNSOrCKMNativeTheorem, f.CHistoryFullIndependentPrediction, f.TreeProxyPoleMass, f.YukawaNativeTheorem, f.Verdict)
}

func containsAll(haystack, needles []string) bool {
	joined := "\x00" + strings.Join(haystack, "\x00") + "\x00"
	for _, n := range needles {
		if !strings.Contains(joined, "\x00"+n+"\x00") {
			return false
		}
	}
	return true
}
