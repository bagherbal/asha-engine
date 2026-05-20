package generation2etarecordendhphimatrixcertificateaudit

import "testing"

func TestGate558EtaRecordEndHPhiMatrixCertificateAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.HPhi.IdentityCertified || a.HPhi.Dimension != 4 || !a.HPhi.ConditionalOnSpontaneousOrientationSeal || a.HPhi.NativeUnsealed {
		t.Fatalf("H_phi certificate failed: %s", FormatHPhi(a.HPhi))
	}
	if !a.Eta.MatrixAvailable || a.Eta.EtaSquaredResidual > 1e-9 || a.Eta.SymmetryResidual > 1e-9 || a.Eta.Rank != 4 {
		t.Fatalf("eta certificate failed: %s", FormatEta(a.Eta))
	}
	if len(a.Records) != 3 || a.Records[0].EtaTrace != 2 || a.Records[1].EtaTrace != -2 || a.Records[2].EtaTrace != 1 {
		t.Fatalf("trace records failed: %s", FormatRecords(a.Records))
	}
	if !a.Closure.Constructed || a.Closure.Dimension != 2 || !a.Closure.Commutative || !a.Closure.Semisimple || !a.Closure.UnitIdentityVerified {
		t.Fatalf("closure failed: %s", FormatClosure(a.Closure))
	}
	if !a.Split.ProjectorsFound || !a.Split.SplitTwoPlusTwo || a.Split.SplitTwoPlusOne || a.Split.IdentifiesWeakPlane || a.Split.IdentifiesFlavor || a.Split.IdentifiesHiggsRadialGoldstone {
		t.Fatalf("split/firewall failed: %s", FormatSplit(a.Split))
	}
	if !a.TraceSpectrum.ValuesAreTraces || a.TraceSpectrum.OperatorWithSpectrumSigned || a.TraceSpectrum.OperatorWithSpectrumAbs {
		t.Fatalf("trace/spectrum failed: %s", FormatTraceSpectrum(a.TraceSpectrum))
	}
	if !a.Gram.MatrixComputed || a.Gram.Rank != 2 || a.Gram.IntrinsicPositiveTwoPlusOne {
		t.Fatalf("eta-Gram failed: %s", FormatGram(a.Gram))
	}
	if a.Transfer.TransferAllowed || a.Transfer.FunctorToWSpatial || a.Transfer.FunctorToGeneration || a.Transfer.PromotedToHiggs || a.Transfer.PromotedToYukawa || a.Transfer.PromotedToCKMPMNS {
		t.Fatalf("transfer firewall failed: %s", FormatTransfer(a.Transfer))
	}
}

func TestGate558Theorem(t *testing.T) {
	res := Generation2EtaRecordEndHPhiMatrixCertificateAndProductClosureAuditTheorem().Run()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected failed route: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %+v", c)
		}
	}
}
