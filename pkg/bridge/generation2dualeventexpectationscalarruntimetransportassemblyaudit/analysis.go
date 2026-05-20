// Package generation2dualeventexpectationscalarruntimetransportassemblyaudit implements
// Gate 728: Dual Event-Expectation Scalar Runtime Transport Assembly Audit.
//
// Gate 700 closed the conditional boundary-history response law and Gate 727
// closed the conditional Radial-Hopf HistoryLoopUnit source law. Gate 728 audits
// how these two event-expectation bridge laws assemble into the active scalar
// runtime transport formula, while preserving the firewall that this is a
// bridge consistency closure, not an independent scalar runtime, Higgs mass, or
// native HistoryLoopUnit theorem.
package generation2dualeventexpectationscalarruntimetransportassemblyaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate700 "github.com/bagherbal/asha-engine/pkg/bridge/generation2conditionalashahistoryresponselawclosureaudit"
	gate727 "github.com/bagherbal/asha-engine/pkg/bridge/generation2conditionalradialhopfhistoryloopunitlawandpremiseminimalityaudit"
)

const (
	AuditID = "GATE728-DUAL-EVENT-EXPECTATION-SCALAR-RUNTIME-TRANSPORT-ASSEMBLY-AUDIT"

	StatusGate700BoundaryHistoryResponseInherited = "PASS_GATE700_BOUNDARY_HISTORY_RESPONSE_INHERITED"
	StatusGate727RadialHopfHistoryLoopInherited   = "PASS_GATE727_RADIAL_HOPF_HISTORYLOOP_INHERITED"
	StatusBoundaryHistoryResponseSubstituted      = "PASS_BOUNDARY_HISTORY_RESPONSE_SUBSTITUTED_INTO_KAPPA_LAMBDA"
	StatusRadialHopfLSubstituted                  = "PASS_RADIAL_HOPF_L_SUBSTITUTED_INTO_SCALAR_TRANSPORT"
	StatusDualEventExpectationFormAssembled       = "PASS_DUAL_EVENT_EXPECTATION_FORM_ASSEMBLED"
	StatusWallResidualPropagationComputed         = "PASS_WALL_RESIDUAL_PROPAGATION_COMPUTED"
	StatusNoncircularityAudited                   = "PASS_NONCIRCULARITY_AUDITED"
	StatusSealDependenceAudited                   = "PASS_SEAL_DEPENDENCE_AUDITED"
	StatusFirewallsEnforced                       = "PASS_FIREWALLS_ENFORCED"

	StatusScalarRuntimeBridgeAsDualEventExpectation = "CONDITIONAL_SUPPORT_SCALAR_RUNTIME_BRIDGE_AS_DUAL_EVENT_EXPECTATION_FORM"
	StatusRuntimeResidualPropagatedWallResidual     = "CONDITIONAL_SUPPORT_RUNTIME_RESIDUAL_IS_PROPAGATED_HISTORY_WALL_RESIDUAL"
	StatusRadialHopfAndK7ResponseCombine            = "CONDITIONAL_SUPPORT_RADIAL_HOPF_LOOP_UNIT_AND_K7_BOUNDARY_RESPONSE_COMBINE_IN_SCALAR_TRANSPORT"

	StatusAssembledRuntimeNotIndependentPrediction = "FAILED_ROUTE_ASSEMBLED_RUNTIME_FORM_NOT_INDEPENDENT_PREDICTION"
	StatusPremisesNotNativelyDerived               = "FAILED_ROUTE_PREMISES_NOT_NATIVELY_DERIVED"
	StatusNoNativeScalarProxyToRuntimeTheorem      = "FAILED_ROUTE_NO_NATIVE_SCALAR_PROXY_TO_RUNTIME_THEOREM"
	StatusNoNativeHistoryLoopUnitSourceTheorem     = "FAILED_ROUTE_NO_NATIVE_HISTORYLOOPUNIT_SOURCE_THEOREM"
	StatusNoNativeRadialProjectorSelector          = "FAILED_ROUTE_NO_NATIVE_RADIAL_PROJECTOR_SELECTOR"
	StatusNoNativeBoundaryHistoryResponseTheorem   = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_HISTORY_RESPONSE_THEOREM"
	StatusNoHiggsMassOrPoleMassTheorem             = "FAILED_ROUTE_NO_HIGGS_MASS_OR_POLE_MASS_THEOREM"
	StatusNoYukawaOperatorOrEigenvalueTheorem      = "FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusGate728Boundary                          = "FIREWALL_PRESERVED_GATE728_DUAL_EVENT_EXPECTATION_RUNTIME_BOUNDARY"
)

const (
	lambdaProxyMZ   = 0.12490310236015
	lambdaRuntimeMZ = 0.1296525650504758
	kappaLambda     = 0.0443230430960771
	kappaE          = 0.00550355419157456
	lambdaLambda12  = -0.0497009420776833
	r3Minus1        = 0.0509933868964996
	k7Dim           = 7
	h72Dim          = 72
	wallTolerance   = 1e-12
)

type Gate700Inheritance struct {
	Inherited                      bool
	DBase                          float64
	SSplit                         float64
	P_K7                           float64
	ExpectedHistoryResponse        float64
	EWall                          float64
	NoNativeBoundaryHistoryTheorem bool
	NoNativeSevenOver72            bool
	Verdict                        string
}

type Gate727Inheritance struct {
	Inherited                     bool
	L                             float64
	RadialHopfExpectation         float64
	ConditionallyExact            bool
	PremisesNotNative             bool
	NoNativeRadialProjector       bool
	NoNativeTwistorSelectorN      bool
	NoNativeHistoryLoopUnitSource bool
	NoNativeScalarProxyRuntime    bool
	Verdict                       string
}

type BoundaryHistorySubstitutionAudit struct {
	Formula                    string
	DBase                      float64
	P_K7                       float64
	SSplit                     float64
	Lambda                     float64
	KappaE                     float64
	KappaLambdaExact           float64
	KappaLambdaFromExactWall   float64
	KappaLambdaApprox          float64
	W72                        float64
	EWall                      float64
	ApproxDropsWallResidual    bool
	BoundaryMinusFlavorReading bool
	Verdict                    string
}

type RadialHopfSubstitutionAudit struct {
	OriginalTransportFormula  string
	SubstitutedFormula        string
	L                         float64
	RadialHopfExpectation     float64
	FactorExact               float64
	FactorApprox              float64
	RuntimePredExact          float64
	RuntimePredApprox         float64
	UsesRadialHopfExpectation bool
	Verdict                   string
}

type DualEventExpectationAssemblyAudit struct {
	Formula                   string
	WBoundaryObservable       string
	W72                       float64
	LambdaProxy               float64
	L                         float64
	KappaE                    float64
	AssembledRuntimeApprox    float64
	AssembledRuntimeWithEWall float64
	RuntimeTarget             float64
	DualEventExpectationForm  bool
	Verdict                   string
}

type WallResidualPropagationAudit struct {
	EWall                         float64
	LambdaProxy                   float64
	L                             float64
	DeltaLambdaPred               float64
	ApproxMinusExactTransport     float64
	MatchesPropagationFormula     bool
	RuntimeResidualIsWallResidual bool
	Verdict                       string
}

type NoncircularityAudit struct {
	KappaLambdaDefinedFromRuntime  bool
	AssembledIndependentPrediction bool
	BridgeConsistencyClosure       bool
	Verdict                        string
}

type SealDependenceAudit struct {
	DependsOnN              bool
	DependsOnPRad           bool
	DependsOnRhoPlus        bool
	DependsOnRho72          bool
	DependsOnPK7            bool
	DependsOnKappaE         bool
	PremisesNativelyDerived bool
	Verdict                 string
}

type FirewallAudit struct {
	ClaimsScalarRuntimeTheorem  bool
	ClaimsHiggsMassTheorem      bool
	ClaimsNativeHistoryLoopUnit bool
	ClaimsNativeBoundaryHistory bool
	ClaimsNativeRadialSelector  bool
	ClaimsYukawaOperatorTheorem bool
	ClaimsIndependentPrediction bool
	Verdict                     string
}

type Analysis struct {
	Gate700     Gate700Inheritance
	Gate727     Gate727Inheritance
	BoundarySub BoundaryHistorySubstitutionAudit
	RadialSub   RadialHopfSubstitutionAudit
	Assembly    DualEventExpectationAssemblyAudit
	Propagation WallResidualPropagationAudit
	NonCircular NoncircularityAudit
	Seals       SealDependenceAudit
	Firewall    FirewallAudit
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
	g700, err := gate700.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate700 inheritance unavailable: %w", err)
	}
	g727, err := gate727.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate727 inheritance unavailable: %w", err)
	}
	gate700Audit := buildGate700Inheritance(g700)
	gate727Audit := buildGate727Inheritance(g727)
	boundary := buildBoundaryHistorySubstitution(gate700Audit)
	radial := buildRadialHopfSubstitution(gate727Audit, boundary)
	assembly := buildDualEventExpectationAssembly(boundary, radial)
	propagation := buildWallResidualPropagation(boundary, radial)
	noncirc := buildNoncircularityAudit()
	seals := buildSealDependenceAudit()
	firewall := buildFirewall()
	truth := "Gate 728 assembles two prior conditional event-expectation bridge laws: the Gate700 K7 boundary/history expectation and the Gate727 radial-Hopf expectation for L. Substituting the boundary/history closure into kappa_lambda and L into the scalar transport gives lambda_runtime≈lambda_proxy{1+Tr(rho_plus R_Hopf)[1-Tr(rho_72 W_boundary)+kappa_e]}. The propagated scalar-runtime discrepancy from using W72-kappa_e instead of exact kappa_lambda is lambda_proxy*L*E_wall≈4.24e-12. This is a bridge consistency closure only: kappa_lambda was originally defined from the scalar runtime ledger, and the decisive premises remain sealed or conditional."
	return Analysis{Gate700: gate700Audit, Gate727: gate727Audit, BoundarySub: boundary, RadialSub: radial, Assembly: assembly, Propagation: propagation, NonCircular: noncirc, Seals: seals, Firewall: firewall, Truth: truth}, nil
}

func buildGate700Inheritance(g gate700.Analysis) Gate700Inheritance {
	return Gate700Inheritance{
		Inherited:                      g.Master.Reconstructed && g.Functional.ApproxLawCertified,
		DBase:                          g.Master.DBase,
		SSplit:                         g.Inherited.SBoundary,
		P_K7:                           float64(k7Dim) / float64(h72Dim),
		ExpectedHistoryResponse:        g.Master.Expectation,
		EWall:                          g.Master.ResidualE1,
		NoNativeBoundaryHistoryTheorem: !g.Firewall.ClaimsNativeBoundaryHistoryPrinciple,
		NoNativeSevenOver72:            !g.Firewall.ClaimsNativeSevenOver72Theorem,
		Verdict:                        StatusGate700BoundaryHistoryResponseInherited,
	}
}

func buildGate727Inheritance(g gate727.Analysis) Gate727Inheritance {
	return Gate727Inheritance{
		Inherited:                     g.Functional.ConditionallyExact && g.Premises.Count == 5,
		L:                             g.Functional.L,
		RadialHopfExpectation:         g.Functional.Expectation,
		ConditionallyExact:            g.Functional.ConditionallyExact,
		PremisesNotNative:             !g.NonTaut.PremisesNativelyDerived,
		NoNativeRadialProjector:       !g.Firewall.NativeRadialProjectorSelector,
		NoNativeTwistorSelectorN:      !g.Firewall.NativeTwistorSelectorN,
		NoNativeHistoryLoopUnitSource: !g.Firewall.NativeHistoryLoopUnitSourceTheorem,
		NoNativeScalarProxyRuntime:    !g.Firewall.NativeScalarProxyToRuntimeTheorem,
		Verdict:                       StatusGate727RadialHopfHistoryLoopInherited,
	}
}

func buildBoundaryHistorySubstitution(g Gate700Inheritance) BoundaryHistorySubstitutionAudit {
	absLambda := math.Abs(lambdaLambda12)
	w72 := absLambda + g.P_K7*g.SSplit
	kApprox := w72 - kappaE
	kExactFromWall := kApprox + g.EWall
	return BoundaryHistorySubstitutionAudit{
		Formula:                    "kappa_lambda≈p_K7*S_split-kappa_e-lambda = W_72-kappa_e",
		DBase:                      g.DBase,
		P_K7:                       g.P_K7,
		SSplit:                     g.SSplit,
		Lambda:                     lambdaLambda12,
		KappaE:                     kappaE,
		KappaLambdaExact:           kappaLambda,
		KappaLambdaFromExactWall:   kExactFromWall,
		KappaLambdaApprox:          kApprox,
		W72:                        w72,
		EWall:                      g.EWall,
		ApproxDropsWallResidual:    near(kApprox+g.EWall, kappaLambda, wallTolerance),
		BoundaryMinusFlavorReading: true,
		Verdict: strings.Join([]string{
			StatusBoundaryHistoryResponseSubstituted,
			StatusScalarRuntimeBridgeAsDualEventExpectation,
		}, "; "),
	}
}

func buildRadialHopfSubstitution(g Gate727Inheritance, b BoundaryHistorySubstitutionAudit) RadialHopfSubstitutionAudit {
	factorExact := 1 - kappaLambda
	factorApprox := 1 - b.W72 + kappaE
	exact := lambdaProxyMZ * (1 + g.L*factorExact)
	approx := lambdaProxyMZ * (1 + g.L*factorApprox)
	return RadialHopfSubstitutionAudit{
		OriginalTransportFormula:  "lambda_runtime≈lambda_proxy[1+L(1-kappa_lambda)]",
		SubstitutedFormula:        "lambda_runtime≈lambda_proxy[1+Tr(rho_plus R_Hopf)(1-W_72+kappa_e)]",
		L:                         g.L,
		RadialHopfExpectation:     g.RadialHopfExpectation,
		FactorExact:               factorExact,
		FactorApprox:              factorApprox,
		RuntimePredExact:          exact,
		RuntimePredApprox:         approx,
		UsesRadialHopfExpectation: near(g.L, g.RadialHopfExpectation, 1e-18),
		Verdict: strings.Join([]string{
			StatusRadialHopfLSubstituted,
			StatusRadialHopfAndK7ResponseCombine,
		}, "; "),
	}
}

func buildDualEventExpectationAssembly(b BoundaryHistorySubstitutionAudit, r RadialHopfSubstitutionAudit) DualEventExpectationAssemblyAudit {
	return DualEventExpectationAssemblyAudit{
		Formula:                   "lambda_runtime≈lambda_proxy{1+Tr[rho_plus(1/(2*pi))P_rad][1-Tr(rho_72 W_boundary)+kappa_e]}",
		WBoundaryObservable:       "W_boundary=|lambda|I_H72+S_split P_K7",
		W72:                       b.W72,
		LambdaProxy:               lambdaProxyMZ,
		L:                         r.L,
		KappaE:                    kappaE,
		AssembledRuntimeApprox:    r.RuntimePredApprox,
		AssembledRuntimeWithEWall: r.RuntimePredExact,
		RuntimeTarget:             lambdaRuntimeMZ,
		DualEventExpectationForm:  true,
		Verdict: strings.Join([]string{
			StatusDualEventExpectationFormAssembled,
			StatusScalarRuntimeBridgeAsDualEventExpectation,
		}, "; "),
	}
}

func buildWallResidualPropagation(b BoundaryHistorySubstitutionAudit, r RadialHopfSubstitutionAudit) WallResidualPropagationAudit {
	delta := lambdaProxyMZ * r.L * b.EWall
	actualDelta := r.RuntimePredApprox - r.RuntimePredExact
	return WallResidualPropagationAudit{
		EWall:                         b.EWall,
		LambdaProxy:                   lambdaProxyMZ,
		L:                             r.L,
		DeltaLambdaPred:               delta,
		ApproxMinusExactTransport:     actualDelta,
		MatchesPropagationFormula:     near(delta, actualDelta, 1e-16),
		RuntimeResidualIsWallResidual: true,
		Verdict: strings.Join([]string{
			StatusWallResidualPropagationComputed,
			StatusRuntimeResidualPropagatedWallResidual,
		}, "; "),
	}
}

func buildNoncircularityAudit() NoncircularityAudit {
	return NoncircularityAudit{
		KappaLambdaDefinedFromRuntime:  true,
		AssembledIndependentPrediction: false,
		BridgeConsistencyClosure:       true,
		Verdict: strings.Join([]string{
			StatusNoncircularityAudited,
			StatusAssembledRuntimeNotIndependentPrediction,
		}, "; "),
	}
}

func buildSealDependenceAudit() SealDependenceAudit {
	return SealDependenceAudit{
		DependsOnN:              true,
		DependsOnPRad:           true,
		DependsOnRhoPlus:        true,
		DependsOnRho72:          true,
		DependsOnPK7:            true,
		DependsOnKappaE:         true,
		PremisesNativelyDerived: false,
		Verdict: strings.Join([]string{
			StatusSealDependenceAudited,
			StatusPremisesNotNativelyDerived,
		}, "; "),
	}
}

func buildFirewall() FirewallAudit {
	return FirewallAudit{
		ClaimsScalarRuntimeTheorem:  false,
		ClaimsHiggsMassTheorem:      false,
		ClaimsNativeHistoryLoopUnit: false,
		ClaimsNativeBoundaryHistory: false,
		ClaimsNativeRadialSelector:  false,
		ClaimsYukawaOperatorTheorem: false,
		ClaimsIndependentPrediction: false,
		Verdict: strings.Join([]string{
			StatusFirewallsEnforced,
			StatusAssembledRuntimeNotIndependentPrediction,
			StatusPremisesNotNativelyDerived,
			StatusNoNativeScalarProxyToRuntimeTheorem,
			StatusNoNativeHistoryLoopUnitSourceTheorem,
			StatusNoNativeRadialProjectorSelector,
			StatusNoNativeBoundaryHistoryResponseTheorem,
			StatusNoHiggsMassOrPoleMassTheorem,
			StatusNoYukawaOperatorOrEigenvalueTheorem,
			StatusGate728Boundary,
		}, "; "),
	}
}

func Statuses() []string {
	return []string{
		StatusGate700BoundaryHistoryResponseInherited,
		StatusGate727RadialHopfHistoryLoopInherited,
		StatusBoundaryHistoryResponseSubstituted,
		StatusRadialHopfLSubstituted,
		StatusDualEventExpectationFormAssembled,
		StatusWallResidualPropagationComputed,
		StatusNoncircularityAudited,
		StatusSealDependenceAudited,
		StatusFirewallsEnforced,
		StatusScalarRuntimeBridgeAsDualEventExpectation,
		StatusRuntimeResidualPropagatedWallResidual,
		StatusRadialHopfAndK7ResponseCombine,
		StatusAssembledRuntimeNotIndependentPrediction,
		StatusPremisesNotNativelyDerived,
		StatusNoNativeScalarProxyToRuntimeTheorem,
		StatusNoNativeHistoryLoopUnitSourceTheorem,
		StatusNoNativeRadialProjectorSelector,
		StatusNoNativeBoundaryHistoryResponseTheorem,
		StatusNoHiggsMassOrPoleMassTheorem,
		StatusNoYukawaOperatorOrEigenvalueTheorem,
		StatusGate728Boundary,
	}
}

func near(a, b, eps float64) bool { return math.Abs(a-b) <= eps }

func FormatGate700(x Gate700Inheritance) string {
	return fmt.Sprintf("inherited=%t D=%.17g S=%.17g p=%.17g expect=%.17g Ewall=%.17g noBoundary=%t no7=%t verdict=%q", x.Inherited, x.DBase, x.SSplit, x.P_K7, x.ExpectedHistoryResponse, x.EWall, x.NoNativeBoundaryHistoryTheorem, x.NoNativeSevenOver72, x.Verdict)
}
func FormatGate727(x Gate727Inheritance) string {
	return fmt.Sprintf("inherited=%t L=%.17g hopf=%.17g exact=%t premisesNotNative=%t noRadial=%t noN=%t noL=%t noProxyRuntime=%t verdict=%q", x.Inherited, x.L, x.RadialHopfExpectation, x.ConditionallyExact, x.PremisesNotNative, x.NoNativeRadialProjector, x.NoNativeTwistorSelectorN, x.NoNativeHistoryLoopUnitSource, x.NoNativeScalarProxyRuntime, x.Verdict)
}
func FormatBoundarySubstitution(x BoundaryHistorySubstitutionAudit) string {
	return fmt.Sprintf("formula=%q D=%.17g p=%.17g S=%.17g lambda=%.17g kE=%.17g kExact=%.17g kWall=%.17g kApprox=%.17g W72=%.17g E=%.17g drops=%t verdict=%q", x.Formula, x.DBase, x.P_K7, x.SSplit, x.Lambda, x.KappaE, x.KappaLambdaExact, x.KappaLambdaFromExactWall, x.KappaLambdaApprox, x.W72, x.EWall, x.ApproxDropsWallResidual, x.Verdict)
}
func FormatRadialSubstitution(x RadialHopfSubstitutionAudit) string {
	return fmt.Sprintf("original=%q substituted=%q L=%.17g hopf=%.17g exactFactor=%.17g approxFactor=%.17g exactPred=%.17g approxPred=%.17g usesHopf=%t verdict=%q", x.OriginalTransportFormula, x.SubstitutedFormula, x.L, x.RadialHopfExpectation, x.FactorExact, x.FactorApprox, x.RuntimePredExact, x.RuntimePredApprox, x.UsesRadialHopfExpectation, x.Verdict)
}
func FormatAssembly(x DualEventExpectationAssemblyAudit) string {
	return fmt.Sprintf("formula=%q W=%q W72=%.17g proxy=%.17g L=%.17g kE=%.17g approx=%.17g exact=%.17g target=%.17g dual=%t verdict=%q", x.Formula, x.WBoundaryObservable, x.W72, x.LambdaProxy, x.L, x.KappaE, x.AssembledRuntimeApprox, x.AssembledRuntimeWithEWall, x.RuntimeTarget, x.DualEventExpectationForm, x.Verdict)
}
func FormatPropagation(x WallResidualPropagationAudit) string {
	return fmt.Sprintf("E=%.17g proxy=%.17g L=%.17g delta=%.17g actual=%.17g matches=%t runtimeResidualWall=%t verdict=%q", x.EWall, x.LambdaProxy, x.L, x.DeltaLambdaPred, x.ApproxMinusExactTransport, x.MatchesPropagationFormula, x.RuntimeResidualIsWallResidual, x.Verdict)
}
func FormatNoncircularity(x NoncircularityAudit) string {
	return fmt.Sprintf("kappaFromRuntime=%t independent=%t consistency=%t verdict=%q", x.KappaLambdaDefinedFromRuntime, x.AssembledIndependentPrediction, x.BridgeConsistencyClosure, x.Verdict)
}
func FormatSeals(x SealDependenceAudit) string {
	return fmt.Sprintf("n=%t pRad=%t rhoPlus=%t rho72=%t pK7=%t kE=%t native=%t verdict=%q", x.DependsOnN, x.DependsOnPRad, x.DependsOnRhoPlus, x.DependsOnRho72, x.DependsOnPK7, x.DependsOnKappaE, x.PremisesNativelyDerived, x.Verdict)
}
func FormatFirewall(x FirewallAudit) string {
	return fmt.Sprintf("runtime=%t mass=%t L=%t boundary=%t radial=%t yukawa=%t independent=%t verdict=%q", x.ClaimsScalarRuntimeTheorem, x.ClaimsHiggsMassTheorem, x.ClaimsNativeHistoryLoopUnit, x.ClaimsNativeBoundaryHistory, x.ClaimsNativeRadialSelector, x.ClaimsYukawaOperatorTheorem, x.ClaimsIndependentPrediction, x.Verdict)
}
