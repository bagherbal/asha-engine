package generation2higgsradialeventweightandphaselooptransportaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate724RadialEventPhaseLoopTransport(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate723.Inherited || !a.Gate723.QuarterPhaseCandidate || !a.Gate723.PhaseLoopMeasureCandidate || !a.Gate723.ScalarTransportNotRepresentationLayer || !a.Gate723.NoNativeHistoryLoopUnit {
		t.Fatalf("bad Gate723 inheritance: %+v", a.Gate723)
	}
	if a.RhoPlus.CarrierDimension != 4 || math.Abs(a.RhoPlus.Trace-1) > 1e-18 || !a.RhoPlus.MaximallyMixed {
		t.Fatalf("bad rho_plus audit: %+v", a.RhoPlus)
	}
	if a.RadialEvent.Rank != 1 || math.Abs(a.RadialEvent.Weight-0.25) > 1e-18 || !a.RadialEvent.ProjectorIdempotent || !a.RadialEvent.ProjectorSymmetric || !a.RadialEvent.ActsInsideK7Plus || a.RadialEvent.NativeSelector {
		t.Fatalf("bad radial event audit: %+v", a.RadialEvent)
	}
	if math.Abs(a.PhasePayoff.PhaseLoopPayoff-1/(2*math.Pi)) > 1e-18 || math.Abs(a.PhasePayoff.Expectation-1/(8*math.Pi)) > 1e-18 || !a.PhasePayoff.ExpectationMatchesL || a.PhasePayoff.NativePayoffTheorem {
		t.Fatalf("bad phase payoff audit: %+v", a.PhasePayoff)
	}
}

func TestGate724AlternativesAndFirewalls(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.Ranks.Alternatives) != 4 || !oneActiveRank(a.Ranks.Alternatives, 1) {
		t.Fatalf("bad alternative rank audit: %+v", a.Ranks)
	}
	if a.Selectors.P_radNativelySelected || a.Selectors.TwistorSelectorSelectsRadial || a.Selectors.QSourcesRadialOrL {
		t.Fatalf("selector firewall failed: %+v", a.Selectors)
	}
	if math.Abs(a.Analogy.K7EventProbability-float64(k7Dim)/float64(h72Dim)) > 1e-18 || math.Abs(a.Analogy.RadialEventProbability-0.25) > 1e-18 || !a.Analogy.AnalogousEventWeights || a.Analogy.SevenOver72SourcesQuarter || a.Analogy.SevenOver72SourcesL {
		t.Fatalf("bad event-weight analogy: %+v", a.Analogy)
	}
	if !a.Placement.BelongsAfterScalarProxy || a.Placement.DerivedFromRepresentationAlone || !a.Placement.NoScalarProxyRuntimeTheorem {
		t.Fatalf("bad scalar placement: %+v", a.Placement)
	}
	if a.Firewall.NativeRadialProjectorSelector || a.Firewall.NativeHistoryLoopUnitSourceTheorem || a.Firewall.NativePhaseLoopPayoffTheorem || a.Firewall.TwistorSelectorSelectsRadialEvent || a.Firewall.QSourcesL || a.Firewall.SevenOver72SourcesOneOver8Pi || a.Firewall.NativeScalarProxyToRuntimeTheorem || a.Firewall.HiggsMassOrPoleMassTheorem || a.Firewall.YukawaOperatorOrEigenvalueTheorem {
		t.Fatalf("firewall failed: %+v", a.Firewall)
	}
	res := Generation2HiggsRadialEventWeightAndPhaseLoopTransportAuditTheorem().Verify()
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
