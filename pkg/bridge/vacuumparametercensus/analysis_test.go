package vacuumparametercensus

import "testing"

func TestFailureLedgerPattern(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Failures.HighestGateInherited != highestInheritedGate || len(a.Failures.Clusters) < 5 || a.Failures.TypeACount != 4 || a.Failures.TypeBCount != 1 || !a.Failures.LandscapeNotVacuum {
		t.Fatalf("bad failure ledger: %s", FormatFailures(a.Failures))
	}
}

func TestLandscapeBoundaryConstraints(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Landscape.NativeBoundaryConstraintCount != 4 || !a.Landscape.ContainsWeakMixing || !a.Landscape.ContainsHiggsGaugeRatio || !a.Landscape.ContainsAlphaGUT || !a.Landscape.ContainsHierarchy || !a.Landscape.ContainsGaugeGroup || !a.Landscape.ContainsMatterContent || !a.Landscape.ContainsGenerations || !a.Landscape.ContainsMoritaSplit {
		t.Fatalf("bad landscape ledger: %s", FormatLandscape(a.Landscape))
	}
}

func TestMinimalSM19Census(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.MinimalSM.BaselineCount != 19 || a.MinimalSM.MinimalInputCount != 15 || a.MinimalSM.RemainingContinuousDim != 15 || len(a.MinimalSM.RemainingVacuumInputs) != 4 {
		t.Fatalf("bad minimal census: %s", FormatMinimal(a.MinimalSM))
	}
	total := 0
	for _, p := range a.MinimalSM.RemainingVacuumInputs {
		total += p.Count
	}
	if total != 15 {
		t.Fatalf("remaining count = %d, want 15", total)
	}
}

func TestExtendedLedgerSeparated(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Extended.IncludeNeutrinos || !a.Extended.IncludeCosmology || a.Extended.AddedContinuousDim != 10 || a.Extended.TotalExtendedDim != 25 || !a.Extended.ModelDependent {
		t.Fatalf("bad extended ledger: %s", FormatExtended(a.Extended))
	}
}

func TestMinimalInputTheoremAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Theorem.ProvesLandscapeOnly || a.Theorem.DerivesVacuumPoint || a.Theorem.MinimalSMVacuumDim != 15 || a.Theorem.ExtendedVacuumDim != 25 || len(a.Theorem.DiscreteSeals) < 4 {
		t.Fatalf("bad theorem: %s", FormatTheorem(a.Theorem))
	}
	if !a.Audit.NoYukawaFitInserted || !a.Audit.NoCKMInvented || !a.Audit.NoPMNSInvented || !a.Audit.NoCosmologicalConstantFit || !a.Audit.NoVacuumDirectionForced || !a.Audit.NoPrecisionClaimInserted || a.Audit.FinalTOEClaimed {
		t.Fatalf("bad audit: %s", FormatAudit(a.Audit))
	}
}

func TestStatuses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	statuses := Statuses(a)
	required := []string{StatusMinimalInputTheoremFormalized, StatusFailedYukawaAmplitudesRemainVacuum, StatusFailedCKMTextureRemainsVacuum, StatusFailedCosmologicalConstantVacuum, StatusFailedFinalVacuumNotDerived}
	for _, req := range required {
		found := false
		for _, s := range statuses {
			if s == req {
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
	res := VacuumParameterCensusMinimalInputTheoremAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
