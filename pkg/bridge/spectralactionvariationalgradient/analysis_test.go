package spectralactionvariationalgradient

import "testing"

func TestModuliFormalization(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Moduli.InheritedGate != 345 || a.Moduli.TotalMinimalCoordinates != minimalVacuumCoordinates || a.Moduli.ContinuousCount != 15 || a.Moduli.ImportedObservedMasses {
		t.Fatalf("bad moduli ledger: %s", FormatModuli(a.Moduli))
	}
}

func TestVariationalActionFlatness(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.Action.Terms) < 4 || a.Action.FlavorFlatTerms < 2 || a.Action.FlavorSelectingTerms < 2 || a.Action.UsesObservedMassFit {
		t.Fatalf("bad action ledger: %s", FormatAction(a.Action))
	}
}

func TestGradientSieve(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gradient.StandardInvariantGradientZero || a.Gradient.PositiveMetricTopMinimum <= 0 || a.Gradient.SignedProjectionRank != 1 || a.Gradient.SignedProjectionNullity != 2 || !a.Gradient.SignedMinimumDegenerate || a.Gradient.SelectsUniqueTopDirection {
		t.Fatalf("bad gradient sieve: %s", FormatGradient(a.Gradient))
	}
}

func TestTopNullingCapacityButNoNativeSelection(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.TopTest.RecoveredGate322Envelope || !nearlyZero(a.TopTest.DotA) || !nearlyZero(a.TopTest.DotB) || a.TopTest.SignedMinimum != 0 || a.TopTest.NativeSelection {
		t.Fatalf("bad top test: %s", FormatTopTest(a.TopTest))
	}
}

func TestVerdictAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Verdict.VariationalVacuumActive || !a.Verdict.GradientFlat || !a.Verdict.NullspaceCapacity || a.Verdict.UniqueVacuumSelected || !a.Verdict.RequiresNewOperator {
		t.Fatalf("bad verdict: %s", FormatVerdict(a.Verdict))
	}
	if !a.Audit.NoObservedYukawasImported || !a.Audit.NoCKMTextureInvented || !a.Audit.NoTopNullingForced || !a.Audit.NoCosmologicalFit || !a.Audit.NoFinalVacuumClaim {
		t.Fatalf("bad audit: %s", FormatAudit(a.Audit))
	}
}

func TestStatuses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	statuses := Statuses(a)
	required := []string{StatusModuliFieldsFormalized, StatusVariationalActionFormalized, StatusGradientSieveExecuted, StatusFailedVariationalVacuumSelection, StatusFailedUniqueCKMTexture, StatusFailedNativeTopSuppression}
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
	res := SpectralActionVariationalGradientPhaseIIIVacuumInitializationSieveTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
