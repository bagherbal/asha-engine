package generation2environmentalflavorsealintegrationhistorytransportaudit

import (
	"strings"
	"testing"
)

func TestGate597IntegratesFlavorSealsIntoHistoryVariables(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Inherited.ChargedLeptonSealName != "ChargedLeptonRootChamberSeal" || a.Inherited.EpsilonNative || a.Inherited.BFlavNative {
		t.Fatalf("bad inherited Gate596 state: %+v", a.Inherited)
	}
	if len(a.SealTable.Rows) < 5 {
		t.Fatalf("seal table too small: %+v", a.SealTable)
	}
	if !strings.Contains(a.Embedding.Verdict, StatusYCoreInserted) || !strings.Contains(a.Embedding.Verdict, StatusOmegaCoreInserted) || !strings.Contains(a.Embedding.Verdict, StatusTCoreDefined) {
		t.Fatalf("embedding statuses missing: %s", a.Embedding.Verdict)
	}
}

func TestGate597FlavorEndMapAndCompression(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.EndMap.BridgeOnly || a.EndMap.NativeDerivation {
		t.Fatalf("end map promoted native derivation: %+v", a.EndMap)
	}
	if len(a.EndMap.RawEnvironmentalInputs) < 5 || len(a.Compression.StillRaw) < 4 {
		t.Fatalf("raw environmental inputs not recorded: end=%+v compression=%+v", a.EndMap, a.Compression)
	}
	if a.Compression.NativeCompression {
		t.Fatalf("unexpected native compression: %+v", a.Compression)
	}
}

func TestGate597TheoremAndFirewalls(t *testing.T) {
	th := Generation2EnvironmentalFlavorSealIntegrationIntoHistoryTransportAuditTheorem()
	res := th.Verify()
	if !res.Passed() {
		t.Fatalf("theorem checks failed: %+v", res)
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusSealIntegrated, StatusYCoreInserted, StatusNoNativeFourthRootTheorem, StatusNoNativeBFlavZero, StatusGate352Preserved, StatusGate597Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status %s", want)
		}
	}
}
