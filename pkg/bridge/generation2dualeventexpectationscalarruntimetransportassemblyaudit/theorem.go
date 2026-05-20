package generation2dualeventexpectationscalarruntimetransportassemblyaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2DualEventExpectationScalarRuntimeTransportAssemblyAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 728 — Dual Event-Expectation Scalar Runtime Transport Assembly Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate728 dual event-expectation runtime audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate700 boundary-history response law", Passed: a.Gate700.Inherited && near(a.Gate700.P_K7, 7.0/72.0, 1e-18) && near(a.Gate700.DBase-a.Gate700.ExpectedHistoryResponse, a.Gate700.EWall, 1e-18) && a.Gate700.NoNativeBoundaryHistoryTheorem && a.Gate700.NoNativeSevenOver72, Detail: FormatGate700(a.Gate700)},
			{Name: "inherit Gate727 radial-Hopf HistoryLoop law", Passed: a.Gate727.Inherited && near(a.Gate727.L, 1/(8*math.Pi), 1e-18) && near(a.Gate727.L, a.Gate727.RadialHopfExpectation, 1e-18) && a.Gate727.ConditionallyExact && a.Gate727.PremisesNotNative && a.Gate727.NoNativeRadialProjector && a.Gate727.NoNativeTwistorSelectorN && a.Gate727.NoNativeHistoryLoopUnitSource && a.Gate727.NoNativeScalarProxyRuntime, Detail: FormatGate727(a.Gate727)},
			{Name: "substitute boundary-history response into kappa_lambda", Passed: near(a.BoundarySub.W72, math.Abs(lambdaLambda12)+a.Gate700.P_K7*a.Gate700.SSplit, 1e-18) && near(a.BoundarySub.KappaLambdaApprox+a.BoundarySub.EWall, kappaLambda, 1e-12) && a.BoundarySub.ApproxDropsWallResidual && a.BoundarySub.BoundaryMinusFlavorReading && strings.Contains(a.BoundarySub.Verdict, StatusBoundaryHistoryResponseSubstituted), Detail: FormatBoundarySubstitution(a.BoundarySub)},
			{Name: "substitute Radial-Hopf L into scalar transport", Passed: near(a.RadialSub.L, a.Gate727.L, 1e-18) && near(a.RadialSub.RadialHopfExpectation, a.Gate727.RadialHopfExpectation, 1e-18) && a.RadialSub.UsesRadialHopfExpectation && near(a.RadialSub.FactorApprox, 1-a.BoundarySub.W72+kappaE, 1e-18) && strings.Contains(a.RadialSub.Verdict, StatusRadialHopfLSubstituted), Detail: FormatRadialSubstitution(a.RadialSub)},
			{Name: "assemble dual event-expectation runtime form", Passed: a.Assembly.DualEventExpectationForm && near(a.Assembly.W72, a.BoundarySub.W72, 1e-18) && near(a.Assembly.L, a.Gate727.L, 1e-18) && near(a.Assembly.AssembledRuntimeApprox, a.RadialSub.RuntimePredApprox, 1e-18) && strings.Contains(a.Assembly.Verdict, StatusDualEventExpectationFormAssembled), Detail: FormatAssembly(a.Assembly)},
			{Name: "compute wall residual propagation", Passed: near(a.Propagation.DeltaLambdaPred, lambdaProxyMZ*a.Gate727.L*a.Gate700.EWall, 1e-18) && a.Propagation.MatchesPropagationFormula && a.Propagation.RuntimeResidualIsWallResidual && strings.Contains(a.Propagation.Verdict, StatusWallResidualPropagationComputed), Detail: FormatPropagation(a.Propagation)},
			{Name: "audit noncircularity", Passed: a.NonCircular.KappaLambdaDefinedFromRuntime && !a.NonCircular.AssembledIndependentPrediction && a.NonCircular.BridgeConsistencyClosure && strings.Contains(a.NonCircular.Verdict, StatusAssembledRuntimeNotIndependentPrediction), Detail: FormatNoncircularity(a.NonCircular)},
			{Name: "audit seal dependence", Passed: a.Seals.DependsOnN && a.Seals.DependsOnPRad && a.Seals.DependsOnRhoPlus && a.Seals.DependsOnRho72 && a.Seals.DependsOnPK7 && a.Seals.DependsOnKappaE && !a.Seals.PremisesNativelyDerived && strings.Contains(a.Seals.Verdict, StatusPremisesNotNativelyDerived), Detail: FormatSeals(a.Seals)},
			{Name: "preserve firewalls", Passed: !a.Firewall.ClaimsScalarRuntimeTheorem && !a.Firewall.ClaimsHiggsMassTheorem && !a.Firewall.ClaimsNativeHistoryLoopUnit && !a.Firewall.ClaimsNativeBoundaryHistory && !a.Firewall.ClaimsNativeRadialSelector && !a.Firewall.ClaimsYukawaOperatorTheorem && !a.Firewall.ClaimsIndependentPrediction && strings.Contains(a.Firewall.Verdict, StatusGate728Boundary), Detail: FormatFirewall(a.Firewall)},
		}
		ok := true
		for _, c := range checks {
			if !c.Passed {
				ok = false
				break
			}
		}
		status := theorem.BridgeRequired
		if !ok {
			status = theorem.FailedRoute
		}
		notes := append([]string{a.Truth}, Statuses()...)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
