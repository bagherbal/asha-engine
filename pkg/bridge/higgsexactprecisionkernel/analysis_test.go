package higgsexactprecisionkernel

import "testing"

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if a.Inputs.HighestInheritedGate != inheritedHighestGate || a.Inputs.ContactShapeRational.Cmp(rat(1197, 4624)) != 0 {
		t.Fatalf("bad inputs: %s", FormatInputs(a.Inputs))
	}
}

func TestHighPrecisionPiKernel(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !nearlyFloat(a.Pi.Pi, 3.141592653589793, 1e-15) || !nearlyFloat(a.Pi.AlphaInverse, 25.132741228718345, 1e-14) {
		t.Fatalf("bad pi kernel: %s", FormatPi(a.Pi))
	}
}

func TestNativeClosedFormMass(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if a.Native.LambdaH.RatString() != "1197/9248" {
		t.Fatalf("bad lambda rational: %s", FormatNative(a.Native))
	}
	if !nearlyFloat(a.Native.MassGeV, 125.27415714969897, 1e-12) {
		t.Fatalf("bad mass: %s", FormatNative(a.Native))
	}
}

func TestExactGap(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if a.Gap.RequiredRePiGeV2.RatString() != "504067437/11560000" {
		t.Fatalf("bad exact RePi target: %s", FormatGap(a.Gap))
	}
	if !nearlyFloat(a.Gap.DeltaMassGeV, 0.17415714969897194, 1e-13) {
		t.Fatalf("bad delta mass: %s", FormatGap(a.Gap))
	}
}

func TestFirewallsAndStatuses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Firewalls.NoPVContraction || !a.Firewalls.NoColliderPoleClaim || a.Efficiency.FullPVContractionExecuted {
		t.Fatalf("bad firewalls: %s %s", FormatFirewalls(a.Firewalls), FormatEfficiency(a.Efficiency))
	}
	statuses := Statuses(a)
	required := []string{StatusNativeClosedFormComputed, StatusSelfEnergyTargetComputed, StatusFailedFullPVContraction, StatusFailedColliderClaim}
	for _, req := range required {
		found := false
		for _, got := range statuses {
			if got == req {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("missing status %s in %v", req, statuses)
		}
	}
}

func TestTheoremPasses(t *testing.T) {
	res := ExactNativeHiggsPredictionArbitraryPrecisionNumericalKernelAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
