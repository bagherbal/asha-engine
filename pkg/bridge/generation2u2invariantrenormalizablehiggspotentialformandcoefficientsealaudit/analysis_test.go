package generation2u2invariantrenormalizablehiggspotentialformandcoefficientsealaudit

import (
	"strings"
	"testing"
)

func TestGate769U2InvariantReductionAndPolynomialForm(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate768.Inherited || !strings.Contains(a.Gate768.RadialReplacement, "P_rad := supp") || !a.Gate768.PotentialStillSupplied || a.Gate768.NativePotentialTheorem || a.Gate768.NativeHistoryLoopTheorem {
		t.Fatalf("bad Gate768 inheritance: %+v", a.Gate768)
	}
	if a.Carrier.ComplexDimension != complexHiggsDim || a.Carrier.RealDimension != realHiggsDim || !strings.Contains(a.Carrier.Carrier, "C^2") || !strings.Contains(a.Carrier.RepresentationSocket, "U(2)") || a.Carrier.NativeSelectorTheorem {
		t.Fatalf("bad carrier audit: %+v", a.Carrier)
	}
	if !a.Reduction.TransitiveOnFixedRadiusSpheres || a.Reduction.InvariantCoordinate != "r^2=phi^dagger phi" || a.Reduction.PotentialFunctionForm != "V(phi)=f(phi^dagger phi)" || !a.Reduction.DependsOnlyOnPhiDaggerPhi || a.Reduction.SelectsCP1Point || !a.Reduction.RequiresNoAnisotropicHermitianAxis {
		t.Fatalf("bad U2 reduction: %+v", a.Reduction)
	}
	if !a.Polynomial.RealPolynomialPremise || !a.Polynomial.QuarticTruncation || a.Polynomial.QuarticDegreeInRealFields != quarticDegree || a.Polynomial.MaxPowerInCoordinate != maxFPower || !strings.Contains(a.Polynomial.PotentialForm, "mu^2 phi^dagger phi") || !strings.Contains(a.Polynomial.PotentialForm, "lambda(phi^dagger phi)^2") || !a.Polynomial.UniqueUnderPremises || a.Polynomial.NativeSpectralActionTheorem || a.Polynomial.NativeScalarPotentialTheorem {
		t.Fatalf("bad polynomial form: %+v", a.Polynomial)
	}
}

func TestGate769CoefficientSealsAndCP1Flatness(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Constant.ConstantSymbol != "c_0" || a.Constant.AffectsGradient || a.Constant.AffectsHessian || a.Constant.AffectsRadialEvent || !a.Constant.IgnoredForLocalScalarDynamics || a.Constant.CosmologicalConstantTheorem {
		t.Fatalf("bad constant offset separation: %+v", a.Constant)
	}
	if !strings.Contains(a.Coefficients.MuSquaredRole, "quadratic") || a.Coefficients.MuSquaredDerived || a.Coefficients.MuSquaredSignDerived || !strings.Contains(a.Coefficients.LambdaRole, "quartic") || a.Coefficients.LambdaDerived || !a.Coefficients.LambdaRuntimeBridgeMaySupply || a.Coefficients.RuntimeLambdaIndependent || a.Coefficients.C0CosmologicalTheorem {
		t.Fatalf("bad coefficient seal audit: %+v", a.Coefficients)
	}
	if !a.CP1.FlatAtFixedRadius || a.CP1.CP1SelectedByPotential || !a.CP1.RadialDirectionNonzero || a.CP1.AngularDirectionsFlat != 3 || !a.CP1.PreservesGate764765 {
		t.Fatalf("bad CP1 flatness audit: %+v", a.CP1)
	}
}

func TestGate769HessianCompatibilityAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Hessian.RealCoordinateConvention != "phi^dagger phi=(1/2)||x||^2" || !strings.Contains(a.Hessian.RealPotentialForm, "lambda/4") || a.Hessian.Gate766HessianFormula != "H_V(x_0)=2 lambda v^2 P_rad" || !strings.Contains(a.Hessian.Gate768SupportReplacement, "supp") || !a.Hessian.HessianNormalizationBelongsHere || a.Hessian.NativeVEVTheorem || a.Hessian.PoleMassTheorem {
		t.Fatalf("bad Hessian compatibility: %+v", a.Hessian)
	}
	if !a.Firewalls.Audited || a.Firewalls.NativeScalarPotentialTheorem || a.Firewalls.NativeMuSquaredTheorem || a.Firewalls.NativeQuarticCoefficientTheorem || a.Firewalls.NativeVEVTheorem || a.Firewalls.NativeSpectralActionTruncationTheorem || a.Firewalls.C0CosmologicalConstantTheorem || a.Firewalls.NativeHistoryLoopUnitTheorem || a.Firewalls.TreeProxyPoleMassTheorem || a.Firewalls.HiggsMassOrPoleMassTheorem || a.Firewalls.YukawaOperatorOrEigenvalueTheorem {
		t.Fatalf("bad firewalls: %+v", a.Firewalls)
	}
}

func TestGate769TheoremStatuses(t *testing.T) {
	res := Generation2U2InvariantRenormalizableHiggsPotentialFormAndCoefficientSealAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status note %s", want)
		}
	}
}
