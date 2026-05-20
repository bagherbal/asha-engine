package generation2boundaryrawmomentgeneratingfunctionandcubicstopaudit

import (
	"strings"
	"testing"
)

func TestGate783LedgerAndOperatorBoard(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate782.Inherited || !strings.Contains(a.Gate782.SelectedSubBottleneck, "F_wall_3_red") {
		t.Fatalf("bad Gate782 inheritance: %+v", a.Gate782)
	}
	if !closeRel(a.Ledger.M1, 0.0001256543573849177, 1e-14) || !closeRel(a.Ledger.M2, 1.624013231638281e-07, 1e-14) || !closeRel(a.Ledger.M3, 2.0989474869200057e-10, 1e-14) {
		t.Fatalf("bad raw moments: %+v", a.Ledger)
	}
	if !closeRel(a.Ledger.FWall3, 0.00012565521035653708, 1e-14) || !a.Ledger.Matches {
		t.Fatalf("bad F_wall ledger: %+v", a.Ledger)
	}
	if !a.Operator.Audited || !a.Operator.ProjectorIdempotent || !strings.Contains(a.Operator.PowerLaw, "s^n P_7") || !a.Operator.AllMomentsSameSupport || a.Operator.IndependentOperatorGeometry || !a.Operator.ScalarRawMomentCoordinate {
		t.Fatalf("bad operator audit: %+v", a.Operator)
	}
}

func TestGate783ResponseCoefficientsAndCubicStop(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Response.Audited || !a.Response.WrittenAsExpectation || !strings.Contains(a.Response.Function, "f_3") || !a.Response.MatchesLedger || a.Response.NativeGeneratingFunction {
		t.Fatalf("bad response audit: %+v", a.Response)
	}
	if !a.Coefficients.Audited || !closeRel(a.Coefficients.TwoP, 7.0/36.0, 1e-15) || !strings.Contains(a.Coefficients.BoundaryPairCandidate, "2*(7/72)") || a.Coefficients.NativeBoundaryPairTheorem || a.Coefficients.TypedOperatorMapSourcesCubic {
		t.Fatalf("bad coefficient audit: %+v", a.Coefficients)
	}
	if !a.CubicStop.Audited || a.CubicStop.ProjectorAlgebraStopsPowers || !a.CubicStop.BoundaryPairExteriorCandidate || a.CubicStop.RawMomentToExteriorMapCertified || a.CubicStop.NativeCubicStopTheorem || !a.CubicStop.UntypedM4FitRejected || a.CubicStop.TypedFourthOrderCoefficient {
		t.Fatalf("bad cubic stop audit: %+v", a.CubicStop)
	}
}

func TestGate783CoordinateRuntimeRelationAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Coordinate.Audited || !strings.Contains(a.Coordinate.ActiveCoordinate, "raw moments") || a.Coordinate.VarianceActive || a.Coordinate.CentralMomentActive || a.Coordinate.NativeRawMomentCoordinate {
		t.Fatalf("bad coordinate audit: %+v", a.Coordinate)
	}
	if !a.Runtime.Audited || !a.Runtime.FormulaLevelRuntimeTargetAbsence || !a.Runtime.EvaluableWithoutDirectHiggsRuntimeVariables || a.Runtime.UsesLambdaRuntime || a.Runtime.UsesLambdaRuntimeEff || a.Runtime.UsesTreeMass || a.Runtime.UsesPoleMass || a.Runtime.UsesCHiggs || a.Runtime.UsesGF || a.Runtime.UsesVEV || a.Runtime.NativeBoundaryResponseTheorem || a.Runtime.NativeCubicCoefficientTheorem || a.Runtime.NativeCubicStopTheorem || !a.Runtime.KappaERedFlavorSealed {
		t.Fatalf("bad runtime audit: %+v", a.Runtime)
	}
	if !a.Relation.Recorded || !strings.Contains(a.Relation.BoundarySubBottleneck, "F_wall_3_red") || !strings.Contains(a.Relation.FWallLevel, "Level B") || !strings.Contains(a.Relation.CHiggsLevel, "not Level C") || a.Relation.CHistoryIndependent {
		t.Fatalf("bad relation audit: %+v", a.Relation)
	}
	if !a.Prediction.Recorded || !strings.Contains(a.Prediction.FWall3Level, "Level B") || !strings.Contains(a.Prediction.CHiggsLevel, "not Level C") {
		t.Fatalf("bad prediction classification: %+v", a.Prediction)
	}
	if !a.Firewalls.Enforced || a.Firewalls.ResponseFunctionNative || a.Firewalls.BoundaryGeneratingFunctionNative || a.Firewalls.RawMomentCoordinateNative || a.Firewalls.BoundaryPairStressPullNative || a.Firewalls.RawMomentExteriorDegreeMapNative || a.Firewalls.CubicStopNative || a.Firewalls.TypedFourthOrderCoefficientSource || a.Firewalls.FWallNativeBoundaryResponse || a.Firewalls.CHistoryFullIndependentPrediction || a.Firewalls.TreeProxyPoleMass || a.Firewalls.YukawaNativeTheorem {
		t.Fatalf("bad firewalls: %+v", a.Firewalls)
	}
	if !strings.Contains(a.FinalStatement, "does not make F_wall_3_red native") || !strings.Contains(a.FinalStatement, "cubic raw-response expectation") || !strings.Contains(a.FinalStatement, "next bottleneck") {
		t.Fatalf("bad final statement: %s", a.FinalStatement)
	}
}

func TestGate783TheoremStatuses(t *testing.T) {
	res := Generation2BoundaryRawMomentGeneratingFunctionAndCubicStopAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status note %s", want)
		}
	}
}
