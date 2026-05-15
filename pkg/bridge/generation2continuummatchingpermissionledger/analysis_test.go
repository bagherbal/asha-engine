package generation2continuummatchingpermissionledger

import (
	"strings"
	"testing"
)

func TestGate504ContinuumMatchingPermissionLedger(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault: %v", err)
	}
	if !a.Inheritance.Gate503ConditionalIndexAccepted || !a.Inheritance.Gate503NonzeroRayAssumed || a.Inheritance.Gate503UnconditionalVacuumProven || a.Inheritance.Gate503WZMassDerived || !a.Inheritance.YukawaTraceBridgeScalarNorm || a.Inheritance.YukawaTraceNativeNumeric {
		t.Fatalf("bad inheritance: %+v", a.Inheritance)
	}
	if !a.Schema.Executed || a.Schema.NativeRows != 0 || a.Schema.BridgeRows != len(a.Schema.Rows) || a.Schema.RowsRequiringExplicitValues < 3 || a.Schema.RowsRequiringSchemeScale != len(a.Schema.Rows) {
		t.Fatalf("bad permission schema: %+v", a.Schema)
	}
	if !a.Formula.TreeLevelWZFormulasDefined || !a.Formula.RequiresExplicitVEV || !a.Formula.RequiresExplicitGaugeCouplings || a.Formula.ComputesNow || !a.Formula.PhotonZeroSymbolic || a.Formula.NativeWeakAngleDerived || a.Formula.NativeWZMassesDerived || a.Formula.NativeKappaPromoted {
		t.Fatalf("formula ledger over-promoted: %+v", a.Formula)
	}
	if !a.Boundary.PermissionLedgerAccepted || !a.Boundary.ContinuumAdapterMayComputeWithExplicitInputs || a.Boundary.NumericalAdapterExecuted || a.Boundary.NativeVEVSelected || a.Boundary.NativeGaugeCouplingsSelected || a.Boundary.NativeWeakAngleDerived || a.Boundary.NativeWZMassesDerived || a.Boundary.NativeKappaSelected || !a.Boundary.YukawaTraceStillEnvironmental {
		t.Fatalf("boundary over-promoted: %+v", a.Boundary)
	}
	if a.Firewall.ObservedVEVImported || a.Firewall.ObservedGaugeCouplingsImported || a.Firewall.ObservedWeakAngleImported || a.Firewall.ObservedWMassImported || a.Firewall.ObservedZMassImported || a.Firewall.ObservedYukawaImported || a.Firewall.NativeVEVWritten || a.Firewall.NativeGaugeCouplingWritten || a.Firewall.NativeWeakAngleWritten || a.Firewall.NativeWZMassWritten || a.Firewall.NativeKappaWritten {
		t.Fatalf("firewall violation: %+v", a.Firewall)
	}
	if a.Next.Gate != 505 {
		t.Fatalf("expected Gate505 redirect, got %+v", a.Next)
	}
	md := Markdown(a)
	for _, want := range []string{"# Gate 504 Registry Audit", "m_W = g2 v / 2", StatusPermissionLedgerConstructed, StatusFailedWZMassesNotNative, "Gate 505"} {
		if !strings.Contains(md, want) {
			t.Fatalf("markdown missing %q", want)
		}
	}
}

func TestGate504TheoremPasses(t *testing.T) {
	res := Generation2ContinuumMatchingPermissionLedgerAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem failed:\n%s", res.Details())
	}
}
