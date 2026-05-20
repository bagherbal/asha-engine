package generation2k7plusmaximumentropyobserverstateandradialeventweightaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate736MaximumEntropyAndNoBias(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate735.Inherited || !a.Gate735.RhoPlusInventoried || !a.Gate735.PRadStillSealed || !a.Gate735.ForecastOnlyBridgeConsistency {
		t.Fatalf("bad Gate735 inheritance: %+v", a.Gate735)
	}
	if a.Entropy.Dimension != 4 || !a.Entropy.PositiveNormalized || !a.Entropy.UniqueMaximumEntropy || !near(a.Entropy.Trace, 1, 1e-15) || !near(a.Entropy.Entropy, math.Log(4), 1e-15) {
		t.Fatalf("bad entropy audit: %+v", a.Entropy)
	}
	if !a.NoBias.SelectsRhoPlus || !near(a.NoBias.Coefficient, 0.25, 1e-15) {
		t.Fatalf("bad no-bias audit: %+v", a.NoBias)
	}
}

func TestGate736WeightsHistoryLoopAndFirewalls(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Radial.Rank != 1 || !near(a.Radial.Weight, 0.25, 1e-15) || !a.Radial.IndependentOfLine || a.Radial.RhoPlusSelectsEvent {
		t.Fatalf("bad radial event audit: %+v", a.Radial)
	}
	if !near(a.Weights.RadialWeight, 0.25, 1e-15) || !near(a.Weights.PhaseWeight, 0.25, 1e-15) || !near(a.Weights.TransverseWeight, 0.5, 1e-15) || !a.Weights.RequiresN || !a.Weights.RequiresPRad {
		t.Fatalf("bad radial/phase/transverse weights: %+v", a.Weights)
	}
	if a.Biased.RhoPlusUniqueAmongAllDensityStates || !a.Biased.BiasedStateCanReproduceWeight || !a.Biased.BiasedReproductionCircular {
		t.Fatalf("biased firewall failed: %+v", a.Biased)
	}
	if a.Selectors.RhoPlusSelectsPRad || a.Selectors.RhoPlusSelectsN || a.Selectors.RhoPlusSelectsPhaseLine || a.Selectors.RhoPlusSelectsHopfFiber || !a.Selectors.UsefulAfterNAndPRad {
		t.Fatalf("selector firewall failed: %+v", a.Selectors)
	}
	if !a.HistoryLoop.MatchesHistoryLoopUnit || !near(a.HistoryLoop.Expectation, 1/(8*math.Pi), 1e-18) || !a.HistoryLoop.UsesMaxEntropyWeight || a.HistoryLoop.NativeTransportTheorem {
		t.Fatalf("bad HistoryLoop placement: %+v", a.HistoryLoop)
	}
	if a.Firewall.RhoPlusSelectsPRad || a.Firewall.RhoPlusSelectsN || a.Firewall.NativeHistoryLoopUnitSource || a.Firewall.NativeScalarRuntimeTheorem || a.Firewall.HiggsMassOrPoleMassTheorem || a.Firewall.YukawaOperatorOrEigenvalueTheorem {
		t.Fatalf("physical firewall failed: %+v", a.Firewall)
	}

	res := Generation2K7PlusMaximumEntropyObserverStateAndRadialEventWeightAuditTheorem().Verify()
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
