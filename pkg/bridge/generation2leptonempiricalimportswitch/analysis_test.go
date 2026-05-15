package generation2leptonempiricalimportswitch

import (
	"strings"
	"testing"
)

func TestBuildDefaultGate477ValidatesLeptonAirlock(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.Gate476PMNSNullResidualAdapter || !a.Policy.RequiresExplicitTrue || !a.Policy.AllowsPMNSResidualTarget || a.Policy.AllowsPMNSAsRayInput {
		t.Fatalf("bad inheritance/policy: %+v %+v", a.Inheritance, a.Policy)
	}
	if a.Sieve.AcceptedCaseCount != 3 || a.Sieve.RejectedCaseCount != 14 || !a.Sieve.QuarantinedChargedLeptonImportAccepted || !a.Sieve.QuarantinedNeutrinoImportAccepted || !a.Sieve.QuarantinedPMNSResidualTargetAccepted {
		t.Fatalf("bad sieve: %+v", a.Sieve)
	}
	if !a.Sieve.PMNSAsRayInputRejected || !a.Sieve.NativePromotionRejected || !a.Sieve.NativeRegistryWriteRejected || !a.Sieve.PMNSNativePredictionRejected || !a.Sieve.ObservedDataAsTheoremRejected {
		t.Fatalf("unsafe routes not rejected: %+v", a.Sieve)
	}
	if a.Firewall.EmpiricalDataInNativeRegistry || a.Firewall.PMNSMatrixConstructed || a.Firewall.PMNSEntryComputed || a.Firewall.IKNativeSelectorFound || a.Firewall.DENuComputedFromObserved || a.Firewall.DENuNativePrediction {
		t.Fatalf("firewall violated: %+v", a.Firewall)
	}
}

func TestEvaluateImportRejectsPMNSRaySmuggling(t *testing.T) {
	res, accepted, verdict, _ := EvaluateImport(pmnsAsRayInputImport())
	if accepted || res.Imported || verdict != StatusFailedPMNSAsRayInput {
		t.Fatalf("PMNS ray smuggling was not rejected: accepted=%t result=%+v verdict=%s", accepted, res, verdict)
	}
}

func TestRenderAuditContainsGate477Language(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderAudit(a)
	for _, want := range []string{"# Gate 477 Registry Audit", StatusLeptonImportSwitchValid, "lepton-sector-comparator-ledger", "PMNS allowed role = residual target only", "PMNS-as-ray"} {
		if !strings.Contains(md, want) {
			t.Fatalf("audit missing %q", want)
		}
	}
}
