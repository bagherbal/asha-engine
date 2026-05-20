package generation2canonicalamplitudeairlockaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate666Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.CoordinateSealInherited || !a.Inherited.AmplitudeNatural || a.Inherited.RGNativeInverseNatural || a.Inherited.CoordinateRobust {
		t.Fatalf("bad inheritance: %+v", a.Inherited)
	}
	if math.Abs(a.Inherited.AmplitudeWBestMinus7Over72) > 1e-6 {
		t.Fatalf("amplitude coordinate should preserve 7/72: %+v", a.Inherited)
	}
	if math.Abs(a.Inherited.InverseWBestMinus7Over72) < 0.01 {
		t.Fatalf("inverse coordinate should not preserve 7/72: %+v", a.Inherited)
	}
	if len(a.CoordinateStack.Rows) != 5 || !a.CoordinateStack.AmplitudeLayerPasses || a.CoordinateStack.InverseKineticLayerPasses || a.CoordinateStack.StrengthLayerPasses || a.CoordinateStack.LogLayerPasses {
		t.Fatalf("bad coordinate stack: %+v", a.CoordinateStack)
	}
	if a.KineticToAmplitude.AmplitudeResidual < 0.04 || a.KineticToAmplitude.AmplitudeResidual > 0.06 {
		t.Fatalf("bad amplitude wound: %+v", a.KineticToAmplitude)
	}
	if a.KineticToAmplitude.InverseOverAmplitude < 1.8 || a.KineticToAmplitude.InverseOverAmplitude > 2.1 {
		t.Fatalf("inverse wound should approximately double the amplitude wound: %+v", a.KineticToAmplitude)
	}
	if a.KineticToAmplitude.AmplitudeScalarScaleGap >= a.KineticToAmplitude.InverseScalarScaleGap {
		t.Fatalf("amplitude wound should be closer to scalar wound than inverse wound: %+v", a.KineticToAmplitude)
	}
	if len(a.Pattern.Rows) != 5 || !strings.Contains(a.Pattern.Verdict, StatusRootAmplitudePatternSupported) {
		t.Fatalf("bad recurring pattern: %+v", a.Pattern)
	}
	if !strings.Contains(a.Target.CandidateTheorem, "CanonicalAmplitudeAirlockTheorem") || !strings.Contains(a.Target.Verdict, StatusNoNativeAmplitudeAirlockTheorem) {
		t.Fatalf("bad theorem target: %+v", a.Target)
	}
	if a.Source.Classification != "BoundaryWeightedDeficitClosureAmplitudeSeal" || len(a.Source.Statements) != 4 {
		t.Fatalf("bad source verdict: %+v", a.Source)
	}
	if a.Discipline.ClaimsNativeAmplitudeAirlockTheorem || a.Discipline.ClaimsNativeDualRootTheorem || a.Discipline.ClaimsNativeSevenOver72Theorem || a.Discipline.ClaimsNativeTransportTheorem || a.Discipline.ClaimsBoundaryStressDerivation || a.Discipline.ClaimsHiggsPrediction || a.Discipline.ClaimsScalarStability || a.Discipline.ClaimsFlavorDerivation || a.Discipline.ClaimsGaugeUnification || a.Discipline.ClaimsCKMPMNSDerivation || a.Discipline.Verdict != StatusGate666Boundary {
		t.Fatalf("firewall breach: %+v", a.Discipline)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2CanonicalAmplitudeAirlockForBoundaryWeightedDeficitClosureAuditTheorem().Verify()
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
