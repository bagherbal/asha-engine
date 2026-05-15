package generation2schwingersourceauthenticityairlock

import (
	"strings"
	"testing"
)

func TestBuildDefaultSchwingerSourceAuthenticityAirlock(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Schema.Executed || a.Schema.RequiredRows != 13 || a.Schema.NativeWriteRows != 0 || !a.Schema.NonSyntheticRequired || !a.Schema.ChecksumRequired {
		t.Fatalf("bad authenticity schema: %+v", a.Schema)
	}
	if !a.Discriminator.SyntheticLedgerRecognized || a.Discriminator.SyntheticLedgerAcceptedAsPhysical || a.Discriminator.NonSyntheticSourceLoaded || a.Discriminator.PhysicalSchwingerAuthenticated {
		t.Fatalf("bad discriminator: %+v", a.Discriminator)
	}
	if !strings.Contains(a.Firewall.Verdict, StatusFailedAuthenticitySchemaNotHamiltonian) || !strings.Contains(a.Truth, "source authenticity") {
		t.Fatalf("missing firewall/truth: firewall=%s truth=%s", a.Firewall.Verdict, a.Truth)
	}
}

func TestAuthenticityRowsAreBridgeOnlyRequired(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	for _, row := range a.Schema.Rows {
		if row.Name == "" || !row.Required || !row.BridgeOnly || row.NativeWrite || row.Reason == "" {
			t.Fatalf("bad row: %+v", row)
		}
	}
}

func TestAsTheorem(t *testing.T) {
	res := Generation2SchwingerSourceAuthenticityComparatorAirlockPreflightTheorem().Verify()
	for _, check := range res.Checks {
		if !check.Passed {
			t.Fatalf("check failed: %s: %s", check.Name, check.Detail)
		}
	}
}
