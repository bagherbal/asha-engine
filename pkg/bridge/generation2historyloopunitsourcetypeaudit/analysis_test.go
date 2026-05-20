package generation2historyloopunitsourcetypeaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate624Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	L := 1 / (8 * math.Pi)
	if a.Inherited.Verdict != StatusGate623Inherited {
		t.Fatalf("bad inheritance verdict: %s", a.Inherited.Verdict)
	}
	if math.Abs(a.Inherited.LoopUnit-L) > 1e-18 {
		t.Fatalf("bad loop unit %.18g", a.Inherited.LoopUnit)
	}
	if !a.Decompositions.AllValuesMatch || !a.Decompositions.AllRowsTyped || len(a.Decompositions.Rows) != 5 {
		t.Fatalf("bad decompositions: %+v", a.Decompositions)
	}
	if !a.HopfPhase.CirclePhaseNormalization || !a.HopfPhase.QuarterProjectionCandidate || a.HopfPhase.QuarterProjectionCertified {
		t.Fatalf("bad Hopf phase audit: %+v", a.HopfPhase)
	}
	if a.HopfPhase.MapToFlavorWallCertified || a.HopfPhase.MapToScalarMatchingCertified || a.HopfPhase.PhysicalTimeClaimed {
		t.Fatalf("Hopf candidate overpromoted: %+v", a.HopfPhase)
	}
	if !a.WeakQuarter.WeakNormalizationTyped || !a.WeakQuarter.PMNSOverlapTyped || a.WeakQuarter.NativeConnectionToL {
		t.Fatalf("bad weak quarter audit: %+v", a.WeakQuarter)
	}
	if a.HeatKernel.AnyCertifiedReduction {
		t.Fatalf("heat-kernel reduction must remain uncertified")
	}
	if math.Abs(a.ScalarRole.KappaLambda-0.0443230430960771) > 1e-14 {
		t.Fatalf("bad kappa_lambda %.18g", a.ScalarRole.KappaLambda)
	}
	if a.ScalarRole.KappaSourceCertified {
		t.Fatalf("kappa_lambda source must not be certified")
	}
	if math.Abs(a.FlavorRole.Residual) > 2e-7 || a.FlavorRole.NativeDerived {
		t.Fatalf("bad flavor role audit: %+v", a.FlavorRole)
	}
	if len(a.CrossSeal.Rows) != 3 || !a.CrossSeal.SharedLBridgeSeal || a.CrossSeal.NativeCrossSeal {
		t.Fatalf("bad cross-seal table: %+v", a.CrossSeal)
	}
	if a.NativeStatus.NativeLTheorem || a.NativeStatus.NativeHeatKernelToLReduction || a.NativeStatus.NativeCrossSealOrientationLaw {
		t.Fatalf("native theorem incorrectly certified: %+v", a.NativeStatus)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2HistoryLoopUnitSourceTypeAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected failed theorem: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusDecompositionsTyped, StatusHopfPhaseAudited, StatusWeakQuarterAudited, StatusHeatKernelAudited, StatusQuarterPhaseCandidate, StatusNoHopfToFlavorTheorem, StatusNoHeatKernelReduction, StatusGate624Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
