package higgspassarinoveltmankernel

import "testing"

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if a.Inputs.HighestInheritedGate != inheritedHighestGate || a.Inputs.NativeRunMassGeV <= 125 {
		t.Fatalf("bad inputs: %s", FormatInputs(a.Inputs))
	}
}

func TestPVBasisAndValues(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Basis.EqualMassLane || !a.Basis.BelowThresholdOnly {
		t.Fatalf("bad basis: %s", FormatBasis(a.Basis))
	}
	if len(a.PV.Values) != 4 || !allFiniteBelowThreshold(a.PV) {
		t.Fatalf("bad PV values: %s", FormatPVLedger(a.PV))
	}
	if !nearlyEqual(a.PV.Values[0].B0Finite, -0.550445463802, 1e-9) {
		t.Fatalf("unexpected top B0 finite: %s", FormatPVValue(a.PV.Values[0]))
	}
}

func TestKernelSlotsAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Slots.AllBasisAvailable || a.Slots.FullKernelClosed || !noFullCoefficients(a.Slots) {
		t.Fatalf("bad slots: %s", FormatSlots(a.Slots))
	}
	if !a.Firewalls.NoCoefficientTable || !a.Firewalls.NoCounterterms || !a.Firewalls.NoExactPoleClaim {
		t.Fatalf("bad firewalls: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestStatuses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	statuses := Statuses(a)
	required := []string{StatusPVBasisFormalized, StatusFinitePVFunctionsComputed, StatusFailedFullCoefficientTable, StatusFailedExactPoleMass}
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
	res := HiggsPassarinoVeltmanPoleKernelFiniteIntegralInstallationAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
