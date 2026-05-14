package resolventvacuum

import "testing"

func TestBuildDefaultResolventVacuum(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.Resolvent.IrreducibleOverQ || a.Resolvent.CanonicalRootSelected {
		t.Fatalf("expected irreducible branch-free resolvent algebra with no selected root: %s", FormatResolvent(a.Resolvent))
	}
	if got := a.Resolvent.MonicCoefficients; len(got) != 4 || got[1] != "-119/60" || got[2] != "8411/6480" || got[3] != "-1637467/5832000" {
		t.Fatalf("unexpected resolvent coefficients: %s", FormatResolvent(a.Resolvent))
	}
	if !a.VacuumOrbit.DegenerateVacuumOrbitDerived || a.VacuumOrbit.CanonicalUniqueVacuumDerived || !a.VacuumOrbit.BranchSelectionIsSpontaneousData {
		t.Fatalf("expected degenerate spontaneous vacuum orbit, not a unique selector: %s", FormatVacuumOrbit(a.VacuumOrbit))
	}
	if !a.Splitting.EveryBranchGivesTwoQuadraticFactors || !a.Splitting.OrderedQuadraticFactorsRequireFurtherAdjunction || a.Splitting.IndividualRootDiagonalizationUsed {
		t.Fatalf("expected conditional two-quadratic branch scheme without root diagonalization: %s", FormatSplitting(a.Splitting))
	}
	if !a.Higgs.ConditionalScalarCarrierOpened || a.Higgs.PhysicalScalarBundleDerived || a.Higgs.CanonicalScalarProjectorDerived {
		t.Fatalf("expected conditional carrier only, not physical scalar bundle/projector: %s", FormatHiggs(a.Higgs))
	}
	if !a.Complex.Gate186GlobalCommutingJObstructed || a.Complex.CanonicalComplexStructureDerived || a.Complex.CanonicalSymplecticStructureDerived {
		t.Fatalf("complex/symplectic structure must remain noncanonical: %s", FormatComplex(a.Complex))
	}
	if a.Firewall.PhysicalConstantsDerived || a.Firewall.AbsoluteCouplingPromoted || a.Firewall.ChernWeilCarrierDerived {
		t.Fatalf("firewall violation: %s", FormatFirewall(a.Firewall))
	}
}

func TestResolventRootSemantics(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if a.Splitting.ResolventRootSemantic == "" {
		t.Fatalf("missing resolvent semantic")
	}
	for _, branch := range a.Splitting.Branches {
		if branch.ResolventMeaning != "z = p + q, where p and q are the two pair-products for this partition" {
			t.Fatalf("unexpected branch semantic for %s: %#v", branch.Label, branch)
		}
		if !branch.TwoPlusTwoScalarShape || branch.IndividualRootDiagonalized || branch.CanonicalBranchSelectedHere {
			t.Fatalf("invalid branch factor scheme for %s: %#v", branch.Label, branch)
		}
	}
}

func TestTheoremChecksPass(t *testing.T) {
	res := ResolventVacuumSpontaneousHiggsPairingTheorem().Verify()
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check %q failed: %s", c.Name, c.Detail)
		}
	}
}
