package ccmpfaffianf0closure

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func SelfConsistentCCMPfaffianCoefficientClosureF0SieveTheorem() theorem.Theorem {
	const id = "BRIDGE-CCM-PFAFFIAN-F0-CLOSURE"
	const name = "Self-Consistent CCM + Pfaffian Coefficient Closure & f0 Sieve"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build CCM+Pfaffian f0 closure ledger", Passed: false, Detail: err.Error()}}}
		}
		c := a.Calculation
		checks := []theorem.Check{
			{Name: "CCM quartic and Pfaffian VEV frameworks are combined", Passed: c.Executed && strings.Contains(c.Input.QuarticFormula, "π²") && c.Input.PfaffianRatio > 0 && c.Input.PfaffianVEVGeV > 240 && c.Input.PfaffianVEVGeV < 250, Detail: c.Input.QuarticFormula + "; " + c.Input.F0TargetFormula},
			{Name: "effective f0 target is extracted from Higgs boundary", Passed: math.Abs(c.F0Targets.UsingStandardEWVEV-effectiveF0(StandardEWVEVGeV, HiggsMassBoundaryGeV)) < 1e-12 && math.Abs(c.F0Targets.UsingPfaffianVEV-effectiveF0(c.Input.PfaffianVEVGeV, HiggsMassBoundaryGeV)) < 1e-12, Detail: c.F0Targets.Verdict},
			{Name: "f0=7 and f0=14 do not close the Higgs mass", Passed: c.StandardVEVPredictions[0].PredictedMass > 140 && c.StandardVEVPredictions[2].PredictedMass < 110 && c.PfaffianVEVPredictions[0].PredictedMass > 140 && c.PfaffianVEVPredictions[2].PredictedMass < 110, Detail: formatPredictions(c.StandardVEVPredictions) + " | " + formatPredictions(c.PfaffianVEVPredictions)},
			{Name: "f0=10 predicts a near-125 GeV Higgs mass", Passed: math.Abs(c.StandardVEVPredictions[1].PredictedMass-HiggsMassBoundaryGeV) < 1.0 && math.Abs(c.PfaffianVEVPredictions[1].PredictedMass-HiggsMassBoundaryGeV) < 0.3, Detail: formatPredictions([]F0Point{c.StandardVEVPredictions[1], c.PfaffianVEVPredictions[1]})},
			{Name: "Rule-of-10 edge sieve finds five finite edge classes and ten J-doubled edge slots", Passed: c.EdgeSieve.FundamentalEdgeCount == 5 && c.EdgeSieve.JDoubledEdgeCount == 10 && c.EdgeSieve.MatchesF0Target, Detail: c.EdgeSieve.Verdict},
			{Name: "f0=10 is not promoted to a native theorem without a moment functional proof", Passed: !c.EdgeSieve.IsSpectralMomentProof && !c.NativeHiggsMassClosed && strings.Contains(StatusLine(c), StatusFailedF0MomentNotDerived), Detail: c.Truth},
			{Name: "full numerical ToE closure is not claimed", Passed: !c.FullNumericalTOEClosure && strings.Contains(StatusLine(c), StatusFailedFullNumericalTOEClosure), Detail: c.Truth},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{c.Truth}}
	}}
}

func formatPredictions(points []F0Point) string {
	out := make([]string, 0, len(points))
	for _, p := range points {
		out = append(out, p.Label+": f0="+FormatFloat(p.F0)+", lambda="+FormatFloat(p.Lambda)+", mH="+FormatFloat(p.PredictedMass)+" GeV, err%="+FormatFloat(p.PercentError))
	}
	return strings.Join(out, "; ")
}
