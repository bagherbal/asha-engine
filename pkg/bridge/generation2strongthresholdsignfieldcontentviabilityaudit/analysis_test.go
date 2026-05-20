package generation2strongthresholdsignfieldcontentviabilityaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate609RequiredSignAndMagnitude(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(a.Inherited.DeltaB3Required+0.933360651351616) > 1e-12 {
		t.Fatalf("unexpected Delta b3: %.15g", a.Inherited.DeltaB3Required)
	}
	if a.Inherited.Delta3ThresholdRequired <= 0 || a.Inherited.B3EffectiveDiagnostic >= a.Inherited.B3SM {
		t.Fatalf("bad sign classification: %s", FormatInherited(a.Inherited))
	}
	if math.Abs(a.Inherited.RelativeB3Deformation-0.133337235907374) > 1e-12 {
		t.Fatalf("unexpected relative deformation: %.15g", a.Inherited.RelativeB3Deformation)
	}
}

func TestGate609ViabilityRows(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.WrongSignMatter.Verdict != StatusExtraColoredWrongSign {
		t.Fatalf("ordinary matter route not blocked: %s", FormatWrongSignMatter(a.WrongSignMatter))
	}
	if !containsOrigin(a.CorrectionOrigins, "boundary-localized") || !containsOrigin(a.CorrectionOrigins, "finite spectral-action") || !containsOrigin(a.CorrectionOrigins, "extra colored") {
		t.Fatalf("missing correction-origin row: %s", FormatCorrectionOrigins(a.CorrectionOrigins))
	}
	if !a.BoundaryThreshold.SignCompatible || a.BoundaryThreshold.RequiredDeltaAlpha <= 4.0 {
		t.Fatalf("bad boundary threshold slot: %s", FormatBoundaryThreshold(a.BoundaryThreshold))
	}
}

func TestGate609TheoremAndFirewalls(t *testing.T) {
	res := Generation2StrongThresholdSignFieldContentViabilityAuditTheorem().Verify()
	if !res.Passed() {
		t.Fatalf("theorem checks failed: %+v", res)
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusGate608Inherited, StatusRequiredSignClassified, StatusOrdinaryMatterAudited, StatusExtraColoredWrongSign, StatusBoundaryThresholdCompatible, StatusFSABoundarySlotDefined, StatusNoNativeStrongThreshold, StatusNoGaugeUnificationClaim, StatusGate609Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status %s", want)
		}
	}
}
