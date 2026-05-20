package generation2scalarquarticcoordinateairlockaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate668Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.ConnectionAmplitudeInherited || !a.Inherited.AmplitudeOnlyPasses || !a.Inherited.InverseKineticFails || !a.Inherited.ScalarSideWasRuntimeShadow {
		t.Fatalf("bad inheritance: %+v", a.Inherited)
	}
	if len(a.Scalars.Rows) != 6 || a.Scalars.ActiveScalarCoordinate != "|lambda(Lambda_12)|" || a.Scalars.HessianCoordinate != "2|lambda(Lambda_12)|" {
		t.Fatalf("bad scalar coordinate family: %+v", a.Scalars)
	}
	if !a.Hessian.TypedAsHessianLayer || math.Abs(a.Hessian.HessianCoordinate-2*absLambda12) > 1e-15 || math.Abs(a.Hessian.InverseKineticWound-0.0946) > 1e-3 {
		t.Fatalf("bad hessian audit: %+v", a.Hessian)
	}
	if len(a.Pairings.Rows) != 5 || !a.Pairings.AmplitudePairPasses || !a.Pairings.InverseHessianShadowMagnitude || a.Pairings.InverseHessianClosurePasses || a.Pairings.MassAmplitudePairPasses {
		t.Fatalf("bad pairings: %+v", a.Pairings)
	}
	if a.Retest.BestTypedPair != "amplitude/quartic" || math.Abs(a.Retest.BestTypedWBestMinus7) > 1e-6 {
		t.Fatalf("bad retest: %+v", a.Retest)
	}
	if a.Source.Classification != "BoundaryWeightedDeficitClosureQuarticWoundSeal" || len(a.Source.Statements) != 4 || !strings.Contains(a.Source.Verdict, StatusNoNativeScalarAirlockTheorem) {
		t.Fatalf("bad source: %+v", a.Source)
	}
	if a.Discipline.ClaimsNativeScalarAirlockTheorem || a.Discipline.ClaimsNativeBoundaryStressTheorem || a.Discipline.ClaimsNativeSevenOver72Theorem || a.Discipline.ClaimsNativeTransportTheorem || a.Discipline.ClaimsHiggsMassPrediction || a.Discipline.ClaimsScalarStability || a.Discipline.ClaimsGaugeUnification || a.Discipline.ClaimsFlavorDerivation || a.Discipline.ClaimsCKMPMNSDerivation || a.Discipline.Verdict != StatusGate668Boundary {
		t.Fatalf("firewall breach: %+v", a.Discipline)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2ScalarQuarticCoordinateAirlockAndHessianDoublingAuditTheorem().Verify()
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

func TestFiniteGuard(t *testing.T) {
	if !finiteNumber(sevenOver72) || finiteNumber(math.Inf(1)) {
		t.Fatal("finite guard broken")
	}
}
