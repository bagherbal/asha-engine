package scalarcontactselector

import "testing"

func TestBuildDefaultScalarContactSelector(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.PreviousGate185.Firewall.QuarticAbstractOperatorDerived {
		t.Fatalf("expected inherited abstract quartic module")
	}
	if a.Partition.CanonicalPartitionDerived || !a.Partition.RequiresBranchChoice {
		t.Fatalf("expected no canonical partition and a required branch choice: %s", FormatPartition(a.Partition))
	}
	if got := a.Resolvent.MonicCoefficients; len(got) != 4 || got[1] != "-119/60" || got[2] != "8411/6480" || got[3] != "-1637467/5832000" {
		t.Fatalf("unexpected resolvent coefficients: %s", FormatResolvent(a.Resolvent))
	}
	if a.ExternalSelector.CanonicalPartitionSelectors != 0 {
		t.Fatalf("external selector should not be derived: %s", FormatExternalSelectors(a.ExternalSelector))
	}
	if a.ComplexStructure.ExistsElementSquareMinusOne || a.ComplexStructure.CanonicalComplexStructureDerived {
		t.Fatalf("commuting complex structure must be obstructed: %s", FormatComplexStructure(a.ComplexStructure))
	}
	if a.Firewall.PhysicalScalarBundleDerived || a.Firewall.PhysicalConstantsDerived {
		t.Fatalf("must not claim physical scalar bundle/constants: %s", FormatFirewall(a.Firewall))
	}
}

func TestTheoremChecksPass(t *testing.T) {
	res := ScalarContactQuarticIdentificationSelectorObstructionTheorem().Verify()
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check %q failed: %s", c.Name, c.Detail)
		}
	}
}
