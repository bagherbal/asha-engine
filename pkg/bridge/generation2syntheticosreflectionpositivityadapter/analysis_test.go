package generation2syntheticosreflectionpositivityadapter

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestBuildDefaultSyntheticOSAdapter(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Output.Ready || !a.Output.SyntheticOSPositivityVerified || !a.Output.FiniteOSPlumbingVerified {
		t.Fatalf("expected synthetic OS plumbing to pass: %+v", a.Output)
	}
	if a.Output.WickRotationGranted || a.Output.PhysicalHilbertSpaceSelected || a.Output.PositiveEnergyHamiltonianDerived || a.Output.UnitaryRealTimeDynamicsDerived || a.Output.GlobalHyperbolicityGranted || a.Output.ArrowOfTimeSelected {
		t.Fatalf("synthetic OS adapter promoted physical data: %+v", a.Output)
	}
	if !strings.Contains(a.Firewall.Verdict, StatusFailedSyntheticOSNotWick) {
		t.Fatalf("missing Wick firewall status: %s", a.Firewall.Verdict)
	}
}

func TestNativePromotionRowRejected(t *testing.T) {
	b, err := os.ReadFile(resolvePath(DefaultLedger))
	if err != nil {
		t.Fatal(err)
	}
	var ledger OSKernelLedger
	if err := json.Unmarshal(b, &ledger); err != nil {
		t.Fatal(err)
	}
	ledger.Rows[0].NativePromotion = true
	path := filepath.Join(t.TempDir(), "bad_gate534.json")
	out, err := json.MarshalIndent(ledger, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, out, 0644); err != nil {
		t.Fatal(err)
	}
	a, err := BuildFromFile(path)
	if err == nil {
		t.Fatal("expected native promotion fixture to fail validation")
	}
	if a.Import.MetadataComplete || !a.Import.NativePromotionRejected {
		t.Fatalf("native promotion was not rejected: %+v", a.Import)
	}
}

func TestAsTheorem(t *testing.T) {
	res := Generation2SyntheticOSReflectionPositivityKernelAdapterDryRunTheorem().Verify()
	for _, check := range res.Checks {
		if !check.Passed {
			t.Fatalf("check failed: %s: %s", check.Name, check.Detail)
		}
	}
}
