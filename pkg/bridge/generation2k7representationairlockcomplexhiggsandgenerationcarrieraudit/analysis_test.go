package generation2k7representationairlockcomplexhiggsandgenerationcarrieraudit

import (
	"strings"
	"testing"
)

func TestRepresentationAirlockCandidates(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.ShadowInherited || a.Inherited.K7PlusDimension != 4 || a.Inherited.K7MinusDimension != 3 || !a.Inherited.FanoCouplingFrameCandidate || a.Inherited.PhysicalHiggsMapCertified || a.Inherited.GenerationMapCertified || a.Inherited.YukawaTheoremCertified || a.Inherited.FlavorHierarchyCertified || a.Inherited.CKMPMNSTheoremCertified || a.Inherited.HiggsFlavorMapCertified || a.Inherited.SevenOver72TheoremCertified {
		t.Fatalf("bad Gate708 inheritance: %+v", a.Inherited)
	}
	if a.Higgs.RealDimension != 4 || a.Higgs.HiggsDoubletRealDimension != 4 || !a.Higgs.QuaternionicTripleVisible || !a.Higgs.CandidateComplexStructure || !a.Higgs.SU2LikeInternalAction || a.Higgs.SU2LMapCertified || a.Higgs.HyperchargeCertified || a.Higgs.PhysicalHiggsDoubletMap || !strings.Contains(a.Higgs.Verdict, StatusK7PlusHiggsRealSpaceCandidate) {
		t.Fatalf("bad K7+ Higgs airlock audit: %+v", a.Higgs)
	}
	if a.Generation.RealDimension != 3 || a.Generation.ChannelCount != 3 || !a.Generation.SO3CovariantInternalFrame || a.Generation.DiscreteFamilyLabelsCertified || a.Generation.ComplexGenerationRealDimension != 6 || a.Generation.C3GenerationMapCertified || a.Generation.FlavorHilbertFactorCertified || a.Generation.YukawaOperatorsCertified || !strings.Contains(a.Generation.Verdict, StatusK7MinusFlavorChannelCandidate) {
		t.Fatalf("bad K7- generation airlock audit: %+v", a.Generation)
	}
	if a.CouplingFrame.Rank != 3 || a.CouplingFrame.FrameSize != 12 || !a.CouplingFrame.QuaternionicFanoCalibration || !a.CouplingFrame.SO3GaugeCovariant || !a.CouplingFrame.ProtoYukawaOrientation || a.CouplingFrame.YukawaOperatorCertified || a.CouplingFrame.SingularValuesCertified || a.CouplingFrame.MixingMatricesCertified || !strings.Contains(a.CouplingFrame.Verdict, StatusFanoNormalFormCouplingFrameCandidate) {
		t.Fatalf("bad Fano coupling-frame audit: %+v", a.CouplingFrame)
	}
}

func TestComplexificationFirewallsAndTheorem(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Complexification.K7MinusRealDimension != 3 || a.Complexification.C3GenerationComplexDimension != 3 || a.Complexification.C3GenerationRealDimension != 6 || a.Complexification.Real3EqualsComplex3 || !a.Complexification.LabelSpacePossibleFuture || a.Complexification.ComplexificationCertified || !strings.Contains(a.Complexification.Verdict, StatusNoTypedK7MinusToComplexGenerationSpace) {
		t.Fatalf("bad complexification firewall: %+v", a.Complexification)
	}
	f := a.Firewalls
	if f.ClaimsK7PlusPhysicalHiggsDoublet || f.ClaimsK7MinusPhysicalGeneration || f.ClaimsOmegaYukawaMatrix || f.ClaimsFanoObservedFlavorTheorem || f.ClaimsFourPlusThreeDerivation || f.ClaimsHiggsMass || f.ClaimsYukawaEigenvalues || f.ClaimsFlavorHierarchy || f.ClaimsCKMPMNS || f.Verdict != StatusGate709RepresentationAirlockBoundary {
		t.Fatalf("physical representation firewall violated: %+v", f)
	}
	if len(a.Missing.Missing) != 5 || !strings.Contains(a.Missing.Verdict, StatusNoTypedK7PlusToSU2HiggsDoubletMap) || !strings.Contains(a.Missing.Verdict, StatusNoTypedK7MinusToComplexGenerationSpace) || !strings.Contains(a.Missing.Verdict, StatusNoTypedFanoToYukawaOperatorMap) || !strings.Contains(a.Missing.Verdict, StatusNoYukawaEigenvalueOrFlavorHierarchy) || !strings.Contains(a.Missing.Verdict, StatusNoHiggsMassOrScalarRuntimeTheorem) {
		t.Fatalf("bad missing representation maps: %+v", a.Missing)
	}
	res := Generation2K7RepresentationAirlockComplexHiggsAndGenerationCarrierAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected construction failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
