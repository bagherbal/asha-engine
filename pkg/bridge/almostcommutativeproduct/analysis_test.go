package almostcommutativeproduct

import (
	"math"
	"testing"
)

func TestBuildDefaultAlmostCommutativeProduct(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.Product.ProductIsMarriage || a.Product.ProductIsDerivation || a.Product.SpacetimeDerivedFromF {
		t.Fatalf("product directionality firewall failed: %+v", a.Product)
	}
	if a.Product.FiniteHilbertDimension != 96 {
		t.Fatalf("unexpected finite Hilbert dimension: %d", a.Product.FiniteHilbertDimension)
	}
	if !a.HeatKernel.ProductFactorization || len(a.HeatKernel.Terms) < 7 {
		t.Fatalf("heat-kernel product factorization missing: %+v", a.HeatKernel)
	}
	if !a.Finite.AllRequiredInvariantsSeen || a.Finite.ChargedFiniteDiracModuli != 13 {
		t.Fatalf("bad finite invariant ledger: %+v", a.Finite)
	}
	if math.Abs(a.Finite.F2LambdaOverMP2-math.Pi/64) > 1e-15 {
		t.Fatalf("bad f2Λ/M_P invariant: %.18g", a.Finite.F2LambdaOverMP2)
	}
	if math.Abs(a.Finite.LambdaHOverGStarSquared-1197.0/4624.0) > 1e-15 {
		t.Fatalf("bad lambda ratio: %.18g", a.Finite.LambdaHOverGStarSquared)
	}
	if !a.Lagrangian.StandardModelRecovered || !a.Lagrangian.EinsteinGravityRecovered || a.Lagrangian.AllCoefficientsFixed {
		t.Fatalf("bad lagrangian assembly: %+v", a.Lagrangian)
	}
	if !a.Interface.EnablesRG || !a.Interface.EnablesBoltzmann || !a.Interface.EnablesBounce || a.Interface.HardCosmologyPredictedNow {
		t.Fatalf("bad continuum interface: %+v", a.Interface)
	}
	if !a.Firewall.DoesNotPredictCosmologicalConstant || !a.Firewall.DoesNotErase13Moduli || !a.Firewall.DoesNotDeriveMFromF {
		t.Fatalf("firewall breach: %+v", a.Firewall)
	}
}

func TestNativeProductConstants(t *testing.T) {
	c := NativeProductConstants()
	if c["dim_HF_doubled"] != 96 {
		t.Fatalf("unexpected dim_HF_doubled: %.0f", c["dim_HF_doubled"])
	}
	if c["f0_contact"] != 7 {
		t.Fatalf("unexpected f0: %.12g", c["f0_contact"])
	}
	if c["charged_DF_moduli"] != 13 || c["external_minimal_ledger"] != 15 {
		t.Fatalf("bad moduli ledger: %+v", c)
	}
	if !(c["lambdaH_over_gstar2"] > 0.25 && c["lambdaH_over_gstar2"] < 0.27) {
		t.Fatalf("unexpected quartic ratio: %.12g", c["lambdaH_over_gstar2"])
	}
}

func TestTheoremPasses(t *testing.T) {
	res := AlmostCommutativeProductGeometryFullSMGravitySpectralActionAssemblyTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem failed:\n%s", res.Details())
	}
}
