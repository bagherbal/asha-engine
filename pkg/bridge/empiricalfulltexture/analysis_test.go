package empiricalfulltexture

import "testing"

func TestGate265SealAndInheritance(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.Inheritance.EmpiricalYukawaSealActive || !a.Inheritance.RestrictedAnsatzViolated || !a.Inheritance.FullEmpiricalMatricesRequired {
		t.Fatalf("bad Gate 264 inheritance: %s", FormatInheritance(a.Inheritance))
	}
	if !a.Seal.Activated || !a.Seal.ExplicitlyQuarantined || a.Seal.DerivedFromFiniteCore || a.Seal.AllowsMassPrediction || a.Seal.AllowsCKMPrediction {
		t.Fatalf("bad full texture seal: %s", FormatSeal(a.Seal))
	}
}

func TestFullTexturesAreColumnOrthogonalSVDWitnesses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.UpSVD.ColumnOrthogonalityResidual > 1e-9 || a.DownSVD.ColumnOrthogonalityResidual > 1e-9 {
		t.Fatalf("texture columns should be orthogonal by sealed basis convention: up=%s down=%s", FormatSVD(a.UpSVD), FormatSVD(a.DownSVD))
	}
	if !a.UpSVD.Passed || !a.DownSVD.Passed {
		t.Fatalf("SVD audits failed: up=%s down=%s", FormatSVD(a.UpSVD), FormatSVD(a.DownSVD))
	}
	if a.UpSVD.ReconstructionResidual > 1e-9 || a.DownSVD.ReconstructionResidual > 1e-9 {
		t.Fatalf("bad reconstruction residuals: up=%s down=%s", FormatSVD(a.UpSVD), FormatSVD(a.DownSVD))
	}
}

func TestMassEigenvalueExtraction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Masses.Verified || !a.Masses.PhenomenologicalInputOnly {
		t.Fatalf("mass extraction must be verified but phenomenological: %s", FormatMasses(a.Masses))
	}
	if a.Masses.UpMaxAbsError > a.Masses.Tolerance || a.Masses.DownMaxAbsError > a.Masses.Tolerance {
		t.Fatalf("mass extraction errors exceed tolerance: %s", FormatMasses(a.Masses))
	}
}

func TestCKMReconstructionFromLeftMisalignment(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.CKM.Verified || a.CKM.DerivedFromFiniteCore || !a.CKM.PhenomenologicalInputOnly {
		t.Fatalf("CKM reconstruction/firewall failed: %s", FormatCKM(a.CKM))
	}
	if a.CKM.FrobeniusResidual > a.CKM.Tolerance || a.CKM.UnitarityResidual > 1e-9 {
		t.Fatalf("CKM residual too large: %s", FormatCKM(a.CKM))
	}
}

func TestFirewallNoNativeDerivation(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Firewall.DoesNotClaimFiniteMassDerivation || !a.Firewall.DoesNotClaimFiniteCKMDerivation || !a.Firewall.DoesNotInferYukawaAction || a.Firewall.FiniteCorePolluted {
		t.Fatalf("firewall failed: %s", FormatFirewall(a.Firewall))
	}
	if a.Summary.NativeDerivation || !a.Summary.EmpiricalBoundaryPreserved || !a.Summary.CKMReconstructed {
		t.Fatalf("bad summary: %s", FormatSummary(a.Summary))
	}
}
