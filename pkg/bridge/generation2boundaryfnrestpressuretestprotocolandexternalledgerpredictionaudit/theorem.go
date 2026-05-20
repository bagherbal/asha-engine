package generation2boundaryfnrestpressuretestprotocolandexternalledgerpredictionaudit

import (
	"fmt"
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

const (
	theoremID   = "GATE-815-BOUNDARY-FN-RESTPRESSURE-TEST-PROTOCOL-EXTERNAL-LEDGER-PREDICTION"
	theoremName = "Gate 815 — Boundary-FN RestPressure Test Protocol and External Ledger Prediction Audit"
)

func Generation2BoundaryFNRestPressureTestProtocolAndExternalLedgerPredictionAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Notes: []string{err.Error()}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 814 boundary-FN status", Passed: a.Inheritance.Gate814Inherited && a.Inheritance.CandidateSelected && math.Abs(a.Inheritance.DeltaNBFN-0.002327375081808316) < 1e-18 && math.Abs(a.Inheritance.RelativeResidual+1.1863116229e-5) < 1e-13 && containsAll(a.Inheritance.Supports, []string{StatusReadyForTest}) && containsAll(a.Inheritance.Failures, []string{StatusNotPromoted}), Detail: FormatInheritance(a.Inheritance)},
			{Name: "freeze H_BFN coefficients before ledger testing", Passed: a.Hypothesis.Frozen && math.Abs(a.Hypothesis.C1-1.8) < 1e-15 && math.Abs(a.Hypothesis.C2-6) < 1e-15 && containsAll(a.Hypothesis.Verdicts, []string{StatusHypothesisFrozen}) && containsAll(a.Hypothesis.Supports, []string{StatusTypedCoeffSources}) && containsAll(a.Hypothesis.Failures, []string{StatusNoRetuning, StatusTypedNotNative}), Detail: FormatHypothesis(a.Hypothesis)},
			{Name: "define external trace-magnitude ledger and top/rest readout requirements", Passed: a.Ledger.Defined && len(a.Ledger.RequiredItems) >= 10 && len(a.Ledger.Readouts) >= 7 && containsAll(a.Ledger.Verdicts, []string{StatusExternalLedgerReqs, StatusTopRestReqs}) && containsAll(a.Ledger.Supports, []string{StatusExternalCanTest}) && containsAll(a.Ledger.Failures, []string{StatusNoTestNoAtoms, StatusNoAlphaBetaQNoTop, StatusExternalNotNative}), Detail: FormatLedger(a.Ledger)},
			{Name: "define aggregate Delta_N and c2_ext tests", Passed: a.Aggregate.Defined && math.Abs(a.Aggregate.C2Observed-5.82999157225) < 1e-9 && strings.Contains(strings.Join(a.Aggregate.Diagnostics, " "), "c2_ext") && containsAll(a.Aggregate.Supports, []string{StatusC2NearSixPrimary}) && containsAll(a.Aggregate.Failures, []string{StatusAggregateNoSectors, StatusC2NoAbsorbRetune}), Detail: FormatAggregate(a.Aggregate)},
			{Name: "compute positive top/rest alpha band", Passed: a.TopRest.Defined && math.Abs(a.TopRest.AlphaMin-0.000387820644542014) < 1e-15 && math.Abs(a.TopRest.AlphaMax-0.000388046602361924) < 1e-15 && math.Abs(a.TopRest.AlphaMinOverS-0.300067468178) < 1e-10 && math.Abs(a.TopRest.AlphaMaxOverS-0.30024229794) < 1e-9 && containsAll(a.TopRest.Supports, []string{StatusAlphaNearThreeTenths}) && containsAll(a.TopRest.Failures, []string{StatusAlphaBandNoSectors, StatusAlphaBandNotOperator}), Detail: FormatTopRest(a.TopRest)},
			{Name: "define rest concentration and sector pressure diagnostics", Passed: a.RestQ.Defined && a.Sector.Defined && len(a.RestQ.Interpretations) == 3 && len(a.Sector.Predictions) >= 5 && containsAll(a.RestQ.Supports, []string{StatusQClassifies}) && containsAll(a.Sector.Supports, []string{StatusSmallPositiveRest}) && containsAll(a.Sector.Failures, []string{StatusNoSectorAssignment, StatusGJKoideSecondary}), Detail: strings.Join(a.RestQ.Interpretations, " | ") + " :: " + strings.Join(a.Sector.Predictions, " | ")},
			{Name: "define boundary-FN spurion diagnostic", Passed: a.Spurion.Defined && math.Abs(a.Spurion.EpsilonBFN-0.2196426096400638) < 1e-15 && math.Abs(a.Spurion.ResidualEpsilon+6.51406623e-7) < 1e-14 && containsAll(a.Spurion.Supports, []string{StatusEpsilonSharp}) && containsAll(a.Spurion.Failures, []string{StatusEpsilonNotNativeFN, StatusEpsilonNoAtoms}), Detail: strings.Join([]string{StatusSpurionTest, "epsilon_BFN=" + fmtFloat(a.Spurion.EpsilonBFN), "R_epsilon=" + fmtFloat(a.Spurion.ResidualEpsilon)}, " ")},
			{Name: "define noncircular protocol", Passed: a.Protocol.Defined && len(a.Protocol.Steps) == 7 && len(a.Protocol.Forbidden) == 5 && containsAll(a.Protocol.Supports, []string{StatusNonCircularlyTestable}) && containsAll(a.Protocol.Failures, []string{StatusCoeffRetuningInvalid, StatusTopSelectorNoForce, StatusNoHiggsDataForYukawa}), Detail: strings.Join(a.Protocol.Steps, " | ")},
			{Name: "define pass/failure criteria", Passed: a.Failure.Defined && a.Pass.Defined && len(a.Failure.Items) == 6 && len(a.Pass.Items) == 6 && containsAll(a.Failure.Supports, []string{StatusClearFalsification}) && containsAll(a.Pass.Supports, []string{StatusCanUpgradeR2R3}) && containsAll(a.Pass.Failures, []string{StatusExternalR3NotNativeR4}), Detail: strings.Join(a.Failure.Items, " | ") + " :: " + strings.Join(a.Pass.Items, " | ")},
			{Name: "record C_Yukawa and C_Higgs candidate impact without update", Passed: a.Impact.Recorded && math.Abs(a.Impact.CYukawaBFN-0.9992248096922658) < 1e-15 && math.Abs(a.Impact.CHiggsBFN-1.0372205108665146) < 1e-15 && math.Abs(a.Impact.OfficialCYukawa-CYukawa) < 1e-15 && containsAll(a.Impact.Supports, []string{StatusCanReduceSealDependence}) && containsAll(a.Impact.Failures, []string{StatusNoCYukawaUpdate, StatusCHiggsLevelB}), Detail: FormatImpact(a.Impact)},
			{Name: "classify pattern lanes, outcome, and branch", Passed: a.Patterns.Classified && len(a.Patterns.Lanes) == 5 && a.Outcome.Recorded && len(a.Outcome.Items) == 6 && a.Branch.Recorded && strings.Contains(a.Branch.NativeNext, "Gate 816") && containsAll(a.Patterns.Supports, []string{StatusFNInterpretsRest, StatusGJTestsDownLepton}) && containsAll(a.Patterns.Failures, []string{StatusPatternsNoMap}) && containsAll(a.Branch.Supports, []string{StatusNativeNextCoeffPrior, StatusEmpiricalNextFrozenTest}), Detail: strings.Join(a.Patterns.Lanes, " | ") + " -> " + a.Branch.NativeNext},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Enforced && a.Firewalls.NoPromotion && a.Firewalls.NoRetuning && a.Firewalls.NoSilentTopChoice && a.Firewalls.NoHiggsData && a.Firewalls.NoPatternSource && a.Firewalls.NoLedgerUpdate && a.Firewalls.NoPoleMass && a.Firewalls.Verdict == StatusFirewallGate815, Detail: a.Firewalls.Verdict},
		}
		ok := true
		for _, c := range checks {
			if !c.Passed {
				ok = false
				break
			}
		}
		status := theorem.BridgeRequired
		if !ok {
			status = theorem.FailedRoute
		}
		notes := []string{a.Truth, FormatInheritance(a.Inheritance), FormatHypothesis(a.Hypothesis), FormatLedger(a.Ledger), FormatAggregate(a.Aggregate), FormatTopRest(a.TopRest), strings.Join(a.RestQ.Interpretations, " | "), strings.Join(a.Sector.Predictions, " | "), FormatImpact(a.Impact), strings.Join(a.Outcome.Items, " | "), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}

func fmtFloat(x float64) string {
	return strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.16g", x), "0"), ".")
}
