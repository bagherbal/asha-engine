package generation2electroweakiksource

import (
	"strings"
	"testing"
)

func TestBuildDefaultGate474FindsNoNativeIKSource(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.Gate473MassClosureFailed || !a.Inheritance.MissingIK {
		t.Fatalf("bad inheritance: %+v", a.Inheritance)
	}
	if a.Sieve.NativeSelectors != 0 || a.Sieve.IKHalfDerived {
		t.Fatalf("must not derive native IK: %+v", a.Sieve)
	}
	if len(a.Sieve.Candidates) != 3 {
		t.Fatalf("expected three source candidates: %+v", a.Sieve.Candidates)
	}
	if a.Sieve.Candidates[0].SuppliesIK || a.Sieve.Candidates[1].SuppliesIK {
		t.Fatalf("Higgs/gauge must not supply IK: %+v", a.Sieve.Candidates)
	}
	if !a.Sieve.Candidates[2].RequiresEmpiricalAirlock || !a.Frontier.CanUsePMNSAsBridgeComparator {
		t.Fatalf("PMNS should be bridge-only frontier: %+v %+v", a.Sieve.Candidates[2], a.Frontier)
	}
	if a.Firewall.NativeRegistryWritten || a.Firewall.IKHalfNative || a.Firewall.PMNSNativePrediction || a.Firewall.CKMNativePrediction {
		t.Fatalf("firewall violated: %+v", a.Firewall)
	}
}

func TestRenderAuditContainsGate474Language(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderAudit(a)
	for _, want := range []string{"# Gate 474 Registry Audit", StatusFailedNoNativeIKSource, StatusFailedHiggsGenerationBlind, "PMNS/lepton", "I_K=0.5 native derivation: not achieved"} {
		if !strings.Contains(md, want) {
			t.Fatalf("audit missing %q", want)
		}
	}
}
