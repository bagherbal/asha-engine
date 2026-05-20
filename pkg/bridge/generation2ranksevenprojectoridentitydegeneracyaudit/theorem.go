package generation2ranksevenprojectoridentitydegeneracyaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2RankSevenProjectorIdentityDegeneracyAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: AuditID, Name: "Gate 684 — Rank-Seven Projector Identity Degeneracy Audit", Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: AuditID, Name: "Gate 684 — Rank-Seven Projector Identity Degeneracy Audit", Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate683 projector response", Passed: a.Inherited.ProjectorResponseInherited && a.Inherited.K7Rank == 7 && a.Inherited.H72Dimension == 72 && a.Inherited.Gate683UsedOrdinaryTrace && a.Inherited.Gate683SignedTraceFailed && a.Inherited.PriorFirewallPreserved && a.Inherited.Verdict == StatusGate683ProjectorResponseInherited, Detail: FormatInherited(a.Inherited)},
			{Name: "audit ordinary trace rank law", Passed: a.RankLaw.DependsOnlyOnRank && !a.RankLaw.CanSelectIdentity && a.RankLaw.TraceIdentity == 72 && a.RankLaw.Verdict == StatusOrdinaryTraceRankLawAudited, Detail: FormatRankLaw(a.RankLaw)},
			{Name: "enumerate typed projector candidates", Passed: len(a.Candidates.Candidates) >= 10 && len(a.Candidates.RankSevenCandidates) >= 2 && strings.Contains(a.Candidates.Verdict, StatusTypedProjectorCandidatesEnumerated) && strings.Contains(a.Candidates.Verdict, StatusNumericalResponseByRankComputed), Detail: FormatCandidates(a.Candidates)},
			{Name: "rank seven is active response", Passed: a.Candidates.BestRank == 7 && math.Abs(a.Candidates.BestResidual) < 1e-8 && contains(a.Candidates.BestNames, "P_K7") && contains(a.Candidates.BestNames, "P_W7") && strings.Contains(a.Candidates.Verdict, StatusActiveResponseSelectsRankSeven), Detail: FormatCandidates(a.Candidates)},
			{Name: "ordinary trace cannot select projector identity", Passed: a.Degeneracy.ActiveRankSelected == 7 && a.Degeneracy.OrdinaryTraceRankOnly && !a.Degeneracy.PK7UniquelySelected && contains(a.Degeneracy.DegenerateRank7Sources, "P_K7") && contains(a.Degeneracy.DegenerateRank7Sources, "P_W7") && strings.Contains(a.Degeneracy.Verdict, StatusOrdinaryTraceCannotSelectIdentity), Detail: FormatDegeneracy(a.Degeneracy)},
			{Name: "P_K7 remains strongest typed candidate but not unique", Passed: !a.PK7Source.UniquelySelected && strings.Contains(a.PK7Source.BestTypedCandidate, "P_K7") && strings.Contains(a.PK7Source.Verdict, StatusPK7StrongestTypedRankSevenCandidate) && strings.Contains(a.PK7Source.Verdict, StatusPK7NotUniquelySelectedByTraceAlone), Detail: FormatPK7Source(a.PK7Source)},
			{Name: "record missing K7 activation theorem", Passed: strings.Contains(a.Missing.Verdict, StatusNoNativeK7ActivationTheorem) && strings.Contains(a.Missing.Verdict, StatusNoNativeProjectorIdentitySelection) && strings.Contains(a.Missing.Verdict, StatusNoNativeSevenOver72Theorem) && a.Missing.PreciseGap != "", Detail: FormatMissing(a.Missing)},
			{Name: "preserve firewalls", Passed: !a.Discipline.ClaimsK7IdentitySelectedByTrace && !a.Discipline.ClaimsNativeK7Activation && !a.Discipline.ClaimsProjectorIdentityTheorem && !a.Discipline.ClaimsNativeSevenOver72 && !a.Discipline.ClaimsBoundaryStressDerivation && !a.Discipline.ClaimsHiggsMass && !a.Discipline.ClaimsGaugeUnification && !a.Discipline.ClaimsFlavorDerivation && a.Discipline.Verdict == StatusGate684Boundary, Detail: FormatDiscipline(a.Discipline)},
		}
		notes := append(Statuses(), a.Truth)
		return theorem.Result{ID: AuditID, Name: "Gate 684 — Rank-Seven Projector Identity Degeneracy Audit", Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}

func contains(xs []string, x string) bool {
	for _, y := range xs {
		if y == x {
			return true
		}
	}
	return false
}
