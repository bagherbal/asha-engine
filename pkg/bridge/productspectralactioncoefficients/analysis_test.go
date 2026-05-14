package productspectralactioncoefficients

import (
	"math"
	"testing"
)

func TestBuildDefaultCoefficientCalculator(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	c := a.Calculation
	if !c.Executed || !c.Product.Valid || !c.StandardModelGravityStructural {
		t.Fatalf("bad basic calculation: %+v", c)
	}
	if c.AllCoefficientsDetermined || c.HardTOEClosure {
		t.Fatalf("firewall breach: %+v", c)
	}
	if math.Abs(c.A0CosmologicalPrefactorPerF4Lambda4.Numeric-(96.0/(16.0*math.Pi*math.Pi))) > 1e-15 {
		t.Fatalf("bad a0 prefactor: %.18g", c.A0CosmologicalPrefactorPerF4Lambda4.Numeric)
	}
	wantRaw := 96.0 * (math.Pi / 64.0) / (192.0 * math.Pi * math.Pi)
	if math.Abs(c.A2RawEinsteinCoefficientPerMP2.Numeric-wantRaw) > 1e-15 {
		t.Fatalf("bad raw EH coefficient: %.18g want %.18g", c.A2RawEinsteinCoefficientPerMP2.Numeric, wantRaw)
	}
	wantSkeleton := 0.5 * 96.0 * (math.Pi / 64.0)
	if math.Abs(c.A2SkeletonEinsteinCoefficientPerMP2.Numeric-wantSkeleton) > 1e-15 {
		t.Fatalf("bad skeleton EH coefficient: %.18g want %.18g", c.A2SkeletonEinsteinCoefficientPerMP2.Numeric, wantSkeleton)
	}
	if math.Abs(c.A2NormalizationNeededToMatchMP-(0.5/wantSkeleton)) > 1e-15 {
		t.Fatalf("bad EH normalization: %.18g", c.A2NormalizationNeededToMatchMP)
	}
}

func TestNativeCoefficientConstants(t *testing.T) {
	m := NativeCoefficientConstants()
	if m["Tr_F_1"] != 96 || m["f0"] != 7 || m["charged_moduli"] != 13 {
		t.Fatalf("bad constants: %+v", m)
	}
	if math.Abs(m["lambdaH_over_gstar2"]-1197.0/4624.0) > 1e-15 {
		t.Fatalf("bad lambda ratio: %.18g", m["lambdaH_over_gstar2"])
	}
	if !(m["a2_skeleton_EH_coeff_per_MP2"] > 2.3 && m["a2_skeleton_EH_coeff_per_MP2"] < 2.4) {
		t.Fatalf("unexpected skeleton EH coefficient: %.18g", m["a2_skeleton_EH_coeff_per_MP2"])
	}
}

func TestTheoremPasses(t *testing.T) {
	res := ProductSpectralActionCoefficientCalculatorClosureAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem failed:\n%s", res.Details())
	}
}
