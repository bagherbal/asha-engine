package fullflavorledgerclosure

import "testing"

func TestBuildDefaultClosesFlavorLedger(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Summary.FullFlavorLedgerClosed {
		t.Fatalf("expected full flavor ledger to close: %+v", a.Summary)
	}
	if a.Summary.NativeFlavorDynamicsDerived {
		t.Fatalf("closure gate must not derive native flavor dynamics")
	}
	if !a.Firewall.KinematicsDerived || !a.Firewall.DynamicsSealed {
		t.Fatalf("expected derived kinematics and sealed dynamics: %+v", a.Firewall)
	}
	if !a.Firewall.NoMassPredictionClaim || !a.Firewall.NoCKMPMNSPredictionClaim {
		t.Fatalf("closure gate must not claim flavor predictions: %+v", a.Firewall)
	}
	if a.Firewall.FiniteCorePolluted {
		t.Fatalf("finite core polluted: %+v", a.Firewall)
	}
}

func TestLedgersSeparateDerivedAndSealedData(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if len(a.Geometric.Items) < 7 {
		t.Fatalf("geometric ledger too small: %d", len(a.Geometric.Items))
	}
	if len(a.Empirical.Items) < 6 {
		t.Fatalf("empirical ledger too small: %d", len(a.Empirical.Items))
	}
	if !a.Geometric.TauEtaSourceMapRecorded || !a.Geometric.AdTauMixingComplementRecorded || !a.Geometric.TrialityHermitianBasisRecorded {
		t.Fatalf("missing derived structural flavor ledgers: %+v", a.Geometric)
	}
	if a.Geometric.YukawaAmplitudeDerived || a.Geometric.FermionMassesDerived || a.Geometric.CKMPMNSDerived {
		t.Fatalf("geometric ledger illegally promoted empirical data: %+v", a.Geometric)
	}
	if !a.Empirical.QuarkTexturesSealed || !a.Empirical.LeptonTexturesSealed || !a.Empirical.MajoranaChoiceSealed {
		t.Fatalf("empirical ledger failed to quarantine flavor inputs: %+v", a.Empirical)
	}
}

func TestFutureCriteriaRequireSpectralActionBeforeSealLift(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if a.FutureCriteria.CurrentGateCanLiftSeal {
		t.Fatalf("Gate 267 must not lift the flavor seals")
	}
	if !a.FutureCriteria.RequiresFiniteSpectralAction || !a.FutureCriteria.RequiresCanonicalFiniteDirac || !a.FutureCriteria.RequiresYukawaAmplitudeMap {
		t.Fatalf("missing future theorem requirements: %+v", a.FutureCriteria)
	}
	unsatisfied := 0
	for _, c := range a.FutureCriteria.Criteria {
		if c.Required && !c.Satisfied {
			unsatisfied++
		}
	}
	if unsatisfied < 6 {
		t.Fatalf("expected all future criteria to remain unsatisfied, got %d", unsatisfied)
	}
}

func TestTheoremPassesWithBridgeRequiredStatus(t *testing.T) {
	th := FullFlavorLedgerClosureQuarkLeptonEmpiricalFirewallSummaryAuditTheorem()
	res := th.Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
