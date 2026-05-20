package generation2k7plusquaternioniccomplexstructureandhiggsdoubletairlockaudit

import (
	"strings"
	"testing"
)

func TestK7PlusQuaternionicComplexStructureAudit(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.RepresentationAirlockInherited || a.Inherited.K7PlusRealDimension != 4 || a.Inherited.K7MinusRealDimension != 3 || !a.Inherited.HiggsRealSpaceCandidate || !a.Inherited.FanoCouplingFrameCandidate || a.Inherited.PhysicalHiggsDoubletCertified || a.Inherited.SU2LHiggsMapCertified || a.Inherited.HyperchargeCertified || a.Inherited.YukawaOperatorCertified || a.Inherited.HiggsMassCertified {
		t.Fatalf("bad Gate709 inheritance: %+v", a.Inherited)
	}
	if !a.FanoTriple.FormsDefineEndomorphisms || !a.FanoTriple.QuaternionicTriple || a.FanoTriple.JIdentityResidual != 0 || a.FanoTriple.WedgeIdentityResidual != 0 || !a.FanoTriple.GaugeCovariant || !a.FanoTriple.InheritedFromGate654 {
		t.Fatalf("bad Fano triple inheritance: %+v", a.FanoTriple)
	}
	if !a.Endomorphisms.SkewAdjointCertified || !a.Endomorphisms.SquaresToMinusIdentity || !a.Endomorphisms.QuaternionicProductCertified || a.Endomorphisms.Residual != 0 || !strings.Contains(a.Endomorphisms.Verdict, StatusK7PlusQuaternionicComplexStructureCandidate) {
		t.Fatalf("bad endomorphism audit: %+v", a.Endomorphisms)
	}
	if !a.ComplexFamily.S2Family || !a.ComplexFamily.C2AfterChoice || a.ComplexFamily.CanonicalSelected || !strings.Contains(a.ComplexFamily.Verdict, StatusNoCanonicalHiggsComplexStructure) {
		t.Fatalf("bad complex-structure family audit: %+v", a.ComplexFamily)
	}
}

func TestHiggsAirlockFirewallsAndTheorem(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.SU2LikeAction.Sp1SU2LikeAlgebra || a.SU2LikeAction.PhysicalSU2LCertified || a.SU2LikeAction.ElectroweakEmbedding || a.SU2LikeAction.HyperchargeAssignment || a.SU2LikeAction.FiniteTripleHiggsOneForm || !strings.Contains(a.SU2LikeAction.Verdict, StatusInternalSU2NotPhysicalSU2L) {
		t.Fatalf("bad SU2-like action audit: %+v", a.SU2LikeAction)
	}
	if a.HiggsCompatibility.K7PlusRealDimension != 4 || a.HiggsCompatibility.CandidateComplexDimension != 2 || !a.HiggsCompatibility.DimensionCompatible || !a.HiggsCompatibility.C2PreCarrierAfterChoice || a.HiggsCompatibility.PhysicalHiggsDoubletMap || a.HiggsCompatibility.ScalarRuntimeTheorem || !strings.Contains(a.HiggsCompatibility.Verdict, StatusNoTypedK7PlusToPhysicalHiggsDoubletMap) {
		t.Fatalf("bad Higgs compatibility audit: %+v", a.HiggsCompatibility)
	}
	if !a.FanoRelation.K7MinusIndexesTriple || !a.FanoRelation.TwoFormsOnK7Plus || !a.FanoRelation.SupportsGate709Candidate || a.FanoRelation.YukawaOperatorCertified || a.FanoRelation.YukawaEigenvaluesCertified || a.FanoRelation.FlavorHierarchyCertified {
		t.Fatalf("bad Fano relation audit: %+v", a.FanoRelation)
	}
	f := a.Firewalls
	if f.ClaimsK7PlusPhysicalHiggsDoublet || f.ClaimsCanonicalComplexStructure || f.ClaimsInternalSU2IsPhysicalSU2L || f.ClaimsHypercharge || f.ClaimsHiggsMass || f.ClaimsScalarRuntime || f.ClaimsYukawaOperator || f.ClaimsYukawaEigenvalues || f.Verdict != StatusGate710K7PlusHiggsAirlockBoundary {
		t.Fatalf("physical Higgs firewall violated: %+v", f)
	}
	if len(a.Missing.Missing) != 6 || !strings.Contains(a.Missing.Verdict, StatusNoCanonicalHiggsComplexStructure) || !strings.Contains(a.Missing.Verdict, StatusNoHyperchargeAssignment) || !strings.Contains(a.Missing.Verdict, StatusNoYukawaOperatorOrEigenvalueTheorem) {
		t.Fatalf("bad missing airlock maps: %+v", a.Missing)
	}
	res := Generation2K7PlusQuaternionicComplexStructureAndHiggsDoubletAirlockAuditTheorem().Verify()
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
