package generation2kinetictoconnectionamplitudeairlockaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate667Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.AmplitudeSealInherited || !a.Inherited.AmplitudeLayerPasses || a.Inherited.InverseKineticLayerPasses {
		t.Fatalf("bad inheritance: %+v", a.Inherited)
	}
	if a.Inherited.InverseOverAmplitude < 1.8 || a.Inherited.InverseOverAmplitude > 2.1 {
		t.Fatalf("expected inverse wound to approximately double amplitude wound: %+v", a.Inherited)
	}
	if math.Abs(a.Inherited.AmplitudeWBestMinus7Over72) > 1e-6 || math.Abs(a.Inherited.InverseWBestMinus7Over72) < 0.01 {
		t.Fatalf("coordinate inheritance should preserve amplitude and block inverse: %+v", a.Inherited)
	}
	if !strings.Contains(a.Kinetic.NativeCoordinate, "1/g_i^2") || !strings.Contains(a.Kinetic.Verdict, StatusInverseKineticStillFails) {
		t.Fatalf("bad kinetic coordinate audit: %+v", a.Kinetic)
	}
	if !a.Rescaling.AmplitudeCoordinateTyped || !strings.Contains(a.Rescaling.AlgebraicRelation, "u_i^{-1/2}") || !strings.Contains(a.Rescaling.Verdict, StatusGaugeAmplitudeSourcedByConnection) {
		t.Fatalf("bad rescaling audit: %+v", a.Rescaling)
	}
	if len(a.Coordinates.Rows) != 5 || !a.Coordinates.AmplitudeOnlyPasses || !a.Coordinates.InverseKineticFails {
		t.Fatalf("bad coordinate comparison: %+v", a.Coordinates)
	}
	if !a.HessianSocket.CompatibleWithClosure || len(a.HessianSocket.AmplitudeObjects) != 4 {
		t.Fatalf("bad hessian socket: %+v", a.HessianSocket)
	}
	if a.ScalarSide.NativeAmplitude || !strings.Contains(a.ScalarSide.Verdict, StatusScalarRuntimeShadow) {
		t.Fatalf("bad scalar side audit: %+v", a.ScalarSide)
	}
	if len(a.Pattern.Rows) != 5 || !strings.Contains(a.Pattern.Verdict, StatusRootAmplitudePatternSupported) {
		t.Fatalf("bad recurring pattern: %+v", a.Pattern)
	}
	if !strings.Contains(a.Target.Name, "Kinetic") || !strings.Contains(a.Target.Verdict, StatusNoNativeKineticAmplitudeTheorem) {
		t.Fatalf("bad target: %+v", a.Target)
	}
	if a.Source.Classification != "BoundaryWeightedDeficitClosureConnectionAmplitudeSeal" || len(a.Source.Statements) != 4 {
		t.Fatalf("bad source classification: %+v", a.Source)
	}
	if a.Discipline.ClaimsNativeKineticAmplitudeTheorem || a.Discipline.ClaimsNativeSevenOver72Theorem || a.Discipline.ClaimsNativeDualRootTheorem || a.Discipline.ClaimsNativeTransportTheorem || a.Discipline.ClaimsBoundaryStressDerivation || a.Discipline.ClaimsHiggsPrediction || a.Discipline.ClaimsScalarStability || a.Discipline.ClaimsFlavorDerivation || a.Discipline.ClaimsGaugeUnification || a.Discipline.ClaimsCKMPMNSDerivation || a.Discipline.Verdict != StatusGate667Boundary {
		t.Fatalf("firewall breach: %+v", a.Discipline)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2KineticToConnectionAmplitudeAirlockSourceAuditTheorem().Verify()
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
