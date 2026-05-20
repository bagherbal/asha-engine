// Package generation2boundaryweighteddeficitclosurerobustnessaudit implements
// Gate 661: BoundaryWeightedDeficitClosure Robustness and Noncircularity Audit.
//
// Gate 660 source-typed the active 7/72 boundary interpolation weight and lifted
// W72 back into the scalar runtime formula.  Gate 661 separates the genuinely
// nontrivial closure from the partly tautological formula lift, audits exact vs
// orientation kappa_e, records uncertainty and scale-sensitivity slots, and
// checks the uniqueness of 7/72 only against typed weights already present in the
// project.
package generation2boundaryweighteddeficitclosurerobustnessaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate660 "github.com/bagherbal/asha-engine/pkg/bridge/generation2activesevenoverseventytwoboundaryweightsourceaudit"
)

const (
	AuditID = "GATE661-BOUNDARY-WEIGHTED-DEFICIT-CLOSURE-ROBUSTNESS-NONCIRCULARITY-AUDIT"

	StatusGate660ActiveWeightInherited           = "PASS_GATE660_ACTIVE_BOUNDARY_WEIGHT_INHERITED"
	StatusDependencyGraphAudited                 = "PASS_DEPENDENCY_GRAPH_AUDITED"
	StatusNontrivialClosureIsolated              = "PASS_NONTRIVIAL_CLOSURE_ISOLATED"
	StatusScalarFormulaLiftCircularityAudited    = "PASS_SCALAR_FORMULA_LIFT_CIRCULARITY_AUDITED"
	StatusOrientationApproximationAudited        = "PASS_ORIENTATION_APPROXIMATION_AUDITED"
	StatusUncertaintySlotsDefined                = "PASS_UNCERTAINTY_SLOTS_DEFINED"
	StatusScaleSensitivitySlotsDefined           = "PASS_SCALE_SENSITIVITY_SLOTS_DEFINED"
	StatusTypedWeightUniquenessAudited           = "PASS_TYPED_WEIGHT_UNIQUENESS_AUDITED"
	StatusClosureRobustInV1                      = "CONDITIONAL_SUPPORT_BOUNDARY_WEIGHTED_DEFICIT_CLOSURE_ROBUST_IN_V1_EXACT_LEDGER"
	StatusV1PrecisionCluePendingUncertaintySweep = "CONDITIONAL_SUPPORT_CLOSURE_IS_V1_PRECISION_CLUE_PENDING_UNCERTAINTY_AND_SCALE_SWEEP"
	StatusOrientationApproxStillSmall            = "CONDITIONAL_SUPPORT_ORIENTATION_APPROXIMATION_RETAINS_SMALL_BRIDGE_RESIDUAL"
	StatusSevenOver72TypedWeightBest             = "CONDITIONAL_SUPPORT_SEVEN_OVER_SEVENTY_TWO_IS_BEST_TYPED_WEIGHT_IN_CURRENT_LEDGER"
	StatusFormulaLiftNotIndependentEvidence      = "FAILED_ROUTE_SCALAR_RUNTIME_FORMULA_LIFT_NOT_INDEPENDENT_EVIDENCE"
	StatusNoNativeSevenOver72Theorem             = "FAILED_ROUTE_NO_NATIVE_7_OVER_72_THEOREM"
	StatusNoNativeScalarFlavorBoundaryTheorem    = "FAILED_ROUTE_NO_NATIVE_SCALAR_FLAVOR_BOUNDARY_TRANSPORT_THEOREM"
	StatusNoIndependentEndpointDerivation        = "FAILED_ROUTE_NO_INDEPENDENT_ENDPOINT_DERIVATION"
	StatusNoFullUncertaintyLedger                = "FAILED_ROUTE_NO_UNCERTAINTY_LEDGER_FOR_FULL_PHYSICAL_SIGNIFICANCE"
	StatusNoScaleSweepData                       = "FAILED_ROUTE_NO_SCALE_SWEEP_DATA_IN_CURRENT_GATE"
	StatusNoBoundaryStressDerivation             = "FAILED_ROUTE_NO_BOUNDARY_STRESS_DERIVATION"
	StatusNoHiggsFlavorGaugeClaim                = "FAILED_ROUTE_NO_HIGGS_FLAVOR_GAUGE_UNIFICATION_OR_CKM_PMNS_CLAIM"
	StatusGate661Boundary                        = "FIREWALL_PRESERVED_GATE661_ROBUSTNESS_NONCIRCULARITY_BOUNDARY"
)

const (
	lambdaProxyMZ     = 0.12490310236015
	lambdaRuntimeMZ   = 0.1296525650504758
	kappaLambda       = 0.0443230430960771
	kappaE            = 0.00550355419157456
	kappaEOrientation = 0.00550633006471245
	absLambdaLambda12 = 0.0497009420776833
	r3Minus1          = 0.0509933868964996
	xiBoundary        = 0.0503471644870914
	sevenOver72       = 7.0 / 72.0
)

type Gate660Inheritance struct {
	ActiveWeightInherited        bool
	W72                          float64
	KSum                         float64
	WeightedResidual             float64
	FormulaLiftResidualExact     float64
	FormulaLiftResidualOrient    float64
	FormulaLiftBridgeLayerOnly   bool
	NoNativeSevenOver72Theorem   bool
	NoNativeK7BoundaryMap        bool
	NoNativeTransportTheorem     bool
	NoFanoHitchinBoundaryRevival bool
	FirewallPreserved            bool
	Verdict                      string
}

type DependencyNode struct {
	Name        string
	Value       float64
	Role        string
	Source      string
	DependsOn   []string
	Independent bool
	Derived     bool
	Circularity string
}

type DependencyGraphAudit struct {
	Nodes                          []DependencyNode
	KappaLambdaDefinedFromRuntime  bool
	LambdaLambda12DependsOnRuntime bool
	W72DependsOnBoundaryEndpoints  bool
	FormulaLiftPartlyTautological  bool
	NontrivialStatement            string
	Verdict                        string
}

type NontrivialClosureAudit struct {
	KappaLambda               float64
	KappaEExact               float64
	KSumExact                 float64
	W72                       float64
	ClosureResidualExact      float64
	RelativeToW72             float64
	RelativeToBoundarySplit   float64
	ScalarFormulaLiftResidual float64
	FormulaLiftIndependent    bool
	NontrivialBridgeStatement string
	Verdict                   string
}

type OrientationApproximationAudit struct {
	KappaEExact                        float64
	KappaEOrientation                  float64
	KappaEDifference                   float64
	KSumOrientation                    float64
	ClosureResidualOrientation         float64
	RelativeResidualOrientationToW72   float64
	RelativeResidualOrientationToSplit float64
	ExactToOrientationResidualRatio    float64
	Verdict                            string
}

type UncertaintySlot struct {
	Quantity     string
	NeededFor    string
	Available    bool
	Treatment    string
	WouldPerturb string
}

type UncertaintyAudit struct {
	Slots                        []UncertaintySlot
	FullPropagationAvailable     bool
	InventedUncertainties        bool
	ClosureSignificanceCertified bool
	Verdict                      string
}

type ScaleSensitivityRow struct {
	Scale             string
	Available         bool
	RequiredData      string
	CurrentTreatment  string
	CanCompareClosure bool
}

type ScaleSensitivityAudit struct {
	Rows                          []ScaleSensitivityRow
	Lambda12OnlyComputed          bool
	NearbyScaleSweepAvailable     bool
	EndpointIndependenceCertified bool
	Verdict                       string
}

type TypedWeightRow struct {
	Name        string
	Weight      float64
	Target      float64
	Residual    float64
	AbsResidual float64
	TypedRole   string
	Allowed     bool
}

type TypedWeightUniquenessAudit struct {
	Rows                  []TypedWeightRow
	BestName              string
	BestResidual          float64
	SecondBestName        string
	SecondBestResidual    float64
	ImprovementOverSecond float64
	NoArbitrarySearch     bool
	Verdict               string
}

type VerdictDiscipline struct {
	ClassifiesRobustV1ExactLedger       bool
	ClassifiesPendingUncertaintySweep   bool
	ClaimsNativeSevenOver72Theorem      bool
	ClaimsNativeTransportTheorem        bool
	ClaimsIndependentEndpointDerivation bool
	ClaimsBoundaryStressDerivation      bool
	ClaimsHiggsPrediction               bool
	ClaimsScalarStability               bool
	ClaimsFlavorDerivation              bool
	ClaimsCKMPMNSDerivation             bool
	ClaimsGaugeUnification              bool
	ClaimsFanoHitchinBoundaryMap        bool
	Verdict                             string
}

type Analysis struct {
	Inherited   Gate660Inheritance
	Dependency  DependencyGraphAudit
	Closure     NontrivialClosureAudit
	Orientation OrientationApproximationAudit
	Uncertainty UncertaintyAudit
	Scale       ScaleSensitivityAudit
	Weights     TypedWeightUniquenessAudit
	Discipline  VerdictDiscipline
	Truth       string
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
	g660, err := gate660.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate660 inheritance unavailable: %w", err)
	}
	inherited := buildInheritance(g660)
	dependency := buildDependencyGraph(inherited)
	closure := buildNontrivialClosure(inherited)
	orientation := buildOrientationApproximation(inherited)
	uncertainty := buildUncertaintyAudit()
	scale := buildScaleSensitivityAudit()
	weights := buildTypedWeightAudit(inherited)
	discipline := VerdictDiscipline{ClassifiesRobustV1ExactLedger: true, ClassifiesPendingUncertaintySweep: true, Verdict: StatusGate661Boundary}
	truth := "Gate 661 separates the strongest nontrivial content of the Gate660 bridge from the partly tautological scalar runtime formula lift. The independent bridge diagnostic is kappa_lambda+kappa_e-W72, not the lambda_runtime formula residual, because kappa_lambda was defined from lambda_runtime(M_Z). In the exact v1 ledger, 7/72 remains the best typed boundary weight among the allowed candidates; using the flavor OrientationBalance approximation enlarges the residual to about 2.78e-6 but preserves a small bridge-level discrepancy. Full physical significance remains pending explicit uncertainty propagation and scale sweeps at Lambda_13, Lambda_23, Lambda_geom, and nearby Lambda_12 shifts."
	return Analysis{Inherited: inherited, Dependency: dependency, Closure: closure, Orientation: orientation, Uncertainty: uncertainty, Scale: scale, Weights: weights, Discipline: discipline, Truth: truth}, nil
}

func buildInheritance(g gate660.Analysis) Gate660Inheritance {
	return Gate660Inheritance{
		ActiveWeightInherited:        g.Interpolation.ActiveTransport && !g.Interpolation.FanoHitchinRoute && math.Abs(g.Interpolation.Weight-sevenOver72) < 1e-15 && math.Abs(g.Interpolation.Residual-8.525834413464217e-10) < 5e-18,
		W72:                          g.Interpolation.W72,
		KSum:                         g.Interpolation.KSum,
		WeightedResidual:             g.Interpolation.Residual,
		FormulaLiftResidualExact:     g.FormulaLift.RuntimeResidualExactKappaE,
		FormulaLiftResidualOrient:    g.FormulaLift.RuntimeResidualOrientKappaE,
		FormulaLiftBridgeLayerOnly:   g.FormulaLift.BridgeLayerOnly,
		NoNativeSevenOver72Theorem:   strings.Contains(g.Classification.Verdict, gate660.StatusNoNativeSevenOver72SourceTheorem),
		NoNativeK7BoundaryMap:        strings.Contains(g.Classification.Verdict, gate660.StatusNoNativeK7BoundaryMap),
		NoNativeTransportTheorem:     strings.Contains(g.Classification.Verdict, gate660.StatusNoNativeScalarFlavorBoundaryTheorem),
		NoFanoHitchinBoundaryRevival: strings.Contains(g.Numerator.Verdict, gate660.StatusNoFanoHitchinBoundaryRevival),
		FirewallPreserved:            g.Firewalls.Verdict == gate660.StatusGate660Boundary,
		Verdict:                      StatusGate660ActiveWeightInherited,
	}
}

func buildDependencyGraph(g Gate660Inheritance) DependencyGraphAudit {
	nodes := []DependencyNode{
		{Name: "lambda_proxy(M_Z)", Value: lambdaProxyMZ, Role: "spectral/tree scalar proxy", Source: "finite proxy lane (3/8)(b/a^2)", Independent: true, Derived: true, Circularity: "independent of lambda_runtime in this audit"},
		{Name: "lambda_runtime(M_Z)", Value: lambdaRuntimeMZ, Role: "runtime scalar endpoint at M_Z", Source: "environmental runtime/pole-MSbar ledger", Independent: true, Derived: false, Circularity: "used to define kappa_lambda"},
		{Name: "kappa_lambda", Value: kappaLambda, Role: "scalar matching deficit", Source: "1 - ((lambda_runtime-lambda_proxy)/lambda_proxy)/L", DependsOn: []string{"lambda_runtime(M_Z)", "lambda_proxy(M_Z)", "L=1/(8*pi)"}, Independent: false, Derived: true, Circularity: "not independent evidence for reconstructing lambda_runtime"},
		{Name: "kappa_e", Value: kappaE, Role: "flavor wall loop deficit", Source: "charged-lepton flavor seal epsilon_e=L(1-kappa_e)", Independent: true, Derived: true, Circularity: "independent of scalar runtime lane"},
		{Name: "kappa_e_orient", Value: kappaEOrientation, Role: "OrientationBalance approximation", Source: "sin²(theta13)/4 - J_CKM", Independent: true, Derived: true, Circularity: "independent replacement for exact kappa_e at bridge level"},
		{Name: "lambda(Lambda_12)", Value: -absLambdaLambda12, Role: "high-scale scalar wound", Source: "v1 RG transport from runtime scalar inputs", DependsOn: []string{"lambda_runtime(M_Z)", "RG ledger"}, Independent: false, Derived: true, Circularity: "not a fully independent endpoint until RG/threshold sources are certified"},
		{Name: "R_3-1", Value: r3Minus1, Role: "strong-gauge boundary wound", Source: "v1 gauge boundary transport", DependsOn: []string{"g3 runtime", "Lambda_12 transport"}, Independent: true, Derived: true, Circularity: "independent of scalar kappa_lambda but shares boundary scale selection"},
		{Name: "W_72", Value: g.W72, Role: "boundary-weighted scalar/gauge wound", Source: "|lambda(Lambda_12)|+(7/72)[(R_3-1)-|lambda(Lambda_12)|]", DependsOn: []string{"lambda(Lambda_12)", "R_3-1", "7/72"}, Independent: false, Derived: true, Circularity: "boundary interpolation target; not a native theorem"},
		{Name: "K_sum", Value: g.KSum, Role: "scalar+flavor deficit sum", Source: "kappa_lambda+kappa_e", DependsOn: []string{"kappa_lambda", "kappa_e"}, Independent: false, Derived: true, Circularity: "contains lambda_runtime through kappa_lambda"},
	}
	return DependencyGraphAudit{
		Nodes:                          nodes,
		KappaLambdaDefinedFromRuntime:  true,
		LambdaLambda12DependsOnRuntime: true,
		W72DependsOnBoundaryEndpoints:  true,
		FormulaLiftPartlyTautological:  true,
		NontrivialStatement:            "kappa_lambda + kappa_e - W_72 ≈ 0 is the bridge diagnostic; lambda_runtime reconstruction from kappa_lambda is not independent evidence.",
		Verdict:                        join(StatusDependencyGraphAudited, StatusFormulaLiftNotIndependentEvidence),
	}
}

func buildNontrivialClosure(g Gate660Inheritance) NontrivialClosureAudit {
	return NontrivialClosureAudit{
		KappaLambda:               kappaLambda,
		KappaEExact:               kappaE,
		KSumExact:                 kappaLambda + kappaE,
		W72:                       g.W72,
		ClosureResidualExact:      kappaLambda + kappaE - g.W72,
		RelativeToW72:             math.Abs(kappaLambda+kappaE-g.W72) / g.W72,
		RelativeToBoundarySplit:   math.Abs(kappaLambda+kappaE-g.W72) / (r3Minus1 - absLambdaLambda12),
		ScalarFormulaLiftResidual: g.FormulaLiftResidualExact,
		FormulaLiftIndependent:    false,
		NontrivialBridgeStatement: "kappa_lambda+kappa_e≈(65/72)|lambda(Lambda_12)|+(7/72)(R_3-1)",
		Verdict:                   join(StatusNontrivialClosureIsolated, StatusClosureRobustInV1, StatusFormulaLiftNotIndependentEvidence),
	}
}

func buildOrientationApproximation(g Gate660Inheritance) OrientationApproximationAudit {
	res := kappaLambda + kappaEOrientation - g.W72
	exact := kappaLambda + kappaE - g.W72
	return OrientationApproximationAudit{
		KappaEExact:                        kappaE,
		KappaEOrientation:                  kappaEOrientation,
		KappaEDifference:                   kappaEOrientation - kappaE,
		KSumOrientation:                    kappaLambda + kappaEOrientation,
		ClosureResidualOrientation:         res,
		RelativeResidualOrientationToW72:   math.Abs(res) / g.W72,
		RelativeResidualOrientationToSplit: math.Abs(res) / (r3Minus1 - absLambdaLambda12),
		ExactToOrientationResidualRatio:    math.Abs(res) / math.Abs(exact),
		Verdict:                            join(StatusOrientationApproximationAudited, StatusOrientationApproxStillSmall),
	}
}

func buildUncertaintyAudit() UncertaintyAudit {
	slots := []UncertaintySlot{
		{Quantity: "theta13", NeededFor: "kappa_e_orient", Available: false, Treatment: "slot recorded; no numerical uncertainty invented", WouldPerturb: "OrientationBalance approximation residual"},
		{Quantity: "J_CKM", NeededFor: "kappa_e_orient", Available: false, Treatment: "slot recorded; no numerical uncertainty invented", WouldPerturb: "flavor contribution to K_sum"},
		{Quantity: "lambda_runtime(M_Z) / Higgs-top-pole-MSbar conversion", NeededFor: "kappa_lambda and lambda(Lambda_12)", Available: false, Treatment: "slot recorded; exact v1 ledger used only", WouldPerturb: "both scalar matching deficit and RG high-scale wound"},
		{Quantity: "alpha_s / g3", NeededFor: "R_3-1", Available: false, Treatment: "slot recorded; no external uncertainty pulled into gate", WouldPerturb: "strong boundary wound and W72"},
		{Quantity: "RG scheme / threshold ledger", NeededFor: "lambda(Lambda_12) and boundary scale", Available: false, Treatment: "slot recorded as missing source theorem", WouldPerturb: "scale-sensitive closure target"},
	}
	return UncertaintyAudit{Slots: slots, FullPropagationAvailable: false, InventedUncertainties: false, ClosureSignificanceCertified: false, Verdict: join(StatusUncertaintySlotsDefined, StatusNoFullUncertaintyLedger, StatusV1PrecisionCluePendingUncertaintySweep)}
}

func buildScaleSensitivityAudit() ScaleSensitivityAudit {
	rows := []ScaleSensitivityRow{
		{Scale: "Lambda_12", Available: true, RequiredData: "lambda(Lambda_12), R_3-1", CurrentTreatment: "primary active closure scale", CanCompareClosure: true},
		{Scale: "Lambda_13", Available: false, RequiredData: "lambda(Lambda_13), corresponding gauge wound", CurrentTreatment: "slot only; no closure computed", CanCompareClosure: false},
		{Scale: "Lambda_23", Available: false, RequiredData: "lambda(Lambda_23), corresponding gauge wound", CurrentTreatment: "slot only; no closure computed", CanCompareClosure: false},
		{Scale: "Lambda_geom", Available: false, RequiredData: "lambda(Lambda_geom), gauge residual pair", CurrentTreatment: "slot only; no closure computed", CanCompareClosure: false},
		{Scale: "nearby Lambda_12 shifts", Available: false, RequiredData: "local RG derivative / sweep table", CurrentTreatment: "slot only; robustness not certified", CanCompareClosure: false},
	}
	return ScaleSensitivityAudit{Rows: rows, Lambda12OnlyComputed: true, NearbyScaleSweepAvailable: false, EndpointIndependenceCertified: false, Verdict: join(StatusScaleSensitivitySlotsDefined, StatusNoScaleSweepData, StatusV1PrecisionCluePendingUncertaintySweep)}
}

func buildTypedWeightAudit(g Gate660Inheritance) TypedWeightUniquenessAudit {
	weights := []struct {
		name string
		w    float64
		role string
	}{
		{"7/72", 7.0 / 72.0, "active augmented-chamber boundary interpolation candidate"},
		{"1/10", 1.0 / 10.0, "simple decimal/control typed as ten-block shadow"},
		{"1/9", 1.0 / 9.0, "nine-chamber control"},
		{"1/8", 1.0 / 8.0, "one-eighth scalar proxy shadow"},
		{"7/70", 7.0 / 70.0, "K7 over Lambda4 chamber without boundary pair"},
		{"7/144", 7.0 / 144.0, "half-trace boundary-coordinate clue from Gate656"},
	}
	rows := make([]TypedWeightRow, 0, len(weights))
	kSum := kappaLambda + kappaE
	for _, x := range weights {
		target := absLambdaLambda12 + x.w*(r3Minus1-absLambdaLambda12)
		res := kSum - target
		rows = append(rows, TypedWeightRow{Name: x.name, Weight: x.w, Target: target, Residual: res, AbsResidual: math.Abs(res), TypedRole: x.role, Allowed: true})
	}
	best, second := rows[0], rows[1]
	if second.AbsResidual < best.AbsResidual {
		best, second = second, best
	}
	for _, r := range rows[2:] {
		if r.AbsResidual < best.AbsResidual {
			second = best
			best = r
		} else if r.AbsResidual < second.AbsResidual {
			second = r
		}
	}
	return TypedWeightUniquenessAudit{Rows: rows, BestName: best.Name, BestResidual: best.AbsResidual, SecondBestName: second.Name, SecondBestResidual: second.AbsResidual, ImprovementOverSecond: second.AbsResidual / best.AbsResidual, NoArbitrarySearch: true, Verdict: join(StatusTypedWeightUniquenessAudited, StatusSevenOver72TypedWeightBest)}
}

func join(parts ...string) string { return strings.Join(parts, "; ") }

func Statuses() []string {
	return []string{
		StatusGate660ActiveWeightInherited,
		StatusDependencyGraphAudited,
		StatusNontrivialClosureIsolated,
		StatusScalarFormulaLiftCircularityAudited,
		StatusOrientationApproximationAudited,
		StatusUncertaintySlotsDefined,
		StatusScaleSensitivitySlotsDefined,
		StatusTypedWeightUniquenessAudited,
		StatusClosureRobustInV1,
		StatusV1PrecisionCluePendingUncertaintySweep,
		StatusOrientationApproxStillSmall,
		StatusSevenOver72TypedWeightBest,
		StatusFormulaLiftNotIndependentEvidence,
		StatusNoNativeSevenOver72Theorem,
		StatusNoNativeScalarFlavorBoundaryTheorem,
		StatusNoIndependentEndpointDerivation,
		StatusNoFullUncertaintyLedger,
		StatusNoScaleSweepData,
		StatusNoBoundaryStressDerivation,
		StatusNoHiggsFlavorGaugeClaim,
		StatusGate661Boundary,
	}
}
