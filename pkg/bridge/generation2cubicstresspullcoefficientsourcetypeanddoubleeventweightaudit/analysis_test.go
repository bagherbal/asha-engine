package generation2cubicstresspullcoefficientsourcetypeanddoubleeventweightaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate731CubicCoefficientSourceType(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate730.Inherited || !a.Gate730.SevenOver36CompressedResidual || !a.Gate730.KappaEPartiallyDependent || !a.Gate730.NoNativeBoundaryMomentExpansion {
		t.Fatalf("bad Gate730 inheritance: %+v", a.Gate730)
	}
	if math.Abs(a.DoubleEvent.CubicCoefficient-7.0/36.0) > 1e-18 || math.Abs(a.DoubleEvent.DoubleK7Weight-2*a.DoubleEvent.K7EventProbability) > 1e-18 || !a.DoubleEvent.IdentityExact {
		t.Fatalf("bad double event coefficient: %+v", a.DoubleEvent)
	}
	if math.Abs(a.BoundaryPair.BoundaryPairTimesK7Weight-a.DoubleEvent.CubicCoefficient) > 1e-18 || !a.BoundaryPair.EqualsCubicCoefficient {
		t.Fatalf("bad boundary pair source: %+v", a.BoundaryPair)
	}
	if !a.StressPull.TwoSidedBoundaryLegs || !a.StressPull.ArbitraryFitRejected {
		t.Fatalf("bad stress-pull source: %+v", a.StressPull)
	}
	if !a.KineticFactor.FactorTwoResonance || a.KineticFactor.DerivesCubicCoeff {
		t.Fatalf("bad kinetic warning: %+v", a.KineticFactor)
	}
}

func TestGate731AlternativesPolynomialAndFirewalls(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Alternatives.NoArbitrarySearch || !a.Alternatives.ClosestAccepted || a.Alternatives.ClosestName != "2p_K7=7/36" {
		t.Fatalf("bad alternatives: %+v", a.Alternatives)
	}
	if !a.Polynomial.UsesDoubleEventForm || math.Abs(a.Polynomial.CubicTerm-a.DoubleEvent.DoubleK7Weight*a.Gate730.M3Wall) > 1e-18 || !strings.Contains(a.Polynomial.PolynomialInS, "2p_K7^2") {
		t.Fatalf("bad polynomial: %+v", a.Polynomial)
	}
	if !a.NonCircular.KappaEPartiallyDependent || !a.NonCircular.TwoPK7TypedButUnexplained || a.NonCircular.BoundaryPairStressPullNative || a.NonCircular.MomentExpansionTheoremNative {
		t.Fatalf("bad noncircularity: %+v", a.NonCircular)
	}
	if a.Firewall.ClaimsNativeScalarRuntime || a.Firewall.ClaimsHiggsMassTheorem || a.Firewall.ClaimsYukawaTheorem || a.Firewall.ClaimsCKMPMNSTheorem || a.Firewall.ClaimsHistoryLoopTheorem {
		t.Fatalf("firewall failed: %+v", a.Firewall)
	}

	res := Generation2CubicStressPullCoefficientSourceTypeAndDoubleEventWeightAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected construction failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
