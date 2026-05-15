package generation2rayinversion

import (
	"math"
	"strings"
	"testing"
)

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if a.Sieve.ValidDomainCount != 3 || a.Sieve.RejectedDomainCount != 2 {
		t.Fatalf("unexpected branch sieve: %s", FormatSieve(a.Sieve))
	}
	if a.Inverse.BranchCountGeneric != PhiBranchCount || a.Inverse.ExportsNativeRay {
		t.Fatalf("inverse leaked native ray or wrong branch count: %s", FormatInverse(a.Inverse))
	}
}

func TestInverseFormulaConsistency(t *testing.T) {
	s := classify(Sample{Name: "test", IK: 0.3, ISpec: 0.05})
	if !s.InsideDomain || s.AtCaustic {
		t.Fatalf("sample should be generic interior: %s", FormatSample(s))
	}
	ikBack := s.Alpha / math.Sqrt(s.Alpha*s.Alpha+3)
	if math.Abs(ikBack-s.IK) > 1e-12 {
		t.Fatalf("alpha inverse not consistent: got I_K %.15g want %.15g", ikBack, s.IK)
	}
	ispecBack := 2 * s.CosThreePhi / math.Pow(s.Alpha*s.Alpha+3, 1.5)
	if math.Abs(ispecBack-s.ISpec) > 1e-12 {
		t.Fatalf("cos3phi inverse not consistent: got I_spec %.15g want %.15g", ispecBack, s.ISpec)
	}
}

func TestCausticAndOutsideDomain(t *testing.T) {
	caustic := classify(Sample{Name: "caustic", IK: 0, ISpec: 2 / (3 * math.Sqrt(3))})
	if !caustic.InsideDomain || !caustic.AtCaustic || caustic.GenericBranchCount != 3 {
		t.Fatalf("caustic not detected: %s", FormatSample(caustic))
	}
	outside := classify(Sample{Name: "outside", IK: 0.2, ISpec: 1})
	if outside.InsideDomain || outside.AllowedAsBridgeDryRun || outside.AllowedAsNativeExport {
		t.Fatalf("outside-domain sample did not fail closed: %s", FormatSample(outside))
	}
}

func TestTheoremPasses(t *testing.T) {
	res := Generation2SymbolicCoefficientRayInversionBranchCausticMapTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem audit checks did not pass:\n%s", res.Details())
	}
}

func TestRenderAuditContainsKeyFormulas(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderAudit(a)
	for _, want := range []string{StatusBridgeOnlyInversionValidated, "alpha = sqrt(3) I_K", "cos(3 phi)", "sin(3 phi)=0", StatusFailedGlobalUniqueRayAbsent} {
		if !strings.Contains(md, want) {
			t.Fatalf("audit missing %q", want)
		}
	}
}
