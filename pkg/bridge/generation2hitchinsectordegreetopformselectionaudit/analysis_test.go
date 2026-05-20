package generation2hitchinsectordegreetopformselectionaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate650Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.ChannelAlgebraInherited || !a.Inherited.TwoComponentSupport || !a.Inherited.AAAChannelAudited || !a.Inherited.AABChannelAudited || !a.Inherited.VanishingAudited || !a.Inherited.OffBlockAudited || !a.Inherited.SlotFormulaDerived || !a.Inherited.DEqualsQCoincidence || a.Inherited.FullSymbolicChannelTheorem || a.Inherited.SplitG2Certified || a.Inherited.BoundaryStressAssignment || a.Inherited.SevenOver72Theorem || a.Inherited.ScalarFlavorTransport || a.Inherited.PhysicalMetric || !a.Inherited.Gate649FirewallPreserved {
		t.Fatalf("bad inheritance: %+v", a.Inherited)
	}
	if a.Ledger.TopDegree != (Degree{4, 3}) || a.Ledger.PositiveDim != 4 || a.Ledger.NegativeDim != 3 || !a.Ledger.AHasDegree21 || !a.Ledger.BHasDegree03 || len(a.Ledger.Rows) != 2 {
		t.Fatalf("bad degree ledger: %+v", a.Ledger)
	}
	if !a.Positive.AAAOnlySurvives || len(a.Positive.SurvivingChannels) != 1 || a.Positive.SurvivingChannels[0] != "AAA" || len(a.Positive.Rows) != 8 {
		t.Fatalf("bad positive audit: %+v", a.Positive)
	}
	for _, r := range a.Positive.Rows {
		if r.Channel == "AAA" {
			if !r.SurvivesByDegree || !r.ReachesTopDegree || r.FirstSlotZero || r.SecondSlotZero || r.TotalDegree != (Degree{4, 3}) {
				t.Fatalf("bad positive AAA row: %+v", r)
			}
		} else if r.SurvivesByDegree {
			t.Fatalf("unexpected positive survivor: %+v", r)
		}
	}
	if !a.Negative.AABPlacementsOnly || a.Negative.AllowedPlacementCount != 3 || !containsAll(a.Negative.SurvivingChannels, []string{"AAB", "ABA", "BAA"}) || len(a.Negative.Rows) != 8 {
		t.Fatalf("bad negative audit: %+v", a.Negative)
	}
	for _, r := range a.Negative.Rows {
		survivor := r.Channel == "AAB" || r.Channel == "ABA" || r.Channel == "BAA"
		if survivor {
			if !r.SurvivesByDegree || !r.ReachesTopDegree || r.TotalDegree != (Degree{4, 3}) {
				t.Fatalf("bad negative survivor: %+v", r)
			}
		} else if r.SurvivesByDegree {
			t.Fatalf("unexpected negative survivor: %+v", r)
		}
	}
	if a.Mixed.AnySurvivesByDegree || !a.Mixed.MixedBlockZeroByDegree || len(a.Mixed.Rows) != 16 {
		t.Fatalf("bad mixed audit: %+v", a.Mixed)
	}
	for _, r := range a.Mixed.Rows {
		if r.SurvivesByDegree {
			t.Fatalf("unexpected mixed survivor: %+v", r)
		}
	}
	if !a.Sign.DegreeRuleCertifiesSupport || !a.Sign.Gate649CertifiesFiniteSigns || a.Sign.EqualUnitWeightCertifiedByDegree || !a.Sign.RequiresCalibrationIdentity {
		t.Fatalf("bad sign audit: %+v", a.Sign)
	}
	if !a.Theorem.DegreeSelectionSupported || a.Theorem.FullSymbolicTheoremCertified {
		t.Fatalf("bad theorem readiness: %+v", a.Theorem)
	}
	if a.Slot.PositiveDim != 4 || a.Slot.NegativeDim != 3 || a.Slot.SlotMultiplicity != 3 || math.Abs(a.Slot.NormSquared-31) > tol || !a.Slot.RecoversGate642Angle || math.Abs(a.Slot.Cosine-13/math.Sqrt(217)) > tol || math.Abs(a.Slot.ResidualSquared-48.0/217.0) > tol {
		t.Fatalf("bad slot formula: %+v", a.Slot)
	}
	if a.Resonance.CubicDegree != 3 || a.Resonance.NegativeDim != 3 || !a.Resonance.EqualInASHACarrier {
		t.Fatalf("bad resonance: %+v", a.Resonance)
	}
	if a.Firewalls.ClaimsFullSymbolicDegreeTheorem || a.Firewalls.ClaimsSplitG2 || a.Firewalls.ClaimsBoundaryStress || a.Firewalls.ClaimsSevenOver72 || a.Firewalls.ClaimsScalarFlavor || a.Firewalls.ClaimsPhysicalMetric || a.Firewalls.ClaimsHiggsMass || a.Firewalls.ClaimsCKMPMNS || a.Firewalls.ClaimsGaugeUnification || a.Firewalls.Verdict != StatusGate650Boundary {
		t.Fatalf("firewall breach: %+v", a.Firewalls)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2HitchinSectorDegreeTopFormSelectionRuleAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected construction failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusGate649ChannelAlgebraInherited, StatusSectorDegreeLedgerDefined, StatusPositiveAAAOnlyByTopFormDegree, StatusNegativeAABByTopFormDegree, StatusMixedBlockZeroByTopFormDegree, StatusDegreeSelectionRuleSupported, StatusMinusThreeFromDegreePlacements, StatusDEqualsQCarrierResonance, StatusSignUnitRequiresCalibration, StatusNoFullSymbolicDegreeTheorem, StatusNoSplitG2, StatusNoBoundaryStress, StatusNoSevenOver72, StatusNoScalarFlavor, StatusNoPhysicalMetric, StatusNoHiggsFlavorGauge, StatusGate650Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
