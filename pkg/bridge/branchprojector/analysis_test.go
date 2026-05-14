package branchprojector

import "testing"

func TestBuildDefaultBranchProjector(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.Factors.FactorizationVerified || !a.Factors.FactorsMonicQuadratic || !a.Factors.OnlyTwoPlusTwoSplitConstructed {
		t.Fatalf("expected exact 2+2 quadratic factorization: %s", FormatFactors(a.Factors))
	}
	if !a.Bezout.IdentityVerified || !a.Bezout.ExactArithmetic || !a.Bezout.NoNumericRootApproximation {
		t.Fatalf("expected exact Bezout identity: %s", FormatBezout(a.Bezout))
	}
	if !a.Projectors.ProjectorPairConstructed || !a.Projectors.ProjectorSumIdentity || !a.Projectors.ProjectorsIdempotent || !a.Projectors.ProjectorsOrthogonal {
		t.Fatalf("expected complementary orthogonal idempotents: %s", FormatProjectors(a.Projectors))
	}
	if !a.Projectors.TraceTwoEach || !a.Projectors.DimensionTwoEach || a.Projectors.IndividualRootProjectors != 0 {
		t.Fatalf("expected two trace-2 projectors and no individual root projectors: %s", FormatProjectors(a.Projectors))
	}
	if a.HiggsCarrier.PhysicalScalarBundleDerived || a.Firewall.PhysicalConstantsDerived || a.Firewall.ChernWeilCarrierDerived {
		t.Fatalf("firewall violation: %s", FormatFirewall(a.Firewall))
	}
}

func TestProjectorCertificateExactAlgebra(t *testing.T) {
	c, err := buildProjectorCertificate()
	if err != nil {
		t.Fatalf("certificate failed: %v", err)
	}
	if !c.factorizationVerified {
		t.Fatalf("factorization not verified")
	}
	if !c.bezoutVerified {
		t.Fatalf("Bezout not verified")
	}
	if !c.projectorAIdempotent || !c.projectorBIdempotent || !c.projectorSumIdentity || !c.projectorsOrthogonal {
		t.Fatalf("projector laws failed: PA=%s PB=%s", c.projectorA.String(), c.projectorB.String())
	}
	if !c.traceA.Equal(intE(2)) || !c.traceB.Equal(intE(2)) {
		t.Fatalf("expected projector traces 2,2; got %s,%s", c.traceA.String(), c.traceB.String())
	}
	if !factorSwapPreserved(c) {
		t.Fatalf("eta involution should exchange the two factors/projectors")
	}
}

func TestTheoremChecksPass(t *testing.T) {
	res := BranchwiseQuadraticIdempotentScalarProjectorTheorem().Verify()
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check %q failed: %s", c.Name, c.Detail)
		}
	}
}
