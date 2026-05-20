package generation2boundarystresssplitpullbackcorrectionaudit

import (
	"fmt"
	"strings"
)

func FormatInherited(x Gate671Inheritance) string {
	return fmt.Sprintf("normal=%t bestExact=%t coordinateSealed=%t noNormal=%t no7=%t noWall=%t noBoundary=%t firewall=%t inheritedResidual=%.15g verdict=%q", x.NormalVectorInherited, x.NormalVectorBestTypedExact, x.CoordinateSealed, x.NoNativeNormalVectorTheorem, x.NoNativeSevenOver72Theorem, x.NoWallDistanceAirlockTheorem, x.NoBoundaryStressDerivation, x.FirewallPreserved, x.InheritedResidual, x.Verdict)
}

func FormatFloatSlice(xs []float64) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%.15g", x))
	}
	return "(" + strings.Join(parts, ",") + ")"
}

func FormatDecomposition(x NormalVectorDecompositionAudit) string {
	return fmt.Sprintf("original=%s base=%s split=%s weight=%.15g label=%q functional=%q passes=%t verdict=%q", FormatFloatSlice(x.OriginalNormal), FormatFloatSlice(x.BaseNormal), FormatFloatSlice(x.StressSplitNormal), x.Weight, x.DecompositionLabel, x.EquivalentFunctional, x.DecompositionPasses, x.Verdict)
}

func FormatBase(x BaseScalarFlavorClosureAudit) string {
	return fmt.Sprintf("kappaLambda=%.15g kappaE=%.15g lambda=%.15g dBase=%.15g meaning=%q verdict=%q", x.KappaLambda, x.KappaE, x.Lambda, x.DBase, x.Meaning, x.Verdict)
}

func FormatStress(x BoundaryStressSplitAudit) string {
	return fmt.Sprintf("r3Minus1=%.15g lambda=%.15g sSplit=%.15g meaning=%q verdict=%q", x.R3Minus1, x.Lambda, x.SSplit, x.Meaning, x.Verdict)
}

func FormatPullback(x PullbackAudit) string {
	return fmt.Sprintf("weight=%.15g pullback=%.15g dBase=%.15g residual=%.15g abs=%.15g ratio=%.15g weightResidual=%.15g passes=%t verdict=%q", x.Weight, x.Pullback, x.DBase, x.Residual, x.AbsResidual, x.RatioDBaseToSplit, x.WeightResidual, x.PassesBridgeWindow, x.Verdict)
}

func FormatReconstruction(x ReconstructionAudit) string {
	return fmt.Sprintf("dBaseMinusPullback=%.15g wall=%.15g equivalent=%t equation=%q verdict=%q", x.DBaseMinusPullback, x.HistoryWallBalance, x.EquivalentToGate670Normal, x.Equation, x.Verdict)
}

func FormatSource(x SourceTypeAudit) string {
	return fmt.Sprintf("dBase=%q sSplit=%q weight=%q fanoFirewall=%q support=[%s] missing=[%s] verdict=%q", x.DBaseRole, x.SSplitRole, x.WeightRole, x.FanoHitchinFirewall, strings.Join(x.CandidateSupport, "; "), strings.Join(x.RequiredMissingMaps, "; "), x.Verdict)
}

func FormatDiscipline(x VerdictDiscipline) string {
	return fmt.Sprintf("claimsStressPullback=%t claims7=%t claimsWall=%t claimsBoundary=%t claimsHiggs=%t claimsStability=%t claimsGauge=%t claimsFlavor=%t claimsCKM=%t verdict=%q", x.ClaimsNativeStressSplitPullback, x.ClaimsNativeSevenOver72, x.ClaimsWallDistanceAirlock, x.ClaimsBoundaryStressDerivation, x.ClaimsHiggsMassPrediction, x.ClaimsScalarStability, x.ClaimsGaugeUnification, x.ClaimsFlavorDerivation, x.ClaimsCKMPMNSDerivation, x.Verdict)
}
