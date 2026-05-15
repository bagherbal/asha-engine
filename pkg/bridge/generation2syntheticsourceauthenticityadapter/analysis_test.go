package generation2syntheticsourceauthenticityadapter

import (
	"strings"
	"testing"
)

func TestBuildDefaultSyntheticSourceAuthenticityAdapter(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Import.Loaded || a.Import.Rows != 13 || a.Import.AcceptedRows != 13 || !a.Import.ChecksumVerified || !a.Import.RequiredSchemaRowsMatched {
		t.Fatalf("bad import: %+v", a.Import)
	}
	if !a.Output.AuthenticityPlumbingVerified || !a.Output.SyntheticFixtureRejectedAsPhysical || a.Output.PhysicalSourceAuthenticated || a.Output.NativePromotionRows != 0 || a.Output.PhysicalClaimRows != 0 {
		t.Fatalf("bad adapter output: %+v", a.Output)
	}
	if !strings.Contains(a.Firewall.Verdict, StatusFailedSyntheticNotPhysicalSource) || !strings.Contains(a.Truth, "synthetic provenance remains synthetic") {
		t.Fatalf("missing firewall/truth: firewall=%s truth=%s", a.Firewall.Verdict, a.Truth)
	}
}

func TestLedgerRowsAreQuarantined(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Import.AllRowsBridgeOnly || !a.Import.AllRowsNoTheoremInput || !a.Import.AllRowsSynthetic || a.Import.AnyPhysicalClaim || a.Import.AnyNativePromotionClaim {
		t.Fatalf("metadata sieve failed: %+v", a.Import)
	}
}

func TestAsTheorem(t *testing.T) {
	res := Generation2SyntheticSourceAuthenticityLedgerAdapterRejectionDryRunTheorem().Verify()
	for _, check := range res.Checks {
		if !check.Passed {
			t.Fatalf("check failed: %s: %s", check.Name, check.Detail)
		}
	}
}
