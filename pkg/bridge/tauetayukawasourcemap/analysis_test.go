package tauetayukawasourcemap

import "testing"

func TestGate261TauEtaYukawaSourceMapAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.Inheritance.EightVRouteClosed || !a.Inheritance.DirectGenerationCarrierOpened || !a.Inheritance.TauEtaYukawaSourceCandidate {
		t.Fatalf("bad Gate 260 inheritance: %s", FormatInheritance(a.Inheritance))
	}
	if !a.BilinearCarrier.BilinearCarrierLawful || a.BilinearCarrier.TextureAlgebraDimension != 9 || a.BilinearCarrier.UsesEightVKernel || a.BilinearCarrier.YukawaAmplitudeInserted {
		t.Fatalf("bad bilinear carrier: %s", FormatBilinear(a.BilinearCarrier))
	}
	if !a.TauEtaAction.ActsWithoutEightVKernel || a.TauEtaAction.SignedDistinctEigenvalueCount != 3 || a.TauEtaAction.MagnitudeDistinctCount != 2 || a.TauEtaAction.Trace != 1 || a.TauEtaAction.Determinant != -4 {
		t.Fatalf("bad tau_eta action: %s", FormatTauEta(a.TauEtaAction))
	}
	if a.TextureAlgebra.CommutantDimension != 3 || a.TextureAlgebra.OffDiagonalComplementDimension != 6 || !a.TextureAlgebra.NonCommutingDirectionsExist || a.TextureAlgebra.CanonicalNonCommutingPartnerSelected {
		t.Fatalf("bad texture algebra decomposition: %s", FormatTextureAlgebra(a.TextureAlgebra))
	}
	if len(a.TextureAlgebra.DistinctAbsMixingGaps) != 3 || a.TextureAlgebra.DistinctAbsMixingGaps[0] != 1 || a.TextureAlgebra.DistinctAbsMixingGaps[1] != 3 || a.TextureAlgebra.DistinctAbsMixingGaps[2] != 4 {
		t.Fatalf("unexpected commutator gaps: %s", FormatTextureAlgebra(a.TextureAlgebra))
	}
	if !a.YukawaSourceMap.ActsOnBilinearCarrier || !a.YukawaSourceMap.DiagonalGenerationTextureSeed || a.YukawaSourceMap.CanProduceMixingByItself || a.YukawaSourceMap.PhysicalYukawaTextureDerived || !a.YukawaSourceMap.RequiresSecondNonCommutingOperator {
		t.Fatalf("bad Yukawa source-map audit: %s", FormatYukawa(a.YukawaSourceMap))
	}
	if a.SealLedger.EmpiricalMassDataUsed || a.SealLedger.ObservedMixingDataUsed || a.SealLedger.EmpiricalYukawaSealActivated {
		t.Fatalf("empirical data leaked: %s", FormatSealLedger(a.SealLedger))
	}
	if !a.Firewall.Gate260EightVNoGoPreserved || !a.Firewall.DoesNotInventNonCommutingPartner || a.Firewall.FiniteCorePolluted {
		t.Fatalf("firewall violation: %s", FormatFirewall(a.Firewall))
	}
}
