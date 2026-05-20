package generation2k7plusu2higgssocketandquaternioniccommutantaudit

import (
	"strings"
	"testing"
)

func TestGate711InheritanceSO4CommutantAndU2Socket(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.QuaternionicK7PlusInherited || a.Inherited.K7PlusRealDimension != 4 || !a.Inherited.HasQuaternionicTriple || !a.Inherited.HasS2ComplexFamily || !a.Inherited.C2PreCarrierAfterChoice || a.Inherited.CanonicalJHSelected || a.Inherited.PhysicalSU2LCertified || a.Inherited.HyperchargeCertified || a.Inherited.PhysicalHiggsDoubletMap {
		t.Fatalf("bad Gate710 inheritance: %+v", a.Inherited)
	}
	if a.SO4.Dimension != 6 || a.SO4.LeftSP1Dimension != 3 || a.SO4.RightSP1Dimension != 3 || !a.SO4.TripleSelectsFactor || a.SO4.PhysicalGaugeGroup || !strings.Contains(a.SO4.Verdict, StatusSO4SplitAudited) {
		t.Fatalf("bad so4 split audit: %+v", a.SO4)
	}
	if a.Commutant.Dimension != 3 || !a.Commutant.ClosesAsSU2Like || a.Commutant.PhysicalSU2LCertified || !strings.Contains(a.Commutant.Commutator, "epsilon") || !strings.Contains(a.Commutant.Verdict, StatusQuaternionicCommutantComputed) {
		t.Fatalf("bad commutant audit: %+v", a.Commutant)
	}
	if a.ChosenJH.ComplexDimension != 2 || a.ChosenJH.CanonicalSelected || !a.ChosenJH.SelectedAfterChoice || !a.ChosenJH.PotentialK7MinusSelector || !strings.Contains(a.ChosenJH.Verdict, StatusNoCanonicalJHSelected) {
		t.Fatalf("bad chosen JH audit: %+v", a.ChosenJH)
	}
	if a.U2Socket.Dimension != 4 || !a.U2Socket.SpanJHInternalU1Candidate || !a.U2Socket.CommutantInternalSU2Candidate || a.U2Socket.PhysicalElectroweakU2 || !strings.Contains(a.U2Socket.Decomposition, "Comm_so4") {
		t.Fatalf("bad u2 socket audit: %+v", a.U2Socket)
	}
}

func TestGate711SelectorFirewallsAndTheorem(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.K7Minus.K7MinusDimension != 3 || !a.K7Minus.CanSelectJH || a.K7Minus.NativeSelectorTheorem || a.K7Minus.GenerationTheorem || a.K7Minus.FlavorOrientationTheorem || !strings.Contains(a.K7Minus.Verdict, StatusRelationToK7MinusSelectorRecorded) {
		t.Fatalf("bad K7- selector audit: %+v", a.K7Minus)
	}
	f := a.Firewalls
	if f.ClaimsInternalU2PhysicalElectroweak || f.ClaimsCommutantPhysicalSU2L || f.ClaimsSpanJHPhysicalHypercharge || f.ClaimsHyperchargeNormalization || f.ClaimsTypedHiggsDoubletMap || f.ClaimsYukawaOperator || f.ClaimsYukawaEigenvalues || f.ClaimsHiggsMass || f.ClaimsScalarRuntime || f.Verdict != StatusGate711K7PlusU2HiggsSocketBoundary {
		t.Fatalf("physical electroweak firewall violated: %+v", f)
	}
	if len(a.Missing.Missing) != 6 || !strings.Contains(a.Missing.Verdict, StatusNoHyperchargeAssignmentOrNormalization) || !strings.Contains(a.Missing.Verdict, StatusNoYukawaOperatorOrEigenvalueTheorem) {
		t.Fatalf("bad missing map ledger: %+v", a.Missing)
	}
	res := Generation2K7PlusU2HiggsSocketAndQuaternionicCommutantAuditTheorem().Verify()
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
