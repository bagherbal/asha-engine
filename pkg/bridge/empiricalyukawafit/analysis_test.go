package empiricalyukawafit

import "testing"

func TestGate264EmpiricalYukawaSealActivation(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.Inheritance.GeometricAnsatzAvailable || !a.Inheritance.BasisOrthogonal || a.Inheritance.FiniteActionCoefficientRule {
		t.Fatalf("bad Gate 263 inheritance: %s", FormatInheritance(a.Inheritance))
	}
	if !a.Seal.Activated || !a.Seal.ExplicitlyQuarantined || a.Seal.DerivedFromFiniteCore || a.Seal.AllowsFinitePrediction {
		t.Fatalf("bad seal activation: %s", FormatSeal(a.Seal))
	}
	if !a.Firewall.DoesNotRewriteFiniteCore || a.Firewall.FiniteCorePolluted {
		t.Fatalf("firewall pollution: %s", FormatFirewall(a.Firewall))
	}
}

func TestRepresentativeQuarkDataLedger(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Data.RepresentativeNotPrecision || !a.Data.MixedScaleWarning || !a.Data.UsesObservedMassHierarchy || !a.Data.UsesObservedCKMParameters {
		t.Fatalf("data must be sealed representative stress data: %s", FormatData(a.Data))
	}
	if a.Data.DataParameterCount != 10 || a.Data.AnsatzQuarkParameterCount != 6 || a.Data.ParameterDeficit != 4 {
		t.Fatalf("bad parameter count: %s", FormatData(a.Data))
	}
}

func TestProjectionFitViolatesRestrictedShell(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.Fits) != 2 {
		t.Fatalf("expected two sector fits, got %d", len(a.Fits))
	}
	for _, f := range a.Fits {
		if f.FitsExactly || f.RelativeResidual <= f.ExactFitTolerance {
			t.Fatalf("sector should not fit exactly: %s", FormatFit(f))
		}
		if f.TargetFrobeniusNorm <= 0 || f.ResidualFrobeniusNorm <= 0 {
			t.Fatalf("bad fit norms: %s", FormatFit(f))
		}
	}
	if a.Fits[0].RelativeResidual < 0.5 || a.Fits[1].RelativeResidual < 0.5 {
		t.Fatalf("representative stress residuals should be large: %s", FormatFits(a.Fits))
	}
	if !a.Fits[0].EqualOffDiagonalFailure || !a.Fits[1].DiagonalShapeFailure {
		t.Fatalf("expected up off-diagonal and down diagonal shape failures: %s", FormatFits(a.Fits))
	}
}

func TestViabilityAndFirewallRemainSealed(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Viability.ViolatesAnsatz || !a.Viability.RequiresFullYukawaMatrices || a.Viability.AllSectorsExactFit {
		t.Fatalf("bad viability verdict: %s", FormatViability(a.Viability))
	}
	if a.Viability.CKMNumericalFitDerived || a.Viability.MassSpectrumDerived || a.Summary.MassesDerived || a.Summary.CKMDerived {
		t.Fatalf("must not derive masses or CKM: %s / %s", FormatViability(a.Viability), FormatSummary(a.Summary))
	}
	if !a.Firewall.FullEmpiricalSealStillRequired || !a.Firewall.Gate263NoGoPreserved || !a.Firewall.DoesNotPromoteProjectionToLaw {
		t.Fatalf("firewall failed: %s", FormatFirewall(a.Firewall))
	}
}
