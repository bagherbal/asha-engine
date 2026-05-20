// Package generation2level1scalarruntimebridgeconsistencyestimateandnonpredictionaudit implements
// Gate 739: Level-1 Scalar Runtime Bridge Consistency Estimate and Non-Prediction Audit.
//
// Gate 738 established the minimal scalar-Higgs seal package (n,q,P_rad).
// Gate 734 stabilized the cubic scalar-runtime consistency bridge. Gate 739
// performs the permitted Level-1 estimate using explicit seals and audits the
// forecast boundary: the bridge reproduces the runtime ledger to near float
// scale, but it is not an independent scalar-runtime or Higgs-mass prediction.
package generation2level1scalarruntimebridgeconsistencyestimateandnonpredictionaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate734 "github.com/bagherbal/asha-engine/pkg/bridge/generation2cubicboundarypolynomialscalarruntimetransportandpredictionboundaryaudit"
	gate738 "github.com/bagherbal/asha-engine/pkg/bridge/generation2minimalscalarhiggssealpackageandindependenceaudit"
)

const (
	AuditID = "GATE739-LEVEL1-SCALAR-RUNTIME-BRIDGE-CONSISTENCY-ESTIMATE-NON-PREDICTION-AUDIT"

	StatusGate738MinimalSealPackageInherited        = "PASS_GATE738_MINIMAL_SCALAR_HIGGS_SEAL_PACKAGE_INHERITED"
	StatusGate734CubicScalarRuntimeBridgeInherited  = "PASS_GATE734_CUBIC_SCALAR_RUNTIME_BRIDGE_INHERITED"
	StatusLevel1ScalarRuntimeBridgeEstimateComputed = "PASS_LEVEL_1_SCALAR_RUNTIME_BRIDGE_ESTIMATE_COMPUTED"
	StatusRuntimeLedgerResidualComputed             = "PASS_RUNTIME_LEDGER_RESIDUAL_COMPUTED"
	StatusAllSealsExplicitlyLabeled                 = "PASS_ALL_SEALS_EXPLICITLY_LABELED"
	StatusNonPredictionFirewallEnforced             = "PASS_NON_PREDICTION_FIREWALL_ENFORCED"
	StatusHiggsMassFirewallEnforced                 = "PASS_HIGGS_MASS_FIREWALL_ENFORCED"

	StatusLevel1BridgeConsistencyEstimateAllowed                          = "CONDITIONAL_SUPPORT_LEVEL_1_BRIDGE_CONSISTENCY_ESTIMATE_IS_ALLOWED"
	StatusCubicBoundaryBridgeReproducesRuntimeLedgerNearFloat             = "CONDITIONAL_SUPPORT_CUBIC_BOUNDARY_BRIDGE_REPRODUCES_RUNTIME_LEDGER_TO_NEAR_FLOAT_SCALE"
	StatusScalarRuntimeLaneStructurallyOrganizedBySealedEventExpectations = "CONDITIONAL_SUPPORT_SCALAR_RUNTIME_LANE_IS_STRUCTURALLY_ORGANIZED_BY_SEALED_EVENT_EXPECTATIONS"

	StatusLevel1EstimateNotIndependentRuntimePrediction     = "FAILED_ROUTE_LEVEL_1_ESTIMATE_IS_NOT_INDEPENDENT_RUNTIME_PREDICTION"
	StatusNoNativeScalarRuntimeTheorem                      = "FAILED_ROUTE_NO_NATIVE_SCALAR_RUNTIME_THEOREM"
	StatusNoNativeHistoryLoopUnitTheorem                    = "FAILED_ROUTE_NO_NATIVE_HISTORYLOOPUNIT_THEOREM"
	StatusNoNativeRadialSelectorTheorem                     = "FAILED_ROUTE_NO_NATIVE_RADIAL_SELECTOR_THEOREM"
	StatusNoNativeBoundaryResponseGeneratingFunctionTheorem = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_RESPONSE_GENERATING_FUNCTION_THEOREM"
	StatusRuntimeLambdaBridgeNotHiggsMassTheorem            = "FAILED_ROUTE_RUNTIME_LAMBDA_BRIDGE_IS_NOT_HIGGS_MASS_THEOREM"
	StatusNoHiggsMassOrPoleMassTheorem                      = "FAILED_ROUTE_NO_HIGGS_MASS_OR_POLE_MASS_THEOREM"
	StatusNoYukawaOperatorOrEigenvalueTheorem               = "FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusGate739Boundary                                   = "FIREWALL_PRESERVED_GATE739_LEVEL1_SCALAR_RUNTIME_ESTIMATE_BOUNDARY"
)

const nearFloatRuntimeTolerance = 1e-13

type Gate738Inheritance struct {
	Inherited                bool
	PackageMinimal           bool
	SealsIndependent         bool
	RequiresThreeSealPackage bool
	NoNativeN                bool
	NoNativeQ                bool
	NoNativePRad             bool
	NoRuntimeTheorem         bool
	NoMassTheorem            bool
	NoYukawaTheorem          bool
	Verdict                  string
}

type Gate734Inheritance struct {
	Inherited                bool
	W3                       float64
	FWall3                   float64
	SSplit                   float64
	KappaE                   float64
	LambdaProxy              float64
	L                        float64
	RuntimeApprox            float64
	RuntimeExact             float64
	RuntimeTarget            float64
	PolynomialResidual       float64
	CubicRuntimeResidual     float64
	NotIndependentPrediction bool
	NoNativeRuntimeTheorem   bool
	NoMassTheorem            bool
	NoYukawaTheorem          bool
	Verdict                  string
}

type Level1Estimate struct {
	Formula           string
	SSplit            float64
	FWall3            float64
	W3                float64
	KappaE            float64
	KappaLambdaBridge float64
	LambdaProxy       float64
	L                 float64
	RuntimeBridge     float64
	RuntimeExact      float64
	RuntimeResidual   float64
	NearFloatScale    bool
	Level1Allowed     bool
	Verdict           string
}

type SealLabel struct {
	Name           string
	Role           string
	Native         bool
	SealedOrBridge bool
}

type SealLabelAudit struct {
	Labels              []SealLabel
	AllExplicit         bool
	AllRequiredByBridge bool
	Verdict             string
}

type NonPredictionFirewall struct {
	KappaLambdaDefinedFromRuntimeLedger bool
	IndependentRuntimePrediction        bool
	ConsistencyClosure                  bool
	Verdict                             string
}

type HiggsMassFirewall struct {
	RuntimeLambdaBridgeIsHiggsMassTheorem bool
	HasScalarPotentialTheorem             bool
	HasVEVOrScaleTheorem                  bool
	HasPoleMassCorrectionTheorem          bool
	HasUncertaintyPropagation             bool
	HasPhysicalMassConventionFirewall     bool
	Verdict                               string
}

type Analysis struct {
	Gate738       Gate738Inheritance
	Gate734       Gate734Inheritance
	Estimate      Level1Estimate
	Seals         SealLabelAudit
	NonPrediction NonPredictionFirewall
	HiggsMass     HiggsMassFirewall
	Truth         string
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
	g738, err := gate738.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate738 inheritance unavailable: %w", err)
	}
	g734, err := gate734.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate734 inheritance unavailable: %w", err)
	}
	inherit738 := buildGate738Inheritance(g738)
	inherit734 := buildGate734Inheritance(g734)
	estimate := buildLevel1Estimate(inherit734)
	seals := buildSealLabels()
	nonPred := buildNonPredictionFirewall()
	massFW := buildHiggsMassFirewall()
	truth := "Gate 739 performs the permitted Level-1 scalar-runtime bridge consistency estimate using the Gate734 cubic boundary polynomial and the Gate738 minimal scalar-Higgs seal package. It computes lambda_runtime_bridge=lambda_proxy[1+L(1-W_3+kappa_e)] and reproduces the runtime ledger to near float scale, but the result is a seal-dependent consistency closure because kappa_lambda/runtime data were already part of the scalar transport ledger. No independent scalar runtime, Higgs mass, HistoryLoopUnit, radial selector, boundary generating-function, or Yukawa theorem is certified."
	return Analysis{Gate738: inherit738, Gate734: inherit734, Estimate: estimate, Seals: seals, NonPrediction: nonPred, HiggsMass: massFW, Truth: truth}, nil
}

func buildGate738Inheritance(g gate738.Analysis) Gate738Inheritance {
	return Gate738Inheritance{
		Inherited:                g.Minimality.AllThreeRequired && g.Available.RuntimeBridgeCompatible && g.Remaining.AllStillBridgeOrSealed,
		PackageMinimal:           g.Minimality.AllThreeRequired,
		SealsIndependent:         g.Independence.NQTypeDistinct && g.Independence.NPRadTypeDistinct && g.Independence.QPRadTypeDistinct,
		RequiresThreeSealPackage: g.Available.NQPRadSupplied && g.Available.HistoryLoopAvailable && g.Available.RuntimeBridgeCompatible,
		NoNativeN:                !g.Firewall.NIsNativeComplexStructureTheorem,
		NoNativeQ:                !g.Firewall.QIsNativeHyperchargeDerivation,
		NoNativePRad:             !g.Firewall.PRadIsElectroweakVacuumTheorem,
		NoRuntimeTheorem:         !g.Firewall.RuntimeBridgeIsHiggsMassPrediction,
		NoMassTheorem:            !g.Firewall.PackageIsPhysicalHiggsTheorem,
		NoYukawaTheorem:          strings.Contains(g.Firewall.Verdict, gate738.StatusNoYukawaOperatorOrEigenvalueTheorem),
		Verdict:                  StatusGate738MinimalSealPackageInherited,
	}
}

func buildGate734Inheritance(g gate734.Analysis) Gate734Inheritance {
	return Gate734Inheritance{
		Inherited:                g.BoundarySub.BoundaryWoundMinusFlavor && g.Runtime.UsesCubicBoundaryWound && g.Propagation.MatchesPropagation && g.Prediction.ConsistencyClosure,
		W3:                       g.BoundarySub.W3,
		FWall3:                   g.BoundarySub.FWall3,
		SSplit:                   g.Gate733.SSplit,
		KappaE:                   g.BoundarySub.KappaE,
		LambdaProxy:              g.Runtime.LambdaProxy,
		L:                        g.Runtime.L,
		RuntimeApprox:            g.Runtime.RuntimeApprox,
		RuntimeExact:             g.Runtime.RuntimeExactTransport,
		RuntimeTarget:            g.Runtime.RuntimeTarget,
		PolynomialResidual:       g.Propagation.EPoly3,
		CubicRuntimeResidual:     g.Propagation.RuntimeResidual,
		NotIndependentPrediction: !g.Prediction.CubicRuntimeIndependentPrediction,
		NoNativeRuntimeTheorem:   !g.Firewall.ClaimsNativeScalarRuntimeTheorem,
		NoMassTheorem:            !g.Firewall.ClaimsHiggsPoleMassPrediction,
		NoYukawaTheorem:          !g.Firewall.ClaimsYukawaEigenvalueTheorem,
		Verdict:                  StatusGate734CubicScalarRuntimeBridgeInherited,
	}
}

func buildLevel1Estimate(g Gate734Inheritance) Level1Estimate {
	kappaLambdaBridge := g.W3 - g.KappaE
	runtimeBridge := g.LambdaProxy * (1 + g.L*(1-g.W3+g.KappaE))
	residual := g.RuntimeExact - runtimeBridge
	return Level1Estimate{
		Formula:           "lambda_runtime_bridge=lambda_proxy[1+L(1-W_3+kappa_e)]",
		SSplit:            g.SSplit,
		FWall3:            g.FWall3,
		W3:                g.W3,
		KappaE:            g.KappaE,
		KappaLambdaBridge: kappaLambdaBridge,
		LambdaProxy:       g.LambdaProxy,
		L:                 g.L,
		RuntimeBridge:     runtimeBridge,
		RuntimeExact:      g.RuntimeExact,
		RuntimeResidual:   residual,
		NearFloatScale:    math.Abs(residual) < nearFloatRuntimeTolerance,
		Level1Allowed:     true,
		Verdict:           strings.Join([]string{StatusLevel1ScalarRuntimeBridgeEstimateComputed, StatusCubicBoundaryBridgeReproducesRuntimeLedgerNearFloat, StatusLevel1BridgeConsistencyEstimateAllowed}, "; "),
	}
}

func buildSealLabels() SealLabelAudit {
	labels := []SealLabel{
		{Name: "n", Role: "twistor selector / complex-structure selector", SealedOrBridge: true},
		{Name: "q", Role: "hypercharge / phase normalization", SealedOrBridge: true},
		{Name: "P_rad", Role: "radial projector / scalar vacuum direction", SealedOrBridge: true},
		{Name: "rho_plus", Role: "maximum-entropy K7+ observer state", SealedOrBridge: true},
		{Name: "rho_72", Role: "maximum-entropy H72 observer state", SealedOrBridge: true},
		{Name: "P_K7", Role: "Boolean-octonionic support-selected event projector", Native: true, SealedOrBridge: true},
		{Name: "kappa_e", Role: "flavor-wall deficit bridge input", SealedOrBridge: true},
		{Name: "lambda_proxy", Role: "scalar proxy lane input", SealedOrBridge: true},
		{Name: "L", Role: "HistoryLoopUnit bridge seal / radial-Hopf source type", SealedOrBridge: true},
		{Name: "F_wall_3", Role: "cubic raw-moment boundary response closure", SealedOrBridge: true},
	}
	return SealLabelAudit{Labels: labels, AllExplicit: len(labels) == 10, AllRequiredByBridge: true, Verdict: strings.Join([]string{StatusAllSealsExplicitlyLabeled, StatusScalarRuntimeLaneStructurallyOrganizedBySealedEventExpectations}, "; ")}
}

func buildNonPredictionFirewall() NonPredictionFirewall {
	return NonPredictionFirewall{KappaLambdaDefinedFromRuntimeLedger: true, IndependentRuntimePrediction: false, ConsistencyClosure: true, Verdict: strings.Join([]string{StatusNonPredictionFirewallEnforced, StatusLevel1EstimateNotIndependentRuntimePrediction}, "; ")}
}

func buildHiggsMassFirewall() HiggsMassFirewall {
	return HiggsMassFirewall{
		RuntimeLambdaBridgeIsHiggsMassTheorem: false,
		HasScalarPotentialTheorem:             false,
		HasVEVOrScaleTheorem:                  false,
		HasPoleMassCorrectionTheorem:          false,
		HasUncertaintyPropagation:             false,
		HasPhysicalMassConventionFirewall:     true,
		Verdict:                               strings.Join([]string{StatusHiggsMassFirewallEnforced, StatusRuntimeLambdaBridgeNotHiggsMassTheorem}, "; "),
	}
}

func Statuses() []string {
	return []string{
		StatusGate738MinimalSealPackageInherited,
		StatusGate734CubicScalarRuntimeBridgeInherited,
		StatusLevel1ScalarRuntimeBridgeEstimateComputed,
		StatusRuntimeLedgerResidualComputed,
		StatusAllSealsExplicitlyLabeled,
		StatusNonPredictionFirewallEnforced,
		StatusHiggsMassFirewallEnforced,
		StatusLevel1BridgeConsistencyEstimateAllowed,
		StatusCubicBoundaryBridgeReproducesRuntimeLedgerNearFloat,
		StatusScalarRuntimeLaneStructurallyOrganizedBySealedEventExpectations,
		StatusLevel1EstimateNotIndependentRuntimePrediction,
		StatusNoNativeScalarRuntimeTheorem,
		StatusNoNativeHistoryLoopUnitTheorem,
		StatusNoNativeRadialSelectorTheorem,
		StatusNoNativeBoundaryResponseGeneratingFunctionTheorem,
		StatusRuntimeLambdaBridgeNotHiggsMassTheorem,
		StatusNoHiggsMassOrPoleMassTheorem,
		StatusNoYukawaOperatorOrEigenvalueTheorem,
		StatusGate739Boundary,
	}
}

func FormatGate738(x Gate738Inheritance) string {
	return fmt.Sprintf("inherited=%t minimal=%t independent=%t requiresThree=%t noN=%t noQ=%t noPRad=%t noRuntime=%t noMass=%t noYukawa=%t verdict=%q", x.Inherited, x.PackageMinimal, x.SealsIndependent, x.RequiresThreeSealPackage, x.NoNativeN, x.NoNativeQ, x.NoNativePRad, x.NoRuntimeTheorem, x.NoMassTheorem, x.NoYukawaTheorem, x.Verdict)
}

func FormatGate734(x Gate734Inheritance) string {
	return fmt.Sprintf("inherited=%t S=%.17g F3=%.17g W3=%.17g kappaE=%.17g runtimeApprox=%.17g runtimeExact=%.17g polyResidual=%.17g cubicRuntimeResidual=%.17g verdict=%q", x.Inherited, x.SSplit, x.FWall3, x.W3, x.KappaE, x.RuntimeApprox, x.RuntimeExact, x.PolynomialResidual, x.CubicRuntimeResidual, x.Verdict)
}

func FormatEstimate(x Level1Estimate) string {
	return fmt.Sprintf("formula=%q S=%.17g F3=%.17g W3=%.17g kappaLambdaBridge=%.17g runtimeBridge=%.17g runtimeExact=%.17g residual=%.17g nearFloat=%t verdict=%q", x.Formula, x.SSplit, x.FWall3, x.W3, x.KappaLambdaBridge, x.RuntimeBridge, x.RuntimeExact, x.RuntimeResidual, x.NearFloatScale, x.Verdict)
}

func FormatSeals(x SealLabelAudit) string {
	names := make([]string, 0, len(x.Labels))
	for _, l := range x.Labels {
		names = append(names, l.Name+":"+l.Role)
	}
	return fmt.Sprintf("labels=[%s] explicit=%t required=%t verdict=%q", strings.Join(names, " | "), x.AllExplicit, x.AllRequiredByBridge, x.Verdict)
}

func FormatNonPrediction(x NonPredictionFirewall) string {
	return fmt.Sprintf("kappaFromRuntime=%t independentPrediction=%t consistency=%t verdict=%q", x.KappaLambdaDefinedFromRuntimeLedger, x.IndependentRuntimePrediction, x.ConsistencyClosure, x.Verdict)
}

func FormatHiggsMass(x HiggsMassFirewall) string {
	return fmt.Sprintf("runtimeIsMass=%t scalarPotential=%t vev=%t pole=%t uncertainty=%t conventionFirewall=%t verdict=%q", x.RuntimeLambdaBridgeIsHiggsMassTheorem, x.HasScalarPotentialTheorem, x.HasVEVOrScaleTheorem, x.HasPoleMassCorrectionTheorem, x.HasUncertaintyPropagation, x.HasPhysicalMassConventionFirewall, x.Verdict)
}
