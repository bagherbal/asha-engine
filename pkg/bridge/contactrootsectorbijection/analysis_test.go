package contactrootsectorbijection

import "testing"

func TestRootsAndMagnitudeFirewall(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault error: %v", err)
	}
	if !RootResidualOK(a) || !a.Magnitude.AllRootsO1 || a.Magnitude.AnyNativeNullRoot || a.Magnitude.MagnitudeBijectionDerived || a.Magnitude.BGapScaleComparableToRoots {
		t.Fatalf("bad magnitude audit: %s", FormatMagnitude(a.Magnitude))
	}
}

func TestConstraintsReachSectorsNotRoots(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault error: %v", err)
	}
	if !a.Constraints.MoritaMultiplicityAvailable || !a.Constraints.BGapTagAvailable || !a.Constraints.TauEtaPairingAvailable || !a.Constraints.ConstraintsReachSectors || a.Constraints.ConstraintsReachRoots || a.Constraints.UsesObservedMasses {
		t.Fatalf("bad constraints: %s", FormatConstraints(a.Constraints))
	}
}

func TestProjectorSemanticsMissing(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault error: %v", err)
	}
	if !a.Projectors.QuarticIrreducibleOverQ || !a.Projectors.ResolventIrreducibleOverQ || a.Projectors.IndividualRootProjectorsOverQ || a.Projectors.TwoPlusTwoPairProjectorsOverQ || !a.Projectors.RequiresSplittingField || !a.Projectors.RequiresResolventRootAdjunction || a.Projectors.RationalContactProjectorDerived {
		t.Fatalf("bad projector audit: %s", FormatProjectors(a.Projectors))
	}
}

func TestPairingsAndBijectionRemainDegenerate(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault error: %v", err)
	}
	if a.Pairings.TotalCandidates != 3 || a.Pairings.CompatibleWithSectorSplit != 3 || a.Pairings.SelectedPairings != 0 || a.Pairings.UniqueRootPairing {
		t.Fatalf("bad pairing audit: %s", FormatPairings(a.Pairings))
	}
	if a.Bijection.TotalRootSectorBijections != 24 || a.Bijection.BijectionsAfterSectorPairing <= 1 || a.Bijection.UniqueBijection || len(a.Bijection.DerivedAssignment) != 0 {
		t.Fatalf("bad bijection audit: %s", FormatBijection(a.Bijection))
	}
}

func TestBranchAndFirewall(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault error: %v", err)
	}
	if a.BranchProjection.RPlus <= 0 || a.BranchProjection.RMinus <= 0 || a.BranchProjection.ResolventRootSelected || a.BranchProjection.RootPairingToRBranchMap || a.BranchProjection.UniqueAmplitudeBranch {
		t.Fatalf("bad branch projection: %s", FormatBranchProjection(a.BranchProjection))
	}
	if !a.Firewall.NoObservedMassesUsed || !a.Firewall.NoRootOrderingPromotion || !a.Firewall.NoArbitraryRootSectorMap || a.Firewall.FiniteCorePolluted || !a.Firewall.NoHiggsRatioClaimed {
		t.Fatalf("bad firewall: %s", FormatFirewall(a.Firewall))
	}
}

func TestTheorem(t *testing.T) {
	res := QuarticRootToYukawaSectorBijectionContactProjectorSemanticsAuditTheorem().Verify()
	if len(res.Checks) == 0 {
		t.Fatalf("no checks")
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check %q failed: %s", c.Name, c.Detail)
		}
	}
}
