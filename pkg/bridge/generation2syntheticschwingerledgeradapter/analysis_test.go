package generation2syntheticschwingerledgeradapter

import (
	"strings"
	"testing"
)

func TestBuildDefaultSyntheticSchwingerAdapter(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Import.Loaded || a.Import.Rows != 19 || a.Import.AcceptedRows != 19 || !a.Import.RequiredSchemaRowsMatched {
		t.Fatalf("bad import: %+v", a.Import)
	}
	if !a.Output.SyntheticSchwingerAdapterVerified || a.Output.NegativeQuadraticVectors != 0 || a.Output.NativePromotionRows != 0 || a.Output.ObservedRows != 0 {
		t.Fatalf("bad adapter output: %+v", a.Output)
	}
	if !strings.Contains(a.Firewall.Verdict, StatusFailedSyntheticNotWick) || !strings.Contains(a.Truth, "not Wick rotation") {
		t.Fatalf("missing firewall/truth language: firewall=%s truth=%s", a.Firewall.Verdict, a.Truth)
	}
}

func TestLedgerRowsAreBridgeOnlyNoTheoremInput(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Import.AllRowsBridgeOnly || !a.Import.AllRowsNoTheoremInput || !a.Import.AllRowsSourceTagged || !a.Import.AllRowsConventionTagged {
		t.Fatalf("metadata sieve failed: %+v", a.Import)
	}
}

func TestAsTheorem(t *testing.T) {
	res := Generation2SyntheticSchwingerFunctionSourceLedgerAdapterDryRunTheorem().Verify()
	for _, check := range res.Checks {
		if !check.Passed {
			t.Fatalf("check failed: %s: %s", check.Name, check.Detail)
		}
	}
}
