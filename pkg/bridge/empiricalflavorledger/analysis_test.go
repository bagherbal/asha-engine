package empiricalflavorledger

import "testing"

func TestGate266SealAndGate265Inheritance(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.Inheritance.EmpiricalYukawaSealActive || !a.Inheritance.QuarkSVDCKMVerified || a.Inheritance.QuarkNativeDerivation || !a.Inheritance.QuarkBoundaryPreserved {
		t.Fatalf("bad Gate 265 inheritance: %s", FormatInheritance(a.Inheritance))
	}
	if !a.Seal.Activated || !a.Seal.ExplicitlyQuarantined || a.Seal.DerivedFromFiniteCore || a.Seal.AllowsMassPrediction || a.Seal.AllowsPMNSPrediction || a.Seal.AllowsNeutrinoNaturePrediction {
		t.Fatalf("bad lepton flavor seal: %s", FormatSeal(a.Seal))
	}
}

func TestChargedLeptonSVD(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.ChargedSVD.Passed || a.ChargedSVD.ReconstructionResidual > 1e-12 || a.ChargedSVD.LeftUnitarityResidual > 1e-12 {
		t.Fatalf("charged SVD failed: %s", FormatSVD(a.ChargedSVD))
	}
	if !a.Masses.Verified || a.Masses.ChargedLeptonMaxAbsError > a.Masses.ChargedTolerance {
		t.Fatalf("charged masses not verified: %s", FormatMasses(a.Masses))
	}
}

func TestMajoranaTakagiAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.NeutrinoTakagi.Passed || a.NeutrinoTakagi.ReconstructionResidual > 1e-11 || a.NeutrinoTakagi.SymmetryResidual > 1e-11 || a.NeutrinoTakagi.OffDiagonalResidual > 1e-11 {
		t.Fatalf("Takagi audit failed: %s", FormatTakagi(a.NeutrinoTakagi))
	}
	if !a.NeutrinoTakagi.MajoranaAssumptionSealed || a.NeutrinoTakagi.DerivedNeutrinoNature {
		t.Fatalf("neutrino nature firewall failed: %s", FormatTakagi(a.NeutrinoTakagi))
	}
}

func TestPMNSReconstructionAndLargeAngles(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.PMNS.Verified || a.PMNS.DerivedFromFiniteCore || !a.PMNS.PhenomenologicalInputOnly || a.PMNS.FrobeniusResidual > a.PMNS.Tolerance {
		t.Fatalf("PMNS reconstruction failed: %s", FormatPMNS(a.PMNS))
	}
	if !a.LargeAngles.LargeAngleStructure || !a.LargeAngles.S12Large || !a.LargeAngles.S23Large || !a.LargeAngles.S13Nonzero {
		t.Fatalf("large-angle audit failed: %s", FormatLargeAngles(a.LargeAngles))
	}
}

func TestLeptonFlavorFirewallNoNativeDerivation(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Firewall.DoesNotClaimFiniteChargedLeptonMass || !a.Firewall.DoesNotClaimFiniteNeutrinoMass || !a.Firewall.DoesNotClaimFinitePMNSDerivation || !a.Firewall.DoesNotClaimFiniteMajoranaDerivation || a.Firewall.FiniteCorePolluted {
		t.Fatalf("firewall failed: %s", FormatFirewall(a.Firewall))
	}
	if a.Summary.NativeDerivation || !a.Summary.EmpiricalBoundaryPreserved || !a.Summary.PMNSReconstructed || !a.Summary.NeutrinoTakagiCompleted {
		t.Fatalf("bad summary: %s", FormatSummary(a.Summary))
	}
}
