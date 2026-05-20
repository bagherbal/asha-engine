package generation2rankthreetopcolorblockandrestpressureoperatorsourceaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate808TopColorBlockAndRestPressure(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.TopColor.Defined || a.TopColor.Name != "RankThreeTopColorBlockSeal" {
		t.Fatalf("bad top seal: %s", FormatTop(a.TopColor))
	}
	neffTop, err := TopColorNEff(0.7)
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(neffTop-3) > 1e-12 || math.Abs(a.TopColor.NEffTop-3) > 1e-12 {
		t.Fatalf("bad top-color limit: %.17g %.17g", neffTop, a.TopColor.NEffTop)
	}
	neff, err := RestPressureNEff(0.01, 0.005)
	if err != nil {
		t.Fatal(err)
	}
	want := 3 * (1.01 * 1.01) / 1.005
	if math.Abs(neff-want) > 1e-15 {
		t.Fatalf("bad rest formula: got %.17g want %.17g", neff, want)
	}
	delta, err := RestPressureDelta(0.01, 0.005)
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(delta-(neff-3)) > 1e-15 {
		t.Fatalf("bad rest delta: %.17g %.17g", delta, neff-3)
	}
	if !containsAll(a.RestPressure.Supports, []string{StatusRestPressureAboveTop, StatusRestPressureDilutesCYukawa}) {
		t.Fatalf("missing rest supports: %s", FormatRest(a.RestPressure))
	}
}

func TestGate808AggregatePositivityCorridorAndConcentration(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	lower, upper, err := TopDominantPositivityCorridor(AInherited, BInherited)
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(lower-0.9471023226011707) > 1e-15 || math.Abs(upper-0.9471025365183062) > 1e-15 {
		t.Fatalf("bad corridor: %.17g %.17g", lower, upper)
	}
	if math.Abs(a.Corridor.AlphaAtUpper-0.00038781604472679744) > 1e-12 || math.Abs(a.Corridor.BetaAtLower-4.5172977535955994e-7) > 1e-16 {
		t.Fatalf("bad corridor values: %s", FormatCorridor(a.Corridor))
	}
	beta, err := RestConcentrationBeta(0.01, 1)
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(beta-0.0003) > 1e-16 {
		t.Fatalf("bad beta from q_rest: %.17g", beta)
	}
	if !containsAll(a.Concentration.Failures, []string{StatusNoRestAtomCount, StatusNoQRestFromAggregate}) {
		t.Fatalf("missing concentration failures: %s", FormatConcentration(a.Concentration))
	}
}

func TestGate808SourceAuditsAndCHiggsImpact(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !containsAll(a.SectorCandidates.Failures, []string{StatusNoSectorAssignment, StatusNeutrinoImplicit, StatusScaleSchemeUntyped}) {
		t.Fatalf("bad sector candidates: %s", FormatAudit(a.SectorCandidates))
	}
	if !containsAll(a.PatternDiagnostics.Failures, []string{StatusKoideNotNEff, StatusFNNotNativeRest, StatusGJNotLowScaleTop}) {
		t.Fatalf("bad pattern firewall: %s", FormatAudit(a.PatternDiagnostics))
	}
	if !containsAll(a.D4Firewall.Failures, []string{StatusNEffNotD4, StatusNoTrialityTraceReadout, StatusNoTrialityRestPressure, StatusNoTrialityRealDescent}) {
		t.Fatalf("bad D4 firewall: %s", FormatAudit(a.D4Firewall))
	}
	if !containsAll(a.FiniteTriple.Supports, []string{StatusFiniteTripleColorShape}) || !containsAll(a.FiniteTriple.Failures, []string{StatusFSTNoTopEigenvalue, StatusFSTNoRestOperator}) {
		t.Fatalf("bad finite triple: %s", FormatAudit(a.FiniteTriple))
	}
	if !containsAll(a.ExternalLedger.Supports, []string{StatusExternalCanTest}) || !containsAll(a.ExternalLedger.Failures, []string{StatusExternalNotNative}) {
		t.Fatalf("bad external ledger: %s", FormatAudit(a.ExternalLedger))
	}
	if !containsAll(a.K7Projective.Failures, []string{StatusK7NotTopBlock, StatusProjectiveNotRest}) {
		t.Fatalf("bad K7 audit: %s", FormatAudit(a.K7Projective))
	}
	if !containsAll(a.ComplexD4.Failures, []string{StatusTD4NotTraceMagnitude, StatusTD4NotTopDominance, StatusTD4NotRestPressure}) {
		t.Fatalf("bad complex D4 audit: %s", FormatAudit(a.ComplexD4))
	}
	if math.Abs(a.CHiggs.DeltaCHiggs-0.0008046575187645733) > 1e-15 || math.Abs(a.CHiggs.TreeShift-TreeProxyShift) > 1e-10 {
		t.Fatalf("bad C_Higgs impact: %s", FormatCHiggs(a.CHiggs))
	}
}

func TestGate808HierarchyOutcomeBranchAndTheorem(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Hierarchy.Selected || !strings.Contains(a.Hierarchy.Name, "HierarchyBreakingOperatorSeal") || !containsAll(a.Hierarchy.Failures, []string{StatusNoHierarchyBreakingOperator, StatusNoNativeTopDominance, StatusNoLightSuppression, StatusNoNativeRestSource}) {
		t.Fatalf("bad hierarchy obstruction: %s", FormatHierarchy(a.Hierarchy))
	}
	if !a.Outcome.Recorded || !containsAll(a.Outcome.Items, []string{"exact N_eff=3", "positive rest spectral pressure", "C_Higgs remains Level B"}) {
		t.Fatalf("bad outcome: %s", FormatOutcome(a.Outcome))
	}
	if !strings.Contains(a.Branch.Next, "Gate 809") || !strings.Contains(a.Branch.Next, "HierarchyBreakingOperatorSeal") {
		t.Fatalf("bad branch: %+v", a.Branch)
	}
	res := Generation2RankThreeTopColorBlockAndRestPressureOperatorSourceAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status %s", want)
		}
	}
}
