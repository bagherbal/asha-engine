package generation2boundaryalphasourceanddomaintransportaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate827CoefficientSources(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(a.Coefficients.LinearCoeff-3.0/10.0) > 1e-15 {
		t.Fatalf("bad linear coefficient: %s", FormatCoefficients(a.Coefficients))
	}
	if math.Abs(a.Coefficients.QuadraticCoeff-7.0/72.0) > 1e-15 {
		t.Fatalf("bad quadratic coefficient: %s", FormatCoefficients(a.Coefficients))
	}
	if !a.Coefficients.DimensionRatiosVerified {
		t.Fatalf("dimension ratios not verified: %s", FormatCoefficients(a.Coefficients))
	}
}

func TestGate827AlphaDecompositionAndNonCircularity(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(a.Alpha.Alpha-0.0003878958469680527) > 1e-18 {
		t.Fatalf("bad alpha: %s", FormatAlpha(a.Alpha))
	}
	if math.Abs(a.Alpha.LinearContribution+a.Alpha.QuadraticContribution-a.Alpha.Alpha) > 1e-21 {
		t.Fatalf("alpha contributions do not add: %s", FormatAlpha(a.Alpha))
	}
	if a.NonCircular.UsesNEffToDefineAlpha || a.NonCircular.UsesObservedYukawas || a.NonCircular.UsesHiggsMass {
		t.Fatalf("noncircularity failed: %s", FormatNonCircularity(a.NonCircular))
	}
}

func TestGate827TransportFirewallAndTheorem(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Transport.UnifiedTransportCertified || a.Transport.NativeAlphaTheoremCertified {
		t.Fatalf("transport over-promoted: %s", FormatTransport(a.Transport))
	}
	if a.Impact.CanUpdateCYukawa || a.Impact.CanUpdateCHiggs || a.Impact.CanUpdateNEff {
		t.Fatalf("impact firewall failed: %s", FormatImpact(a.Impact))
	}
	res := Generation2BoundaryAlphaSourceAndDomainTransportAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("failed check %s: %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status %s", want)
		}
	}
}
