package generation2nativeupdownsource

import (
	"strings"
	"testing"
)

func TestBuildDefaultGate488NativeUpDownSourceAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Ledger.Executed || a.Ledger.CandidateCount < 7 {
		t.Fatalf("expected source ledger: %+v", a.Ledger)
	}
	if a.Ledger.NativeUpDownLabelSources == 0 || a.Ledger.NativeUniversalFamilyAxes == 0 || a.Ledger.GenerationAwareCandidates == 0 {
		t.Fatalf("expected native labels and universal family axes: %+v", a.Ledger)
	}
	if a.Ledger.SourcesPassingAllRequirements != 0 {
		t.Fatalf("unexpected native CKM source: %+v", a.Ledger)
	}
	if a.Requirements.NativeUpDownOperatorsDerived || a.Requirements.NativeDiagonalizersDerived || a.Requirements.CKMInvariantConstraintsDerived != 0 || a.Requirements.NativeCKMSourceFound {
		t.Fatalf("unexpected native operator theorem: %+v", a.Requirements)
	}
	if !a.Socket.UpDownSectorLabelsNative || !a.Socket.YukawaSlotsNative || a.Socket.YukawaMatrixValuesNative || a.Socket.CanPopulateOuOdNatively || a.Socket.CanComputeUuDaggerUd || !a.Socket.BridgeAirlockRequired {
		t.Fatalf("operator socket firewall failed: %+v", a.Socket)
	}
	if a.Firewall.ObservedCKMImported || a.Firewall.ObservedYukawaEntriesImported || a.Firewall.NativeUpOperatorWritten || a.Firewall.NativeDownOperatorWritten || a.Firewall.CKMMatrixNativePrediction || a.Firewall.NativeRegistryWritten {
		t.Fatalf("firewall leak: %+v", a.Firewall)
	}
}

func TestRenderAuditGate488(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderAudit(a)
	for _, want := range []string{
		"# Gate 488 Registry Audit",
		StatusNativeUpDownSectorLabelsFound,
		StatusNoNativeUpDownEigenbasisSource,
		"O_u/O_d",
		"Yukawa",
		"Gate 489",
	} {
		if !strings.Contains(md, want) {
			t.Fatalf("audit missing %q", want)
		}
	}
}
