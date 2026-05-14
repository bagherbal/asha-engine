package scalaryukawasupport

import "testing"

func TestBuildDefaultTensorLiftedYukawaSupport(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.TensorLift.DoublyGradedFunctionalConstructed || a.TensorLift.TotalTensorDimension != 64 {
		t.Fatalf("bad tensor lift: %s", FormatTensorLift(a.TensorLift))
	}
	if len(a.ScalarBranches) != 2 || !a.ScalarBranches[0].NonzeroSupport || !a.ScalarBranches[1].NonzeroSupport {
		t.Fatalf("bad scalar branches: %s", FormatBranches(a.ScalarBranches))
	}
	if a.BilinearSupport.SupportedChannels != 8 || !a.BilinearSupport.AllSupportNonzero {
		t.Fatalf("bad bilinear support: %s", FormatBilinear(a.BilinearSupport))
	}
	if !a.Neutrality.TotalEtaSupportBalances || !a.Neutrality.UpDownQuarkBalance || !a.Neutrality.NeutrinoElectronBalance || !a.Neutrality.BLWeightedEtaSupportBalances {
		t.Fatalf("eta support should balance: %s", FormatNeutrality(a.Neutrality))
	}
	if a.Firewall.PhysicalYukawaAmplitudesDerived || a.Firewall.FermionMassesDerived || a.Firewall.GenerationTextureValuesDerived || a.Firewall.CKMMatrixDerived || a.Firewall.PMNSMatrixDerived || a.Firewall.PhysicalConstantsDerived {
		t.Fatalf("firewall leaked: %s", FormatFirewall(a.Firewall))
	}
}

func TestSupportSignPattern(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	positive := 0
	negative := 0
	for _, r := range a.BilinearSupport.Records {
		if r.TensorSupportSignature > 0 {
			positive++
		}
		if r.TensorSupportSignature < 0 {
			negative++
		}
		if r.AmplitudeDerived || r.MassDerived {
			t.Fatalf("channel should be support-only: %+v", r)
		}
	}
	if positive != 4 || negative != 4 {
		t.Fatalf("expected four positive and four negative eta-supported channels, got +%d -%d", positive, negative)
	}
}

func TestTheoremChecksPass(t *testing.T) {
	res := TensorLiftedScalarFundamentalClassYukawaBilinearSupportTheorem().Verify()
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check %q failed: %s", c.Name, c.Detail)
		}
	}
}
