package modulartimeflowvacuumselector

import "testing"

func TestPathBActivated(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Span.InheritedGate != 361 || a.Span.Path != "B: minimal dynamical extension" || a.Span.AddsFit {
		t.Fatalf("bad span: %s", FormatSpan(a.Span))
	}
}

func TestDocumentationShift(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Docs.Required || !a.Docs.AppliedInReadme || !a.Docs.AppliedInDocs {
		t.Fatalf("documentation shift not formalized: %s", FormatDocs(a.Docs))
	}
}

func TestFlowAdmissibilitySieve(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Sieve.Executed || len(a.Sieve.Axioms) < 5 || !a.Sieve.Candidate.NewOperatorClass || !a.Sieve.PreservesLandscape {
		t.Fatalf("bad sieve: %s", FormatSieve(a.Sieve))
	}
}

func TestKernelStillMissingAndCensusPreserved(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Sieve.ExplicitKernelConstructed || a.Sieve.VacuumSelected || a.Census.RemainingInputs != 15 || a.Census.SevenSealReached {
		t.Fatalf("bad kernel/census state: %s / %s", FormatSieve(a.Sieve), FormatCensus(a.Census))
	}
}

func TestStatuses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	statuses := Statuses(a)
	required := []string{StatusPathBActivated, StatusReadmeShiftFormalized, StatusFlowOperatorClassIntroduced, StatusAdmissibilitySieveExecuted, StatusFailedExplicitFlowKernelMissing, StatusFailedVacuumPointNotSelected}
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
	res := ModularTimeFlowVacuumSelectorExtensionAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
