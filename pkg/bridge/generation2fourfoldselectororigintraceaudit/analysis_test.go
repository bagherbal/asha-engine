package generation2fourfoldselectororigintraceaudit

import "testing"

func TestGate555FourfoldSelectorAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Selector.AllRowsVerified || a.Selector.CommutantDimension != 10 {
		t.Fatalf("selector theorem failed: %+v", a.Selector)
	}
	if a.BMinusL.Split == "" || a.BMinusL.CommutantDimension != 10 || !a.BMinusL.AllBridgeDeltasPMFourThird {
		t.Fatalf("B-L audit failed: %+v", a.BMinusL)
	}
	if a.WeakPlane.RejectedCount != 3 || a.WeakPlane.PreservedCount != 3 || a.WeakPlane.UniqueWeakPlane {
		t.Fatalf("unexpected weak-plane sieve: %+v", a.WeakPlane)
	}
	if a.TauEta.ExistingUnitPreservingPullback || !a.TauEta.CanSelectTwoPlusOneIfPulledBack || a.TauEta.NativeThreeToTwoPlusOne {
		t.Fatalf("tau_eta firewall failed: %+v", a.TauEta)
	}
	if !a.Contact.RegularRepresentationUnit || !a.Contact.IrreducibleOverQ || a.Contact.NontrivialRationalIdempotents != 0 || a.Contact.NativeCarrierAction {
		t.Fatalf("contact quartic firewall failed: %+v", a.Contact)
	}
	if !a.Final.NativeSelectorAlgebraTheorem || a.Final.NativeThreeToTwoPlusOneSelector || !a.Final.TauEtaSealed || !a.Final.ContactRemainsContactOnly {
		t.Fatalf("bad final verdict: %+v", a.Final)
	}
}

func TestGate555Theorem(t *testing.T) {
	res := Generation2FourfoldSelectorOriginAndTraceTransferAuditTheorem().Run()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected failed route: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %+v", c)
		}
	}
}
