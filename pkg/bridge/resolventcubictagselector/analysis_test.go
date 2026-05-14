package resolventcubictagselector

import "testing"

func TestResolventRetrieved(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault error: %v", err)
	}
	if !a.Resolvent.Retrieved || !a.Resolvent.EncodesTwoPlusTwo || len(a.Resolvent.QuarticRoots) != 4 || len(a.Resolvent.Branches) != 3 {
		t.Fatalf("bad resolvent: %s", FormatResolvent(a.Resolvent))
	}
	if a.Resolvent.CanonicalRootPreviouslySelected {
		t.Fatalf("resolvent root should not be preselected: %s", FormatResolvent(a.Resolvent))
	}
}

func TestTopologicalTagsReachOnlySectorLabels(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault error: %v", err)
	}
	if !a.Tags.TauEtaBindsUD || !a.Tags.BGapTagsNeutrino || !a.Tags.TagsReachSectorLabels || a.Tags.TagsReachQuarticRoots {
		t.Fatalf("bad tags: %s", FormatTags(a.Tags))
	}
}

func TestSectorPairingSelectedButContactRootNot(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault error: %v", err)
	}
	if a.Sieve.TotalCandidates != 3 || a.Sieve.SurvivingSectorPairings != 1 || a.Sieve.SelectedSectorPairing != "{u,d}|{e,nu}" || !a.Sieve.UniqueSectorPairing {
		t.Fatalf("sector selector failed: %s", FormatSieve(a.Sieve))
	}
	if a.Sieve.UniqueContactRoot {
		t.Fatalf("contact root was overpromoted: %s", FormatSieve(a.Sieve))
	}
}

func TestGate275BranchNotLocked(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault error: %v", err)
	}
	if len(a.BranchProjection.BranchesInherited) != 2 || !BranchResidualOK(a.BranchProjection.BranchesInherited) {
		t.Fatalf("bad Gate275 branches: %s", FormatProjection(a.BranchProjection))
	}
	if a.BranchProjection.ResolventRootSelected || a.BranchProjection.ResolventRootToRBranchMap || a.BranchProjection.UniqueRBranchSelected {
		t.Fatalf("branch was overpromoted: %s", FormatProjection(a.BranchProjection))
	}
}

func TestFirewallsAndSummary(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault error: %v", err)
	}
	if !a.Firewall.NoObservedMassesUsed || !a.Firewall.NoArbitraryRootSectorMap || a.Firewall.FiniteCorePolluted || !a.Firewall.HiggsRatioNotClaimed {
		t.Fatalf("bad firewall: %s", FormatFirewall(a.Firewall))
	}
	if !a.Summary.UniqueSectorPairing || a.Summary.UniqueContactRoot || a.Summary.UniqueAmplitudeBranch || a.Summary.HiggsRatioDerived {
		t.Fatalf("bad summary: %s", FormatSummary(a.Summary))
	}
}

func TestTheorem(t *testing.T) {
	res := ResolventCubicSelectorBGapTauEtaSymmetryBreakingAuditTheorem().Verify()
	if len(res.Checks) == 0 {
		t.Fatalf("no checks")
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check %q failed: %s", c.Name, c.Detail)
		}
	}
}
