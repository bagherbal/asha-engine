package generation2boundarystresssplitlinepullbacksourceaudit

import (
	"fmt"
	"strings"
)

func FormatInherited(x Gate672Inheritance) string {
	return fmt.Sprintf("inherited=%t base=%t split=%t noStress=%t no7=%t noWall=%t noBoundary=%t firewall=%t dBase=%.15g sSplit=%.15g residual=%.15g verdict=%q", x.InheritedStressSplitPullback, x.BaseClosureComputed, x.StressSplitComputed, x.NoNativeStressSplitTheorem, x.NoNativeSevenOver72Theorem, x.NoWallDistanceAirlockTheorem, x.NoBoundaryStressDerivation, x.FirewallPreserved, x.DBase, x.SSplit, x.Residual, x.Verdict)
}

func FormatBoundaryLine(x BoundarySplitLineAudit) string {
	return fmt.Sprintf("r3Minus1=%.15g lambda=%.15g sSplit=%.15g line=%q anti=%q verdict=%q", x.R3Minus1, x.Lambda, x.SSplit, x.LineDefinition, x.AntiAlignment, x.Verdict)
}

func FormatBaseLine(x BaseDefectLineAudit) string {
	return fmt.Sprintf("kappaLambda=%.15g kappaE=%.15g lambda=%.15g dBase=%.15g line=%q direct=%q verdict=%q", x.KappaLambda, x.KappaE, x.Lambda, x.DBase, x.LineDefinition, x.DirectClosure, x.Verdict)
}

func FormatCandidate(x CandidateWeight) string {
	return fmt.Sprintf("%s weight=%.15g pullback=%.15g residual=%.15g abs=%.15g source=%q", x.Name, x.Weight, x.Pullback, x.Residual, x.AbsResidual, x.SourceTyping)
}

func FormatCoefficient(x PullbackCoefficientAudit) string {
	parts := make([]string, 0, len(x.Candidates))
	for _, c := range x.Candidates {
		parts = append(parts, FormatCandidate(c))
	}
	return fmt.Sprintf("qPull=%.15g best=%q bestWeight=%.15g bestResidual=%.15g sevenResidual=%.15g candidates=[%s] verdict=%q", x.QPull, x.BestTypedCandidate, x.BestTypedWeight, x.BestTypedResidual, x.SevenOver72Residual, strings.Join(parts, "; "), x.Verdict)
}

func FormatSource(x LineMapSourceAudit) string {
	return fmt.Sprintf("augmented=%q k7=%q split=%q response=%q coordinate=%q support=[%s] missing=[%s] verdict=%q", x.AugmentedChamberTrace, x.K7DefectResponse, x.BoundarySplitProject, x.StressResponse, x.CoordinateArtifact, strings.Join(x.CandidateSupport, "; "), strings.Join(x.MissingTheorems, "; "), x.Verdict)
}

func FormatFirewall(x FullBoundaryMapFirewallAudit) string {
	return fmt.Sprintf("fullK7Failed=%t fanoSealed=%t linePossible=%t distinction=%q verdict=%q", x.FullK7ToBoundaryMapFailed, x.FanoHitchinRouteRemainsSealed, x.LinePullbackStillPossible, x.Distinction, x.Verdict)
}

func FormatScale(x ScaleLocalAudit) string {
	return fmt.Sprintf("lambda12Local=%t crossing=%t stationaryRejected=%t qNearOnlyAtLambda12=%t statement=%q verdict=%q", x.Lambda12Local, x.CrossingBased, x.StationarityRejected, x.QPullNearSevenOver72OnlyAtLambda12, x.ScaleStatement, x.Verdict)
}

func FormatDiscipline(x VerdictDiscipline) string {
	return fmt.Sprintf("claimsStress=%t claims7=%t claimsK7Map=%t claimsWall=%t claimsBoundary=%t claimsHiggs=%t claimsStability=%t claimsGauge=%t claimsFlavor=%t claimsCKM=%t verdict=%q", x.ClaimsNativeStressSplitPullback, x.ClaimsNativeSevenOver72, x.ClaimsFullK7BoundaryMap, x.ClaimsWallDistanceAirlock, x.ClaimsBoundaryStressDerivation, x.ClaimsHiggsMassPrediction, x.ClaimsScalarStability, x.ClaimsGaugeUnification, x.ClaimsFlavorDerivation, x.ClaimsCKMPMNSDerivation, x.Verdict)
}
