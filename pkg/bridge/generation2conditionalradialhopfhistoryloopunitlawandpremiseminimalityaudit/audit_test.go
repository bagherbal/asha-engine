package generation2conditionalradialhopfhistoryloopunitlawandpremiseminimalityaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate727ConditionalRadialHopfLaw(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate726.Inherited || !a.Gate726.RadialPhaseTransverseDefined || !a.Gate726.HopfFiberPhaseLoopAudited || !a.Gate726.EventWeightsComputed || !a.Gate726.SelectorDependenceAudited {
		t.Fatalf("bad Gate726 inheritance: %+v", a.Gate726)
	}
	if a.Observable.RadialRank != 1 || a.Observable.CarrierDimension != 4 || math.Abs(a.Observable.PhaseLoopUnit-1/(2*math.Pi)) > 1e-18 || !a.Observable.UsesRadialEvent || !a.Observable.UsesHopfPhaseUnit {
		t.Fatalf("bad payoff observable: %+v", a.Observable)
	}
	if math.Abs(a.Functional.L-1/(8*math.Pi)) > 1e-18 || math.Abs(a.Functional.Expectation-1/(8*math.Pi)) > 1e-18 || math.Abs(a.Functional.Residual) > 1e-18 || !a.Functional.ConditionallyExact {
		t.Fatalf("bad conditional functional: %+v", a.Functional)
	}
	if a.Premises.Count != 5 || !a.Premises.FourRealCarrier || !a.Premises.RhoPlusNoBiasState || !a.Premises.RankOneRadialEvent || !a.Premises.TwistorJHPhaseLoop || !a.Premises.FirstExpectationPayoff {
		t.Fatalf("bad premise ladder: %+v", a.Premises)
	}
}

func TestGate727MinimalityAndFirewalls(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Removal.WithoutRhoPlusWeightFixed || a.Removal.WithoutPRadRadialEventDefined || a.Removal.WithoutNPhaseLoopDefined || a.Removal.WithoutPhaseUnitGivesL || math.Abs(a.Removal.RankTwoValue-1/(4*math.Pi)) > 1e-18 || math.Abs(a.Removal.FullEventValue-1/(2*math.Pi)) > 1e-18 || a.Removal.RankTwoMatchesL || a.Removal.FullEventMatchesL || a.Removal.QuadraticMomentIsActive || !a.Removal.EachPremiseDoesWork {
		t.Fatalf("bad premise removal audit: %+v", a.Removal)
	}
	if !a.NonTaut.ConditionallyExact || a.NonTaut.PremisesNativelyDerived || a.NonTaut.NativeRadialProjectorSelector || a.NonTaut.NativeTwistorSelectorN || a.NonTaut.HistoryTransportUsesHopfProved || a.NonTaut.RhoPlusPhysicalHistoryTheorem {
		t.Fatalf("non-tautology firewall failed: %+v", a.NonTaut)
	}
	if !a.Transport.AfterScalarProxyLane || a.Transport.NativeRuntimeTransportTheorem || a.Transport.NativeScalarProxyToRuntimeTheorem {
		t.Fatalf("transport placement failed: %+v", a.Transport)
	}
	if !a.Analogy.NeitherDerivesOther || math.Abs(a.Analogy.K7EventWeight-7.0/72.0) > 1e-18 || math.Abs(a.Analogy.RadialHopfValue-1/(8*math.Pi)) > 1e-18 {
		t.Fatalf("analogy audit failed: %+v", a.Analogy)
	}
	if a.Firewall.NativeRadialProjectorSelector || a.Firewall.NativeTwistorSelectorN || a.Firewall.HistoryTransportUsesHopfPhasePayoff || a.Firewall.NativeHistoryLoopUnitSourceTheorem || a.Firewall.NativeScalarProxyToRuntimeTheorem || a.Firewall.HiggsMassOrPoleMassTheorem || a.Firewall.YukawaOperatorOrEigenvalueTheorem {
		t.Fatalf("firewall failed: %+v", a.Firewall)
	}

	res := Generation2ConditionalRadialHopfHistoryLoopUnitLawAndPremiseMinimalityAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected construction failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
