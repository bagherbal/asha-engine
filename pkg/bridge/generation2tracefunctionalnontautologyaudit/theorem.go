package generation2tracefunctionalnontautologyaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2AugmentedChamberTraceResponseFunctionalNonTautologyAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 675 — AugmentedChamber Trace-Response Functional Non-Tautology Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate675 trace-functional non-tautology audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate674 trace-response candidate", Passed: a.Inherited.TraceCandidateInherited && a.Inherited.AugmentedChamberDefined && a.Inherited.RankSevenSourceAudited && a.Inherited.ScalarTraceCandidateDefined && a.Inherited.DenominatorAlternativesDone && a.Inherited.FullK7BoundaryMapFailed && a.Inherited.NoNativeTraceResponseTheorem && a.Inherited.NoNativeSevenOver72Theorem && a.Inherited.NoBoundaryStressDerivation && a.Inherited.FirewallPreserved, Detail: FormatInherited(a.Inherited)},
			{Name: "define augmented-chamber defect projector", Passed: a.Projector.TotalDimension == 72 && a.Projector.Lambda4Dimension == 70 && a.Projector.BoundaryDimension == 2 && a.Projector.RankPDefect == 7 && a.Projector.BoundaryActionRank == 0 && !a.Projector.BoundaryVectorMapNeeded && a.Projector.Verdict == StatusAugmentedChamberProjectorDefined, Detail: FormatProjector(a.Projector)},
			{Name: "compute normalized defect trace", Passed: math.Abs(a.Trace.TauDefect-7.0/72.0) < 1e-15 && strings.Contains(a.Trace.Verdict, StatusTauDefectEqualsSevenOver72), Detail: FormatTrace(a.Trace)},
			{Name: "define boundary split line and alternatives", Passed: a.BoundaryLine.ChosenLine.Name == "split line" && math.Abs(a.BoundaryLine.SSplit-a.Inherited.SSplit) < 1e-15 && len(a.BoundaryLine.Candidates) == 5 && a.BoundaryLine.Verdict == StatusBoundarySplitLineDefined, Detail: FormatBoundaryLine(a.BoundaryLine)},
			{Name: "test trace-response ansatz", Passed: math.Abs(a.Ansatz.TauDefect-7.0/72.0) < 1e-15 && math.Abs(a.Ansatz.Residual-8.52583439801e-10) < 1e-14 && a.Ansatz.RequiresScalarFunctional && !a.Ansatz.RequiresVectorBoundaryMap && strings.Contains(a.Ansatz.Verdict, StatusScalarFunctionalNotVectorMap), Detail: FormatAnsatz(a.Ansatz)},
			{Name: "audit non-tautology criteria", Passed: a.NonTautology.CertifiedCriteriaCount == 4 && a.NonTautology.RequiredCriteriaCount == 5 && !a.NonTautology.PromotableToNativeTheorem && strings.Contains(a.NonTautology.Verdict, StatusNoNativeReasonTraceActsOnSplitLine), Detail: FormatNonTautology(a.NonTautology)},
			{Name: "audit source routes", Passed: len(a.Sources) == 5 && a.Sources[0].Route == "augmented chamber trace" && a.Sources[4].Status == "missing theorem", Detail: FormatSources(a.Sources)},
			{Name: "record missing theorem targets", Passed: len(a.Missing.NativeTheoremTargets) == 3 && len(a.Missing.MissingTheorems) == 4 && len(a.Missing.AllowedSupport) == 4 && strings.Contains(a.Missing.Verdict, StatusNoNativeTraceResponseTheorem) && strings.Contains(a.Missing.Verdict, StatusNoNativeWallDistanceAirlockTheorem), Detail: FormatMissing(a.Missing)},
			{Name: "preserve firewalls", Passed: !a.Discipline.ClaimsNativeTraceResponse && !a.Discipline.ClaimsTraceActsOnSplitLine && !a.Discipline.ClaimsNativeWallDistanceAirlock && !a.Discipline.ClaimsNativeSevenOver72 && !a.Discipline.ClaimsNativeStressSplitPullback && !a.Discipline.ClaimsFullK7BoundaryMap && !a.Discipline.ClaimsBoundaryStressDerivation && !a.Discipline.ClaimsHiggsMassPrediction && !a.Discipline.ClaimsScalarStability && !a.Discipline.ClaimsGaugeUnification && !a.Discipline.ClaimsFlavorDerivation && !a.Discipline.ClaimsCKMPMNSDerivation && a.Discipline.Verdict == StatusGate675Boundary, Detail: FormatDiscipline(a.Discipline)},
		}
		notes := append(Statuses(), a.Truth)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
