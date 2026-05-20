package generation2conditionalradialhopfhistoryloopunitlawandpremiseminimalityaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2ConditionalRadialHopfHistoryLoopUnitLawAndPremiseMinimalityAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 727 — Conditional Radial-Hopf HistoryLoopUnit Law and Premise-Minimality Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate727 conditional radial-Hopf audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate726 radial phase Hopf decomposition", Passed: a.Gate726.Inherited && a.Gate726.RadialPhaseTransverseDefined && a.Gate726.HopfFiberPhaseLoopAudited && a.Gate726.EventWeightsComputed && a.Gate726.SelectorDependenceAudited && a.Gate726.NoNativeRadialProjectorSelector && a.Gate726.NoNativeTwistorSelectorN && a.Gate726.NoNativeHistoryLoopUnit && a.Gate726.NoHiggsMassTheorem && a.Gate726.NoYukawaTheorem && a.Gate726.Verdict == StatusGate726RadialPhaseHopfInherited, Detail: FormatGate726(a.Gate726)},
			{Name: "define radial-Hopf payoff observable", Passed: a.Observable.RadialRank == radialRank && a.Observable.CarrierDimension == k7PlusRealDim && near(a.Observable.PhaseLoopUnit, phaseLoopUnit, 1e-18) && a.Observable.UsesRadialEvent && a.Observable.UsesHopfPhaseUnit && strings.Contains(a.Observable.Verdict, StatusRadialHopfPayoffObservableDefined), Detail: FormatObservable(a.Observable)},
			{Name: "define conditional HistoryLoop functional and reproduce 1/(8*pi)", Passed: near(a.Functional.L, historyLoopUnit, 1e-18) && near(a.Functional.Expectation, historyLoopUnit, 1e-18) && near(a.Functional.Residual, 0, 1e-18) && a.Functional.ConditionallyExact && strings.Contains(a.Functional.Verdict, StatusExpectationReproducesOneOver8Pi), Detail: FormatFunctional(a.Functional)},
			{Name: "construct premise ladder", Passed: a.Premises.Count == 5 && a.Premises.FourRealCarrier && a.Premises.RhoPlusNoBiasState && a.Premises.RankOneRadialEvent && a.Premises.TwistorJHPhaseLoop && a.Premises.FirstExpectationPayoff && strings.Contains(a.Premises.Verdict, StatusPremiseLadderConstructed), Detail: FormatPremises(a.Premises)},
			{Name: "compute premise-removal audit", Passed: !a.Removal.WithoutRhoPlusWeightFixed && !a.Removal.WithoutPRadRadialEventDefined && !a.Removal.WithoutNPhaseLoopDefined && !a.Removal.WithoutPhaseUnitGivesL && near(a.Removal.RankTwoValue, 1/(4*math.Pi), 1e-18) && near(a.Removal.FullEventValue, 1/(2*math.Pi), 1e-18) && !a.Removal.RankTwoMatchesL && !a.Removal.FullEventMatchesL && !a.Removal.QuadraticMomentIsActive && a.Removal.EachPremiseDoesWork, Detail: FormatRemoval(a.Removal)},
			{Name: "audit non-tautology and missing native premises", Passed: a.NonTaut.ConditionallyExact && !a.NonTaut.PremisesNativelyDerived && !a.NonTaut.NativeRadialProjectorSelector && !a.NonTaut.NativeTwistorSelectorN && !a.NonTaut.HistoryTransportUsesHopfProved && !a.NonTaut.RhoPlusPhysicalHistoryTheorem && strings.Contains(a.NonTaut.Verdict, StatusPremisesNotNativelyDerived), Detail: FormatNonTautology(a.NonTaut)},
			{Name: "preserve scalar transport placement", Passed: a.Transport.AfterScalarProxyLane && !a.Transport.NativeRuntimeTransportTheorem && !a.Transport.NativeScalarProxyToRuntimeTheorem && strings.Contains(a.Transport.Verdict, StatusScalarTransportPlacementPreserved), Detail: FormatTransport(a.Transport)},
			{Name: "audit event-weight analogy to 7/72", Passed: a.Analogy.NeitherDerivesOther && near(a.Analogy.K7EventWeight, float64(k7EventNumerator)/float64(h72Dimension), 1e-18) && near(a.Analogy.RadialHopfValue, historyLoopUnit, 1e-18) && strings.Contains(a.Analogy.Verdict, StatusEventWeightAnalogyTo7Over72Audited), Detail: FormatAnalogy(a.Analogy)},
			{Name: "preserve firewalls", Passed: !a.Firewall.NativeRadialProjectorSelector && !a.Firewall.NativeTwistorSelectorN && !a.Firewall.HistoryTransportUsesHopfPhasePayoff && !a.Firewall.NativeHistoryLoopUnitSourceTheorem && !a.Firewall.NativeScalarProxyToRuntimeTheorem && !a.Firewall.HiggsMassOrPoleMassTheorem && !a.Firewall.YukawaOperatorOrEigenvalueTheorem && strings.Contains(a.Firewall.Verdict, StatusGate727Boundary), Detail: FormatFirewall(a.Firewall)},
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
