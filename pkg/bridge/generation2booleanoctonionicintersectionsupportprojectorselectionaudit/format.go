package generation2booleanoctonionicintersectionsupportprojectorselectionaudit

import (
	"fmt"
	"strings"
)

func FormatInherited(x Gate684Inheritance) string {
	return fmt.Sprintf("rankDegeneracy=%t traceRankOnly=%t rank7=%t traceNoIdentity=%t dBase=%.15g sSplit=%.15g residual=%.15g h72=%d k7=%d priorFirewall=%t verdict=%q", x.RankDegeneracyInherited, x.OrdinaryTraceRankOnly, x.RankSevenSelected, x.TraceCannotSelectIdentity, x.DBase, x.SSplit, x.TraceResidual, x.H72Dimension, x.K7Dimension, x.PriorFirewallPreserved, x.Verdict)
}

func FormatSupport(x NativeSupportConstraintAudit) string {
	return fmt.Sprintf("constraints=[%s] boolean=%q octonionic=%q intersection=%q ranks=(PB:%d PG:%d K7:%d) imagePB=%t imagePG=%t imageIntersection=%t verdict=%q", strings.Join(x.ProjectorConstraints, ", "), x.BooleanSupport, x.OctonionicSupport, x.Intersection, x.PBRank, x.PGRank, x.IntersectionDimension, x.ImpliesImageInPB, x.ImpliesImageInPG, x.ImpliesImageInIntersection, x.Verdict)
}

func FormatChamber(x ChamberDimensionAudit) string {
	return fmt.Sprintf("lambda4=%d boundary=%d h72=%d rankPB=%d rankPG=%d intersection=%d uPlusV=%d w7=%d grassmannDegenerate=%t ledgerOK=%t verdict=%q", x.Lambda4Dimension, x.BoundaryDimension, x.H72Dimension, x.PBRank, x.PGRank, x.IntersectionDim, x.UPlusVDim, x.OrthogonalW7Dim, x.GrassmannDegeneracy, x.DimensionalLedgerOK, x.Verdict)
}

func FormatSelection(x SelectionProofAudit) string {
	return fmt.Sprintf("assumptions=[%s] subsetK7=%t rankEqualsK7=%t imageEqualsK7=%t symmetric=%t orthogonalUnique=%t selected=%q verdict=%q", strings.Join(x.Assumptions, "; "), x.ImageSubsetK7, x.RankEqualsIntersectionDim, x.ImageEqualsK7, x.SymmetricProjectorRequired, x.OrthogonalProjectorUnique, x.SelectedProjector, x.Verdict)
}

func FormatCandidates(x CandidateComparisonAudit) string {
	parts := make([]string, 0, len(x.Candidates))
	for _, c := range x.Candidates {
		parts = append(parts, fmt.Sprintf("%s rank=%d inPB=%d inPG=%d inK7=%d inW7=%d inBoundary=%d PBP=P:%t PGP=P:%t pass:%t selected:%t meaning=%s reason=%s", c.Name, c.Rank, c.InBooleanRank, c.InOctonionicRank, c.InIntersectionRank, c.InOrthogonalW7Rank, c.InBoundaryRank, c.PBPEqualsP, c.PGPEqualsP, c.PassesNativeSupport, c.SelectedAsPK7, c.TypedMeaning, c.RejectionReason))
	}
	return fmt.Sprintf("passing=[%s] rejectedRank7=[%s] pk7Passes=%t w7Rejected=%t arbitraryRejected=%t allPassingPK7=%t candidates=[%s] verdict=%q", strings.Join(x.PassingCandidates, ","), strings.Join(x.RejectedRankSeven, ","), x.PK7Passes, x.W7Rejected, x.ArbitraryRejected, x.AllPassingArePK7, strings.Join(parts, " | "), x.Verdict)
}

func FormatResponseUpdate(x ResponseUpdateAudit) string {
	return fmt.Sprintf("rankOnly=%q supportSelected=%q reason=%q residual=%.15g resolved=%t conditional=%t activationUnproved=%t verdict=%q", x.RankOnlyResponse, x.SupportSelectedResponse, x.SelectionReason, x.TraceResidual, x.DegeneracyResolved, x.SelectionIsConditional, x.ActivationStillUnproved, x.Verdict)
}

func FormatMissing(x MissingTheoremAudit) string {
	return fmt.Sprintf("missing=[%s] precise=%q verdict=%q", strings.Join(x.Missing, "; "), x.PreciseGap, x.Verdict)
}

func FormatDiscipline(x VerdictDiscipline) string {
	return fmt.Sprintf("claimsTraceSelectsPK7=%t claimsSSplitActivatesSupport=%t claimsProjectorActivation=%t claims7=%t claimsBoundary=%t claimsScalarRG=%t claimsHiggs=%t claimsGauge=%t claimsFlavor=%t verdict=%q", x.ClaimsTraceSelectsPK7, x.ClaimsSSplitActivatesSupport, x.ClaimsProjectorActivation, x.ClaimsNativeSevenOver72, x.ClaimsBoundaryStressDerivation, x.ClaimsScalarRGMatching, x.ClaimsHiggsMass, x.ClaimsGaugeUnification, x.ClaimsFlavorDerivation, x.Verdict)
}
