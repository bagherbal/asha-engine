package generation2sectordeformationsource

import (
	"strings"
	"testing"
)

func TestBuildDefaultGate482NoNativeSource(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.Gate481BaselineCancellation || a.Inheritance.IKVac != 0.5 {
		t.Fatalf("bad inheritance: %+v", a.Inheritance)
	}
	if a.Candidates.NativeSourceFound {
		t.Fatalf("unexpected native source: %+v", a.Candidates)
	}
	if a.Sieve.CandidatesPassingNativeSource != 0 || a.Sieve.DeltaAlphaNative || a.Sieve.DeltaPhiNative {
		t.Fatalf("bad source sieve: %+v", a.Sieve)
	}
	if a.Sieve.AllZeroPerturbationDistance != 0 || !a.Sieve.AllZeroWouldPredictNoMixing {
		t.Fatalf("bad all-zero limit: %+v", a.Sieve)
	}
	if !a.BridgeSlot.RequiresAirlock || !a.BridgeSlot.RequiresBranchTags || !a.BridgeSlot.RejectsCKMPMNSAsInput || a.BridgeSlot.CanComputePhysicalResidual {
		t.Fatalf("bad bridge slot: %+v", a.BridgeSlot)
	}
	if a.Firewall.SectorPerturbationsNative || a.Firewall.PhysicalDUDComputed || a.Firewall.PhysicalDENuComputed || a.Firewall.NativeRegistryWritten {
		t.Fatalf("firewall leak: %+v", a.Firewall)
	}
}
func TestRenderAuditGate482(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderAudit(a)
	for _, want := range []string{StatusFailedNativeSourceAbsent, "finite orientation / triality family address", "sector-perturbation-source-ledger", "physical d_ud = undefined"} {
		if !strings.Contains(md, want) {
			t.Fatalf("audit missing %q", want)
		}
	}
}
