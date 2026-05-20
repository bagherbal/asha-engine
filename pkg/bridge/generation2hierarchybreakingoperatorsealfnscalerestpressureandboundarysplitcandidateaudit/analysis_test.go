package generation2hierarchybreakingoperatorsealfnscalerestpressureandboundarysplitcandidateaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate809NumericalBlinkAndBoundarySplit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Blink.Computed || math.Abs(a.Blink.EpsilonN-0.21964195823344188) > 2e-15 {
		t.Fatalf("bad epsilon_N: %s", FormatBlink(a.Blink))
	}
	if got, err := FourthRootRestPressure(DeltaN); err != nil || math.Abs(got-a.Blink.EpsilonN) > 1e-15 {
		t.Fatalf("fourth root failed: %.17g %v", got, err)
	}
	if !containsAll(a.Blink.Supports, []string{StatusDeltaNHasFNScale, StatusEpsilonNStrong}) || !containsAll(a.Blink.Failures, []string{StatusEpsilonNotNative, StatusEpsilon4NotTheorem, StatusNoFNChargeOperator}) {
		t.Fatalf("bad blink statuses: %s", FormatBlink(a.Blink))
	}
	if math.Abs(a.Boundary.DeltaOverS-1.8007325638446063) > 5e-14 || math.Abs(a.Boundary.NineFifthsResid-9.467983454135818e-7) > 1e-15 {
		t.Fatalf("bad boundary resonance: %s", FormatBoundary(a.Boundary))
	}
	if math.Abs(AlphaFromDeltaSmallRest(DeltaN)-0.0003878912453691245) > 1e-15 || math.Abs(AlphaFromBoundarySplit(SBoundary)-0.00038773344564488885) > 1e-15 {
		t.Fatalf("bad alpha approximations")
	}
	if !containsAll(a.Boundary.Failures, []string{StatusNoNineFifthsSource, StatusNoThreeTenthsSource, StatusNoBoundaryYukawaMap, StatusBoundaryNotTheorem}) {
		t.Fatalf("missing boundary failures: %s", FormatBoundary(a.Boundary))
	}
}

func TestGate809BoundaryFNSynthesisAndCandidates(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	epsB, err := BoundaryFNScale(SBoundary)
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(epsB-0.21961961644976352) > 1e-15 || math.Abs(a.BoundaryFN.EpsilonB-epsB) > 1e-15 {
		t.Fatalf("bad epsilon_B: %s", FormatBoundaryFN(a.BoundaryFN))
	}
	if math.Abs(a.BoundaryFN.Difference-2.234178367835349e-05) > 2e-15 {
		t.Fatalf("bad epsilon difference: %s", FormatBoundaryFN(a.BoundaryFN))
	}
	if !containsAll(a.BoundaryFN.Supports, []string{StatusBoundaryMaySourceFN, StatusEpsilonClose}) || !containsAll(a.BoundaryFN.Failures, []string{StatusBoundaryFNNotExact, StatusNoBoundaryFNCoeff, StatusNoEpsilonBSpurion, StatusNoRestReadoutEpsilonB}) {
		t.Fatalf("bad boundary-FN statuses: %s", FormatBoundaryFN(a.BoundaryFN))
	}
	if !a.FN.Defined || !containsAll(a.FN.Supports, []string{StatusFNCompatible, StatusEpsilonFourCandidate}) || !containsAll(a.FN.Failures, []string{StatusFNNotNativeNoCharge, StatusFNEpsilonNoSilentFit, StatusFNNoSectorAssignment, StatusFNNoTopDominance}) {
		t.Fatalf("bad FN candidate: %s", FormatFN(a.FN))
	}
}

func TestGate809SourceAuditsRankingAndCHiggsFirewall(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !containsAll(a.Projective.Supports, []string{StatusProjectiveResonance, StatusK743Candidate}) || !containsAll(a.Projective.Failures, []string{StatusProjectiveNoEigenvalue, StatusProjectiveNotTopBlock, StatusK7NotTraceMagnitude, StatusNoProjectiveHFMap}) {
		t.Fatalf("bad projective audit: %s", FormatAudit(a.Projective))
	}
	if !containsAll(a.GeorgiJarlskog.Supports, []string{StatusGJRestStructure}) || !containsAll(a.GeorgiJarlskog.Failures, []string{StatusGJNotLowScale, StatusGJNotTopColorThree, StatusSingleScaleNoGJ}) {
		t.Fatalf("bad GJ audit: %s", FormatAudit(a.GeorgiJarlskog))
	}
	if !containsAll(a.Koide.Failures, []string{StatusKoideNotTop, StatusKoideNotRest, StatusKoideNotNative}) {
		t.Fatalf("bad Koide firewall: %s", FormatAudit(a.Koide))
	}
	if !containsAll(a.D4.Failures, []string{StatusD4NotHierarchy, StatusD4NotTop, StatusD4NotRest, StatusNoD4TraceReadout}) {
		t.Fatalf("bad D4 firewall: %s", FormatAudit(a.D4))
	}
	if len(a.Ranking.Ranks) != 7 || !containsAll(a.Ranking.Ranks, []string{"Boundary-FN", "External Yukawa", "D4/triality"}) || !containsAll(a.Ranking.Supports, []string{StatusBoundaryFNSharpest, StatusFNEpsilonSerious, StatusProjectiveSearch}) {
		t.Fatalf("bad ranking: %s", FormatRanking(a.Ranking))
	}
	if !a.CHiggs.Preserved || !strings.Contains(a.CHiggs.CandidateRewrite, "(9/5)s") || math.Abs(ApproxCYukawaFromBoundarySplit(SBoundary)-0.9992251339916449) > 1e-15 {
		t.Fatalf("bad C_Higgs firewall: %s", FormatCHiggs(a.CHiggs))
	}
	if !containsAll(a.CHiggs.Failures, []string{StatusNoCYukawaUpdate, StatusBoundaryRewriteNotCert, StatusCHiggsLevelB}) {
		t.Fatalf("missing C_Higgs failures: %s", FormatCHiggs(a.CHiggs))
	}
}

func TestGate809OutcomeBranchAndTheorem(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Hierarchy.Defined || !containsAll(a.Hierarchy.Failures, []string{StatusNoCurrentHierarchy, StatusNoCurrentTopDominance, StatusNoCurrentRestSuppress}) {
		t.Fatalf("bad hierarchy seal: %s", FormatHierarchy(a.Hierarchy))
	}
	if !a.Outcome.Recorded || !containsAll(a.Outcome.Items, []string{"epsilon_N", "(9/5)s", "Boundary-FN"}) {
		t.Fatalf("bad outcome: %s", FormatOutcome(a.Outcome))
	}
	if !strings.Contains(a.Branch.Next, "Gate 810") || !strings.Contains(a.Branch.Next, "Boundary-FN") || !containsAll(a.Branch.Supports, []string{StatusNextBoundaryFN}) {
		t.Fatalf("bad branch: %+v", a.Branch)
	}
	res := Generation2HierarchyBreakingOperatorSealFNScaleRestPressureAndBoundarySplitCandidateAuditTheorem().Verify()
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
