package productspectralactioncoefficients

import (
	"math"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func ProductSpectralActionCoefficientCalculatorClosureAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-PRODUCT-SPECTRAL-ACTION-COEFFICIENT-CALCULATOR"
	const name = "Product Spectral Action Coefficient Calculator / Almost-Commutative Closure Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build coefficient audit", Passed: false, Detail: err.Error()}}}
		}
		c := a.Calculation
		checks := []theorem.Check{
			{Name: "Gate 376 criticism is accepted: this gate performs coefficient arithmetic", Passed: c.Executed && c.Product.Valid && c.Convention.Dimension == 4, Detail: c.Verdict},
			{Name: "product spectral triple is reconstructed", Passed: c.Product.Algebra != "" && c.Product.HilbertSpace != "" && c.Product.Dirac == "D_total = D_M ⊗ 1_F + γ₅ ⊗ D_F", Detail: c.Product.Algebra + "; " + c.Product.Dirac},
			{Name: "heat-kernel convention is declared before reading coefficients", Passed: c.Convention.IncludesRaw16Pi2 && c.Convention.Expansion != "" && c.Convention.A0Density != "" && c.Convention.A2DiracRDensity != "", Detail: c.Convention.Verdict},
			{Name: "finite ASHA invariants are substituted", Passed: c.Finite.TrOne == 96 && c.Finite.F0 == 7 && math.Abs(c.Finite.F2LambdaOverMP2-math.Pi/64) < 1e-15 && math.Abs(c.Finite.LambdaHOverGStarSq-1197.0/4624.0) < 1e-15 && c.Finite.ChargedModuli == 13, Detail: FormatCalculation(c)},
			{Name: "a0 cosmological channel is computed but not physically predicted", Passed: c.A0CosmologicalPrefactorPerF4Lambda4.Numeric > 0 && !c.A0CosmologicalPrefactorPerF4Lambda4.FullyPhysical, Detail: FormatCoefficient(c.A0CosmologicalPrefactorPerF4Lambda4)},
			{Name: "a2 Einstein-Hilbert channels are computed and expose the normalization seal", Passed: c.A2RawEinsteinCoefficientPerMP2.Numeric > 0 && c.A2SkeletonEinsteinCoefficientPerMP2.Numeric > 0 && c.A2NormalizationNeededToMatchMP > 0 && !c.A2RawEinsteinCoefficientPerMP2.FullyPhysical, Detail: FormatCoefficient(c.A2RawEinsteinCoefficientPerMP2) + "\n" + FormatCoefficient(c.A2SkeletonEinsteinCoefficientPerMP2)},
			{Name: "gauge and Higgs finite ratios are read off", Passed: c.GaugeSin2ThetaW.Numeric == 3.0/8.0 && math.Abs(c.HiggsQuarticRatio.Numeric-1197.0/4624.0) < 1e-15, Detail: FormatCoefficient(c.GaugeSin2ThetaW) + "\n" + FormatCoefficient(c.HiggsQuarticRatio)},
			{Name: "fermionic/Yukawa term is present but preserves the 13 moduli", Passed: c.YukawaTerm.Numeric == 13 && !c.YukawaTerm.FullyPhysical, Detail: FormatCoefficient(c.YukawaTerm)},
			{Name: "SM+gravity structure is recovered but full numerical ToE closure is not", Passed: c.StandardModelGravityStructural && !c.AllCoefficientsDetermined && !c.HardTOEClosure, Detail: c.Lagrangian},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{a.Truth}}
	}}
}
