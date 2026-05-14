package contactprojectorcompanion

import "testing"

func TestGate279CompanionProjectorAudit(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatalf("Build failed: %v", err)
	}
	if !RootResidualOK(a) {
		t.Fatalf("expected companion + irreducibility + centralizer preflight to pass: %+v", a.Summary)
	}
	if !CompanionTraceOK(a) {
		t.Fatalf("unexpected companion trace/determinant: trace=%s det=%s", RatString(a.Companion.Trace), RatString(a.Companion.Determinant))
	}
	if a.Centralizer.NontrivialIdempotentsOverQ != 0 {
		t.Fatalf("nontrivial rational idempotents must not be derived")
	}
	if a.Centralizer.BlockDiagonalizes2x2OverQ {
		t.Fatalf("irreducible companion module cannot be block-diagonalized over Q")
	}
	if a.Resolvent.ResolventRootAlreadySelected || a.Branch.RBranchMapDerived {
		t.Fatalf("resolvent root or amplitude branch was over-selected")
	}
	if err := AssertNoSuccessOverreach(a); err != nil {
		t.Fatalf("firewall overreach: %v", err)
	}
}

func TestGate279NativeActionsRemainDiagnostics(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatalf("Build failed: %v", err)
	}
	if len(a.NativeActions.Candidates) != 3 {
		t.Fatalf("expected 3 native action candidates, got %d", len(a.NativeActions.Candidates))
	}
	if a.NativeActions.AnyLegalAction || a.NativeActions.AnyCommutingProjector || a.NativeActions.AnyPairSelector {
		t.Fatalf("native action candidates should not become projectors: %+v", a.NativeActions)
	}
	var sawMultiplicity, sawBGap bool
	for _, c := range a.NativeActions.Candidates {
		if c.Name == "Morita multiplicity diag(1,3,3,3)" {
			sawMultiplicity = true
			if c.CommutesWithCompanion || c.Residual <= 0 {
				t.Fatalf("multiplicity diagnostic should fail companion commutation with positive residual: %+v", c)
			}
		}
		if c.Name == "B_gap scalar scale" {
			sawBGap = true
			if !c.CommutesWithCompanion || c.CanSelectRootPair {
				t.Fatalf("B_gap identity action should commute but not select: %+v", c)
			}
		}
	}
	if !sawMultiplicity || !sawBGap {
		t.Fatalf("missing expected diagnostic candidates")
	}
}

func TestGate279TheoremChecksPass(t *testing.T) {
	th := ContactProjectorActionQuarticCompanionModuleSemanticsAuditTheorem()
	res := th.Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("theorem build failed: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check %q failed: %s", c.Name, c.Detail)
		}
	}
}
