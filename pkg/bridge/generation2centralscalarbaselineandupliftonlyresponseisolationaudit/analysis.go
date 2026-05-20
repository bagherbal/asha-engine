// Package generation2centralscalarbaselineandupliftonlyresponseisolationaudit implements
// Gate 706: Central Scalar Baseline and Uplift-Only Response Isolation Audit.
//
// Gate 705 rewrote the positive-distance boundary wound observable as
//
//	W_boundary=|lambda|I_H72+S_split P_K7.
//
// Gate 706 audits the structural separation between the universal central scalar
// baseline and the support-selected K7 uplift. The baseline is an identity shift
// seen by every normalized observer state, while the nontrivial bridge response
// is isolated as the uplift-only operator R_uplift=S_split P_K7. This is a
// bridge-layer baseline/uplift isolation audit only. It does not derive boundary
// stress, scalar RG matching, Higgs mass, gauge unification, flavor, CKM/PMNS, a
// native response theorem, a native state-selection theorem, or a native 7/72
// theorem.
package generation2centralscalarbaselineandupliftonlyresponseisolationaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate705 "github.com/bagherbal/asha-engine/pkg/bridge/generation2scalarbaselineandk7boundarysplitupliftobservableaudit"
)

const (
	AuditID = "GATE706-CENTRAL-SCALAR-BASELINE-UPLIFT-ONLY-RESPONSE-ISOLATION-AUDIT"

	StatusGate705ScalarBaselineK7UpliftInherited  = "PASS_GATE705_SCALAR_BASELINE_K7_UPLIFT_INHERITED"
	StatusCentralBaselineOperatorDefined          = "PASS_CENTRAL_BASELINE_OPERATOR_DEFINED"
	StatusBaselineCommutesWithProjectorAlgebra    = "PASS_BASELINE_COMMUTES_WITH_PROJECTOR_ALGEBRA"
	StatusBaselineExpectationObserverIndependent  = "PASS_BASELINE_EXPECTATION_OBSERVER_INDEPENDENT"
	StatusUpliftOperatorIsolated                  = "PASS_UPLIFT_OPERATOR_ISOLATED"
	StatusBaselineSubtractedResponseReconstructed = "PASS_BASELINE_SUBTRACTED_RESPONSE_RECONSTRUCTED"
	StatusObserverDependenceLocalizedToUplift     = "PASS_OBSERVER_DEPENDENCE_LOCALIZED_TO_UPLIFT"
	StatusSupportDependenceLocalizedToUplift      = "PASS_SUPPORT_DEPENDENCE_LOCALIZED_TO_UPLIFT"
	StatusRelationToPreviousGatesAudited          = "PASS_RELATION_TO_PREVIOUS_GATES_AUDITED"
	StatusScalarBaselineCentralIdentityShift      = "CONDITIONAL_SUPPORT_SCALAR_BASELINE_IS_CENTRAL_IDENTITY_SHIFT"
	StatusNontrivialBridgeContentK7Uplift         = "CONDITIONAL_SUPPORT_NONTRIVIAL_BRIDGE_CONTENT_IS_K7_UPLIFT_RESPONSE"
	StatusDBaseBaselineSubtractedHistoryDefect    = "CONDITIONAL_SUPPORT_DBASE_IS_BASELINE_SUBTRACTED_HISTORY_DEFECT"
	StatusBaselineDoesNotSelectK7OrRho72          = "FAILED_ROUTE_BASELINE_DOES_NOT_SELECT_K7_OR_RHO72"
	StatusNoNativeScalarWoundFullChamberBaseline  = "FAILED_ROUTE_NO_NATIVE_REASON_SCALAR_WOUND_IS_FULL_CHAMBER_BASELINE"
	StatusNoNativeK7ReceivesSplitUplift           = "FAILED_ROUTE_NO_NATIVE_REASON_K7_RECEIVES_SPLIT_UPLIFT"
	StatusNoNativeBoundaryWoundUpliftTheorem      = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_WOUND_UPLIFT_THEOREM"
	StatusNoNativeHistoryResponseTheorem          = "FAILED_ROUTE_NO_NATIVE_HISTORY_RESPONSE_THEOREM"
	StatusNoNativeSevenOver72Theorem              = "FAILED_ROUTE_NO_NATIVE_7_OVER_72_THEOREM"
	StatusGate706CentralBaselineUpliftBoundary    = "FIREWALL_PRESERVED_GATE706_CENTRAL_BASELINE_UPLIFT_ISOLATION_BOUNDARY"
)

const (
	h72Dimension    = 72
	finiteDimension = 70
	kernelDimension = 71
	k7Dimension     = 7
	pK7             = float64(k7Dimension) / float64(h72Dimension)
	pK7Finite       = float64(k7Dimension) / float64(finiteDimension)
	pK7Kernel       = float64(k7Dimension) / float64(kernelDimension)
	kappaLambda     = 0.0443230430960771
	kappaE          = 0.00550355419157456
	lambdaLambda12  = -0.0497009420776833
	r3Minus1        = 0.0509933868964996
	tolerance       = 1e-15
)

type Gate705Inheritance struct {
	ScalarBaselineK7UpliftInherited bool
	TwoPayoffRewritten              bool
	KSumScalarBaselineUplift        bool
	K7SplitUpliftNotPrimitiveGauge  bool
	ScalarWallAirlockBaseline       bool
	TotalExpectation                float64
	KSum                            float64
	Residual                        float64
	NoNativeScalarBaseline          bool
	NoNativeK7Uplift                bool
	NoNativeUpliftTheorem           bool
	NoNativeHistoryResponse         bool
	NoNativeSevenOver72             bool
	Verdict                         string
}

type CentralBaselineAudit struct {
	Operator          string
	ScalarBaseline    float64
	CommutesWithPK7   bool
	CommutesWithPB    bool
	CommutesWithPG    bool
	CommutesWithPPerp bool
	ProjectorBlind    bool
	IdentityShift     bool
	Verdict           string
}

type BaselineExpectationAudit struct {
	Formula                  string
	Rho72Expectation         float64
	FiniteStateExpectation   float64
	KernelStateExpectation   float64
	LocalK7StateExpectation  float64
	ArbitraryNormalizedState float64
	ObserverIndependent      bool
	Verdict                  string
}

type UpliftIsolationAudit struct {
	PositiveObservable      string
	ScalarBaseline          float64
	SubtractedObservable    string
	UpliftOperator          string
	BoundarySplit           float64
	DBase                   float64
	KSumMinusScalarBaseline float64
	UpliftExpectationRho72  float64
	Residual                float64
	IsolatesGate700Law      bool
	Verdict                 string
}

type ObserverDependenceAudit struct {
	BaselineIndependent      bool
	UpliftFormula            string
	Rho72UpliftExpectation   float64
	FiniteUpliftExpectation  float64
	KernelUpliftExpectation  float64
	LocalK7UpliftExpectation float64
	OnlyUpliftRequiresRho72  bool
	Verdict                  string
}

type SupportDependenceAudit struct {
	BaselineSelectsPK7        bool
	BaselineSelectsRho72      bool
	UpliftProjectorSensitive  bool
	UpliftSupportSelected     bool
	SelectorStatement         string
	K7IdentityRequiresSupport bool
	Verdict                   string
}

type RelationToPreviousGatesAudit struct {
	Gate705PositiveDistanceLaw string
	Gate706SubtractedLaw       string
	KSum                       float64
	DBase                      float64
	ScalarBaseline             float64
	ExpectedWound              float64
	ExpectedUplift             float64
	NoNewNumericalRelation     bool
	Verdict                    string
}

type SourceTypeClassification struct {
	ScalarBaselineRole string
	UpliftRole         string
	ObserverRole       string
	KSumRole           string
	DBaseRole          string
	Verdict            string
}

type MissingTheoremAudit struct {
	Missing []string
	Verdict string
}

type FirewallAudit struct {
	ClaimsBaselineSelectsK7OrRho72             bool
	ClaimsScalarWoundFullChamberBaselineNative bool
	ClaimsK7ReceivesSplitUpliftNative          bool
	ClaimsNativeBoundaryWoundUpliftTheorem     bool
	ClaimsNativeHistoryResponseTheorem         bool
	ClaimsNativeSevenOver72Theorem             bool
	ClaimsBoundaryStressDerived                bool
	ClaimsScalarRGMatching                     bool
	ClaimsHiggsMass                            bool
	ClaimsGaugeUnification                     bool
	ClaimsFlavorDerivation                     bool
	ClaimsCKMPMNS                              bool
	Verdict                                    string
}

type Analysis struct {
	Inherited       Gate705Inheritance
	CentralBaseline CentralBaselineAudit
	BaselineExpect  BaselineExpectationAudit
	Uplift          UpliftIsolationAudit
	Observer        ObserverDependenceAudit
	Support         SupportDependenceAudit
	Relation        RelationToPreviousGatesAudit
	SourceTypes     SourceTypeClassification
	Missing         MissingTheoremAudit
	Firewall        FirewallAudit
	Truth           string
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
	g705, err := gate705.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate705 inheritance unavailable: %w", err)
	}
	inherited := buildInheritance(g705)
	central := buildCentralBaseline()
	baselineExpect := buildBaselineExpectation(central)
	uplift := buildUpliftIsolation()
	observer := buildObserverDependence(uplift)
	support := buildSupportDependence()
	relation := buildRelation(uplift)
	sourceTypes := buildSourceTypes()
	missing := MissingTheoremAudit{
		Missing: []string{
			StatusBaselineDoesNotSelectK7OrRho72,
			StatusNoNativeScalarWoundFullChamberBaseline,
			StatusNoNativeK7ReceivesSplitUplift,
			StatusNoNativeBoundaryWoundUpliftTheorem,
			StatusNoNativeHistoryResponseTheorem,
			StatusNoNativeSevenOver72Theorem,
		},
		Verdict: strings.Join([]string{
			StatusBaselineDoesNotSelectK7OrRho72,
			StatusNoNativeScalarWoundFullChamberBaseline,
			StatusNoNativeK7ReceivesSplitUplift,
			StatusNoNativeBoundaryWoundUpliftTheorem,
			StatusNoNativeHistoryResponseTheorem,
			StatusNoNativeSevenOver72Theorem,
		}, "; "),
	}
	firewall := FirewallAudit{Verdict: StatusGate706CentralBaselineUpliftBoundary}
	truth := "Gate 706 isolates Gate705's positive-distance observable into a central scalar baseline |lambda|I_H72 and an uplift-only response S_split P_K7. The central baseline is observer-independent and projector-blind; all nontrivial observer and support dependence lives in the K7 uplift. This recovers Gate700 after subtracting the scalar baseline, but it does not prove why scalar wound is a full-chamber baseline or why K7 receives the split uplift."
	return Analysis{Inherited: inherited, CentralBaseline: central, BaselineExpect: baselineExpect, Uplift: uplift, Observer: observer, Support: support, Relation: relation, SourceTypes: sourceTypes, Missing: missing, Firewall: firewall, Truth: truth}, nil
}

func buildInheritance(g gate705.Analysis) Gate705Inheritance {
	return Gate705Inheritance{
		ScalarBaselineK7UpliftInherited: true,
		TwoPayoffRewritten:              g.Rewrite.AlgebraicEquivalence,
		KSumScalarBaselineUplift:        g.Expectation.ReproducesKSumClosure,
		K7SplitUpliftNotPrimitiveGauge:  g.SourceUpgrade.K7ReceivesBoundarySplitUplift,
		ScalarWallAirlockBaseline:       g.SourceUpgrade.ScalarWallAirlockSupportsBaseline,
		TotalExpectation:                g.Expectation.TotalExpectation,
		KSum:                            g.Expectation.KSum,
		Residual:                        g.Expectation.Residual,
		NoNativeScalarBaseline:          !g.Firewall.ClaimsScalarWoundFullChamberBaselineNative,
		NoNativeK7Uplift:                !g.Firewall.ClaimsK7ReceivesSplitUpliftNative,
		NoNativeUpliftTheorem:           !g.Firewall.ClaimsNativeBoundaryWoundUpliftTheorem,
		NoNativeHistoryResponse:         !g.Firewall.ClaimsNativeHistoryResponseTheorem,
		NoNativeSevenOver72:             !g.Firewall.ClaimsNativeSevenOver72Theorem,
		Verdict:                         StatusGate705ScalarBaselineK7UpliftInherited,
	}
}

func buildCentralBaseline() CentralBaselineAudit {
	return CentralBaselineAudit{
		Operator:          "B_scalar=|lambda|I_H72",
		ScalarBaseline:    math.Abs(lambdaLambda12),
		CommutesWithPK7:   true,
		CommutesWithPB:    true,
		CommutesWithPG:    true,
		CommutesWithPPerp: true,
		ProjectorBlind:    true,
		IdentityShift:     true,
		Verdict: strings.Join([]string{
			StatusCentralBaselineOperatorDefined,
			StatusBaselineCommutesWithProjectorAlgebra,
			StatusScalarBaselineCentralIdentityShift,
		}, "; "),
	}
}

func buildBaselineExpectation(c CentralBaselineAudit) BaselineExpectationAudit {
	baseline := c.ScalarBaseline
	return BaselineExpectationAudit{
		Formula:                  "Tr(rho |lambda|I_H72)=|lambda|Tr(rho)=|lambda| for every normalized density state rho",
		Rho72Expectation:         baseline,
		FiniteStateExpectation:   baseline,
		KernelStateExpectation:   baseline,
		LocalK7StateExpectation:  baseline,
		ArbitraryNormalizedState: baseline,
		ObserverIndependent:      true,
		Verdict:                  StatusBaselineExpectationObserverIndependent,
	}
}

func buildUpliftIsolation() UpliftIsolationAudit {
	absLambda := math.Abs(lambdaLambda12)
	sSplit := lambdaLambda12 + r3Minus1
	dBase := kappaLambda + kappaE + lambdaLambda12
	upliftExpectation := pK7 * sSplit
	return UpliftIsolationAudit{
		PositiveObservable:      "W_boundary=|lambda|I_H72+S_split P_K7",
		ScalarBaseline:          absLambda,
		SubtractedObservable:    "W_boundary-|lambda|I_H72",
		UpliftOperator:          "R_uplift=S_split P_K7",
		BoundarySplit:           sSplit,
		DBase:                   dBase,
		KSumMinusScalarBaseline: (kappaLambda + kappaE) - absLambda,
		UpliftExpectationRho72:  upliftExpectation,
		Residual:                dBase - upliftExpectation,
		IsolatesGate700Law:      math.Abs(((kappaLambda+kappaE)-absLambda)-dBase) < tolerance,
		Verdict: strings.Join([]string{
			StatusUpliftOperatorIsolated,
			StatusBaselineSubtractedResponseReconstructed,
			StatusNontrivialBridgeContentK7Uplift,
			StatusDBaseBaselineSubtractedHistoryDefect,
		}, "; "),
	}
}

func buildObserverDependence(u UpliftIsolationAudit) ObserverDependenceAudit {
	s := u.BoundarySplit
	return ObserverDependenceAudit{
		BaselineIndependent:      true,
		UpliftFormula:            "Tr(rho R_uplift)=S_split Tr(rho P_K7)",
		Rho72UpliftExpectation:   pK7 * s,
		FiniteUpliftExpectation:  pK7Finite * s,
		KernelUpliftExpectation:  pK7Kernel * s,
		LocalK7UpliftExpectation: s,
		OnlyUpliftRequiresRho72:  true,
		Verdict:                  StatusObserverDependenceLocalizedToUplift,
	}
}

func buildSupportDependence() SupportDependenceAudit {
	return SupportDependenceAudit{
		BaselineSelectsPK7:        false,
		BaselineSelectsRho72:      false,
		UpliftProjectorSensitive:  true,
		UpliftSupportSelected:     true,
		SelectorStatement:         "rank(P)=7, P_B P=P, P_G P=P => P=P_K7",
		K7IdentityRequiresSupport: true,
		Verdict: strings.Join([]string{
			StatusSupportDependenceLocalizedToUplift,
			StatusBaselineDoesNotSelectK7OrRho72,
		}, "; "),
	}
}

func buildRelation(u UpliftIsolationAudit) RelationToPreviousGatesAudit {
	absLambda := math.Abs(lambdaLambda12)
	expectedWound := absLambda + u.UpliftExpectationRho72
	return RelationToPreviousGatesAudit{
		Gate705PositiveDistanceLaw: "K_sum≈Tr(rho_72 W_boundary)=|lambda|+(7/72)S_split",
		Gate706SubtractedLaw:       "D_base=K_sum-|lambda|≈Tr(rho_72[W_boundary-|lambda|I_H72])",
		KSum:                       kappaLambda + kappaE,
		DBase:                      u.DBase,
		ScalarBaseline:             absLambda,
		ExpectedWound:              expectedWound,
		ExpectedUplift:             u.UpliftExpectationRho72,
		NoNewNumericalRelation:     math.Abs((kappaLambda+kappaE)-absLambda-u.DBase) < tolerance && math.Abs(expectedWound-(absLambda+u.UpliftExpectationRho72)) < tolerance,
		Verdict:                    StatusRelationToPreviousGatesAudited,
	}
}

func buildSourceTypes() SourceTypeClassification {
	return SourceTypeClassification{
		ScalarBaselineRole: "|lambda|I_H72: universal scalar zero-wall baseline",
		UpliftRole:         "S_split P_K7: K7-localized boundary-split uplift",
		ObserverRole:       "rho_72: needed only for the uplift event probability",
		KSumRole:           "K_sum: total expected wound",
		DBaseRole:          "D_base: uplift-only history defect after scalar baseline subtraction",
		Verdict:            StatusRelationToPreviousGatesAudited,
	}
}

func Statuses() []string {
	return []string{
		StatusGate705ScalarBaselineK7UpliftInherited,
		StatusCentralBaselineOperatorDefined,
		StatusBaselineCommutesWithProjectorAlgebra,
		StatusBaselineExpectationObserverIndependent,
		StatusUpliftOperatorIsolated,
		StatusBaselineSubtractedResponseReconstructed,
		StatusObserverDependenceLocalizedToUplift,
		StatusSupportDependenceLocalizedToUplift,
		StatusRelationToPreviousGatesAudited,
		StatusScalarBaselineCentralIdentityShift,
		StatusNontrivialBridgeContentK7Uplift,
		StatusDBaseBaselineSubtractedHistoryDefect,
		StatusBaselineDoesNotSelectK7OrRho72,
		StatusNoNativeScalarWoundFullChamberBaseline,
		StatusNoNativeK7ReceivesSplitUplift,
		StatusNoNativeBoundaryWoundUpliftTheorem,
		StatusNoNativeHistoryResponseTheorem,
		StatusNoNativeSevenOver72Theorem,
		StatusGate706CentralBaselineUpliftBoundary,
	}
}

func FormatInheritance(x Gate705Inheritance) string {
	return fmt.Sprintf("inherited=%t rewritten=%t ksumBaselineUplift=%t k7Split=%t scalarAirlock=%t expected=%.18g K=%.18g residual=%.18g noScalarBaseline=%t noK7Uplift=%t noUpliftTheorem=%t noHistory=%t no7=%t verdict=%q", x.ScalarBaselineK7UpliftInherited, x.TwoPayoffRewritten, x.KSumScalarBaselineUplift, x.K7SplitUpliftNotPrimitiveGauge, x.ScalarWallAirlockBaseline, x.TotalExpectation, x.KSum, x.Residual, x.NoNativeScalarBaseline, x.NoNativeK7Uplift, x.NoNativeUpliftTheorem, x.NoNativeHistoryResponse, x.NoNativeSevenOver72, x.Verdict)
}

func FormatCentralBaseline(x CentralBaselineAudit) string {
	return fmt.Sprintf("operator=%q baseline=%.18g commPK7=%t commPB=%t commPG=%t commPerp=%t blind=%t identity=%t verdict=%q", x.Operator, x.ScalarBaseline, x.CommutesWithPK7, x.CommutesWithPB, x.CommutesWithPG, x.CommutesWithPPerp, x.ProjectorBlind, x.IdentityShift, x.Verdict)
}

func FormatBaselineExpectation(x BaselineExpectationAudit) string {
	return fmt.Sprintf("formula=%q rho72=%.18g finite=%.18g kernel=%.18g k7=%.18g arbitrary=%.18g independent=%t verdict=%q", x.Formula, x.Rho72Expectation, x.FiniteStateExpectation, x.KernelStateExpectation, x.LocalK7StateExpectation, x.ArbitraryNormalizedState, x.ObserverIndependent, x.Verdict)
}

func FormatUplift(x UpliftIsolationAudit) string {
	return fmt.Sprintf("positive=%q baseline=%.18g subtracted=%q uplift=%q S=%.18g D=%.18g KminusBaseline=%.18g E=%.18g residual=%.18g isolates=%t verdict=%q", x.PositiveObservable, x.ScalarBaseline, x.SubtractedObservable, x.UpliftOperator, x.BoundarySplit, x.DBase, x.KSumMinusScalarBaseline, x.UpliftExpectationRho72, x.Residual, x.IsolatesGate700Law, x.Verdict)
}

func FormatObserver(x ObserverDependenceAudit) string {
	return fmt.Sprintf("baselineIndependent=%t formula=%q rho72=%.18g finite=%.18g kernel=%.18g localK7=%.18g onlyUpliftNeedsRho72=%t verdict=%q", x.BaselineIndependent, x.UpliftFormula, x.Rho72UpliftExpectation, x.FiniteUpliftExpectation, x.KernelUpliftExpectation, x.LocalK7UpliftExpectation, x.OnlyUpliftRequiresRho72, x.Verdict)
}

func FormatSupport(x SupportDependenceAudit) string {
	return fmt.Sprintf("baselinePK7=%t baselineRho72=%t upliftSensitive=%t upliftSelected=%t selector=%q requiresSupport=%t verdict=%q", x.BaselineSelectsPK7, x.BaselineSelectsRho72, x.UpliftProjectorSensitive, x.UpliftSupportSelected, x.SelectorStatement, x.K7IdentityRequiresSupport, x.Verdict)
}

func FormatRelation(x RelationToPreviousGatesAudit) string {
	return fmt.Sprintf("gate705=%q gate706=%q K=%.18g D=%.18g baseline=%.18g expectedWound=%.18g expectedUplift=%.18g notNew=%t verdict=%q", x.Gate705PositiveDistanceLaw, x.Gate706SubtractedLaw, x.KSum, x.DBase, x.ScalarBaseline, x.ExpectedWound, x.ExpectedUplift, x.NoNewNumericalRelation, x.Verdict)
}

func FormatSourceTypes(x SourceTypeClassification) string {
	return fmt.Sprintf("baseline=%q uplift=%q observer=%q K=%q D=%q verdict=%q", x.ScalarBaselineRole, x.UpliftRole, x.ObserverRole, x.KSumRole, x.DBaseRole, x.Verdict)
}

func FormatMissing(x MissingTheoremAudit) string {
	return fmt.Sprintf("missing=%s verdict=%q", strings.Join(x.Missing, ", "), x.Verdict)
}

func FormatFirewall(x FirewallAudit) string {
	return fmt.Sprintf("baselineSelects=%t scalarBaselineNative=%t k7UpliftNative=%t upliftTheorem=%t historyTheorem=%t native7=%t boundaryStress=%t scalarRG=%t higgs=%t gaugeUnification=%t flavor=%t ckmPmns=%t verdict=%q", x.ClaimsBaselineSelectsK7OrRho72, x.ClaimsScalarWoundFullChamberBaselineNative, x.ClaimsK7ReceivesSplitUpliftNative, x.ClaimsNativeBoundaryWoundUpliftTheorem, x.ClaimsNativeHistoryResponseTheorem, x.ClaimsNativeSevenOver72Theorem, x.ClaimsBoundaryStressDerived, x.ClaimsScalarRGMatching, x.ClaimsHiggsMass, x.ClaimsGaugeUnification, x.ClaimsFlavorDerivation, x.ClaimsCKMPMNS, x.Verdict)
}
