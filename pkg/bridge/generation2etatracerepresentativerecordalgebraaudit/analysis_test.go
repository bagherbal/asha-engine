package generation2etatracerepresentativerecordalgebraaudit

import "testing"

func TestGate557EtaTraceRepresentativeRecordAlgebraAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Eta.ActsAsTraceGradingOnHPhi || a.Eta.NativeEndHPhiMatrixCertified || a.Eta.SpectrumKnown || a.Eta.RankKnown {
		t.Fatalf("eta type boundary failed: %s", FormatEta(a.Eta))
	}
	if len(a.RecordAlgebra.Records) != 3 || a.RecordAlgebra.Records[0].EtaTraceValue != 2 || a.RecordAlgebra.Records[1].EtaTraceValue != -2 || a.RecordAlgebra.Records[2].EtaTraceValue != 1 {
		t.Fatalf("trace records not recovered: %s", FormatRecordAlgebra(a.RecordAlgebra))
	}
	if a.RecordAlgebra.ConstructedAsEndHPhiAlgebra || a.RecordAlgebra.ProductClosureKnown || a.RecordAlgebra.NontrivialIdempotentsKnown {
		t.Fatalf("record algebra should be blocked: %s", FormatRecordAlgebra(a.RecordAlgebra))
	}
	if a.HPhiSplit.SplitOnePlusThree || a.HPhiSplit.SplitTwoPlusTwo || a.HPhiSplit.SplitTwoPlusOnePlusOne || a.HPhiSplit.ProjectorsAvailable {
		t.Fatalf("H_phi split should not be derived: %s", FormatHPhiSplit(a.HPhiSplit))
	}
	if !a.TraceSpectrum.ValuesAreEtaTraces || a.TraceSpectrum.NativeOperatorWithSpectrum || a.TraceSpectrum.NativeOperatorWithAbsSpectrum {
		t.Fatalf("trace/spectrum boundary failed: %s", FormatTraceSpectrum(a.TraceSpectrum))
	}
	if a.EtaGram.MatrixComputed || a.EtaGram.IntrinsicTwoPlusOneSplit || a.EtaGram.ProductTracesAvailable {
		t.Fatalf("eta-Gram should be unavailable: %s", FormatEtaGram(a.EtaGram))
	}
	if a.Transfer.NativeTransferAllowed || a.Transfer.FunctorToWSpatialExists || a.Transfer.FunctorToGenerationCarrierExists {
		t.Fatalf("transfer firewall failed: %s", FormatTransfer(a.Transfer))
	}
	if a.Firewall.PollutedNativeRegistry || a.Firewall.PromotedTraceValuesToSpectrum || a.Firewall.PromotedTauEtaToHiggs || a.Firewall.PromotedTauEtaToYukawa || a.Firewall.PromotedTauEtaToCKMPMNS {
		t.Fatalf("firewall polluted: %s", FormatFirewall(a.Firewall))
	}
}

func TestGate557Theorem(t *testing.T) {
	res := Generation2EtaTraceRepresentativeAndRecordAlgebraAuditTheorem().Run()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected failed route: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %+v", c)
		}
	}
}
