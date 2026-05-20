package generation2cubicslotmultiplicityversusnegativesectordimensionaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate648Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.ContractionLedgerInherited || a.Inherited.RouteCount != 3 || a.Inherited.PositiveDim != 4 || a.Inherited.NegativeDim != 3 || a.Inherited.CubicDegree != 3 || !a.Inherited.HasThreeNegativeChannels || a.Inherited.GeneralPQDimensionClaim || a.Inherited.FullSymbolicTheorem || a.Inherited.SplitG2Certified || a.Inherited.BoundaryStressAssignment || a.Inherited.SevenOver72Theorem || a.Inherited.ScalarFlavorTransport || a.Inherited.PhysicalMetric || !a.Inherited.Gate647FirewallPreserved {
		t.Fatalf("bad inheritance: %+v", a.Inherited)
	}
	if !a.TraceAudit.AllRoutesPass || len(a.TraceAudit.Rows) != 3 {
		t.Fatalf("bad trace audit: %+v", a.TraceAudit)
	}
	for _, r := range a.TraceAudit.Rows {
		if !r.Passed || math.Abs(r.PerDirectionRatio+3) > tol || math.Abs(r.TotalTraceRatio+2.25) > tol || math.Abs(r.ExpectedTraceRatio+2.25) > tol {
			t.Fatalf("bad trace row: %+v", r)
		}
	}
	if !a.SlotAudit.SlotSourceSupported || !a.SlotAudit.EachNegativeChannelUnit || a.SlotAudit.NegativeChannelCount != 3 || a.SlotAudit.ExpectedCubicSlotCount != 3 || len(a.SlotAudit.NegativeChannels) != 9 || len(a.SlotAudit.PositiveChannels) != 3 {
		t.Fatalf("bad slot audit: %+v", a.SlotAudit)
	}
	for _, r := range a.SlotAudit.NegativeChannels {
		if !r.ContributesUnitNegative || math.Abs(r.NegativeMeanUnit+1) > tol || math.Abs(r.PositiveMeanUnit) > tol {
			t.Fatalf("bad negative channel: %+v", r)
		}
	}
	for _, r := range a.SlotAudit.PositiveChannels {
		if math.Abs(r.PositiveMeanUnit-1) > tol || math.Abs(r.NegativeMeanUnit) > tol {
			t.Fatalf("bad positive channel: %+v", r)
		}
	}
	if !a.NegativeIndexAudit.AllRoutesUniformByIndex || len(a.NegativeIndexAudit.Rows) != 3 {
		t.Fatalf("bad negative index audit: %+v", a.NegativeIndexAudit)
	}
	for _, r := range a.NegativeIndexAudit.Rows {
		if r.NegativeDirections != 3 || math.Abs(r.PerNegativeDirectionWeight+3) > tol || math.Abs(r.PerChannelPerDirectionWeight+1) > tol || math.Abs(r.TotalNegativeTraceWeight+9) > tol || !r.DimensionChangesTraceOnly {
			t.Fatalf("bad negative index row: %+v", r)
		}
	}
	if !a.FormulaAudit.FinalRayCannotDistinguish || !a.FormulaAudit.LedgerSelectsSlotSource || a.FormulaAudit.GeneralPQDimensionTheorem || !a.FormulaAudit.Row.CoincideInASHA || math.Abs(a.FormulaAudit.Row.DimensionNormSq-31) > tol || math.Abs(a.FormulaAudit.Row.SlotNormSq-31) > tol {
		t.Fatalf("bad formula audit: %+v", a.FormulaAudit)
	}
	if !a.Ablations.AllDiagnosticOnly || !a.Ablations.SlotSourceDominates || len(a.Ablations.Diagnostics) < 4 {
		t.Fatalf("bad ablations: %+v", a.Ablations)
	}
	if !a.TheoremTarget.ASHACoincidenceCertified || a.TheoremTarget.GeneralPQDimensionTheorem || a.TheoremTarget.CubicSlotTheoremCertified {
		t.Fatalf("bad theorem target: %+v", a.TheoremTarget)
	}
	if a.Firewalls.ClaimsGeneralPQDimensionTheorem || a.Firewalls.ClaimsCubicSlotTheorem || a.Firewalls.ClaimsFullSymbolicHitchin || a.Firewalls.ClaimsSplitG2 || a.Firewalls.ClaimsBoundaryStress || a.Firewalls.ClaimsSevenOver72 || a.Firewalls.ClaimsScalarFlavor || a.Firewalls.ClaimsPhysicalMetric || a.Firewalls.ClaimsHiggsMass || a.Firewalls.ClaimsCKMPMNS || a.Firewalls.ClaimsGaugeUnification || a.Firewalls.Verdict != StatusGate648Boundary {
		t.Fatalf("firewall breach: %+v", a.Firewalls)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2CubicSlotMultiplicityVersusNegativeSectorDimensionAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected construction failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusGate647LedgerInherited, StatusPerDirectionTraceComputed, StatusOrderedSlotContributions, StatusNegativeIndexComputed, StatusFormulaDisambiguationAudited, StatusAblativeDiagnosticsComputed, StatusMinusThreeFromCubicSlots, StatusDimMinusEqualsCubicDegree, StatusHitchinTheoremRefined, StatusNoGeneralPQDimensionTheorem, StatusNoFullSymbolicTheorem, StatusNoSplitG2, StatusNoBoundaryStress, StatusNoSevenOver72, StatusNoScalarFlavor, StatusNoPhysicalMetric, StatusNoHiggsFlavorGauge, StatusGate648Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
