package generation2tauetacarrierpullbackobstructionaudit

import "testing"

func TestGate556TauEtaCarrierPullbackObstructionAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Type.IsTraceValueVector || !a.Type.IsSealedBookkeepingDatum || a.Type.IsSpectrumOfNativeOperator || a.Type.IsDiagonalEndomorphism {
		t.Fatalf("tau_eta type classification failed: %s", FormatType(a.Type))
	}
	if a.SourceAlgebra.NativeSourceAlgebraExists || len(a.SourceAlgebra.Candidates) != 2 {
		t.Fatalf("source algebra audit failed: %s", FormatSource(a.SourceAlgebra))
	}
	if a.Representation.AnyValidUnitPreservingRepresentation || a.Representation.RhoOneIsIdentity {
		t.Fatalf("representation firewall failed: %s", FormatRepresentation(a.Representation))
	}
	if a.Selector.FormalCommutantDimension != 5 || a.Selector.ProducesNativeSelector || a.Selector.CanonicalU12Selected || !a.Selector.BasisDependentIfForced {
		t.Fatalf("selector consequence audit failed: %s", FormatSelector(a.Selector))
	}
	if a.SpectralTriple.NativeSpectralTriplePromotionAllowed || len(a.SpectralTriple.MissingData) == 0 {
		t.Fatalf("spectral triple firewall failed: %s", FormatSpectralTriple(a.SpectralTriple))
	}
	if !a.Final.TauEtaOnlyTraceVector || a.Final.TauEtaOperator || a.Final.NativeSourceAlgebraExists || a.Final.UnitPreservingRepresentationExists || a.Final.CanonicalTwoPlusOneSelectorOnWSpatial {
		t.Fatalf("bad final verdict: %s", FormatFinal(a.Final))
	}
}

func TestGate556Theorem(t *testing.T) {
	res := Generation2TauEtaCarrierPullbackObstructionAuditTheorem().Run()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected failed route: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %+v", c)
		}
	}
}
