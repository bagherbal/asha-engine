package weakquaternionicnormalization

import "testing"

func TestBuildDefaultComputesTraceMultiplicityLedger(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Summary.TraceWeightsComputed || a.InnerProduct.KappaCRatio != 1 || a.InnerProduct.KappaQRatio != 3 {
		t.Fatalf("expected κ_C:κ_Q=1:3 multiplicity ledger: %s", FormatInnerProduct(a.InnerProduct))
	}
	if !a.InnerProduct.MultiplicitiesGeometric || a.InnerProduct.EdgeNormsDerived {
		t.Fatalf("expected multiplicities but not edge norms: %s", FormatInnerProduct(a.InnerProduct))
	}
}

func TestWeakQuaternionicSelectorIsNotNativeToCPlusM3(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if a.Sieve.WeakQuaternionicNative || a.Sieve.PhysicalSMHilbertDerived {
		t.Fatalf("must not import quaternionic/SM Hilbert structure as native theorem: %s", FormatSieve(a.Sieve))
	}
	if len(a.Sieve.ChiralOrderOneEdgesRetained) != 2 {
		t.Fatalf("expected two retained chiral order-one edges: %s", FormatSieve(a.Sieve))
	}
}

func TestXYRatioAndHiggsRemainBlocked(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if a.XYRatio.XOverYLocked || a.Summary.CanonicalDFDerived || a.Summary.A2A4Derived || a.Summary.HiggsRatioDerived {
		t.Fatalf("Gate 273 over-promoted normalization to amplitude theorem: %s", FormatSummary(a.Summary))
	}
	if !a.SpectralTrace.RatioDependsOnXOverY || a.SpectralTrace.StableInvariant {
		t.Fatalf("expected spectral ratio to vary with x:y: %s", FormatSpectralTrace(a.SpectralTrace))
	}
}

func TestFirewallAndFutureCriteria(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Firewall.MultiplicityNotAmplitude || a.Firewall.FiniteCorePolluted || !a.Firewall.NoSMQuaternionImportedAsTheorem {
		t.Fatalf("firewall violation: %s", FormatFirewall(a.Firewall))
	}
	missing := 0
	for _, c := range a.Future.Criteria {
		if c.Required && !c.Satisfied {
			missing++
		}
	}
	if missing < 6 || !a.Future.NeedNativeWeakQuaternionicAlgebra || !a.Future.NeedEdgeNormOrAmplitudeAction {
		t.Fatalf("expected missing weak/H and amplitude criteria: %s", FormatFuture(a.Future))
	}
}

func TestTheoremPassesWithBridgeRequiredStatus(t *testing.T) {
	res := WeakQuaternionicSubBimoduleSelectorFiniteInnerProductNormalizationAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
