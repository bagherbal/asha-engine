package generation2commonscaleledger

import (
	"strings"
	"testing"
)

func TestBuildDefaultGate467(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Schema.AcceptedLedgers != RequiredSectors || !a.Schema.CompleteUSectorAccepted || !a.Schema.CompleteDSectorAccepted || !a.Schema.BothSectorsReady {
		t.Fatalf("expected complete u and d schema ledgers: %+v", a.Schema)
	}
	if !a.Schema.MixedScaleRejected || !a.Schema.MissingIKRejected || !a.Schema.MissingBranchRejected || !a.Schema.MissingUncertaintyRejected || !a.Schema.MassOnlyRejected || !a.Schema.CabibboRayInputRejected || !a.Schema.NativePromotionRejected {
		t.Fatalf("expected fail-closed rejection probes: %+v", a.Schema)
	}
	if a.Schema.DUDComputedNow || a.Firewall.DUDNativePrediction || a.Firewall.CKMNativePrediction || a.Firewall.NativeRegistryWritten || a.Firewall.CabibboUsedAsRayInput {
		t.Fatalf("Gate467 must not compute or promote d_ud/CKM: schema=%+v firewall=%+v", a.Schema, a.Firewall)
	}
}

func TestEvaluateLedgerRejectsMassOnlyAndCabibboRay(t *testing.T) {
	massOnly := canonicalLedgers()[0]
	massOnly.HasIK = false
	massOnly.HasCPOddSign = false
	massOnly.HasC3Sheet = false
	res := EvaluateLedger(massOnly)
	if res.Accepted || !contains(res.Failures, StatusFailedMissingIKRejected) || !contains(res.Failures, StatusFailedMassOnlyStillRankOne) {
		t.Fatalf("mass-only ledger should fail rank completion: %+v", res)
	}
	cab := canonicalLedgers()[0]
	cab.Sector = "u-d"
	cab.MassLabels = []string{"|V_us|"}
	cab.CabibboAsRayInput = true
	res = EvaluateLedger(cab)
	if res.Accepted || !contains(res.Failures, StatusFailedCabibboAsRayInput) {
		t.Fatalf("Cabibbo-as-ray input should fail closed: %+v", res)
	}
}

func TestRenderAuditContainsGate467Language(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderAudit(a)
	for _, want := range []string{"# Gate 467 Registry Audit", StatusBridgeOnlyDesignValidated, "I_spec", "I_K", "{sigma_CP,n_C3}", StatusFailedCabibboAsRayInput} {
		if !strings.Contains(md, want) {
			t.Fatalf("audit missing %q", want)
		}
	}
}

func contains(xs []string, want string) bool {
	for _, x := range xs {
		if x == want {
			return true
		}
	}
	return false
}
