package generation2boundarystresssplitpullbackcorrectionaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate672Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.NormalVectorInherited || !a.Inherited.NormalVectorBestTypedExact || !a.Inherited.CoordinateSealed || !a.Inherited.NoNativeSevenOver72Theorem || !a.Inherited.FirewallPreserved {
		t.Fatalf("bad inheritance: %+v", a.Inherited)
	}
	if !a.Decomposition.DecompositionPasses || len(a.Decomposition.OriginalNormal) != 4 || math.Abs(a.Decomposition.Weight-sevenOver72) > 1e-15 {
		t.Fatalf("bad decomposition: %+v", a.Decomposition)
	}
	if math.Abs(a.BaseClosure.DBase-0.00012565520996836) > 1e-14 {
		t.Fatalf("bad base closure: %+v", a.BaseClosure)
	}
	if math.Abs(a.StressSplit.SSplit-0.0012924448188163) > 1e-14 {
		t.Fatalf("bad stress split: %+v", a.StressSplit)
	}
	if !a.Pullback.PassesBridgeWindow || math.Abs(a.Pullback.Pullback-0.000125654357384641) > 1e-14 || math.Abs(a.Pullback.Residual-8.52583727234e-10) > 1e-14 {
		t.Fatalf("bad pullback: %+v", a.Pullback)
	}
	if !a.Reconstruction.EquivalentToGate670Normal || math.Abs(a.Reconstruction.DBaseMinusPullback-a.Reconstruction.HistoryWallBalance) > 1e-15 {
		t.Fatalf("bad reconstruction: %+v", a.Reconstruction)
	}
	if len(a.Source.CandidateSupport) != 3 || len(a.Source.RequiredMissingMaps) != 4 || !strings.Contains(a.Source.Verdict, StatusStressSplitCorrectedScalarFlavorClosure) {
		t.Fatalf("bad source audit: %+v", a.Source)
	}
	if a.Discipline.ClaimsNativeStressSplitPullback || a.Discipline.ClaimsNativeSevenOver72 || a.Discipline.ClaimsWallDistanceAirlock || a.Discipline.ClaimsBoundaryStressDerivation || a.Discipline.ClaimsHiggsMassPrediction || a.Discipline.ClaimsScalarStability || a.Discipline.ClaimsGaugeUnification || a.Discipline.ClaimsFlavorDerivation || a.Discipline.ClaimsCKMPMNSDerivation || a.Discipline.Verdict != StatusGate672Boundary {
		t.Fatalf("firewall breach: %+v", a.Discipline)
	}
}

func TestNormalVectorDecompositionIdentity(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	for i := range a.Decomposition.OriginalNormal {
		got := a.Decomposition.BaseNormal[i] - sevenOver72*a.Decomposition.StressSplitNormal[i]
		if math.Abs(got-a.Decomposition.OriginalNormal[i]) > 1e-15 {
			t.Fatalf("component %d mismatch: got %.17g want %.17g", i, got, a.Decomposition.OriginalNormal[i])
		}
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2BoundaryStressSplitPullbackCorrectionAuditTheorem().Verify()
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
