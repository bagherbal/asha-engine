package generation2ranksevenprojectoridentitydegeneracyaudit

import (
	"fmt"
	"strings"
)

func FormatInherited(x Gate683Inheritance) string {
	return fmt.Sprintf("projectorInherited=%t dBase=%.15g sSplit=%.15g h72=%d k7=%d gate683Residual=%.15g ordinary=%t signedFailed=%t firewall=%t verdict=%q", x.ProjectorResponseInherited, x.DBase, x.SSplit, x.H72Dimension, x.K7Rank, x.Gate683Residual, x.Gate683UsedOrdinaryTrace, x.Gate683SignedTraceFailed, x.PriorFirewallPreserved, x.Verdict)
}
func FormatRankLaw(x RankLawAudit) string {
	return fmt.Sprintf("formula=%q rankOnly=%t canSelectIdentity=%t TrI=%d verdict=%q", x.Formula, x.DependsOnlyOnRank, x.CanSelectIdentity, x.TraceIdentity, x.Verdict)
}
func FormatCandidates(x ProjectorCandidateAudit) string {
	parts := make([]string, 0, len(x.Candidates))
	for _, c := range x.Candidates {
		parts = append(parts, fmt.Sprintf("%s rank=%d coeff=%.15g pred=%.15g residual=%.15g source=%s typed=%s", c.Name, c.Rank, c.Coefficient, c.Prediction, c.Residual, c.Source, c.TypedStatus))
	}
	return fmt.Sprintf("bestRank=%d bestResidual=%.15g bestNames=[%s] rank7=[%s] candidates=[%s] verdict=%q", x.BestRank, x.BestResidual, strings.Join(x.BestNames, ","), strings.Join(x.RankSevenCandidates, ","), strings.Join(parts, " | "), x.Verdict)
}
func FormatDegeneracy(x RankDegeneracyAudit) string {
	return fmt.Sprintf("activeRank=%d rank7Residual=%.15g rank7Names=[%s] rankOnly=%t pk7Unique=%t degenerate=[%s] verdict=%q", x.ActiveRankSelected, x.RankSevenResidual, strings.Join(x.RankSevenNames, ","), x.OrdinaryTraceRankOnly, x.PK7UniquelySelected, strings.Join(x.DegenerateRank7Sources, ","), x.Verdict)
}
func FormatPK7Source(x PK7SourceAudit) string {
	return fmt.Sprintf("reasons=[%s] alternative=%q best=%q unique=%t verdict=%q", strings.Join(x.Reasons, "; "), x.AlternativeWarning, x.BestTypedCandidate, x.UniquelySelected, x.Verdict)
}
func FormatMissing(x MissingTheoremAudit) string {
	return fmt.Sprintf("missing=[%s] precise=%q verdict=%q", strings.Join(x.Missing, "; "), x.PreciseGap, x.Verdict)
}
func FormatDiscipline(x VerdictDiscipline) string {
	return fmt.Sprintf("claimsK7TraceIdentity=%t claimsActivation=%t claimsIdentityTheorem=%t claims7=%t claimsBoundary=%t claimsHiggs=%t claimsGauge=%t claimsFlavor=%t verdict=%q", x.ClaimsK7IdentitySelectedByTrace, x.ClaimsNativeK7Activation, x.ClaimsProjectorIdentityTheorem, x.ClaimsNativeSevenOver72, x.ClaimsBoundaryStressDerivation, x.ClaimsHiggsMass, x.ClaimsGaugeUnification, x.ClaimsFlavorDerivation, x.Verdict)
}
