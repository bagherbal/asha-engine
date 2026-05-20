package generation2hitchinchannelalgebraselectionruleaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate649Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.SlotMultiplicityInherited || a.Inherited.PositiveDim != 4 || a.Inherited.NegativeDim != 3 || a.Inherited.CubicDegree != 3 || !a.Inherited.SlotSourceSupported || !a.Inherited.ASHADimEqualsDegree || a.Inherited.GeneralPQDimensionTheorem || a.Inherited.CubicSlotTheoremCertified || a.Inherited.FullSymbolicHitchin || a.Inherited.SplitG2Certified || a.Inherited.BoundaryStressAssignment || a.Inherited.SevenOver72Theorem || a.Inherited.ScalarFlavorTransport || a.Inherited.PhysicalMetric || !a.Inherited.Gate648FirewallPreserved {
		t.Fatalf("bad inheritance: %+v", a.Inherited)
	}
	if !a.Support.OnlyAAndBSupported || len(a.Support.Rows) != 4 || a.Support.AName != "Ω++-" || a.Support.BName != "Ω---" {
		t.Fatalf("bad support audit: %+v", a.Support)
	}
	supported := map[string]bool{}
	for _, r := range a.Support.Rows {
		supported[r.Family] = r.Supported
		if r.ResidualLeak {
			t.Fatalf("unexpected residual family support: %+v", r)
		}
	}
	if !supported["Ω++-"] || !supported["Ω---"] || supported["Ω+++"] || supported["Ω+--"] {
		t.Fatalf("unexpected family support map: %+v", supported)
	}
	if a.Expansion.RouteCount != 3 || a.Expansion.ChannelsPerRoute != 8 || a.Expansion.NonzeroChannelsPerRoute != 4 || !a.Expansion.AAAOnlyPositive || !a.Expansion.AABOnlyNegative || !a.Expansion.ABBBBBClean || !a.Expansion.MixedBlocksClean || len(a.Expansion.Rows) != 24 {
		t.Fatalf("bad expansion: %+v", a.Expansion)
	}
	for _, r := range a.Expansion.Rows {
		switch r.Class {
		case "AAA":
			if !r.Nonzero || math.Abs(r.PlusMeanUnit-1) > tol || math.Abs(r.MinusMeanUnit) > tol || r.MixedFrobenius > tol {
				t.Fatalf("bad AAA row: %+v", r)
			}
		case "AAB", "ABA", "BAA":
			if !r.Nonzero || math.Abs(r.MinusMeanUnit+1) > tol || math.Abs(r.PlusMeanUnit) > tol || r.MixedFrobenius > tol {
				t.Fatalf("bad AAB row: %+v", r)
			}
		case "ABB", "BAB", "BBA", "BBB":
			if r.Nonzero || math.Abs(r.PlusMeanUnit) > tol || math.Abs(r.MinusMeanUnit) > tol || r.MixedFrobenius > tol {
				t.Fatalf("bad vanishing row: %+v", r)
			}
		default:
			t.Fatalf("unexpected class: %+v", r)
		}
	}
	if !a.PositiveAAA.AAAContributesUnit || !a.PositiveAAA.AAAContributesOnlyPlus || len(a.PositiveAAA.Rows) != 3 {
		t.Fatalf("bad positive audit: %+v", a.PositiveAAA)
	}
	if !a.NegativeAAB.EachAABContributesMinusUnit || a.NegativeAAB.NegativeOrderedChannelCount != 3 || math.Abs(a.NegativeAAB.CombinedNegativeCoefficient+3) > tol || len(a.NegativeAAB.Rows) != 9 {
		t.Fatalf("bad negative audit: %+v", a.NegativeAAB)
	}
	if !a.Vanishing.AllVanishOrProjectAway || a.Vanishing.SymbolicMechanismCertified || len(a.Vanishing.Rows) != 12 {
		t.Fatalf("bad vanishing audit: %+v", a.Vanishing)
	}
	if !a.OffBlock.ChannelwiseZero || a.OffBlock.MaxMixedFrobenius > tol {
		t.Fatalf("bad off block: %+v", a.OffBlock)
	}
	if a.SlotFormula.PositiveDim != 4 || a.SlotFormula.NegativeDim != 3 || a.SlotFormula.SlotMultiplicity != 3 || math.Abs(a.SlotFormula.NormSquared-31) > tol || !a.SlotFormula.RecoversGate642Angle || math.Abs(a.SlotFormula.Cosine-13/math.Sqrt(217)) > tol || math.Abs(a.SlotFormula.ResidualSquared-48.0/217.0) > tol {
		t.Fatalf("bad slot formula: %+v", a.SlotFormula)
	}
	if !a.Coincidence.EqualInASHACarrier || !a.Coincidence.SupportsSlotTheoremOnly || a.Coincidence.SupportsDimensionTheorem {
		t.Fatalf("bad coincidence audit: %+v", a.Coincidence)
	}
	if !a.Readiness.FiniteChannelRuleSupported || a.Readiness.FullSymbolicTheoremCertified {
		t.Fatalf("bad readiness: %+v", a.Readiness)
	}
	if a.Firewalls.ClaimsFullSymbolicChannelSelection || a.Firewalls.ClaimsSplitG2 || a.Firewalls.ClaimsBoundaryStress || a.Firewalls.ClaimsSevenOver72 || a.Firewalls.ClaimsScalarFlavor || a.Firewalls.ClaimsPhysicalMetric || a.Firewalls.ClaimsHiggsMass || a.Firewalls.ClaimsCKMPMNS || a.Firewalls.ClaimsGaugeUnification || a.Firewalls.Verdict != StatusGate649Boundary {
		t.Fatalf("firewall breach: %+v", a.Firewalls)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2HitchinChannelAlgebraSelectionRuleAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected construction failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusGate648SlotMultiplicityInherited, StatusTwoComponentTensorSupportAudited, StatusOrderedCubicExpansionComputed, StatusAAAPositiveChannelAudited, StatusAABNegativeChannelsAudited, StatusABBBBBVanishingAudited, StatusOffBlockCancellationAudited, StatusSlotFormulaDerived, StatusSlotTheoremPrimary, StatusDEqualsQCoincidence, StatusChannelSelectionSharpened, StatusNoFullSymbolicChannelSelectionTheorem, StatusNoSplitG2, StatusNoBoundaryStress, StatusNoSevenOver72, StatusNoScalarFlavor, StatusNoPhysicalMetric, StatusNoHiggsFlavorGauge, StatusGate649Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
