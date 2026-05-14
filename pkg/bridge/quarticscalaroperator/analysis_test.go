package quarticscalaroperator

import "testing"

func TestBuildDefaultQuarticScalarOperator(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.Companion.PolynomialIdentityZero {
		t.Fatalf("expected q4(T)=0 exactly: %s", FormatCompanion(a.Companion))
	}
	if !a.Companion.CyclicModule || a.Companion.CyclicVectorRank != 4 {
		t.Fatalf("expected cyclic rank 4")
	}
	if !a.Moments.AllMatch {
		t.Fatalf("moments did not match: %s", FormatMoments(a.Moments))
	}
	if a.Moments.TraceT != "71/30" || a.Moments.TraceT2 != "1471/900" {
		t.Fatalf("unexpected first moments: %s", FormatMoments(a.Moments))
	}
	if !a.Gate37Comparison.Gate37PairDegenerate {
		t.Fatalf("expected Gate37 pair degeneracy")
	}
	if a.Gate37Comparison.HasQuarticMinimalPolynomial || a.Gate37Comparison.PhysicalScalarBundleDerived {
		t.Fatalf("must not promote Gate37 scalar operator to quartic physical bundle")
	}
	if !a.Firewall.QuarticAbstractOperatorDerived || !a.Firewall.QuarticMomentsVerified {
		t.Fatalf("abstract quartic module should be derived")
	}
	if a.Firewall.PhysicalScalarBundleDerived || a.Firewall.PhysicalConstantsDerived {
		t.Fatalf("must not claim physical scalar bundle/constants")
	}
}

func TestTheoremChecksPass(t *testing.T) {
	res := QuarticScalarOperatorMinimalPolynomialTheorem().Verify()
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check %q failed: %s", c.Name, c.Detail)
		}
	}
}
