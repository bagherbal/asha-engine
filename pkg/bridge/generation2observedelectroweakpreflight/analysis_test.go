package generation2observedelectroweakpreflight

import (
	"strings"
	"testing"
)

func TestGate506ObservedElectroweakPreflight(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault: %v", err)
	}
	if !a.Inheritance.SyntheticAdapterExecuted || !a.Inheritance.SyntheticOnly || a.Inheritance.Gate505ObservedDataImported || a.Inheritance.Gate505NativeDataImported || !a.Inheritance.Gate505NativeWriteBlocked || !a.Inheritance.Gate506RedirectDefined {
		t.Fatalf("bad inheritance: %+v", a.Inheritance)
	}
	if !a.Preflight.Executed || a.Preflight.AcceptedSchemaCases != 1 || a.Preflight.RejectedCases != 10 || a.Preflight.ReadyForNumericalAdapterCases != 0 || a.Preflight.NumericalAdapterRun || a.Preflight.ObservedNumbersImported || !a.Preflight.AllAcceptedBridgeOnlyObserved {
		t.Fatalf("bad preflight counts: %+v", a.Preflight)
	}
	if !a.Preflight.SwitchClosedRejected || !a.Preflight.MissingVEVRejected || !a.Preflight.MissingGaugeCouplingRejected || !a.Preflight.MissingScaleSchemeRejected || !a.Preflight.MissingSourceUncertaintyRejected || !a.Preflight.ObservedMassAsNativeInputRejected || !a.Preflight.WeakAngleNativePromotionRejected || !a.Preflight.KappaPromotionRejected || !a.Preflight.NativePromotionRejected {
		t.Fatalf("missing rejection coverage: %+v", a.Preflight)
	}
	if a.Firewall.ObservedNumbersImported || a.Firewall.NumericalAdapterExecuted || a.Firewall.NativeVEVWritten || a.Firewall.NativeGaugeCouplingWritten || a.Firewall.NativeWeakAngleWritten || a.Firewall.NativeWZMassWritten || a.Firewall.NativeKappaWritten || a.Firewall.NativeRegistryWritten || a.Firewall.NativeElectroweakPredictionMade {
		t.Fatalf("firewall violation: %+v", a.Firewall)
	}
	if a.Next.Gate != 507 {
		t.Fatalf("expected Gate507 redirect, got %+v", a.Next)
	}
	md := Markdown(a)
	for _, want := range []string{"# Gate 506 Registry Audit", StatusObservedEWPreflightValidated, StatusFailedObservedMassAsNativeInput, StatusFirewallNativeWriteBlocked, "Gate 507"} {
		if !strings.Contains(md, want) {
			t.Fatalf("markdown missing %q", want)
		}
	}
}

func TestGate506TheoremPasses(t *testing.T) {
	res := Generation2ObservedElectroweakComparatorAirlockPreflightTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem failed:\n%s", res.Details())
	}
}
