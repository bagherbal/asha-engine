package nativeweakquaternionicalgebra

import "testing"

func TestLocalQuaternionicClosureIsExact(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	q := a.Quaternionic
	if !q.LocalHExtracted || !q.NativeToSelectedDoublet {
		t.Fatalf("expected local H on selected doublet: %s", FormatQuaternionic(q))
	}
	if q.ISquareResidual != 0 || q.JSquareResidual != 0 || q.KSquareResidual != 0 || q.IJMinusKResidual != 0 || q.JIMinusNegativeKResidual != 0 || q.AntiCommutatorResidual != 0 {
		t.Fatalf("quaternionic closure residuals must vanish: %s", FormatQuaternionic(q))
	}
}

func TestCandidateAlgebraNotPromotedToExactSMAlgebra(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Algebra.AssembledOnlyUnderSelector || a.Algebra.ExactSMFiniteAlgebraDerived || a.Quaternionic.GlobalHSummandDerived {
		t.Fatalf("local H was over-promoted: %s :: %s", FormatQuaternionic(a.Quaternionic), FormatAlgebra(a.Algebra))
	}
	if a.Hilbert.ExactPhysicalHFDerived || a.Hilbert.OppositeActionJDerived {
		t.Fatalf("physical H_F/J must remain missing: %s", FormatHilbert(a.Hilbert))
	}
}

func TestQuaternionicDoesNotLockXYOrHiggsRatio(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if a.Amplitude.XOverYLocked || a.Amplitude.EdgeNormCSelected || a.Amplitude.EdgeNormQSelected || a.Summary.A2A4Derived || a.Summary.HiggsRatioDerived {
		t.Fatalf("Gate 274 over-derived dynamics: %s", FormatSummary(a.Summary))
	}
	if !a.SpectralTrace.RatioDependsOnXOverY || a.SpectralTrace.StableInvariant {
		t.Fatalf("spectral ratio should still depend on x:y: %s", FormatSpectralTrace(a.SpectralTrace))
	}
}

func TestFirewallAndFutureCriteria(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if a.Firewall.FiniteCorePolluted || !a.Firewall.LocalHNotPromotedToGlobalH || !a.Firewall.MultiplicityNotAmplitude {
		t.Fatalf("firewall violation: %s", FormatFirewall(a.Firewall))
	}
	missing := 0
	for _, c := range a.Future.Criteria {
		if c.Required && !c.Satisfied {
			missing++
		}
	}
	if missing < 5 || !a.Future.NeedPhysicalFiniteHF || !a.Future.NeedEdgeNormAction {
		t.Fatalf("expected remaining future obligations: %s", FormatFuture(a.Future))
	}
}

func TestTheoremPassesWithBridgeRequiredStatus(t *testing.T) {
	res := NativeWeakQuaternionicAlgebraPhysicalFiniteHilbertSpaceReconstructionAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
