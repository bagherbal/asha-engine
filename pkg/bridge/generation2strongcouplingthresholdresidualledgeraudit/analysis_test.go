package generation2strongcouplingthresholdresidualledgeraudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate607ResidualConversions(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.GaugeSpinePresent || a.Inherited.R3 <= 1 || a.Inherited.Delta3Runtime >= 0 {
		t.Fatalf("bad inheritance: %+v", a.Inherited)
	}
	if !containsResidual(a.ResidualConversions, "Delta g3") || !containsResidual(a.ResidualConversions, "required positive inverse correction") || !containsResidual(a.ResidualConversions, "required Delta alpha3^-1") {
		t.Fatalf("missing residual rows: %s", FormatResidualTable(a.ResidualConversions))
	}
	req := a.ThresholdSlots[0].RequiredValue
	if math.Abs(req-0.32739043299998416) > 1e-12 {
		t.Fatalf("required inverse correction drifted: %.15g", req)
	}
}

func TestGate607BetaAndMeetingTriangle(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.BetaDeformation.DeltaB3Required >= 0 {
		t.Fatalf("expected negative delta b3: %+v", a.BetaDeformation)
	}
	if math.Abs(a.BetaDeformation.FractionOfAbsSMb3-0.133337235907374) > 1e-9 {
		t.Fatalf("unexpected beta deformation fraction: %.15g", a.BetaDeformation.FractionOfAbsSMb3)
	}
	if len(a.MeetingScales) != 3 || a.MeetingScales[0].Pair != "Lambda_12" || a.MeetingScales[1].ScaleGeV <= a.MeetingScales[0].ScaleGeV || a.MeetingScales[2].ScaleGeV <= a.MeetingScales[1].ScaleGeV {
		t.Fatalf("bad meeting triangle: %s", FormatMeetingScales(a.MeetingScales))
	}
}

func TestGate607FirewallsAndStatuses(t *testing.T) {
	res := Generation2StrongCouplingThresholdResidualLedgerAuditTheorem().Verify()
	if !res.Passed() {
		t.Fatalf("theorem checks failed: %+v", res)
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusGate606Inherited, StatusStrongResidualConverted, StatusStrongThresholdSlotDefined, StatusRequiredThresholdQuantified, StatusBetaDeformationComputed, StatusNoNativeStrongThresholdTheorem, StatusNoFullGaugeUnificationClaim, StatusGate607Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status %s", want)
		}
	}
}
