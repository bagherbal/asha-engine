package generation2historyloopdeficitclosuretriangleaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate625Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Inherited.Verdict != StatusGate624Inherited || !a.Inherited.Gate624QuarterPhase {
		t.Fatalf("bad Gate624 inheritance: %+v", a.Inherited)
	}
	if a.Inherited.NativeHistoryLoopUnit || a.Inherited.NativeHopfToScalar || a.Inherited.NativeHopfToFlavor {
		t.Fatalf("Gate624 native theorem overpromoted: %+v", a.Inherited)
	}
	if !a.Kappas.BothPositive || !a.Kappas.ScalarDeficitLarger {
		t.Fatalf("bad kappa definitions: %+v", a.Kappas)
	}
	wantSum := 0.0498265972876517
	if math.Abs(a.ClosureTable.KappaSum-wantSum) > 1e-14 {
		t.Fatalf("bad kappa sum %.18g", a.ClosureTable.KappaSum)
	}
	if a.ClosureTable.ClosestTarget != "|lambda(Lambda_12)|" || !a.ClosureTable.ClosesOnAbsLambda {
		t.Fatalf("closure did not select abs lambda: %+v", a.ClosureTable)
	}
	if math.Abs(a.ClosureTable.ClosestResidual-0.000125655209968385) > 1e-14 {
		t.Fatalf("bad closure residual %.18g", a.ClosureTable.ClosestResidual)
	}
	if a.ClosureTable.ClosestRelative > 0.0026 {
		t.Fatalf("closure not sharp enough: %+v", a.ClosureTable)
	}
	if math.Abs(a.ScalarFormula.PredictedKappaLambdaExact-0.0441973878861087) > 1e-14 {
		t.Fatalf("bad predicted kappa_lambda: %+v", a.ScalarFormula)
	}
	if len(a.ScalarPrediction.Rows) != 2 || !a.ScalarPrediction.ImprovesGate623RawLAnsatz {
		t.Fatalf("bad scalar prediction: %+v", a.ScalarPrediction)
	}
	if math.Abs(a.ScalarPrediction.Rows[0].PredictedLambda-0.129653189523764) > 1e-14 {
		t.Fatalf("bad exact closure scalar prediction %.18g", a.ScalarPrediction.Rows[0].PredictedLambda)
	}
	if math.Abs(a.ScalarPrediction.Rows[0].Residual) > 7e-7 {
		t.Fatalf("exact closure scalar prediction residual too large: %+v", a.ScalarPrediction.Rows[0])
	}
	if !a.ResidualScales.ClosureSharperThanRawScalarAnsatz || a.ResidualScales.ScalarImprovementFactor < 100 {
		t.Fatalf("closure did not improve raw scalar ansatz: %+v", a.ResidualScales)
	}
	if !a.SignRole.OpposedRGWoundSign || a.SignRole.NativeTheoremClaimed {
		t.Fatalf("bad sign/role audit: %+v", a.SignRole)
	}
	if a.NativeStatus.NativeKappaClosureTheorem || a.NativeStatus.NativeHistoryLoopDeficitClosureTheorem {
		t.Fatalf("native closure theorem incorrectly certified: %+v", a.NativeStatus)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2HistoryLoopDeficitClosureTriangleAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected failed theorem: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusGate624Inherited, StatusKappasDefined, StatusDeficitClosureComputed, StatusClosureOnAbsLambda12, StatusScalarPredictionComputed, StatusClosureSealDefined, StatusNoNativeKappaClosure, StatusNoNativeScalarRGMatching, StatusNoNativeFlavorOrientation, StatusGate625Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
